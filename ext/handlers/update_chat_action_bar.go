package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatActionBar struct {
	Filter   filters.UpdateChatActionBar
	Response func(ctx *ext.Context) error
}

func NewUpdateChatActionBar(filter filters.UpdateChatActionBar, response func(ctx *ext.Context) error) *UpdateChatActionBar {
	return &UpdateChatActionBar{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatActionBar) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatActionBar
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatActionBar) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
