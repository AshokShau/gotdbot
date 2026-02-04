package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdatePoll struct {
	Filter   filters.UpdatePoll
	Response func(ctx *ext.Context) error
}

func NewUpdatePoll(filter filters.UpdatePoll, response func(ctx *ext.Context) error) *UpdatePoll {
	return &UpdatePoll{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdatePoll) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdatePoll
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdatePoll) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
