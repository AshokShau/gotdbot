package gotdbot_test

import (
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
