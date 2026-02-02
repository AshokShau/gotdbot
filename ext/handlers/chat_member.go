package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type ChatMember struct {
	Filter   filters.ChatMember
	Response func(ctx *ext.Context) error
}

func NewChatMember(filter filters.ChatMember, response func(ctx *ext.Context) error) *ChatMember {
	return &ChatMember{
		Filter:   filter,
		Response: response,
	}
}

func (h *ChatMember) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdateChatMember
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *ChatMember) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
