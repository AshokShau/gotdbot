package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateInstalledStickerSets struct {
	Filter   filters.UpdateInstalledStickerSets
	Response func(ctx *ext.Context) error
}

func NewUpdateInstalledStickerSets(filter filters.UpdateInstalledStickerSets, response func(ctx *ext.Context) error) *UpdateInstalledStickerSets {
	return &UpdateInstalledStickerSets{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateInstalledStickerSets) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateInstalledStickerSets
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateInstalledStickerSets) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
