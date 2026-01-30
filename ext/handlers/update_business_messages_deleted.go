package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateBusinessMessagesDeleted struct {
	Filter   filters.UpdateBusinessMessagesDeleted
	Response func(ctx *ext.Context) error
}

func NewUpdateBusinessMessagesDeleted(filter filters.UpdateBusinessMessagesDeleted, response func(ctx *ext.Context) error) *UpdateBusinessMessagesDeleted {
	return &UpdateBusinessMessagesDeleted{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateBusinessMessagesDeleted) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateBusinessMessagesDeleted)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateBusinessMessagesDeleted) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
