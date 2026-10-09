package gotdbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSendWithContextValidation(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(int, []byte) { t.Error("invalid request reached transport") }
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	var nilRequest *GetMe
	for _, tc := range []struct {
		name string
		ctx  context.Context
		req  TlObject
		want error
	}{
		{"nil context", nil, &GetMe{}, nil},
		{"nil request", context.Background(), nil, nil},
		{"typed nil", context.Background(), nilRequest, nil},
		{"object", context.Background(), &Ok{}, nil},
		{"function value", context.Background(), GetMe{}, nil},
		{"cancelled", cancelled, &GetMe{}, context.Canceled},
	} {
		t.Run(tc.name, func(t *testing.T) {
			obj, err := c.SendWithContext(tc.ctx, tc.req)
			if obj != nil || err == nil || (tc.want != nil && !errors.Is(err, tc.want)) {
				t.Fatalf("SendWithContext = (%v, %v), want error %v", obj, err, tc.want)
			}
			assertNoPendingCalls(t, c)
		})
	}
	c.closing.Store(true)
	if _, err := c.Send(&GetMe{}); !errors.Is(err, ErrClientClosed) {
		t.Fatalf("closed client: %v", err)
	}
}

func TestSendConcurrentSharedRequestCorrelation(t *testing.T) {
	c := newTestClient(t)
	const count = 64
	requests := make(chan testWireRequest, count)
	c.sendRequest = func(id int, data []byte) {
		if id != c.clientID {
			t.Errorf("client ID = %d", id)
		}
		requests <- decodeTestRequest(t, data)
	}
	req := &GetChat{ChatId: 9007199254740993}
	before, _ := json.Marshal(req)
	type result struct {
		obj TlObject
		err error
	}
	results := make(chan result, count)
	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() { defer wg.Done(); obj, err := c.Send(req); results <- result{obj, err} }()
	}
	seen := make(map[string]bool)
	wires := make([]testWireRequest, count)
	for i := range wires {
		wires[i] = receiveTest(t, requests)
		w := wires[i]
		if w.Type != "getChat" || w.ChatID != req.ChatId || w.Extra == "" || seen[w.Extra] {
			t.Fatalf("bad correlation or precision: %+v", w)
		}
		seen[w.Extra] = true
	}
	// Deliver in reverse order to catch responses accidentally shared by calls.
	for i := len(wires) - 1; i >= 0; i-- {
		c.acceptResponse(&Chat{Id: int64(i + 1)}, wires[i].Extra, nil)
	}
	ids := make(map[int64]bool)
	for range count {
		r := receiveTest(t, results)
		chat, ok := r.obj.(*Chat)
		if r.err != nil || !ok || ids[chat.Id] {
			t.Fatalf("lost/duplicate response: %+v", r)
		}
		ids[chat.Id] = true
	}
	wg.Wait()
	after, _ := json.Marshal(req)
	if string(before) != string(after) {
		t.Fatal("Send mutated shared request")
	}
	assertNoPendingCalls(t, c)
}

func TestSendAbortCleansCorrelation(t *testing.T) {
	for _, reason := range []string{"cancel", "timeout", "closed", "stop"} {
		t.Run(reason, func(t *testing.T) {
			c := newTestClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if reason == "timeout" {
				c.config.RequestTimeout = 20 * time.Millisecond
			}
			requests := make(chan testWireRequest, 1)
			c.sendRequest = func(_ int, data []byte) { requests <- decodeTestRequest(t, data) }
			done := make(chan error, 1)
			go func() { _, err := c.SendWithContext(ctx, &GetMe{}); done <- err }()
			wire := receiveTest(t, requests)
			want := ErrClientClosed
			switch reason {
			case "cancel":
				cancel()
				want = context.Canceled
			case "timeout":
				want = context.DeadlineExceeded
			case "closed":
				c.closedOnce.Do(func() { close(c.closed) })
			case "stop":
				c.stopOnce.Do(func() { close(c.stop) })
			}
			err := receiveTest(t, done)
			if !errors.Is(err, want) {
				t.Fatalf("error = %v, want %v", err, want)
			}
			if reason == "timeout" && !errors.Is(err, SendTimeout) {
				t.Fatalf("missing SendTimeout: %v", err)
			}
			assertNoPendingCalls(t, c)
			c.deliverResponse(wire.Extra, &User{Id: 1}) // Late replies must be harmless.
			assertNoPendingCalls(t, c)
			if c.requestID.Load() != 1 {
				t.Fatal("transport abort retried a request")
			}
		})
	}
}

