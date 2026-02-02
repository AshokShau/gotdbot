package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateQuickReplyShortcut struct {
	Filter   filters.UpdateQuickReplyShortcut
	Response func(ctx *ext.Context) error
}

func NewUpdateQuickReplyShortcut(filter filters.UpdateQuickReplyShortcut, response func(ctx *ext.Context) error) *UpdateQuickReplyShortcut {
	return &UpdateQuickReplyShortcut{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateQuickReplyShortcut) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateQuickReplyShortcut
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateQuickReplyShortcut) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
