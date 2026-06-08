package gotdbot

//go:generate go run ./internal/gen

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"os/signal"
	"regexp"
	"sort"
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
	manager  *ClientManager
	clientID int

	apiID    int32
	apiHash  string
	botToken string
	// phoneNumber is set if the user is logging in with a phone number.
	phoneNumber string
	config      *ClientOpts
	Logger      *logger.Logger

	updates chan TlObject
	stop    chan struct{}
	closed  chan struct{}
	wg      sync.WaitGroup

	handlers     atomic.Pointer[map[int][]Handler]
	hMu          sync.Mutex
	panicHandler func(client *Client, update TlObject, r interface{})
	errorHandler func(client *Client, update TlObject, err error) error

	pendingRequests sync.Map // map[string]chan TlObject
	pendingMessages sync.Map // map[string]chan TlObject

	waiters     map[string]*Waiter
	waiterCount int64
	wMu         sync.RWMutex

	// Auth state management
	authErrorChan chan error
	isAuthorized  bool
	startOnce     sync.Once
	closeOnce     sync.Once

	// Me is the bot's User info, as returned by client.GetMe.
	// Populated when authorization is ready.
	Me *User
}

func NewClient(apiID int32, apiHash, tokenOrPhone string, config *ClientOpts) (*Client, error) {
	tokenOrPhone = strings.TrimSpace(tokenOrPhone)
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

		if config.AutoRetry == nil {
			config.AutoRetry = def.AutoRetry
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

	handlers := make(map[int][]Handler)
	c.handlers.Store(&handlers)
	c.waiters = make(map[string]*Waiter)
	c.panicHandler = config.PanicHandler
	c.errorHandler = config.ErrorHandler

	c.AddUpdateAuthorizationStateHandlerGroup(c.authHandler, nil, -100)
	c.AddUpdateMessageSendSucceededHandlerGroup(c.messageSendSucceededHandler, nil, -99)
	c.AddUpdateMessageSendFailedHandlerGroup(c.messageSendFailedHandler, nil, -98)
	c.AddUpdateConnectionStateHandlerGroup(c.connectionStateHandler, nil, -97)
	return c, nil

}

// initClient initializes the client's internal state and processor loop exactly once.
func (c *Client) initClient() {
	tdjson.Send(c.clientID, `{"@type": "getOption", "name": "version"}`)

	c.startOnce.Do(func() {
		if c.manager == nil {
			m := GetDefaultManager(c.config.LibraryPath)
			m.AddClient(c)
		}

		c.wg.Add(1)
		go c.processor()
	})
}

// Start initializes the client and blocks until authorization is successful or fails.
func (c *Client) Start() error {
	c.initClient()

	select {
	case err := <-c.authErrorChan:
		return err
	case <-time.After(c.config.AuthorizationTimeout):
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
	<-sig
	c.Close()
}

// Debug logs a message at debug level.
func (c *Client) Debug(msg string, args ...any) { c.Logger.Debug(msg, args...) }

// Info logs a message at info level.
func (c *Client) Info(msg string, args ...any) { c.Logger.Info(msg, args...) }

// Warn logs a message at warn level.
func (c *Client) Warn(msg string, args ...any) { c.Logger.Warn(msg, args...) }

// Error logs a message at error level.
func (c *Client) Error(msg string, args ...any) { c.Logger.Error(msg, args...) }

// Debugf logs a formatted message at debug level.
func (c *Client) Debugf(format string, args ...any) { c.Logger.Debugf(format, args...) }

// Infof logs a formatted message at info level.
func (c *Client) Infof(format string, args ...any) { c.Logger.Infof(format, args...) }

// Warnf logs a formatted message at warn level.
func (c *Client) Warnf(format string, args ...any) { c.Logger.Warnf(format, args...) }

// Errorf logs a formatted message at error level.
func (c *Client) Errorf(format string, args ...any) { c.Logger.Errorf(format, args...) }

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (c *Client) Close() {
	c.closeOnce.Do(func() {
		_, _ = c.Send(&Close{})

		select {
		case <-c.closed:
		case <-time.After(5 * time.Second):
			c.Warn("Timeout waiting for TDLib to close")
		}

		close(c.stop)
		c.wg.Wait()
	})
}
func (c *Client) processor() {
	defer c.wg.Done()
	for {
		select {
		case <-c.stop:
			return
		case update := <-c.updates:
			updateType := update.GetType()

			go func() {
				defer func() {
					if r := recover(); r != nil {
						if c.panicHandler != nil {
							c.panicHandler(c, update, r)
						} else {
							c.Error("Panic in handler", "panic", r)
						}
					}
				}()

				c.wMu.RLock()
				if len(c.waiters) > 0 {
					var matchedWaiters []*Waiter
					for _, w := range c.waiters {
						if w.Filter(c, update) {
							matchedWaiters = append(matchedWaiters, w)
						}
					}
					c.wMu.RUnlock()

					for _, w := range matchedWaiters {
						select {
						case w.C <- update:
						default:
						}
					}
				} else {
					c.wMu.RUnlock()
				}

				handlersMap := *c.handlers.Load()
				groups := make([]int, 0, len(handlersMap))
				for k := range handlersMap {
					groups = append(groups, k)
				}
				sort.Ints(groups)

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
					c.Error("Handler error", "error", err, "type", updateType)
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
								continue // Move to next handler in the same group
							}
							// Handler matched and succeeded, move to next group
							goto nextGroup
						}
					}
				nextGroup:
				}
			}()
		}
	}
}

