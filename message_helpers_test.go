package gotdbot

import (
	"encoding/json"
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
	for _, content := range []MessageContent{&MessagePhoto{Caption: formatted}, &MessageVideo{Caption: formatted}, &MessageAnimation{Caption: formatted}, &MessageAudio{Caption: formatted}, &MessageDocument{Caption: formatted}, &MessageVoiceNote{Caption: formatted}, &MessagePaidMedia{Caption: formatted}} {
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

func TestMessageAccessorsHandleMissingContent(t *testing.T) {
	messages := []*Message{nil, {}, {Content: (*MessageText)(nil)}, {Content: &MessageText{}},
		{Content: (*MessagePhoto)(nil)}, {Content: &MessagePhoto{}},
		{Content: (*MessageVideo)(nil)}, {Content: &MessageVideo{}},
		{Content: &MessageAnimation{}}, {Content: &MessageAudio{}}, {Content: &MessageDocument{}},
		{Content: &MessageVoiceNote{}}, {Content: &MessageVideoNote{}}, {Content: &MessagePaidMedia{}},
		{SenderId: (*MessageSenderUser)(nil)}, {SenderId: (*MessageSenderChat)(nil)}, {ReplyTo: (*MessageReplyToMessage)(nil)},
	}
	for _, m := range messages {
		if m.GetText() != "" || m.Text() != "" || m.Caption() != "" || len(m.GetEntities()) != 0 || !m.Empty() || m.IsCommand() {
			t.Fatalf("unexpected text for %#v", m)
		}
		if m.ChatID() != 0 || m.SenderID() != 0 || m.ReplyToMessageID() != 0 || m.IsOutgoingMessage() || m.IsEditedMessage() || m.RemoteDuration() != 0 {
			t.Fatalf("unexpected metadata for %#v", m)
		}
		if _, err := m.GetFormattedText(); !errors.Is(err, ErrNoFormattedText) {
			t.Fatal(err)
		}
		if _, err := m.GetFormattedCaption(); !errors.Is(err, ErrNoFormattedText) {
			t.Fatal(err)
		}
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

func TestStopLiveLocationAndTypingPreserveOptions(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	location := &LiveLocation{Location: &Location{Latitude: 1, Longitude: 2}, LivePeriod: 60}
	markup := &ReplyMarkupInlineKeyboard{}
	opts := &EditMessageLiveLocationOpts{Location: location, ReplyMarkup: markup}
	before := *opts
	if _, err := (&Message{ChatId: 42, Id: 20}).StopLiveLocation(c, opts); err != nil {
		t.Fatal(err)
	}
	if (*requests)[0].Location != nil || !reflect.DeepEqual(helperRequestMarkup(t, (*requests)[0]), markup) || !reflect.DeepEqual(*opts, before) {
		t.Fatal("stop lost markup or changed caller location")
	}
	var request map[string]any
	c.sendRequest = func(_ int, data []byte) {
		if err := json.Unmarshal(data, &request); err != nil {
			t.Fatal(err)
		}
		c.deliverResponse(request["@extra"].(string), &Ok{})
	}
	actionOpts := &SendChatActionOpts{TopicId: &MessageTopicForum{ForumTopicId: 9}}
	if err := (&Message{ChatId: 42}).Action(c, actionOpts); err != nil {
		t.Fatal(err)
	}
	if actionOpts.Action != nil || request["action"].(map[string]any)["@type"] != "chatActionTyping" {
		t.Fatal("typing default changed caller options")
	}
}

func TestDownloadUsesAttachedFileIDWithoutRemoteLookup(t *testing.T) {
	c := newTestClient(t)
	file := &File{Id: 9, Remote: &RemoteFile{Id: "remote-id"}}
	message := &Message{Content: &MessagePhoto{Photo: &Photo{Sizes: []PhotoSize{
		{Width: 1280, Height: 720, Photo: file},
		{Width: 4000, Height: 3000},
		{Width: 100, Height: 100, Photo: &File{Id: 2}},
	}}}}
	if message.GetFile() != file {
		t.Fatal("file getter did not choose the largest available photo")
	}
	calls := 0
	c.sendRequest = func(_ int, data []byte) {
		calls++
		var req struct {
			Type        string `json:"@type"`
			Extra       string `json:"@extra"`
			FileID      int32  `json:"file_id"`
			Priority    int32  `json:"priority"`
			Offset      int64  `json:"offset"`
			Limit       int64  `json:"limit"`
			Synchronous bool   `json:"synchronous"`
		}
		if err := json.Unmarshal(data, &req); err != nil {
			t.Fatal(err)
		}
		if req.Type != "downloadFile" || req.FileID != 9 || req.Priority != 10 || req.Offset != 11 || req.Limit != 12 || !req.Synchronous {
			t.Fatalf("unexpected download request: %s", data)
		}
		c.deliverResponse(req.Extra, file)
	}
	got, err := message.Download(c, 10, &DownloadFileOpts{Offset: 11, Limit: 12, Synchronous: true})
	if err != nil || got != file || calls != 1 {
		t.Fatalf("download = %v, %v; calls = %d", got, err, calls)
	}
}

func TestDownloadWithoutMediaReturnsAnErrorWithoutSending(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(int, []byte) { t.Fatal("missing media sent a request") }
	for _, message := range []*Message{nil, {}, testMessage(42, "text"), {Content: &MessagePhoto{Photo: &Photo{}}}, {Content: (*MessageDocument)(nil)}} {
		if file, err := message.Download(c, 10, nil); file != nil || !errors.Is(err, ErrNoMedia) {
			t.Fatalf("missing media = %v, %v", file, err)
		}
	}
}

func TestMentionDispatchesBySenderType(t *testing.T) {
	for _, sender := range []MessageSender{
		&MessageSenderUser{UserId: 7}, MessageSenderUser{UserId: 7},
		&MessageSenderChat{ChatId: -42}, MessageSenderChat{ChatId: -42},
	} {
		c := newTestClient(t)
		calls := 0
		c.sendRequest = func(_ int, data []byte) {
			request := decodeTestRequest(t, data)
			calls++
			switch sender.(type) {
			case *MessageSenderUser, MessageSenderUser:
				if request.Type != "getUser" {
					t.Fatalf("user lookup used %s", request.Type)
				}
				c.deliverResponse(request.Extra, &User{Id: 7, FirstName: "<User>"})
			case *MessageSenderChat, MessageSenderChat:
				if request.Type != "getChat" {
					t.Fatalf("chat lookup used %s", request.Type)
				}
				c.deliverResponse(request.Extra, &Chat{Title: "<Chat>"})
			}
		}
		got, err := (&Message{SenderId: sender}).Mention(c, "HTML")
		want := "<Chat>"
		switch sender.(type) {
		case *MessageSenderUser, MessageSenderUser:
			want = `<a href="tg://user?id=7">&lt;User&gt;</a>`
		}
		if err != nil || got != want || calls != 1 {
			t.Fatalf("%T: mention/error/calls=%q/%v/%d", sender, got, err, calls)
		}
	}
}

func TestMentionAbsentSenderAndLookupError(t *testing.T) {
	c := newTestClient(t)
	for _, message := range []*Message{nil, {}, {SenderId: (*MessageSenderUser)(nil)}, {SenderId: (*MessageSenderChat)(nil)}} {
		if _, err := message.Mention(c, "HTML"); err == nil {
			t.Fatal("missing sender accepted")
		}
	}
	failure := &Error{Code: 400, Message: "USER_NOT_FOUND"}
	calls := 0
	c.sendRequest = func(_ int, data []byte) {
		request := decodeTestRequest(t, data)
		calls++
		if request.Type != "getUser" {
			t.Fatalf("failed user lookup fell back to %s", request.Type)
		}
		c.deliverResponse(request.Extra, failure)
	}
	if _, err := (&Message{SenderId: &MessageSenderUser{UserId: 7}}).Mention(c, "HTML"); err != failure || calls != 1 {
		t.Fatalf("error/calls=%v/%d", err, calls)
	}
}

func TestMessageWebAppData(t *testing.T) {
	content := MessageWebAppDataReceived{ButtonText: "Open app", Data: `{"item":"hat","quantity":2}`}
	for _, value := range []MessageContent{content, &content} {
		message := &Message{Content: value}
		data, err := message.GetWebAppData()
		if err != nil || data.ButtonText != content.ButtonText || data.Data != content.Data {
			t.Fatalf("Web App data = %v, %v", data, err)
		}
		var payload struct {
			Item     string `json:"item"`
			Quantity int    `json:"quantity"`
		}
		if err := message.DecodeWebAppData(&payload); err != nil || payload.Item != "hat" || payload.Quantity != 2 {
			t.Fatalf("decoded payload = %+v, %v", payload, err)
		}
		if err := message.DecodeWebAppData(nil); err == nil {
			t.Fatal("nil destination accepted")
		}
	}
	for _, message := range []*Message{nil, {}, {Content: (*MessageWebAppDataReceived)(nil)}, {Content: MessageText{}}, {Content: MessageWebAppDataSent{}}} {
		if _, err := message.GetWebAppData(); !errors.Is(err, ErrNoWebAppData) {
			t.Fatalf("missing Web App data = %v", err)
		}
		if err := message.DecodeWebAppData(new(any)); !errors.Is(err, ErrNoWebAppData) {
			t.Fatalf("decoded missing data = %v", err)
		}
	}
	message := &Message{Content: MessageWebAppDataReceived{Data: "not JSON"}}
	if err := message.DecodeWebAppData(new(any)); err == nil {
		t.Fatal("invalid JSON accepted")
	}
}
