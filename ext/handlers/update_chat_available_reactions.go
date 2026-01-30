package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatAvailableReactions struct {
	Filter   filters.UpdateChatAvailableReactions
	Response func(ctx *ext.Context) error
}

func NewUpdateChatAvailableReactions(filter filters.UpdateChatAvailableReactions, response func(ctx *ext.Context) error) *UpdateChatAvailableReactions {
	return &UpdateChatAvailableReactions{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatAvailableReactions) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatAvailableReactions)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatAvailableReactions) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
