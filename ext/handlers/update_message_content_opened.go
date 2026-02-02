package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageContentOpened struct {
	Filter   filters.UpdateMessageContentOpened
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageContentOpened(filter filters.UpdateMessageContentOpened, response func(ctx *ext.Context) error) *UpdateMessageContentOpened {
	return &UpdateMessageContentOpened{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageContentOpened) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateMessageContentOpened
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageContentOpened) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
