package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatPhoto struct {
	Filter   filters.UpdateChatPhoto
	Response func(ctx *ext.Context) error
}

func NewUpdateChatPhoto(filter filters.UpdateChatPhoto, response func(ctx *ext.Context) error) *UpdateChatPhoto {
	return &UpdateChatPhoto{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatPhoto) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatPhoto
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatPhoto) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
