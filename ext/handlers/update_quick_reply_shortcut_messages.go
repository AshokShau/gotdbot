package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateQuickReplyShortcutMessages struct {
	Filter   filters.UpdateQuickReplyShortcutMessages
	Response func(ctx *ext.Context) error
}

func NewUpdateQuickReplyShortcutMessages(filter filters.UpdateQuickReplyShortcutMessages, response func(ctx *ext.Context) error) *UpdateQuickReplyShortcutMessages {
	return &UpdateQuickReplyShortcutMessages{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateQuickReplyShortcutMessages) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateQuickReplyShortcutMessages)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateQuickReplyShortcutMessages) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
