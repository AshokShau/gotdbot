package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateActiveNotifications struct {
	Filter   filters.UpdateActiveNotifications
	Response func(ctx *ext.Context) error
}

func NewUpdateActiveNotifications(filter filters.UpdateActiveNotifications, response func(ctx *ext.Context) error) *UpdateActiveNotifications {
	return &UpdateActiveNotifications{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateActiveNotifications) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateActiveNotifications)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateActiveNotifications) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
