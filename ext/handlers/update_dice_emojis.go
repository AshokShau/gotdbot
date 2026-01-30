package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateDiceEmojis struct {
	Filter   filters.UpdateDiceEmojis
	Response func(ctx *ext.Context) error
}

func NewUpdateDiceEmojis(filter filters.UpdateDiceEmojis, response func(ctx *ext.Context) error) *UpdateDiceEmojis {
	return &UpdateDiceEmojis{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateDiceEmojis) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateDiceEmojis)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateDiceEmojis) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
