package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStarRevenueStatus struct {
	Filter   filters.UpdateStarRevenueStatus
	Response func(ctx *ext.Context) error
}

func NewUpdateStarRevenueStatus(filter filters.UpdateStarRevenueStatus, response func(ctx *ext.Context) error) *UpdateStarRevenueStatus {
	return &UpdateStarRevenueStatus{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStarRevenueStatus) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStarRevenueStatus)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStarRevenueStatus) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
