package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNotification struct {
	Filter   filters.UpdateNotification
	Response func(ctx *ext.Context) error
}

func NewUpdateNotification(filter filters.UpdateNotification, response func(ctx *ext.Context) error) *UpdateNotification {
	return &UpdateNotification{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNotification) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNotification
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNotification) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
