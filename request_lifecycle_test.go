package gotdbot

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestWaitsForSend(t *testing.T) {
	for _, req := range []TlObject{&SendMessage{}, &SendMessageAlbum{}, &SendEphemeralMessage{}, &ForwardMessages{}, &SendBotStartMessage{}, &SendInlineQueryResultMessage{}, &SendQuickReplyShortcutMessages{}, &ResendMessages{}, &AddOffer{}} {
		if !waitsForSend(req) {
			t.Errorf("%T must await completion", req)
		}
	}
	for _, req := range []TlObject{nil, &GetMessage{}, &GetChat{}, &EditMessageText{}, &Close{}} {
		if waitsForSend(req) {
			t.Errorf("%T must not await completion", req)
		}
	}
}

func TestSendImmediateCompletion(t *testing.T) {
	for _, outcome := range []string{"success", "failure"} {
		t.Run(outcome, func(t *testing.T) {
			c := newTestClient(t)
			pending := &Message{ChatId: 12, Id: -1, SendingState: &MessageSendingStatePending{}}
			final := &Message{ChatId: 12, Id: 100}
			failure := &Error{Code: 400, Message: "send failed"}
			c.sendRequest = func(_ int, data []byte) {
				c.deliverResponse(decodeTestRequest(t, data).Extra, pending)
				var completion TlObject = &UpdateMessageSendSucceeded{Message: final, OldMessageId: -1}
				if outcome == "failure" {
					completion = &UpdateMessageSendFailed{Message: final, OldMessageId: -1, Error: failure}
				}
				// Completion happens before Send even starts waiting for its reply.
				c.acceptResponse(completion, "", nil)
				c.acceptResponse(completion, "", nil) // Duplicate must never block.
			}
			obj, err := c.Send(&SendMessage{ChatId: 12})
			if obj != final {
				t.Fatalf("result = %v, want final message", obj)
			}
			if outcome == "success" && err != nil {
				t.Fatal(err)
			}
			if outcome == "failure" && err != failure {
				t.Fatalf("error = %v", err)
			}
			assertNoPendingCalls(t, c)
		})
	}
}

func TestSendAlbumOutOfOrderAndChatIsolation(t *testing.T) {
	c := newTestClient(t)
	failure := &Error{Code: 400, Message: "second message failed"}
	album := &Messages{TotalCount: 3, Messages: []Message{
		{ChatId: 11, Id: -1, SendingState: &MessageSendingStatePending{}},
		{ChatId: 22, Id: -1, SendingState: &MessageSendingStatePending{}},
		{ChatId: 11, Id: 3},
	}}
	c.sendRequest = func(_ int, data []byte) {
		c.deliverResponse(decodeTestRequest(t, data).Extra, album)
		c.acceptResponse(&UpdateMessageSendFailed{Message: &Message{ChatId: 22, Id: 202}, OldMessageId: -1, Error: failure}, "", nil)
		c.acceptResponse(&UpdateMessageSendSucceeded{Message: &Message{ChatId: 11, Id: 101}, OldMessageId: -1}, "", nil)
	}
	obj, err := c.Send(&SendMessageAlbum{})
	if obj != album || err != failure {
		t.Fatalf("album = (%v, %v)", obj, err)
	}
	for i, want := range []int64{101, 202, 3} {
		if album.Messages[i].Id != want {
			t.Errorf("message %d = %d, want %d", i, album.Messages[i].Id, want)
		}
	}
	assertNoPendingCalls(t, c)
}

func TestReadPendingMessageDoesNotWaitForSend(t *testing.T) {
	c := newTestClient(t)
	m := &Message{ChatId: 12, Id: -1, SendingState: &MessageSendingStatePending{}}
	c.sendRequest = func(_ int, data []byte) { c.deliverResponse(decodeTestRequest(t, data).Extra, m) }
	obj, err := c.Send(&GetMessage{ChatId: 12, MessageId: -1})
	if obj != m || err != nil {
		t.Fatalf("pending read = (%v, %v)", obj, err)
	}
	assertNoPendingCalls(t, c)
}

