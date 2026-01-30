package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGroupCallMessageLevels struct {
	Filter   filters.UpdateGroupCallMessageLevels
	Response func(ctx *ext.Context) error
}

func NewUpdateGroupCallMessageLevels(filter filters.UpdateGroupCallMessageLevels, response func(ctx *ext.Context) error) *UpdateGroupCallMessageLevels {
	return &UpdateGroupCallMessageLevels{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGroupCallMessageLevels) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateGroupCallMessageLevels)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGroupCallMessageLevels) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
