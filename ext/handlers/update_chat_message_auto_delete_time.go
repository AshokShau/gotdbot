package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatMessageAutoDeleteTime struct {
	Filter   filters.UpdateChatMessageAutoDeleteTime
	Response func(ctx *ext.Context) error
}

func NewUpdateChatMessageAutoDeleteTime(filter filters.UpdateChatMessageAutoDeleteTime, response func(ctx *ext.Context) error) *UpdateChatMessageAutoDeleteTime {
	return &UpdateChatMessageAutoDeleteTime{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatMessageAutoDeleteTime) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatMessageAutoDeleteTime)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatMessageAutoDeleteTime) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
