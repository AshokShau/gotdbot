package filters

import (
	"strings"

	"github.com/AshokShau/gotdbot"
)

var All Message = func(_ *gotdbot.Message) bool {
	return true
}

var Incoming Message = func(msg *gotdbot.Message) bool {
	return !msg.IsOutgoing
}

var Outgoing Message = func(msg *gotdbot.Message) bool {
	return msg.IsOutgoing
}

var Text Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageText != nil
}

var Photo Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessagePhoto != nil
}

var Video Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageVideo != nil
}

var Animation Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageAnimation != nil
}

var Audio Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageAudio != nil
}

var Document Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageDocument != nil
}

var Sticker Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageSticker != nil
}

var VideoNote Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageVideoNote != nil
}

var VoiceNote Message = func(msg *gotdbot.Message) bool {
	return msg.Content != nil && msg.Content.MessageVoiceNote != nil
}

var Private Message = func(msg *gotdbot.Message) bool {
	return msg.ChatId > 0
}

func ChatID(id int64) Message {
	return func(msg *gotdbot.Message) bool {
		return msg.ChatId == id
	}
}

func SenderID(id int64) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.SenderId == nil {
			return false
		}
		if msg.SenderId.MessageSenderUser != nil {
			return msg.SenderId.MessageSenderUser.UserId == id
		}
		if msg.SenderId.MessageSenderChat != nil {
			return msg.SenderId.MessageSenderChat.ChatId == id
		}
		return false
	}
}

// Command checks if the message is a command.
func Command(command string, prefixes ...rune) Message {
	if len(prefixes) == 0 {
		prefixes = []rune{'/'}
	}
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil || msg.Content.MessageText == nil {
			return false
		}
		text := msg.Content.MessageText.Text.Text
		if text == "" {
			return false
		}

		for _, prefix := range prefixes {
			if strings.HasPrefix(text, string(prefix)+command) {
				return true
			}
		}
		return false
	}
}

func Contains(text string) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil || msg.Content.MessageText == nil {
			return false
		}
		return strings.Contains(msg.Content.MessageText.Text.Text, text)
	}
}

func HasPrefix(prefix string) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil || msg.Content.MessageText == nil {
			return false
		}
		return strings.HasPrefix(msg.Content.MessageText.Text.Text, prefix)
	}
}

func Equal(text string) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil || msg.Content.MessageText == nil {
			return false
		}
		return msg.Content.MessageText.Text.Text == text
	}
}
