package gotdbot

import (
	"strings"
	"unicode/utf8"
)

// CommandHandler handles bot commands such as /start and /help.
type CommandHandler struct {
	Options  MatchingOptions
	Triggers []rune
	Command  string
	Aliases  []string
	Handler  func(client *Client, message *Message) error
	Filter   MessageFilter
}

// SetOptions changes the handler's matching options. Configure handlers before
// registration or starting the client.
func (h *CommandHandler) SetOptions(options MatchingOptions) *CommandHandler {
	h.Options = options
	return h
}

// SetTriggers sets the characters that can prefix this command, such as "/!.".
// Each character is a separate trigger; an empty string disables matching.
// Configure triggers before registration or starting the client.
func (h *CommandHandler) SetTriggers(triggers string) *CommandHandler {
	h.Triggers = []rune(triggers)
	return h
}

// AddAliases adds alternative command names that call the same handler.
// Provide names without a trigger, such as "help" instead of "/help".
// Names are trimmed and matched case-insensitively; empty names are ignored.
// Aliases share the handler's triggers, matching options, and filters.
// Configure aliases before registration or starting the client.
func (h *CommandHandler) AddAliases(aliases ...string) *CommandHandler {
	for _, alias := range aliases {
		alias = strings.ToLower(strings.TrimSpace(alias))
		if alias != "" {
			h.Aliases = append(h.Aliases, alias)
		}
	}
	return h
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
// command triggers and matches this handler's command or an alias.
func (h *CommandHandler) matchCommand(client *Client, text string) bool {
	parts := strings.Fields(text)
	if len(parts) == 0 {
		return false
	}

	commandPart := parts[0]
	first, size := utf8.DecodeRuneInString(commandPart)

	for _, trigger := range h.Triggers {
		if first != trigger {
			continue
		}

		commandName, mentionedBot, hasMention := strings.Cut(commandPart[size:], "@")
		if commandName == "" {
			return false
		}
		if hasMention && (mentionedBot == "" || strings.ContainsRune(mentionedBot, '@')) {
			return false
		}
		matched := strings.EqualFold(commandName, h.Command)
		if !matched {
			for _, alias := range h.Aliases {
				if strings.EqualFold(commandName, alias) {
					matched = true
					break
				}
			}
		}
		if !matched {
			continue
		}

		if mentionedBot != "" && !h.matchesBot(client, mentionedBot) {
			return false
		}

		return true
	}

	return false
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
