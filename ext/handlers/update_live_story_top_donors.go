package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateLiveStoryTopDonors struct {
	Filter   filters.UpdateLiveStoryTopDonors
	Response func(ctx *ext.Context) error
}

func NewUpdateLiveStoryTopDonors(filter filters.UpdateLiveStoryTopDonors, response func(ctx *ext.Context) error) *UpdateLiveStoryTopDonors {
	return &UpdateLiveStoryTopDonors{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateLiveStoryTopDonors) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateLiveStoryTopDonors)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateLiveStoryTopDonors) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
