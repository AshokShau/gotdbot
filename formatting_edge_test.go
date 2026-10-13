package gotdbot

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf16"
	"unicode/utf8"
)

func TestUnparseUnicodeAndNestedEntities(t *testing.T) {
	for _, tc := range []struct {
		name, text, mode, want string
		entities               []TextEntity
	}{
		{"UTF16 emoji", "A😀B", "html", "A<b>😀</b>B", []TextEntity{{Offset: 1, Length: 2, Type: &TextEntityTypeBold{}}}},
		{"combining mark", "e\u0301!", "html", "<i>e\u0301</i>!", []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeItalic{}}}},
		{"nested", "hello", "html", "<b>h<i>ell</i>o</b>", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeBold{}}, {Offset: 1, Length: 3, Type: &TextEntityTypeItalic{}}}},
		{"same span", "hello", "html", "<b><i>hello</i></b>", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeBold{}}, {Offset: 0, Length: 5, Type: &TextEntityTypeItalic{}}}},
		{"adjacent", "abcd", "html", "<b>ab</b><i>cd</i>", []TextEntity{{Offset: 0, Length: 2, Type: &TextEntityTypeBold{}}, {Offset: 2, Length: 2, Type: &TextEntityTypeItalic{}}}},
		{"code escapes", "a`b\\c", "markdownv2", "`a\\`b\\\\c`", []TextEntity{{Offset: 0, Length: 5, Type: &TextEntityTypeCode{}}}},
		{"HTML escaping", "<&>", "html", "<b>&lt;&amp;&gt;</b>", []TextEntity{{Offset: 0, Length: 3, Type: &TextEntityTypeBold{}}}},
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
