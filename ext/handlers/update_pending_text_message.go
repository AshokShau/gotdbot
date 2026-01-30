package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdatePendingTextMessage struct {
	Filter   filters.UpdatePendingTextMessage
	Response func(ctx *ext.Context) error
}

func NewUpdatePendingTextMessage(filter filters.UpdatePendingTextMessage, response func(ctx *ext.Context) error) *UpdatePendingTextMessage {
	return &UpdatePendingTextMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdatePendingTextMessage) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdatePendingTextMessage)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdatePendingTextMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
