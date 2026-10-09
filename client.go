package gotdbot

//go:generate go run ./internal/gen

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"os/signal"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/AshokShau/gotdbot/logger"

	"github.com/AshokShau/gotdbot/internal/qrcode"
	"github.com/AshokShau/gotdbot/internal/tdjson"
)

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

	updates *updateQueue
	stop    chan struct{}
	closed  chan struct{}
	wg      sync.WaitGroup

	handlers     atomic.Pointer[handlersData]
	hMu          sync.Mutex
	panicHandler func(client *Client, update TlObject, r interface{})
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
	lifecycleMu sync.Mutex

	// Me is the bot's User info, as returned by client.GetMe.
	// Populated when authorization is ready.
	Me *User
}

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
	if config.RequestTimeout <= 0 {
		config.RequestTimeout = 30 * time.Second
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
	}

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
		c.wg.Add(workers)
		for i := 0; i < workers; i++ {
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
		return fmt.Errorf("authorization timeout")
	}
}

// SendDummyRequest sends a dummy request to TDLib to initialize its state.
// This is useful if you need to call unauthenticated TDLib methods before calling Start().
func (c *Client) SendDummyRequest() {
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
			_, _ = c.SendWithContext(ctx, &Close{})
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

func (c *Client) processor() {
	defer c.wg.Done()
	for {
		update, err := c.updates.next(c.stop)
		if errors.Is(err, ErrClientClosed) {
			return
		}
		if err != nil {
			c.Logger.Error("Cannot read queued TDLib update", "error", err)
			if update == nil {
				timer := time.NewTimer(time.Second)
				select {
				case <-c.stop:
					timer.Stop()
					return
				case <-timer.C:
				}
				continue
			}
		}
		select {
		case <-c.stop:
			return
		default:
		}
		func() {
			defer func() {
				if r := recover(); r != nil {
					if c.panicHandler != nil {
						c.panicHandler(c, update, r)
					} else {
						c.Logger.Error("Panic in handler", "panic", r)
					}
				}
			}()

			// Internal state handling is independent of application handler groups.
			switch u := update.(type) {
			case *UpdateAuthorizationState:
				if u != nil && u.AuthorizationState != nil {
					c.authHandler(u)
				}
			case *UpdateConnectionState:
				if u != nil && u.State != nil {
					c.connectionStateHandler(u)
				}
			}

			// Snapshot waiters before invoking application filters. Filters may
			// register/remove waiters and must never execute under wMu.
			c.wMu.RLock()
			var candidates []*waiter
			chatID := getChatIDFromUpdate(update)
			for id, inner := range c.waiters {
				if id == 0 || id == chatID {
					for _, w := range inner {
						candidates = append(candidates, w)
					}
				}
			}
			c.wMu.RUnlock()
			for _, w := range candidates {
				if w.filter(c, update) {
					select {
					case w.response <- update:
					default:
					}
				}
			}

			hData := c.handlers.Load()
			if hData == nil {
				return
			}
			handlersMap := hData.handlers
			groups := hData.groups

			handleError := func(err error) error {
				if err == nil {
					return nil
				}
				if errors.Is(err, EndGroups) || errors.Is(err, ContinueGroups) || errors.Is(err, ContinueHandlers) {
					return err
				}
				if c.errorHandler != nil {
					return c.errorHandler(c, update, err)
				}
				c.Logger.Error("Handler error", "error", err, "type", update.GetType())
				return nil
			}

			for _, group := range groups {
				groupHandlers := handlersMap[group]
				for _, h := range groupHandlers {
					if h.CheckUpdate(c, update) {
						err := h.HandleUpdate(c, update)
						action := handleError(err)

						if errors.Is(action, EndGroups) {
							return
						}
						if errors.Is(action, ContinueGroups) {
							break // Move to next group
						}
						if errors.Is(action, ContinueHandlers) {
							continue // Move to next handler in same group
						}
						// Run at most one matching handler in each group.
						break
					}
				}
			}
		}()
	}
}

// sendAuthError publishes the first authorization result to all Start callers.
func (c *Client) sendAuthError(err error) {
	c.authOnce.Do(func() { c.authErr = err; close(c.authDone) })
}

func (c *Client) authHandler(authState *UpdateAuthorizationState) {
	c.Logger.Debug("Authorization state update", "state", authState.AuthorizationState.GetType())

	switch state := authState.AuthorizationState.(type) {
	case *AuthorizationStateWaitTdlibParameters:
		if c.config.TDLibOptions != nil {
			c.config.TDLibOptions.forEachSet(func(k string, v interface{}) {
				if opt := toOptionValue(v); opt != nil {
					err := c.SetOption(k, &SetOptionOpts{Value: opt})
					if err != nil {
						c.Logger.Error("Error setting option", "option", k, "error", err)
					}
				}
			})
		}

		err := c.SetTdlibParameters(
			c.apiHash,
			c.apiID,
			c.config.ApplicationVersion,
			c.config.DatabaseDirectory,
			[]byte(c.config.DatabaseEncryptionKey),
			c.config.DeviceModel,
			c.config.FilesDirectory,
			c.config.SystemLanguageCode,
			c.config.SystemVersion,
			&SetTdlibParametersOpts{
				UseChatInfoDatabase: *c.config.UseChatInfoDatabase,
				UseFileDatabase:     *c.config.UseFileDatabase,
				UseMessageDatabase:  *c.config.UseMessageDatabase,
				UseSecretChats:      c.config.UseSecretChats,
				UseTestDc:           c.config.UseTestDC,
			},
		)
		if err != nil {
			c.Logger.Error("Error setting tdlib parameters", "error", err)
			c.sendAuthError(err)
		}

	case *AuthorizationStateWaitPhoneNumber:
		if c.botToken != "" {
			err := c.CheckAuthenticationBotToken(c.botToken)
			if err != nil {
				c.Logger.Error("Error checking bot token", "error", err)
				c.sendAuthError(err)
			}
		} else {
			if c.config.QrMode {
				err := c.RequestQrCodeAuthentication(nil)
				if err != nil {
					c.Logger.Error("Error requesting QR code", "error", err)
					c.sendAuthError(err)
				}
			} else if c.phoneNumber != "" {
				err := c.SetAuthenticationPhoneNumber(c.phoneNumber, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.sendAuthError(err)
				}
			} else {
				fmt.Print("Enter phone number: ")
				reader := bufio.NewReader(os.Stdin)
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)
				err := c.SetAuthenticationPhoneNumber(phone, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.sendAuthError(err)
				}
			}
		}

	case *AuthorizationStateWaitOtherDeviceConfirmation:
		link := state.Link
		fmt.Printf("Scan the QR code below: or open the link in Telegram: %s\n", link)
		q, err := qrcode.NewQRCode(link)
		if err != nil {
			fmt.Printf("Failed to generate QR: %v\nLink: %s\n", err, link)
		} else {
			fmt.Println(q.ToSmallString(false))
		}

	case *AuthorizationStateWaitCode:
		codeInfo := state.CodeInfo
		codeType := strings.TrimPrefix(codeInfo.Type.GetType(), "authenticationCodeType")
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("Enter the code received via %s: ", codeType)
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)
			if code == "" {
				continue
			}
			err := c.CheckAuthenticationCode(code)
			if err != nil {
				fmt.Printf("Error checking code: %v\n", err)
				continue
			}
			break
		}

	case *AuthorizationStateWaitPassword:
		hint := state.PasswordHint
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("Enter your 2FA password (hint: %s): ", hint)
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password == "" {
				continue
			}
			err := c.CheckAuthenticationPassword(password)
			if err != nil {
				fmt.Printf("Error checking password: %v\n", err)
				continue
			}
			break
		}

	case *AuthorizationStateWaitRegistration:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter first name: ")
		firstName, _ := reader.ReadString('\n')
		firstName = strings.TrimSpace(firstName)
		fmt.Print("Enter last name: ")
		lastName, _ := reader.ReadString('\n')
		lastName = strings.TrimSpace(lastName)

		err := c.RegisterUser(firstName, lastName, nil)
		if err != nil {
			c.Logger.Error("Error registering user", "error", err)
			c.sendAuthError(err)
		}

	case *AuthorizationStateWaitPremiumPurchase:
		c.Logger.Warn("Account is limited and requires Telegram Premium.")
		c.sendAuthError(WaitPremiumPurchase)

	case *AuthorizationStateReady:
		me, err := c.GetMe()
		if err != nil {
			c.Logger.Error("Failed to get me", "error", err)
			c.sendAuthError(err)
			return
		}

		c.meMu.Lock()
		c.Me = me
		c.meMu.Unlock()
		c.isAuthorized.Store(true)

		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		c.Logger.Info("Logged in", "user_id", me.Id, "username", username)

		c.sendAuthError(nil)

	}
	return
}

