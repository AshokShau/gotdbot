package gotdbot

import (
	"fmt"
	"html"
	"os"
	"strings"
)

func Bool(b bool) *bool {
	return &b
}

func GetFormattedText(c *Client, text string, entities []*TextEntity, parseMode string) *FormattedText {
	if len(entities) > 0 {
		return &FormattedText{
			Text:     text,
			Entities: entities,
		}
	} else if parseMode != "" {
		ft, err := c.ParseText(text, parseMode)
		if err == nil {
			return ft
		}
	}
	return &FormattedText{Text: text}
}

func GetInputFile(path string) InputFile {
	if _, err := os.Stat(path); err == nil {
		return &InputFileLocal{Path: path}
	}

	return &InputFileRemote{Id: path}
}

// EscapeHTML escapes HTML characters in the given text.
func EscapeHTML(text string) string {
	return html.EscapeString(text)
}

// EscapeMarkdown escapes Markdown characters in the given text.
func EscapeMarkdown(text string, version int) string {
	var chars string
	if version == 1 {
		chars = "_*`[\\"
	} else {
		chars = "_*[]()~`>#+-=|{}.!\\"
	}
	var b strings.Builder
	for _, c := range text {
		if strings.ContainsRune(chars, c) {
			b.WriteRune('\\')
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Mention returns a text mention for the given user ID.
func Mention(text string, userId int64, isHtml bool, escape bool) string {
	if escape {
		if isHtml {
			text = EscapeHTML(text)
		} else {
			text = EscapeMarkdown(text, 2)
		}
	}
	if isHtml {
		return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", userId, text)
	}
	return fmt.Sprintf("[%s](tg://user?id=%d)", text, userId)
}
