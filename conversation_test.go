package gotdbot

import (
	"context"
	"errors"
	"testing"
	"time"
)

func waitForTestWaiters(t *testing.T, c *Client, count int) {
	t.Helper()
	eventuallyTest(t, func() bool {
		c.wMu.RLock()
		defer c.wMu.RUnlock()
		n := 0
		for _, inner := range c.waiters {
			n += len(inner)
		}
		return n == count
	})
}

func TestWaitForChatRoutingAndReentrantFilter(t *testing.T) {
	c := newTestClient(t)
	startTestProcessors(c, 2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	type result struct {
		update TlObject
		err    error
	}
	global := make(chan result, 1)
	chat := make(chan result, 1)
	other := make(chan result, 1)
	filter := func(c *Client, u TlObject) bool {
		// Registration/removal from inside a filter must not run under wMu.
		_, err := c.WaitForContext(context.Background(), 99, func(*Client, TlObject) bool { return false }, 0)
		if !errors.Is(err, ConversationTimeout) {
			t.Errorf("nested waiter: %v", err)
		}
		return u.(*UpdateNewMessage).Message.Id == 1
	}
	go func() { u, err := c.WaitFor(filter, time.Second); global <- result{u, err} }()
	go func() { u, err := c.WaitForChat(42, filter, time.Second); chat <- result{u, err} }()
	go func() {
		u, err := c.WaitForContext(ctx, 43, func(*Client, TlObject) bool { t.Error("wrong-chat filter called"); return true }, time.Second)
		other <- result{u, err}
	}()
	waitForTestWaiters(t, c, 3)
	update := queueTestUpdate(1)
	if err := c.updates.push(update); err != nil {
		t.Fatal(err)
	}
	for _, ch := range []chan result{global, chat} {
		r := receiveTest(t, ch)
		if r.update != update || r.err != nil {
			t.Fatalf("waiter = %+v", r)
		}
	}
	cancel()
	if r := receiveTest(t, other); !errors.Is(r.err, context.Canceled) {
		t.Fatal(r.err)
	}
	waitForTestWaiters(t, c, 0)
}

func TestWaitForExitPaths(t *testing.T) {
	for _, reason := range []string{"timeout", "cancel", "closed", "stop"} {
		t.Run(reason, func(t *testing.T) {
			c := newTestClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			timeout := time.Second
			if reason == "timeout" {
				timeout = 20 * time.Millisecond
			}
			done := make(chan error, 1)
			go func() {
				_, err := c.WaitForContext(ctx, 42, func(*Client, TlObject) bool { return true }, timeout)
				done <- err
			}()
			waitForTestWaiters(t, c, 1)
			want := ErrClientClosed
			switch reason {
			case "timeout":
				want = ConversationTimeout
			case "cancel":
				cancel()
				want = context.Canceled
			case "closed":
				c.closedOnce.Do(func() { close(c.closed) })
			case "stop":
				c.stopOnce.Do(func() { close(c.stop) })
			}
			if err := receiveTest(t, done); !errors.Is(err, want) {
				t.Fatalf("error = %v, want %v", err, want)
			}
			waitForTestWaiters(t, c, 0)
		})
	}
	c := newTestClient(t)
	filter := func(*Client, TlObject) bool { return true }
	if _, err := c.WaitForContext(nil, 0, filter, time.Second); err == nil {
		t.Fatal("nil context accepted")
	}
	if _, err := c.WaitFor(nil, time.Second); err == nil {
		t.Fatal("nil filter accepted")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.WaitForContext(ctx, 0, filter, time.Second); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	c.closing.Store(true)
	if _, err := c.WaitFor(filter, time.Second); !errors.Is(err, ErrClientClosed) {
		t.Fatal(err)
	}
}

func TestAskFiltersAndCancellation(t *testing.T) {
	for _, cancelConversation := range []bool{false, true} {
		t.Run(map[bool]string{false: "reply", true: "cancel"}[cancelConversation], func(t *testing.T) {
			c := newTestClient(t)
			startTestProcessors(c, 1)
			type result struct {
				msg *Message
				err error
			}
			done := make(chan result, 1)
			opts := &WaitMessageOpts{Timeout: time.Second, Filter: func(m *Message) bool { return m.GetText() == "answer" }, CancellationFilter: func(m *Message) bool { return m.GetText() == "cancel" }}
			go func() { msg, err := c.Ask(42, opts); done <- result{msg, err} }()
			waitForTestWaiters(t, c, 1)
			for _, u := range []TlObject{&UpdateMessageEdited{ChatId: 42}, &UpdateNewMessage{Message: testMessage(43, "answer")}, &UpdateNewMessage{Message: testMessage(42, "ignored")}} {
				if err := c.updates.push(u); err != nil {
					t.Fatal(err)
				}
			}
			text := "answer"
			if cancelConversation {
				text = "cancel"
			}
			msg := testMessage(42, text)
			if err := c.updates.push(&UpdateNewMessage{Message: msg}); err != nil {
				t.Fatal(err)
			}
			r := receiveTest(t, done)
			if cancelConversation {
				if r.msg != nil || !errors.Is(r.err, ConversationCancelled) {
					t.Fatalf("cancellation = %+v", r)
				}
			} else if r.msg != msg || r.err != nil {
				t.Fatalf("reply = %+v", r)
			}
			waitForTestWaiters(t, c, 0)
		})
	}
}
