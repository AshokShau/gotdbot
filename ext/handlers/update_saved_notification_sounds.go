package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSavedNotificationSounds struct {
	Filter   filters.UpdateSavedNotificationSounds
	Response func(ctx *ext.Context) error
}

func NewUpdateSavedNotificationSounds(filter filters.UpdateSavedNotificationSounds, response func(ctx *ext.Context) error) *UpdateSavedNotificationSounds {
	return &UpdateSavedNotificationSounds{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSavedNotificationSounds) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSavedNotificationSounds)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSavedNotificationSounds) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
