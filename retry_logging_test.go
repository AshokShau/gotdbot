package gotdbot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/AshokShau/gotdbot/logger"
)

func captureRequestLogs(c *Client, level slog.Level) *bytes.Buffer {
	out := new(bytes.Buffer)
	c.Logger = logger.New(logger.WithHandler(slog.NewJSONHandler(out, &slog.HandlerOptions{Level: level})))
	return out
}

func assertRequestLog(t *testing.T, out *bytes.Buffer, level, message string, attrs map[string]any) {
	t.Helper()
	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("invalid log record: %v: %s", err, line)
		}
		if record["msg"] != message || record["level"] != level {
			continue
		}
		for key, want := range attrs {
			if got := record[key]; got != want {
				t.Errorf("%s: %s = %v, want %v", message, key, got, want)
			}
		}
		return
	}
	t.Fatalf("missing %s log %q in:\n%s", level, message, out.String())
}

func TestFloodWaitDoesNotBlockConcurrentWork(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := newTestClient(t)
		c.me = &User{Id: 12345}
		c.config.AutoRetry = &AutoRetry{MaxFloodWait: 10 * time.Second}
		c.config.RequestTimeout = 30 * time.Second
		out := captureRequestLogs(c, slog.LevelDebug)
		var attempts atomic.Int32
		c.sendRequest = func(_ int, data []byte) {
			w := decodeTestRequest(t, data)
			if w.Type == "getMe" {
				if attempts.Add(1) == 1 {
					c.acceptResponse(&Error{Code: 429, Message: "Too Many Requests: retry after 10"}, w.Extra, nil)
					return
				}
			}
			c.acceptResponse(&Ok{}, w.Extra, nil)
		}
		finished := make(chan int64, 2)
		c.AddHandler(testHandler{handle: func(c *Client, u TlObject) error {
			id := u.(*UpdateNewMessage).Message.Id
			var req TlObject = &GetMe{}
			if id == 2 {
				req = &GetChat{ChatId: 42}
			}
			_, err := c.Send(req)
			if err != nil {
				t.Errorf("handler request: %v", err)
			}
			finished <- id
			return nil
		}})
		startTestProcessors(c, 2)
		c.acceptResponse(queueTestUpdate(1), "", nil)
		synctest.Wait() // The first handler is now waiting on its flood timer.
		assertRequestLog(t, out, "WARN", "TDLib flood wait; delaying request retry", map[string]any{
			"user_id": float64(c.Me().Id), "method": "getMe", "retry_after": float64(10 * time.Second),
		})
		assertNoPendingCalls(t, c) // The wait retains no request correlation or lock.

		started := time.Now()
		if _, err := c.Send(&GetOption{Name: "version"}); err != nil {
			t.Fatal(err)
		}
		c.acceptResponse(queueTestUpdate(2), "", nil)
		if id := receiveTest(t, finished); id != 2 {
			t.Fatalf("waiting handler finished before unrelated update: %d", id)
		}
		other := newTestClient(t)
		other.sendRequest = func(_ int, data []byte) {
			other.acceptResponse(&Ok{}, decodeTestRequest(t, data).Extra, nil)
		}
		if _, err := other.Send(&GetMe{}); err != nil {
			t.Fatal(err)
		}
		if elapsed := time.Since(started); elapsed != 0 {
			t.Fatalf("unrelated work waited for virtual time to advance: %s", elapsed)
		}
		time.Sleep(9 * time.Second)
		synctest.Wait()
		if attempts.Load() != 1 {
			t.Fatalf("request retried before its delay: %d attempts", attempts.Load())
		}
		time.Sleep(time.Second)
		if id := receiveTest(t, finished); id != 1 {
			t.Fatalf("waiting handler did not finish: %d", id)
		}
		synctest.Wait()
		if attempts.Load() != 2 {
			t.Fatalf("attempts = %d, want 2", attempts.Load())
		}
		assertRequestLog(t, out, "DEBUG", "Retrying TDLib request after flood wait", map[string]any{"method": "getMe"})
		assertNoPendingCalls(t, c)
	})
}