func TestSendServerErrorAndWrongResponseType(t *testing.T) {
	c := newTestClient(t)
	serverError := &Error{Code: 403, Message: "Forbidden"}
	c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, serverError) }
	if obj, err := c.Send(&GetMe{}); obj != nil || err != serverError {
		t.Fatalf("server error = (%v, %v)", obj, err)
	}
	c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, &Ok{}) }
	if _, err := c.GetMe(); err == nil {
		t.Fatal("typed method accepted wrong response")
	}
	assertNoPendingCalls(t, c)
}

func TestSendLookupRetries(t *testing.T) {
	for _, tc := range []struct {
		name, rejection, lookup string
		policy                  *AutoRetry
	}{
		{"chat", "Chat not found", "getChat", &AutoRetry{ChatNotFound: true}},
		{"message", "Message not found", "getMessage", &AutoRetry{MessageNotFound: true}},
	} {
		for _, scenario := range []string{"success", "lookup failure", "repeated rejection"} {
			t.Run(tc.name+"/"+scenario, func(t *testing.T) {
				c := newTestClient(t)
				c.config.AutoRetry = tc.policy
				serverErr := &Error{Code: 400, Message: tc.rejection}
				var calls []string
				c.sendRequest = func(_ int, data []byte) {
					w := decodeTestRequest(t, data)
					calls = append(calls, w.Type)
					var response TlObject = &Ok{}
					if len(calls) == 1 || (w.Type == "deleteMessages" && scenario == "repeated rejection") {
						response = serverErr
					}
					if w.Type == tc.lookup {
						if scenario == "lookup failure" {
							response = &Error{Code: 404, Message: "missing"}
						}
						if w.ChatID != 9007199254740993 || (tc.lookup == "getMessage" && w.MessageID != 9223372036854775806) {
							t.Errorf("lookup lost IDs: %+v", w)
						}
					}
					c.deliverResponse(w.Extra, response)
				}
				// A small function fixture carries both IDs; DeleteMessages uses a vector.
				_, err := c.Send(&testIDRequest{ChatID: 9007199254740993, MessageID: 9223372036854775806})
				want := []string{"deleteMessages", tc.lookup, "deleteMessages"}
				if scenario == "lookup failure" {
					want = want[:2]
				}
				if !reflect.DeepEqual(calls, want) {
					t.Fatalf("calls = %v, want %v", calls, want)
				}
				if scenario == "success" && err != nil {
					t.Fatal(err)
				}
				if scenario != "success" && err != serverErr {
					t.Fatalf("error = %v, want original %v", err, serverErr)
				}
				assertNoPendingCalls(t, c)
			})
		}
	}
}

type testIDRequest struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}

func (*testIDRequest) isFunction()     {}
func (*testIDRequest) GetType() string { return "deleteMessages" }
func (r *testIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type      string `json:"@type"`
		ChatID    int64  `json:"chat_id"`
		MessageID int64  `json:"message_id"`
	}{r.GetType(), r.ChatID, r.MessageID})
}

