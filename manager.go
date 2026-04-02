package gotdbot

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/AshokShau/gotdbot/internal/tdjson"
)

var (
	defaultManager *ClientManager
	managerOnce    sync.Once
)

// GetDefaultManager returns the global singleton ClientManager.
func GetDefaultManager(libPath string) *ClientManager {
	managerOnce.Do(func() {
		defaultManager = NewClientManager(libPath)
		defaultManager.Start()
	})
	return defaultManager
}

type ClientManager struct {
	LibraryPath string

	clients   map[int]*Client
	mu        sync.RWMutex
	stop      chan struct{}
	wg        sync.WaitGroup
	once      sync.Once
	closeOnce sync.Once
}

func NewClientManager(LibraryPath string) *ClientManager {
	return &ClientManager{
		LibraryPath: LibraryPath,
		clients:     make(map[int]*Client),
		stop:        make(chan struct{}),
	}
}

func (m *ClientManager) AddClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[c.clientID] = c
	c.manager = m
	m.Start()
}

func (m *ClientManager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, c.clientID)
	c.manager = nil
}

func (m *ClientManager) GetClient(clientID int) *Client {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.clients[clientID]
}

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
	if config.LibraryPath == "" {
		config.LibraryPath = m.LibraryPath
	}

	client, err := NewClient(apiID, apiHash, tokenOrPhone, config)
	if err != nil {
		return nil, err
	}

	m.AddClient(client)
	if err := client.Start(); err != nil {
		client.Close()
		m.RemoveClient(client)
		return nil, err
	}

	return client, nil
}

func (m *ClientManager) Start() {
	m.once.Do(func() {
		_ = tdjson.Init(m.LibraryPath)
		m.wg.Add(1)
		go m.receiver()
	})
}

func (m *ClientManager) Stop() {
	m.closeOnce.Do(func() {
		m.mu.RLock()
		for _, c := range m.clients {
			c.Close()
		}
		m.mu.RUnlock()
		close(m.stop)
		m.wg.Wait()
	})
}

// Idle blocks the current goroutine until a SIGINT or SIGTERM signal is received.
func (m *ClientManager) Idle() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
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
					continue
				}
				if obj == nil {
					continue
				}

				m.mu.RLock()
				c, ok := m.clients[clientId]
				m.mu.RUnlock()

				if ok {
					if extra != "" {
						if ch, ok := c.pendingRequests.Load(extra); ok {
							ch.(chan TlObject) <- obj
							c.pendingRequests.Delete(extra)
							continue
						}
					}
					c.updates <- obj
				}
			}
		}
	}
}
