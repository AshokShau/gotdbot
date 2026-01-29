package filters

import (
	"strings"

	"github.com/AshokShau/gotdbot/ext"
)

// MessageAll allows all messages.
func MessageAll(_ *ext.Message) bool {
	return true
}

// MessageFromUserID returns a filter that checks if the message sender's user ID matches the given ID.
func MessageFromUserID(id int64) ext.MessageFilter {
	return func(msg *ext.Message) bool {
		return msg.FromID() == id
	}
}

// MessageText returns a filter that checks if the message text equals the given text.
func MessageText(text string) ext.MessageFilter {
	return func(msg *ext.Message) bool {
		return msg.Text() == text
	}
}

// MessageTextContains returns a filter that checks if the message text contains the given substring.
func MessageTextContains(text string) ext.MessageFilter {
	return func(msg *ext.Message) bool {
		return strings.Contains(msg.Text(), text)
	}
}

// Command returns a filter that checks if the message starts with the given command.
func Command(command string, prefixes string) ext.MessageFilter {
	if prefixes == "" {
		prefixes = "/"
	}
	return func(msg *ext.Message) bool {
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
