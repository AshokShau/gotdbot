package gotdbot

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf16"
)

// getText extracts the *FormattedText from the message content if it has text.
func (m *Message) getText() (*FormattedText, error) {
	if m == nil || m.Content == nil {
		return nil, ErrNoFormattedText
	}
	switch content := m.Content.(type) {
	case *MessageText:
		return content.Text, nil
	default:
		return nil, ErrNoFormattedText
	}
}

// getCaption extracts the *FormattedText from the message content if it has a caption.
func (m *Message) getCaption() (*FormattedText, error) {
	if m == nil || m.Content == nil {
		return nil, ErrNoFormattedText
	}
	switch content := m.Content.(type) {
	case *MessageAnimation:
		return content.Caption, nil
	case *MessageAudio:
		return content.Caption, nil
	case *MessageDocument:
		return content.Caption, nil
	case *MessagePhoto:
		return content.Caption, nil
	case *MessageVideo:
		return content.Caption, nil
	case *MessageVoiceNote:
		return content.Caption, nil
	default:
		return nil, ErrNoFormattedText
	}
}

// OriginalMD gets the original Markdown formatting of a message text.
func (m *Message) OriginalMD() string {
	text, err := m.getText()
	if err != nil {
		return ""
	}
	return UnparseEntities(text.Text, text.Entities, "markdown")
}

// OriginalMDV2 gets the original markdownV2 formatting of a message text.
func (m *Message) OriginalMDV2() string {
	text, err := m.getText()
	if err != nil {
		return ""
	}
	return UnparseEntities(text.Text, text.Entities, "markdownv2")
}

// OriginalHTML gets the original HTML formatting of a message text.
func (m *Message) OriginalHTML() string {
	text, err := m.getText()
	if err != nil {
		return ""
	}
	return UnparseEntities(text.Text, text.Entities, "html")
}

// OriginalCaptionMD gets the original markdown formatting of a message caption.
func (m *Message) OriginalCaptionMD() string {
	caption, err := m.getCaption()
	if err != nil {
		return ""
	}
	return UnparseEntities(caption.Text, caption.Entities, "markdown")
}

// OriginalCaptionMDV2 gets the original markdownV2 formatting of a message caption.
func (m *Message) OriginalCaptionMDV2() string {
	caption, err := m.getCaption()
	if err != nil {
		return ""
	}
	return UnparseEntities(caption.Text, caption.Entities, "markdownv2")
}

// OriginalCaptionHTML gets the original HTML formatting of a message caption.
func (m *Message) OriginalCaptionHTML() string {
	caption, err := m.getCaption()
	if err != nil {
		return ""
	}
	return UnparseEntities(caption.Text, caption.Entities, "html")
}

// entityEvent represents the start or end of an entity
type entityEvent struct {
	Index   int
	IsStart bool
	Entity  *TextEntity
	Level   int
}

