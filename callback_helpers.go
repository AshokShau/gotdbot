package gotdbot

// GetMessage returns the message carried by the business callback.
// It returns ErrNilUpdate for a nil callback and ErrUpdateHasNoMessage if absent.
func (t *UpdateNewBusinessCallbackQuery) GetMessage() (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	if t.Message == nil || t.Message.Message == nil {
		return nil, ErrUpdateHasNoMessage
	}
	return t.Message.Message, nil
}

// CallbackData returns the business callback's data payload.
func (t *UpdateNewBusinessCallbackQuery) CallbackData() []byte {
	if t == nil {
		return nil
	}
	return callbackPayloadData(t.Payload)
}

// DataString returns the business callback's data as text.
func (t *UpdateNewBusinessCallbackQuery) DataString() string { return string(t.CallbackData()) }

// CallbackDataWithPassword returns the business callback's password-protected data.
func (t *UpdateNewBusinessCallbackQuery) CallbackDataWithPassword() ([]byte, string) {
	if t == nil {
		return nil, ""
	}
	return callbackPayloadPassword(t.Payload)
}

// GameShortName returns the business callback's game name, or an empty string.
func (t *UpdateNewBusinessCallbackQuery) GameShortName() string {
	if t == nil {
		return ""
	}
	return callbackPayloadGame(t.Payload)
}

// Answer sends a notification or alert for the business callback.
func (t *UpdateNewBusinessCallbackQuery) Answer(c *Client, text string, opts *CallbackAnswerOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.answerCallbackQuery(t.Id, text, opts)
}

