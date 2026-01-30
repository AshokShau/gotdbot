package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatBlockList struct {
	Filter   filters.UpdateChatBlockList
	Response func(ctx *ext.Context) error
}

func NewUpdateChatBlockList(filter filters.UpdateChatBlockList, response func(ctx *ext.Context) error) *UpdateChatBlockList {
	return &UpdateChatBlockList{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatBlockList) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatBlockList)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatBlockList) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
