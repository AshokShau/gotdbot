package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatVideoChat struct {
	Filter   filters.UpdateChatVideoChat
	Response func(ctx *ext.Context) error
}

func NewUpdateChatVideoChat(filter filters.UpdateChatVideoChat, response func(ctx *ext.Context) error) *UpdateChatVideoChat {
	return &UpdateChatVideoChat{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatVideoChat) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatVideoChat
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatVideoChat) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
