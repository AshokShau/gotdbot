package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateNewChosenInlineResult struct {
	Filter   filters.UpdateNewChosenInlineResult
	Response func(ctx *ext.Context) error
}

func NewUpdateNewChosenInlineResult(filter filters.UpdateNewChosenInlineResult, response func(ctx *ext.Context) error) *UpdateNewChosenInlineResult {
	return &UpdateNewChosenInlineResult{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateNewChosenInlineResult) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateNewChosenInlineResult
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateNewChosenInlineResult) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
