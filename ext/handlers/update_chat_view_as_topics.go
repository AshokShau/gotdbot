package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatViewAsTopics struct {
	Filter   filters.UpdateChatViewAsTopics
	Response func(ctx *ext.Context) error
}

func NewUpdateChatViewAsTopics(filter filters.UpdateChatViewAsTopics, response func(ctx *ext.Context) error) *UpdateChatViewAsTopics {
	return &UpdateChatViewAsTopics{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatViewAsTopics) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatViewAsTopics)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatViewAsTopics) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
