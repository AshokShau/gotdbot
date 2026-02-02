package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUnreadMessageCount struct {
	Filter   filters.UpdateUnreadMessageCount
	Response func(ctx *ext.Context) error
}

func NewUpdateUnreadMessageCount(filter filters.UpdateUnreadMessageCount, response func(ctx *ext.Context) error) *UpdateUnreadMessageCount {
	return &UpdateUnreadMessageCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUnreadMessageCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateUnreadMessageCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUnreadMessageCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
