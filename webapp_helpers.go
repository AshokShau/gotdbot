package gotdbot

import (
	"encoding/json"

	"github.com/AshokShau/gotdbot/webapp"
)

func (c *Client) ValidateTelegramInitData(raw string, opts *webapp.ValidationOptions) (*webapp.WebAppInitData, error) {
	if c == nil {
		return nil, webapp.ErrBotTokenRequired
	}
	return webapp.ValidateTelegramInitData(raw, c.botToken, opts)
}

func (m *Message) GetWebAppData() (*MessageWebAppDataReceived, error) {
	if m != nil {
		if content, ok := formattingValue(m.Content).(MessageWebAppDataReceived); ok {
			return &content, nil
		}
	}
	return nil, ErrNoWebAppData
}

func (m *Message) DecodeWebAppData(dst any) error {
	data, err := m.GetWebAppData()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data.Data), dst)
}
