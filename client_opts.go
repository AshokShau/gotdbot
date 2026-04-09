package gotdbot

import (
	"log/slog"
	"os"
	"runtime"
	"time"
)

type AutoRetry struct {
	ChatNotFound    bool
	MessageNotFound bool
}

type ClientOpts struct {
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
	AutoRetry               *AutoRetry

	// Dispatcher is the dispatcher to use for this client.
	// If nil, a new dispatcher will be created.
	Dispatcher *Dispatcher
}

func DefaultClientConfig() *ClientOpts {
	return &ClientOpts{
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
		AutoRetry:               &AutoRetry{},
	}
}
