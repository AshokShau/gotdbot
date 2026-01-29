package client

import "github.com/AshokShau/gotdbot/types"

// MessageSender is a wrapper around types.MessageSender with helper methods.
type MessageSender struct {
	*types.MessageSender
	client *Client
}

// NewMessageSender creates a new bound MessageSender.
func NewMessageSender(client *Client, sender *types.MessageSender) *MessageSender {
	return &MessageSender{
		MessageSender: sender,
		client:        client,
	}
}

// FromID returns the user ID or chat ID of the sender.
func (s *MessageSender) FromID() int64 {
	if s.MessageSenderUser != nil {
		return s.MessageSenderUser.UserId
	}
	if s.MessageSenderChat != nil {
		return s.MessageSenderChat.ChatId
	}
	return 0
}

// IsUser returns true if the sender is a user.
func (s *MessageSender) IsUser() bool {
	return s.MessageSenderUser != nil
}

// IsChat returns true if the sender is a chat.
func (s *MessageSender) IsChat() bool {
	return s.MessageSenderChat != nil
}
