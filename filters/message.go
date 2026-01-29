package filters

import (
	"strings"

	"github.com/AshokShau/gotdbot/client"
)

// MessageAll allows all messages.
func MessageAll(_ *client.Message) bool {
	return true
}

// MessageFromUserID returns a filter that checks if the message sender's user ID matches the given ID.
func MessageFromUserID(id int64) client.MessageFilter {
	return func(msg *client.Message) bool {
		return msg.FromID() == id
	}
}

// MessageText returns a filter that checks if the message text equals the given text.
func MessageText(text string) client.MessageFilter {
	return func(msg *client.Message) bool {
		return msg.Text() == text
	}
}

// MessageTextContains returns a filter that checks if the message text contains the given substring.
func MessageTextContains(text string) client.MessageFilter {
	return func(msg *client.Message) bool {
		return strings.Contains(msg.Text(), text)
	}
}

// Command returns a filter that checks if the message starts with the given command.
func Command(command string, prefixes string) client.MessageFilter {
	if prefixes == "" {
		prefixes = "/"
	}
	return func(msg *client.Message) bool {
		text := msg.Text()
		if text == "" {
			return false
		}
		for _, prefix := range prefixes {
			if strings.HasPrefix(text, string(prefix)+command) {
				return true
			}
		}
		return false
	}
}
