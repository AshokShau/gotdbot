package gotdbot

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestRichInlineFormatting(t *testing.T) {
	for _, tc := range []struct {
		name            string
		text            RichText
		html, md, plain string
	}{
		{"escaped text", &RichTextPlain{Text: "<b>_*~=$\n😀"}, "&lt;b&gt;_*~=$<br/>😀", "&lt;b&gt;\\_\\*\\~\\=\\$<br/>😀", "<b>_*~=$\n😀"},
		{"bold", RichTextBold{Text: &RichTextItalic{Text: RichTextPlain{Text: "hi"}}}, "<b><i>hi</i></b>", "**_hi_**", "hi"},
		{"underline literal", RichTextUnderline{Text: RichTextPlain{Text: "*hi*"}}, "<u>*hi*</u>", "<u>\\*hi\\*</u>", "*hi*"},
		{"code", &RichTextFixed{Text: RichTextPlain{Text: "`_*<&\n"}}, "<code>`_*&lt;&amp;\n</code>", "<code>`_*&lt;&amp;\n</code>", "`_*<&\n"},
		{"link", RichTextUrl{Url: `https://x/?q="&a=(b)`, Text: RichTextPlain{Text: "*x*"}}, `<a href="https://x/?q=&#34;&amp;a=(b)">*x*</a>`, `<a href="https://x/?q=&#34;&amp;a=(b)">\*x\*</a>`, "*x*"},
		{"custom emoji", RichTextCustomEmoji{CustomEmojiId: 123, AlternativeText: "😀"}, `<tg-emoji emoji-id="123">😀</tg-emoji>`, `<tg-emoji emoji-id="123">😀</tg-emoji>`, "😀"},
		{"time", RichTextDateTime{UnixTime: 42, FormattingType: DateTimeFormattingTypeRelative{}, Text: RichTextPlain{Text: "now"}}, `<tg-time unix="42" format="r">now</tg-time>`, `<tg-time unix="42" format="r">now</tg-time>`, "now"},
		{"math", RichTextMathematicalExpression{Expression: "x<y"}, "<tg-math>x&lt;y</tg-math>", "<tg-math>x&lt;y</tg-math>", "x<y"},
		{"array and nil", RichTexts{Texts: []RichText{nil, (*RichTextBold)(nil), RichTextPlain{Text: "x"}}}, "x", "x", "x"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := RichTextHTML(tc.text); got != tc.html {
				t.Errorf("HTML = %q, want %q", got, tc.html)
			}
			if got := RichTextMarkdown(tc.text); got != tc.md {
				t.Errorf("MD = %q, want %q", got, tc.md)
			}
			if got := RichTextContent(tc.text); got != tc.plain {
				t.Errorf("plain = %q, want %q", got, tc.plain)
			}
		})
	}
}

