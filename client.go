package gotdbot

//go:generate go run ./internal/gen

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/AshokShau/gotdbot/internal/tdjson"
	"github.com/AshokShau/gotdbot/logger"
)

// Client manages requests, authentication, and updates for one Telegram account.
type Client struct {
	manager  atomic.Pointer[ClientManager]
	clientID int

	apiID    int32
	apiHash  string
	botToken string
	// phoneNumber is set if the user is logging in with a phone number.
	phoneNumber string
	config      *ClientOpts
	Logger      *logger.Logger

	updates        *updateQueue
	botUpdates     botUpdateState
	startupUpdates pendingUpdateState
	stop           chan struct{}
	closed         chan struct{}
	wg             sync.WaitGroup

	handlers     atomic.Pointer[handlersData]
	hMu          sync.Mutex
	panicHandler func(client *Client, update TlObject, r any)
	errorHandler func(client *Client, update TlObject, err error) error

	requestMu       sync.Mutex
	pendingRequests map[string]*pendingCall
	pendingMessages map[[2]int64]chan TlObject

	waiters     map[int64]map[string]*waiter
	waiterCount atomic.Int64
	wMu         sync.RWMutex

	// Auth state management
	authDone     chan struct{}
	authOnce     sync.Once
	authErr      error
	isAuthorized atomic.Bool
	startOnce    sync.Once
	closeOnce    sync.Once
	closedOnce   sync.Once
	requestID    atomic.Uint64
	sendRequest  func(int, []byte) // test seam; nil uses tdjson.SendBytes
	closing      atomic.Bool

	stopOnce    sync.Once
	meMu        sync.RWMutex
	me          *User
	lifecycleMu sync.Mutex
}

// Me returns a non-nil, read-only snapshot of the cached account.
func (c *Client) Me() *User {
	c.meMu.RLock()
	me := c.me
	c.meMu.RUnlock()
	if me == nil {
		return &User{}
	}
	return me
}

// NewClient creates a client using a bot token or user phone number.
// A nil config uses DefaultClientConfig. Call Start to authenticate the account.
func NewClient(apiID int32, apiHash, tokenOrPhone string, config *ClientOpts) (*Client, error) {
	tokenOrPhone = strings.TrimSpace(tokenOrPhone)
	if config == nil {
		config = DefaultClientConfig()
	} else {
		copyConfig := *config
		config = &copyConfig
		def := DefaultClientConfig()
		if config.UseFileDatabase == nil {
			config.UseFileDatabase = def.UseFileDatabase
		}
		if config.UseChatInfoDatabase == nil {
			config.UseChatInfoDatabase = def.UseChatInfoDatabase
		}
		if config.UseMessageDatabase == nil {
			config.UseMessageDatabase = def.UseMessageDatabase
		}
		if config.SystemLanguageCode == "" {
			config.SystemLanguageCode = def.SystemLanguageCode
		}
		if config.DeviceModel == "" {
			config.DeviceModel = def.DeviceModel
		}
		if config.SystemVersion == "" {
			config.SystemVersion = def.SystemVersion
		}
		if config.ApplicationVersion == "" {
			config.ApplicationVersion = def.ApplicationVersion
		}
		if config.DatabaseDirectory == "" {
			config.DatabaseDirectory = def.DatabaseDirectory
		}
		if config.FilesDirectory == "" {
			config.FilesDirectory = def.FilesDirectory
		}
		if config.Logger == nil {
			config.Logger = def.Logger
		}
		if config.AuthorizationTimeout <= 0 {
			config.AuthorizationTimeout = def.AuthorizationTimeout
		}
		if config.AutoRetry == nil {
			config.AutoRetry = def.AutoRetry
		}
		if config.ParseMode == "" {
			config.ParseMode = def.ParseMode
		}
	}

	if config.UpdateWorkers <= 0 {
		config.UpdateWorkers = 2
	}
	if config.UpdateQueueSize <= 0 {
		config.UpdateQueueSize = 1000
	}
	if err := tdjson.Init(config.LibraryPath); err != nil {
		return nil, err
	}

	botTokenRegex := regexp.MustCompile(`^\d+:[a-zA-Z0-9_-]+$`)
	var botToken, phoneNumber string
	if botTokenRegex.MatchString(tokenOrPhone) {
		botToken = tokenOrPhone
	} else {
		phoneNumber = tokenOrPhone
	}

	c := &Client{
		clientID:        tdjson.CreateClientID(),
		apiID:           apiID,
		apiHash:         apiHash,
		botToken:        botToken,
		phoneNumber:     phoneNumber,
		config:          config,
		Logger:          config.Logger,
		updates:         newUpdateQueue(config.UpdateQueueSize),
		stop:            make(chan struct{}),
		closed:          make(chan struct{}),
		authDone:        make(chan struct{}),
		pendingRequests: make(map[string]*pendingCall),
		pendingMessages: make(map[[2]int64]chan TlObject),
		me:              &User{},
	}
	c.me.Id = c.botTokenUserID()

	c.handlers.Store(&handlersData{
		handlers: make(map[int][]Handler),
	})
	c.waiters = make(map[int64]map[string]*waiter)
	c.panicHandler = config.PanicHandler
	c.errorHandler = config.ErrorHandler

	return c, nil
}