func TestFloodWaitInterrupted(t *testing.T) {
	for _, reason := range []string{"cancel", "timeout", "closed", "stop"} {
		t.Run(reason, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				c := newTestClient(t)
				c.config.AutoRetry = &AutoRetry{MaxFloodWait: time.Minute}
				out := captureRequestLogs(c, slog.LevelDebug)
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				c.sendRequest = func(_ int, data []byte) {
					c.acceptResponse(&Error{Code: 429, Message: "Too Many Requests: retry after 60"}, decodeTestRequest(t, data).Extra, nil)
				}
				done := make(chan error, 1)
				go func() { _, err := c.SendWithContext(ctx, &GetMe{}); done <- err }()
				synctest.Wait()
				want := ErrClientClosed
				switch reason {
				case "cancel":
					cancel()
					want = context.Canceled
				case "timeout":
					time.Sleep(c.config.RequestTimeout)
					want = context.DeadlineExceeded
				case "closed":
					c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
				case "stop":
					c.stopOnce.Do(func() { close(c.stop) })
				}
				err := receiveTest(t, done)
				if !errors.Is(err, want) {
					t.Fatalf("error = %v, want %v", err, want)
				}
				if reason == "timeout" {
					if !errors.Is(err, SendTimeout) {
						t.Fatalf("missing SendTimeout: %v", err)
					}
					assertRequestLog(t, out, "WARN", "TDLib request timed out", map[string]any{"method": "getMe"})
				} else {
					assertRequestLog(t, out, "DEBUG", "TDLib request failed", map[string]any{"method": "getMe"})
				}
				if c.requestID.Load() != 1 {
					t.Fatal("interrupted flood wait retried the request")
				}
				assertNoPendingCalls(t, c)
			})
		})
	}
}

func TestFloodRetrySkippedLogging(t *testing.T) {
	for _, tc := range []struct {
		name, message, reason string
		limit                 time.Duration
	}{
		{"disabled", "Too Many Requests: retry after 1", "automatic flood retry disabled", 0},
		{"over limit", "Too Many Requests: retry after 2", "retry delay exceeds MaxFloodWait", time.Second},
		{"malformed", "Too Many Requests", "invalid retry delay", time.Second},
		{"negative", "Too Many Requests: retry after -1", "invalid retry delay", time.Second},
		{"large delay", "Too Many Requests: retry after " + strconv.Itoa(int(^uint(0)>>1)), "retry delay exceeds MaxFloodWait", time.Second},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t)
			c.config.AutoRetry = &AutoRetry{MaxFloodWait: tc.limit}
			out := captureRequestLogs(c, slog.LevelInfo)
			serverErr := &Error{Code: 429, Message: tc.message}
			c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, serverErr) }
			if _, err := c.Send(&GetMe{}); err != serverErr {
				t.Fatalf("error = %v, want original rejection", err)
			}
			assertRequestLog(t, out, "WARN", "TDLib flood wait; request will not be retried", map[string]any{"reason": tc.reason})
			if strings.Contains(out.String(), "DEBUG") || c.requestID.Load() != 1 {
				t.Fatalf("unexpected debug output or retry: %s", out.String())
			}
		})
	}
}

func TestRequestTraceOmitsBody(t *testing.T) {
	c := newTestClient(t)
	out := captureRequestLogs(c, slog.LevelDebug)
	c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, &Ok{}) }
	if _, err := c.Send(&CheckAuthenticationBotToken{Token: "private-bot-token"}); err != nil {
		t.Fatal(err)
	}
	assertRequestLog(t, out, "DEBUG", "Sending TDLib request", map[string]any{"method": "checkAuthenticationBotToken", "request_id": "1"})
	assertRequestLog(t, out, "DEBUG", "TDLib request completed", map[string]any{"method": "checkAuthenticationBotToken"})
	if strings.Contains(out.String(), "private-bot-token") {
		t.Fatal("request trace exposed authentication credentials")
	}
	if strings.Contains(out.String(), "tdlib_client_id") || strings.Contains(out.String(), `"client_id"`) {
		t.Fatal("request trace includes an internal TDLib client ID")
	}
}

func TestQRCodeFailureIncludesLink(t *testing.T) {
	c := newTestClient(t)
	out := captureRequestLogs(c, slog.LevelWarn)
	link := "tg://login?token=" + strings.Repeat("x", 512)
	c.authHandler(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateWaitOtherDeviceConfirmation{Link: link}})
	assertRequestLog(t, out, "WARN", "Cannot generate authentication QR code; open the link in Telegram", map[string]any{"link": link})
}

