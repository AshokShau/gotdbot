package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateSpeechRecognitionTrial struct {
	Filter   filters.UpdateSpeechRecognitionTrial
	Response func(ctx *ext.Context) error
}

func NewUpdateSpeechRecognitionTrial(filter filters.UpdateSpeechRecognitionTrial, response func(ctx *ext.Context) error) *UpdateSpeechRecognitionTrial {
	return &UpdateSpeechRecognitionTrial{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateSpeechRecognitionTrial) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateSpeechRecognitionTrial)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateSpeechRecognitionTrial) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
