package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateLanguagePackStrings struct {
	Filter   filters.UpdateLanguagePackStrings
	Response func(ctx *ext.Context) error
}

func NewUpdateLanguagePackStrings(filter filters.UpdateLanguagePackStrings, response func(ctx *ext.Context) error) *UpdateLanguagePackStrings {
	return &UpdateLanguagePackStrings{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateLanguagePackStrings) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateLanguagePackStrings)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateLanguagePackStrings) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
