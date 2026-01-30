package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateApplicationRecaptchaVerificationRequired struct {
	Filter   filters.UpdateApplicationRecaptchaVerificationRequired
	Response func(ctx *ext.Context) error
}

func NewUpdateApplicationRecaptchaVerificationRequired(filter filters.UpdateApplicationRecaptchaVerificationRequired, response func(ctx *ext.Context) error) *UpdateApplicationRecaptchaVerificationRequired {
	return &UpdateApplicationRecaptchaVerificationRequired{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateApplicationRecaptchaVerificationRequired) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateApplicationRecaptchaVerificationRequired)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateApplicationRecaptchaVerificationRequired) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
