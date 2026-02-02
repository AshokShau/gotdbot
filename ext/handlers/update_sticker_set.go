package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStickerSet struct {
	Filter   filters.UpdateStickerSet
	Response func(ctx *ext.Context) error
}

func NewUpdateStickerSet(filter filters.UpdateStickerSet, response func(ctx *ext.Context) error) *UpdateStickerSet {
	return &UpdateStickerSet{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStickerSet) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateStickerSet
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStickerSet) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
