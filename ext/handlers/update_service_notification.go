package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateServiceNotification struct {
	Filter   filters.UpdateServiceNotification
	Response func(ctx *ext.Context) error
}

func NewUpdateServiceNotification(filter filters.UpdateServiceNotification, response func(ctx *ext.Context) error) *UpdateServiceNotification {
	return &UpdateServiceNotification{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateServiceNotification) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateServiceNotification)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateServiceNotification) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
