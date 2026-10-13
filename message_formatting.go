package gotdbot

// OriginalMD gets the original markdown formatting of a message text.
func (m *Message) OriginalMD() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format("markdown")
}

// OriginalMDV2 gets the original markdownV2 formatting of a message text.
func (m *Message) OriginalMDV2() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format("markdownv2")
}

// OriginalHTML gets the original HTML formatting of a message text.
func (m *Message) OriginalHTML() string {
	text, err := m.GetFormattedText()
	if err != nil {
		return ""
	}
	return text.Format("html")
}

// OriginalCaptionMD gets the original markdown formatting of a message caption.
func (m *Message) OriginalCaptionMD() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format("markdown")
}

// OriginalCaptionMDV2 gets the original markdownV2 formatting of a message caption.
func (m *Message) OriginalCaptionMDV2() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format("markdownv2")
}

// OriginalCaptionHTML gets the original HTML formatting of a message caption.
func (m *Message) OriginalCaptionHTML() string {
	caption, err := m.GetFormattedCaption()
	if err != nil {
		return ""
	}
	return caption.Format("html")
}
