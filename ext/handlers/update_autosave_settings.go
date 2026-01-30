package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateAutosaveSettings struct {
	Filter   filters.UpdateAutosaveSettings
	Response func(ctx *ext.Context) error
}

func NewUpdateAutosaveSettings(filter filters.UpdateAutosaveSettings, response func(ctx *ext.Context) error) *UpdateAutosaveSettings {
	return &UpdateAutosaveSettings{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateAutosaveSettings) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateAutosaveSettings)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateAutosaveSettings) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
