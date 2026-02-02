package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatReadInbox struct {
	Filter   filters.UpdateChatReadInbox
	Response func(ctx *ext.Context) error
}

func NewUpdateChatReadInbox(filter filters.UpdateChatReadInbox, response func(ctx *ext.Context) error) *UpdateChatReadInbox {
	return &UpdateChatReadInbox{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatReadInbox) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatReadInbox
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatReadInbox) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
