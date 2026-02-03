package gotdbot

//go:generate go run ./scripts/generate

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/AshokShau/gotdbot/tdjson"
)

var StopHandlers = fmt.Errorf("stop handlers")

type ClientConfig struct {
	LibraryPath             string
	UseTestDC               bool
	DatabaseDirectory       string
	FilesDirectory          string
	DatabaseEncryptionKey   string
	UseFileDatabase         *bool
	UseChatInfoDatabase     *bool
	UseMessageDatabase      *bool
	UseSecretChats          *bool
	LoadMessagesBeforeReply bool
	SystemLanguageCode      string
	DeviceModel             string
	SystemVersion           string
	ApplicationVersion      string
	TDLibOptions            map[string]interface{}
	Logger                  *slog.Logger
}

func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		DatabaseDirectory:       "database",
		FilesDirectory:          "database",
		DatabaseEncryptionKey:   "default",
		UseFileDatabase:         Bool(true),
		UseChatInfoDatabase:     Bool(true),
		UseMessageDatabase:      Bool(true),
		UseSecretChats:          Bool(false),
		LoadMessagesBeforeReply: false,
		SystemLanguageCode:      "en",
		DeviceModel:             "Gotdbot",
		SystemVersion:           runtime.GOOS,
		ApplicationVersion:      "Gotdbot " + Version,
		Logger:                  slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

type Client struct {
	clientID int

	apiID    int32
	apiHash  string
	botToken string
	config   *ClientConfig
	Logger   *slog.Logger

	updates chan TlObject
	stop    chan struct{}
	closed  chan struct{}
	wg      sync.WaitGroup

	handlers     map[string][]*Handler
	initializers []*Handler
	finalizers   []*Handler
	hMu          sync.RWMutex

	pendingRequests sync.Map // map[string]chan TlObject
	pendingMessages sync.Map // map[string]chan TlObject

	// Cache
	me   *User
	meMu sync.RWMutex

	// Auth state management
	authErrorChan chan error
	isAuthorized  bool
}

