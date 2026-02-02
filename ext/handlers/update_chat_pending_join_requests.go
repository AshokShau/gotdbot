package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatPendingJoinRequests struct {
	Filter   filters.UpdateChatPendingJoinRequests
	Response func(ctx *ext.Context) error
}

func NewUpdateChatPendingJoinRequests(filter filters.UpdateChatPendingJoinRequests, response func(ctx *ext.Context) error) *UpdateChatPendingJoinRequests {
	return &UpdateChatPendingJoinRequests{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatPendingJoinRequests) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatPendingJoinRequests
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatPendingJoinRequests) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
