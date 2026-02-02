package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatFolders struct {
	Filter   filters.UpdateChatFolders
	Response func(ctx *ext.Context) error
}

func NewUpdateChatFolders(filter filters.UpdateChatFolders, response func(ctx *ext.Context) error) *UpdateChatFolders {
	return &UpdateChatFolders{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatFolders) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatFolders
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatFolders) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
