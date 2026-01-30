package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateProfileAccentColors struct {
	Filter   filters.UpdateProfileAccentColors
	Response func(ctx *ext.Context) error
}

func NewUpdateProfileAccentColors(filter filters.UpdateProfileAccentColors, response func(ctx *ext.Context) error) *UpdateProfileAccentColors {
	return &UpdateProfileAccentColors{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateProfileAccentColors) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateProfileAccentColors)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateProfileAccentColors) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
