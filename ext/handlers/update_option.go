package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateOption struct {
	Filter   filters.UpdateOption
	Response func(ctx *ext.Context) error
}

func NewUpdateOption(filter filters.UpdateOption, response func(ctx *ext.Context) error) *UpdateOption {
	return &UpdateOption{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateOption) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateOption
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateOption) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
