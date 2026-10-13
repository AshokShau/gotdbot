package gotdbot

import (
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/AshokShau/gotdbot/internal/tdjson"
)

var (
	defaultManager  *ClientManager
	managerOnce     sync.Once
	receiverOwnerMu sync.Mutex
	receiverOwner   *ClientManager
)

// GetDefaultManager returns the global singleton ClientManager.
func GetDefaultManager(libPath string) *ClientManager {
	receiverOwnerMu.Lock()
	owner := receiverOwner
	receiverOwnerMu.Unlock()
	if owner != nil {
		return owner
	}
	managerOnce.Do(func() {
		defaultManager = NewClientManager(libPath)
	})
	return defaultManager
}

// ClientManager routes TDLib responses and updates to multiple account clients.
type ClientManager struct {
	LibraryPath string

	clients   map[int]*Client
	mu        sync.RWMutex
	stop      chan struct{}
	wg        sync.WaitGroup
	once      sync.Once
	closeOnce sync.Once
	startErr  error
	stopping  atomic.Bool
}

// NewClientManager creates a manager using the specified TDLib shared library.
func NewClientManager(LibraryPath string) *ClientManager {
	return &ClientManager{
		LibraryPath: LibraryPath,
		clients:     make(map[int]*Client),
		stop:        make(chan struct{}),
	}
}

// AddClient attaches a client to this manager and starts its shared receiver.
// The client must not belong to another manager.
func (m *ClientManager) AddClient(c *Client) error {
	if c == nil {
		return ErrClientClosed
	}
	if err := m.Start(); err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.stopping.Load() {
		return ErrClientClosed
	}
	if previous := c.manager.Load(); previous != nil && previous != m {
		return errors.New("client belongs to another manager")
	}
	m.clients[c.clientID] = c
	c.manager.Store(m)
	return nil
}

// RemoveClient detaches a client without closing its TDLib instance.
func (m *ClientManager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, c.clientID)
	c.manager.CompareAndSwap(m, nil)
}

// GetClient returns the registered client with this TDLib identifier, or nil.
func (m *ClientManager) GetClient(clientID int) *Client {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.clients[clientID]
}

// GetClients returns a snapshot of the registered clients in unspecified order.
func (m *ClientManager) GetClients() []*Client {
	m.mu.RLock()
	defer m.mu.RUnlock()
	clients := make([]*Client, 0, len(m.clients))
	for _, c := range m.clients {
		clients = append(clients, c)
	}
	return clients
}

// RegisterClient creates a new client, adds it to the manager, and starts it.
func (m *ClientManager) RegisterClient(apiID int32, apiHash, tokenOrPhone string, config *ClientOpts) (*Client, error) {
	if config == nil {
		config = DefaultClientConfig()
	}
	copyConfig := *config
	config = &copyConfig
	if config.LibraryPath == "" {
		config.LibraryPath = m.LibraryPath
	}

	client, err := NewClient(apiID, apiHash, tokenOrPhone, config)
	if err != nil {
		return nil, err
	}

	if err := m.AddClient(client); err != nil {
		client.Close()
		return nil, err
	}
	if err := client.Start(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}

// Start loads TDLib and starts the shared receiver. Repeated calls are safe.
func (m *ClientManager) Start() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.stopping.Load() {
		return ErrClientClosed
	}
	m.once.Do(func() {
		receiverOwnerMu.Lock()
		defer receiverOwnerMu.Unlock()
		if receiverOwner != nil && receiverOwner != m {
			m.startErr = errors.New("TDLib already has an active receiver; use the existing ClientManager")
			return
		}
		if err := tdjson.Init(m.LibraryPath); err != nil {
			m.startErr = err
			return
		}
		receiverOwner = m
		m.wg.Add(1)
		go m.receiver()
	})
	return m.startErr
}

// Stop closes registered clients and stops the shared receiver.
func (m *ClientManager) Stop() {
	m.closeOnce.Do(func() {
		m.mu.Lock()
		m.stopping.Store(true)
		m.mu.Unlock()
		for _, c := range m.GetClients() {
			c.Close()
		}
		close(m.stop)
		m.wg.Wait()
		receiverOwnerMu.Lock()
		if receiverOwner == m {
			receiverOwner = nil
		}
		receiverOwnerMu.Unlock()
	})
}

// Idle blocks the current goroutine until a SIGINT or SIGTERM signal is received.
func (m *ClientManager) Idle() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)
	select {
	case <-sig:
	case <-m.stop:
	}
	m.Stop()
}

func (m *ClientManager) receiver() {
	defer m.wg.Done()
	for {
		select {
		case <-m.stop:
			return
		default:
			res := tdjson.Receive(0.1)
			if res != "" {
				data := []byte(res)
				obj, clientId, extra, err := UnmarshalWithClient(data)
				if err != nil {
					slog.Error("Cannot decode TDLib response; check the library/schema version", "error", err)
					continue
				}
				if obj == nil {
					continue
				}

				m.mu.RLock()
				c, ok := m.clients[clientId]
				m.mu.RUnlock()

				if ok {
					c.acceptResponse(obj, extra, m.stop)
				}
			}
		}
	}
}
