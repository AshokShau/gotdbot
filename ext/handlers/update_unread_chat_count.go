package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUnreadChatCount struct {
	Filter   filters.UpdateUnreadChatCount
	Response func(ctx *ext.Context) error
}

func NewUpdateUnreadChatCount(filter filters.UpdateUnreadChatCount, response func(ctx *ext.Context) error) *UpdateUnreadChatCount {
	return &UpdateUnreadChatCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUnreadChatCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateUnreadChatCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUnreadChatCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
