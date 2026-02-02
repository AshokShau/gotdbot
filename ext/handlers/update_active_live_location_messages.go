package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateActiveLiveLocationMessages struct {
	Filter   filters.UpdateActiveLiveLocationMessages
	Response func(ctx *ext.Context) error
}

func NewUpdateActiveLiveLocationMessages(filter filters.UpdateActiveLiveLocationMessages, response func(ctx *ext.Context) error) *UpdateActiveLiveLocationMessages {
	return &UpdateActiveLiveLocationMessages{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateActiveLiveLocationMessages) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateActiveLiveLocationMessages
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateActiveLiveLocationMessages) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
