package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStakeDiceState struct {
	Filter   filters.UpdateStakeDiceState
	Response func(ctx *ext.Context) error
}

func NewUpdateStakeDiceState(filter filters.UpdateStakeDiceState, response func(ctx *ext.Context) error) *UpdateStakeDiceState {
	return &UpdateStakeDiceState{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStakeDiceState) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStakeDiceState)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStakeDiceState) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
