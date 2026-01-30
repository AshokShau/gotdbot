package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateVideoPublished struct {
	Filter   filters.UpdateVideoPublished
	Response func(ctx *ext.Context) error
}

func NewUpdateVideoPublished(filter filters.UpdateVideoPublished, response func(ctx *ext.Context) error) *UpdateVideoPublished {
	return &UpdateVideoPublished{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateVideoPublished) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateVideoPublished)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateVideoPublished) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
