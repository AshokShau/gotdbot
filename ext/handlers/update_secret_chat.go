package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSecretChat struct {
	Filter   filters.UpdateSecretChat
	Response func(ctx *ext.Context) error
}

func NewUpdateSecretChat(filter filters.UpdateSecretChat, response func(ctx *ext.Context) error) *UpdateSecretChat {
	return &UpdateSecretChat{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSecretChat) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSecretChat)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSecretChat) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
