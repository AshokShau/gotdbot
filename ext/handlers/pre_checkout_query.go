package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type PreCheckoutQuery struct {
	Filter   filters.PreCheckoutQuery
	Response func(ctx *ext.Context) error
}

func NewPreCheckoutQuery(filter filters.PreCheckoutQuery, response func(ctx *ext.Context) error) *PreCheckoutQuery {
	return &PreCheckoutQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *PreCheckoutQuery) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdateNewPreCheckoutQuery
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *PreCheckoutQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
