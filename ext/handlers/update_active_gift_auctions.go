package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateActiveGiftAuctions struct {
	Filter   filters.UpdateActiveGiftAuctions
	Response func(ctx *ext.Context) error
}

func NewUpdateActiveGiftAuctions(filter filters.UpdateActiveGiftAuctions, response func(ctx *ext.Context) error) *UpdateActiveGiftAuctions {
	return &UpdateActiveGiftAuctions{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateActiveGiftAuctions) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateActiveGiftAuctions)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateActiveGiftAuctions) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
