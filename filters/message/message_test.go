package message

import (
	"sync"
	"testing"

	"github.com/AshokShau/gotdbot"
)

func TestMessageFilters(t *testing.T) {
	m := &gotdbot.Message{ChatId: 42, SenderId: &gotdbot.MessageSenderUser{UserId: 7}, EditDate: 1, Content: &gotdbot.MessageText{Text: &gotdbot.FormattedText{Text: "/start hello😀", Entities: []gotdbot.TextEntity{{Offset: 0, Length: 6, Type: &gotdbot.TextEntityTypeBotCommand{}}}}}}
	for _, tc := range []struct {
		name   string
		filter gotdbot.MessageFilter
		want   bool
	}{
		{"prefix", Prefix("/start"), true}, {"wrong prefix", Prefix("hello"), false},
		{"suffix", Suffix("😀"), true}, {"equal", Equal("/start hello😀"), true}, {"case sensitive", Equal("/START hello😀"), false},
		{"regex", Regex(`^/start\s+hello`), true}, {"regex no match", Regex(`^/help`), false}, {"invalid regex", Regex("["), false},
		{"private", Private, true}, {"group", Group, false}, {"supergroup", Supergroup, false}, {"secret", Secret, false}, {"channel", Channel, false},
		{"incoming", Incoming, true}, {"outgoing", Outgoing, false}, {"edited", Edited, true}, {"command", Command, true},
		{"sender", FromUserID(7), true}, {"wrong sender", FromUserID(8), false}, {"chat", ChatID(42), true}, {"wrong chat", ChatID(43), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.filter(m); got != tc.want {
				t.Fatalf("filter = %v, want %v", got, tc.want)
			}
		})
	}
	m.Content = &gotdbot.MessagePhoto{Caption: &gotdbot.FormattedText{Text: "caption"}}
	if !Prefix("cap")(m) || !Suffix("tion")(m) || !Equal("caption")(m) || !Regex(`^caption$`)(m) {
		t.Fatal("filters lost media caption")
	}
}

func TestRegexFilterConcurrentReuse(t *testing.T) {
	filter := Regex(`^hello\s+\d+$`)
	m := &gotdbot.Message{Content: &gotdbot.MessageText{Text: &gotdbot.FormattedText{Text: "hello 42"}}}
	var wg sync.WaitGroup
	for range 32 {
		wg.Go(func() {
			for range 32 {
				if !filter(m) {
					t.Error("shared regex changed result")
				}
			}
		})
	}
	wg.Wait()
}