func TestRichBlockStructureAndMedia(t *testing.T) {
	p := func(s string) RichText { return &RichTextPlain{Text: s} }
	rich := &RichMessage{IsFull: true, IsRtl: true, Blocks: []PageBlock{
		&PageBlockSectionHeading{Size: 2, Text: p("Title")},
		&PageBlockDetails{Header: p("More"), IsOpen: true, Blocks: []PageBlock{
			PageBlockPreformatted{Language: "go", Text: p("<&\n`code`")},
			PageBlockList{Items: []PageBlockListItem{{Value: 3, Type: "i", Blocks: []PageBlock{PageBlockParagraph{Text: p("three")}}}, {Value: 7, Blocks: []PageBlock{PageBlockParagraph{Text: p("seven")}}}, {HasCheckbox: true, IsChecked: true, Blocks: []PageBlock{PageBlockParagraph{Text: p("done")}}}}},
			PageBlockPhoto{HasSpoiler: true, Photo: &Photo{Sizes: []PhotoSize{{Width: 100, Height: 100, Photo: &File{Id: 10}}, {Width: 1, Height: 1, Photo: &File{Id: 11}}}}, Caption: &PageBlockCaption{Text: p("photo"), Credit: p("credit")}},
		}},
		PageBlockTable{IsBordered: true, IsStriped: true, IsCompact: true, Caption: p("Table"), Cells: [][]PageBlockTableCell{{{IsHeader: true, Text: p("head"), Colspan: 2, Rowspan: 3, Align: PageBlockHorizontalAlignmentRight{}, Valign: &PageBlockVerticalAlignmentTop{}}, {Text: nil}}}},
		PageBlockBlockQuote{Blocks: []PageBlock{PageBlockParagraph{Text: p("a\nb")}}, Credit: p("author")},
		PageBlockExpandableBlockQuote{Text: p("hidden")},
		PageBlockAudio{Audio: &Audio{Audio: &File{Id: 20}, Duration: 5, Title: "song"}},
	}}
	before := append([]PageBlock(nil), rich.Blocks...)
	for _, markdown := range []bool{false, true} {
		text, media := rich.HTML()
		if markdown {
			text, media = rich.Markdown()
		}
		for _, want := range []string{`<details open><summary>More</summary>`, "<pre><code class=\"language-go\">&lt;&amp;\n`code`</code></pre>", `<li value="3" type="i">`, `<li value="7">`, `<input type="checkbox" checked>`, `<table bordered striped compact>`, `<caption>Table</caption>`, `<th colspan="2" rowspan="3" align="right" valign="top">head</th>`, `<blockquote expandable>hidden</blockquote>`, `<cite>author</cite>`, `<img src="tg://photo?id=1" tg-spoiler/>`, `<audio src="tg://audio?id=2">`, `<figcaption>photo<cite>credit</cite></figcaption>`} {
			if !strings.Contains(text, want) {
				t.Errorf("markdown=%v: missing %q in %s", markdown, want, text)
			}
		}
		if strings.Contains(text, "<td></td>") {
			t.Fatal("rendered invisible spanning cell")
		}
		if len(media) != 2 || media[0].Id != "1" || media[1].Id != "2" {
			t.Fatalf("media = %#v", media)
		}
		if got := media[0].Media.(*InputMessagePhoto).Photo.Photo.(*InputFileId).Id; got != 10 {
			t.Fatalf("wrong photo size: %d", got)
		}
		if a := media[1].Media.(*InputMessageAudio).Audio; a.Audio.(*InputFileId).Id != 20 || a.Title != "song" {
			t.Fatal("audio lost")
		}
		if !reflect.DeepEqual(before, rich.Blocks) {
			t.Fatal("render mutated blocks")
		}
	}
}

func TestOriginalRichFormatting(t *testing.T) {
	formatted := &FormattedText{Text: "A😀B", Entities: []TextEntity{{Offset: 1, Length: 2, Type: TextEntityTypeBold{}}}}
	for _, content := range []MessageContent{MessageText{Text: formatted}, &MessageText{Text: formatted}, &MessagePhoto{Caption: formatted}, &MessageVideo{Caption: formatted}, &MessageDocument{Caption: formatted}, &MessageAudio{Caption: formatted}, &MessageAnimation{Caption: formatted}, &MessageVoiceNote{Caption: formatted}, MessagePaidMedia{Caption: formatted}} {
		m := &Message{Content: content}
		if got, media := m.OriginalRichHTML(); got != "<p>A<b>😀</b>B</p>\n" || len(media) != 0 {
			t.Errorf("%T: %q %v", content, got, media)
		}
		if got, _ := m.OriginalRichMD(); got != "A<b>😀</b>B\n\n" {
			t.Errorf("%T: %q", content, got)
		}
	}
	for _, m := range []*Message{nil, {}, {Content: (*MessageRichMessage)(nil)}, {Content: &MessageRichMessage{}}} {
		if _, err := m.GetRichMessage(); !errors.Is(err, ErrNoRichMessage) {
			t.Fatal(err)
		}
		if got, media := m.OriginalRichHTML(); got != "" || len(media) != 0 {
			t.Fatal("nil content rendered")
		}
	}
}

func TestRenderedRichMessageCanBeResent(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	rich := &RichMessage{IsFull: true, IsRtl: true, Blocks: []PageBlock{PageBlockParagraph{Text: RichTextBold{Text: RichTextPlain{Text: "hi"}}}, PageBlockDocument{Document: &Document{Document: &File{Id: 123}}}}}
	m := &Message{Id: 42, ChatId: 7, Content: &MessageRichMessage{Message: rich}}
	for _, mode := range []string{"html", "markdown"} {
		text, media := m.OriginalRichHTML()
		var err error
		if mode == "html" {
			_, err = m.ReplyRichMessageHTML(c, text, &SendRichTextOpts{Media: media, IsRtl: rich.IsRtl})
		} else {
			text, media = m.OriginalRichMD()
			_, err = m.ReplyRichMessageMarkdown(c, text, &SendRichTextOpts{Media: media, IsRtl: rich.IsRtl})
		}
		if err != nil {
			t.Fatal(err)
		}
		req := (*requests)[len(*requests)-1]
		content := helperRequestContent(t, req).(*InputMessageRichMessage).Message
		var gotText string
		var gotMedia []InputRichMessageMedia
		switch source := content.Source.(type) {
		case *RichMessageSourceHtml:
			gotText, gotMedia = source.Text, source.Media
		case *RichMessageSourceMarkdown:
			gotText, gotMedia = source.Text, source.Media
		default:
			t.Fatalf("source = %T", source)
		}
		if gotText != text || !reflect.DeepEqual(gotMedia, media) || !content.IsRtl || helperReplyId(t, req) != 42 {
			t.Fatal("lost source/media/reply/direction")
		}
	}
	if len(*requests) != 2 {
		t.Fatal("rich source went through ordinary text parser")
	}
}

