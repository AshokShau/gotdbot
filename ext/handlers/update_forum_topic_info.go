package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateForumTopicInfo struct {
	Filter   filters.UpdateForumTopicInfo
	Response func(ctx *ext.Context) error
}

func NewUpdateForumTopicInfo(filter filters.UpdateForumTopicInfo, response func(ctx *ext.Context) error) *UpdateForumTopicInfo {
	return &UpdateForumTopicInfo{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateForumTopicInfo) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateForumTopicInfo)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateForumTopicInfo) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
