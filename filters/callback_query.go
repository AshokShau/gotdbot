package filters

import (
	"strings"

	"github.com/AshokShau/gotdbot/ext"
)

// CallbackQueryData returns a filter that checks if the callback data equals the given string.
func CallbackQueryData(data string) ext.CallbackQueryFilter {
	return func(cb *ext.CallbackQuery) bool {
		return cb.Text() == data
	}
}

// CallbackQueryDataPrefix returns a filter that checks if the callback data starts with the given prefix.
func CallbackQueryDataPrefix(prefix string) ext.CallbackQueryFilter {
	return func(cb *ext.CallbackQuery) bool {
		return strings.HasPrefix(cb.Text(), prefix)
	}
}
