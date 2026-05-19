package gotdbot

import (
	"time"
)

type Context struct {
	// Client is the client that received the update.
	Client *Client
	// RawUpdate is the original update object.
	RawUpdate TlObject
	// Update contains pointers to all possible update types.
	Update *ContextUpdates
	// Dispatcher is the dispatcher that created this context.
	Dispatcher *Dispatcher

	// Data is a map for storing data across handlers.
	Data map[string]interface{}

	// Helper fields derived from the update
	EffectiveMessage *Message
	EffectiveChatId  int64
}

func NewContext(client *Client, update TlObject, dispatcher *Dispatcher) *Context {
	ctx := &Context{
		Client:     client,
		RawUpdate:  update,
		Update:     NewContextUpdates(update),
		Dispatcher: dispatcher,
		Data:       make(map[string]interface{}),
	}
	extractGeneratedEffectiveFields(update, ctx)
	return ctx
}

// WaitFor waits for an update that matches the filter.
func (c *Context) WaitFor(filter UpdateFilter, timeout time.Duration) (TlObject, error) {
	return c.Dispatcher.WaitFor(c.Client, filter, timeout)
}
