package gotdbot

import (
	"fmt"
	"html"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// RichTextHTML renders Telegram rich HTML.
func RichTextHTML(text RichText) string { return renderRichText(text, ParseModeHTML) }

// RichTextMarkdown renders rich Markdown with HTML fallbacks.
func RichTextMarkdown(text RichText) string { return renderRichText(text, ParseModeMarkdown) }

// RichTextContent returns the visible text, without formatting.
func RichTextContent(text RichText) string { return renderRichText(text, ParseModeNone) }

var richMarkdownEscaper = strings.NewReplacer(
	"\\", "\\\\", "`", "\\`", "*", "\\*", "_", "\\_", "[", "\\[", "]", "\\]",
	"(", "\\(", ")", "\\)", "#", "\\#", "+", "\\+", "-", "\\-", "!", "\\!",
	"|", "\\|", "~", "\\~", "=", "\\=", "$", "\\$", ".", "\\.", "{", "\\{", "}", "\\}",
)

func richEscape(text, mode string) string {
	if mode == ParseModeNone {
		return text
	}
	if strings.HasPrefix(mode, ParseModeMarkdown) {
		text = richMarkdownEscaper.Replace(text)
	}
	text = html.EscapeString(text)
	return strings.ReplaceAll(text, "\n", "<br/>")
}

func renderRichText(text RichText, mode string) string {
	wrap := func(tag, delimiter string, child RichText) string {
		if mode == ParseModeNone {
			return RichTextContent(child)
		}
		if mode == ParseModeMarkdown && delimiter != "" {
			content := RichTextContent(child)
			if content != "" && content == strings.TrimSpace(content) {
				return delimiter + renderRichText(child, mode) + delimiter
			}
			// CommonMark delimiters can't enclose leading/trailing spaces.
			return "<" + tag + ">" + renderRichText(child, ParseModeMarkdown+"_html") + "</" + tag + ">"
		}
		return "<" + tag + ">" + renderRichText(child, mode) + "</" + tag + ">"
	}
	link := func(url string, child RichText) string {
		if mode == ParseModeNone {
			return RichTextContent(child)
		}
		return `<a href="` + html.EscapeString(url) + `">` + renderRichText(child, mode) + `</a>`
	}
	switch v := formattingValue(text).(type) {
	case RichTextPlain:
		return richEscape(v.Text, mode)
	case RichTexts:
		var out strings.Builder
		for i, child := range v.Texts {
			childMode := mode
			if mode == ParseModeMarkdown && richMarkdownStyle(child) {
				// Avoid ambiguous Markdown delimiters between adjacent text runs.
				for j := i - 1; j >= 0; j-- {
					if content := RichTextContent(v.Texts[j]); content != "" {
						last, _ := utf8.DecodeLastRuneInString(content)
						if !unicode.IsSpace(last) {
							childMode = ParseModeMarkdown + "_html"
						}
						break
					}
				}
				for j := i + 1; j < len(v.Texts); j++ {
					if content := RichTextContent(v.Texts[j]); content != "" {
						first, _ := utf8.DecodeRuneInString(content)
						if !unicode.IsSpace(first) {
							childMode = ParseModeMarkdown + "_html"
						}
						break
					}
				}
			}
			out.WriteString(renderRichText(child, childMode))
		}
		return out.String()
	case RichTextBold:
		return wrap("b", "**", v.Text)
	case RichTextItalic:
		return wrap("i", "_", v.Text)
	case RichTextUnderline:
		return wrap("u", "", v.Text)
	case RichTextStrikethrough:
		return wrap("s", "~~", v.Text)
	case RichTextSpoiler:
		return wrap("tg-spoiler", "||", v.Text)
	case RichTextMarked:
		return wrap("mark", "==", v.Text)
	case RichTextSubscript:
		return wrap("sub", "", v.Text)
	case RichTextSuperscript:
		return wrap("sup", "", v.Text)
	case RichTextFixed:
		if mode == ParseModeNone {
			return RichTextContent(v.Text)
		}
		return "<code>" + html.EscapeString(RichTextContent(v.Text)) + "</code>"
	case RichTextUrl:
		return link(v.Url, v.Text)
	case RichTextEmailAddress:
		return link("mailto:"+v.EmailAddress, v.Text)
	case RichTextPhoneNumber:
		return link("tel:"+v.PhoneNumber, v.Text)
	case RichTextMentionName:
		return link(fmt.Sprintf("tg://user?id=%d", v.UserId), v.Text)
	case RichTextAnchorLink:
		return link("#"+v.AnchorName, v.Text)
	case RichTextReferenceLink:
		return link("#"+v.ReferenceName, v.Text)
	case RichTextReference:
		if mode == ParseModeNone {
			return RichTextContent(v.Text)
		}
		return `<tg-reference name="` + html.EscapeString(v.Name) + `">` + renderRichText(v.Text, mode) + `</tg-reference>`
	case RichTextAnchor:
		if mode == ParseModeNone {
			return ""
		}
		return `<a name="` + html.EscapeString(v.Name) + `"></a>`
	case RichTextCustomEmoji:
		if mode == ParseModeNone {
			return v.AlternativeText
		}
		return fmt.Sprintf(`<tg-emoji emoji-id="%d">%s</tg-emoji>`, v.CustomEmojiId, richEscape(v.AlternativeText, mode))
	case RichTextDateTime:
		if mode == ParseModeNone {
			return RichTextContent(v.Text)
		}
		open, close := entityTags(TextEntityTypeDateTime{UnixTime: v.UnixTime, FormattingType: v.FormattingType}, ParseModeHTML)
		return open + renderRichText(v.Text, mode) + close
	case RichTextMathematicalExpression:
		if mode == ParseModeNone {
			return v.Expression
		}
		return "<tg-math>" + html.EscapeString(v.Expression) + "</tg-math>"
	case RichTextMention:
		return renderRichText(v.Text, mode)
	case RichTextHashtag:
		return renderRichText(v.Text, mode)
	case RichTextCashtag:
		return renderRichText(v.Text, mode)
	case RichTextBotCommand:
		return renderRichText(v.Text, mode)
	case RichTextBankCardNumber:
		return renderRichText(v.Text, mode)
	case RichTextTonAddress:
		return renderRichText(v.Text, mode)
	case RichTextDiff:
		return renderRichText(v.Text, mode)
	case RichTextIcon:
		if v.Document != nil {
			return richEscape(v.Document.FileName, mode)
		}
	case RichTextButton:
		if v.Button == nil {
			return ""
		}
		if mode == ParseModeNone {
			return RichTextContent(v.Button.Text)
		}
		return richButtonHTML(v.Button)
	}
	return ""
}

func richMarkdownStyle(text RichText) bool {
	switch formattingValue(text).(type) {
	case RichTextBold, RichTextItalic, RichTextStrikethrough, RichTextSpoiler, RichTextMarked:
		return true
	}
	return false
}

// HTML renders available blocks as rich HTML with reusable media references.
func (m *RichMessage) HTML() (string, []InputRichMessageMedia) {
	return renderRichMessage(m, ParseModeHTML)
}

// Markdown renders rich Markdown with HTML fallbacks and reusable media references.
func (m *RichMessage) Markdown() (string, []InputRichMessageMedia) {
	return renderRichMessage(m, ParseModeMarkdown)
}

func renderRichMessage(m *RichMessage, mode string) (string, []InputRichMessageMedia) {
	if m == nil {
		return "", nil
	}
	r := richRenderer{}
	return r.blocks(m.Blocks, mode), r.media
}

// RichBlockHTML renders a TDLib page block and collects referenced media.
func RichBlockHTML(block PageBlock) (string, []InputRichMessageMedia) {
	r := richRenderer{}
	return r.block(block, ParseModeHTML), r.media
}

// RichBlockMarkdown renders a TDLib page block as rich Markdown and media.
func RichBlockMarkdown(block PageBlock) (string, []InputRichMessageMedia) {
	r := richRenderer{}
	return r.block(block, ParseModeMarkdown), r.media
}

func formattedTextRichMessage(f *FormattedText) *RichMessage {
	if f == nil {
		return nil
	}
	units, roots := entityTree(f.Text, f.Entities)
	plain := func(start, end int) RichText { return RichTextPlain{Text: string(utf16.Decode(units[start:end]))} }
	var sequence func(int, int, []*formatEntity) RichText
	sequence = func(start, end int, nodes []*formatEntity) RichText {
		var texts []RichText
		pos := start
		for _, node := range nodes {
			if node.end <= start || node.start >= end {
				continue
			}
			left, right := max(start, node.start), min(end, node.end)
			if left > pos {
				texts = append(texts, plain(pos, left))
			}
			inner := sequence(left, right, node.children)
			switch e := node.entity.Type.(type) {
			case TextEntityTypeBold:
				inner = RichTextBold{Text: inner}
			case TextEntityTypeItalic:
				inner = RichTextItalic{Text: inner}
			case TextEntityTypeUnderline:
				inner = RichTextUnderline{Text: inner}
			case TextEntityTypeStrikethrough:
				inner = RichTextStrikethrough{Text: inner}
			case TextEntityTypeSpoiler:
				inner = RichTextSpoiler{Text: inner}
			case TextEntityTypeCode:
				inner = RichTextFixed{Text: inner}
			case TextEntityTypeTextUrl:
				inner = RichTextUrl{Text: inner, Url: e.Url}
			case TextEntityTypeUrl:
				inner = RichTextUrl{Text: inner, Url: RichTextContent(inner)}
			case TextEntityTypeEmailAddress:
				inner = RichTextEmailAddress{Text: inner, EmailAddress: RichTextContent(inner)}
			case TextEntityTypePhoneNumber:
				inner = RichTextPhoneNumber{Text: inner, PhoneNumber: RichTextContent(inner)}
			case TextEntityTypeMentionName:
				inner = RichTextMentionName{Text: inner, UserId: e.UserId}
			case TextEntityTypeCustomEmoji:
				inner = RichTextCustomEmoji{AlternativeText: RichTextContent(inner), CustomEmojiId: e.CustomEmojiId}
			case TextEntityTypeDateTime:
				inner = RichTextDateTime{Text: inner, UnixTime: e.UnixTime, FormattingType: e.FormattingType}
			}
			texts = append(texts, inner)
			pos = right
		}
		if pos < end {
			texts = append(texts, plain(pos, end))
		}
		return RichTexts{Texts: texts}
	}
	// Clip inline styles around structural blocks to preserve surrounding formatting.
	var structural func([]*formatEntity) []*formatEntity
	structural = func(nodes []*formatEntity) []*formatEntity {
		var result []*formatEntity
		for _, node := range nodes {
			children := structural(node.children)
			switch node.entity.Type.(type) {
			case TextEntityTypePre, TextEntityTypePreCode, TextEntityTypeBlockQuote, TextEntityTypeExpandableBlockQuote:
				result = append(result, &formatEntity{entity: node.entity, start: node.start, end: node.end, children: children})
			default:
				result = append(result, children...)
			}
		}
		return result
	}
	var buildBlocks func(int, int, []*formatEntity) []PageBlock
	buildBlocks = func(start, end int, nodes []*formatEntity) []PageBlock {
		var blocks []PageBlock
		pos := start
		paragraph := func(left, right int) {
			if right > left {
				blocks = append(blocks, PageBlockParagraph{Text: sequence(left, right, roots)})
			}
		}
		for _, node := range nodes {
			paragraph(pos, node.start)
			switch e := node.entity.Type.(type) {
			case TextEntityTypePre:
				blocks = append(blocks, PageBlockPreformatted{Text: plain(node.start, node.end)})
			case TextEntityTypePreCode:
				blocks = append(blocks, PageBlockPreformatted{Text: plain(node.start, node.end), Language: e.Language})
			case TextEntityTypeBlockQuote:
				blocks = append(blocks, PageBlockBlockQuote{Blocks: buildBlocks(node.start, node.end, node.children)})
			case TextEntityTypeExpandableBlockQuote:
				blocks = append(blocks, PageBlockExpandableBlockQuote{Text: sequence(node.start, node.end, roots)})
			}
			pos = node.end
		}
		paragraph(pos, end)
		return blocks
	}
	return &RichMessage{Blocks: buildBlocks(0, len(units), structural(roots)), IsFull: true}
}