func (c *Client) connectionStateHandler(u *UpdateConnectionState) {
	state := strings.TrimPrefix(u.State.GetType(), "connectionState")
	c.Logger.Info("Connection state changed", "state", state)
}

// Send dispatches req to TDLib and waits for the response.
func (c *Client) Send(req TlObject) (TlObject, error) {
	return c.SendWithContext(context.Background(), req)
}

// SendWithContext dispatches req to TDLib and waits for the response, honouring ctx.
func (c *Client) SendWithContext(ctx context.Context, req TlObject) (TlObject, error) {
	if ctx == nil {
		return nil, fmt.Errorf("nil context")
	}
	if req == nil || (reflect.ValueOf(req).Kind() == reflect.Pointer && reflect.ValueOf(req).IsNil()) {
		return nil, fmt.Errorf("nil TDLib request")
	}
	if c.closing.Load() && req.GetType() != "close" {
		return nil, ErrClientClosed
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if _, isFn := req.(tlFunction); !isFn {
		return nil, fmt.Errorf("TDLib requests must be function pointers, got %T", req)
	}
	// Serialize without mutating the request; correlation belongs to the transport.
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var wire map[string]json.RawMessage
	if err := json.Unmarshal(data, &wire); err != nil {
		return nil, err
	}
	// Bound retries, lookup requests, send completion and flood waits together.
	ctx, cancel := context.WithTimeout(ctx, c.config.RequestTimeout)
	defer cancel()
	chatLoaded := req.GetType() == "getChat"
	messageLoaded := strings.HasPrefix(req.GetType(), "getMessage")
	floodRetried := false
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		select {
		case <-c.stop:
			return nil, ErrClientClosed
		default:
		}
		extra := strconv.FormatUint(c.requestID.Add(1), 10)
		wire["@extra"], _ = json.Marshal(extra)
		data, err := json.Marshal(wire)
		if err != nil {
			return nil, err
		}
		call := &pendingCall{response: make(chan TlObject, 1), waitForSend: waitsForSend(req)}
		c.requestMu.Lock()
		c.pendingRequests[extra] = call
		c.requestMu.Unlock()
		if c.sendRequest != nil {
			c.sendRequest(c.clientID, data)
		} else {
			tdjson.SendBytes(c.clientID, data)
		}
		var result TlObject
		select {
		case result = <-call.response:
		case <-ctx.Done():
			c.finishCall(extra)
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return nil, errors.Join(SendTimeout, ctx.Err())
			}
			return nil, ctx.Err()
		case <-c.closed:
			c.finishCall(extra)
			return nil, ErrClientClosed
		case <-c.stop:
			c.finishCall(extra)
			return nil, ErrClientClosed
		}
		if errObj, ok := result.(*Error); ok {
			c.finishCall(extra)
			retry, retryErr := c.retryRequest(ctx, req, errObj, &chatLoaded, &messageLoaded, &floodRetried)
			if retryErr != nil {
				return nil, retryErr
			}
			if retry {
				continue
			}
			return nil, errObj
		}
		if waitsForSend(req) {
			result, err = c.waitSend(ctx, result, call)
			if errors.Is(err, context.DeadlineExceeded) {
				err = errors.Join(SendTimeout, err)
			}
		}
		c.finishCall(extra)
		return result, err
	}
}

