package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateTopicMessageCount struct {
	Filter   filters.UpdateTopicMessageCount
	Response func(ctx *ext.Context) error
}

func NewUpdateTopicMessageCount(filter filters.UpdateTopicMessageCount, response func(ctx *ext.Context) error) *UpdateTopicMessageCount {
	return &UpdateTopicMessageCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateTopicMessageCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateTopicMessageCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateTopicMessageCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
