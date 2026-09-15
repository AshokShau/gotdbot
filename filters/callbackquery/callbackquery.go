package callbackquery

import (
	"regexp"
	"strings"

	"github.com/AshokShau/gotdbot"
)

// Prefix checks if the callback query data starts with the given prefix.
func Prefix(prefix string) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return strings.HasPrefix(u.DataString(), prefix)
	}
}

// Suffix checks if the callback query data ends with the given suffix.
func Suffix(suffix string) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return strings.HasSuffix(u.DataString(), suffix)
	}
}

// Equal checks if the callback query data equals the given match string.
func Equal(match string) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.DataString() == match
	}
}

// Regex checks if the callback query data matches the given regex pattern.
func Regex(pattern string) gotdbot.UpdateNewCallbackQueryFilter {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return func(u *gotdbot.UpdateNewCallbackQuery) bool {
			return false
		}
	}
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		data := u.CallbackData()
		if data == nil {
			return false
		}
		return reg.Match(data)
	}
}

// FromUserID checks if the callback query sender's user ID matches the given ID.
func FromUserID(id int64) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.SenderUserId == id
	}
}

// ChatInstance checks if the callback query chat instance matches the given instance string.
func ChatInstance(instance int64) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.ChatInstance == instance
	}
}

// GameName checks if the callback query is for a game and matches the given short name.
func GameName(name string) gotdbot.UpdateNewCallbackQueryFilter {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.GameShortName() == name
	}
}

// Private checks if the callback query originated from a private chat.
func Private(u *gotdbot.UpdateNewCallbackQuery) bool {
	return u.IsPrivate()
}
