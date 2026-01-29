package ext

import "github.com/AshokShau/gotdbot"

// MessageFilter is a function that filters messages.
type MessageFilter func(*Message) bool

// CallbackQueryFilter is a function that filters callback queries.
type CallbackQueryFilter func(*CallbackQuery) bool

// OnMessage registers a handler for new messages that provides a bound Message object.
func OnMessage(c *gotdbot.Client, fn func(client *gotdbot.Client, message *Message), filters ...MessageFilter) {
	c.AddHandler("updateNewMessage", func(cl *gotdbot.Client, update gotdbot.TlObject) error {
		up, ok := update.(*gotdbot.UpdateNewMessage)
		if !ok {
			return nil
		}

		boundMsg := NewMessage(cl, up.Message)
		for _, f := range filters {
			if f != nil && !f(boundMsg) {
				return nil
			}
		}

		fn(cl, boundMsg)
		return nil
	}, nil, 0)
}

// OnCallbackQuery registers a handler for new callback queries that provides a bound CallbackQuery object.
func OnCallbackQuery(c *gotdbot.Client, fn func(client *gotdbot.Client, callbackQuery *CallbackQuery), filters ...CallbackQueryFilter) {
	c.AddHandler("updateNewCallbackQuery", func(cl *gotdbot.Client, update gotdbot.TlObject) error {
		up, ok := update.(*gotdbot.UpdateNewCallbackQuery)
		if !ok {
			return nil
		}

		boundCb := NewCallbackQuery(cl, up)
		for _, f := range filters {
			if f != nil && !f(boundCb) {
				return nil
			}
		}

		fn(cl, boundCb)
		return nil
	}, nil, 0)
}
