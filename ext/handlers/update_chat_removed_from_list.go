package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatRemovedFromList struct {
	Filter   filters.UpdateChatRemovedFromList
	Response func(ctx *ext.Context) error
}

func NewUpdateChatRemovedFromList(filter filters.UpdateChatRemovedFromList, response func(ctx *ext.Context) error) *UpdateChatRemovedFromList {
	return &UpdateChatRemovedFromList{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatRemovedFromList) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatRemovedFromList)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatRemovedFromList) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
