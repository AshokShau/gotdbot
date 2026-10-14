package gotdbot

import (
	"maps"
	"sort"
)

// Handler handles Telegram updates.
type Handler interface {
	CheckUpdate(client *Client, update TlObject) bool
	HandleUpdate(client *Client, update TlObject) error
}

// MatchingOptions controls which messages a handler accepts.
type MatchingOptions struct {
	AllowOutgoing   bool
	AllowIncoming   bool
	AllowEdited     bool
	AllowChannel    bool
	AllowGroup      bool // Basic groups only.
	AllowSupergroup bool // Supergroups, excluding channels.
	AllowPrivate    bool
	AllowSecret     bool
}

// DefaultMatchingOptions returns the default message matching options.
func DefaultMatchingOptions() MatchingOptions {
	return MatchingOptions{
		AllowIncoming:   true,
		AllowEdited:     true,
		AllowChannel:    true,
		AllowGroup:      true,
		AllowSupergroup: true,
		AllowPrivate:    true,
		AllowSecret:     true,
	}
}

// check reports whether msg satisfies the matching options.
func (o MatchingOptions) check(msg *Message) bool {
	if msg == nil {
		return false
	}

	// Incoming/outgoing.
	if msg.IsOutgoing {
		if !o.AllowOutgoing {
			return false
		}
	} else if !o.AllowIncoming {
		return false
	}

	// Edited messages.
	if msg.EditDate > 0 && !o.AllowEdited {
		return false
	}

	// Chat type.
	switch {
	case msg.IsChannel():
		return o.AllowChannel
	case msg.IsGroup():
		return o.AllowGroup
	case msg.IsSupergroup():
		return o.AllowSupergroup
	case msg.IsPrivate():
		return o.AllowPrivate
	case msg.IsSecret():
		return o.AllowSecret
	default:
		// Unknown chat type.
		return false
	}
}

// AddHandler adds a handler to the default group (0).
func (c *Client) AddHandler(handler Handler) {
	c.AddHandlerGroup(handler, 0)
}

// AddHandlerGroup adds a handler to the specified group.
//
// Handlers within a group are executed in registration order.
// Groups themselves are processed in ascending order.
func (c *Client) AddHandlerGroup(handler Handler, group int) {
	if handler == nil {
		return
	}

	c.hMu.Lock()
	defer c.hMu.Unlock()

	var oldHandlers map[int][]Handler

	if data := c.handlers.Load(); data != nil {
		oldHandlers = data.handlers
	}

	newHandlers := make(map[int][]Handler, len(oldHandlers))

	maps.Copy(newHandlers, oldHandlers)

	handlers := append(
		[]Handler(nil),
		newHandlers[group]...,
	)

	handlers = append(handlers, handler)
	newHandlers[group] = handlers

	groups := make([]int, 0, len(newHandlers))

	for groupID := range newHandlers {
		groups = append(groups, groupID)
	}

	sort.Ints(groups)

	c.handlers.Store(&handlersData{
		handlers: newHandlers,
		groups:   groups,
	})
}
