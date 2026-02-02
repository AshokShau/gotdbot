package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateTrendingStickerSets struct {
	Filter   filters.UpdateTrendingStickerSets
	Response func(ctx *ext.Context) error
}

func NewUpdateTrendingStickerSets(filter filters.UpdateTrendingStickerSets, response func(ctx *ext.Context) error) *UpdateTrendingStickerSets {
	return &UpdateTrendingStickerSets{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateTrendingStickerSets) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateTrendingStickerSets
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateTrendingStickerSets) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
