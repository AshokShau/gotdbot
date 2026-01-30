package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUnconfirmedSession struct {
	Filter   filters.UpdateUnconfirmedSession
	Response func(ctx *ext.Context) error
}

func NewUpdateUnconfirmedSession(filter filters.UpdateUnconfirmedSession, response func(ctx *ext.Context) error) *UpdateUnconfirmedSession {
	return &UpdateUnconfirmedSession{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUnconfirmedSession) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateUnconfirmedSession)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUnconfirmedSession) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
