package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSpeedLimitNotification struct {
	Filter   filters.UpdateSpeedLimitNotification
	Response func(ctx *ext.Context) error
}

func NewUpdateSpeedLimitNotification(filter filters.UpdateSpeedLimitNotification, response func(ctx *ext.Context) error) *UpdateSpeedLimitNotification {
	return &UpdateSpeedLimitNotification{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSpeedLimitNotification) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateSpeedLimitNotification
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSpeedLimitNotification) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
