package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewBusinessMessage struct {
	Filter   filters.UpdateNewBusinessMessage
	Response func(ctx *ext.Context) error
}

func NewUpdateNewBusinessMessage(filter filters.UpdateNewBusinessMessage, response func(ctx *ext.Context) error) *UpdateNewBusinessMessage {
	return &UpdateNewBusinessMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewBusinessMessage) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateNewBusinessMessage)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewBusinessMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
