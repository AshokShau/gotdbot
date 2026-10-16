package gotdbot

import (
	"fmt"
	"strings"
)

// ParseButtonMarkup extracts [Name](buttonurl:...) buttons, preserving code and escaped links.
func ParseButtonMarkup(text, parseMode string) (string, *ReplyMarkupInlineKeyboard, error) {
	mode := strings.ToLower(parseMode)
	if mode != ParseModeNone && mode != ParseModeHTML && mode != ParseModeMarkdown && mode != ParseModeMarkdownV2 {
		return "", nil, fmt.Errorf("unsupported text parse mode %q", parseMode)
	}

	var out strings.Builder
	var rows [][]InlineKeyboardButton
	for pos := 0; pos < len(text); {
		end := pos + 1
		var name, target string
		var matched bool
		if mode == ParseModeHTML && text[pos] == '<' {
			if tag, ok := readButtonHTMLTag(text, pos); ok {
				end = tag.end
				if !tag.closing && (tag.name == "code" || tag.name == "pre" || tag.name == "a") {
					end = buttonHTMLElementEnd(text, tag)
				}
			}
		} else {
			switch text[pos] {
			case '\\':
				if end < len(text) {
					end++
					if mode == ParseModeNone && (text[pos+1] == '\\' || text[pos+1] == '[') {
						out.WriteString(text[pos+1 : end])
						pos = end
						continue
					}
				}
			case '`':
				if mode == ParseModeMarkdown || mode == ParseModeMarkdownV2 {
					end = buttonMarkdownCodeEnd(text, pos)
				}
			case '[':
				if label, destination, linkEnd, ok := readButtonMarkdownLink(text, pos); ok {
					end = linkEnd
					image := pos > 0 && text[pos-1] == '!'
					if image {
						for before := pos - 2; before >= 0 && text[before] == '\\'; before-- {
							image = !image
						}
					}
					if !image && isButtonTarget(destination) {
						name, target, matched = unescapeButtonMarkdown(label), destination, true
					}
				}
			}
		}
		if matched {
			button, same, err := parseInlineButton(name, target)
			if err != nil {
				return "", nil, err
			}
			if same && len(rows) > 0 {
				rows[len(rows)-1] = append(rows[len(rows)-1], button)
			} else {
				rows = append(rows, []InlineKeyboardButton{button})
			}
		} else {
			out.WriteString(text[pos:end])
		}
		pos = end
	}
	if len(rows) == 0 {
		return out.String(), nil, nil
	}
	return strings.TrimSpace(out.String()), &ReplyMarkupInlineKeyboard{Rows: rows}, nil
}

// ParseTextWithButtons parses ordinary text and returns its inline keyboard.
func (c *Client) ParseTextWithButtons(text, parseMode string) (*FormattedText, *ReplyMarkupInlineKeyboard, error) {
	body, keyboard, err := ParseButtonMarkup(text, parseMode)
	if err != nil {
		return nil, nil, err
	}
	formatted, err := c.ParseText(body, parseMode)
	if err != nil {
		return nil, nil, err
	}
	return formatted, keyboard, nil
}

// Format renders an inline keyboard using the same button syntax in every text mode.
func (k *ReplyMarkupInlineKeyboard) Format() (string, error) {
	if k == nil {
		return "", nil
	}
	var out strings.Builder
	for _, row := range k.Rows {
		for column, button := range row {
			target, err := inlineButtonTarget(button)
			if err != nil {
				return "", err
			}
			if column > 0 {
				target += ":same"
			}
			if out.Len() > 0 {
				out.WriteByte('\n')
			}
			label := strings.NewReplacer("\\", "\\\\", "[", "\\[", "]", "\\]").Replace(button.Text)
			out.WriteString("[" + label + "](" + target + ")")
		}
	}
	return out.String(), nil
}

// OriginalTextWithButtons renders ordinary text or a caption with its inline keyboard.
func (m *Message) OriginalTextWithButtons(parseMode string) (string, error) {
	mode := strings.ToLower(parseMode)
	if mode != ParseModeNone && mode != ParseModeHTML && mode != ParseModeMarkdown && mode != ParseModeMarkdownV2 {
		return "", fmt.Errorf("unsupported text parse mode %q", parseMode)
	}
	text, err := m.GetFormattedText()
	if err != nil {
		text, err = m.GetFormattedCaption()
	}
	if err != nil {
		return "", err
	}
	var keyboard *ReplyMarkupInlineKeyboard
	switch markup := formattingValue(m.ReplyMarkup).(type) {
	case nil:
	case ReplyMarkupInlineKeyboard:
		keyboard = &markup
	default:
		return "", fmt.Errorf("button markup requires an inline keyboard")
	}
	buttons, err := keyboard.Format()
	if err != nil {
		return "", err
	}
	body := text.Format(parseMode)
	if mode == ParseModeHTML {
		body = strings.ReplaceAll(body, "[", "&#91;")
	} else if mode == ParseModeNone {
		body = strings.NewReplacer("\\", "\\\\", "[", "\\[").Replace(body)
	}
	if buttons != "" {
		body += "\n" + buttons
	}
	return body, nil
}
