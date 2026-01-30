package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFileDownloads struct {
	Filter   filters.UpdateFileDownloads
	Response func(ctx *ext.Context) error
}

func NewUpdateFileDownloads(filter filters.UpdateFileDownloads, response func(ctx *ext.Context) error) *UpdateFileDownloads {
	return &UpdateFileDownloads{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFileDownloads) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateFileDownloads)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFileDownloads) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
