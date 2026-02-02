package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatNotificationSettings struct {
	Filter   filters.UpdateChatNotificationSettings
	Response func(ctx *ext.Context) error
}

func NewUpdateChatNotificationSettings(filter filters.UpdateChatNotificationSettings, response func(ctx *ext.Context) error) *UpdateChatNotificationSettings {
	return &UpdateChatNotificationSettings{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatNotificationSettings) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatNotificationSettings
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatNotificationSettings) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
