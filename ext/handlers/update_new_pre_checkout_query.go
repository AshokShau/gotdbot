package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewPreCheckoutQuery struct {
	Filter   filters.UpdateNewPreCheckoutQuery
	Response func(ctx *ext.Context) error
}

func NewUpdateNewPreCheckoutQuery(filter filters.UpdateNewPreCheckoutQuery, response func(ctx *ext.Context) error) *UpdateNewPreCheckoutQuery {
	return &UpdateNewPreCheckoutQuery{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewPreCheckoutQuery) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewPreCheckoutQuery
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewPreCheckoutQuery) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
