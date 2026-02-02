package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatOnlineMemberCount struct {
	Filter   filters.UpdateChatOnlineMemberCount
	Response func(ctx *ext.Context) error
}

func NewUpdateChatOnlineMemberCount(filter filters.UpdateChatOnlineMemberCount, response func(ctx *ext.Context) error) *UpdateChatOnlineMemberCount {
	return &UpdateChatOnlineMemberCount{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatOnlineMemberCount) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatOnlineMemberCount
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatOnlineMemberCount) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
