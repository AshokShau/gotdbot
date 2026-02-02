package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatIsMarkedAsUnread struct {
	Filter   filters.UpdateChatIsMarkedAsUnread
	Response func(ctx *ext.Context) error
}

func NewUpdateChatIsMarkedAsUnread(filter filters.UpdateChatIsMarkedAsUnread, response func(ctx *ext.Context) error) *UpdateChatIsMarkedAsUnread {
	return &UpdateChatIsMarkedAsUnread{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatIsMarkedAsUnread) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatIsMarkedAsUnread
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatIsMarkedAsUnread) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