// Retry only explicit server rejections; never retry a transport timeout.
func (c *Client) retryRequest(ctx context.Context, req TlObject, e *Error, chatLoaded, messageLoaded, floodRetried *bool) (bool, error) {
	if c.config == nil || c.config.AutoRetry == nil {
		return false, nil
	}
	policy := c.config.AutoRetry
	if e.Code == 429 && !*floodRetried && policy.MaxFloodWait > 0 {
		seconds := e.GetRetryAfter()
		if seconds > 0 && int64(seconds) <= int64(policy.MaxFloodWait/time.Second) {
			*floodRetried = true
			timer := time.NewTimer(time.Duration(seconds) * time.Second)
			defer timer.Stop()
			select {
			case <-timer.C:
				return true, nil
			case <-ctx.Done():
				return false, ctx.Err()
			case <-c.stop:
				return false, ErrClientClosed
			}
		}
	}
	if e.Code != 400 {
		return false, nil
	}
	if e.Message == "Message not found" && !*messageLoaded && policy.MessageNotFound {
		*messageLoaded = true
		chatID, messageID := extractChatAndMessageID(req)
		if chatID != 0 && messageID != 0 {
			_, err := c.SendWithContext(ctx, &GetMessage{ChatId: chatID, MessageId: messageID})
			return err == nil, nil
		}
	}
	if e.Message == "Chat not found" && !*chatLoaded && policy.ChatNotFound {
		*chatLoaded = true
		if id, _ := extractChatAndMessageID(req); id != 0 {
			_, err := c.SendWithContext(ctx, &GetChat{ChatId: id})
			return err == nil, nil
		}
	}
	return false, nil
}

