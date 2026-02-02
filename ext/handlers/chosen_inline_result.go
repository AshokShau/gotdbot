package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type ChosenInlineResult struct {
	Filter   filters.ChosenInlineResult
	Response func(ctx *ext.Context) error
}

func NewChosenInlineResult(filter filters.ChosenInlineResult, response func(ctx *ext.Context) error) *ChosenInlineResult {
	return &ChosenInlineResult{
		Filter:   filter,
		Response: response,
	}
}

func (h *ChosenInlineResult) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdateNewChosenInlineResult
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *ChosenInlineResult) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
