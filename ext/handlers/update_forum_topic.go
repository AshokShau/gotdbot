package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateForumTopic struct {
	Filter   filters.UpdateForumTopic
	Response func(ctx *ext.Context) error
}

func NewUpdateForumTopic(filter filters.UpdateForumTopic, response func(ctx *ext.Context) error) *UpdateForumTopic {
	return &UpdateForumTopic{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateForumTopic) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateForumTopic
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateForumTopic) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