func TestFormattedTextRichBlockEntities(t *testing.T) {
	f := &FormattedText{Text: "a\nbcd", Entities: []TextEntity{{Length: 1, Type: TextEntityTypePreCode{Language: "go"}}, {Offset: 2, Length: 1, Type: TextEntityTypeBlockQuote{}}, {Offset: 3, Length: 1, Type: TextEntityTypeExpandableBlockQuote{}}, {Offset: 4, Length: 1, Type: TextEntityTypeItalic{}}}}
	text, _ := (&Message{Content: &MessageText{Text: f}}).OriginalRichHTML()
	for _, want := range []string{`<pre><code class="language-go">a</code></pre>`, "<blockquote>\n<p>b</p>\n</blockquote>", `<blockquote expandable>c</blockquote>`, `<p><i>d</i></p>`} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q in %q", want, text)
		}
	}
}

func TestRichConversionKeepsInlineStylesAroundBlocks(t *testing.T) {
	f := &FormattedText{Text: "abcde", Entities: []TextEntity{
		{Length: 5, Type: TextEntityTypeBold{}},
		{Offset: 1, Length: 3, Type: TextEntityTypeBlockQuote{}},
		{Offset: 2, Length: 1, Type: TextEntityTypeItalic{}},
	}}
	text, _ := formattedTextRichMessage(f).HTML()
	want := "<p><b>a</b></p>\n<blockquote>\n<p><b>b<i>c</i>d</b></p>\n</blockquote>\n<p><b>e</b></p>\n"
	if text != want {
		t.Fatalf("got %q, want %q", text, want)
	}
}

func TestRichMarkdownDelimiterBoundaries(t *testing.T) {
	for _, tc := range []struct {
		text RichText
		want string
	}{
		{RichTexts{Texts: []RichText{RichTextPlain{Text: "a"}, RichTextItalic{Text: RichTextPlain{Text: "b"}}, RichTextPlain{Text: "c"}}}, "a<i>b</i>c"},
		{RichTextBold{Text: RichTextPlain{Text: " hi "}}, "<b> hi </b>"},
		{RichTexts{Texts: []RichText{RichTextPlain{Text: "a "}, RichTextBold{Text: RichTextPlain{Text: "b"}}, RichTextPlain{Text: " c"}}}, "a **b** c"},
		{RichTexts{Texts: []RichText{RichTextBold{Text: RichTextPlain{Text: "*"}}, RichTextItalic{Text: RichTextPlain{Text: "_"}}}}, `<b>\*</b><i>\_</i>`},
	} {
		if got := RichTextMarkdown(tc.text); got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}
}

func TestRichMediaTypesAndMissingFiles(t *testing.T) {
	file := &File{Id: 123}
	for _, tc := range []struct {
		block     PageBlock
		kind, typ string
	}{
		{PageBlockPhoto{Photo: &Photo{Sizes: []PhotoSize{{Photo: file}}}}, "photo", "inputMessagePhoto"},
		{PageBlockAnimation{Animation: &Animation{Animation: file}}, "video", "inputMessageAnimation"},
		{PageBlockVideo{Video: &Video{Video: file}}, "video", "inputMessageVideo"},
		{PageBlockAudio{Audio: &Audio{Audio: file}}, "audio", "inputMessageAudio"},
		{PageBlockVoiceNote{VoiceNote: &VoiceNote{Voice: file}}, "audio", "inputMessageVoiceNote"},
		{PageBlockDocument{Document: &Document{Document: file}}, "document", "inputMessageDocument"},
	} {
		text, media := RichBlockHTML(tc.block)
		if !strings.Contains(text, "tg://"+tc.kind+"?id=1") || len(media) != 1 || media[0].Media.GetType() != tc.typ {
			t.Errorf("%T: %s, %#v", tc.block, text, media)
		}
	}
	for _, block := range []PageBlock{nil, (*PageBlockParagraph)(nil), PageBlockUnsupported{}, PageBlockPhoto{}, PageBlockPhoto{Photo: &Photo{}}, PageBlockPhoto{Photo: &Photo{Sizes: []PhotoSize{{Photo: nil}}}}, PageBlockAnimation{}, PageBlockVideo{}, PageBlockAudio{}, PageBlockVoiceNote{}, PageBlockDocument{Document: &Document{}}, PageBlockMap{}} {
		for _, render := range []func(PageBlock) (string, []InputRichMessageMedia){RichBlockHTML, RichBlockMarkdown} {
			if text, media := render(block); text != "" || len(media) != 0 {
				t.Errorf("%T: %q, %#v", block, text, media)
			}
		}
	}
}

