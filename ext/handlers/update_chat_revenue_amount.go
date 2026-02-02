package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatRevenueAmount struct {
	Filter   filters.UpdateChatRevenueAmount
	Response func(ctx *ext.Context) error
}

func NewUpdateChatRevenueAmount(filter filters.UpdateChatRevenueAmount, response func(ctx *ext.Context) error) *UpdateChatRevenueAmount {
	return &UpdateChatRevenueAmount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatRevenueAmount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatRevenueAmount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatRevenueAmount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
