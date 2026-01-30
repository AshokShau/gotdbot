package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatReplyMarkup struct {
	Filter   filters.UpdateChatReplyMarkup
	Response func(ctx *ext.Context) error
}

func NewUpdateChatReplyMarkup(filter filters.UpdateChatReplyMarkup, response func(ctx *ext.Context) error) *UpdateChatReplyMarkup {
	return &UpdateChatReplyMarkup{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatReplyMarkup) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatReplyMarkup)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatReplyMarkup) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
