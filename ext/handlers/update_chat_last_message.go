package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatLastMessage struct {
	Filter   filters.UpdateChatLastMessage
	Response func(ctx *ext.Context) error
}

func NewUpdateChatLastMessage(filter filters.UpdateChatLastMessage, response func(ctx *ext.Context) error) *UpdateChatLastMessage {
	return &UpdateChatLastMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatLastMessage) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatLastMessage)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatLastMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
