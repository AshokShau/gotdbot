package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageReaction struct {
	Filter   filters.UpdateMessageReaction
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageReaction(filter filters.UpdateMessageReaction, response func(ctx *ext.Context) error) *UpdateMessageReaction {
	return &UpdateMessageReaction{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageReaction) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateMessageReaction)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageReaction) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
