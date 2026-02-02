package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type ShippingQuery struct {
	Filter   filters.ShippingQuery
	Response func(ctx *ext.Context) error
}

func NewShippingQuery(filter filters.ShippingQuery, response func(ctx *ext.Context) error) *ShippingQuery {
	return &ShippingQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *ShippingQuery) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdateNewShippingQuery
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *ShippingQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
