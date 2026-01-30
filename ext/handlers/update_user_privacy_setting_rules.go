package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUserPrivacySettingRules struct {
	Filter   filters.UpdateUserPrivacySettingRules
	Response func(ctx *ext.Context) error
}

func NewUpdateUserPrivacySettingRules(filter filters.UpdateUserPrivacySettingRules, response func(ctx *ext.Context) error) *UpdateUserPrivacySettingRules {
	return &UpdateUserPrivacySettingRules{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUserPrivacySettingRules) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateUserPrivacySettingRules)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUserPrivacySettingRules) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
