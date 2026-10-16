package gotdbot

import (
	"fmt"
	"html"
	"strings"
)

// GetFormattedText returns a FormattedText object from the given text, entities and parse mode.
func (c *Client) GetFormattedText(text string, entities []TextEntity, parseMode string) (*FormattedText, error) {
	if len(entities) > 0 {
		return &FormattedText{
			Text:     text,
			Entities: entities,
		}, nil
	}

	if parseMode == ParseModeNone {
		if c.config != nil {
			parseMode = c.config.ParseMode
		}
	}

	return c.ParseText(text, parseMode)
}

// Mention returns an escaped text mention using the user's full name.
// Use HTML for an HTML link; other parse modes use MarkdownV2 escaping.
func (u *User) Mention(parseMode string) string {
	if u == nil {
		return ""
	}
	text := strings.TrimSpace(u.FirstName + " " + u.LastName)
	if strings.EqualFold(parseMode, ParseModeHTML) {
		return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", u.Id, html.EscapeString(text))
	}
	name := (&FormattedText{Text: text}).Format(ParseModeMarkdownV2)
	return fmt.Sprintf("[%s](tg://user?id=%d)", name, u.Id)
}

// ParseText parses the text using the specified parse mode.
func (c *Client) ParseText(text string, parseMode string) (*FormattedText, error) {
	var mode TextParseMode

	switch strings.ToLower(parseMode) {
	case ParseModeMarkdown:
		mode = &TextParseModeMarkdown{Version: 1}
	case ParseModeMarkdownV2:
		mode = &TextParseModeMarkdown{Version: 2}
	case ParseModeHTML:
		mode = &TextParseModeHTML{}
	default:
		return &FormattedText{Text: text}, nil
	}

	return c.ParseTextEntities(mode, text)
}
