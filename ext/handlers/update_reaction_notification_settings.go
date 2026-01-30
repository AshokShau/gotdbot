package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateReactionNotificationSettings struct {
	Filter   filters.UpdateReactionNotificationSettings
	Response func(ctx *ext.Context) error
}

func NewUpdateReactionNotificationSettings(filter filters.UpdateReactionNotificationSettings, response func(ctx *ext.Context) error) *UpdateReactionNotificationSettings {
	return &UpdateReactionNotificationSettings{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateReactionNotificationSettings) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateReactionNotificationSettings)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateReactionNotificationSettings) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
