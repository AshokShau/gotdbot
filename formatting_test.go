package gotdbot

import (
	"math"
	"reflect"
	"strings"
	"testing"
	"unicode/utf16"
	"unicode/utf8"
)

func TestFormattedTextFormat(t *testing.T) {
	for _, tc := range []struct {
		name, text, mode, want string
		entities               []TextEntity
	}{
		{"Plain text HTML", "Hello <World> & Everyone", "html", "Hello &lt;World&gt; &amp; Everyone", nil},
		{"Plain text MDV2", "Hello. World!", "markdownv2", "Hello\\. World\\!", nil},
		{"Bold and Italic HTML", "Bold and Italic", "html", "<b>Bold</b> and <i>Italic</i>", []TextEntity{{Offset: 0, Length: 4, Type: &TextEntityTypeBold{}}, {Offset: 9, Length: 6, Type: &TextEntityTypeItalic{}}}},
		{"Bold and Italic MDV2", "Bold and Italic", "markdownv2", "*Bold* and _Italic_", []TextEntity{{Offset: 0, Length: 4, Type: &TextEntityTypeBold{}}, {Offset: 9, Length: 6, Type: &TextEntityTypeItalic{}}}},
		{"Mention Name HTML", "User Name", "html", "<a href=\"tg://user?id=123456789\">User Name</a>", []TextEntity{{Offset: 0, Length: 9, Type: &TextEntityTypeMentionName{UserId: 123456789}}}},
		{"Mention Name MDV2", "User Name", "markdownv2", "[User Name](tg://user?id=123456789)", []TextEntity{{Offset: 0, Length: 9, Type: &TextEntityTypeMentionName{UserId: 123456789}}}},
		{"Text URL HTML", "Open Link", "html", "<a href=\"https://example.com\">Open Link</a>", []TextEntity{{Offset: 0, Length: 9, Type: &TextEntityTypeTextUrl{Url: "https://example.com"}}}},
		{"Text URL MDV2", "Open Link", "markdownv2", "[Open Link](https://example.com/test?a=1&b=2)", []TextEntity{{Offset: 0, Length: 9, Type: &TextEntityTypeTextUrl{Url: "https://example.com/test?a=1&b=2"}}}},
		{"Custom Emoji HTML", "Smile", "html", "<tg-emoji emoji-id=\"123456\">Smile</tg-emoji>", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeCustomEmoji{CustomEmojiId: 123456}}}},
		{"Custom Emoji MDV2", "Smile", "markdownv2", "![Smile](tg://emoji?id=123456)", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeCustomEmoji{CustomEmojiId: 123456}}}},
		{"Date Time HTML", "tomorrow", "html", "<tg-time unix=\"1234567890\" format=\"r\">tomorrow</tg-time>", []TextEntity{{Offset: 0, Length: 8, Type: &TextEntityTypeDateTime{UnixTime: 1234567890, FormattingType: &DateTimeFormattingTypeRelative{}}}}},
		{"Date Time MDV2", "tomorrow", "markdownv2", "![tomorrow](tg://time?unix=1234567890&format=r)", []TextEntity{{Offset: 0, Length: 8, Type: &TextEntityTypeDateTime{UnixTime: 1234567890, FormattingType: &DateTimeFormattingTypeRelative{}}}}},
		{"Implicit Entity Mention MDV2", "@username", "markdownv2", "@username", []TextEntity{{Offset: 0, Length: 9, Type: &TextEntityTypeMention{}}}},
		{"Implicit Entity URL MDV2", "https://example.com", "markdownv2", `https://example\.com`, []TextEntity{{Offset: 0, Length: 19, Type: &TextEntityTypeUrl{}}}},
		{"Pre Code HTML", "print('hello')", "html", "<pre><code class=\"language-python\">print('hello')</code></pre>", []TextEntity{{Offset: 0, Length: 14, Type: &TextEntityTypePreCode{Language: "python"}}}},
		{"Pre Code MDV2", "print('hello')", "markdownv2", "```python\nprint('hello')```", []TextEntity{{Offset: 0, Length: 14, Type: &TextEntityTypePreCode{Language: "python"}}}},
		{"Underline MDV2 Ambiguity Test", "under", "markdownv2", "__under__", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeUnderline{}}}},
		{"UTF16 emoji", "A😀B", "html", "A<b>😀</b>B", []TextEntity{{Offset: 1, Length: 2, Type: &TextEntityTypeBold{}}}},
		{"combining mark", "e\u0301!", "html", "<i>e\u0301</i>!", []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeItalic{}}}},
		{"nested", "hello", "html", "<b>h<i>ell</i>o</b>", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeBold{}}, {Offset: 1, Length: 3, Type: &TextEntityTypeItalic{}}}},
		{"same span", "hello", "html", "<b><i>hello</i></b>", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeBold{}}, {Offset: 0, Length: 5, Type: &TextEntityTypeItalic{}}}},
		{"adjacent", "abcd", "html", "<b>ab</b><i>cd</i>", []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeBold{}}, {Offset: 2, Length: 2, Type: &TextEntityTypeItalic{}}}},
		{"code escapes", "a`b\\c", "markdownv2", "`a\\`b\\\\c`", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeCode{}}}},
		{"HTML escaping", "<&>", "html", "<b>&lt;&amp;&gt;</b>", []TextEntity{{Offset: 0, Length: 3, Type: &TextEntityTypeBold{}}}},
		{"unsorted nested", "hello", "html", "<b>h<i>ell</i>o</b>", []TextEntity{{Offset: 1, Length: 3, Type: TextEntityTypeItalic{}}, {Length: 5, Type: TextEntityTypeBold{}}}},
		{"same start wider last", "hello", "html", "<b><i>hel</i>lo</b>", []TextEntity{{Length: 3, Type: TextEntityTypeItalic{}}, {Length: 5, Type: TextEntityTypeBold{}}}},
		{"same end", "hello", "html", "<b>he<i>llo</i></b>", []TextEntity{{Offset: 2, Length: 3, Type: TextEntityTypeItalic{}}, {Length: 5, Type: TextEntityTypeBold{}}}},
		{"implicit HTML entity", "a&<b", "html", "a&amp;&lt;b", []TextEntity{{Length: 4, Type: TextEntityTypeUrl{}}}},
		{"implicit Markdown entity", "@a_b", "markdownv2", `@a\_b`, []TextEntity{{Length: 4, Type: TextEntityTypeMention{}}}},
		{"HTML attribute", "x", "html", `<a href="https://x/?a=&#34;&amp;b=1">x</a>`, []TextEntity{{Length: 1, Type: TextEntityTypeTextUrl{Url: `https://x/?a="&b=1`}}}},
		{"code language attribute", "x", "html", `<pre><code class="language-&#34;&gt;">x</code></pre>`, []TextEntity{{Length: 1, Type: TextEntityTypePreCode{Language: `">`}}}},
		{"Markdown URL", "x", "markdownv2", `[x](https://x/\)\\)`, []TextEntity{{Length: 1, Type: TextEntityTypeTextUrl{Url: "https://x/)\\"}}}},
		{"underline italic separation", "x", "markdownv2", "__\r_x_\r__", []TextEntity{{Length: 1, Type: TextEntityTypeUnderline{}}, {Length: 1, Type: TextEntityTypeItalic{}}}},
		{"italic underline separation", "x", "markdownv2", "_\r__x__\r_", []TextEntity{{Length: 1, Type: TextEntityTypeItalic{}}, {Length: 1, Type: TextEntityTypeUnderline{}}}},
		{"multiline quote with nested style", "a\nb", "markdownv2", ">*a\n>b*", []TextEntity{{Length: 3, Type: TextEntityTypeBlockQuote{}}, {Length: 3, Type: TextEntityTypeBold{}}}},
		{"pre doesn't append text", "x!", "markdownv2", "```\nx```\\!", []TextEntity{{Length: 1, Type: TextEntityTypePre{}}}},
		{"legacy drops nesting", "abc", "markdown", "*abc*", []TextEntity{{Length: 3, Type: TextEntityTypeBold{}}, {Offset: 1, Length: 1, Type: TextEntityTypeItalic{}}}},
		{"value date precision", "x", "html", `<tg-time unix="1" format="wDT">x</tg-time>`, []TextEntity{{Length: 1, Type: TextEntityTypeDateTime{UnixTime: 1, FormattingType: DateTimeFormattingTypeAbsolute{ShowDayOfWeek: true, DatePrecision: DateTimePartPrecisionLong{}, TimePrecision: DateTimePartPrecisionLong{}}}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			before := append([]TextEntity(nil), tc.entities...)
			if got := (&FormattedText{Text: tc.text, Entities: tc.entities}).Format(tc.mode); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
			if !reflect.DeepEqual(before, tc.entities) {
				t.Fatal("formatting mutated entities")
			}
		})
	}
}

