package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageSendAcknowledged struct {
	Filter   filters.UpdateMessageSendAcknowledged
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageSendAcknowledged(filter filters.UpdateMessageSendAcknowledged, response func(ctx *ext.Context) error) *UpdateMessageSendAcknowledged {
	return &UpdateMessageSendAcknowledged{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageSendAcknowledged) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateMessageSendAcknowledged)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageSendAcknowledged) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
