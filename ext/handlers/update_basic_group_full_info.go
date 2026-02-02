package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateBasicGroupFullInfo struct {
	Filter   filters.UpdateBasicGroupFullInfo
	Response func(ctx *ext.Context) error
}

func NewUpdateBasicGroupFullInfo(filter filters.UpdateBasicGroupFullInfo, response func(ctx *ext.Context) error) *UpdateBasicGroupFullInfo {
	return &UpdateBasicGroupFullInfo{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateBasicGroupFullInfo) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateBasicGroupFullInfo
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateBasicGroupFullInfo) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
