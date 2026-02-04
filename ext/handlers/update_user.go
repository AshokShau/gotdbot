package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type UpdateUser struct {
	Filter   filters.UpdateUser
	Response func(ctx *ext.Context) error
}

func NewUpdateUser(filter filters.UpdateUser, response func(ctx *ext.Context) error) *UpdateUser {
	return &UpdateUser{
		Filter:   filter,
		Response: response,
	}
}

func (h *UpdateUser) CheckUpdate(ctx *ext.Context) bool {
	u := ctx.Update.UpdateUser
	if u == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(u)
}

func (h *UpdateUser) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
