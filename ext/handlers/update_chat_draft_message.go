package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatDraftMessage struct {
	Filter   filters.UpdateChatDraftMessage
	Response func(ctx *ext.Context) error
}

func NewUpdateChatDraftMessage(filter filters.UpdateChatDraftMessage, response func(ctx *ext.Context) error) *UpdateChatDraftMessage {
	return &UpdateChatDraftMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatDraftMessage) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatDraftMessage)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatDraftMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
