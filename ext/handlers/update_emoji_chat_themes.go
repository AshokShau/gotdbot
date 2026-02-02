package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateEmojiChatThemes struct {
	Filter   filters.UpdateEmojiChatThemes
	Response func(ctx *ext.Context) error
}

func NewUpdateEmojiChatThemes(filter filters.UpdateEmojiChatThemes, response func(ctx *ext.Context) error) *UpdateEmojiChatThemes {
	return &UpdateEmojiChatThemes{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateEmojiChatThemes) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateEmojiChatThemes
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateEmojiChatThemes) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
