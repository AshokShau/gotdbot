package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageLiveLocationViewed struct {
	Filter   filters.UpdateMessageLiveLocationViewed
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageLiveLocationViewed(filter filters.UpdateMessageLiveLocationViewed, response func(ctx *ext.Context) error) *UpdateMessageLiveLocationViewed {
	return &UpdateMessageLiveLocationViewed{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageLiveLocationViewed) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageLiveLocationViewed
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageLiveLocationViewed) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
