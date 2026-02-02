package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateDefaultReactionType struct {
	Filter   filters.UpdateDefaultReactionType
	Response func(ctx *ext.Context) error
}

func NewUpdateDefaultReactionType(filter filters.UpdateDefaultReactionType, response func(ctx *ext.Context) error) *UpdateDefaultReactionType {
	return &UpdateDefaultReactionType{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateDefaultReactionType) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateDefaultReactionType
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateDefaultReactionType) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
