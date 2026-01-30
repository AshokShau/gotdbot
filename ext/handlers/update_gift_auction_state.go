package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateGiftAuctionState struct {
	Filter   filters.UpdateGiftAuctionState
	Response func(ctx *ext.Context) error
}

func NewUpdateGiftAuctionState(filter filters.UpdateGiftAuctionState, response func(ctx *ext.Context) error) *UpdateGiftAuctionState {
	return &UpdateGiftAuctionState{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateGiftAuctionState) CheckUpdate(ctx *ext.Context) bool {
	u, ok := ctx.RawUpdate.(*gotdbot.UpdateGiftAuctionState)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateGiftAuctionState) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
