package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateDirectMessagesChatTopic struct {
	Filter   filters.UpdateDirectMessagesChatTopic
	Response func(ctx *ext.Context) error
}

func NewUpdateDirectMessagesChatTopic(filter filters.UpdateDirectMessagesChatTopic, response func(ctx *ext.Context) error) *UpdateDirectMessagesChatTopic {
	return &UpdateDirectMessagesChatTopic{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateDirectMessagesChatTopic) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateDirectMessagesChatTopic)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateDirectMessagesChatTopic) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
