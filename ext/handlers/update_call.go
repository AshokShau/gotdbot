package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateCall struct {
	Filter   filters.UpdateCall
	Response func(ctx *ext.Context) error
}

func NewUpdateCall(filter filters.UpdateCall, response func(ctx *ext.Context) error) *UpdateCall {
	return &UpdateCall{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateCall) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateCall)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateCall) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
