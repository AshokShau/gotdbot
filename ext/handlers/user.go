package handlers

import (
	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type User struct {
	Filter   filters.User
	Response func(ctx *ext.Context) error
}

func NewUser(filter filters.User, response func(ctx *ext.Context) error) *User {
	return &User{
		Filter:   filter,
		Response: response,
	}
}

func (h *User) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateUser)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *User) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}

type UserStatus struct {
	Filter   filters.UserStatus
	Response func(ctx *ext.Context) error
}

func NewUserStatus(filter filters.UserStatus, response func(ctx *ext.Context) error) *UserStatus {
	return &UserStatus{
		Filter:   filter,
		Response: response,
	}
}

func (h *UserStatus) CheckUpdate(ctx *ext.Context) bool {
	update, ok := ctx.RawUpdate.(*gotdbot.UpdateUserStatus)
	if !ok {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *UserStatus) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