func extractChatAndMessageID(req TlObject) (int64, int64) {
	data, err := json.Marshal(req)
	if err != nil {
		return 0, 0
	}
	// TDLib accepts int64 values as JSON numbers or strings. json.Number keeps
	// their precision in either form; decoding through float64 loses large IDs.
	var ids struct {
		ChatID    json.Number `json:"chat_id"`
		MessageID json.Number `json:"message_id"`
	}
	if err := json.Unmarshal(data, &ids); err != nil {
		return 0, 0
	}
	chatID, _ := ids.ChatID.Int64()
	messageID, _ := ids.MessageID.Int64()
	return chatID, messageID
}

func toOptionValue(v interface{}) OptionValue {
	switch val := v.(type) {
	case bool:
		return &OptionValueBoolean{Value: val}
	case int:
		return &OptionValueInteger{Value: int64(val)}
	case int32:
		return &OptionValueInteger{Value: int64(val)}
	case int64:
		return &OptionValueInteger{Value: val}
	case string:
		return &OptionValueString{Value: val}
	case nil:
		return &OptionValueEmpty{}
	default:
		return nil
	}
}

func Bool(b bool) *bool {
	return &b
}

type handlersData struct {
	handlers map[int][]Handler
	groups   []int
}

type waiter struct {
	filter   func(client *Client, update TlObject) bool
	response chan TlObject
}

// WaitFor registers a waiter and blocks until a matching update arrives or timeout occurs.
func (c *Client) WaitFor(filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	return c.WaitForChat(0, filter, timeout)
}

// WaitForContext registers a waiter and blocks until a matching update arrives,
// timeout occurs, or ctx is cancelled.
func (c *Client) WaitForContext(ctx context.Context, chatId int64, filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	if ctx == nil || filter == nil {
		return nil, fmt.Errorf("context and filter required")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c.closing.Load() {
		return nil, ErrClientClosed
	}
	ch := make(chan TlObject, 1)
	idNum := c.waiterCount.Add(1)
	id := fmt.Sprintf("%d", idNum)

	c.wMu.Lock()
	if c.waiters == nil {
		c.waiters = make(map[int64]map[string]*waiter)
	}
	if c.waiters[chatId] == nil {
		c.waiters[chatId] = make(map[string]*waiter)
	}
	c.waiters[chatId][id] = &waiter{filter: filter, response: ch}
	c.wMu.Unlock()

	defer func() {
		c.wMu.Lock()
		delete(c.waiters[chatId], id)
		if len(c.waiters[chatId]) == 0 {
			delete(c.waiters, chatId)
		}
		c.wMu.Unlock()
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case u := <-ch:
		return u, nil
	case <-c.closed:
		return nil, ErrClientClosed
	case <-c.stop:
		return nil, ErrClientClosed
	case <-timer.C:
		return nil, ConversationTimeout
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// WaitForChat registers a waiter for a specific chat and blocks until a matching
// update arrives or timeout occurs.
func (c *Client) WaitForChat(chatId int64, filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	return c.WaitForContext(context.Background(), chatId, filter, timeout)
}
