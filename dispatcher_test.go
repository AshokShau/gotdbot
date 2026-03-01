package gotdbot_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

func TestDispatcher_Message(t *testing.T) {
	d := gotdbot.NewDispatcher(nil, nil)

	var wg sync.WaitGroup
	wg.Add(1)
	called := false
	h := handlers.NewMessage(filters.Text, func(client *gotdbot.Client, ctx *gotdbot.Context) error {
		called = true
		wg.Done()
		return nil
	})

	d.AddHandler(h)

	update := &gotdbot.UpdateNewMessage{
		Message: &gotdbot.Message{
			ChatId: 123,
			Content: &gotdbot.MessageText{
				Text: &gotdbot.FormattedText{
					Text: "hello",
				},
			},
		},
	}

	d.ProcessUpdate(update)

	if waitTimeout(&wg, 1*time.Second) {
		t.Errorf("Handler was not called (timeout)")
	}
	if !called {
		t.Errorf("Handler was not called")
	}
}

func TestDispatcher_Command(t *testing.T) {
	d := gotdbot.NewDispatcher(&gotdbot.Client{}, nil)
	var wg sync.WaitGroup
	wg.Add(1)
	called := false
	cmd := handlers.NewCommand("start", func(client *gotdbot.Client, ctx *gotdbot.Context) error {
		called = true
		wg.Done()
		return nil
	})
	d.AddHandler(cmd)

	update := &gotdbot.UpdateNewMessage{
		Message: &gotdbot.Message{
			ChatId: 123,
			Content: &gotdbot.MessageText{
				Text: &gotdbot.FormattedText{
					Text: "/start",
				},
			},
		},
	}

	d.ProcessUpdate(update)
	if waitTimeout(&wg, 1*time.Second) {
		t.Errorf("Command handler was not called (timeout)")
	}
	if !called {
		t.Errorf("Command handler was not called")
	}
}

func TestDispatcher_InlineQuery(t *testing.T) {
	d := gotdbot.NewDispatcher(nil, nil)
	var wg sync.WaitGroup
	wg.Add(1)
	called := false
	h := handlers.NewUpdateNewInlineQuery(nil, func(client *gotdbot.Client, ctx *gotdbot.Context) error {
		called = true
		wg.Done()
		return nil
	})
	d.AddHandler(h)

	update := &gotdbot.UpdateNewInlineQuery{
		Id:    1,
		Query: "test",
	}

	d.ProcessUpdate(update)
	if waitTimeout(&wg, 1*time.Second) {
		t.Errorf("InlineQuery handler was not called (timeout)")
	}
	if !called {
		t.Errorf("InlineQuery handler was not called")
	}
}

