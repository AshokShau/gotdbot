package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallParticipant struct {
	Filter   filters.UpdateGroupCallParticipant
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallParticipant(filter filters.UpdateGroupCallParticipant, response func(ctx *ext.Context) error) *UpdateGroupCallParticipant {
	return &UpdateGroupCallParticipant{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallParticipant) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateGroupCallParticipant
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallParticipant) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