// EditMessageText replaces the business callback message's formatted text.
func (t *UpdateNewBusinessCallbackQuery) EditMessageText(c *Client, text string, opts *EditMessageTextOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}
	content, err := c.textContentForEdit(text, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessMessageText(t.ConnectionId, message.ChatId, content, message.Id, &EditBusinessMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageCaption replaces or removes the business callback message's caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessageCaption(c *Client, caption string, opts *EditMessageCaptionOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}
	formatted, err := c.GetFormattedText(caption, opts.Entities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessMessageCaption(t.ConnectionId, message.ChatId, message.Id, &EditBusinessMessageCaptionOpts{Caption: formatted, ReplyMarkup: opts.ReplyMarkup, ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia})
}

// EditMessageReplyMarkup replaces or removes the business callback's keyboard.
func (t *UpdateNewBusinessCallbackQuery) EditMessageReplyMarkup(c *Client, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	return c.EditBusinessMessageReplyMarkup(t.ConnectionId, message.ChatId, message.Id, opts)
}

// EditMessageContent replaces the business callback message's editable content.
func (t *UpdateNewBusinessCallbackQuery) EditMessageContent(c *Client, content InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, opts)
}

// EditRichMessage replaces the business callback message with rich content.
func (t *UpdateNewBusinessCallbackQuery) EditRichMessage(c *Client, message *InputRichMessage, opts *EditRichMessageOpts) (*BusinessMessage, error) {
	if opts == nil {
		opts = &EditRichMessageOpts{}
	}
	return t.EditMessageContent(c, &InputMessageRichMessage{Message: message}, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditRichMessageMarkdown replaces the business callback message with rich Markdown.
func (t *UpdateNewBusinessCallbackQuery) EditRichMessageMarkdown(c *Client, text string, opts *EditRichTextOpts) (*BusinessMessage, error) {
	message, markup := richMarkdownForEdit(text, opts)
	return t.EditRichMessage(c, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditRichMessageHTML replaces the business callback message with rich HTML.
func (t *UpdateNewBusinessCallbackQuery) EditRichMessageHTML(c *Client, text string, opts *EditRichTextOpts) (*BusinessMessage, error) {
	message, markup := richHTMLForEdit(text, opts)
	return t.EditRichMessage(c, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditMessagePhoto replaces the business callback message's photo and caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessagePhoto(c *Client, photo InputFile, opts *EditMessagePhotoOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessagePhotoOpts{}
	}
	content, err := c.photoContentForEdit(photo, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageVideo replaces the business callback message's video and caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessageVideo(c *Client, video InputFile, opts *EditMessageVideoOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageVideoOpts{}
	}
	content, err := c.videoContentForEdit(video, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageAnimation replaces the business callback message's animation and caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessageAnimation(c *Client, animation InputFile, opts *EditMessageAnimationOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageAnimationOpts{}
	}
	content, err := c.animationContentForEdit(animation, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageAudio replaces the business callback message's audio and caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessageAudio(c *Client, audio InputFile, opts *EditMessageAudioOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageAudioOpts{}
	}
	content, err := c.audioContentForEdit(audio, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageDocument replaces the business callback message's document and caption.
func (t *UpdateNewBusinessCallbackQuery) EditMessageDocument(c *Client, document InputFile, opts *EditMessageDocumentOpts) (*BusinessMessage, error) {
	message, err := t.GetMessage()
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageDocumentOpts{}
	}
	content, err := c.documentContentForEdit(document, opts)
	if err != nil {
		return nil, err
	}
	return c.EditBusinessContent(t.ConnectionId, message.ChatId, message.Id, content, &EditBusinessMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageLiveLocation updates or stops the callback message's live location.
func (t *UpdateNewCallbackQuery) EditMessageLiveLocation(c *Client, opts *EditMessageLiveLocationOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessageLiveLocation(t.ChatId, t.MessageId, opts)
}

// StopMessageLiveLocation stops the callback message's live location.
func (t *UpdateNewCallbackQuery) StopMessageLiveLocation(c *Client, opts *EditMessageLiveLocationOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.StopMessageLiveLocation(t.ChatId, t.MessageId, opts)
}

// EditMessageLiveLocation updates or stops the callback message's live location.
func (t *UpdateNewInlineCallbackQuery) EditMessageLiveLocation(c *Client, opts *EditInlineMessageLiveLocationOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessageLiveLocation(t.InlineMessageId, opts)
}

// StopMessageLiveLocation stops the callback message's live location.
func (t *UpdateNewInlineCallbackQuery) StopMessageLiveLocation(c *Client, opts *EditInlineMessageLiveLocationOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.StopInlineMessageLiveLocation(t.InlineMessageId, opts)
}

// CallbackAnswerOpts controls a callback notification or alert.
type CallbackAnswerOpts struct {
	// CacheTime is the maximum time, in seconds, the answer can be cached.
	CacheTime int32
	// ShowAlert displays an alert instead of a notification.
	ShowAlert bool
	// Url opens a game or a supported bot deep link.
	Url string
}

func (c *Client) answerCallbackQuery(id int64, text string, opts *CallbackAnswerOpts) error {
	if opts == nil {
		opts = &CallbackAnswerOpts{}
	}
	return c.AnswerCallbackQuery(id, &AnswerCallbackQueryOpts{
		Text: text, Url: opts.Url, CacheTime: opts.CacheTime, ShowAlert: opts.ShowAlert,
	})
}

// EditRichMessage replaces the callback's message with rich content.
func (t *UpdateNewCallbackQuery) EditRichMessage(c *Client, message *InputRichMessage, opts *EditRichMessageOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditRichMessage(t.ChatId, t.MessageId, message, opts)
}

// EditRichMessageMarkdown replaces the callback's message with rich Markdown.
func (t *UpdateNewCallbackQuery) EditRichMessageMarkdown(c *Client, text string, opts *EditRichTextOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditRichMessageMarkdown(t.ChatId, t.MessageId, text, opts)
}

// EditRichMessageHTML replaces the callback's message with rich HTML.
func (t *UpdateNewCallbackQuery) EditRichMessageHTML(c *Client, text string, opts *EditRichTextOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditRichMessageHTML(t.ChatId, t.MessageId, text, opts)
}

// EditRichMessage replaces the inline callback's message with rich content.
func (t *UpdateNewInlineCallbackQuery) EditRichMessage(c *Client, message *InputRichMessage, opts *EditRichMessageOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineRichMessage(t.InlineMessageId, message, opts)
}

// EditRichMessageMarkdown replaces the inline callback's message with rich Markdown.
func (t *UpdateNewInlineCallbackQuery) EditRichMessageMarkdown(c *Client, text string, opts *EditRichTextOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineRichMessageMarkdown(t.InlineMessageId, text, opts)
}

// EditRichMessageHTML replaces the inline callback's message with rich HTML.
func (t *UpdateNewInlineCallbackQuery) EditRichMessageHTML(c *Client, text string, opts *EditRichTextOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineRichMessageHTML(t.InlineMessageId, text, opts)
}

func callbackPayloadData(payload CallbackQueryPayload) []byte {
	switch value := payload.(type) {
	case *CallbackQueryPayloadData:
		if value != nil {
			return value.Data
		}
	case CallbackQueryPayloadData:
		return value.Data
	case *CallbackQueryPayloadDataWithPassword:
		if value != nil {
			return value.Data
		}
	case CallbackQueryPayloadDataWithPassword:
		return value.Data
	}
	return nil
}

func callbackPayloadPassword(payload CallbackQueryPayload) ([]byte, string) {
	switch value := payload.(type) {
	case *CallbackQueryPayloadDataWithPassword:
		if value != nil {
			return value.Data, value.Password
		}
	case CallbackQueryPayloadDataWithPassword:
		return value.Data, value.Password
	}
	return nil, ""
}

func callbackPayloadGame(payload CallbackQueryPayload) string {
	switch value := payload.(type) {
	case *CallbackQueryPayloadGame:
		if value != nil {
			return value.GameShortName
		}
	case CallbackQueryPayloadGame:
		return value.GameShortName
	}
	return ""
}

// CallbackData returns the data payload, including password-protected data.
func (t *UpdateNewCallbackQuery) CallbackData() []byte {
	if t == nil {
		return nil
	}
	return callbackPayloadData(t.Payload)
}

// DataString returns the callback data as text.
func (t *UpdateNewCallbackQuery) DataString() string { return string(t.CallbackData()) }

// CallbackDataWithPassword returns password-protected data and its password.
func (t *UpdateNewCallbackQuery) CallbackDataWithPassword() ([]byte, string) {
	if t == nil {
		return nil, ""
	}
	return callbackPayloadPassword(t.Payload)
}

// GameShortName returns the callback's game name, or an empty string.
func (t *UpdateNewCallbackQuery) GameShortName() string {
	if t == nil {
		return ""
	}
	return callbackPayloadGame(t.Payload)
}

// IsPrivate checks if the callback query originated from a private chat.
func (t *UpdateNewCallbackQuery) IsPrivate() bool {
	return t != nil && t.ChatId > 0
}

// GetMessage returns the message that originated the query.
func (t *UpdateNewCallbackQuery) GetMessage(c *Client) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	if t.MessageId == 0 {
		return nil, nil
	}
	return c.GetMessage(t.ChatId, t.MessageId)
}

// Answer sends an answer to the callback query.
func (t *UpdateNewCallbackQuery) Answer(c *Client, text string, opts *CallbackAnswerOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.answerCallbackQuery(t.Id, text, opts)
}

// EditMessageText edits the text of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageText(c *Client, text string, opts *EditMessageTextOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditMessageText(t.ChatId, t.MessageId, text, opts)
}

// EditMessageCaption edits the caption of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageCaption(c *Client, caption string, opts *EditMessageCaptionOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditMessageCaption(t.ChatId, t.MessageId, caption, opts)
}

// EditMessageReplyMarkup edits the reply markup of the message associated with the callback query.
func (t *UpdateNewCallbackQuery) EditMessageReplyMarkup(c *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditMessageReplyMarkup(t.ChatId, t.MessageId, opts)
}

// DeleteMessage deletes the message from which a chat callback originated.
func (t *UpdateNewCallbackQuery) DeleteMessage(c *Client, opts *DeleteMessagesOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.DeleteMessage(t.ChatId, t.MessageId, opts)
}

// EditMessageMedia replaces the callback message's editable content.
func (t *UpdateNewCallbackQuery) EditMessageMedia(c *Client, content InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}

	return c.EditContent(t.ChatId, t.MessageId, content, opts)
}

// CallbackData returns the payload for an inline callback, including password data.
func (t *UpdateNewInlineCallbackQuery) CallbackData() []byte {
	if t == nil {
		return nil
	}
	return callbackPayloadData(t.Payload)
}

// DataString returns the inline callback's data as text.
func (t *UpdateNewInlineCallbackQuery) DataString() string { return string(t.CallbackData()) }

// GameShortName returns the short name of the callback's game.
func (t *UpdateNewInlineCallbackQuery) GameShortName() string {
	if t == nil {
		return ""
	}
	return callbackPayloadGame(t.Payload)
}

// CallbackDataWithPassword returns the inline callback's password-protected data.
func (t *UpdateNewInlineCallbackQuery) CallbackDataWithPassword() ([]byte, string) {
	if t == nil {
		return nil, ""
	}
	return callbackPayloadPassword(t.Payload)
}

// Answer answers an inline callback without fetching a chat message.
func (t *UpdateNewInlineCallbackQuery) Answer(c *Client, text string, opts *CallbackAnswerOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.answerCallbackQuery(t.Id, text, opts)
}

// EditMessageText replaces the inline callback message's formatted text.
func (t *UpdateNewInlineCallbackQuery) EditMessageText(c *Client, text string, opts *EditMessageTextOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineText(t.InlineMessageId, text, opts)
}

// EditMessageCaption edits an inline callback's caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageCaption(c *Client, caption string, opts *EditMessageCaptionOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineCaption(t.InlineMessageId, caption, opts)
}

// EditMessageReplyMarkup edits an inline callback's keyboard.
func (t *UpdateNewInlineCallbackQuery) EditMessageReplyMarkup(c *Client, opts *EditInlineMessageReplyMarkupOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineMessageReplyMarkup(t.InlineMessageId, opts)
}

// EditMessageMedia replaces the inline callback message's editable content.
func (t *UpdateNewInlineCallbackQuery) EditMessageMedia(c *Client, content InputMessageContent, opts *EditInlineMessageMediaOpts) error {
	if t == nil {
		return ErrNilUpdate
	}

	return c.EditInlineContent(t.InlineMessageId, content, opts)
}

// EditMessagePhoto replaces the callback message's photo and caption.
func (t *UpdateNewCallbackQuery) EditMessagePhoto(c *Client, photo InputFile, opts *EditMessagePhotoOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessagePhoto(t.ChatId, t.MessageId, photo, opts)
}

// EditMessagePhoto replaces the callback message's photo and caption.
func (t *UpdateNewInlineCallbackQuery) EditMessagePhoto(c *Client, photo InputFile, opts *EditMessagePhotoOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessagePhoto(t.InlineMessageId, photo, opts)
}

// EditMessageVideo replaces the callback message's video and caption.
func (t *UpdateNewCallbackQuery) EditMessageVideo(c *Client, video InputFile, opts *EditMessageVideoOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessageVideo(t.ChatId, t.MessageId, video, opts)
}

// EditMessageVideo replaces the callback message's video and caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageVideo(c *Client, video InputFile, opts *EditMessageVideoOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessageVideo(t.InlineMessageId, video, opts)
}

// EditMessageAnimation replaces the callback message's animation and caption.
func (t *UpdateNewCallbackQuery) EditMessageAnimation(c *Client, animation InputFile, opts *EditMessageAnimationOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessageAnimation(t.ChatId, t.MessageId, animation, opts)
}

// EditMessageAnimation replaces the callback message's animation and caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageAnimation(c *Client, animation InputFile, opts *EditMessageAnimationOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessageAnimation(t.InlineMessageId, animation, opts)
}

// EditMessageAudio replaces the callback message's audio and caption.
func (t *UpdateNewCallbackQuery) EditMessageAudio(c *Client, audio InputFile, opts *EditMessageAudioOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessageAudio(t.ChatId, t.MessageId, audio, opts)
}

// EditMessageAudio replaces the callback message's audio and caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageAudio(c *Client, audio InputFile, opts *EditMessageAudioOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessageAudio(t.InlineMessageId, audio, opts)
}

// EditMessageDocument replaces the callback message's document and caption.
func (t *UpdateNewCallbackQuery) EditMessageDocument(c *Client, document InputFile, opts *EditMessageDocumentOpts) (*Message, error) {
	if t == nil {
		return nil, ErrNilUpdate
	}
	return c.EditMessageDocument(t.ChatId, t.MessageId, document, opts)
}

// EditMessageDocument replaces the callback message's document and caption.
func (t *UpdateNewInlineCallbackQuery) EditMessageDocument(c *Client, document InputFile, opts *EditMessageDocumentOpts) error {
	if t == nil {
		return ErrNilUpdate
	}
	return c.EditInlineMessageDocument(t.InlineMessageId, document, opts)
}
