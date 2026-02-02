package filters

import (
	"strings"

	"github.com/AshokShau/gotdbot"
)

var CallbackQueryAll CallbackQuery = func(cq *gotdbot.UpdateNewCallbackQuery) bool {
	return true
}

func CallbackQueryData(data string) CallbackQuery {
	return func(cq *gotdbot.UpdateNewCallbackQuery) bool {
		if cq.Payload == nil {
			return false
		}
		if p, ok := cq.Payload.(*gotdbot.CallbackQueryPayloadData); ok {
			return strings.Contains(string(p.Data), data)
		}
		return false
	}
}

var InlineQueryAll InlineQuery = func(iq *gotdbot.UpdateNewInlineQuery) bool {
	return true
}

func InlineQueryQuery(query string) InlineQuery {
	return func(iq *gotdbot.UpdateNewInlineQuery) bool {
		return iq.Query == query
	}
}

var ChatMemberAll ChatMember = func(cm *gotdbot.UpdateChatMember) bool {
	return true
}
