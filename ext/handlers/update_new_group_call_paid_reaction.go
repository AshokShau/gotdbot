package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewGroupCallPaidReaction struct {
	Filter   filters.UpdateNewGroupCallPaidReaction
	Response func(ctx *ext.Context) error
}

func NewUpdateNewGroupCallPaidReaction(filter filters.UpdateNewGroupCallPaidReaction, response func(ctx *ext.Context) error) *UpdateNewGroupCallPaidReaction {
	return &UpdateNewGroupCallPaidReaction{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewGroupCallPaidReaction) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewGroupCallPaidReaction
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewGroupCallPaidReaction) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
