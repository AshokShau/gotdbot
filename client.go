package gotdbot

//go:generate go run ./scripts/generate

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/AshokShau/gotdbot/internal/qrcode"
	"github.com/AshokShau/gotdbot/tdjson"
)

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
	QrMode                  bool
	AuthorizationTimeout    time.Duration
	LogVerbosityLevel       int32
	LogStream               LogStream

	// DispatcherOpts configures the dispatcher
	DispatcherOpts *DispatcherOpts
}

func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		DatabaseDirectory:       "database",
		FilesDirectory:          "",
		DatabaseEncryptionKey:   "",
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
		QrMode:                  false,
		AuthorizationTimeout:    60 * time.Second,
		LogVerbosityLevel:       2,
	}
}

type Client struct {
	clientID int

	apiID    int32
	apiHash  string
	botToken string
	// phoneNumber is set if the user is logging in with a phone number.
	phoneNumber string
	config      *ClientConfig
	Logger      *slog.Logger

	updates chan TlObject
	stop    chan struct{}
	closed  chan struct{}
	wg      sync.WaitGroup

	Dispatcher *Dispatcher

	pendingRequests sync.Map // map[string]chan TlObject
	pendingMessages sync.Map // map[string]chan TlObject

	// Cache
	me   *User
	meMu sync.RWMutex

	// Auth state management
	authErrorChan chan error
	isAuthorized  bool
}

func NewClient(apiID int32, apiHash, tokenOrPhone string, config *ClientConfig) (*Client, error) {
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
		if config.AuthorizationTimeout == 0 {
			config.AuthorizationTimeout = def.AuthorizationTimeout
		}
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
		clientID:      tdjson.CreateClientID(),
		apiID:         apiID,
		apiHash:       apiHash,
		botToken:      botToken,
		phoneNumber:   phoneNumber,
		config:        config,
		Logger:        config.Logger,
		updates:       make(chan TlObject, 1000),
		stop:          make(chan struct{}),
		closed:        make(chan struct{}),
		authErrorChan: make(chan error, 1),
	}

	c.Dispatcher = NewDispatcher(c, config.DispatcherOpts)

	c.Dispatcher.AddHandlerToGroup(&internalHandler{handleFunc: c.authHandler, updateType: "updateAuthorizationState"}, -1)
	c.Dispatcher.AddHandlerToGroup(&internalHandler{handleFunc: c.userHandler, updateType: "updateUser"}, -1)
	c.Dispatcher.AddHandlerToGroup(&internalHandler{handleFunc: c.connectionStateHandler, updateType: "updateConnectionState"}, -1)
	c.Dispatcher.AddHandlerToGroup(&internalHandler{handleFunc: c.messageSendSucceededHandler, updateType: "updateMessageSendSucceeded"}, -1)
	c.Dispatcher.AddHandlerToGroup(&internalHandler{handleFunc: c.messageSendFailedHandler, updateType: "updateMessageSendFailed"}, -1)
	return c, nil
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
	case <-time.After(c.config.AuthorizationTimeout):
		return fmt.Errorf("authorization timeout")
	}
}

// Idle blocks the current goroutine until a SIGINT or SIGTERM signal is received.
func (c *Client) Idle() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	c.Close()
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

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (c *Client) Close() {
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
			c.Dispatcher.ProcessUpdate(update)
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
					err := c.SetOption(k, &SetOptionOpts{Value: opt})
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
				UseSecretChats:      *c.config.UseSecretChats,
				UseTestDc:           c.config.UseTestDC,
			},
		)
		if err != nil {
			c.Logger.Error("Error setting tdlib parameters", "error", err)
			c.authErrorChan <- err
		}

	case "authorizationStateWaitPhoneNumber":
		if c.botToken != "" {
			err := c.CheckAuthenticationBotToken(c.botToken)
			if err != nil {
				c.Logger.Error("Error checking bot token", "error", err)
				c.authErrorChan <- err
			}
		} else {
			// User login
			if c.config.QrMode {
				err := c.RequestQrCodeAuthentication(nil)
				if err != nil {
					c.Logger.Error("Error requesting QR code", "error", err)
					c.authErrorChan <- err
				}
			} else if c.phoneNumber != "" {
				err := c.SetAuthenticationPhoneNumber(c.phoneNumber, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.authErrorChan <- err
				}
			} else {
				fmt.Print("Enter phone number: ")
				reader := bufio.NewReader(os.Stdin)
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)
				err := c.SetAuthenticationPhoneNumber(phone, nil)
				if err != nil {
					c.Logger.Error("Error setting phone number", "error", err)
					c.authErrorChan <- err
				}
			}
		}

	case "authorizationStateWaitOtherDeviceConfirmation":
		link := authState.AuthorizationState.(*AuthorizationStateWaitOtherDeviceConfirmation).Link
		fmt.Printf("Scan the QR code below: or open the link in Telegram: %s\n", link)
		q, err := qrcode.NewQRCode(link)
		if err != nil {
			fmt.Printf("Failed to generate QR: %v\nLink: %s\n", err, link)
		} else {
			fmt.Println(q.ToSmallString(false))
		}

	case "authorizationStateWaitCode":
		codeInfo := authState.AuthorizationState.(*AuthorizationStateWaitCode).CodeInfo
		codeType := codeInfo.TypeField.Type()
		codeType = strings.TrimPrefix(codeType, "authenticationCodeType")
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

	case "authorizationStateWaitPassword":
		hint := authState.AuthorizationState.(*AuthorizationStateWaitPassword).PasswordHint
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

	case "authorizationStateWaitRegistration":
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
			c.authErrorChan <- err
		}
	case "authorizationStateWaitPremiumPurchase":
		c.Logger.Info("Account is limited and requires Telegram Premium. Please purchase Telegram Premium to continue.")
		c.authErrorChan <- WaitPremiumPurchase
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
				return nil, errObj
			}
		}
		return res, nil
	case <-time.After(30 * time.Second):
		c.pendingRequests.Delete(extra)
		return nil, SendTimeout
	}
}

// Me returns the cached User object for the current bot.
func (c *Client) Me() *User {
	c.meMu.RLock()
	defer c.meMu.RUnlock()
	return c.me
}

// WaitMessage waits for the message to be sent and returns the final message.
func (c *Client) WaitMessage(msg *Message) (*Message, error) {
	if msg.SendingState != nil && msg.SendingState.Type() == "messageSendingStatePending" {
		key := fmt.Sprintf("%d:%d", msg.ChatId, msg.Id)
		ch := make(chan TlObject, 1)
		c.pendingMessages.Store(key, ch)
		defer c.pendingMessages.Delete(key)

		select {
		case res := <-ch:
			if errObj, ok := res.(*Error); ok {
				return nil, errObj
			}
			if finalMsg, ok := res.(*Message); ok {
				return finalMsg, nil
			}
			return nil, fmt.Errorf("unexpected response type from waiter: %T", res)
		case <-time.After(30 * time.Second):
			return msg, nil
		}
	}

	return msg, nil
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