func (c *Client) authHandler(client *Client, authState *UpdateAuthorizationState) error {
	c.Debug("Authorization state update", "state", authState.AuthorizationState.GetType())

	switch authState.AuthorizationState.GetType() {
	case "authorizationStateWaitTdlibParameters":
		if c.config.TDLibOptions != nil {
			c.config.TDLibOptions.forEachSet(func(k string, v interface{}) {
				if opt := toOptionValue(v); opt != nil {
					err := c.SetOption(k, &SetOptionOpts{Value: opt})
					if err != nil {
						c.Error("Error setting option", "option", k, "error", err)
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
			c.Error("Error setting tdlib parameters", "error", err)
			c.authErrorChan <- err
		}

	case "authorizationStateWaitPhoneNumber":
		if c.botToken != "" {
			err := c.CheckAuthenticationBotToken(c.botToken)
			if err != nil {
				c.Error("Error checking bot token", "error", err)
				c.authErrorChan <- err
			}
		} else {
			// User login
			if c.config.QrMode {
				err := c.RequestQrCodeAuthentication(nil)
				if err != nil {
					c.Error("Error requesting QR code", "error", err)
					c.authErrorChan <- err
				}
			} else if c.phoneNumber != "" {
				err := c.SetAuthenticationPhoneNumber(c.phoneNumber, nil)
				if err != nil {
					c.Error("Error setting phone number", "error", err)
					c.authErrorChan <- err
				}
			} else {
				fmt.Print("Enter phone number: ")
				reader := bufio.NewReader(os.Stdin)
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)
				err := c.SetAuthenticationPhoneNumber(phone, nil)
				if err != nil {
					c.Error("Error setting phone number", "error", err)
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
		codeType := codeInfo.Type.GetType()
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
			c.Error("Error registering user", "error", err)
			c.authErrorChan <- err
		}
	case "authorizationStateWaitPremiumPurchase":
		c.Info("Account is limited and requires Telegram Premium. Please purchase Telegram Premium to continue.")
		c.authErrorChan <- WaitPremiumPurchase
	case "authorizationStateReady":
		c.isAuthorized = true
		me, err := c.GetMe()
		if err != nil {
			c.Error("Failed to get me", "error", err)
			c.authErrorChan <- err
			return nil
		}

		c.Me = me

		username := ""
		if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
			username = me.Usernames.ActiveUsernames[0]
		}
		c.Info("Logged in", "user_id", me.Id, "username", username)

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

func (c *Client) connectionStateHandler(client *Client, u *UpdateConnectionState) error {
	state := u.State.GetType()
	state = strings.TrimPrefix(state, "connectionState")
	c.Info("Connection state changed", "state", state)
	return nil
}

func (c *Client) messageSendSucceededHandler(client *Client, u *UpdateMessageSendSucceeded) error {
	key := fmt.Sprintf("%d:%d", u.Message.ChatId, u.OldMessageId)
	if ch, ok := c.pendingMessages.Load(key); ok {
		ch.(chan TlObject) <- u
		c.pendingMessages.Delete(key)
	}

	return nil
}

func (c *Client) messageSendFailedHandler(client *Client, u *UpdateMessageSendFailed) error {
	key := fmt.Sprintf("%d:%d", u.Message.ChatId, u.OldMessageId)
	if ch, ok := c.pendingMessages.Load(key); ok {
		ch.(chan TlObject) <- u
		c.pendingMessages.Delete(key)
	}
	return nil
}

func (c *Client) Send(req TlObject) (TlObject, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return nil, err
	}

	reqType, _ := tmp["@type"].(string)
	reqType = strings.ToLower(reqType)

	isChatAttemptedLoad := reqType == "getchat"
	isMessageAttemptedLoad := reqType == "getmessage" || reqType == "getmessagelocally" || reqType == "getrepliedmessage" || reqType == "getcallbackquerymessage"

	for {
		extra := fmt.Sprintf("%d", time.Now().UnixNano())
		tmp["@extra"] = extra

		sendData, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}

		ch := make(chan TlObject, 1)
		c.pendingRequests.Store(extra, ch)

		tdjson.Send(c.clientID, string(sendData))

		var result TlObject
		var resultErr error

		select {
		case res := <-ch:
			if res.GetType() == "error" {
				if errObj, ok := res.(*Error); ok {
					resultErr = errObj
					result = res
				}
			} else {
				result = res
			}
		case <-time.After(30 * time.Second):
			c.pendingRequests.Delete(extra)
			return nil, SendTimeout
		}

		if resultErr == nil {
			return result, nil
		}

		errObj := resultErr.(*Error)

		if c.handleAutoRetry(tmp, errObj, &isChatAttemptedLoad, &isMessageAttemptedLoad) {
			continue
		}

		return nil, resultErr
	}
}

func (c *Client) handleAutoRetry(req map[string]interface{}, errObj *Error, isChatAttemptedLoad, isMessageAttemptedLoad *bool) bool {
	if c.config.AutoRetry == nil {
		return false
	}

	if errObj.Code == 429 && c.config.AutoRetry.MaxFloodWait > 0 {
		prefix := "Too Many Requests: retry after "
		if strings.HasPrefix(errObj.Message, prefix) {
			seconds, err := strconv.Atoi(strings.TrimPrefix(errObj.Message, prefix))
			if err == nil {
				waitTime := time.Duration(seconds) * time.Second
				if waitTime <= c.config.AutoRetry.MaxFloodWait {
					c.Info("Flood wait", "seconds", seconds)
					time.Sleep(waitTime)
					return true
				}
			}
		}
	}

	if errObj.Code != 400 {
		return false
	}

	if !*isMessageAttemptedLoad && errObj.Message == "Message not found" && c.config.AutoRetry.MessageNotFound {
		*isMessageAttemptedLoad = true
		var chatId, messageId int64
		if cId, ok := req["chat_id"].(float64); ok {
			chatId = int64(cId)
		}
		if mId, ok := req["message_id"].(float64); ok {
			messageId = int64(mId)
		}
		if chatId != 0 && messageId != 0 {
			c.Debug("Attempting to load message", "chat_id", chatId, "message_id", messageId)
			msg, loadErr := c.GetMessage(chatId, messageId)
			if loadErr == nil && msg != nil {
				c.Debug("Successfully loaded message, retrying request", "chat_id", chatId, "message_id", messageId)
				return true
			}

			c.Debug("Failed to load message", "chat_id", chatId, "message_id", messageId, "error", loadErr)
		}
	}

	if !*isChatAttemptedLoad && errObj.Message == "Chat not found" && c.config.AutoRetry.ChatNotFound {
		*isChatAttemptedLoad = true
		var chatId int64
		if cId, ok := req["chat_id"].(float64); ok {
			chatId = int64(cId)
		}
		if chatId != 0 {
			c.Debug("Attempting to load chat", "chat_id", chatId)
			chat, loadErr := c.GetChat(chatId)
			if loadErr == nil && chat != nil {
				c.Debug("Successfully loaded chat, retrying request", "chat_id", chatId)
				if replyToMap, ok := req["reply_to"].(map[string]interface{}); ok {
					if mId, ok := replyToMap["message_id"].(float64); ok {
						replyToMessageId := int64(mId)
						if replyToMessageId > 0 {
							_, _ = c.GetMessage(chatId, replyToMessageId)
						}
					}
				}

				return true
			}

			c.Debug("Failed to load chat", "chat_id", chatId, "error", loadErr)
		}
	}
	return false
}

// waitMessage waits for the message to be sent and returns the final message.
func (c *Client) waitMessage(msg *Message) (*Message, error) {
	if msg.SendingState != nil && msg.SendingState.GetType() == "messageSendingStatePending" {
		key := fmt.Sprintf("%d:%d", msg.ChatId, msg.Id)
		ch := make(chan TlObject, 1)
		c.pendingMessages.Store(key, ch)
		defer c.pendingMessages.Delete(key)

		select {
		case res := <-ch:
			if errObj, ok := res.(*Error); ok {
				return nil, errObj
			}
			if u, ok := res.(*UpdateMessageSendFailed); ok {
				return u.Message, u.Error
			}
			if finalMsg, ok := res.(*Message); ok {
				return finalMsg, nil
			}
			if u, ok := res.(*UpdateMessageSendSucceeded); ok {
				return u.Message, nil
			}
			return nil, fmt.Errorf("unexpected response type from waiter: %T", res)
		case <-time.After(30 * time.Second):
			return msg, nil
		}
	}

	return msg, nil
}

// waitMessages waits for all messages in the album to be sent and returns the final messages.
func (c *Client) waitMessages(msgs *Messages) (*Messages, error) {
	if msgs == nil || len(msgs.Messages) == 0 {
		return msgs, nil
	}

	totalCount := len(msgs.Messages)
	ch := make(chan TlObject, totalCount)
	errs := make([]error, totalCount)

	for i := range msgs.Messages {
		msg := &msgs.Messages[i]
		if msg.SendingState != nil && msg.SendingState.GetType() == "messageSendingStatePending" {
			key := fmt.Sprintf("%d:%d", msg.ChatId, msg.Id)
			c.pendingMessages.Store(key, ch)
		} else {
			ch <- msg
		}
	}

	defer func() {
		for i := range msgs.Messages {
			msg := &msgs.Messages[i]
			key := fmt.Sprintf("%d:%d", msg.ChatId, msg.Id)
			c.pendingMessages.Delete(key)
		}
	}()

	var receivedCount int
	timeout := time.After(30 * time.Second)

	for receivedCount < totalCount {
		select {
		case res := <-ch:
			if errObj, ok := res.(*Error); ok {
				errs[receivedCount] = errObj
			} else if u, ok := res.(*UpdateMessageSendFailed); ok {
				errs[receivedCount] = u.Error
				for i := range msgs.Messages {
					if msgs.Messages[i].Id == u.OldMessageId {
						msgs.Messages[i] = *u.Message
						break
					}
				}
			} else if finalMsg, ok := res.(*Message); ok {
				for i := range msgs.Messages {
					if msgs.Messages[i].Id == finalMsg.Id {
						msgs.Messages[i] = *finalMsg
						break
					}
				}
			} else if u, ok := res.(*UpdateMessageSendSucceeded); ok {
				for i := range msgs.Messages {
					if msgs.Messages[i].Id == u.OldMessageId {
						msgs.Messages[i] = *u.Message
						break
					}
				}
			}
			receivedCount++
		case <-timeout:
			return msgs, errors.Join(errs...)
		}
	}

	return msgs, errors.Join(errs...)
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

type Waiter struct {
	Filter func(client *Client, update TlObject) bool
	C      chan TlObject
}

// WaitFor registers a waiter for a specific client and blocks until a matching update arrives or timeout occurs.
func (c *Client) WaitFor(filter func(client *Client, update TlObject) bool, timeout time.Duration) (TlObject, error) {
	ch := make(chan TlObject, 1)
	idNum := atomic.AddInt64(&c.waiterCount, 1)
	id := fmt.Sprintf("%d", idNum)

	c.wMu.Lock()
	c.waiters[id] = &Waiter{Filter: filter, C: ch}
	c.wMu.Unlock()

	defer func() {
		c.wMu.Lock()
		delete(c.waiters, id)
		c.wMu.Unlock()
	}()

	select {
	case u := <-ch:
		return u, nil
	case <-time.After(timeout):
		return nil, ConversationTimeout
	}
}
