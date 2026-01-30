package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAgeVerificationParameters struct {
	Filter   filters.UpdateAgeVerificationParameters
	Response func(ctx *ext.Context) error
}

func NewUpdateAgeVerificationParameters(filter filters.UpdateAgeVerificationParameters, response func(ctx *ext.Context) error) *UpdateAgeVerificationParameters {
	return &UpdateAgeVerificationParameters{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAgeVerificationParameters) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateAgeVerificationParameters)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAgeVerificationParameters) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
