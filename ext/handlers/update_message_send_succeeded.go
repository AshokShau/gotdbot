package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageSendSucceeded struct {
	Filter   filters.UpdateMessageSendSucceeded
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageSendSucceeded(filter filters.UpdateMessageSendSucceeded, response func(ctx *ext.Context) error) *UpdateMessageSendSucceeded {
	return &UpdateMessageSendSucceeded{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageSendSucceeded) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageSendSucceeded
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageSendSucceeded) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
