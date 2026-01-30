package gotdbot

// CallbackData returns the callback data payload as bytes.
func (t *UpdateNewCallbackQuery) CallbackData() []byte {
	if t.Payload == nil {
		return nil
	}
	if t.Payload.CallbackQueryPayloadData != nil {
		return t.Payload.CallbackQueryPayloadData.Data
	}
	return nil
}

// GetMessage returns the message that originated the query.
func (t *UpdateNewCallbackQuery) GetMessage(c *Client) (*Message, error) {
	if t.MessageId == 0 {
		return nil, nil
	}
	return c.GetMessage(t.ChatId, t.MessageId)
}

// Answer sends an answer to the callback query.
func (t *UpdateNewCallbackQuery) Answer(c *Client, text string, showAlert bool, url string, cacheTime int32) error {
	_, err := c.AnswerCallbackQuery(t.Id, text, showAlert, url, cacheTime)
	return err
}

// EditMessageText edits the text of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageText(c *Client, text string, opts *EditTextMessageOpts) (*Message, error) {
	return c.EditTextMessage(t.ChatId, t.MessageId, text, opts)
}

// EditMessageReplyMarkup edits the reply markup of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageReplyMarkup(c *Client, replyMarkup *ReplyMarkup) (*Message, error) {
	return c.EditMessageReplyMarkup(t.ChatId, t.MessageId, &EditMessageReplyMarkupOpts{ReplyMarkup: replyMarkup})
}
