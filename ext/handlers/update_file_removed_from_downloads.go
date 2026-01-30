package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFileRemovedFromDownloads struct {
	Filter   filters.UpdateFileRemovedFromDownloads
	Response func(ctx *ext.Context) error
}

func NewUpdateFileRemovedFromDownloads(filter filters.UpdateFileRemovedFromDownloads, response func(ctx *ext.Context) error) *UpdateFileRemovedFromDownloads {
	return &UpdateFileRemovedFromDownloads{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFileRemovedFromDownloads) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateFileRemovedFromDownloads)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFileRemovedFromDownloads) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
