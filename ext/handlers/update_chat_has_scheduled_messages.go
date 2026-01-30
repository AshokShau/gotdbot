package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatHasScheduledMessages struct {
	Filter   filters.UpdateChatHasScheduledMessages
	Response func(ctx *ext.Context) error
}

func NewUpdateChatHasScheduledMessages(filter filters.UpdateChatHasScheduledMessages, response func(ctx *ext.Context) error) *UpdateChatHasScheduledMessages {
	return &UpdateChatHasScheduledMessages{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatHasScheduledMessages) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatHasScheduledMessages)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatHasScheduledMessages) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
