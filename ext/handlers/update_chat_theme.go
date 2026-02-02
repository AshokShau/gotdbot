package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatTheme struct {
	Filter   filters.UpdateChatTheme
	Response func(ctx *ext.Context) error
}

func NewUpdateChatTheme(filter filters.UpdateChatTheme, response func(ctx *ext.Context) error) *UpdateChatTheme {
	return &UpdateChatTheme{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatTheme) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatTheme
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatTheme) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
