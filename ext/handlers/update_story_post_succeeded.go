package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStoryPostSucceeded struct {
	Filter   filters.UpdateStoryPostSucceeded
	Response func(ctx *ext.Context) error
}

func NewUpdateStoryPostSucceeded(filter filters.UpdateStoryPostSucceeded, response func(ctx *ext.Context) error) *UpdateStoryPostSucceeded {
	return &UpdateStoryPostSucceeded{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStoryPostSucceeded) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateStoryPostSucceeded
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStoryPostSucceeded) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
