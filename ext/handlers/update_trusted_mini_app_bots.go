package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateTrustedMiniAppBots struct {
	Filter   filters.UpdateTrustedMiniAppBots
	Response func(ctx *ext.Context) error
}

func NewUpdateTrustedMiniAppBots(filter filters.UpdateTrustedMiniAppBots, response func(ctx *ext.Context) error) *UpdateTrustedMiniAppBots {
	return &UpdateTrustedMiniAppBots{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateTrustedMiniAppBots) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateTrustedMiniAppBots
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateTrustedMiniAppBots) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
