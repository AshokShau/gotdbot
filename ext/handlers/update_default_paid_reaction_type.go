package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateDefaultPaidReactionType struct {
	Filter   filters.UpdateDefaultPaidReactionType
	Response func(ctx *ext.Context) error
}

func NewUpdateDefaultPaidReactionType(filter filters.UpdateDefaultPaidReactionType, response func(ctx *ext.Context) error) *UpdateDefaultPaidReactionType {
	return &UpdateDefaultPaidReactionType{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateDefaultPaidReactionType) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateDefaultPaidReactionType)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateDefaultPaidReactionType) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
