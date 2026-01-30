package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallParticipants struct {
	Filter   filters.UpdateGroupCallParticipants
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallParticipants(filter filters.UpdateGroupCallParticipants, response func(ctx *ext.Context) error) *UpdateGroupCallParticipants {
	return &UpdateGroupCallParticipants{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallParticipants) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateGroupCallParticipants)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallParticipants) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
