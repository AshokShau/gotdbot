package ext

import (
	"encoding/base64"

	"github.com/AshokShau/gotdbot"
)

// CallbackQuery is a wrapper around gotdbot.UpdateNewCallbackQuery with bound methods.
// TDLib sends `updateNewCallbackQuery`. There is no `CallbackQuery` object in TDLib API directly,
type CallbackQuery struct {
	*gotdbot.UpdateNewCallbackQuery
	client *gotdbot.Client
}

// NewCallbackQuery creates a new bound CallbackQuery.
func NewCallbackQuery(client *gotdbot.Client, cb *gotdbot.UpdateNewCallbackQuery) *CallbackQuery {
	return &CallbackQuery{
		UpdateNewCallbackQuery: cb,
		client:                 client,
	}
}

// CallbackData returns the callback data payload as bytes.
func (c *CallbackQuery) CallbackData() []byte {
	if c.Payload == nil {
		return nil
	}
	if c.Payload.CallbackQueryPayloadData != nil {
		data, err := base64.StdEncoding.DecodeString(c.Payload.CallbackQueryPayloadData.Data)
		if err == nil {
			return data
		}

		return []byte(c.Payload.CallbackQueryPayloadData.Data)
	}
	return nil
}

// Text returns the callback data payload as string.
func (c *CallbackQuery) Text() string {
	return string(c.CallbackData())
}

// GetMessage returns the message that originated the query.
func (c *CallbackQuery) GetMessage() (*Message, error) {
	if c.MessageId == 0 {
		return nil, nil
	}
	msg, err := c.client.GetMessage(c.ChatId, c.MessageId)
	if err != nil {
		return nil, err
	}
	return NewMessage(c.client, msg), nil
}

// Answer sends an answer to the callback query.
func (c *CallbackQuery) Answer(text string, showAlert bool, url string, cacheTime int32) error {
	_, err := c.client.AnswerCallbackQuery(c.Id, text, showAlert, url, cacheTime)
	return err
}
