package message

import (
	"regexp"

	"github.com/AshokShau/gotdbot"
)

// Prefix checks if the message text starts with the given prefix.
func Prefix(prefix string) gotdbot.MessageFilter {
	return func(msg *gotdbot.Message) bool {
		return msg.StartsWith(prefix)
	}
}

// Suffix checks if the message text ends with the given suffix.
func Suffix(suffix string) gotdbot.MessageFilter {
	return func(msg *gotdbot.Message) bool {
		return msg.EndsWith(suffix)
	}
}

// Equal checks if the message text equals the given match string.
func Equal(match string) gotdbot.MessageFilter {
	return func(msg *gotdbot.Message) bool {
		return msg.GetText() == match
	}
}

// Regex checks if the message text matches the given regex pattern.
// Invalid patterns match nothing; use CompileRegex to receive the error.
func Regex(pattern string) gotdbot.MessageFilter {
	filter, err := CompileRegex(pattern)
	if err != nil {
		return func(msg *gotdbot.Message) bool {
			return false
		}
	}
	return filter
}

// CompileRegex creates a message filter or returns a pattern compilation error.
func CompileRegex(pattern string) (gotdbot.MessageFilter, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return func(msg *gotdbot.Message) bool {
		return msg != nil && reg.MatchString(msg.GetText())
	}, nil
}

// Private filters messages sent in private chats.
func Private(msg *gotdbot.Message) bool {
	return msg.IsPrivate()
}

// Group filters messages sent in basic groups, excluding supergroups.
func Group(msg *gotdbot.Message) bool {
	return msg.IsGroup()
}

// Supergroup filters messages sent in supergroups, excluding channels.
func Supergroup(msg *gotdbot.Message) bool {
	return msg.IsSupergroup()
}

// Secret filters messages sent in secret chats.
func Secret(msg *gotdbot.Message) bool {
	return msg.IsSecret()
}

// Channel filters messages sent in channel chats.
func Channel(msg *gotdbot.Message) bool {
	return msg.IsChannel()
}

// Incoming filters incoming messages.
func Incoming(msg *gotdbot.Message) bool {
	return !msg.IsOutgoingMessage()
}

// Outgoing filters outgoing messages.
func Outgoing(msg *gotdbot.Message) bool {
	return msg.IsOutgoingMessage()
}

// Edited filters edited messages.
func Edited(msg *gotdbot.Message) bool {
	return msg.IsEditedMessage()
}

// Command filters messages that are commands.
func Command(msg *gotdbot.Message) bool {
	return msg.IsCommand()
}

// FromUserID checks if the message sender's user ID matches the given ID.
func FromUserID(id int64) gotdbot.MessageFilter {
	return func(msg *gotdbot.Message) bool {
		return msg.SenderID() == id
	}
}

// ChatID checks if the message chat ID matches the given ID.
func ChatID(id int64) gotdbot.MessageFilter {
	return func(msg *gotdbot.Message) bool {
		return msg.ChatID() == id
	}
}
