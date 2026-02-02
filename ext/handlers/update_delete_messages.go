package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateDeleteMessages struct {
	Filter   filters.UpdateDeleteMessages
	Response func(ctx *ext.Context) error
}

func NewUpdateDeleteMessages(filter filters.UpdateDeleteMessages, response func(ctx *ext.Context) error) *UpdateDeleteMessages {
	return &UpdateDeleteMessages{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateDeleteMessages) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateDeleteMessages
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateDeleteMessages) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
