package callbackquery

import (
	"testing"

	"github.com/AshokShau/gotdbot"
)

func TestCallbackFilters(t *testing.T) {
	for _, payload := range []gotdbot.CallbackQueryPayload{&gotdbot.CallbackQueryPayloadData{Data: []byte("action:42")}, &gotdbot.CallbackQueryPayloadDataWithPassword{Data: []byte("action:42"), Password: "password"}} {
		u := &gotdbot.UpdateNewCallbackQuery{Payload: payload, ChatId: 42, SenderUserId: 7, ChatInstance: 9007199254740993}
		for _, tc := range []struct {
			name   string
			filter gotdbot.UpdateNewCallbackQueryFilter
			want   bool
		}{
			{"prefix", Prefix("action:"), true}, {"wrong prefix", Prefix("other:"), false}, {"suffix", Suffix(":42"), true},
			{"equal", Equal("action:42"), true}, {"case sensitive", Equal("ACTION:42"), false},
			{"regex", Regex(`^action:\d+$`), true}, {"bad regex", Regex("["), false},
			{"user", FromUserID(7), true}, {"wrong user", FromUserID(8), false},
			{"instance", ChatInstance(9007199254740993), true}, {"wrong instance", ChatInstance(1), false},
			{"private", Private, true}, {"game", GameName("game"), false},
		} {
			if got := tc.filter(u); got != tc.want {
				t.Errorf("%T/%s = %v, want %v", payload, tc.name, got, tc.want)
			}
		}
	}
	u := &gotdbot.UpdateNewCallbackQuery{Payload: &gotdbot.CallbackQueryPayloadGame{GameShortName: "game"}, ChatId: -42}
	if !GameName("game")(u) || Private(u) || Regex(`.*`)(u) {
		t.Fatal("game payload classified as callback data/private")
	}
	if Regex(`.*`)(&gotdbot.UpdateNewCallbackQuery{}) {
		t.Fatal("regex matched absent payload")
	}
}
