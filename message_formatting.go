package gotdbot

// GetRichMessage returns rich content; fetch GetFullRichMessage if IsFull is false.
func (m *Message) GetRichMessage() (*RichMessage, error) {
	if m != nil {
		switch content := formattingValue(m.Content).(type) {
		case MessageRichMessage:
			if content.Message != nil {
				return content.Message, nil
			}
		}
	}
	return nil, ErrNoRichMessage
}

// OriginalRichHTML renders rich content, text or a caption as rich HTML with media references.
func (m *Message) OriginalRichHTML() (string, []InputRichMessageMedia) {
	return m.richFormattingContent().HTML()
}

// OriginalRichMD renders rich content, text or a caption as rich Markdown with media references.
func (m *Message) OriginalRichMD() (string, []InputRichMessageMedia) {
	return m.richFormattingContent().Markdown()
}

func (m *Message) richFormattingContent() *RichMessage {
	if rich, err := m.GetRichMessage(); err == nil {
		return rich
	}
	if text, err := m.GetFormattedText(); err == nil {
		return formattedTextRichMessage(text)
	}
	if caption, err := m.GetFormattedCaption(); err == nil {
		return formattedTextRichMessage(caption)
	}
	return nil
}

// OriginalTextHTML renders ordinary message text or a media caption as HTML.
func (m *Message) OriginalTextHTML() string { return m.originalTextFormat(ParseModeHTML) }

// OriginalTextMD renders ordinary message text or a caption as legacy Markdown.
func (m *Message) OriginalTextMD() string { return m.originalTextFormat(ParseModeMarkdown) }

// OriginalTextMDV2 renders ordinary message text or a caption as MarkdownV2.
func (m *Message) OriginalTextMDV2() string { return m.originalTextFormat(ParseModeMarkdownV2) }

func (m *Message) originalTextFormat(mode string) string {
	if text, err := m.GetFormattedText(); err == nil {
		return text.Format(mode)
	}
	if caption, err := m.GetFormattedCaption(); err == nil {
		return caption.Format(mode)
	}
	return ""
}

// OriginalMD gets the original markdown formatting of a message text.
func (m *Message) OriginalMD() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format(ParseModeMarkdown)
}

// OriginalMDV2 gets the original markdownV2 formatting of a message text.
func (m *Message) OriginalMDV2() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format(ParseModeMarkdownV2)
}

// OriginalHTML gets the original HTML formatting of a message text.
func (m *Message) OriginalHTML() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format(ParseModeHTML)
}

// OriginalCaptionMD gets the original markdown formatting of a message caption.
func (m *Message) OriginalCaptionMD() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format(ParseModeMarkdown)
}

// OriginalCaptionMDV2 gets the original markdownV2 formatting of a message caption.
func (m *Message) OriginalCaptionMDV2() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format(ParseModeMarkdownV2)
}

// OriginalCaptionHTML gets the original HTML formatting of a message caption.
func (m *Message) OriginalCaptionHTML() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format(ParseModeHTML)
}
