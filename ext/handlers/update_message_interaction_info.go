package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateMessageInteractionInfo struct {
	Filter   filters.UpdateMessageInteractionInfo
	Response func(ctx *ext.Context) error
}

func NewUpdateMessageInteractionInfo(filter filters.UpdateMessageInteractionInfo, response func(ctx *ext.Context) error) *UpdateMessageInteractionInfo {
	return &UpdateMessageInteractionInfo{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateMessageInteractionInfo) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateMessageInteractionInfo)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateMessageInteractionInfo) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
