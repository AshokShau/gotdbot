package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatMessageSender struct {
	Filter   filters.UpdateChatMessageSender
	Response func(ctx *ext.Context) error
}

func NewUpdateChatMessageSender(filter filters.UpdateChatMessageSender, response func(ctx *ext.Context) error) *UpdateChatMessageSender {
	return &UpdateChatMessageSender{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatMessageSender) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatMessageSender)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatMessageSender) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
