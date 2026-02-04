package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewInlineQuery struct {
	Filter   filters.UpdateNewInlineQuery
	Response func(ctx *ext.Context) error
}

func NewUpdateNewInlineQuery(filter filters.UpdateNewInlineQuery, response func(ctx *ext.Context) error) *UpdateNewInlineQuery {
	return &UpdateNewInlineQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewInlineQuery) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewInlineQuery
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewInlineQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
