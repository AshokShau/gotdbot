package gotdbot

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
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

func NewClientManager(LibraryPath string) *ClientManager {
	return &ClientManager{
		LibraryPath: LibraryPath,
		clients:     make(map[int]*Client),
		stop:        make(chan struct{}),
	}
}

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

func (m *ClientManager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, c.clientID)
	c.manager.CompareAndSwap(m, nil)
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

// Protocol completion must not depend on an available application handler.
func (c *Client) acceptResponse(obj TlObject, extra string, managerStop <-chan struct{}) {
	if obj == nil {
		return
	}
	if extra != "" {
		c.deliverResponse(extra, obj)
		return
	}
	switch u := obj.(type) {
	case *UpdateMessageSendSucceeded:
		if u.Message != nil {
			c.completeMessage(u.Message.ChatId, u.OldMessageId, u)
		}
	case *UpdateMessageSendFailed:
		if u.Message != nil {
			c.completeMessage(u.Message.ChatId, u.OldMessageId, u)
		}
	case *UpdateAuthorizationState:
		if _, ok := u.AuthorizationState.(*AuthorizationStateClosed); ok {
			c.closing.Store(true)
			if !c.isAuthorized.Load() {
				c.sendAuthError(ErrClientClosed)
			}
			c.closedOnce.Do(func() { close(c.closed) })
		}
	}

	if !strings.HasPrefix(obj.GetType(), "update") {
		if err, ok := obj.(*Error); ok {
			c.Logger.Error("Uncorrelated TDLib error", "error", err)
		}
		return
	}
	select {
	case <-c.stop:
		return
	case <-managerStop:
		return
	default:
	}
	if err := c.updates.push(obj); err != nil && !errors.Is(err, ErrClientClosed) {
		c.Logger.Error("Cannot queue TDLib update", "type", obj.GetType(), "error", err)
	}
}

// updateQueue keeps a bounded number of decoded updates in memory.
type updateQueue struct {
	mu     sync.Mutex
	memory chan TlObject
	ready  chan struct{}
	closed bool

	dir                     string
	reader, writer          *os.File
	head, tail              uint64
	readOffset, writeOffset int64
	spilled                 int
}

const updateSegmentSize = 8 << 20

func newUpdateQueue(size int) *updateQueue {
	return &updateQueue{memory: make(chan TlObject, size), ready: make(chan struct{}, 1)}
}

func (q *updateQueue) notify() {
	select {
	case q.ready <- struct{}{}:
	default:
	}
}

func (q *updateQueue) segment(id uint64) string {
	return filepath.Join(q.dir, strconv.FormatUint(id, 10))
}

func (q *updateQueue) push(update TlObject) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return ErrClientClosed
	}

	// Once spilling starts, append every new update to disk until it drains.
	// Otherwise newer updates could overtake the older disk backlog.
	if q.dir == "" {
		select {
		case q.memory <- update:
			q.notify()
			return nil
		default:
		}
	}
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}
	if uint64(len(data)) > uint64(^uint32(0)) {
		return fmt.Errorf("update exceeds spool record size")
	}

	if q.dir == "" {
		q.dir, err = os.MkdirTemp("", "gotdbot-updates-*")
		if err != nil {
			return err
		}
		q.head, q.tail = 1, 1
	}

	if q.writer == nil || q.writeOffset >= updateSegmentSize {
		id := q.tail
		if q.writer != nil {
			id++
		}
		writer, openErr := os.OpenFile(q.segment(id), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if openErr != nil {
			if q.spilled == 0 {
				openErr = errors.Join(openErr, q.clearSpool())
			}
			return openErr
		}
		if q.writer != nil {
			err = q.writer.Close()
		}
		q.writer, q.tail, q.writeOffset = writer, id, 0
	}
	record := binary.BigEndian.AppendUint32(nil, uint32(len(data)))
	record = append(record, data...)
	n, writeErr := q.writer.WriteAt(record, q.writeOffset)
	if writeErr == nil && n != len(record) {
		writeErr = io.ErrShortWrite
	}
	if writeErr != nil {
		writeErr = errors.Join(err, writeErr, q.writer.Truncate(q.writeOffset))
		if q.spilled == 0 {
			writeErr = errors.Join(writeErr, q.clearSpool())
		}
		return writeErr
	}
	q.writeOffset += int64(n)
	q.spilled++
	q.notify()
	return err
}

func (q *updateQueue) next(stop <-chan struct{}) (TlObject, error) {
	for {
		select {
		case <-stop:
			return nil, ErrClientClosed
		default:
		}
		update, err := q.take()
		if update != nil || err != nil {
			return update, err
		}
		select {
		case <-q.ready:
		case <-stop:
			return nil, ErrClientClosed
		}
	}
}

func (q *updateQueue) take() (TlObject, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return nil, ErrClientClosed
	}
	defer func() {
		if len(q.memory) > 0 || q.spilled > 0 {
			q.notify()
		}
	}()
	select {
	case update := <-q.memory:
		return update, nil
	default:
	}
	if q.spilled == 0 {
		return nil, nil
	}
	for {
		if q.reader == nil {
			var err error
			q.reader, err = os.Open(q.segment(q.head))
			if err != nil {
				return nil, err
			}
		}
		info, err := q.reader.Stat()
		if err != nil {
			return nil, err
		}
		if q.readOffset == info.Size() && q.head < q.tail {
			if err := q.retireSegment(); err != nil {
				return nil, err
			}
			continue
		}
		var header [4]byte
		if _, err := q.reader.ReadAt(header[:], q.readOffset); err != nil {
			return nil, err
		}
		size := int64(binary.BigEndian.Uint32(header[:]))
		if size == 0 || size > info.Size()-q.readOffset-4 {
			return nil, fmt.Errorf("invalid update spool record length %d", size)
		}
		data := make([]byte, int(size))
		if _, err := q.reader.ReadAt(data, q.readOffset+4); err != nil {
			return nil, err
		}
		update, _, err := Unmarshal(data)
		if err != nil {
			return nil, err
		}
		q.readOffset += 4 + size
		q.spilled--
		if q.spilled == 0 {
			err = q.clearSpool()
		} else if q.readOffset == info.Size() && q.head < q.tail {
			err = q.retireSegment()
		}
		return update, err
	}
}

func (q *updateQueue) retireSegment() error {
	err := q.reader.Close()
	err = errors.Join(err, os.Remove(q.segment(q.head)))
	q.reader = nil
	q.head++
	q.readOffset = 0
	return err
}

func (q *updateQueue) clearSpool() error {
	var err error
	for _, file := range []*os.File{q.reader, q.writer} {
		if file != nil {
			err = errors.Join(err, file.Close())
		}
	}
	if q.dir != "" {
		err = errors.Join(err, os.RemoveAll(q.dir))
	}
	q.dir, q.reader, q.writer = "", nil, nil
	q.head, q.tail, q.readOffset, q.writeOffset, q.spilled = 0, 0, 0, 0, 0
	return err
}

func (q *updateQueue) close() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return nil
	}
	q.closed = true
	close(q.ready)
	for len(q.memory) > 0 {
		<-q.memory
	}
	return q.clearSpool()
}
