package gotdbot

import "github.com/AshokShau/gotdbot/filters"

// MessageHandler handles Message updates
type MessageHandler struct {
	Options  MatchingOptions
	Filter   MessageFilter
	Response func(client *Client, message *Message) error
}

// NewMessageHandler creates a new MessageHandler
func NewMessageHandler(response func(client *Client, message *Message) error, filter MessageFilter) *MessageHandler {
	return &MessageHandler{
		Options:  DefaultMatchingOptions(),
		Filter:   filter,
		Response: response,
	}
}

// CheckUpdate reports whether a message update matches this handler's options and filter.
func (h *MessageHandler) CheckUpdate(client *Client, update TlObject) bool {
	msg, ok := getMessageFromUpdate(update)
	if !ok || msg == nil || h.Response == nil {
		return false
	}
	if !h.Options.check(msg) {
		return false
	}
	if h.Filter != nil && !h.Filter(msg) {
		return false
	}
	return true
}

// HandleUpdate calls the response for a message carried by the update.
func (h *MessageHandler) HandleUpdate(client *Client, update TlObject) error {
	msg, ok := getMessageFromUpdate(update)
	if !ok || msg == nil || h.Response == nil {
		return nil
	}
	return h.Response(client, msg)
}

// SetOptions replaces the handler's message matching options.
// Configure handlers before registration or starting the client.
func (h *MessageHandler) SetOptions(options MatchingOptions) *MessageHandler {
	h.Options = options
	return h
}

// AddFilter combines a message filter with existing filters using AND.
// Configure handlers before registration or starting the client.
func (h *MessageHandler) AddFilter(filter MessageFilter) *MessageHandler {
	if filter != nil {
		h.Filter = h.Filter.And(filter)
	}
	return h
}

// OnMessage registers a handler in group 0 for new messages and business message edits.
func (c *Client) OnMessage(handler func(client *Client, message *Message) error, filter MessageFilter) *MessageHandler {
	return c.OnMessageGroup(handler, filter, 0)
}

// OnMessageGroup registers a message handler in the specified group.
func (c *Client) OnMessageGroup(handler func(client *Client, message *Message) error, filter MessageFilter, group int) *MessageHandler {
	h := NewMessageHandler(handler, filter)
	c.AddHandlerGroup(h, group)
	return h
}

// ForNewMessage adapts this message filter for a typed new-message handler.
func (filter MessageFilter) ForNewMessage() UpdateNewMessageFilter {
	return func(u *UpdateNewMessage) bool {
		if u == nil || u.Message == nil || filter == nil {
			return false
		}
		return filter(u.Message)
	}
}

// MessageFilter filters Message objects.
type MessageFilter func(msg *Message) bool

// And combines two message filters, ignoring nil filters.
func (f MessageFilter) And(other MessageFilter) MessageFilter {
	return filters.AllFilters[*Message](f, other)
}

// Or combines two message filters, ignoring nil filters.
func (f MessageFilter) Or(other MessageFilter) MessageFilter {
	return filters.AnyFilters[*Message](f, other)
}

// Not negates this message filter. A nil filter matches nothing.
func (f MessageFilter) Not() MessageFilter {
	return filters.NotFilter[*Message](f)
}
