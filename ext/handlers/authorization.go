package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type AuthorizationState struct {
	Filter   filters.AuthorizationState
	Response func(ctx *ext.Context) error
}

func NewAuthorizationState(filter filters.AuthorizationState, response func(ctx *ext.Context) error) *AuthorizationState {
	return &AuthorizationState{
		Filter:   filter,
		Response: response,
	}
}

func (h *AuthorizationState) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateAuthorizationState)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *AuthorizationState) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
