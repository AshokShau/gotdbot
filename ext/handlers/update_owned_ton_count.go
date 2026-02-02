package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateOwnedTonCount struct {
	Filter   filters.UpdateOwnedTonCount
	Response func(ctx *ext.Context) error
}

func NewUpdateOwnedTonCount(filter filters.UpdateOwnedTonCount, response func(ctx *ext.Context) error) *UpdateOwnedTonCount {
	return &UpdateOwnedTonCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateOwnedTonCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateOwnedTonCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateOwnedTonCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
