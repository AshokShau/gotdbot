package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateQuickReplyShortcutDeleted struct {
	Filter   filters.UpdateQuickReplyShortcutDeleted
	Response func(ctx *ext.Context) error
}

func NewUpdateQuickReplyShortcutDeleted(filter filters.UpdateQuickReplyShortcutDeleted, response func(ctx *ext.Context) error) *UpdateQuickReplyShortcutDeleted {
	return &UpdateQuickReplyShortcutDeleted{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateQuickReplyShortcutDeleted) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateQuickReplyShortcutDeleted)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateQuickReplyShortcutDeleted) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