// escapeText escapes special characters for HTML and MDV2.
func escapeText(text, mode string) string {
	if mode == "html" {
		text = strings.ReplaceAll(text, "&", "&amp;")
		text = strings.ReplaceAll(text, "<", "&lt;")
		text = strings.ReplaceAll(text, ">", "&gt;")
		return text
	} else if mode == "markdownv2" {
		specials := []string{"\\", "_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == "markdownv2_url" {
		specials := []string{"\\", ")"}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	} else if mode == "markdown" {
		specials := []string{"*", "_", "`", "["}
		for _, s := range specials {
			text = strings.ReplaceAll(text, s, "\\"+s)
		}
		return text
	}
	return text
}

// UnparseEntities converts plain text with entities into a formatted string (Markdown, MarkdownV2, HTML).
func UnparseEntities(text string, entities []TextEntity, mode string) string {
	if len(entities) == 0 {
		return escapeText(text, mode)
	}

	utf16Str := utf16.Encode([]rune(text))
	var events []entityEvent

	for i, ent := range entities {
		e := ent
		events = append(events, entityEvent{Index: int(e.Offset), IsStart: true, Entity: &e, Level: i})
		events = append(events, entityEvent{Index: int(e.Offset + e.Length), IsStart: false, Entity: &e, Level: i})
	}

	sort.SliceStable(events, func(i, j int) bool {
		if events[i].Index != events[j].Index {
			return events[i].Index < events[j].Index
		}
		if events[i].IsStart != events[j].IsStart {
			return !events[i].IsStart
		}

		if events[i].IsStart {
			return events[i].Level < events[j].Level
		}

		return events[i].Level > events[j].Level
	})

	var result strings.Builder
	lastIndex := 0

	for _, ev := range events {
		if ev.Index > lastIndex {
			if lastIndex < len(utf16Str) {
				end := ev.Index
				if end > len(utf16Str) {
					end = len(utf16Str)
				}
				chunk := string(utf16.Decode(utf16Str[lastIndex:end]))
				result.WriteString(escapeText(chunk, mode))
			}
			lastIndex = ev.Index
		}

		tag := getTag(ev.Entity, ev.IsStart, mode)
		result.WriteString(tag)
	}

	if lastIndex < len(utf16Str) {
		chunk := string(utf16.Decode(utf16Str[lastIndex:]))
		result.WriteString(escapeText(chunk, mode))
	}

	if mode == "markdownv2" || mode == "markdown" {
		res := result.String()
		res = strings.ReplaceAll(res, "_\r", "_")
		res = strings.ReplaceAll(res, "\r_", "_")
		res = strings.ReplaceAll(res, "__\r", "__")
		res = strings.ReplaceAll(res, "\r__", "__")
		return res
	}

	return result.String()
}

func getTag(ent *TextEntity, isStart bool, mode string) string {
	switch e := ent.TypeField.(type) {
	case *TextEntityTypeBold, TextEntityTypeBold:
		if mode == "html" {
			if isStart {
				return "<b>"
			} else {
				return "</b>"
			}
		} else {
			return "*"
		}
	case *TextEntityTypeItalic, TextEntityTypeItalic:
		if mode == "html" {
			if isStart {
				return "<i>"
			} else {
				return "</i>"
			}
		} else {
			if isStart {
				return "_\r"
			} else {
				return "\r_"
			}
		}
	case *TextEntityTypeUnderline, TextEntityTypeUnderline:
		if mode == "html" {
			if isStart {
				return "<u>"
			} else {
				return "</u>"
			}
		} else if mode == "markdownv2" {
			if isStart {
				return "__\r"
			} else {
				return "\r__"
			}
		}
	case *TextEntityTypeStrikethrough, TextEntityTypeStrikethrough:
		if mode == "html" {
			if isStart {
				return "<s>"
			} else {
				return "</s>"
			}
		} else if mode == "markdownv2" {
			return "~"
		}
	case *TextEntityTypeSpoiler, TextEntityTypeSpoiler:
		if mode == "html" {
			if isStart {
				return "<tg-spoiler>"
			} else {
				return "</tg-spoiler>"
			}
		} else if mode == "markdownv2" {
			return "||"
		}
	case *TextEntityTypeCode, TextEntityTypeCode:
		if mode == "html" {
			if isStart {
				return "<code>"
			} else {
				return "</code>"
			}
		} else {
			return "`"
		}
	case *TextEntityTypePre, TextEntityTypePre:
		if mode == "html" {
			if isStart {
				return "<pre>"
			} else {
				return "</pre>"
			}
		} else {
			return "```\n"
		}
	case *TextEntityTypePreCode:
		lang := e.Language
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<pre><code class=\"language-%s\">", lang)
			} else {
				return "</code></pre>"
			}
		} else {
			if isStart {
				return fmt.Sprintf("```%s\n", lang)
			} else {
				return "```\n"
			}
		}
	case TextEntityTypePreCode:
		lang := e.Language
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<pre><code class=\"language-%s\">", lang)
			} else {
				return "</code></pre>"
			}
		} else {
			if isStart {
				return fmt.Sprintf("```%s\n", lang)
			} else {
				return "```\n"
			}
		}
	case *TextEntityTypeTextUrl:
		url := e.Url
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<a href=\"%s\">", url)
			} else {
				return "</a>"
			}
		} else {
			if isStart {
				return "["
			} else {
				return fmt.Sprintf("](%s)", escapeText(url, mode+"_url"))
			}
		}
	case TextEntityTypeTextUrl:
		url := e.Url
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<a href=\"%s\">", url)
			} else {
				return "</a>"
			}
		} else {
			if isStart {
				return "["
			} else {
				return fmt.Sprintf("](%s)", escapeText(url, mode+"_url"))
			}
		}
	case *TextEntityTypeMentionName:
		userId := e.UserId
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<a href=\"tg://user?id=%d\">", userId)
			} else {
				return "</a>"
			}
		} else {
			if isStart {
				return "["
			} else {
				return fmt.Sprintf("](tg://user?id=%d)", userId)
			}
		}
	case TextEntityTypeMentionName:
		userId := e.UserId
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<a href=\"tg://user?id=%d\">", userId)
			} else {
				return "</a>"
			}
		} else {
			if isStart {
				return "["
			} else {
				return fmt.Sprintf("](tg://user?id=%d)", userId)
			}
		}
	case *TextEntityTypeCustomEmoji:
		emojiId := e.CustomEmojiId
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<tg-emoji emoji-id=\"%d\">", emojiId)
			} else {
				return "</tg-emoji>"
			}
		} else if mode == "markdownv2" {
			if isStart {
				return "!["
			} else {
				return fmt.Sprintf("](tg://emoji?id=%d)", emojiId)
			}
		}
	case TextEntityTypeCustomEmoji:
		emojiId := e.CustomEmojiId
		if mode == "html" {
			if isStart {
				return fmt.Sprintf("<tg-emoji emoji-id=\"%d\">", emojiId)
			} else {
				return "</tg-emoji>"
			}
		} else if mode == "markdownv2" {
			if isStart {
				return "!["
			} else {
				return fmt.Sprintf("](tg://emoji?id=%d)", emojiId)
			}
		}
	case *TextEntityTypeBlockQuote, TextEntityTypeBlockQuote:
		if mode == "html" {
			if isStart {
				return "<blockquote>"
			} else {
				return "</blockquote>"
			}
		} else if mode == "markdownv2" {
			// This might need line-by-line escaping, but we'll use block tags
			if isStart {
				return "**>"
			} else {
				return ""
			}
		}
	case *TextEntityTypeExpandableBlockQuote, TextEntityTypeExpandableBlockQuote:
		if mode == "html" {
			if isStart {
				return "<blockquote expandable>"
			} else {
				return "</blockquote>"
			}
		} else if mode == "markdownv2" {
			if isStart {
				return "**>"
			} else {
				return ""
			}
		}

	case *TextEntityTypeUrl, TextEntityTypeUrl,
		*TextEntityTypeEmailAddress, TextEntityTypeEmailAddress,
		*TextEntityTypePhoneNumber, TextEntityTypePhoneNumber,
		*TextEntityTypeHashtag, TextEntityTypeHashtag,
		*TextEntityTypeCashtag, TextEntityTypeCashtag,
		*TextEntityTypeBankCardNumber, TextEntityTypeBankCardNumber,
		*TextEntityTypeBotCommand, TextEntityTypeBotCommand,
		*TextEntityTypeMention, TextEntityTypeMention,
		*TextEntityTypeDateTime, TextEntityTypeDateTime,
		*TextEntityTypeMediaTimestamp, TextEntityTypeMediaTimestamp:
		return ""
	}
	return ""
}
