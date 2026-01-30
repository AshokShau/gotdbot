package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateHavePendingNotifications struct {
	Filter   filters.UpdateHavePendingNotifications
	Response func(ctx *ext.Context) error
}

func NewUpdateHavePendingNotifications(filter filters.UpdateHavePendingNotifications, response func(ctx *ext.Context) error) *UpdateHavePendingNotifications {
	return &UpdateHavePendingNotifications{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateHavePendingNotifications) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateHavePendingNotifications)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateHavePendingNotifications) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
