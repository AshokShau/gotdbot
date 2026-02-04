package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewCallbackQuery struct {
	Filter   filters.UpdateNewCallbackQuery
	Response func(ctx *ext.Context) error
}

func NewUpdateNewCallbackQuery(filter filters.UpdateNewCallbackQuery, response func(ctx *ext.Context) error) *UpdateNewCallbackQuery {
	return &UpdateNewCallbackQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewCallbackQuery) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewCallbackQuery
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewCallbackQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
