package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateQuickReplyShortcuts struct {
	Filter   filters.UpdateQuickReplyShortcuts
	Response func(ctx *ext.Context) error
}

func NewUpdateQuickReplyShortcuts(filter filters.UpdateQuickReplyShortcuts, response func(ctx *ext.Context) error) *UpdateQuickReplyShortcuts {
	return &UpdateQuickReplyShortcuts{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateQuickReplyShortcuts) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateQuickReplyShortcuts)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateQuickReplyShortcuts) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