func TestFormattedTextFormatHandlesModesAndNil(t *testing.T) {
	var absent *FormattedText
	if absent.Format("HTML") != "" {
		t.Fatal("nil formatted text produced output")
	}
	text := &FormattedText{Text: "<text>", Entities: []TextEntity{{Offset: 0, Length: 6, Type: &TextEntityTypeBold{}}}}
	if text.Format("HTML") != "<b>&lt;text&gt;</b>" || text.Format("MarkdownV2") != "*<text\\>*" || text.Format("unknown") != "<text>" {
		t.Fatal("format mode handling failed")
	}
}

func TestFormattingConvenienceHelpers(t *testing.T) {
	if got := (&FormattedText{Text: "_*[]()!\\"}).Format("MarkdownV2"); got != "\\_\\*\\[\\]\\(\\)\\!\\\\" {
		t.Fatalf("Markdown escaping: %q", got)
	}
	if got := (&User{Id: 42, FirstName: "<name>"}).Mention("HTML"); got != `<a href="tg://user?id=42">&lt;name&gt;</a>` {
		t.Fatal(got)
	}
	if got := (&User{Id: 42, FirstName: "a*b"}).Mention("MarkdownV2"); got != `[a\*b](tg://user?id=42)` {
		t.Fatal(got)
	}
	m := &Message{Content: &MessageText{Text: &FormattedText{Text: "hi", Entities: []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeBold{}}}}}}
	if m.OriginalHTML() != "<b>hi</b>" || m.OriginalMD() != "*hi*" || m.OriginalMDV2() != "*hi*" {
		t.Fatal("original text formatting failed")
	}
	m.Content = &MessagePhoto{Caption: &FormattedText{Text: "hi", Entities: []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeBold{}}}}}
	if m.OriginalCaptionHTML() != "<b>hi</b>" || m.OriginalCaptionMD() != "*hi*" || m.OriginalCaptionMDV2() != "*hi*" {
		t.Fatal("original caption formatting failed")
	}
}