func NewClient(apiID int32, apiHash, botToken string, config *ClientConfig) *Client {
	if config == nil {
		config = DefaultClientConfig()
	} else {
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
		if config.UseSecretChats == nil {
			config.UseSecretChats = def.UseSecretChats
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
	}

	if err := tdjson.Init(config.LibraryPath); err != nil {
		config.Logger.Error("Failed to initialize TDLib", "error", err)
		os.Exit(1)
	}

	c := &Client{
		clientID:      tdjson.CreateClientID(),
		apiID:         apiID,
		apiHash:       apiHash,
		botToken:      botToken,
		config:        config,
		Logger:        config.Logger,
		updates:       make(chan TlObject, 1000),
		stop:          make(chan struct{}),
		closed:        make(chan struct{}),
		handlers:      make(map[string][]*Handler),
		authErrorChan: make(chan error, 1),
	}

	c.AddHandler("updateAuthorizationState", c.authHandler, nil, 0)
	c.AddHandler("updateUser", c.userHandler, nil, 0)
	c.AddHandler("updateConnectionState", c.connectionStateHandler, nil, 0)
	c.AddHandler("updateMessageSendSucceeded", c.messageSendSucceededHandler, nil, 0)
	c.AddHandler("updateMessageSendFailed", c.messageSendFailedHandler, nil, 0)
	return c
}

// Start initializes the client and blocks until authorization is successful or fails.
func (c *Client) Start() error {
	c.wg.Add(1)
	go c.receiver()

	c.wg.Add(1)
	go c.processor()

	tdjson.Send(c.clientID, `{"@type": "getOption", "name": "version"}`)

	select {
	case err := <-c.authErrorChan:
		return err
	case <-time.After(60 * time.Second):
		return fmt.Errorf("authorization timeout")
	}
}

// Idle blocks the current goroutine until a SIGINT or SIGTERM signal is received.
func (c *Client) Idle() {
	c.Logger.Info("Bot is running. Press Ctrl+C to stop.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	c.Logger.Info("Stopping...")
	c.Stop()
}

// SetTdlibLogVerbosityLevel sets the verbosity level of TDLib's internal logging.
func SetTdlibLogVerbosityLevel(level int32) {
	req := fmt.Sprintf(`{"@type": "setLogVerbosityLevel", "new_verbosity_level": %d}`, level)
	tdjson.Execute(req)
}

// SetTdlibLogStreamFile redirects internal TDLib logs to a file.
func SetTdlibLogStreamFile(path string, maxFileSize int64, redirectStderr bool) {
	req := fmt.Sprintf(`{"@type": "setLogStream", "log_stream": {"@type": "logStreamFile", "path": "%s", "max_file_size": %d, "redirect_stderr": %v}}`, path, maxFileSize, redirectStderr)
	tdjson.Execute(req)
}

// SetTdlibLogStreamEmpty disables internal TDLib logging.
func SetTdlibLogStreamEmpty() {
	req := `{"@type": "setLogStream", "log_stream": {"@type": "logStreamEmpty"}}`
	tdjson.Execute(req)
}

func (c *Client) Stop() {
	_, _ = c.Send(&Close{})

	select {
	case <-c.closed:
	case <-time.After(5 * time.Second):
		c.Logger.Warn("Timeout waiting for TDLib to close")
	}

	close(c.stop)
	c.wg.Wait()
}

func (c *Client) receiver() {
	defer c.wg.Done()
	for {
		select {
		case <-c.stop:
			return
		default:
			res := tdjson.Receive(0.1)
			if res != "" {
				data := []byte(res)
				obj, extra, err := Unmarshal(data)
				if err != nil {
					// fmt.Println("Error unmarshaling:", err)
					continue
				}
				if obj == nil {
					continue
				}

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

func (c *Client) processor() {
	defer c.wg.Done()
	for {
		select {
		case <-c.stop:
			return
		case update := <-c.updates:
			go c.handleUpdate(update)
		}
	}
}

func (c *Client) handleUpdate(update TlObject) {
	c.hMu.RLock()
	defer c.hMu.RUnlock()

	// Initializers
	for _, h := range c.initializers {
		if h.Filter != nil && !h.Filter(update) {
			continue
		}
		if err := h.Func(c, update); errors.Is(err, StopHandlers) {
			return
		}
	}

	// Main Handlers
	if typeHandlers, ok := c.handlers[update.Type()]; ok {
		for _, h := range typeHandlers {
			if h.Filter != nil && !h.Filter(update) {
				continue
			}
			if err := h.Func(c, update); errors.Is(err, StopHandlers) {
				return
			}
		}
	}

	// Finalizers
	for _, h := range c.finalizers {
		if h.Filter != nil && !h.Filter(update) {
			continue
		}
		if err := h.Func(c, update); errors.Is(err, StopHandlers) {
			return
		}
	}
}

func (c *Client) authHandler(client *Client, update TlObject) error {
	authState, ok := update.(*UpdateAuthorizationState)
	if !ok {
		return nil
	}

	c.Logger.Debug("Authorization state update", "state", authState.AuthorizationState.Type())

	switch authState.AuthorizationState.Type() {
	case "authorizationStateWaitTdlibParameters":
		if len(c.config.TDLibOptions) > 0 {
			for k, v := range c.config.TDLibOptions {
				if opt := toOptionValue(v); opt != nil {
					_, err := c.SetOption(k, &SetOptionOpts{Value: opt})
					if err != nil {
						c.Logger.Error("Error setting option", "option", k, "error", err)
					}
				}
			}
		}

		c.Logger.Debug("Setting TDLib parameters",
			"use_test_dc", c.config.UseTestDC,
			"database_directory", c.config.DatabaseDirectory,
			"files_directory", c.config.FilesDirectory,
			"use_file_database", *c.config.UseFileDatabase,
			"use_chat_info_database", *c.config.UseChatInfoDatabase,
			"use_message_database", *c.config.UseMessageDatabase,
			"use_secret_chats", *c.config.UseSecretChats,
			"api_id", c.apiID,
			"system_language_code", c.config.SystemLanguageCode,
			"device_model", c.config.DeviceModel,
			"system_version", c.config.SystemVersion,
			"application_version", c.config.ApplicationVersion,
		)

		_, err := c.SetTdlibParameters(
			c.config.UseTestDC,
			c.config.DatabaseDirectory,
			c.config.FilesDirectory,
			[]byte(c.config.DatabaseEncryptionKey),
			*c.config.UseFileDatabase,
			*c.config.UseChatInfoDatabase,
			*c.config.UseMessageDatabase,
			*c.config.UseSecretChats,
			c.apiID,
			c.apiHash,
			c.config.SystemLanguageCode,
			c.config.DeviceModel,
			c.config.SystemVersion,
			c.config.ApplicationVersion,
		)
		if err != nil {
			c.Logger.Error("Error setting tdlib parameters", "error", err)
			c.authErrorChan <- err
		}

	case "authorizationStateWaitPhoneNumber":
		_, err := c.CheckAuthenticationBotToken(c.botToken)
		if err != nil {
			c.Logger.Error("Error checking bot token", "error", err)
			c.authErrorChan <- err
		}

	case "authorizationStateReady":
		c.isAuthorized = true
		me, err := c.GetMe()
		if err != nil {
			c.Logger.Error("Failed to get me", "error", err)
			c.authErrorChan <- err
			return nil
		}
		c.meMu.Lock()
		c.me = me
		c.meMu.Unlock()

		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		c.Logger.Info("Logged in", "user_id", me.Id, "username", username)

		select {
		case c.authErrorChan <- nil:
		default:
		}

	case "authorizationStateClosed":
		if !c.isAuthorized {
			c.authErrorChan <- fmt.Errorf("authorization closed unexpectedly")
		}
		select {
		case <-c.closed:
		default:
			close(c.closed)
		}
	}
	return nil
}

func (c *Client) userHandler(client *Client, update TlObject) error {
	u, ok := update.(*UpdateUser)
	if !ok {
		return nil
	}

	c.meMu.RLock()
	isMe := c.me != nil && c.me.Id == u.User.Id
	c.meMu.RUnlock()

	if isMe {
		c.meMu.Lock()
		c.me = u.User
		c.meMu.Unlock()
	}
	return nil
}

func (c *Client) connectionStateHandler(client *Client, update TlObject) error {
	u, ok := update.(*UpdateConnectionState)
	if !ok {
		return nil
	}

	state := u.State.Type()
	state = strings.TrimPrefix(state, "connectionState")
	c.Logger.Info("Connection state changed", "state", state)
	return nil
}

func (c *Client) messageSendSucceededHandler(client *Client, update TlObject) error {
	u, ok := update.(*UpdateMessageSendSucceeded)
	if !ok {
		return nil
	}
	key := fmt.Sprintf("%d:%d", u.Message.ChatId, u.OldMessageId)
	if ch, ok := c.pendingMessages.Load(key); ok {
		ch.(chan TlObject) <- u.Message
		c.pendingMessages.Delete(key)
	}
	return nil
}

func (c *Client) messageSendFailedHandler(client *Client, update TlObject) error {
	u, ok := update.(*UpdateMessageSendFailed)
	if !ok {
		return nil
	}
	key := fmt.Sprintf("%d:%d", u.Message.ChatId, u.OldMessageId)
	if ch, ok := c.pendingMessages.Load(key); ok {
		ch.(chan TlObject) <- u.Error
		c.pendingMessages.Delete(key)
	}
	return nil
}

func (c *Client) Send(req TlObject) (TlObject, error) {
	extra := fmt.Sprintf("%d", time.Now().UnixNano())

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return nil, err
	}
	tmp["@extra"] = extra
	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	ch := make(chan TlObject, 1)
	c.pendingRequests.Store(extra, ch)

	tdjson.Send(c.clientID, string(data))

	select {
	case res := <-ch:
		if res.Type() == "error" {
			if errObj, ok := res.(*Error); ok {
				return nil, fmt.Errorf("TDLib error %d: %s", errObj.Code, errObj.Message)
			}
		}
		return res, nil
	case <-time.After(30 * time.Second):
		c.pendingRequests.Delete(extra)
		return nil, fmt.Errorf("timeout")
	}
}

// AddHandler registers a handler for updates.
func (c *Client) AddHandler(updateType string, fn HandlerFunc, filter FilterFunc, position int) {
	c.hMu.Lock()
	defer c.hMu.Unlock()

	h := &Handler{Func: fn, UpdateType: updateType, Filter: filter, Position: position}

	if updateType == "initializer" {
		c.initializers = append(c.initializers, h)
		sortHandlers(c.initializers)
	} else if updateType == "finalizer" {
		c.finalizers = append(c.finalizers, h)
		sortHandlers(c.finalizers)
	} else {
		c.handlers[updateType] = append(c.handlers[updateType], h)
		sortHandlers(c.handlers[updateType])
	}
}

func sortHandlers(handlers []*Handler) {
	sort.Slice(handlers, func(i, j int) bool {
		return handlers[i].Position < handlers[j].Position
	})
}

// Me returns the cached User object for the current bot.
func (c *Client) Me() *User {
	c.meMu.RLock()
	defer c.meMu.RUnlock()
	return c.me
}

func toOptionValue(v interface{}) OptionValue {
	switch val := v.(type) {
	case bool:
		return &OptionValueBoolean{Value: val}
	case int:
		return &OptionValueInteger{Value: fmt.Sprintf("%d", val)}
	case int32:
		return &OptionValueInteger{Value: fmt.Sprintf("%d", val)}
	case int64:
		return &OptionValueInteger{Value: fmt.Sprintf("%d", val)}
	case string:
		return &OptionValueString{Value: val}
	case nil:
		return &OptionValueEmpty{}
	default:
		return nil
	}
}
