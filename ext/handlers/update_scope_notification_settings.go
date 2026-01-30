package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateScopeNotificationSettings struct {
	Filter   filters.UpdateScopeNotificationSettings
	Response func(ctx *ext.Context) error
}

func NewUpdateScopeNotificationSettings(filter filters.UpdateScopeNotificationSettings, response func(ctx *ext.Context) error) *UpdateScopeNotificationSettings {
	return &UpdateScopeNotificationSettings{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateScopeNotificationSettings) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateScopeNotificationSettings)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateScopeNotificationSettings) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
