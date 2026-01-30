package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatDefaultDisableNotification struct {
	Filter   filters.UpdateChatDefaultDisableNotification
	Response func(ctx *ext.Context) error
}

func NewUpdateChatDefaultDisableNotification(filter filters.UpdateChatDefaultDisableNotification, response func(ctx *ext.Context) error) *UpdateChatDefaultDisableNotification {
	return &UpdateChatDefaultDisableNotification{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatDefaultDisableNotification) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatDefaultDisableNotification)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatDefaultDisableNotification) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
