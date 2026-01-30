package ext_test

import (
	"testing"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

func TestDispatcher_Message(t *testing.T) {
	d := ext.NewDispatcher(nil)

	called := false
	h := handlers.NewMessage(filters.Text, func(ctx *ext.Context) error {
		called = true
		return nil
	})

	d.AddHandler(h)

	update := &gotdbot.UpdateNewMessage{
		Message: &gotdbot.Message{
			ChatId: 123,
			Content: &gotdbot.MessageContent{
				MessageText: &gotdbot.MessageText{
					Text: &gotdbot.FormattedText{
						Text: "hello",
					},
				},
			},
		},
	}

	d.ProcessUpdate(update)

	if !called {
		t.Errorf("Handler was not called")
	}
}

func TestDispatcher_Command(t *testing.T) {
	d := ext.NewDispatcher(nil)
	called := false
	cmd := handlers.NewCommand("start", func(ctx *ext.Context) error {
		called = true
		return nil
	})
	d.AddHandler(cmd)

	update := &gotdbot.UpdateNewMessage{
		Message: &gotdbot.Message{
			ChatId: 123,
			Content: &gotdbot.MessageContent{
				MessageText: &gotdbot.MessageText{
					Text: &gotdbot.FormattedText{
						Text: "/start",
					},
				},
			},
		},
	}

	d.ProcessUpdate(update)
	if !called {
		t.Errorf("Command handler was not called")
	}
}

func TestDispatcher_InlineQuery(t *testing.T) {
	d := ext.NewDispatcher(nil)
	called := false
	h := handlers.NewInlineQuery(nil, func(ctx *ext.Context) error {
		called = true
		return nil
	})
	d.AddHandler(h)

	update := &gotdbot.UpdateNewInlineQuery{
		Id:    "1",
		Query: "test",
	}

	d.ProcessUpdate(update)
	if !called {
		t.Errorf("InlineQuery handler was not called")
	}
}
