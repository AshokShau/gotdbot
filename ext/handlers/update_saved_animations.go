package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSavedAnimations struct {
	Filter   filters.UpdateSavedAnimations
	Response func(ctx *ext.Context) error
}

func NewUpdateSavedAnimations(filter filters.UpdateSavedAnimations, response func(ctx *ext.Context) error) *UpdateSavedAnimations {
	return &UpdateSavedAnimations{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSavedAnimations) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSavedAnimations)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSavedAnimations) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
