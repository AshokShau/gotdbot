package client

import "github.com/AshokShau/gotdbot/types"

// Message is a wrapper around types.Message with bound methods and helper getters.
type Message struct {
	*types.Message
	client *Client
}

// NewMessage creates a new bound Message.
func NewMessage(client *Client, msg *types.Message) *Message {
	return &Message{
		Message: msg,
		client:  client,
	}
}

// FromID returns the user ID or chat ID of the sender.
func (m *Message) FromID() int64 {
	if m.SenderId == nil {
		return 0
	}
	if m.SenderId.MessageSenderUser != nil {
		return m.SenderId.MessageSenderUser.UserId
	}
	if m.SenderId.MessageSenderChat != nil {
		return m.SenderId.MessageSenderChat.ChatId
	}
	return 0
}

// ChatID returns the chat ID.
func (m *Message) ChatID() int64 {
	return m.ChatId
}

// ReplyToMessageID returns the ID of the replied message.
func (m *Message) ReplyToMessageID() int64 {
	if m.ReplyTo == nil {
		return 0
	}
	if m.ReplyTo.MessageReplyToMessage != nil {
		return m.ReplyTo.MessageReplyToMessage.MessageId
	}
	return 0
}

// Text returns the text of the message.
func (m *Message) Text() string {
	if m.Content == nil {
		return ""
	}
	if m.Content.MessageText != nil {
		return m.Content.MessageText.Text.Text
	}
	return ""
}

// Entities returns the entities of the message text.
func (m *Message) Entities() []*types.TextEntity {
	if m.Content != nil && m.Content.MessageText != nil {
		return m.Content.MessageText.Text.Entities
	}
	return nil
}

// Caption returns the caption of the message.
func (m *Message) Caption() string {
	if m.Content == nil {
		return ""
	}
	if m.Content.MessagePhoto != nil {
		return m.Content.MessagePhoto.Caption.Text
	}
	if m.Content.MessageVideo != nil {
		return m.Content.MessageVideo.Caption.Text
	}
	if m.Content.MessageAnimation != nil {
		return m.Content.MessageAnimation.Caption.Text
	}
	if m.Content.MessageAudio != nil {
		return m.Content.MessageAudio.Caption.Text
	}
	if m.Content.MessageDocument != nil {
		return m.Content.MessageDocument.Caption.Text
	}
	if m.Content.MessageVoiceNote != nil {
		return m.Content.MessageVoiceNote.Caption.Text
	}
	return ""
}

// CaptionEntities returns the entities of the message caption.
func (m *Message) CaptionEntities() []*types.TextEntity {
	if m.Content == nil {
		return nil
	}
	if m.Content.MessagePhoto != nil {
		return m.Content.MessagePhoto.Caption.Entities
	}
	if m.Content.MessageVideo != nil {
		return m.Content.MessageVideo.Caption.Entities
	}
	if m.Content.MessageAnimation != nil {
		return m.Content.MessageAnimation.Caption.Entities
	}
	if m.Content.MessageAudio != nil {
		return m.Content.MessageAudio.Caption.Entities
	}
	if m.Content.MessageDocument != nil {
		return m.Content.MessageDocument.Caption.Entities
	}
	if m.Content.MessageVoiceNote != nil {
		return m.Content.MessageVoiceNote.Caption.Entities
	}
	return nil
}

