package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUserFullInfo struct {
	Filter   filters.UpdateUserFullInfo
	Response func(ctx *ext.Context) error
}

func NewUpdateUserFullInfo(filter filters.UpdateUserFullInfo, response func(ctx *ext.Context) error) *UpdateUserFullInfo {
	return &UpdateUserFullInfo{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUserFullInfo) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateUserFullInfo)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUserFullInfo) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
