package gotdbot

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestParseButtonMarkup(t *testing.T) {
	for _, tc := range []struct {
		name, input, mode, body string
		labels, urls            []string
		rowSizes                []int
	}{
		{"markdown rows", "*Hello*\n[One](buttonurl:https://one.test)\n[Two](buttonurl:https://two.test:same)\n[Three](buttonurl://three.test)", "Markdown", "*Hello*", []string{"One", "Two", "Three"}, []string{"https://one.test", "https://two.test", "https://three.test"}, []int{2, 1}},
		{"balanced delimiters", `Hi [A [😀] B](buttonurl:https://example.test/a(b))`, "MarkdownV2", "Hi", []string{"A [😀] B"}, []string{"https://example.test/a(b)"}, []int{1}},
		{"escaped delimiters", `Hi [A \[B\] \\](buttonurl:https://example.test/a\)b)`, "MarkdownV2", "Hi", []string{`A [B] \`}, []string{"https://example.test/a)b"}, []int{1}},
		{"HTML shared syntax", `<b>Hi</b>[A & 😀](buttonurl:https://example.test/?a=1&b=2)`, "HTML", "<b>Hi</b>", []string{"A & 😀"}, []string{"https://example.test/?a=1&b=2"}, []int{1}},
		{"HTML anchors stay text", `<a href="buttonurl:https://example.test">Literal</a>`, "HTML", `<a href="buttonurl:https://example.test">Literal</a>`, nil, nil, nil},
		{"HTML protected shared syntax", `<pre><code>[Literal](buttonurl:https://example.test)</code></pre>`, "HTML", `<pre><code>[Literal](buttonurl:https://example.test)</code></pre>`, nil, nil, nil},
		{"first same", `[First](buttonurl:https://one.test:same)`, "markdown", "", []string{"First"}, []string{"https://one.test"}, []int{1}},
		{"normal link", "  [Normal](https://example.test)  ", "markdown", "  [Normal](https://example.test)  ", nil, nil, nil},
		{"escaped link", `\[Literal](buttonurl:https://example.test)`, "markdownv2", `\[Literal](buttonurl:https://example.test)`, nil, nil, nil},
		{"inline code", "`[Literal](buttonurl:https://example.test)`", "markdown", "`[Literal](buttonurl:https://example.test)`", nil, nil, nil},
		{"fenced code", "```go\n[Literal](buttonurl:https://example.test)\n```", "markdownv2", "```go\n[Literal](buttonurl:https://example.test)\n```", nil, nil, nil},
		{"HTML code", `<PRE><code><a href="buttonurl:https://example.test">Literal</a></code></PRE>`, "html", `<PRE><code><a href="buttonurl:https://example.test">Literal</a></code></PRE>`, nil, nil, nil},
		{"HTML escaped code", `<code>&lt;a href="buttonurl:https://example.test"&gt;Literal&lt;/a&gt;</code>`, "html", `<code>&lt;a href="buttonurl:https://example.test"&gt;Literal&lt;/a&gt;</code>`, nil, nil, nil},
		{"nested inside normal link", `[Outer [Button](buttonurl:https://example.test)](https://outer.test)`, "markdown", `[Outer [Button](buttonurl:https://example.test)](https://outer.test)`, nil, nil, nil},
		{"malformed link", "[Broken](buttonurl:https://example.test", "markdown", "[Broken](buttonurl:https://example.test", nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body, keyboard, err := ParseButtonMarkup(tc.input, tc.mode)
			if err != nil || body != tc.body {
				t.Fatalf("body = %q, err = %v; want %q", body, err, tc.body)
			}
			if tc.rowSizes == nil {
				if keyboard != nil {
					t.Fatal("unexpected keyboard")
				}
				return
			}
			var sizes []int
			var labels, urls []string
			for _, row := range keyboard.Rows {
				sizes = append(sizes, len(row))
				for _, button := range row {
					labels = append(labels, button.Text)
					urls = append(urls, button.Type.(*InlineKeyboardButtonTypeUrl).Url)
				}
			}
			if !reflect.DeepEqual(sizes, tc.rowSizes) || !reflect.DeepEqual(labels, tc.labels) || !reflect.DeepEqual(urls, tc.urls) {
				t.Fatalf("keyboard = %#v, %v, %v", sizes, labels, urls)
			}
		})
	}
}

func TestButtonMarkupActionsAndErrors(t *testing.T) {
	_, keyboard, err := ParseButtonMarkup(`Hello [Done](buttoncallback:menu:done) [Copy](buttoncopy:a&b:same)`, "MarkdownV2")
	if err != nil {
		t.Fatal(err)
	}
	if data := keyboard.Rows[0][0].Type.(*InlineKeyboardButtonTypeCallback).Data; string(data) != "menu:done" {
		t.Fatal(string(data))
	}
	if text := keyboard.Rows[0][1].Type.(*InlineKeyboardButtonTypeCopyText).Text; text != "a&b" {
		t.Fatal(text)
	}
	for _, source := range []string{`[ ](buttonurl:https://example.test)`, `[Empty](buttonurl:)`, `[Bad](buttonurl:javascript:alert(1))`, `[Bad](buttonurl:https://)`, `[Empty](buttoncallback:)`, `[Empty](buttoncopy:)`, `[Long](buttoncallback:` + strings.Repeat("😀", 17) + `)`} {
		if _, _, err := ParseButtonMarkup(source, "markdown"); err == nil {
			t.Errorf("accepted %q", source)
		}
	}
	if _, _, err := ParseButtonMarkup("text", "richmd"); err == nil {
		t.Fatal("accepted rich parse mode")
	}
}

func TestButtonMarkupRoundTrip(t *testing.T) {
	keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
		{{Text: `A [😀] & "<`, Type: &InlineKeyboardButtonTypeUrl{Url: "https://example.test/a(b)?a=1&b=2"}}, {Text: "Done", Type: &InlineKeyboardButtonTypeCallback{Data: []byte(`a\b(c)`)}}},
		{{Text: "Copy", Type: &InlineKeyboardButtonTypeCopyText{Text: "a(b)&c"}}},
	}}
	for _, mode := range []string{"HTML", "Markdown", "MarkdownV2"} {
		formatted := &FormattedText{Text: "Hello 😀", Entities: []TextEntity{{Length: 5, Type: TextEntityTypeBold{}}}}
		for _, content := range []MessageContent{&MessageText{Text: formatted}, &MessagePhoto{Caption: formatted}} {
			m := &Message{Content: content, ReplyMarkup: keyboard}
			source, err := m.OriginalTextWithButtons(mode)
			if err != nil {
				t.Fatal(err)
			}
			body, gotKeyboard, err := ParseButtonMarkup(source, mode)
			if err != nil || body != formatted.Format(mode) || !reflect.DeepEqual(gotKeyboard, keyboard) {
				t.Fatalf("%s: %q, %#v, %v", mode, body, gotKeyboard, err)
			}
		}
	}
}

