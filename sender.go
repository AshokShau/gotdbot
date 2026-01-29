package gotdbot

// FromID returns the user ID or chat ID of the sender.
func (t *MessageSender) FromID() int64 {
	if t.MessageSenderUser != nil {
		return t.MessageSenderUser.UserId
	}
	if t.MessageSenderChat != nil {
		return t.MessageSenderChat.ChatId
	}
	return 0
}

// IsUser returns true if the sender is a user.
func (t *MessageSender) IsUser() bool {
	return t.MessageSenderUser != nil
}

// IsChat returns true if the sender is a chat.
func (t *MessageSender) IsChat() bool {
	return t.MessageSenderChat != nil
}
