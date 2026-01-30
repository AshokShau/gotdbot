package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatHasProtectedContent struct {
	Filter   filters.UpdateChatHasProtectedContent
	Response func(ctx *ext.Context) error
}

func NewUpdateChatHasProtectedContent(filter filters.UpdateChatHasProtectedContent, response func(ctx *ext.Context) error) *UpdateChatHasProtectedContent {
	return &UpdateChatHasProtectedContent{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatHasProtectedContent) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatHasProtectedContent)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatHasProtectedContent) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
