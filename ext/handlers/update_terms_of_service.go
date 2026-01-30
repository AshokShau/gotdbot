package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateTermsOfService struct {
	Filter   filters.UpdateTermsOfService
	Response func(ctx *ext.Context) error
}

func NewUpdateTermsOfService(filter filters.UpdateTermsOfService, response func(ctx *ext.Context) error) *UpdateTermsOfService {
	return &UpdateTermsOfService{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateTermsOfService) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateTermsOfService)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateTermsOfService) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
