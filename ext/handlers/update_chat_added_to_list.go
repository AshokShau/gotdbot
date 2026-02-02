package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatAddedToList struct {
	Filter   filters.UpdateChatAddedToList
	Response func(ctx *ext.Context) error
}

func NewUpdateChatAddedToList(filter filters.UpdateChatAddedToList, response func(ctx *ext.Context) error) *UpdateChatAddedToList {
	return &UpdateChatAddedToList{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatAddedToList) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatAddedToList
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatAddedToList) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
