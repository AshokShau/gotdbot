package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatUnreadReactionCount struct {
	Filter   filters.UpdateChatUnreadReactionCount
	Response func(ctx *ext.Context) error
}

func NewUpdateChatUnreadReactionCount(filter filters.UpdateChatUnreadReactionCount, response func(ctx *ext.Context) error) *UpdateChatUnreadReactionCount {
	return &UpdateChatUnreadReactionCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatUnreadReactionCount) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatUnreadReactionCount)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatUnreadReactionCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