func TestParseTextWithButtonsAndNativeSend(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(_ int, data []byte) {
		request := decodeTestRequest(t, data)
		var parsed struct {
			Text      string                 `json:"text"`
			ParseMode *TextParseModeMarkdown `json:"parse_mode"`
		}
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatal(err)
		}
		if request.Type != "parseTextEntities" || parsed.Text != "*Hi 😀*" || parsed.ParseMode.Version != 2 {
			t.Fatalf("wrong parser request: %s", data)
		}
		c.deliverResponse(request.Extra, &FormattedText{Text: "Hi 😀", Entities: []TextEntity{{Length: 5, Type: &TextEntityTypeBold{}}}})
	}
	text, keyboard, err := c.ParseTextWithButtons("*Hi 😀*\n[Go](buttonurl:https://example.test)", "MarkdownV2")
	if err != nil {
		t.Fatal(err)
	}
	requests := captureHelperRequests(t, c)
	m := &Message{ChatId: 10, Id: 20}
	if _, err := m.ReplyContent(c, &InputMessageText{Text: text}, &SendMessageContentOpts{ReplyMarkup: keyboard}); err != nil {
		t.Fatal(err)
	}
	if _, err := m.ReplyContent(c, &InputMessagePhoto{Photo: &InputPhoto{Photo: &InputFileId{Id: 123}}, Caption: text}, &SendMessageContentOpts{ReplyMarkup: keyboard}); err != nil {
		t.Fatal(err)
	}
	for _, request := range *requests {
		if !reflect.DeepEqual(helperRequestMarkup(t, request), keyboard) || helperReplyId(t, request) != 20 {
			t.Fatal("keyboard or reply lost")
		}
		switch content := helperRequestContent(t, request).(type) {
		case *InputMessageText:
			if !reflect.DeepEqual(content.Text, text) {
				t.Fatal("text entities lost")
			}
		case *InputMessagePhoto:
			if !reflect.DeepEqual(content.Caption, text) {
				t.Fatal("caption entities lost")
			}
		default:
			t.Fatalf("unexpected content %T", content)
		}
	}
}

