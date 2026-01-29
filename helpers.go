package gotdbot

import (
	"os"
	"strings"
)

// ParseText parses the text using the specified parse mode.
func (c *Client) ParseText(text string, parseMode string) (*FormattedText, error) {
	var mode *TextParseMode

	switch strings.ToLower(parseMode) {
	case "markdown":
		mode = &TextParseMode{
			TextParseModeMarkdown: &TextParseModeMarkdown{Version: 1},
		}
	case "markdownv2":
		mode = &TextParseMode{
			TextParseModeMarkdown: &TextParseModeMarkdown{Version: 2},
		}
	case "html":
		mode = &TextParseMode{
			TextParseModeHTML: &TextParseModeHTML{},
		}
	default:
		return &FormattedText{Text: text}, nil
	}

	return c.ParseTextEntities(text, mode)
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

func GetInputFile(path string) *InputFile {
	if _, err := os.Stat(path); err == nil {
		return &InputFile{
			InputFileLocal: &InputFileLocal{Path: path},
		}
	}

	return &InputFile{
		InputFileRemote: &InputFileRemote{Id: path},
	}
}
