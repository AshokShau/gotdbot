package handlers

import (
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext"
)

type Command struct {
	Triggers []rune
	Command  string
	Response func(ctx *ext.Context) error
}

func NewCommand(command string, response func(ctx *ext.Context) error) *Command {
	return &Command{
		Triggers: []rune{'/'},
		Command:  strings.ToLower(command),
		Response: response,
	}
}

func (c *Command) SetTriggers(triggers []rune) *Command {
	c.Triggers = triggers
	return c
}

func (c *Command) CheckUpdate(ctx *ext.Context) bool {
	if ctx.EffectiveMessage == nil || ctx.EffectiveMessage.Content == nil {
		return false
	}
	
	msgText, ok := ctx.EffectiveMessage.Content.(*gotdbot.MessageText)
	if !ok || msgText.Text == nil {
		return false
	}

	// Ignore outgoing messages
	if ctx.EffectiveMessage.IsOutgoing {
		return false
	}

	text := msgText.Text.Text
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
		// Handle bot username mentions: command@username
		if idx := strings.Index(cmdPart, "@"); idx != -1 {
			cmdPart = cmdPart[:idx]
		}

		if strings.ToLower(cmdPart) == c.Command {
			return true
		}
	}
	return false
}

func (c *Command) HandleUpdate(ctx *ext.Context) error {
	return c.Response(ctx)
}