func TestFormattedTextExplicitEntitiesAndParseMode(t *testing.T) {
	c := newTestClient(t)
	entities := []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeBold{}}}
	c.sendRequest = func(int, []byte) { t.Error("explicit entities/plain text unexpectedly reached TDLib") }
	formatted, err := c.GetFormattedText("hi", entities, "invalid mode")
	if err != nil || formatted.Text != "hi" || !reflect.DeepEqual(formatted.Entities, entities) {
		t.Fatalf("explicit entities = (%v,%v)", formatted, err)
	}
	formatted, err = c.GetFormattedText("plain", nil, "")
	if err != nil || formatted.Text != "plain" {
		t.Fatalf("plain = (%v,%v)", formatted, err)
	}
}

func FuzzUnparseEntitiesUnicode(f *testing.F) {
	for _, text := range []string{"hello", "😀 привет & <world>", "e\u0301", "日本語", "", "a`b\\c"} {
		f.Add(text)
	}
	f.Fuzz(func(t *testing.T, text string) {
		if !utf8.ValidString(text) || len(text) > 4096 || text == "" {
			return
		}
		entities := []TextEntity{{Offset: 0, Length: int32(len(utf16.Encode([]rune(text)))), Type: &TextEntityTypeBold{}}}
		want := "<b>" + strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(text) + "</b>"
		if got := (&FormattedText{Text: text, Entities: entities}).Format("html"); got != want {
			t.Fatalf("Unicode formatting = %q, want %q", got, want)
		}
	})
}

func TestMalformedFormattingRanges(t *testing.T) {
	for _, entity := range []TextEntity{
		{Offset: -1, Length: 1, Type: TextEntityTypeBold{}},
		{Length: -1, Type: TextEntityTypeBold{}},
		{Length: 0, Type: TextEntityTypeBold{}},
		{Offset: math.MaxInt32, Length: math.MaxInt32, Type: TextEntityTypeBold{}},
		{Length: 100, Type: TextEntityTypeBold{}},
		{Offset: 1, Length: 1, Type: TextEntityTypeBold{}}, // half an emoji
		{Offset: 2, Length: 1, Type: TextEntityTypeBold{}},
		{Length: 1, Type: (*TextEntityTypeBold)(nil)},
		{Length: 1},
	} {
		f := &FormattedText{Text: "a😀b", Entities: []TextEntity{entity}}
		if got := f.Format("HTML"); got != f.Text {
			t.Errorf("%+v: got %q", entity, got)
		}
	}
	f := &FormattedText{Text: "abcd", Entities: []TextEntity{{Length: 3, Type: TextEntityTypeBold{}}, {Offset: 1, Length: 3, Type: TextEntityTypeItalic{}}}}
	if got := f.Format("html"); got != "<b>abc</b>d" {
		t.Fatal(got)
	}
}

func TestOriginalTextSelectsTextOrCaption(t *testing.T) {
	f := &FormattedText{Text: "hi", Entities: []TextEntity{{Length: 2, Type: TextEntityTypeBold{}}}}
	for _, content := range []MessageContent{&MessageText{Text: f}, MessagePhoto{Caption: f}} {
		m := &Message{Content: content}
		if m.OriginalTextHTML() != "<b>hi</b>" || m.OriginalTextMD() != "*hi*" || m.OriginalTextMDV2() != "*hi*" {
			t.Errorf("%T: conversion failed", content)
		}
	}
	var missing *Message
	if missing.OriginalTextHTML() != "" || missing.OriginalTextMD() != "" || missing.OriginalTextMDV2() != "" {
		t.Fatal("nil message produced text")
	}
}

func FuzzFormattingRanges(f *testing.F) {
	f.Add("a😀b", int32(-1), int32(3))
	f.Add("<&_", int32(0), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, text string, offset, length int32) {
		if !utf8.ValidString(text) || len(text) > 4096 {
			return
		}
		formatted := &FormattedText{Text: text, Entities: []TextEntity{{Offset: offset, Length: length, Type: TextEntityTypeBold{}}}}
		if got := formatted.Format("html"); !utf8.ValidString(got) {
			t.Fatal("invalid UTF-8")
		}
		rich := formattedTextRichMessage(formatted)
		for _, block := range rich.Blocks {
			if p, ok := block.(PageBlockParagraph); ok && strings.Contains(RichTextContent(p.Text), "\ufffd") && !strings.Contains(text, "\ufffd") {
				t.Fatal("corrupted text")
			}
		}
	})
}
