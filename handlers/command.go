package handlers

import (
	"strings"

	"github.com/AshokShau/gotdbot"
)

type Command struct {
	Triggers []rune
	Command  string
	Response func(b *gotdbot.Client, ctx *gotdbot.Context) error
}

func NewCommand(command string, response func(b *gotdbot.Client, ctx *gotdbot.Context) error) *Command {
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

func (c *Command) CheckUpdate(b *gotdbot.Client, ctx *gotdbot.Context) bool {
	if ctx.EffectiveMessage == nil || ctx.EffectiveMessage.Content == nil {
		return false
	}

	if ctx.EffectiveMessage.IsOutgoing {
		return false
	}

	text := ctx.EffectiveMessage.Text()
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
			me := b.Me()
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

func (c *Command) HandleUpdate(b *gotdbot.Client, ctx *gotdbot.Context) error {
	return c.Response(b, ctx)
}
