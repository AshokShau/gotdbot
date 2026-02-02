package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageFactCheck struct {
	Filter   filters.UpdateMessageFactCheck
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageFactCheck(filter filters.UpdateMessageFactCheck, response func(ctx *ext.Context) error) *UpdateMessageFactCheck {
	return &UpdateMessageFactCheck{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageFactCheck) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageFactCheck
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageFactCheck) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
