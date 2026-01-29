package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
)

type EditedMessage struct {
	Filter   func(u *gotdbot.UpdateMessageEdited) bool
	Response func(ctx *ext.Context) error
}

func NewEditedMessage(filter func(u *gotdbot.UpdateMessageEdited) bool, response func(ctx *ext.Context) error) *EditedMessage {
	return &EditedMessage{
		Filter:   filter,
		Response: response,
	}
}

func (h *EditedMessage) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateMessageEdited)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *EditedMessage) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
