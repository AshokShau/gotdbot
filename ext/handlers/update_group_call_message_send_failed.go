package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallMessageSendFailed struct {
	Filter   filters.UpdateGroupCallMessageSendFailed
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallMessageSendFailed(filter filters.UpdateGroupCallMessageSendFailed, response func(ctx *ext.Context) error) *UpdateGroupCallMessageSendFailed {
	return &UpdateGroupCallMessageSendFailed{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallMessageSendFailed) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateGroupCallMessageSendFailed)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallMessageSendFailed) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