func TestFloodRetryBoundedAndCancelable(t *testing.T) {
	for _, tc := range []struct {
		name, message string
		limit         time.Duration
		want          error
	}{
		{"disabled", "Too Many Requests: retry after 1", 0, nil},
		{"over limit", "Too Many Requests: retry after 2", time.Second, nil},
		{"malformed", "Too Many Requests", time.Second, nil},
		{"budget", "Too Many Requests: retry after 1", time.Second, context.DeadlineExceeded},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.config.AutoRetry = &AutoRetry{MaxFloodWait: tc.limit}
			c.config.RequestTimeout = 20 * time.Millisecond
			serverErr := &Error{Code: 429, Message: tc.message}
			c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, serverErr) }
			_, err := c.Send(&GetMe{})
			if tc.want == nil && err != serverErr {
				t.Fatalf("error = %v, want %v", err, serverErr)
			}
			if tc.want != nil && !errors.Is(err, tc.want) {
				t.Fatalf("error = %v, want %v", err, tc.want)
			}
			if c.requestID.Load() != 1 {
				t.Fatal("unbounded flood retry")
			}
			assertNoPendingCalls(t, c)
		})
	}
	t.Run("one retry", func(t *testing.T) {
		c := newTestClient(t)
		c.config.AutoRetry = &AutoRetry{MaxFloodWait: time.Second}
		serverErr := &Error{Code: 429, Message: "Too Many Requests: retry after 1"}
		c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, serverErr) }
		_, err := c.Send(&GetMe{})
		if err != serverErr || c.requestID.Load() != 2 {
			t.Fatalf("error/attempts = %v/%d", err, c.requestID.Load())
		}
		assertNoPendingCalls(t, c)
	})
}

func TestStartContextAndAuthorizationBroadcast(t *testing.T) {
	c := newTestClient(t)
	c.startOnce.Do(func() {}) // Test authorization waiting without loading TDLib.
	if err := c.StartContext(nil); err == nil {
		t.Fatal("nil context accepted")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := c.StartContext(ctx); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	const count = 32
	done := make(chan error, count)
	for range count {
		go func() { done <- c.Start() }()
	}
	first := fmt.Errorf("authorization failed")
	c.sendAuthError(first)
	c.sendAuthError(nil)
	for range count {
		if err := receiveTest(t, done); err != first {
			t.Fatalf("authorization = %v", err)
		}
	}
	c.isAuthorized.Store(true)
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	c.closing.Store(true)
	if err := c.Start(); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestCloseConcurrentIdempotent(t *testing.T) {
	c := newTestClient(t)
	var sent atomic.Int32
	c.sendRequest = func(_ int, data []byte) {
		if w := decodeTestRequest(t, data); w.Type != "close" {
			t.Errorf("unexpected close request: %+v", w)
		}
		sent.Add(1)
		c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
	}
	var wg sync.WaitGroup
	for range 32 {
		wg.Add(1)
		go func() { defer wg.Done(); c.Close() }()
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	receiveTest(t, done)
	if sent.Load() != 1 || !c.closing.Load() {
		t.Fatalf("close count = %d", sent.Load())
	}
	select {
	case <-c.stop:
	default:
		t.Fatal("stop not closed")
	}
	select {
	case <-c.closed:
	default:
		t.Fatal("closed not signalled")
	}
	if err := c.updates.push(&UpdateNewMessage{}); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestCloseFromHandlerDoesNotDeadlock(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(_ int, _ []byte) {
		c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
	}
	done := make(chan struct{}, 1)
	c.AddHandler(testHandler{handle: func(c *Client, _ TlObject) error { c.Close(); done <- struct{}{}; return nil }})
	startTestProcessors(c, 1)
	if err := c.updates.push(queueTestUpdate(1)); err != nil {
		t.Fatal(err)
	}
	receiveTest(t, done)
	if _, err := c.Send(&GetMe{}); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestStartContextTimeoutAndCancellation(t *testing.T) {
	for _, reason := range []string{"timeout", "cancel", "stop"} {
		t.Run(reason, func(t *testing.T) {
			c := newTestClient(t)
			c.startOnce.Do(func() {})
			c.config.AuthorizationTimeout = 20 * time.Millisecond
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if reason == "cancel" {
				c.config.AuthorizationTimeout = time.Second
			}
			done := make(chan error, 1)
			go func() { done <- c.StartContext(ctx) }()
			if reason == "cancel" {
				cancel()
			}
			if reason == "stop" {
				c.stopOnce.Do(func() { close(c.stop) })
			}
			err := receiveTest(t, done)
			if reason == "timeout" && err == nil {
				t.Fatal("authorization timeout ignored")
			}
			if reason == "cancel" && !errors.Is(err, context.Canceled) {
				t.Fatal(err)
			}
			if reason == "stop" && !errors.Is(err, ErrClientClosed) {
				t.Fatal(err)
			}
		})
	}
}
