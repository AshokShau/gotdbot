package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatTitle struct {
	Filter   filters.UpdateChatTitle
	Response func(ctx *ext.Context) error
}

func NewUpdateChatTitle(filter filters.UpdateChatTitle, response func(ctx *ext.Context) error) *UpdateChatTitle {
	return &UpdateChatTitle{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatTitle) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateChatTitle)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatTitle) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
