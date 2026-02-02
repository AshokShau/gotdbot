package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageContent struct {
	Filter   filters.UpdateMessageContent
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageContent(filter filters.UpdateMessageContent, response func(ctx *ext.Context) error) *UpdateMessageContent {
	return &UpdateMessageContent{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageContent) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageContent
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageContent) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
