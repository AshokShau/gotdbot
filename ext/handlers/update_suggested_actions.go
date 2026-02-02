package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSuggestedActions struct {
	Filter   filters.UpdateSuggestedActions
	Response func(ctx *ext.Context) error
}

func NewUpdateSuggestedActions(filter filters.UpdateSuggestedActions, response func(ctx *ext.Context) error) *UpdateSuggestedActions {
	return &UpdateSuggestedActions{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSuggestedActions) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateSuggestedActions
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSuggestedActions) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
