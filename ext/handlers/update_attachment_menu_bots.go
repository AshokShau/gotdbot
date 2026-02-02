package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAttachmentMenuBots struct {
	Filter   filters.UpdateAttachmentMenuBots
	Response func(ctx *ext.Context) error
}

func NewUpdateAttachmentMenuBots(filter filters.UpdateAttachmentMenuBots, response func(ctx *ext.Context) error) *UpdateAttachmentMenuBots {
	return &UpdateAttachmentMenuBots{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAttachmentMenuBots) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateAttachmentMenuBots
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAttachmentMenuBots) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
