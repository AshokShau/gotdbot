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

func TestRichLinkTargets(t *testing.T) {
	for _, tc := range []struct {
		name, target string
		text         RichText
	}{
		{"mention", "https://t.me/telegram", RichTextMention{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}, Username: "telegram"}},
		{"local anchor", "#a%20%2B%23%25%26%22%2F%C3%A9", &RichTextAnchorLink{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}, AnchorName: "a +#%&\"/é"}},
		{"local reference", "#a%20%2B%23%25%26%22%2F%C3%A9", RichTextReferenceLink{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}, ReferenceName: "a +#%&\"/é"}},
		{"external anchor", "https://example.org/?a=1&amp;b=2#part", RichTextAnchorLink{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}, AnchorName: "part", Url: "https://example.org/?a=1&b=2#part"}},
		{"external reference", "https://example.org/#note", &RichTextReferenceLink{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}, ReferenceName: "note", Url: "https://example.org/#note"}},
		{"top", "#", RichTextAnchorLink{Text: RichTextBold{Text: RichTextPlain{Text: "Telegram"}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := RichTextHTML(tc.text), `<a href="`+tc.target+`"><b>Telegram</b></a>`; got != want {
				t.Errorf("HTML = %q, want %q", got, want)
			}
			if got, want := RichTextMarkdown(tc.text), `<a href="`+tc.target+`">**Telegram**</a>`; got != want {
				t.Errorf("Markdown = %q, want %q", got, want)
			}
			if got := RichTextContent(tc.text); got != "Telegram" {
				t.Errorf("content = %q", got)
			}
		})
	}
	if got := RichTextHTML(RichTextMention{Text: RichTextPlain{Text: "unknown"}}); got != "unknown" {
		t.Fatalf("missing username = %q", got)
	}
	message := &Message{Content: MessageText{Text: &FormattedText{
		Text: "@telegram", Entities: []TextEntity{{Length: 9, Type: &TextEntityTypeMention{}}},
	}}}
	if got, _ := message.OriginalRichHTML(); got != "<p><a href=\"https://t.me/telegram\">@telegram</a></p>\n" {
		t.Fatalf("converted mention = %q", got)
	}
}

func TestRichMarkdownNestedDelimiters(t *testing.T) {
	for _, tc := range []struct {
		name, want string
		text       RichText
	}{
		{"bold", "<b><b>nested</b></b>", RichTextBold{Text: RichTextBold{Text: RichTextPlain{Text: "nested"}}}},
		{"italic", "<i><i>nested</i></i>", RichTextItalic{Text: RichTextItalic{Text: RichTextPlain{Text: "nested"}}}},
		{"spoiler", "<tg-spoiler><tg-spoiler>nested</tg-spoiler></tg-spoiler>", RichTextSpoiler{Text: RichTextSpoiler{Text: RichTextPlain{Text: "nested"}}}},
		{"strikethrough", "<s><s>nested</s></s>", RichTextStrikethrough{Text: RichTextStrikethrough{Text: RichTextPlain{Text: "nested"}}}},
		{"marked", "<mark><mark>nested</mark></mark>", RichTextMarked{Text: RichTextMarked{Text: RichTextPlain{Text: "nested"}}}},
		{"nested through italic", "<b><i><b>nested</b></i></b>", RichTextBold{Text: RichTextItalic{Text: RichTextBold{Text: RichTextPlain{Text: "nested"}}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := RichTextMarkdown(tc.text); got != tc.want {
				t.Fatalf("Markdown = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestRichPhotoSkipsUnavailableSizes(t *testing.T) {
	block := PageBlockPhoto{Photo: &Photo{Sizes: []PhotoSize{
		{Width: 4000, Height: 3000, Photo: &File{}},
		{Width: 1280, Height: 720, Photo: &File{Id: 123}},
		{Width: 320, Height: 180, Photo: &File{Id: 456}},
	}}}
	for _, render := range []func(PageBlock) (string, []InputRichMessageMedia){RichBlockHTML, RichBlockMarkdown} {
		text, media := render(block)
		if !strings.Contains(text, `<img src="tg://photo?id=1"/>`) || len(media) != 1 {
			t.Fatalf("photo = %q, %#v", text, media)
		}
		photo := media[0].Media.(*InputMessagePhoto).Photo
		if photo.Photo.(*InputFileId).Id != 123 || photo.Width != 1280 || photo.Height != 720 {
			t.Fatalf("photo = %#v", photo)
		}
	}
}
