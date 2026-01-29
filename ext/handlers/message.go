package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type Message struct {
	Filter   filters.Message
	Response func(ctx *ext.Context) error
}

func NewMessage(filter filters.Message, response func(ctx *ext.Context) error) *Message {
	return &Message{
		Filter:   filter,
		Response: response,
	}
}

func (m *Message) CheckUpdate(ctx *ext.Context) bool {
	if ctx.EffectiveMessage == nil {
		return false
	}
	return m.Filter(ctx.EffectiveMessage)
}

func (m *Message) HandleUpdate(ctx *ext.Context) error {
	return m.Response(ctx)
}
