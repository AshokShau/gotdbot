package handlers

import (
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/handlers/filters"
)

type Command struct {
	Triggers      []rune
	Command       string
	Response      func(b *gotdbot.Client, u *gotdbot.Message) error
	AllowOutgoing bool
	Filter        filters.Message
}

func NewCommand(command string, response func(b *gotdbot.Client, u *gotdbot.Message) error) *Command {
	return &Command{
		Triggers:      []rune{'/'},
		Command:       strings.ToLower(command),
		Response:      response,
		AllowOutgoing: false,
		Filter:        nil,
	}
}

func (c *Command) SetTriggers(triggers []rune) *Command {
	c.Triggers = triggers
	return c
}

func (c *Command) SetAllowOutgoing(allow bool) *Command {
	c.AllowOutgoing = allow
	return c
}

func (c *Command) SetFilter(filter filters.Message) *Command {
	c.Filter = filter
	return c
}

func (c *Command) CheckUpdate(b *gotdbot.Client, update gotdbot.TlObject) bool {
	msg := c.extractMessage(update)
	if msg == nil || msg.Content == nil {
		return false
	}

	if msg.IsOutgoing && !c.AllowOutgoing {
		return false
	}

	if c.Filter != nil && !c.Filter(msg) {
		return false
	}

	text := msg.Text()
	if text == "" {
		return false
	}

	for _, trigger := range c.Triggers {
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

		if strings.ToLower(cmdName) != c.Command {
			continue
		}

		if mentionedBot != "" {
			me, _ := b.GetMe()
			if me == nil {
				return false
			}
			botUsername := ""
			if me.Usernames != nil && len(me.Usernames.ActiveUsernames) > 0 {
				botUsername = strings.ToLower(me.Usernames.ActiveUsernames[0])
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

func (c *Command) HandleUpdate(b *gotdbot.Client, update gotdbot.TlObject) error {
	return c.Response(b, c.extractMessage(update))
}

// extractMessage extracts a gotdbot.Message from an update if possible.
func (c *Command) extractMessage(u gotdbot.TlObject) *gotdbot.Message {
	switch t := u.(type) {
	case *gotdbot.UpdateMessageSendSucceeded:
		return t.Message
	case *gotdbot.UpdateNewMessage:
		return t.Message
	default:
		return nil
	}
}
