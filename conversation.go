package gotdbot

import (
	"time"

)

// WaitMessageOpts holds optional parameters for Ask.
type WaitMessageOpts struct {
	Filter             func(*Message) bool
	CancellationFilter func(*Message) bool
	Timeout            time.Duration
}

// Ask waits for a new message in the specified chat.
func (c *Context) Ask(chatId int64, opts *WaitMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &WaitMessageOpts{Timeout: 1 * time.Minute}
	}

	filter := opts.Filter
	cancellationFilter := opts.CancellationFilter
	timeout := opts.Timeout

	msgFilter := func(ctx *Context) bool {
		if ctx.EffectiveChatId != chatId {
			return false
		}
		if ctx.Update.UpdateNewMessage == nil {
			return false
		}

		msg := ctx.Update.UpdateNewMessage.Message
		if cancellationFilter != nil && cancellationFilter(msg) {
			return true
		}

		if filter != nil && !filter(msg) {
			return false
		}
		return true
	}

	u, err := c.WaitFor(msgFilter, timeout)
	if err != nil {
		return nil, err
	}

	msg := u.(*UpdateNewMessage).Message
	if cancellationFilter != nil && cancellationFilter(msg) {
		return nil, ConversationCancelled
	}

	return msg, nil
}
