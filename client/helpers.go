package client

import (
	"os"
	"strings"

	"github.com/AshokShau/gotdbot/types"
)

// ParseText parses the text using the specified parse mode.
func (c *Client) ParseText(text string, parseMode string) (*types.FormattedText, error) {
	var mode *types.TextParseMode

	switch strings.ToLower(parseMode) {
	case "markdown":
		mode = &types.TextParseMode{
			TextParseModeMarkdown: &types.TextParseModeMarkdown{Version: 1},
		}
	case "markdownv2":
		mode = &types.TextParseMode{
			TextParseModeMarkdown: &types.TextParseModeMarkdown{Version: 2},
		}
	case "html":
		mode = &types.TextParseMode{
			TextParseModeHTML: &types.TextParseModeHTML{},
		}
	default:
		return &types.FormattedText{Text: text}, nil
	}

	return c.ParseTextEntities(text, mode)
}

func getFormattedText(c *Client, text string, entities []*types.TextEntity, parseMode string) *types.FormattedText {
	if len(entities) > 0 {
		return &types.FormattedText{
			Text:     text,
			Entities: entities,
		}
	} else if parseMode != "" {
		ft, err := c.ParseText(text, parseMode)
		if err == nil {
			return ft
		}
	}
	return &types.FormattedText{Text: text}
}

func getInputFile(path string) *types.InputFile {
	if _, err := os.Stat(path); err == nil {
		return &types.InputFile{
			InputFileLocal: &types.InputFileLocal{Path: path},
		}
	}

	return &types.InputFile{
		InputFileRemote: &types.InputFileRemote{Id: path},
	}
}
