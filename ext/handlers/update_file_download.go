package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFileDownload struct {
	Filter   filters.UpdateFileDownload
	Response func(ctx *ext.Context) error
}

func NewUpdateFileDownload(filter filters.UpdateFileDownload, response func(ctx *ext.Context) error) *UpdateFileDownload {
	return &UpdateFileDownload{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFileDownload) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateFileDownload)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFileDownload) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