func FuzzButtonMarkup(f *testing.F) {
	for _, input := range []string{"text", `[Go](buttonurl:https://example.test)`, `<a href="buttonurl:https://example.test">Go</a>`, "`[literal](buttonurl:https://example.test)`", `[x](buttonurl:a(b))`, `<a href='>`, string([]byte{0xff})} {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, text string) {
		if len(text) > 4096 {
			return
		}
		for _, mode := range []string{"html", "markdown", "markdownv2"} {
			body, keyboard, err := ParseButtonMarkup(text, mode)
			if err == nil && keyboard == nil && body != text {
				t.Fatal("changed text without buttons")
			}
			if err == nil && keyboard != nil {
				if source, formatErr := keyboard.Format(); formatErr == nil {
					_, restored, parseErr := ParseButtonMarkup(source, mode)
					if parseErr != nil || !reflect.DeepEqual(restored, keyboard) {
						t.Fatalf("keyboard round trip failed: %v", parseErr)
					}
				}
			}
		}
	})
}

func TestCompactButtonStorage(t *testing.T) {
	keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
		{{Text: "Docs", Type: &InlineKeyboardButtonTypeUrl{Url: "https://example.com"}, Style: &ButtonStylePrimary{}, IconCustomEmojiId: 123},
			{Text: "Done", Type: &InlineKeyboardButtonTypeCallback{Data: []byte("done")}}},
	}}
	source, err := keyboard.Format()
	if err != nil {
		t.Fatal(err)
	}
	want := "[Docs](buttonurl:https://example.com|style=primary|icon=123)\n[Done](buttoncallback:done:same)"
	if source != want {
		t.Fatalf("storage = %q; want %q", source, want)
	}
	for _, mode := range []string{ParseModeHTML, ParseModeMarkdown, ParseModeMarkdownV2} {
		body, restored, err := ParseButtonMarkup(source, mode)
		if err != nil || body != "" || !reflect.DeepEqual(restored, keyboard) {
			t.Fatalf("%s: body=%q keyboard=%#v err=%v", mode, body, restored, err)
		}
	}
}

func TestCompactButtonValueEscaping(t *testing.T) {
	for _, value := range []string{"a|style=danger", `a\|b(c)`, "a:same", `a\:same`, "a\r\nb", "a%20b%7Cc"} {
		keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
			{{Text: "Copy", Type: &InlineKeyboardButtonTypeCopyText{Text: value}},
				{Text: "Login", Type: &InlineKeyboardButtonTypeLoginUrl{Url: "https://example.test/a%20b", Id: -123, ForwardText: value}}},
		}}
		source, err := keyboard.Format()
		if err != nil {
			t.Fatal(err)
		}
		for _, mode := range []string{ParseModeHTML, ParseModeMarkdown, ParseModeMarkdownV2} {
			_, restored, err := ParseButtonMarkup(source, mode)
			if err != nil || !reflect.DeepEqual(restored, keyboard) {
				t.Fatalf("%s %q: source=%q keyboard=%#v err=%v", mode, value, source, restored, err)
			}
		}
	}
	_, keyboard, err := ParseButtonMarkup(`[Copy](buttoncopy:a\|b\(c\):same)`, ParseModeHTML)
	if err != nil || len(keyboard.Rows) != 1 || keyboard.Rows[0][0].Type.(*InlineKeyboardButtonTypeCopyText).Text != "a|b(c)" {
		t.Fatalf("manual escaped value: %#v, %v", keyboard, err)
	}
}

