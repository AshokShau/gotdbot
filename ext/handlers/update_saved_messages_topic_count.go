package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSavedMessagesTopicCount struct {
	Filter   filters.UpdateSavedMessagesTopicCount
	Response func(ctx *ext.Context) error
}

func NewUpdateSavedMessagesTopicCount(filter filters.UpdateSavedMessagesTopicCount, response func(ctx *ext.Context) error) *UpdateSavedMessagesTopicCount {
	return &UpdateSavedMessagesTopicCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSavedMessagesTopicCount) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSavedMessagesTopicCount)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSavedMessagesTopicCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
