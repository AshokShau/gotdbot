package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type CallbackQuery struct {
	Filter   filters.CallbackQuery
	Response func(ctx *ext.Context) error
}

func NewCallbackQuery(filter filters.CallbackQuery, response func(ctx *ext.Context) error) *CallbackQuery {
	return &CallbackQuery{
		Filter:   filter,
		Response: response,
	}
}

func (c *CallbackQuery) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateNewCallbackQuery)
	if !ok {
		return false
	}
	if c.Filter == nil {
		return true
	}
	return c.Filter(update)
}

func (c *CallbackQuery) HandleUpdate(ctx *ext.Context) error {
	return c.Response(ctx)
}
