package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStoryListChatCount struct {
	Filter   filters.UpdateStoryListChatCount
	Response func(ctx *ext.Context) error
}

func NewUpdateStoryListChatCount(filter filters.UpdateStoryListChatCount, response func(ctx *ext.Context) error) *UpdateStoryListChatCount {
	return &UpdateStoryListChatCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStoryListChatCount) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStoryListChatCount)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStoryListChatCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
