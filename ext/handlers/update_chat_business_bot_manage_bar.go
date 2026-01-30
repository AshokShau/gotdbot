package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatBusinessBotManageBar struct {
	Filter   filters.UpdateChatBusinessBotManageBar
	Response func(ctx *ext.Context) error
}

func NewUpdateChatBusinessBotManageBar(filter filters.UpdateChatBusinessBotManageBar, response func(ctx *ext.Context) error) *UpdateChatBusinessBotManageBar {
	return &UpdateChatBusinessBotManageBar{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatBusinessBotManageBar) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatBusinessBotManageBar)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatBusinessBotManageBar) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