func TestRichButtonsAndCaptions(t *testing.T) {
	text, media := RichBlockHTML(PageBlockButtonRow{Align: &PageBlockHorizontalAlignmentCenter{}, Buttons: []InlineButton{
		{Style: ButtonStylePrimary{}, Text: RichTextPlain{Text: "<&"}, Type: &InlineKeyboardButtonTypeCallback{Data: []byte(`"&`)}},
		{Text: RichTextPlain{Text: "current"}, Type: InlineKeyboardButtonTypeSwitchInline{Query: "a&b", TargetChat: &TargetChatCurrent{}}},
		{Text: RichTextPlain{Text: "chosen"}, Type: InlineKeyboardButtonTypeSwitchInline{TargetChat: TargetChatChosen{Types: &TargetChatTypes{AllowUserChats: true}}}},
	}})
	for _, want := range []string{`<tg-button-row align="center">`, `<tg-button type="callback_data" data="&#34;&amp;" style="primary">&lt;&amp;</tg-button>`, `type="switch_inline_query_current_chat" query="a&amp;b"`, `type="switch_inline_query_chosen_chat" allow-user-chats query=""`} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q in %q", want, text)
		}
	}
	if len(media) != 0 {
		t.Fatal("button produced media")
	}
}

func TestRichMarkdownEntities(t *testing.T) {
	for _, tc := range []struct {
		text, want string
	}{
		{"I don't know what this is but apparently it's also math 👇", "I don&#39;t know what this is but apparently it&#39;s also math 👇"},
		{`It's "quoted"`, `It&#39;s &#34;quoted&#34;`},
		{`&#39; &#x27; &apos;`, `&amp;\#39; &amp;\#x27; &amp;apos;`},
		{"# * [literal] $math$", `\# \* \[literal\] \$math\$`},
		{"<script> & 'hello'\nnext", "&lt;script&gt; &amp; &#39;hello&#39;<br/>next"},
	} {
		t.Run(tc.text, func(t *testing.T) {
			plain := RichTextPlain{Text: tc.text}
			if got := RichTextMarkdown(plain); got != tc.want {
				t.Fatalf("Markdown = %q; want %q", got, tc.want)
			}
			if got := RichTextMarkdown(RichTextBold{Text: plain}); got != "**"+tc.want+"**" {
				t.Fatalf("bold Markdown = %q", got)
			}
			if got := RichTextMarkdown(RichTextUnderline{Text: plain}); got != "<u>"+tc.want+"</u>" {
				t.Fatalf("HTML fallback = %q", got)
			}
			if got := RichTextContent(plain); got != tc.text {
				t.Fatalf("plain text changed: %q", got)
			}
		})
	}
}

func TestRichMessageApostrophesInReply(t *testing.T) {
	c := newTestClient(t)
	requests := captureHelperRequests(t, c)
	plain := "I don't know what this is but apparently it's also math 👇"
	escaped := "I don&#39;t know what this is but apparently it&#39;s also math 👇"
	message := &Message{ChatId: 10, Id: 20, Content: &MessageRichMessage{Message: &RichMessage{
		IsFull: true, Blocks: []PageBlock{&PageBlockParagraph{Text: &RichTextPlain{Text: plain}}},
	}}}
	html, _ := message.OriginalRichHTML()
	if html != "<p>"+escaped+"</p>\n" {
		t.Fatalf("HTML = %q", html)
	}
	markdown, media := message.OriginalRichMD()
	if markdown != escaped+"\n\n" {
		t.Fatalf("Markdown = %q", markdown)
	}
	if _, err := message.ReplyRichMessageMarkdown(c, markdown, &SendRichTextOpts{Media: media}); err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 1 {
		t.Fatalf("sent %d requests", len(*requests))
	}
	request := (*requests)[0]
	content := helperRequestContent(t, request).(*InputMessageRichMessage)
	source := content.Message.Source.(*RichMessageSourceMarkdown)
	if source.Text != escaped+"\n\n" || helperReplyId(t, request) != 20 {
		t.Fatalf("reply changed source or destination: %#v", source)
	}
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
