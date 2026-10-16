package gotdbot

import (
	"strings"
)

// IsPrivate checks if the message is from a private chat (user).
func (m *Message) IsPrivate() bool {
	return m != nil && m.ChatId > 0 && m.ChatId < 1<<40
}

// IsGroup checks if the message is from a basic group, excluding supergroups.
func (m *Message) IsGroup() bool {
	return m != nil && m.ChatId < 0 && m.ChatId > supergroupChatIDOffset
}

// IsSupergroup checks if the message is from a supergroup, excluding channels.
func (m *Message) IsSupergroup() bool {
	return m != nil && isSupergroupOrChannelID(m.ChatId) && !m.IsChannelPost
}

// IsSecret checks if the message is from a secret chat.
func (m *Message) IsSecret() bool {
	return m != nil && m.ChatId >= minSecretChatID && m.ChatId <= maxSecretChatID && m.ChatId != secretChatIDOffset
}

// IsOutgoingMessage returns true if the message is outgoing.
func (m *Message) IsOutgoingMessage() bool {
	return m != nil && m.IsOutgoing
}

// IsEditedMessage returns true if the message is edited.
func (m *Message) IsEditedMessage() bool {
	return m != nil && m.EditDate > 0
}

// IsChannel checks if the message is from a channel.
func (m *Message) IsChannel() bool {
	return m != nil && isSupergroupOrChannelID(m.ChatId) && m.IsChannelPost
}

func isSupergroupOrChannelID(id int64) bool {
	return (id < supergroupChatIDOffset && id > maxSecretChatID) ||
		(id < minSecretChatID && id >= -4000000000000)
}

// GetText returns the message text, for both text messages and media messages
func (m *Message) GetText() string {
	if m.Caption() != "" {
		return m.Caption()
	}

	return m.Text()
}

// GetFormattedText extracts the *FormattedText from the message content if it has text.
func (m *Message) GetFormattedText() (*FormattedText, error) {
	if m != nil {
		if content, ok := formattingValue(m.Content).(MessageText); ok && content.Text != nil {
			return content.Text, nil
		}
	}
	return nil, ErrNoFormattedText
}

// GetEntities returns the message entities, for both text messages and media messages.
func (m *Message) GetEntities() []TextEntity {
	if len(m.CaptionEntities()) > 0 {
		return m.CaptionEntities()
	}

	return m.Entities()
}

// IsCommand returns true if the message is a command.
func (m *Message) IsCommand() bool {
	entities := m.Entities()
	if len(entities) == 0 {
		return false
	}

	for _, entity := range entities {
		if _, ok := entity.Type.(*TextEntityTypeBotCommand); ok && entity.Offset == 0 {
			return true
		}
	}
	return false
}

// Args returns the message arguments excluding the command itself.
func (m *Message) Args() string {
	return strings.Join(m.ArgsList(), " ")
}

// ArgsList returns all arguments excluding the command name.
func (m *Message) ArgsList() []string {
	parts := strings.Fields(m.Text())

	if len(parts) < 2 {
		return nil
	}

	return parts[1:]
}

// Arg returns the argument at the given index.
func (m *Message) Arg(index int) string {
	args := m.ArgsList()
	if index < 0 || index >= len(args) {
		return ""
	}

	return args[index]
}

// StartsWith reports whether the message starts with prefix.
func (m *Message) StartsWith(prefix string) bool {
	return strings.HasPrefix(m.GetText(), prefix)
}

// EndsWith reports whether the message ends with suffix.
func (m *Message) EndsWith(suffix string) bool {
	return strings.HasSuffix(m.GetText(), suffix)
}

// TrimmedText returns the message without surrounding spaces.
func (m *Message) TrimmedText() string {
	return strings.TrimSpace(m.GetText())
}

// Empty reports whether the message has any content.
func (m *Message) Empty() bool {
	return strings.TrimSpace(m.GetText()) == ""
}

// SenderID returns the user ID or chat ID of the sender.
func (m *Message) SenderID() int64 {
	if m == nil {
		return 0
	}
	switch sender := m.SenderId.(type) {
	case MessageSenderUser:
		return sender.UserId
	case MessageSenderChat:
		return sender.ChatId
	case *MessageSenderUser:
		if sender != nil {
			return sender.UserId
		}
	case *MessageSenderChat:
		if sender != nil {
			return sender.ChatId
		}
	}
	return 0
}

// ChatID returns the chat ID.
func (m *Message) ChatID() int64 {
	if m == nil {
		return 0
	}
	return m.ChatId
}

// ReplyToMessageID returns the ID of the replied message.
func (m *Message) ReplyToMessageID() int64 {
	if m != nil {
		if reply, ok := m.ReplyTo.(*MessageReplyToMessage); ok && reply != nil {
			return reply.MessageId
		}
	}
	return 0
}

// Text returns the text of the message.
func (m *Message) Text() string {
	if text, err := m.GetFormattedText(); err == nil {
		return text.Text
	}
	return ""
}

// Entities returns the entities of the message text.
func (m *Message) Entities() []TextEntity {
	if text, err := m.GetFormattedText(); err == nil {
		return text.Entities
	}
	return nil
}

// Caption returns the caption of the message.
func (m *Message) Caption() string {
	if caption, err := m.GetFormattedCaption(); err == nil {
		return caption.Text
	}
	return ""
}

// GetFormattedCaption extracts the *FormattedText from the message content if it has a caption.
func (m *Message) GetFormattedCaption() (*FormattedText, error) {
	if m == nil {
		return nil, ErrNoFormattedText
	}
	var caption *FormattedText
	switch content := formattingValue(m.Content).(type) {
	case MessageAnimation:
		caption = content.Caption
	case MessageAudio:
		caption = content.Caption
	case MessageDocument:
		caption = content.Caption
	case MessagePhoto:
		caption = content.Caption
	case MessageVideo:
		caption = content.Caption
	case MessageVoiceNote:
		caption = content.Caption
	case MessagePaidMedia:
		caption = content.Caption
	}
	if caption == nil {
		return nil, ErrNoFormattedText
	}
	return caption, nil
}

// CaptionEntities returns the entities of the message caption.
func (m *Message) CaptionEntities() []TextEntity {
	if caption, err := m.GetFormattedCaption(); err == nil {
		return caption.Entities
	}
	return nil
}
