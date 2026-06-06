package gotdbot

import (
	"strings"
)

// Handler is the interface for all update handlers.
type Handler interface {
	// CheckUpdate checks if the update should be handled by this handler.
	CheckUpdate(client *Client, update TlObject) bool
	// HandleUpdate handles the update.
	HandleUpdate(client *Client, update TlObject) error
}

// MatchingOptions contains options for matching updates.
type MatchingOptions struct {
	AllowOutgoing bool
	AllowIncoming bool
	AllowEdited   bool
	AllowChannel  bool
	AllowGroup    bool
	AllowPrivate  bool
}

// DefaultMatchingOptions returns the default matching options.
func DefaultMatchingOptions() MatchingOptions {
	return MatchingOptions{
		AllowIncoming: true,
		AllowEdited:   true,
		AllowChannel:  true,
		AllowGroup:    true,
		AllowPrivate:  true,
	}
}

// checkMessageOptions checks if the message matches the matching options.
func (o MatchingOptions) check(msg *Message) bool {
	if msg == nil {
		return false
	}

	if msg.IsOutgoing && !o.AllowOutgoing {
		return false
	}
	if !msg.IsOutgoing && !o.AllowIncoming {
		return false
	}
	if msg.EditDate > 0 && !o.AllowEdited {
		return false
	}

	if msg.IsChannel() && !o.AllowChannel {
		return false
	}
	if msg.IsGroup() && !o.AllowGroup {
		return false
	}
	if msg.IsPrivate() && !o.AllowPrivate {
		return false
	}

	return true
}

// CommandHandler handles bot commands.
type CommandHandler struct {
	Options  MatchingOptions
	Triggers []rune
	Command  string
	Handler  func(client *Client, message *Message) error
	Filter   func(msg *Message) bool
}

// NewCommandHandler creates a new CommandHandler.
func NewCommandHandler(command string, handler func(client *Client, message *Message) error) *CommandHandler {
	return &CommandHandler{
		Options:  DefaultMatchingOptions(),
		Triggers: []rune{'/'},
		Command:  strings.ToLower(command),
		Handler:  handler,
	}
}

// CheckUpdate checks if the update is a command that matches this handler.
func (h *CommandHandler) CheckUpdate(client *Client, update TlObject) bool {
	msg, msgOk := getMessageFromUpdate(update)
	if !msgOk || !h.Options.check(msg) {
		return false
	}

	u, ok := update.(*UpdateNewMessage)
	if !ok || u.Message == nil || u.Message.Content == nil {
		return false
	}

	msg = u.Message
	if h.Filter != nil && !h.Filter(msg) {
		return false
	}

	var text string
	if textContent, ok := msg.Content.(*MessageText); ok {
		text = textContent.Text.Text
	} else if photoContent, ok := msg.Content.(*MessagePhoto); ok {
		text = photoContent.Caption.Text
	} else if videoContent, ok := msg.Content.(*MessageVideo); ok {
		text = videoContent.Caption.Text
	} else if documentContent, ok := msg.Content.(*MessageDocument); ok {
		text = documentContent.Caption.Text
	} else if audioContent, ok := msg.Content.(*MessageAudio); ok {
		text = audioContent.Caption.Text
	} else if animationContent, ok := msg.Content.(*MessageAnimation); ok {
		text = animationContent.Caption.Text
	} else if voiceContent, ok := msg.Content.(*MessageVoiceNote); ok {
		text = voiceContent.Caption.Text
	}

	if text == "" {
		return false
	}

	for _, trigger := range h.Triggers {
		prefix := string(trigger)
		if !strings.HasPrefix(text, prefix) {
			continue
		}

		parts := strings.Fields(text)
		if len(parts) == 0 {
			continue
		}

		cmdPart := parts[0][len(prefix):]

		var (
			cmdName      string
			mentionedBot string
		)

		if i := strings.Index(cmdPart, "@"); i != -1 {
			cmdName = cmdPart[:i]
			mentionedBot = cmdPart[i+1:]
		} else {
			cmdName = cmdPart
		}

		if strings.ToLower(cmdName) != h.Command {
			continue
		}

		if mentionedBot != "" {
			if client.Me == nil {
				return false
			}
			botUsername := ""
			if client.Me.Usernames != nil && len(client.Me.Usernames.ActiveUsernames) > 0 {
				botUsername = strings.ToLower(client.Me.Usernames.ActiveUsernames[0])
			}

			if botUsername == "" {
				return false
			}
			if strings.ToLower(mentionedBot) != botUsername {
				return false
			}
		}

		return true
	}

	return false
}

// HandleUpdate handles the command update.
func (h *CommandHandler) HandleUpdate(client *Client, update TlObject) error {
	u := update.(*UpdateNewMessage)
	return h.Handler(client, u.Message)
}

func getMessageFromUpdate(update TlObject) (*Message, bool) {
	switch u := update.(type) {
	case *UpdateNewMessage:
		return u.Message, true
	case *UpdateNewBusinessMessage:
		if u.Message != nil {
			return u.Message.Message, true
		}
	case *UpdateBusinessMessageEdited:
		if u.Message != nil {
			return u.Message.Message, true
		}
	}
	return nil, false
}

type funcHandler struct {
	check  func(client *Client, update TlObject) bool
	handle func(client *Client, update TlObject) error
}

func (h *funcHandler) CheckUpdate(client *Client, update TlObject) bool {
	return h.check(client, update)
}

func (h *funcHandler) HandleUpdate(client *Client, update TlObject) error {
	return h.handle(client, update)
}

// On registers a handler for a specific update type with the default group (0).
func (c *Client) On(updateType string, handler func(client *Client, update TlObject) error) {
	c.OnGroup(updateType, handler, 0)
}

// OnGroup registers a handler for a specific update type with a specific group.
func (c *Client) OnGroup(updateType string, handler func(client *Client, update TlObject) error, group int) {
	c.AddHandlerGroup(&funcHandler{
		check: func(client *Client, update TlObject) bool {
			return update.GetType() == updateType
		},
		handle: handler,
	}, group)
}

// AddHandler adds a handler with the default group (0).
func (c *Client) AddHandler(handler Handler) {
	c.AddHandlerGroup(handler, 0)
}

// AddHandlerGroup adds a handler with a specific group.
func (c *Client) AddHandlerGroup(handler Handler, group int) {
	c.hMu.Lock()
	defer c.hMu.Unlock()

	oldMap := *c.handlers.Load()
	newMap := make(map[int][]Handler, len(oldMap))
	for k, v := range oldMap {
		newMap[k] = v
	}

	handlers := make([]Handler, len(newMap[group]))
	copy(handlers, newMap[group])
	handlers = append(handlers, handler)

	newMap[group] = handlers
	c.handlers.Store(&newMap)
}

// OnCommand registers a new command handler with the default group (0).
func (c *Client) OnCommand(command string, handler func(client *Client, message *Message) error) *CommandHandler {
	return c.OnCommandGroup(command, handler, 0)
}

// OnCommandGroup registers a new command handler with a specific group.
func (c *Client) OnCommandGroup(command string, handler func(client *Client, message *Message) error, group int) *CommandHandler {
	ch := NewCommandHandler(command, handler)
	c.AddHandlerGroup(ch, group)
	return ch
}