// RemoteFileID returns the remote file ID.
func (m *Message) RemoteFileID() string {
	if m.Content == nil {
		return ""
	}
	getFileId := func(f *types.File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.Id
		}
		return ""
	}

	if m.Content.MessagePhoto != nil && len(m.Content.MessagePhoto.Photo.Sizes) > 0 {
		return getFileId(m.Content.MessagePhoto.Photo.Sizes[len(m.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	}
	if m.Content.MessageVideo != nil {
		return getFileId(m.Content.MessageVideo.Video.Video)
	}
	if m.Content.MessageSticker != nil {
		return getFileId(m.Content.MessageSticker.Sticker.Sticker)
	}
	if m.Content.MessageAnimation != nil {
		return getFileId(m.Content.MessageAnimation.Animation.Animation)
	}
	if m.Content.MessageAudio != nil {
		return getFileId(m.Content.MessageAudio.Audio.Audio)
	}
	if m.Content.MessageDocument != nil {
		return getFileId(m.Content.MessageDocument.Document.Document)
	}
	if m.Content.MessageVoiceNote != nil {
		return getFileId(m.Content.MessageVoiceNote.VoiceNote.Voice)
	}
	if m.Content.MessageVideoNote != nil {
		return getFileId(m.Content.MessageVideoNote.VideoNote.Video)
	}
	return ""
}

// RemoteUniqueFileID returns the remote unique file ID.
func (m *Message) RemoteUniqueFileID() string {
	if m.Content == nil {
		return ""
	}
	getUniqueId := func(f *types.File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.UniqueId
		}
		return ""
	}

	if m.Content.MessagePhoto != nil && len(m.Content.MessagePhoto.Photo.Sizes) > 0 {
		return getUniqueId(m.Content.MessagePhoto.Photo.Sizes[len(m.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	}
	if m.Content.MessageVideo != nil {
		return getUniqueId(m.Content.MessageVideo.Video.Video)
	}
	if m.Content.MessageSticker != nil {
		return getUniqueId(m.Content.MessageSticker.Sticker.Sticker)
	}
	if m.Content.MessageAnimation != nil {
		return getUniqueId(m.Content.MessageAnimation.Animation.Animation)
	}
	if m.Content.MessageAudio != nil {
		return getUniqueId(m.Content.MessageAudio.Audio.Audio)
	}
	if m.Content.MessageDocument != nil {
		return getUniqueId(m.Content.MessageDocument.Document.Document)
	}
	if m.Content.MessageVoiceNote != nil {
		return getUniqueId(m.Content.MessageVoiceNote.VoiceNote.Voice)
	}
	if m.Content.MessageVideoNote != nil {
		return getUniqueId(m.Content.MessageVideoNote.VideoNote.Video)
	}
	return ""
}

// Delete deletes the message.
func (m *Message) Delete(revoke bool) error {
	_, err := m.client.DeleteMessages(m.ChatId, []int64{m.Id}, revoke)
	return err
}

// Pin pins the message.
func (m *Message) Pin(disableNotification bool, onlyForSelf bool) error {
	_, err := m.client.PinChatMessage(m.ChatId, m.Id, disableNotification, onlyForSelf)
	return err
}

// Unpin unpins the message.
func (m *Message) Unpin() error {
	_, err := m.client.UnpinChatMessage(m.ChatId, m.Id)
	return err
}

// GetChat returns information about the chat where the message was sent.
func (m *Message) GetChat() (*types.Chat, error) {
	return m.client.GetChat(m.ChatId)
}

// GetUser returns information about the sender of the message.
func (m *Message) GetUser() (*types.User, error) {
	userId := m.FromID()
	if userId == 0 {
		return nil, nil
	}
	return m.client.GetUser(userId)
}

// LeaveChat leaves the chat where the message was sent.
func (m *Message) LeaveChat() error {
	_, err := m.client.LeaveChat(m.ChatId)
	return err
}

// Download downloads the media file attached to the message.
// Returns a bound *File or nil if no downloadable media is present.
func (m *Message) Download(priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	if m.Content == nil {
		return nil, nil
	}

	resolve := func(f *types.File) (*types.File, error) {
		if f == nil {
			return nil, nil
		}
		if f.Remote != nil {
			fi, err := m.client.GetRemoteFile(f.Remote.Id, nil)
			if err != nil {
				return nil, err
			}
			return fi, nil
		}
		return f, nil
	}

	var (
		fileInfo *types.File
		err      error
	)

	if m.Content.MessagePhoto != nil && len(m.Content.MessagePhoto.Photo.Sizes) > 0 {
		fileInfo, err = resolve(m.Content.MessagePhoto.Photo.Sizes[len(m.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	} else if m.Content.MessageVideo != nil {
		fileInfo, err = resolve(m.Content.MessageVideo.Video.Video)
	} else if m.Content.MessageSticker != nil {
		fileInfo, err = resolve(m.Content.MessageSticker.Sticker.Sticker)
	} else if m.Content.MessageAnimation != nil {
		fileInfo, err = resolve(m.Content.MessageAnimation.Animation.Animation)
	} else if m.Content.MessageAudio != nil {
		fileInfo, err = resolve(m.Content.MessageAudio.Audio.Audio)
	} else if m.Content.MessageDocument != nil {
		fileInfo, err = resolve(m.Content.MessageDocument.Document.Document)
	} else if m.Content.MessageVoiceNote != nil {
		fileInfo, err = resolve(m.Content.MessageVoiceNote.VoiceNote.Voice)
	} else if m.Content.MessageVideoNote != nil {
		fileInfo, err = resolve(m.Content.MessageVideoNote.VideoNote.Video)
	}

	if err != nil {
		return nil, err
	}
	if fileInfo == nil {
		return nil, nil
	}

	return NewFile(m.client, fileInfo).Download(priority, offset, limit, synchronous)
}
