package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

type Message struct {
	Filter        filters.Message
	Response      func(b *gotdbot.Client, ctx *gotdbot.Context) error
	AllowOutgoing bool
}

func NewMessage(filter filters.Message, response func(b *gotdbot.Client, ctx *gotdbot.Context) error) *Message {
	return &Message{
		Filter:        filter,
		Response:      response,
		AllowOutgoing: false,
	}
}

func (m *Message) SetAllowOutgoing(allow bool) *Message {
	m.AllowOutgoing = allow
	return m
}

func (m *Message) CheckUpdate(b *gotdbot.Client, ctx *gotdbot.Context) bool {
	if ctx.EffectiveMessage == nil {
		return false
	}
	if ctx.EffectiveMessage.IsOutgoing && !m.AllowOutgoing {
		return false
	}

	if m.Filter == nil {
		return true
	}

	return m.Filter(ctx.EffectiveMessage)
}

func (m *Message) HandleUpdate(b *gotdbot.Client, ctx *gotdbot.Context) error {
	return m.Response(b, ctx)
}
