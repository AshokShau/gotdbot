package gotdbot

import (
	"fmt"
	"html"
	"reflect"
	"sort"
	"strings"
	"unicode/utf16"
)

// Normalize TDLib values, pointers and typed nils.
func formattingValue[T any](value T) T {
	v := reflect.ValueOf(value)
	if v.IsValid() && v.Kind() == reflect.Pointer {
		if v.IsNil() {
			var zero T
			return zero
		}
		return v.Elem().Interface().(T)
	}
	return value
}

func buildFormatString(format DateTimeFormattingType) string {
	switch f := formattingValue(format).(type) {
	case DateTimeFormattingTypeRelative:
		return "r"
	case DateTimeFormattingTypeAbsolute:
		result := ""
		if f.ShowDayOfWeek {
			result += "w"
		}
		switch formattingValue(f.DatePrecision).(type) {
		case DateTimePartPrecisionShort:
			result += "d"
		case DateTimePartPrecisionLong:
			result += "D"
		}
		switch formattingValue(f.TimePrecision).(type) {
		case DateTimePartPrecisionShort:
			result += "t"
		case DateTimePartPrecisionLong:
			result += "T"
		}
		return result
	}
	return ""
}

// escapeText escapes special characters for HTML and MDV2.
func escapeText(text, mode string) string {
	if mode == ParseModeHTML {
		text = strings.ReplaceAll(text, "&", "&amp;")
		text = strings.ReplaceAll(text, "<", "&lt;")
		text = strings.ReplaceAll(text, ">", "&gt;")
		return text
	} else if mode == ParseModeMarkdownV2 {
		specials := []string{"\\", "_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == ParseModeMarkdownV2+"_url" {
		specials := []string{"\\", ")"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == ParseModeMarkdown+"_url" {
		return strings.NewReplacer("\\", "\\\\", ")", "\\)").Replace(text)
	} else if mode == ParseModeMarkdownV2+"_code" {
		specials := []string{"\\", "`"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == ParseModeMarkdown+"_code" {
		specials := []string{"\\", "`"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == ParseModeMarkdown {
		specials := []string{"\\", "*", "_", "`", "["}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	}
	return text
}

type formatEntity struct {
	entity     TextEntity
	start, end int
	children   []*formatEntity
}

// entityTree sorts entities and skips invalid UTF-16 ranges.
func entityTree(text string, entities []TextEntity) ([]uint16, []*formatEntity) {
	units := utf16.Encode([]rune(text))
	validBoundary := func(i int) bool {
		return i >= 0 && i <= len(units) && (i == 0 || i == len(units) ||
			!(units[i-1] >= 0xd800 && units[i-1] <= 0xdbff && units[i] >= 0xdc00 && units[i] <= 0xdfff))
	}
	nodes := make([]*formatEntity, 0, len(entities))
	for _, entity := range entities {
		start, end := int(entity.Offset), int(int64(entity.Offset)+int64(entity.Length))
		entity.Type = formattingValue(entity.Type)
		if entity.Length <= 0 || entity.Type == nil || !validBoundary(start) || !validBoundary(end) {
			continue
		}
		nodes = append(nodes, &formatEntity{entity: entity, start: start, end: end})
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		if nodes[i].start != nodes[j].start {
			return nodes[i].start < nodes[j].start
		}
		return nodes[i].end > nodes[j].end
	})
	var roots, stack []*formatEntity
	for _, node := range nodes {
		for len(stack) > 0 && node.start >= stack[len(stack)-1].end {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			roots = append(roots, node)
		} else {
			parent := stack[len(stack)-1]
			if node.end > parent.end {
				continue
			}
			parent.children = append(parent.children, node)
		}
		stack = append(stack, node)
	}
	return units, roots
}

// Format renders Telegram HTML, Markdown or MarkdownV2; unknown modes return plain text.
func (f *FormattedText) Format(parseMode string) string {
	if f == nil {
		return ""
	}
	mode := strings.ToLower(parseMode)
	if mode != ParseModeHTML && mode != ParseModeMarkdown && mode != ParseModeMarkdownV2 {
		return f.Text
	}
	units, roots := entityTree(f.Text, f.Entities)
	var sequence func(int, int, []*formatEntity, bool) string
	sequence = func(start, end int, nodes []*formatEntity, nested bool) string {
		var out strings.Builder
		write := func(s string) {
			// Separate italic and underline delimiters with an ignored carriage return.
			if mode == ParseModeMarkdownV2 && strings.HasSuffix(out.String(), "_") && strings.HasPrefix(s, "_") {
				out.WriteByte('\r')
			}
			out.WriteString(s)
		}
		pos := start
		for _, node := range nodes {
			write(escapeText(string(utf16.Decode(units[pos:node.start])), mode))
			typ := node.entity.Type
			open, close := entityTags(typ, mode)
			if mode == ParseModeMarkdown && nested {
				open, close = "", ""
			}
			var content string
			switch typ.(type) {
			case TextEntityTypeCode, TextEntityTypePre, TextEntityTypePreCode:
				codeMode := mode
				if mode != ParseModeHTML && open != "" {
					codeMode += "_code"
				}
				content = escapeText(string(utf16.Decode(units[node.start:node.end])), codeMode)
			default:
				content = sequence(node.start, node.end, node.children, nested || open != "")
			}
			if mode == ParseModeMarkdownV2 {
				switch typ.(type) {
				case TextEntityTypeBlockQuote, TextEntityTypeExpandableBlockQuote:
					content = strings.ReplaceAll(content, "\n", "\n>")
				}
			}
			write(open)
			write(content)
			write(close)
			pos = node.end
		}
		write(escapeText(string(utf16.Decode(units[pos:end])), mode))
		return out.String()
	}
	return sequence(0, len(units), roots, false)
}

func entityTags(typ TextEntityType, mode string) (string, string) {
	wrap := func(tag, delimiter string) (string, string) {
		if mode == ParseModeHTML {
			return "<" + tag + ">", "</" + tag + ">"
		}
		return delimiter, delimiter
	}
	link := func(url string) (string, string) {
		if mode == ParseModeHTML {
			return `<a href="` + html.EscapeString(url) + `">`, "</a>"
		}
		return "[", "](" + escapeText(url, mode+"_url") + ")"
	}
	switch e := typ.(type) {
	case TextEntityTypeBold:
		return wrap("b", "*")
	case TextEntityTypeItalic:
		return wrap("i", "_")
	case TextEntityTypeCode:
		return wrap("code", "`")
	case TextEntityTypeUnderline:
		if mode != ParseModeMarkdown {
			return wrap("u", "__")
		}
	case TextEntityTypeStrikethrough:
		if mode != ParseModeMarkdown {
			return wrap("s", "~")
		}
	case TextEntityTypeSpoiler:
		if mode != ParseModeMarkdown {
			return wrap("tg-spoiler", "||")
		}
	case TextEntityTypePre:
		if mode == ParseModeHTML {
			return "<pre>", "</pre>"
		}
		return "```\n", "```"
	case TextEntityTypePreCode:
		if mode == ParseModeHTML {
			return `<pre><code class="language-` + html.EscapeString(e.Language) + `">`, "</code></pre>"
		}
		return "```" + strings.NewReplacer("\n", "", "\r", "", "`", "").Replace(e.Language) + "\n", "```"
	case TextEntityTypeTextUrl:
		return link(e.Url)
	case TextEntityTypeMentionName:
		return link(fmt.Sprintf("tg://user?id=%d", e.UserId))
	case TextEntityTypeCustomEmoji:
		if mode == ParseModeHTML {
			return fmt.Sprintf(`<tg-emoji emoji-id="%d">`, e.CustomEmojiId), "</tg-emoji>"
		}
		if mode == ParseModeMarkdownV2 {
			return "![", fmt.Sprintf("](tg://emoji?id=%d)", e.CustomEmojiId)
		}
	case TextEntityTypeDateTime:
		format := buildFormatString(e.FormattingType)
		if mode == ParseModeHTML {
			attr := ""
			if format != "" {
				attr = ` format="` + format + `"`
			}
			return fmt.Sprintf(`<tg-time unix="%d"%s>`, e.UnixTime, attr), "</tg-time>"
		}
		if mode == ParseModeMarkdownV2 {
			if format != "" {
				format = "&format=" + format
			}
			return "![", fmt.Sprintf("](tg://time?unix=%d%s)", e.UnixTime, format)
		}
	case TextEntityTypeBlockQuote:
		if mode == ParseModeHTML {
			return "<blockquote>", "</blockquote>"
		}
		if mode == ParseModeMarkdownV2 {
			return ">", ""
		}
	case TextEntityTypeExpandableBlockQuote:
		if mode == ParseModeHTML {
			return "<blockquote expandable>", "</blockquote>"
		}
		if mode == ParseModeMarkdownV2 {
			return "**>", "||"
		}
	}
	return "", ""
}
