package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateOwnedStarCount struct {
	Filter   filters.UpdateOwnedStarCount
	Response func(ctx *ext.Context) error
}

func NewUpdateOwnedStarCount(filter filters.UpdateOwnedStarCount, response func(ctx *ext.Context) error) *UpdateOwnedStarCount {
	return &UpdateOwnedStarCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateOwnedStarCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateOwnedStarCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateOwnedStarCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
