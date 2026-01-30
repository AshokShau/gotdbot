package handlers

import (
	"github.com/AshokShau/gotdbot"
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
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateQuickReplyShortcut)
	if !ok {
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
