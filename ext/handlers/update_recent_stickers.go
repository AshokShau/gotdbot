package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateRecentStickers struct {
	Filter   filters.UpdateRecentStickers
	Response func(ctx *ext.Context) error
}

func NewUpdateRecentStickers(filter filters.UpdateRecentStickers, response func(ctx *ext.Context) error) *UpdateRecentStickers {
	return &UpdateRecentStickers{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateRecentStickers) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateRecentStickers)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateRecentStickers) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
