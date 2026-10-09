package gotdbot

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestDefaultClientConfigIndependent(t *testing.T) {
	a, b := DefaultClientConfig(), DefaultClientConfig()
	if a.RequestTimeout != 30*time.Second || a.AuthorizationTimeout != time.Minute || a.UpdateWorkers != 2 || a.UpdateQueueSize != 1000 || a.Logger == nil || a.AutoRetry == nil {
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

func TestTDLibOptionsEnumeratesSetValues(t *testing.T) {
	var nilOptions *TDLibOptions
	nilOptions.forEachSet(func(string, any) { t.Fatal("nil options emitted value") })
	(&TDLibOptions{}).forEachSet(func(string, any) { t.Fatal("default options emitted value") })
	opts := &TDLibOptions{}
	value := reflect.ValueOf(opts).Elem()
	typ := value.Type()
	expected := make(map[string]any)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			field.SetBool(true)
		case reflect.String:
			field.SetString("value")
		case reflect.Int64:
			field.SetInt(-42)
		default:
			t.Fatalf("uncovered option kind: %v", field.Kind())
		}
		name := typ.Field(i).Tag.Get("json")
		for j, c := range name {
			if c == ',' {
				name = name[:j]
				break
			}
		}
		expected[name] = field.Interface()
	}
	got := make(map[string]any)
	opts.forEachSet(func(name string, v any) {
		if _, exists := got[name]; exists {
			t.Errorf("duplicate option %s", name)
		}
		got[name] = v
	})
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("options = %v, want %v", got, expected)
	}
	opts.forEachSet(nil)
}

func TestToOptionValue(t *testing.T) {
	for _, tc := range []struct {
		input any
		want  OptionValue
	}{
		{true, &OptionValueBoolean{Value: true}}, {false, &OptionValueBoolean{}},
		{int(-1), &OptionValueInteger{Value: -1}}, {int32(42), &OptionValueInteger{Value: 42}}, {int64(math.MaxInt64), &OptionValueInteger{Value: math.MaxInt64}},
		{"value", &OptionValueString{Value: "value"}}, {nil, &OptionValueEmpty{}}, {1.5, nil}, {[]byte("x"), nil},
	} {
		if got := toOptionValue(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("toOptionValue(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

type testRawIDRequest string

func (r testRawIDRequest) GetType() string              { return "test" }
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
	// Ensure the fixture also observes encoding/json's validation of MarshalJSON.
	if _, err := json.Marshal(testRawIDRequest(`invalid`)); err == nil {
		t.Fatal("invalid fixture unexpectedly serialized")
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
