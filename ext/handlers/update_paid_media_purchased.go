package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdatePaidMediaPurchased struct {
	Filter   filters.UpdatePaidMediaPurchased
	Response func(ctx *ext.Context) error
}

func NewUpdatePaidMediaPurchased(filter filters.UpdatePaidMediaPurchased, response func(ctx *ext.Context) error) *UpdatePaidMediaPurchased {
	return &UpdatePaidMediaPurchased{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdatePaidMediaPurchased) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdatePaidMediaPurchased)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdatePaidMediaPurchased) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
