package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageReactions struct {
	Filter   filters.UpdateMessageReactions
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageReactions(filter filters.UpdateMessageReactions, response func(ctx *ext.Context) error) *UpdateMessageReactions {
	return &UpdateMessageReactions{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageReactions) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageReactions
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageReactions) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
