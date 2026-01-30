package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatBackground struct {
	Filter   filters.UpdateChatBackground
	Response func(ctx *ext.Context) error
}

func NewUpdateChatBackground(filter filters.UpdateChatBackground, response func(ctx *ext.Context) error) *UpdateChatBackground {
	return &UpdateChatBackground{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatBackground) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatBackground)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatBackground) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
