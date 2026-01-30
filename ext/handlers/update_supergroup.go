package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSupergroup struct {
	Filter   filters.UpdateSupergroup
	Response func(ctx *ext.Context) error
}

func NewUpdateSupergroup(filter filters.UpdateSupergroup, response func(ctx *ext.Context) error) *UpdateSupergroup {
	return &UpdateSupergroup{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSupergroup) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSupergroup)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSupergroup) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
