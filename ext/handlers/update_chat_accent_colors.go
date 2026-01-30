package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatAccentColors struct {
	Filter   filters.UpdateChatAccentColors
	Response func(ctx *ext.Context) error
}

func NewUpdateChatAccentColors(filter filters.UpdateChatAccentColors, response func(ctx *ext.Context) error) *UpdateChatAccentColors {
	return &UpdateChatAccentColors{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatAccentColors) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatAccentColors)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatAccentColors) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
