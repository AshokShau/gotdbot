package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

type Message struct {
	Filter   filters.Message
	Response func(b *gotdbot.Client, ctx *gotdbot.Context) error
}

func NewMessage(filter filters.Message, response func(b *gotdbot.Client, ctx *gotdbot.Context) error) *Message {
	return &Message{
		Filter:   filter,
		Response: response,
	}
}

func (m *Message) CheckUpdate(b *gotdbot.Client, ctx *gotdbot.Context) bool {
	if ctx.EffectiveMessage == nil {
		return false
	}
	return m.Filter(ctx.EffectiveMessage)
}

func (m *Message) HandleUpdate(b *gotdbot.Client, ctx *gotdbot.Context) error {
	return m.Response(b, ctx)
}
