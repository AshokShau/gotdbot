package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSavedMessagesTags struct {
	Filter   filters.UpdateSavedMessagesTags
	Response func(ctx *ext.Context) error
}

func NewUpdateSavedMessagesTags(filter filters.UpdateSavedMessagesTags, response func(ctx *ext.Context) error) *UpdateSavedMessagesTags {
	return &UpdateSavedMessagesTags{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSavedMessagesTags) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSavedMessagesTags)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSavedMessagesTags) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