func TestDispatcher_ErrorHandler_ControlFlow_Message(t *testing.T) {
	type flowResult struct {
		ranHandlers     int
		errorHandlerErr error
	}

	type testcase struct {
		name              string
		configure         func(*gotdbot.Dispatcher, *sync.Mutex, *int)
		errorHandler      func(error) error
		expectErrorCalled bool
		expect            flowResult
	}

	endGroupsErr := gotdbot.EndGroups
	continueGroupsErr := gotdbot.ContinueGroups

	tests := []testcase{
		{
			name: "handler error routed through ErrorHandler, handler return nil continues",
			configure: func(d *gotdbot.Dispatcher, mu *sync.Mutex, handC *int) {
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return fmt.Errorf("handler failure")
				}), 0)
			},
			errorHandler: func(err error) error {
				// swallow error, allow processing to continue
				return nil
			},
			expectErrorCalled: true,
			expect: flowResult{
				ranHandlers: 1,
			},
		},
		{
			name: "handler directly returns EndGroups, ErrorHandler not called, next group doesn't run",
			configure: func(d *gotdbot.Dispatcher, mu *sync.Mutex, handC *int) {
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return endGroupsErr
				}), 0)
				// this handler should _not_ run because EndGroups was returned
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					t.Errorf("second handler should not be executed when EndGroups is returned")
					return nil
				}), 1)
			},
			errorHandler:      nil,
			expectErrorCalled: false,
			expect: flowResult{
				ranHandlers: 1,
			},
		},
		{
			name: "handler directly returns ContinueGroups, ErrorHandler not called, next group runs",
			configure: func(d *gotdbot.Dispatcher, mu *sync.Mutex, handC *int) {
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return continueGroupsErr
				}), 0)
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return nil
				}), 1)
			},
			errorHandler:      nil,
			expectErrorCalled: false,
			expect: flowResult{
				ranHandlers: 2, // one handler in each group
			},
		},
		{
			name: "nil ErrorHandler: errors are logged, flow controlled by dispatcher defaults",
			configure: func(d *gotdbot.Dispatcher, mu *sync.Mutex, handC *int) {
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return fmt.Errorf("handler failure")
				}), 0)
				d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
					mu.Lock()
					*handC++
					mu.Unlock()
					return nil
				}), 1)
			},
			errorHandler:      nil,
			expectErrorCalled: false,
			expect: flowResult{
				ranHandlers: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			var (
				mu              sync.Mutex
				ranHandlers     int
				errorCalled     bool
				errorHandlerErr error
			)

			// Wrap user-provided ErrorHandler so we can observe calls.
			var wrappedErrHandler func(client *gotdbot.Client, ctx *gotdbot.Context, err error) error
			if tt.errorHandler != nil {
				wrappedErrHandler = func(client *gotdbot.Client, ctx *gotdbot.Context, err error) error {
					mu.Lock()
					defer mu.Unlock()
					errorCalled = true
					res := tt.errorHandler(err)
					errorHandlerErr = res
					return res
				}
			}

			d := gotdbot.NewDispatcher(nil, &gotdbot.DispatcherOpts{ErrorHandler: wrappedErrHandler})

			// Let the test case override with more specific behavior, including
			// returning errors and control-flow markers.
			if tt.configure != nil {
				tt.configure(d, &mu, &ranHandlers)
			}

			msg := &gotdbot.UpdateNewMessage{
				Message: &gotdbot.Message{},
			}

			// Waitgroup to ensure dispatcher routine runs
			var wg sync.WaitGroup
			wg.Add(1)
			d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
				wg.Done()
				return nil
			}), 100) // use a high group number to ensure it runs last

			// Dispatch a single message to drive the pipeline.
			d.ProcessUpdate(msg)

			// Try to wait, if it times out just continue
			waitTimeout(&wg, 100*time.Millisecond)

			mu.Lock()
			defer mu.Unlock()

			if tt.expectErrorCalled != errorCalled {
				t.Errorf("ErrorHandler called = %v, want %v", errorCalled, tt.expectErrorCalled)
			}

			if tt.expect.errorHandlerErr != nil && tt.expect.errorHandlerErr != errorHandlerErr {
				t.Errorf("ErrorHandler returned = %v, want %v", errorHandlerErr, tt.expect.errorHandlerErr)
			}

			if tt.expect.ranHandlers != 0 && ranHandlers != tt.expect.ranHandlers {
				t.Errorf("ranHandlers = %d, want %d", ranHandlers, tt.expect.ranHandlers)
			}
		})
	}
}

func TestDispatcher_ErrorHandler_ControlFlow_Command(t *testing.T) {
	t.Run("command pipeline respects ErrorHandler semantics", func(t *testing.T) {
		var (
			mu          sync.Mutex
			ranHandlers int
			errorCalled bool
		)

		errHandler := func(client *gotdbot.Client, ctx *gotdbot.Context, err error) error {
			mu.Lock()
			errorCalled = true
			mu.Unlock()
			// ensure commands continue even on handler error
			return nil
		}

		d := gotdbot.NewDispatcher(&gotdbot.Client{}, &gotdbot.DispatcherOpts{ErrorHandler: errHandler})

		d.AddHandlerToGroup(handlers.NewCommand("start", func(client *gotdbot.Client, ctx *gotdbot.Context) error {
			mu.Lock()
			ranHandlers++
			mu.Unlock()
			return fmt.Errorf("command handler failure")
		}), 0)

		cmd := &gotdbot.UpdateNewMessage{
			Message: &gotdbot.Message{
				Content: &gotdbot.MessageText{
					Text: &gotdbot.FormattedText{
						Text: "/start",
					},
				},
			},
		}

		var wg sync.WaitGroup
		wg.Add(1)
		d.AddHandlerToGroup(handlers.NewMessage(filters.Message(func(msg *gotdbot.Message) bool { return true }), func(client *gotdbot.Client, ctx *gotdbot.Context) error {
			wg.Done()
			return nil
		}), 100)

		d.ProcessUpdate(cmd)
		waitTimeout(&wg, 100*time.Millisecond)

		mu.Lock()
		defer mu.Unlock()

		if !errorCalled {
			t.Fatalf("expected ErrorHandler to be called for command handler error")
		}
		if ranHandlers == 0 {
			t.Fatalf("expected command handlers to run at least once")
		}
	})
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
