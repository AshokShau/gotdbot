package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

type Message struct {
	Filter        filters.Message
	Response      func(b *gotdbot.Client, u *gotdbot.Message) error
	AllowOutgoing bool
}

func NewMessage(filter filters.Message, response func(b *gotdbot.Client, u *gotdbot.Message) error) *Message {
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

func (m *Message) CheckUpdate(b *gotdbot.Client, update gotdbot.TlObject) bool {
	msg := gotdbot.ExtractMessage(update)
	if msg == nil {
		return false
	}
	if msg.IsOutgoing && !m.AllowOutgoing {
		return false
	}

	if m.Filter == nil {
		return true
	}

	return m.Filter(msg)
}

func (m *Message) HandleUpdate(b *gotdbot.Client, update gotdbot.TlObject) error {
	return m.Response(b, gotdbot.ExtractMessage(update))
}
