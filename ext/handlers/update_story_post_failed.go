package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStoryPostFailed struct {
	Filter   filters.UpdateStoryPostFailed
	Response func(ctx *ext.Context) error
}

func NewUpdateStoryPostFailed(filter filters.UpdateStoryPostFailed, response func(ctx *ext.Context) error) *UpdateStoryPostFailed {
	return &UpdateStoryPostFailed{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStoryPostFailed) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStoryPostFailed)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStoryPostFailed) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