func TestReceivedHTMLButtonLiteral(t *testing.T) {
	literal := "[Literal](buttonurl:https://example.test)"
	keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
		{{Text: "Real", Type: &InlineKeyboardButtonTypeCallback{Data: []byte("real")}}},
	}}
	for _, markup := range []ReplyMarkup{nil, keyboard} {
		msg := &Message{Content: &MessageText{Text: &FormattedText{Text: literal}}, ReplyMarkup: markup}
		source, err := msg.OriginalTextWithButtons(ParseModeHTML)
		if err != nil {
			t.Fatal(err)
		}
		body, restored, err := ParseButtonMarkup(source, ParseModeHTML)
		if err != nil || body != "&#91;Literal](buttonurl:https://example.test)" {
			t.Fatalf("body=%q, err=%v", body, err)
		}
		if markup == nil && restored != nil || markup != nil && !reflect.DeepEqual(restored, keyboard) {
			t.Fatalf("literal produced a button: %#v", restored)
		}
	}
}

func TestSharedButtonsReplyText(t *testing.T) {
	for _, mode := range []string{ParseModeHTML, ParseModeMarkdown, ParseModeMarkdownV2} {
		for _, caption := range []bool{false, true} {
			t.Run(mode+"/caption="+fmt.Sprint(caption), func(t *testing.T) {
				c := newTestClient(t)
				formatted := &FormattedText{Text: "Hi 😀", Entities: []TextEntity{{Length: 5, Type: &TextEntityTypeBold{}}}}
				keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
					{{Text: "A & <B>", Type: &InlineKeyboardButtonTypeUrl{Url: "https://example.test/?a=1&b=2"}, Style: &ButtonStylePrimary{}, IconCustomEmojiId: 123}},
				}}
				original := &Message{Content: &MessageText{Text: formatted}, ReplyMarkup: keyboard}
				if caption {
					original.Content = &MessagePhoto{Caption: formatted}
				}
				source, err := original.OriginalTextWithButtons(mode)
				if err != nil {
					t.Fatal(err)
				}
				body, restored, err := ParseButtonMarkup(source, mode)
				if err != nil {
					t.Fatal(err)
				}
				requests := 0
				c.sendRequest = func(_ int, data []byte) {
					requests++
					var request helperWireRequest
					if err := json.Unmarshal(data, &request); err != nil {
						t.Fatal(err)
					}
					switch request.Type {
					case "parseTextEntities":
						var parse struct {
							Text string          `json:"text"`
							Mode json.RawMessage `json:"parse_mode"`
						}
						if err := json.Unmarshal(data, &parse); err != nil {
							t.Fatal(err)
						}
						expectedMode := `{"@type":"textParseModeHTML"}`
						if mode == ParseModeMarkdown {
							expectedMode = `{"@type":"textParseModeMarkdown","version":1}`
						} else if mode == ParseModeMarkdownV2 {
							expectedMode = `{"@type":"textParseModeMarkdown","version":2}`
						}
						if parse.Text != formatted.Format(mode) || string(parse.Mode) != expectedMode {
							t.Fatalf("wrong parse request: %s", data)
						}
						c.deliverResponse(request.Extra, formatted)
					case "sendMessage":
						content := helperRequestContent(t, request).(*InputMessageText)
						if request.ChatId != 10 || helperReplyId(t, request) != 20 || !reflect.DeepEqual(content.Text, formatted) || !reflect.DeepEqual(helperRequestMarkup(t, request), keyboard) {
							t.Fatalf("reply lost text, keyboard or destination: %s", data)
						}
						c.deliverResponse(request.Extra, &Message{ChatId: 10, Id: 99})
					default:
						t.Fatalf("unexpected request: %s", data)
					}
				}
				command := &Message{ChatId: 10, Id: 20}
				if _, err := command.ReplyText(c, body, &SendMessageOpts{ParseMode: mode, ReplyMarkup: restored}); err != nil {
					t.Fatal(err)
				}
				if requests != 2 {
					t.Fatalf("sent %d requests", requests)
				}
			})
		}
	}
}

