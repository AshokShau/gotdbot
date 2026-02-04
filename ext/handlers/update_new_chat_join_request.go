package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewChatJoinRequest struct {
	Filter   filters.UpdateNewChatJoinRequest
	Response func(ctx *ext.Context) error
}

func NewUpdateNewChatJoinRequest(filter filters.UpdateNewChatJoinRequest, response func(ctx *ext.Context) error) *UpdateNewChatJoinRequest {
	return &UpdateNewChatJoinRequest{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewChatJoinRequest) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewChatJoinRequest
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewChatJoinRequest) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
