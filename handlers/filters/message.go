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
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageText)
	return ok
}

var Photo Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessagePhoto)
	return ok
}

var Video Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageVideo)
	return ok
}

var Animation Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageAnimation)
	return ok
}

var Audio Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageAudio)
	return ok
}

var Document Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageDocument)
	return ok
}

var Sticker Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageSticker)
	return ok
}

var VideoNote Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageVideoNote)
	return ok
}

var VoiceNote Message = func(msg *gotdbot.Message) bool {
	if msg.Content == nil {
		return false
	}
	_, ok := msg.Content.(*gotdbot.MessageVoiceNote)
	return ok
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
		if u, ok := msg.SenderId.(*gotdbot.MessageSenderUser); ok {
			return u.UserId == id
		}
		if c, ok := msg.SenderId.(*gotdbot.MessageSenderChat); ok {
			return c.ChatId == id
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
		if msg.Content == nil {
			return false
		}
		txt, ok := msg.Content.(*gotdbot.MessageText)
		if !ok || txt.Text == nil {
			return false
		}

		text := txt.Text.Text
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
		if msg.Content == nil {
			return false
		}
		txt, ok := msg.Content.(*gotdbot.MessageText)
		if !ok || txt.Text == nil {
			return false
		}
		return strings.Contains(txt.Text.Text, text)
	}
}

func HasPrefix(prefix string) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil {
			return false
		}
		txt, ok := msg.Content.(*gotdbot.MessageText)
		if !ok || txt.Text == nil {
			return false
		}
		return strings.HasPrefix(txt.Text.Text, prefix)
	}
}

func Equal(text string) Message {
	return func(msg *gotdbot.Message) bool {
		if msg.Content == nil {
			return false
		}
		txt, ok := msg.Content.(*gotdbot.MessageText)
		if !ok || txt.Text == nil {
			return false
		}
		return txt.Text.Text == text
	}
}
