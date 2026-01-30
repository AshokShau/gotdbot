package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFavoriteStickers struct {
	Filter   filters.UpdateFavoriteStickers
	Response func(ctx *ext.Context) error
}

func NewUpdateFavoriteStickers(filter filters.UpdateFavoriteStickers, response func(ctx *ext.Context) error) *UpdateFavoriteStickers {
	return &UpdateFavoriteStickers{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFavoriteStickers) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateFavoriteStickers)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFavoriteStickers) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
