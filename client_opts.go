package gotdbot

import (
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/AshokShau/gotdbot/logger"
)

// Bool returns a pointer for options that distinguish an unset value from false.
func Bool(b bool) *bool {
	return &b
}

// AutoRetry controls missing-object lookups and the maximum automatic flood wait.
type AutoRetry struct {
	ChatNotFound    bool
	MessageNotFound bool
	MaxFloodWait    time.Duration
}

// ClientOpts configures account storage, authentication, logging, and update handling.
type ClientOpts struct {
	PanicHandler func(client *Client, update TlObject, r any)
	ErrorHandler func(client *Client, update TlObject, err error) error

	LibraryPath             string
	UseTestDC               bool
	DatabaseDirectory       string
	FilesDirectory          string
	DatabaseEncryptionKey   string
	UseFileDatabase         *bool
	UseChatInfoDatabase     *bool
	UseMessageDatabase      *bool
	UseSecretChats          bool
	LoadMessagesBeforeReply bool
	SystemLanguageCode      string
	DeviceModel             string
	SystemVersion           string
	ApplicationVersion      string
	TDLibOptions            *TDLibOptions
	Logger                  *logger.Logger
	QrMode                  bool
	AuthorizationTimeout    time.Duration
	// RequestTimeout bounds each request and its send completion; default 30s.
	RequestTimeout time.Duration
	// UpdateWorkers bounds concurrent user handlers; default 2. Use 1 for order.
	UpdateWorkers int
	// UpdateQueueSize bounds updates buffered in memory per client; default 1000.
	// Excess updates spill to temporary files, reclaimed as the queue drains.
	UpdateQueueSize int
	// ReceiveOldMessages disables historical-message filtering for bots.
	ReceiveOldMessages bool
	// DropPendingUpdates skips startup updates until auth and connection are ready.
	DropPendingUpdates bool
	LogVerbosityLevel  int32
	LogStream          LogStream
	AutoRetry          *AutoRetry
	ParseMode          string
}

// DefaultClientConfig returns a fresh configuration with the library defaults.
func DefaultClientConfig() *ClientOpts {
	return &ClientOpts{
		DatabaseDirectory:       "database",
		FilesDirectory:          "",
		DatabaseEncryptionKey:   "",
		UseFileDatabase:         Bool(true),
		UseChatInfoDatabase:     Bool(true),
		UseMessageDatabase:      Bool(true),
		UseSecretChats:          false,
		LoadMessagesBeforeReply: false,
		SystemLanguageCode:      "en",
		DeviceModel:             "Gotdbot",
		SystemVersion:           runtime.GOOS,
		ApplicationVersion:      "Gotdbot " + Version,
		Logger:                  logger.New(logger.WithLevel(slog.LevelInfo), logger.WithOutput(os.Stdout)),
		QrMode:                  false,
		AuthorizationTimeout:    60 * time.Second,
		RequestTimeout:          30 * time.Second,
		UpdateWorkers:           2,
		UpdateQueueSize:         1000,
		LogVerbosityLevel:       2,
		AutoRetry:               &AutoRetry{},
		ParseMode:               "",
	}
}
