package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageMentionRead struct {
	Filter   filters.UpdateMessageMentionRead
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageMentionRead(filter filters.UpdateMessageMentionRead, response func(ctx *ext.Context) error) *UpdateMessageMentionRead {
	return &UpdateMessageMentionRead{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageMentionRead) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateMessageMentionRead)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageMentionRead) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
