package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatPermissions struct {
	Filter   filters.UpdateChatPermissions
	Response func(ctx *ext.Context) error
}

func NewUpdateChatPermissions(filter filters.UpdateChatPermissions, response func(ctx *ext.Context) error) *UpdateChatPermissions {
	return &UpdateChatPermissions{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatPermissions) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatPermissions)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatPermissions) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
