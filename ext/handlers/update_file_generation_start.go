package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFileGenerationStart struct {
	Filter   filters.UpdateFileGenerationStart
	Response func(ctx *ext.Context) error
}

func NewUpdateFileGenerationStart(filter filters.UpdateFileGenerationStart, response func(ctx *ext.Context) error) *UpdateFileGenerationStart {
	return &UpdateFileGenerationStart{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFileGenerationStart) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateFileGenerationStart
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFileGenerationStart) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
