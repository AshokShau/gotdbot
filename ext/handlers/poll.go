package handlers

import (
	"github.com/AshokShau/gotdbot/ext"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

type Poll struct {
	Filter   filters.Poll
	Response func(ctx *ext.Context) error
}

func NewPoll(filter filters.Poll, response func(ctx *ext.Context) error) *Poll {
	return &Poll{
		Filter:   filter,
		Response: response,
	}
}

func (h *Poll) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdatePoll
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *Poll) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}

type PollAnswer struct {
	Filter   filters.PollAnswer
	Response func(ctx *ext.Context) error
}

func NewPollAnswer(filter filters.PollAnswer, response func(ctx *ext.Context) error) *PollAnswer {
	return &PollAnswer{
		Filter:   filter,
		Response: response,
	}
}

func (h *PollAnswer) CheckUpdate(ctx *ext.Context) bool {
	update := ctx.Update.UpdatePollAnswer
	if update == nil {
		return false
	}
	if h.Filter == nil {
		return true
	}
	return h.Filter(update)
}

func (h *PollAnswer) HandleUpdate(ctx *ext.Context) error {
	return h.Response(ctx)
}
