package gotdbot

// CallbackData returns the callback data payload as bytes.
func (t *UpdateNewCallbackQuery) CallbackData() []byte {
	if t.Payload == nil {
		return nil
	}
	if p, ok := t.Payload.(*CallbackQueryPayloadData); ok {
		return p.Data
	}
	if p, ok := t.Payload.(*CallbackQueryPayloadDataWithPassword); ok {
		return p.Data
	}
	return nil
}

// DataString returns the callback data payload as a string.
func (t *UpdateNewCallbackQuery) DataString() string {
	if t.CallbackData() != nil {
		return string(t.CallbackData())
	}
	return ""
}

// CallbackDataWithPassword returns the callback data payload as bytes, along with the password if it exists.
func (t *UpdateNewCallbackQuery) CallbackDataWithPassword() ([]byte, string) {
	if t.Payload == nil {
		return nil, ""
	}
	if p, ok := t.Payload.(*CallbackQueryPayloadDataWithPassword); ok {
		return p.Data, p.Password
	}
	return nil, ""
}

// GameShortName returns the short name of the game associated with the callback query.
func (t *UpdateNewCallbackQuery) GameShortName() string {
	if t.Payload == nil {
		return ""
	}
	if p, ok := t.Payload.(*CallbackQueryPayloadGame); ok {
		return p.GameShortName
	}
	return ""
}

// IsPrivate checks if the callback query originated from a private chat.
func (t *UpdateNewCallbackQuery) IsPrivate() bool {
	return t.ChatId > 0
}

// GetMessage returns the message that originated the query.
func (t *UpdateNewCallbackQuery) GetMessage(c *Client) (*Message, error) {
	if t.MessageId == 0 {
		return nil, nil
	}
	return c.GetMessage(t.ChatId, t.MessageId)
}

// Answer sends an answer to the callback query.
func (t *UpdateNewCallbackQuery) Answer(c *Client, cacheTime int32, showAlert bool, text string, url string) error {
	err := c.AnswerCallbackQuery(cacheTime, t.Id, text, url, &AnswerCallbackQueryOpts{ShowAlert: showAlert})
	return err
}

// EditMessageText edits the text of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageText(c *Client, text string, opts *EditMessageTextOpts) (*Message, error) {
	return c.EditMessageText(t.ChatId, t.MessageId, text, opts)
}

// EditMessageCaption edits the caption of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageCaption(c *Client, caption string, opts *EditMessageCaptionOpts) (*Message, error) {
	return c.EditMessageCaption(t.ChatId, t.MessageId, caption, opts)
}

// EditMessageReplyMarkup edits the reply markup of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageReplyMarkup(c *Client, replyMarkup ReplyMarkup) (*Message, error) {
	return c.EditMessageReplyMarkup(t.ChatId, t.MessageId, &EditMessageReplyMarkupOpts{ReplyMarkup: replyMarkup})
}

// DeleteMessage deletes the message from which a chat callback originated.
func (t *UpdateNewCallbackQuery) DeleteMessage(c *Client) error {
	return c.DeleteMessage(t.ChatId, t.MessageId)
}

// EditMessageMedia replaces the callback message's native media content.
func (t *UpdateNewCallbackQuery) EditMessageMedia(c *Client, content InputMessageContent, replyMarkup ReplyMarkup) (*Message, error) {
	return c.EditContent(t.ChatId, t.MessageId, content, replyMarkup)
}

// CallbackData returns the payload for an inline callback, including password data.
func (t *UpdateNewInlineCallbackQuery) CallbackData() []byte {
	return (&UpdateNewCallbackQuery{Payload: t.Payload}).CallbackData()
}

// DataString returns the inline callback's data as text.
func (t *UpdateNewInlineCallbackQuery) DataString() string { return string(t.CallbackData()) }

// GameShortName returns the short name of the callback's game.
func (t *UpdateNewInlineCallbackQuery) GameShortName() string {
	return (&UpdateNewCallbackQuery{Payload: t.Payload}).GameShortName()
}

// Answer answers an inline callback without fetching a chat message.
func (t *UpdateNewInlineCallbackQuery) Answer(c *Client, cacheTime int32, showAlert bool, text, url string) error {
	return c.AnswerCallbackQuery(cacheTime, t.Id, text, url, &AnswerCallbackQueryOpts{ShowAlert: showAlert})
}

// EditMessageText edits an inline callback's text with native formatting options.
func (t *UpdateNewInlineCallbackQuery) EditMessageText(c *Client, text string, opts *EditMessageTextOpts) error {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}
	formatted, err := c.GetFormattedText(text, opts.Entities, opts.ParseMode)
	if err != nil {
		return err
	}
	content := &InputMessageText{Text: formatted, LinkPreviewOptions: &LinkPreviewOptions{
		IsDisabled: opts.DisableWebPagePreview, Url: opts.Url, ForceSmallMedia: opts.ForceSmallMedia,
		ForceLargeMedia: opts.ForceLargeMedia, ShowAboveText: opts.ShowAboveText,
	}}
	return c.EditInlineMessageText(t.InlineMessageId, content, &EditInlineMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageCaption edits an inline callback's caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageCaption(c *Client, caption string, opts *EditMessageCaptionOpts) error {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}
	formatted, err := c.GetFormattedText(caption, opts.Entities, opts.ParseMode)
	if err != nil {
		return err
	}
	return c.EditInlineMessageCaption(t.InlineMessageId, &EditInlineMessageCaptionOpts{
		Caption: formatted, ReplyMarkup: opts.ReplyMarkup, ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	})
}

// EditMessageReplyMarkup edits an inline callback's keyboard.
func (t *UpdateNewInlineCallbackQuery) EditMessageReplyMarkup(c *Client, replyMarkup ReplyMarkup) error {
	return c.EditInlineMessageReplyMarkup(t.InlineMessageId, &EditInlineMessageReplyMarkupOpts{ReplyMarkup: replyMarkup})
}

// EditMessageMedia replaces an inline callback's native media.
func (t *UpdateNewInlineCallbackQuery) EditMessageMedia(c *Client, content InputMessageContent, replyMarkup ReplyMarkup) error {
	return c.EditInlineMessageMedia(t.InlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: replyMarkup})
}
