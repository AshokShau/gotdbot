package callbackquery

import (
	"regexp"
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/ext/handlers/filters"
)

// Prefix checks if the callback query data starts with the given prefix.
func Prefix(prefix string) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return strings.HasPrefix(u.DataString(), prefix)
	}
}

// Suffix checks if the callback query data ends with the given suffix.
func Suffix(suffix string) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return strings.HasSuffix(u.DataString(), suffix)
	}
}

// Equal checks if the callback query data equals the given match string.
func Equal(match string) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.DataString() == match
	}
}

// Regex checks if the callback query data matches the given regex pattern.
// It panics if the regex pattern is invalid.
func Regex(pattern string) filters.CallbackQuery {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		panic("gotdbot: invalid regex pattern in filter: " + err.Error())
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
func FromUserID(id int64) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.SenderUserId == id
	}
}

// ChatInstance checks if the callback query chat instance matches the given instance string.
func ChatInstance(instance string) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.ChatInstance == instance
	}
}

// GameName checks if the callback query is for a game and matches the given short name.
func GameName(name string) filters.CallbackQuery {
	return func(u *gotdbot.UpdateNewCallbackQuery) bool {
		return u.GameShortName() == name
	}
}
