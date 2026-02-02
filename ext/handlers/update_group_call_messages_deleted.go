package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallMessagesDeleted struct {
	Filter   filters.UpdateGroupCallMessagesDeleted
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallMessagesDeleted(filter filters.UpdateGroupCallMessagesDeleted, response func(ctx *ext.Context) error) *UpdateGroupCallMessagesDeleted {
	return &UpdateGroupCallMessagesDeleted{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallMessagesDeleted) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateGroupCallMessagesDeleted
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallMessagesDeleted) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
