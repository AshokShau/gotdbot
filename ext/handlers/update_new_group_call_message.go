package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewGroupCallMessage struct {
	Filter   filters.UpdateNewGroupCallMessage
	Response func(ctx *ext.Context) error
}

func NewUpdateNewGroupCallMessage(filter filters.UpdateNewGroupCallMessage, response func(ctx *ext.Context) error) *UpdateNewGroupCallMessage {
	return &UpdateNewGroupCallMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewGroupCallMessage) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateNewGroupCallMessage)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewGroupCallMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
