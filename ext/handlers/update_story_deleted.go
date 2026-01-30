package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateStoryDeleted struct {
	Filter   filters.UpdateStoryDeleted
	Response func(ctx *ext.Context) error
}

func NewUpdateStoryDeleted(filter filters.UpdateStoryDeleted, response func(ctx *ext.Context) error) *UpdateStoryDeleted {
	return &UpdateStoryDeleted{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateStoryDeleted) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateStoryDeleted)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateStoryDeleted) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
