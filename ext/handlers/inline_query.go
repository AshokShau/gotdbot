package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type InlineQuery struct {
	Filter   filters.InlineQuery
	Response func(ctx *ext.Context) error
}

func NewInlineQuery(filter filters.InlineQuery, response func(ctx *ext.Context) error) *InlineQuery {
	return &InlineQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *InlineQuery) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdateNewInlineQuery
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *InlineQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