func TestAuthorizationUpdatesLogIdentity(t *testing.T) {
	for _, initialID := range []int64{0, 12345} {
		t.Run(strconv.FormatInt(initialID, 10), func(t *testing.T) {
			c := newTestClient(t)
			c.me = &User{Id: initialID}
			out := captureRequestLogs(c, slog.LevelDebug)
			me := &User{Id: 98765, FirstName: "Account"}
			c.sendRequest = func(_ int, data []byte) {
				c.deliverResponse(decodeTestRequest(t, data).Extra, me)
			}
			c.authHandler(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateReady{}})
			if c.Me() != me || c.clientID != 7 {
				t.Fatal("authorization did not update the account identity independently of TDLib routing")
			}
			out.Reset()
			if _, err := c.GetMe(); err != nil {
				t.Fatal(err)
			}
			assertRequestLog(t, out, "DEBUG", "Sending TDLib request", map[string]any{
				"user_id": float64(me.Id),
			})
		})
	}
}

func TestAuthorizationFailureKeepsMe(t *testing.T) {
	var nilUser *User
	for _, response := range []TlObject{nilUser, &Ok{}, &Error{Code: 500, Message: "unavailable"}} {
		c := newTestClient(t)
		initial := c.Me()
		c.sendRequest = func(_ int, data []byte) {
			c.deliverResponse(decodeTestRequest(t, data).Extra, response)
		}
		c.authHandler(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateReady{}})
		if c.Me() == nil || c.Me() != initial || c.Me().Id != 0 {
			t.Fatal("failed authorization replaced the initialized identity")
		}
		if c.authErr == nil || c.isAuthorized.Load() {
			t.Fatal("invalid account response was accepted")
		}
		assertNoPendingCalls(t, c)
	}
}

func TestFloodRetryExhaustionLogging(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := newTestClient(t)
		c.config.AutoRetry = &AutoRetry{MaxFloodWait: time.Second}
		out := captureRequestLogs(c, slog.LevelInfo)
		serverErr := &Error{Code: 429, Message: "Too Many Requests: retry after 1"}
		c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, serverErr) }
		started := time.Now()
		if _, err := c.Send(&GetMe{}); err != serverErr {
			t.Fatalf("error = %v, want original rejection", err)
		}
		if c.requestID.Load() != 2 || time.Since(started) != time.Second {
			t.Fatal("flood retry did not stop after one wait and retry")
		}
		assertRequestLog(t, out, "WARN", "TDLib flood wait; request will not be retried", map[string]any{"reason": "retry already attempted"})
		assertNoPendingCalls(t, c)
	})
}

func TestCombinedRetryLimit(t *testing.T) {
	for _, rejection := range []*Error{
		{Code: 400, Message: "Chat not found"},
		{Code: 400, Message: "Message not found"},
		{Code: 429, Message: "Too Many Requests: retry after 1"},
	} {
		t.Run(rejection.Message, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				c := newTestClient(t)
				c.config.RequestTimeout = 0
				c.config.AutoRetry = &AutoRetry{ChatNotFound: true, MessageNotFound: true, MaxFloodWait: time.Second}
				calls := make(map[string]int)
				rejections := []*Error{
					{Code: 400, Message: "Chat not found"},
					{Code: 400, Message: "Message not found"},
					{Code: 429, Message: "Too Many Requests: retry after 1"},
					rejection,
				}
				c.sendRequest = func(_ int, data []byte) {
					w := decodeTestRequest(t, data)
					calls[w.Type]++
					var response TlObject = &Ok{}
					if w.Type == "deleteMessages" {
						if calls[w.Type] > len(rejections) {
							t.Fatal("original request exceeded its retry limit")
						}
						response = rejections[calls[w.Type]-1]
					} else if calls[w.Type] > 1 {
						t.Fatal("missing-object lookup repeated")
					}
					c.deliverResponse(w.Extra, response)
				}
				if _, err := c.Send(&testIDRequest{ChatID: 42, MessageID: 123}); err != rejection {
					t.Fatalf("error = %v, want final rejection", err)
				}
				if calls["deleteMessages"] != 4 || calls["getChat"] != 1 || calls["getMessage"] != 1 {
					t.Fatalf("unexpected attempts: %v", calls)
				}
				assertNoPendingCalls(t, c)
			})
		})
	}
}

func TestLookupRetryAbort(t *testing.T) {
	for _, lookup := range []string{"getChat", "getMessage"} {
		for _, reason := range []string{"cancel", "timeout", "closed", "stop"} {
			t.Run(lookup+"/"+reason, func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					c := newTestClient(t)
					c.config.AutoRetry = &AutoRetry{ChatNotFound: true, MessageNotFound: true}
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()
					want := ErrClientClosed
					if reason == "cancel" {
						want = context.Canceled
					} else if reason == "timeout" {
						want = context.DeadlineExceeded
					}
					c.sendRequest = func(_ int, data []byte) {
						w := decodeTestRequest(t, data)
						if w.Type == lookup {
							switch reason {
							case "cancel":
								cancel()
							case "closed":
								c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
							case "stop":
								c.stopOnce.Do(func() { close(c.stop) })
							}
							return
						}
						message := "Chat not found"
						if lookup == "getMessage" {
							message = "Message not found"
						}
						c.deliverResponse(w.Extra, &Error{Code: 400, Message: message})
					}
					if _, err := c.SendWithContext(ctx, &testIDRequest{ChatID: 42, MessageID: 123}); !errors.Is(err, want) {
						t.Fatalf("lookup error = %v, want %v", err, want)
					}
					assertNoPendingCalls(t, c)
				})
			})
		}
	}
}

