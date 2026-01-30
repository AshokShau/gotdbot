package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateFile struct {
	Filter   filters.UpdateFile
	Response func(ctx *ext.Context) error
}

func NewUpdateFile(filter filters.UpdateFile, response func(ctx *ext.Context) error) *UpdateFile {
	return &UpdateFile{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateFile) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateFile)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateFile) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