func TestPlainTextButtons(t *testing.T) {
	c := newTestClient(t)
	c.sendRequest = func(_ int, _ []byte) {
		t.Fatal("plain text must not call TDLib's text parser")
	}
	plain := "*Hi* <b>there</b> `code` _text_"
	text, keyboard, err := c.ParseTextWithButtons(plain+"\n[Docs](buttonurl:https://example.test|style=primary)", ParseModeNone)
	if err != nil || text.Text != plain || len(text.Entities) != 0 || len(keyboard.Rows) != 1 {
		t.Fatalf("plain text=%#v keyboard=%#v err=%v", text, keyboard, err)
	}
	if _, _, err := ParseButtonMarkup("text", "unknown"); err == nil {
		t.Fatal("accepted unknown mode")
	}
}

func TestPlainTextButtonsRoundTrip(t *testing.T) {
	keyboard := &ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{
		{{Text: "Real", Type: &InlineKeyboardButtonTypeCallback{Data: []byte("real")}}},
	}}
	for _, plain := range []string{
		"Hello 😀", "*Hi* <b>there</b> `code` _text_", "[Literal](buttonurl:https://example.test)",
		`\[Literal](buttoncallback:data)`, `\\server\path [brackets]`, "`[Literal](buttoncopy:text)`",
	} {
		formatted := &FormattedText{Text: plain, Entities: []TextEntity{{Length: 1, Type: &TextEntityTypeBold{}}}}
		for _, content := range []MessageContent{&MessageText{Text: formatted}, &MessagePhoto{Caption: formatted}} {
			for _, markup := range []ReplyMarkup{nil, keyboard} {
				msg := &Message{Content: content, ReplyMarkup: markup}
				source, err := msg.OriginalTextWithButtons(ParseModeNone)
				if err != nil {
					t.Fatal(err)
				}
				body, restored, err := ParseButtonMarkup(source, ParseModeNone)
				if err != nil || body != plain {
					t.Fatalf("%q: body=%q err=%v", plain, body, err)
				}
				if markup == nil && restored != nil || markup != nil && !reflect.DeepEqual(restored, keyboard) {
					t.Fatalf("%q: literal became a button: %#v", plain, restored)
				}
			}
		}
	}
}

func allButtonTypes() []InlineKeyboardButtonType {
	return []InlineKeyboardButtonType{
		InlineKeyboardButtonTypeBuy{},
		&InlineKeyboardButtonTypeCallback{Data: []byte{0, 255, '&', ':', 's', 'a', 'm', 'e'}},
		InlineKeyboardButtonTypeCallbackGame{},
		&InlineKeyboardButtonTypeCallbackWithPassword{Data: []byte{0, 254, '\\', ')'}},
		InlineKeyboardButtonTypeCopyText{Text: "Copy 😀 & <text>:same"},
		InlineKeyboardButtonTypeDisabled{},
		&InlineKeyboardButtonTypeLoginUrl{Url: "https://example.test/?a=1&b=2", Id: -math.MaxInt64, ForwardText: "Forward 😀"},
		InlineKeyboardButtonTypeSwitchInline{Query: "inline 😀:same", TargetChat: &TargetChatCurrent{}},
		&InlineKeyboardButtonTypeUrl{Url: "https://example.test/:same"},
		InlineKeyboardButtonTypeUser{UserId: math.MaxInt64},
		&InlineKeyboardButtonTypeWebApp{Url: "https://example.test/app?a=1&b=2"},
		InlineKeyboardButtonTypeSwitchInline{Query: "chosen", TargetChat: &TargetChatChosen{Types: &TargetChatTypes{AllowUserChats: true, AllowBotChats: true, AllowGroupChats: true, AllowChannelChats: true}}},
		&InlineKeyboardButtonTypeSwitchInline{Query: "internal", TargetChat: TargetChatInternalLink{Link: &InternalLinkTypePublicChat{ChatUsername: "example", DraftText: "Hello 😀 & <text>", OpenProfile: true}}},
		InlineKeyboardButtonTypeSwitchInline{Query: "nil target"},
	}
}

