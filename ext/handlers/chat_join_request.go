package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type ChatJoinRequest struct {
	Filter   filters.ChatJoinRequest
	Response func(ctx *ext.Context) error
}

func NewChatJoinRequest(filter filters.ChatJoinRequest, response func(ctx *ext.Context) error) *ChatJoinRequest {
	return &ChatJoinRequest{
		Filter:   filter,
		Response: response,
	}
}

func (h *ChatJoinRequest) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateNewChatJoinRequest)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *ChatJoinRequest) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
