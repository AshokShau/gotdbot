package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatPosition struct {
	Filter   filters.UpdateChatPosition
	Response func(ctx *ext.Context) error
}

func NewUpdateChatPosition(filter filters.UpdateChatPosition, response func(ctx *ext.Context) error) *UpdateChatPosition {
	return &UpdateChatPosition{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatPosition) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatPosition)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatPosition) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
