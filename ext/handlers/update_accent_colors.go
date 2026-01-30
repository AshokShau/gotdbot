package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAccentColors struct {
	Filter   filters.UpdateAccentColors
	Response func(ctx *ext.Context) error
}

func NewUpdateAccentColors(filter filters.UpdateAccentColors, response func(ctx *ext.Context) error) *UpdateAccentColors {
	return &UpdateAccentColors{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAccentColors) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateAccentColors)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAccentColors) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
