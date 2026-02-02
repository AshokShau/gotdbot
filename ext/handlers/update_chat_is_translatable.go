package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateChatIsTranslatable struct {
	Filter   filters.UpdateChatIsTranslatable
	Response func(ctx *ext.Context) error
}

func NewUpdateChatIsTranslatable(filter filters.UpdateChatIsTranslatable, response func(ctx *ext.Context) error) *UpdateChatIsTranslatable {
	return &UpdateChatIsTranslatable{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateChatIsTranslatable) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateChatIsTranslatable
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateChatIsTranslatable) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
