package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageSuggestedPostInfo struct {
	Filter   filters.UpdateMessageSuggestedPostInfo
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageSuggestedPostInfo(filter filters.UpdateMessageSuggestedPostInfo, response func(ctx *ext.Context) error) *UpdateMessageSuggestedPostInfo {
	return &UpdateMessageSuggestedPostInfo{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageSuggestedPostInfo) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageSuggestedPostInfo
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageSuggestedPostInfo) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
