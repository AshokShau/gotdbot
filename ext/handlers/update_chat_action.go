package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatAction struct {
	Filter   filters.UpdateChatAction
	Response func(ctx *ext.Context) error
}

func NewUpdateChatAction(filter filters.UpdateChatAction, response func(ctx *ext.Context) error) *UpdateChatAction {
	return &UpdateChatAction{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatAction) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatAction
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatAction) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
