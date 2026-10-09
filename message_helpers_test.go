package gotdbot

import (
	"errors"
	"reflect"
	"testing"
)

func TestChatTypeBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name    string
		id      int64
		channel bool
		kind    string
	}{
		{"zero", 0, false, ""}, {"private minimum", 1, false, "private"}, {"private maximum", (1 << 40) - 1, false, "private"}, {"private overflow", 1 << 40, false, ""},
		{"group minimum", supergroupChatIDOffset + 1, false, "group"}, {"group maximum", -1, false, "group"}, {"supergroup offset", supergroupChatIDOffset, false, ""},
		{"supergroup", supergroupChatIDOffset - 1, false, "supergroup"}, {"channel", supergroupChatIDOffset - 1, true, "channel"},
		{"secret minimum", minSecretChatID, false, "secret"}, {"secret maximum", maxSecretChatID, false, "secret"}, {"secret offset", secretChatIDOffset, false, ""},
		{"supergroup below secret", minSecretChatID - 1, false, "supergroup"}, {"supergroup above secret", maxSecretChatID + 1, false, "supergroup"}, {"supergroup lower bound", -4000000000000, false, "supergroup"}, {"below lower bound", -4000000000001, false, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := &Message{ChatId: tc.id, IsChannelPost: tc.channel}
			matches := map[string]bool{"private": m.IsPrivate(), "group": m.IsGroup(), "supergroup": m.IsSupergroup(), "channel": m.IsChannel(), "secret": m.IsSecret()}
			for kind, matched := range matches {
				if matched != (kind == tc.kind) {
					t.Errorf("%s = %v for ID %d", kind, matched, tc.id)
				}
			}
		})
	}
	var m *Message
	if m.IsPrivate() || m.IsGroup() || m.IsSupergroup() || m.IsChannel() || m.IsSecret() {
		t.Fatal("nil message classified")
	}
}

func TestMessageTextCaptionAndArguments(t *testing.T) {
	entities := []TextEntity{{Offset: 0, Length: 6, Type: &TextEntityTypeBotCommand{}}}
	formatted := &FormattedText{Text: "/start  one\t two\nthree", Entities: entities}
	m := &Message{Content: &MessageText{Text: formatted}}
	if m.GetText() != formatted.Text || !m.IsCommand() || m.Args() != "one two three" || !reflect.DeepEqual(m.ArgsList(), []string{"one", "two", "three"}) {
		t.Fatal("message text/arguments mismatch")
	}
	for index, want := range map[int]string{-1: "", 0: "one", 1: "two", 2: "three", 3: ""} {
		if got := m.Arg(index); got != want {
			t.Errorf("Arg(%d) = %q", index, got)
		}
	}
	if got, err := m.GetFormattedText(); got != formatted || err != nil {
		t.Fatalf("formatted = (%v,%v)", got, err)
	}
	if !reflect.DeepEqual(m.GetEntities(), entities) {
		t.Fatal("text entities missing")
	}
	for _, content := range []MessageContent{&MessagePhoto{Caption: formatted}, &MessageVideo{Caption: formatted}, &MessageAnimation{Caption: formatted}, &MessageAudio{Caption: formatted}, &MessageDocument{Caption: formatted}, &MessageVoiceNote{Caption: formatted}} {
		m := &Message{Content: content}
		if m.GetText() != formatted.Text || m.Text() != "" {
			t.Errorf("%T text/caption mismatch", content)
		}
		if got, err := m.GetFormattedCaption(); got != formatted || err != nil {
			t.Errorf("%T caption = (%v,%v)", content, got, err)
		}
		if !reflect.DeepEqual(m.GetEntities(), entities) {
			t.Errorf("%T lost caption entities", content)
		}
	}
	for _, m := range []*Message{nil, {}, {Content: &MessagePhoto{Caption: formatted}}} {
		if _, err := m.GetFormattedText(); !errors.Is(err, ErrNoFormattedText) {
			t.Errorf("unexpected text error: %v", err)
		}
	}
	if _, err := (&Message{Content: &MessageText{Text: formatted}}).GetFormattedCaption(); !errors.Is(err, ErrNoFormattedText) {
		t.Fatal(err)
	}
	plain := testMessage(42, "  hello world  ")
	if !plain.StartsWith("  hello") || !plain.EndsWith("world  ") || plain.TrimmedText() != "hello world" {
		t.Fatal("text predicates failed")
	}
	if plain.IsCommand() || testMessage(42, "/start").IsCommand() {
		t.Fatal("command inferred without entity")
	}
}

func TestSenderAndReplyHelpers(t *testing.T) {
	for _, tc := range []struct {
		sender MessageSender
		want   int64
	}{{nil, 0}, {&MessageSenderUser{UserId: 42}, 42}, {&MessageSenderChat{ChatId: -42}, -42}} {
		if got := (&Message{SenderId: tc.sender}).SenderID(); got != tc.want {
			t.Errorf("sender %T = %d", tc.sender, got)
		}
	}
	m := &Message{ChatId: 42, ReplyTo: &MessageReplyToMessage{MessageId: 7}}
	if m.ChatID() != 42 || m.ReplyToMessageID() != 7 {
		t.Fatal("chat/reply IDs missing")
	}
	m.ReplyTo = nil
	if m.ReplyToMessageID() != 0 {
		t.Fatal("nil reply returned ID")
	}
}

func TestCallbackPayloadHelpers(t *testing.T) {
	for _, tc := range []struct {
		name                 string
		payload              CallbackQueryPayload
		data, password, game string
	}{
		{"empty", nil, "", "", ""}, {"data", &CallbackQueryPayloadData{Data: []byte("hello😀")}, "hello😀", "", ""},
		{"password", &CallbackQueryPayloadDataWithPassword{Data: []byte("hello"), Password: "secret"}, "hello", "secret", ""},
		{"game", &CallbackQueryPayloadGame{GameShortName: "game"}, "", "", "game"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			u := &UpdateNewCallbackQuery{Payload: tc.payload}
			inline := &UpdateNewInlineCallbackQuery{Payload: tc.payload}
			if u.DataString() != tc.data || string(u.CallbackData()) != tc.data || u.GameShortName() != tc.game || inline.DataString() != tc.data || inline.GameShortName() != tc.game {
				t.Fatal("payload helpers changed data")
			}
			data, password := u.CallbackDataWithPassword()
			if password != tc.password || (tc.password != "" && string(data) != tc.data) {
				t.Fatal("password payload missing")
			}
		})
	}
}
