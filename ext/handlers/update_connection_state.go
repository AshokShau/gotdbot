package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateConnectionState struct {
	Filter   filters.UpdateConnectionState
	Response func(ctx *ext.Context) error
}

func NewUpdateConnectionState(filter filters.UpdateConnectionState, response func(ctx *ext.Context) error) *UpdateConnectionState {
	return &UpdateConnectionState{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateConnectionState) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateConnectionState)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateConnectionState) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
