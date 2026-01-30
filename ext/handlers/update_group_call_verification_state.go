package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallVerificationState struct {
	Filter   filters.UpdateGroupCallVerificationState
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallVerificationState(filter filters.UpdateGroupCallVerificationState, response func(ctx *ext.Context) error) *UpdateGroupCallVerificationState {
	return &UpdateGroupCallVerificationState{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallVerificationState) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateGroupCallVerificationState)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallVerificationState) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