func assertButtonRowsJSON(t *testing.T, got, want [][]InlineKeyboardButton) {
	t.Helper()
	gotJSON, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if string(gotJSON) != string(wantJSON) {
		t.Fatalf("got %s, want %s", gotJSON, wantJSON)
	}
}

func TestEveryInlineButtonTypeAndMetadataRoundTrip(t *testing.T) {
	styles := []ButtonStyle{nil, ButtonStyleDefault{}, &ButtonStylePrimary{}, ButtonStyleSuccess{}, &ButtonStyleDanger{}, ButtonStyleLink{}}
	for i, typ := range allButtonTypes() {
		t.Run(typ.GetType(), func(t *testing.T) {
			var icon int64
			if i%2 != 0 {
				icon = math.MaxInt64
			}
			button := InlineKeyboardButton{Text: "A [😀] <&> \\ `", Type: typ, Style: styles[i%len(styles)], IconCustomEmojiId: icon}
			keyboard := &ReplyMarkupInlineKeyboard{ForceReply: true, Rows: [][]InlineKeyboardButton{{button, {Text: "Next", Type: &InlineKeyboardButtonTypeDisabled{}}}, {{Text: "Last", Type: &InlineKeyboardButtonTypeCopyText{Text: "last:same"}}}}}
			before, _ := json.Marshal(keyboard)
			source, err := keyboard.Format()
			if err != nil {
				t.Fatal(err)
			}
			for _, mode := range []string{ParseModeNone, ParseModeHTML, ParseModeMarkdown, ParseModeMarkdownV2} {
				body, restored, err := ParseButtonMarkup("body\n"+source, mode)
				if err != nil || body != "body" || restored == nil {
					t.Fatalf("%s: body=%q err=%v", mode, body, err)
				}
				assertButtonRowsJSON(t, restored.Rows, keyboard.Rows)
				if restored.ForceReply {
					t.Fatal("ForceReply must be ignored")
				}
			}
			after, _ := json.Marshal(keyboard)
			if string(before) != string(after) {
				t.Fatal("render mutated keyboard")
			}
		})
	}
}

func TestButtonPayloadErrors(t *testing.T) {
	for _, target := range []string{
		"buttonurl:https://example.test|style=unknown", "buttonurl:https://example.test|icon=invalid",
		"buttonurl:https://example.test|icon=1|icon=2", "buttoncopy:text|unknown=x", "buttoncopy:text|style",
		"buttonbuy:value", "buttondisabled:value", "buttongame:value", "buttonuser:invalid",
		"buttonlogin:https://example.test|id=invalid", "buttonswitch:query|target=unknown",
		"buttonswitch:query|target=chosen|chats=unknown", "buttonswitch:query|target=link|link=null",
		"buttonswitch:query|target=link|link={}", "buttoncallback64:invalid!", "buttonpassword64:invalid!",
	} {
		if _, _, err := ParseButtonMarkup("[Button]("+target+")", "Markdown"); err == nil {
			t.Errorf("accepted %s", target)
		}
	}
	for _, typ := range []InlineKeyboardButtonType{nil, (*InlineKeyboardButtonTypeUrl)(nil)} {
		if _, err := (&ReplyMarkupInlineKeyboard{Rows: [][]InlineKeyboardButton{{{Text: "Button", Type: typ}}}}).Format(); err == nil {
			t.Fatal("accepted nil button type")
		}
	}
}
