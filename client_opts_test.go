package gotdbot

import (
	"encoding/json"
	"math"
	"testing"
	"time"
)

func TestDefaultClientConfigIndependent(t *testing.T) {
	a, b := DefaultClientConfig(), DefaultClientConfig()
	if a.RequestTimeout != 0 || a.AuthorizationTimeout != time.Minute || a.UpdateWorkers != 2 || a.UpdateQueueSize != 1000 || a.Logger == nil || a.AutoRetry == nil {
		t.Fatalf("unexpected defaults: %+v", a)
	}
	for _, enabled := range []*bool{a.UseFileDatabase, a.UseChatInfoDatabase, a.UseMessageDatabase} {
		if enabled == nil || !*enabled {
			t.Fatal("database disabled by default")
		}
	}
	*a.UseFileDatabase = false
	a.AutoRetry.ChatNotFound = true
	if !*b.UseFileDatabase || b.AutoRetry.ChatNotFound {
		t.Fatal("default configuration shares mutable state")
	}
}

type testRawIDRequest string

func (r testRawIDRequest) GetType() string { return "test" }

func (r testRawIDRequest) MarshalJSON() ([]byte, error) { return []byte(r), nil }

func TestExtractIDsPreservesPrecision(t *testing.T) {
	for _, req := range []TlObject{
		&GetMessage{ChatId: 9007199254740993, MessageId: math.MaxInt64},
		testRawIDRequest(`{"chat_id":"9007199254740993","message_id":"9223372036854775807"}`),
	} {
		chat, message := extractChatAndMessageID(req)
		if chat != 9007199254740993 || message != math.MaxInt64 {
			t.Fatalf("IDs = %d/%d", chat, message)
		}
	}
	for _, req := range []TlObject{&GetMe{}, testRawIDRequest(`{"chat_id":1.5}`), testRawIDRequest(`{"chat_id":true}`), testRawIDRequest(`invalid`)} {
		chat, message := extractChatAndMessageID(req)
		if chat != 0 || message != 0 {
			t.Fatalf("invalid IDs = %d/%d", chat, message)
		}
	}
}

func TestErrorRetryAfter(t *testing.T) {
	for _, tc := range []struct {
		code    int32
		message string
		want    int
	}{
		{429, "Too Many Requests: retry after 12", 12}, {400, "Too Many Requests: retry after 12", 0},
		{429, "Too Many Requests: retry after x", 0}, {429, "retry after 12", 0}, {429, "Too Many Requests: retry after 999999999999999999999999", 0},
	} {
		err := &Error{Code: tc.code, Message: tc.message}
		if got := err.GetRetryAfter(); got != tc.want {
			t.Errorf("retry after %q = %d", tc.message, got)
		}
	}
	if got := (&Error{Code: 400, Message: "bad request"}).Error(); got != "TDLib error 400: bad request" {
		t.Fatal(got)
	}
}

func TestAuthenticationInitializationPreservesStringOptions(t *testing.T) {
	c := newTestClient(t)
	c.apiID, c.apiHash = 123, "api-hash"
	c.config.DatabaseDirectory = "custom/database"
	c.config.FilesDirectory = "custom/files"
	c.config.SystemVersion = "custom system"
	c.config.ApplicationVersion = "app version"
	c.config.DeviceModel = "device model"
	c.config.SystemLanguageCode = "en"
	var request *SetTdlibParameters
	c.sendRequest = func(_ int, data []byte) {
		wire := decodeTestRequest(t, data)
		if wire.Type != "setTdlibParameters" {
			t.Fatalf("unexpected request: %s", wire.Type)
		}
		if err := json.Unmarshal(data, &request); err != nil {
			t.Fatal(err)
		}
		c.deliverResponse(wire.Extra, &Ok{})
	}
	c.authHandler(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateWaitTdlibParameters{}})
	if c.authErr != nil || request == nil || request.ApiId != 123 || request.ApiHash != "api-hash" || request.DatabaseDirectory != c.config.DatabaseDirectory || request.FilesDirectory != c.config.FilesDirectory || request.SystemVersion != c.config.SystemVersion || request.ApplicationVersion != c.config.ApplicationVersion || request.DeviceModel != c.config.DeviceModel || request.SystemLanguageCode != "en" || request.UseChatInfoDatabase != *c.config.UseChatInfoDatabase || request.UseFileDatabase != *c.config.UseFileDatabase {
		t.Fatalf("initialization lost config: %+v, %v", request, c.authErr)
	}
}
