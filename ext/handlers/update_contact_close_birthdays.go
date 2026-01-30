package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateContactCloseBirthdays struct {
	Filter   filters.UpdateContactCloseBirthdays
	Response func(ctx *ext.Context) error
}

func NewUpdateContactCloseBirthdays(filter filters.UpdateContactCloseBirthdays, response func(ctx *ext.Context) error) *UpdateContactCloseBirthdays {
	return &UpdateContactCloseBirthdays{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateContactCloseBirthdays) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateContactCloseBirthdays)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateContactCloseBirthdays) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
