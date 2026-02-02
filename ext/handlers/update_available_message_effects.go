package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAvailableMessageEffects struct {
	Filter   filters.UpdateAvailableMessageEffects
	Response func(ctx *ext.Context) error
}

func NewUpdateAvailableMessageEffects(filter filters.UpdateAvailableMessageEffects, response func(ctx *ext.Context) error) *UpdateAvailableMessageEffects {
	return &UpdateAvailableMessageEffects{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAvailableMessageEffects) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateAvailableMessageEffects
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAvailableMessageEffects) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
