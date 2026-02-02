package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFileAddedToDownloads struct {
	Filter   filters.UpdateFileAddedToDownloads
	Response func(ctx *ext.Context) error
}

func NewUpdateFileAddedToDownloads(filter filters.UpdateFileAddedToDownloads, response func(ctx *ext.Context) error) *UpdateFileAddedToDownloads {
	return &UpdateFileAddedToDownloads{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFileAddedToDownloads) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateFileAddedToDownloads
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFileAddedToDownloads) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
