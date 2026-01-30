package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAnimationSearchParameters struct {
	Filter   filters.UpdateAnimationSearchParameters
	Response func(ctx *ext.Context) error
}

func NewUpdateAnimationSearchParameters(filter filters.UpdateAnimationSearchParameters, response func(ctx *ext.Context) error) *UpdateAnimationSearchParameters {
	return &UpdateAnimationSearchParameters{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAnimationSearchParameters) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateAnimationSearchParameters)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAnimationSearchParameters) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
