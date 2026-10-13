package gotdbot

import (
	"fmt"
	"reflect"
)

// MessageID returns the chosen result's inline message ID, or an error if unknown.
func (u *UpdateNewChosenInlineResult) MessageID() (string, error) {
	if u == nil {
		return "", ErrNilUpdate
	}
	if u.InlineMessageId == "" {
		return "", ErrUpdateHasNoMessage
	}
	return u.InlineMessageId, nil
}

// EditText edits the chosen inline result's text when its message ID is known.
func (u *UpdateNewChosenInlineResult) EditText(c *Client, text string, opts *EditMessageTextOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineText(id, text, opts)
}

// EditCaption edits the chosen inline result's caption when its message ID is known.
func (u *UpdateNewChosenInlineResult) EditCaption(c *Client, caption string, opts *EditMessageCaptionOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineCaption(id, caption, opts)
}

// EditContent replaces the chosen inline result's editable content.
func (u *UpdateNewChosenInlineResult) EditContent(c *Client, content InputMessageContent, opts *EditInlineMessageMediaOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineContent(id, content, opts)
}

// EditReplyMarkup replaces or removes the chosen inline result's keyboard.
func (u *UpdateNewChosenInlineResult) EditReplyMarkup(c *Client, opts *EditInlineMessageReplyMarkupOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessageReplyMarkup(id, opts)
}

// EditRichMessage replaces the chosen inline result with rich content.
func (u *UpdateNewChosenInlineResult) EditRichMessage(c *Client, message *InputRichMessage, opts *EditRichMessageOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineRichMessage(id, message, opts)
}

// EditRichMessageMarkdown replaces the chosen inline result with rich Markdown.
func (u *UpdateNewChosenInlineResult) EditRichMessageMarkdown(c *Client, text string, opts *EditRichTextOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineRichMessageMarkdown(id, text, opts)
}

// EditRichMessageHTML replaces the chosen inline result with rich HTML.
func (u *UpdateNewChosenInlineResult) EditRichMessageHTML(c *Client, text string, opts *EditRichTextOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineRichMessageHTML(id, text, opts)
}

// EditPhoto replaces the chosen inline result's photo and caption.
func (u *UpdateNewChosenInlineResult) EditPhoto(c *Client, photo InputFile, opts *EditMessagePhotoOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessagePhoto(id, photo, opts)
}

// EditVideo replaces the chosen inline result's video and caption.
func (u *UpdateNewChosenInlineResult) EditVideo(c *Client, video InputFile, opts *EditMessageVideoOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessageVideo(id, video, opts)
}

// EditAnimation replaces the chosen inline result's animation and caption.
func (u *UpdateNewChosenInlineResult) EditAnimation(c *Client, animation InputFile, opts *EditMessageAnimationOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessageAnimation(id, animation, opts)
}

// EditAudio replaces the chosen inline result's audio and caption.
func (u *UpdateNewChosenInlineResult) EditAudio(c *Client, audio InputFile, opts *EditMessageAudioOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessageAudio(id, audio, opts)
}

// EditDocument replaces the chosen inline result's document and caption.
func (u *UpdateNewChosenInlineResult) EditDocument(c *Client, document InputFile, opts *EditMessageDocumentOpts) error {
	id, err := u.MessageID()
	if err != nil {
		return err
	}
	return c.EditInlineMessageDocument(id, document, opts)
}