func TestSendCompletionAbortCleanup(t *testing.T) {
	for _, reason := range []string{"cancel", "timeout", "closed", "stop"} {
		t.Run(reason, func(t *testing.T) {
			c := newTestClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if reason == "timeout" {
				c.config.RequestTimeout = 20 * time.Millisecond
			}
			ready := make(chan struct{}, 1)
			pending := &Message{ChatId: 12, Id: -1, SendingState: &MessageSendingStatePending{}}
			c.sendRequest = func(_ int, data []byte) {
				c.deliverResponse(decodeTestRequest(t, data).Extra, pending)
				ready <- struct{}{}
			}
			done := make(chan error, 1)
			go func() { _, err := c.SendWithContext(ctx, &SendMessage{}); done <- err }()
			receiveTest(t, ready)
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
			if err := receiveTest(t, done); !errors.Is(err, want) || (reason == "timeout" && !errors.Is(err, SendTimeout)) {
				t.Fatalf("completion error = %v", err)
			}
			assertNoPendingCalls(t, c)
			c.completeMessage(12, -1, &UpdateMessageSendSucceeded{Message: &Message{ChatId: 12, Id: 1}})
		})
	}
}

func TestWaitSendResponseVariants(t *testing.T) {
	c := newTestClient(t)
	failure := &Error{Code: 403, Message: "denied"}
	for _, tc := range []struct {
		name    string
		obj     TlObject
		wantErr bool
	}{
		{"nil message", (*Message)(nil), true},
		{"nil messages", (*Messages)(nil), true},
		{"empty album", &Messages{}, false},
		{"failed state", &Message{SendingState: &MessageSendingStateFailed{Error: failure}}, true},
		{"sent message", &Message{Id: 1}, false},
		{"other result", &Ok{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := c.waitSend(context.Background(), tc.obj, &pendingCall{})
			if (err != nil) != tc.wantErr {
				t.Fatalf("waitSend error = %v", err)
			}
		})
	}
	ch := make(chan TlObject, 1)
	ch <- &Ok{}
	if _, err := c.waitSend(context.Background(), &Message{ChatId: 1, Id: -1, SendingState: &MessageSendingStatePending{}}, &pendingCall{messages: map[[2]int64]chan TlObject{{1, -1}: ch}}); err == nil {
		t.Fatal("unexpected completion accepted")
	}
}

func TestConcurrentCompletionAndCancellation(t *testing.T) {
	c := newTestClient(t)
	const count = 128
	var wg sync.WaitGroup
	for i := range count {
		extra := string(rune(i + 1))
		p := &pendingCall{response: make(chan TlObject, 1), waitForSend: true}
		c.requestMu.Lock()
		c.pendingRequests[extra] = p
		c.requestMu.Unlock()
		c.deliverResponse(extra, &Message{ChatId: int64(i + 1), Id: -1, SendingState: &MessageSendingStatePending{}})
		wg.Add(3)
		go func() { defer wg.Done(); c.finishCall(extra) }()
		go func() {
			defer wg.Done()
			c.completeMessage(int64(i+1), -1, &UpdateMessageSendSucceeded{Message: &Message{Id: 1}})
		}()
		go func() { defer wg.Done(); c.deliverResponse(extra, &Ok{}) }()
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	receiveTest(t, done)
	assertNoPendingCalls(t, c)
}

func TestResponseAs(t *testing.T) {
	want := &User{Id: 7}
	if got, err := responseAs[*User](want); got != want || err != nil {
		t.Fatalf("responseAs = (%v, %v)", got, err)
	}
	if got, err := responseAs[*User](&Ok{}); got != nil || err == nil {
		t.Fatal("wrong response type accepted")
	}
	if _, err := responseAs[*User](nil); err == nil {
		t.Fatal("nil response accepted")
	}
}
