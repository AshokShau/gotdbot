package gotdbot

// StopLiveLocation stops sharing this message's live location.
func (m *Message) StopLiveLocation(c *Client, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return c.StopMessageLiveLocation(m.ChatId, m.Id, opts)
}

// EditText edits a text message.
func (m *Message) EditText(c *Client, text string, opts *EditMessageTextOpts) (*Message, error) {
	return c.EditMessageText(m.ChatId, m.Id, text, opts)
}

// EditCaption edits the caption of a message.
func (m *Message) EditCaption(c *Client, caption string, opts *EditMessageCaptionOpts) (*Message, error) {
	return c.EditMessageCaption(m.ChatId, m.Id, caption, opts)
}

// EditReplyMarkup replaces or removes this message's reply markup.
func (m *Message) EditReplyMarkup(c *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return c.EditMessageReplyMarkup(m.ChatId, m.Id, opts)
}

// EditContent edits the message with the given content.
func (m *Message) EditContent(c *Client, content InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return c.EditContent(m.ChatId, m.Id, content, opts)
}

// EditPhoto replaces this message's photo and caption.
func (m *Message) EditPhoto(c *Client, photo InputFile, opts *EditMessagePhotoOpts) (*Message, error) {
	return c.EditMessagePhoto(m.ChatId, m.Id, photo, opts)
}

// EditVideo replaces this message's video and caption.
func (m *Message) EditVideo(c *Client, video InputFile, opts *EditMessageVideoOpts) (*Message, error) {
	return c.EditMessageVideo(m.ChatId, m.Id, video, opts)
}

// EditAnimation replaces this message's animation and caption.
func (m *Message) EditAnimation(c *Client, animation InputFile, opts *EditMessageAnimationOpts) (*Message, error) {
	return c.EditMessageAnimation(m.ChatId, m.Id, animation, opts)
}

// EditAudio replaces this message's audio and caption.
func (m *Message) EditAudio(c *Client, audio InputFile, opts *EditMessageAudioOpts) (*Message, error) {
	return c.EditMessageAudio(m.ChatId, m.Id, audio, opts)
}

// EditDocument replaces this message's document and caption.
func (m *Message) EditDocument(c *Client, document InputFile, opts *EditMessageDocumentOpts) (*Message, error) {
	return c.EditMessageDocument(m.ChatId, m.Id, document, opts)
}