// messageFromUpdate extracts full messages from new-message, business,
// guest-query, and send-completion updates. Updates containing only message IDs
// return nil, false; no request is sent.
func messageFromUpdate(update TlObject) (*Message, bool) {
	switch u := update.(type) {
	case *UpdateNewMessage:
		if u != nil && u.Message != nil {
			return u.Message, true
		}

	case *UpdateNewBusinessMessage:
		if u != nil && u.Message != nil && u.Message.Message != nil {
			return u.Message.Message, true
		}

	case *UpdateBusinessMessageEdited:
		if u != nil && u.Message != nil && u.Message.Message != nil {
			return u.Message.Message, true
		}
	case *UpdateNewBusinessCallbackQuery:
		if u != nil && u.Message != nil && u.Message.Message != nil {
			return u.Message.Message, true
		}
	case *UpdateNewGuestQuery:
		if u != nil && u.Message != nil {
			return u.Message, true
		}
	case *UpdateMessageSendSucceeded:
		if u != nil && u.Message != nil {
			return u.Message, true
		}
	case *UpdateMessageSendFailed:
		if u != nil && u.Message != nil {
			return u.Message, true
		}
	}

	return nil, false
}

// GetUpdateMessage returns a carried message, or fetches a message when the
// update contains only its chat and message IDs. Updates without a single message
// return an error; no request is sent for those updates.
func (c *Client) GetUpdateMessage(update Update) (*Message, error) {
	if _, ok := nativeUpdate(update); !ok {
		return nil, ErrNilUpdate
	}
	if message, ok := messageFromUpdate(update); ok {
		return message, nil
	}
	chatId, messageId := c.UpdateChatID(update), c.UpdateMessageID(update)
	if chatId == 0 || messageId == 0 {
		return nil, fmt.Errorf("%w: %s", ErrUpdateHasNoMessage, update.GetType())
	}
	return c.GetMessage(chatId, messageId)
}

// GetMessage returns the message carried by this guest query, or nil if absent.
func (u *UpdateNewGuestQuery) GetMessage() *Message {
	if u == nil {
		return nil
	}
	return u.Message
}

// Answer sends the result for this guest query.
func (u *UpdateNewGuestQuery) Answer(c *Client, result InputInlineQueryResult) (*InlineMessageId, error) {
	if u == nil {
		return nil, ErrNilUpdate
	}
	return c.AnswerGuestQuery(u.Id, result)
}

func getMessageFromUpdate(update TlObject) (*Message, bool) {
	// Message/command handlers receive incoming message updates, not query or
	// send-completion events that happen to carry a message.
	switch update.(type) {
	case *UpdateNewMessage, *UpdateNewBusinessMessage, *UpdateBusinessMessageEdited:
		return messageFromUpdate(update)
	default:
		return nil, false
	}
}

// UpdateChatID returns the chat ID from a message or a native update's ChatId
// field. It returns 0 when no chat is available, including for nil updates.
func (c *Client) UpdateChatID(object TlObject) int64 {
	update, ok := nativeUpdate(object)
	if !ok {
		return 0
	}
	if msg, ok := messageFromUpdate(update); ok {
		return msg.ChatId
	}

	return updateInt64Field(update, "ChatId")
}

// UpdateMessageID returns a message ID from a message or a native update's
// MessageId field. Send-completion updates return the current message ID.
// Updates without a single message ID, such as deletions, return 0.
func (c *Client) UpdateMessageID(object TlObject) int64 {
	update, ok := nativeUpdate(object)
	if !ok {
		return 0
	}
	if message, ok := messageFromUpdate(update); ok {
		return message.Id
	}
	return updateInt64Field(update, "MessageId")
}

// UpdateSenderID returns the sender's user or chat ID from a message,
// callback, query, or join request. It returns 0 when no sender is available.
func (c *Client) UpdateSenderID(object TlObject) int64 {
	update, ok := nativeUpdate(object)
	if !ok {
		return 0
	}
	if u, ok := update.(*UpdateNewChatJoinRequest); ok && u.Request != nil {
		return u.Request.UserId
	}
	value := updateValue(update)
	if value.Kind() == reflect.Struct {
		field := value.FieldByName("SenderUserId")
		if field.IsValid() && field.Kind() == reflect.Int64 {
			return field.Int()
		}
	}
	if message, ok := messageFromUpdate(update); ok {
		return message.SenderID()
	}
	if value.Kind() == reflect.Struct {
		field := value.FieldByName("SenderId")
		if field.IsValid() && field.CanInterface() {
			if sender, ok := field.Interface().(MessageSender); ok {
				return (&Message{SenderId: sender}).SenderID()
			}
		}
	}
	return 0
}

