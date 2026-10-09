package gotdbot

import (
	"sort"
	"strings"
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

// CommandHandler handles bot commands such as /start and /help.
type CommandHandler struct {
	Options  MatchingOptions
	Triggers []rune
	Command  string
	Handler  func(client *Client, message *Message) error
	Filter   MessageFilter
}

// NewCommandHandler creates a new command handler.
//
// The command should be provided without the trigger.
// For example: "start" instead of "/start".
func NewCommandHandler(command string, handler func(client *Client, message *Message) error) *CommandHandler {
	return &CommandHandler{
		Options:  DefaultMatchingOptions(),
		Triggers: []rune{'/'},
		Command:  strings.ToLower(strings.TrimSpace(command)),
		Handler:  handler,
	}
}

// AddFilter adds a message filter.
//
// All filters must return true for the command to execute.
func (h *CommandHandler) AddFilter(filter MessageFilter) *CommandHandler {
	if filter == nil {
		return h
	}

	if h.Filter == nil {
		h.Filter = filter
		return h
	}

	previous := h.Filter
	h.Filter = func(msg *Message) bool {
		return previous(msg) && filter(msg)
	}

	return h
}

// CheckUpdate reports whether update matches this command handler.
func (h *CommandHandler) CheckUpdate(client *Client, update TlObject) bool {
	msg, ok := getMessageFromUpdate(update)
	if !ok || !h.Options.check(msg) {
		return false
	}

	if msg.Content == nil {
		return false
	}

	if h.Filter != nil && !h.Filter(msg) {
		return false
	}

	text := strings.TrimSpace(msg.GetText())
	if text == "" {
		return false
	}

	return h.matchCommand(client, text)
}

// matchCommand checks whether text starts with one of the configured
// command triggers and matches this handler's command.
func (h *CommandHandler) matchCommand(client *Client, text string) bool {
	parts := strings.Fields(text)
	if len(parts) == 0 {
		return false
	}

	commandPart := parts[0]

	for _, trigger := range h.Triggers {
		prefix := string(trigger)

		if !strings.HasPrefix(commandPart, prefix) {
			continue
		}

		commandPart = commandPart[len(prefix):]
		if commandPart == "" {
			continue
		}

		commandName, mentionedBot := splitCommand(commandPart)
		if !strings.EqualFold(commandName, h.Command) {
			continue
		}

		if mentionedBot != "" && !h.matchesBot(client, mentionedBot) {
			return false
		}

		return true
	}

	return false
}

// splitCommand separates "start@mybot" into "start" and "mybot".
func splitCommand(command string) (name, username string) {
	name, username, ok := strings.Cut(command, "@")
	if !ok {
		return command, ""
	}

	return name, username
}

// matchesBot reports whether username refers to the current bot.
func (h *CommandHandler) matchesBot(client *Client, username string) bool {
	if client == nil {
		return false
	}
	client.meMu.RLock()
	defer client.meMu.RUnlock()
	if client.Me == nil {
		return false
	}

	if client.Me.Usernames == nil {
		return false
	}

	for _, botUsername := range client.Me.Usernames.ActiveUsernames {
		if strings.EqualFold(botUsername, username) {
			return true
		}
	}

	return false
}

// HandleUpdate handles a matched command update.
func (h *CommandHandler) HandleUpdate(client *Client, update TlObject) error {
	msg, ok := getMessageFromUpdate(update)
	if !ok || msg == nil || h.Handler == nil {
		return nil
	}

	return h.Handler(client, msg)
}

// getMessageFromUpdate extracts a Message from supported update types.
func getMessageFromUpdate(update TlObject) (*Message, bool) {
	switch u := update.(type) {
	case *UpdateNewMessage:
		if u != nil && u.Message != nil {
			return u.Message, true
		}

	case *UpdateNewBusinessMessage:
		if u != nil && u.Message != nil && u.Message.Message != nil {
			return u.Message.Message, true
		}

	case *UpdateBusinessMessageEdited:
		if u != nil && u.Message != nil && u.Message.Message != nil {
			return u.Message.Message, true
		}
	}

	return nil, false
}

// getChatIDFromUpdate returns the chat ID associated with an update.
func getChatIDFromUpdate(update TlObject) int64 {
	if msg, ok := getMessageFromUpdate(update); ok && msg != nil {
		return msg.ChatId
	}

	switch u := update.(type) {
	case *UpdateNewCallbackQuery:
		return u.ChatId

	case *UpdateMessageEdited:
		return u.ChatId

	case *UpdateMessageContent:
		return u.ChatId

	case *UpdateMessageEphemeralContent:
		return u.ChatId

	case *UpdateMessageLiveLocationViewed:
		return u.ChatId

	case *UpdateMessageInteractionInfo:
		return u.ChatId

	case *UpdateMessageSendSucceeded:
		if u != nil && u.Message != nil {
			return u.Message.ChatId
		}

	case *UpdateMessageSendFailed:
		if u != nil && u.Message != nil {
			return u.Message.ChatId
		}
	}

	return 0
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

	for groupID, handlers := range oldHandlers {
		newHandlers[groupID] = handlers
	}

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

// OnCommand registers a command handler in the default group (0).
func (c *Client) OnCommand(command string, handler func(client *Client, message *Message) error) *CommandHandler {
	return c.OnCommandGroup(command, handler, 0)
}

// OnCommandGroup registers a command handler in the specified group.
func (c *Client) OnCommandGroup(command string, handler func(client *Client, message *Message) error, group int) *CommandHandler {
	commandHandler := NewCommandHandler(command, handler)
	c.AddHandlerGroup(commandHandler, group)

	return commandHandler
}
