package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStory struct {
	Filter   filters.UpdateStory
	Response func(ctx *ext.Context) error
}

func NewUpdateStory(filter filters.UpdateStory, response func(ctx *ext.Context) error) *UpdateStory {
	return &UpdateStory{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStory) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStory)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStory) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