func updateValue(update Update) reflect.Value {
	value := reflect.ValueOf(update)
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}

func updateInt64Field(update Update, name string) int64 {
	value := updateValue(update)
	if value.Kind() == reflect.Struct {
		field := value.FieldByName(name)
		if field.IsValid() && field.Kind() == reflect.Int64 {
			return field.Int()
		}
	}
	return 0
}

// GetMessage returns the message carried by the update, or nil if absent.
func (u *UpdateNewMessage) GetMessage() *Message {
	if u == nil {
		return nil
	}
	return u.Message
}

// GetMessage returns the business message carried by the update, or nil if absent.
func (u *UpdateNewBusinessMessage) GetMessage() *Message {
	if u == nil || u.Message == nil {
		return nil
	}
	return u.Message.Message
}

// GetMessage returns the edited business message, or nil if absent.
func (u *UpdateBusinessMessageEdited) GetMessage() *Message {
	if u == nil || u.Message == nil {
		return nil
	}
	return u.Message.Message
}

// GetMessage fetches the current message identified by this edit update.
func (u *UpdateMessageEdited) GetMessage(c *Client) (*Message, error) {
	if u == nil {
		return nil, ErrNilUpdate
	}
	return c.GetMessage(u.ChatId, u.MessageId)
}

// GetMessage fetches the current message identified by this content update.
func (u *UpdateMessageContent) GetMessage(c *Client) (*Message, error) {
	if u == nil {
		return nil, ErrNilUpdate
	}
	return c.GetMessage(u.ChatId, u.MessageId)
}

// GetMessage returns the successfully sent message, or nil if absent.
func (u *UpdateMessageSendSucceeded) GetMessage() *Message {
	if u == nil {
		return nil
	}
	return u.Message
}

// GetMessage returns the message whose send failed, or nil if absent.
func (u *UpdateMessageSendFailed) GetMessage() *Message {
	if u == nil {
		return nil
	}
	return u.Message
}

// Answer sends results for this inline query.
func (u *UpdateNewInlineQuery) Answer(c *Client, results []InputInlineQueryResult, opts *AnswerInlineQueryOpts) error {
	if u == nil {
		return ErrNilUpdate
	}
	return c.AnswerInlineQuery(u.Id, results, opts)
}

// Answer accepts the checkout when errorMessage is empty, or rejects it with
// the supplied explanation. Telegram applies the payment validation rules.
func (u *UpdateNewPreCheckoutQuery) Answer(c *Client, errorMessage string) error {
	if u == nil {
		return ErrNilUpdate
	}
	return c.AnswerPreCheckoutQuery(u.Id, &AnswerPreCheckoutQueryOpts{ErrorMessage: errorMessage})
}

// Answer supplies available shipping options, or rejects the query when
// opts.ErrorMessage is non-empty.
func (u *UpdateNewShippingQuery) Answer(c *Client, shippingOptions []ShippingOption, opts *AnswerShippingQueryOpts) error {
	if u == nil {
		return ErrNilUpdate
	}
	return c.AnswerShippingQuery(shippingOptions, u.Id, opts)
}

// Approve accepts this user's join request.
func (u *UpdateNewChatJoinRequest) Approve(c *Client) error {
	if u == nil {
		return ErrNilUpdate
	}
	if u.Request == nil {
		return fmt.Errorf("chat join request is missing")
	}
	return c.ApproveChatJoinRequest(u.ChatId, u.Request.UserId)
}

// Decline rejects this user's join request.
func (u *UpdateNewChatJoinRequest) Decline(c *Client) error {
	if u == nil {
		return ErrNilUpdate
	}
	if u.Request == nil {
		return fmt.Errorf("chat join request is missing")
	}
	return c.DeclineChatJoinRequest(u.ChatId, u.Request.UserId)
}
