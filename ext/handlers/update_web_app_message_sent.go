package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateWebAppMessageSent struct {
	Filter   filters.UpdateWebAppMessageSent
	Response func(ctx *ext.Context) error
}

func NewUpdateWebAppMessageSent(filter filters.UpdateWebAppMessageSent, response func(ctx *ext.Context) error) *UpdateWebAppMessageSent {
	return &UpdateWebAppMessageSent{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateWebAppMessageSent) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateWebAppMessageSent
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateWebAppMessageSent) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