func TestRequestsWithoutTimeout(t *testing.T) {
	for _, timeout := range []time.Duration{0, -time.Second} {
		for _, tc := range []struct {
			name     string
			call     func(*Client) (TlObject, error)
			response TlObject
		}{
			{"request", func(c *Client) (TlObject, error) { return c.GetMe() }, &User{Id: 7}},
			{"download", func(c *Client) (TlObject, error) {
				return c.DownloadFile(42, 1, &DownloadFileOpts{Synchronous: true})
			}, &File{Id: 42, Local: &LocalFile{IsDownloadingCompleted: true}}},
			{"send file", func(c *Client) (TlObject, error) {
				return c.SendDocument(12, &InputFileLocal{Path: "large.zip"}, nil)
			}, &Message{ChatId: 12, Id: 100}},
		} {
			t.Run(tc.name+"/"+timeout.String(), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					c := newTestClient(t)
					c.config.RequestTimeout = timeout
					c.sendRequest = func(_ int, data []byte) {
						w := decodeTestRequest(t, data)
						if w.Type == "sendMessage" {
							c.acceptResponse(&Message{ChatId: 12, Id: -1, SendingState: &MessageSendingStatePending{}}, w.Extra, nil)
						}
						go func() {
							time.Sleep(2 * time.Minute)
							if w.Type == "sendMessage" {
								c.acceptResponse(&UpdateMessageSendSucceeded{OldMessageId: -1, Message: tc.response.(*Message)}, "", nil)
							} else {
								c.acceptResponse(tc.response, w.Extra, nil)
							}
						}()
					}
					started := time.Now()
					result, err := tc.call(c)
					if err != nil || result != tc.response || time.Since(started) != 2*time.Minute {
						t.Fatalf("slow request = (%v, %v), elapsed %s", result, err, time.Since(started))
					}
					if c.requestID.Load() != 1 {
						t.Fatal("slow request was retried")
					}
					assertNoPendingCalls(t, c)
				})
			})
		}
	}
}

func TestNoRequestTimeoutStillStopsWaiting(t *testing.T) {
	for _, req := range []TlObject{&DownloadFile{FileId: 42, Priority: 1, Synchronous: true}, &SendMessage{ChatId: 12}} {
		for _, reason := range []string{"cancel", "deadline", "closed", "stop"} {
			t.Run(req.GetType()+"/"+reason, func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					c := newTestClient(t)
					c.config.RequestTimeout = DefaultClientConfig().RequestTimeout
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()
					if reason == "deadline" {
						var deadlineCancel context.CancelFunc
						ctx, deadlineCancel = context.WithTimeout(ctx, time.Second)
						defer deadlineCancel()
					}
					c.sendRequest = func(_ int, data []byte) {
						if waitsForSend(req) {
							c.acceptResponse(&Message{ChatId: 12, Id: -1, SendingState: &MessageSendingStatePending{}}, decodeTestRequest(t, data).Extra, nil)
						}
					}
					done := make(chan error, 1)
					go func() { _, err := c.SendWithContext(ctx, req); done <- err }()
					synctest.Wait()
					want := ErrClientClosed
					switch reason {
					case "cancel":
						cancel()
						want = context.Canceled
					case "deadline":
						time.Sleep(time.Second)
						want = context.DeadlineExceeded
					case "closed":
						c.acceptResponse(&UpdateAuthorizationState{AuthorizationState: &AuthorizationStateClosed{}}, "", nil)
					case "stop":
						c.stopOnce.Do(func() { close(c.stop) })
					}
					err := receiveTest(t, done)
					if !errors.Is(err, want) {
						t.Fatalf("error = %v, want %v", err, want)
					}
					if reason == "deadline" && (!errors.Is(err, SendTimeout) || strings.Contains(err.Error(), "\n")) {
						t.Fatalf("unexpected timeout error: %v", err)
					}
					assertNoPendingCalls(t, c)
				})
			})
		}
	}
}
