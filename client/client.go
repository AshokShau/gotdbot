package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/AshokShau/gotdbot/tdjson"
	"github.com/AshokShau/gotdbot/types"
)

var StopHandlers = fmt.Errorf("stop handlers")

type ClientConfig struct {
	LibraryPath           string
	UseTestDC             bool
	DatabaseDirectory     string
	FilesDirectory        string
	DatabaseEncryptionKey string
	UseFileDatabase       bool
	UseChatInfoDatabase   bool
	UseMessageDatabase    bool
	UseSecretChats        bool
	SystemLanguageCode    string
	DeviceModel           string
	SystemVersion         string
	ApplicationVersion    string
}

func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		DatabaseDirectory:   "tdlib_db",
		FilesDirectory:      "tdlib_files",
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      true,
		SystemLanguageCode:  "en",
		DeviceModel:         "Gotdbot",
		SystemVersion:       "1.0",
		ApplicationVersion:  "1.0",
	}
}

type Client struct {
	clientID int

	apiID    int32
	apiHash  string
	botToken string
	config   *ClientConfig

	updates chan types.TlObject
	stop    chan struct{}
	wg      sync.WaitGroup

	handlers     map[string][]*Handler
	initializers []*Handler
	finalizers   []*Handler
	hMu          sync.RWMutex

	pendingRequests sync.Map // map[string]chan types.TlObject
}

func NewClient(apiID int32, apiHash, botToken string, config *ClientConfig) *Client {
	if config == nil {
		config = DefaultClientConfig()
	} else {
		def := DefaultClientConfig()
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
	}

	if err := tdjson.Init(config.LibraryPath); err != nil {
		log.Fatalf("Failed to initialize TDLib: %v", err)
	}

	c := &Client{
		clientID: tdjson.CreateClientID(),
		apiID:    apiID,
		apiHash:  apiHash,
		botToken: botToken,
		config:   config,
		updates:  make(chan types.TlObject, 1000),
		stop:     make(chan struct{}),
		handlers: make(map[string][]*Handler),
	}

	c.AddHandler("updateAuthorizationState", c.authHandler, nil, 0)
	tdjson.Send(c.clientID, `{"@type": "getOption", "name": "version"}`)

	return c
}

// SetLogVerbosityLevel sets the verbosity level of TDLib's internal logging.
func SetLogVerbosityLevel(level int32) {
	req := fmt.Sprintf(`{"@type": "setLogVerbosityLevel", "new_verbosity_level": %d}`, level)
	tdjson.Execute(req)
}

// SetLogStreamFile redirects internal TDLib logs to a file.
func SetLogStreamFile(path string, maxFileSize int64, redirectStderr bool) {
	req := fmt.Sprintf(`{"@type": "setLogStream", "log_stream": {"@type": "logStreamFile", "path": "%s", "max_file_size": %d, "redirect_stderr": %v}}`, path, maxFileSize, redirectStderr)
	tdjson.Execute(req)
}

// SetLogStreamEmpty disables internal TDLib logging.
func SetLogStreamEmpty() {
	req := `{"@type": "setLogStream", "log_stream": {"@type": "logStreamEmpty"}}`
	tdjson.Execute(req)
}

func (c *Client) Run() {
	c.wg.Add(1)
	go c.receiver()

	c.wg.Add(1)
	go c.processor()
}

func (c *Client) Stop() {
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
			// Receive with short timeout
			res := tdjson.Receive(0.1)
			if res != "" {
				data := []byte(res)
				obj, err := types.Unmarshal(data)
				if err != nil {
					// fmt.Println("Error unmarshaling:", err)
					continue
				}
				if obj == nil {
					continue
				}

				extra := obj.GetExtra()
				if extra != "" {
					if ch, ok := c.pendingRequests.Load(extra); ok {
						ch.(chan types.TlObject) <- obj
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

func (c *Client) handleUpdate(update types.TlObject) {
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

func (c *Client) authHandler(client *Client, update types.TlObject) error {
	authState, ok := update.(*types.UpdateAuthorizationState)
	if !ok {
		return nil
	}

	switch authState.AuthorizationState.Type() {
	case "authorizationStateWaitTdlibParameters":
		_, err := c.SetTdlibParameters(
			c.config.UseTestDC,
			c.config.DatabaseDirectory,
			c.config.FilesDirectory,
			c.config.DatabaseEncryptionKey,
			c.config.UseFileDatabase,
			c.config.UseChatInfoDatabase,
			c.config.UseMessageDatabase,
			c.config.UseSecretChats,
			c.apiID,
			c.apiHash,
			c.config.SystemLanguageCode,
			c.config.DeviceModel,
			c.config.SystemVersion,
			c.config.ApplicationVersion,
		)
		if err != nil {
			log.Printf("Error setting tdlib parameters: %v", err)
		}

	case "authorizationStateWaitPhoneNumber":
		_, err := c.CheckAuthenticationBotToken(c.botToken)
		if err != nil {
			log.Printf("Error checking bot token: %v", err)
		}
	}
	return nil
}

func (c *Client) Send(req types.TlObject) (types.TlObject, error) {
	extra := fmt.Sprintf("%d", time.Now().UnixNano())
	req.SetExtra(extra)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan types.TlObject, 1)
	c.pendingRequests.Store(extra, ch)

	tdjson.Send(c.clientID, string(data))

	select {
	case res := <-ch:
		// check for error
		if res.Type() == "error" {
			if errObj, ok := res.(*types.Error); ok {
				return nil, fmt.Errorf("TDLib error %d: %s", errObj.Code, errObj.Message)
			}
		}
		return res, nil
	case <-time.After(30 * time.Second): // Timeout
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