// initClient initializes the client's internal state and processor loop exactly once.
func (c *Client) initClient() {
	c.startOnce.Do(func() {
		c.lifecycleMu.Lock()
		defer c.lifecycleMu.Unlock()
		if c.closing.Load() {
			return
		}
		if c.manager.Load() == nil {
			m := GetDefaultManager(c.config.LibraryPath)
			if err := m.AddClient(c); err != nil {
				c.sendAuthError(err)
				return
			}
		}

		workers := c.config.UpdateWorkers
		c.Logger.Debug("Starting TDLib client workers", "user_id", c.Me().Id, "workers", workers)
		c.wg.Add(workers)
		for range workers {
			go c.processor()
		}
		tdjson.Send(c.clientID, `{"@type": "getOption", "name": "version"}`)
	})
}

// Start initializes the client and blocks until authorization is successful or fails.
func (c *Client) Start() error { return c.StartContext(context.Background()) }

// StartContext starts TDLib and waits for authentication or cancellation.
// Cancellation leaves the instance running; call Close to stop it.
func (c *Client) StartContext(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("nil context")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.closing.Load() {
		return ErrClientClosed
	}
	if c.isAuthorized.Load() {
		return nil
	}
	c.initClient()
	timer := time.NewTimer(c.config.AuthorizationTimeout)
	defer timer.Stop()
	select {
	case <-c.authDone:
		return c.authErr
	case <-ctx.Done():
		return ctx.Err()
	case <-c.stop:
		return ErrClientClosed
	case <-timer.C:
		c.Logger.Warn("TDLib authorization timed out", "user_id", c.Me().Id, "timeout", c.config.AuthorizationTimeout)
		return fmt.Errorf("authorization timeout")
	}
}

// Init starts TDLib and the client workers without waiting for authentication.
// Use it for requests that do not require authorization; Start authenticates the account.
func (c *Client) Init() {
	c.initClient()
}

// Idle blocks the current goroutine until a SIGINT or SIGTERM signal is received.
func (c *Client) Idle() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)
	select {
	case <-sig:
	case <-c.stop:
	case <-c.closed:
	}
	c.Close()
}

// Close closes the TDLib instance. All databases will be flushed to disk and properly
// closed. After the close completes, updateAuthorizationState with authorizationStateClosed
// will be sent. Can be called before initialization or from a handler.
// In-flight application handlers are not waited on; future requests fail closed.
func (c *Client) Close() {
	c.closeOnce.Do(func() {
		c.Logger.Debug("Closing TDLib client", "user_id", c.Me().Id)
		c.lifecycleMu.Lock()
		c.closing.Store(true)
		select {
		case <-c.closed:
			c.stopOnce.Do(func() { close(c.stop) })
			if err := c.updates.close(); err != nil {
				c.Logger.Error("Cannot clean up update queue", "error", err)
			}
			if manager := c.manager.Load(); manager != nil {
				manager.RemoveClient(c)
			}
			c.lifecycleMu.Unlock()
			return
		default:
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if c.manager.Load() == nil && c.sendRequest == nil {
			GetDefaultManager(c.config.LibraryPath).AddClient(c)
		}
		c.lifecycleMu.Unlock()
		if c.manager.Load() != nil {
			if _, err := c.SendWithContext(ctx, &Close{}); err != nil {
				c.Logger.Warn("TDLib close request failed", "user_id", c.Me().Id, "error", err)
			}
		} else if c.sendRequest != nil {
			c.sendRequest(c.clientID, []byte(`{"@type":"close"}`))
		} else {
			tdjson.Send(c.clientID, `{"@type":"close"}`)
		}

		select {
		case <-c.closed:
		case <-ctx.Done():
			c.Logger.Warn("Timeout waiting for TDLib to close")
		}

		c.stopOnce.Do(func() { close(c.stop) })
		if err := c.updates.close(); err != nil {
			c.Logger.Error("Cannot clean up update queue", "error", err)
		}
		if manager := c.manager.Load(); manager != nil {
			manager.RemoveClient(c)
		}
	})
}

func (c *Client) connectionStateHandler(u *UpdateConnectionState) {
	state := strings.TrimPrefix(u.State.GetType(), "connectionState")
	c.Logger.Info("Connection state changed", "state", state)
}
