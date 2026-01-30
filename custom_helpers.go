package gotdbot

import (
	"os"
	"strings"
)

// ParseText parses the text using the specified parse mode.
func (c *Client) ParseText(text string, parseMode string) (*FormattedText, error) {
	var mode *TextParseMode

	switch strings.ToLower(parseMode) {
	case "markdown":
		mode = &TextParseMode{
			TextParseModeMarkdown: &TextParseModeMarkdown{Version: 1},
		}
	case "markdownv2":
		mode = &TextParseMode{
			TextParseModeMarkdown: &TextParseModeMarkdown{Version: 2},
		}
	case "html":
		mode = &TextParseMode{
			TextParseModeHTML: &TextParseModeHTML{},
		}
	default:
		return &FormattedText{Text: text}, nil
	}

	return c.ParseTextEntities(text, mode)
}

func GetFormattedText(c *Client, text string, entities []*TextEntity, parseMode string) *FormattedText {
	if len(entities) > 0 {
		return &FormattedText{
			Text:     text,
			Entities: entities,
		}
	} else if parseMode != "" {
		ft, err := c.ParseText(text, parseMode)
		if err == nil {
			return ft
		}
	}
	return &FormattedText{Text: text}
}

func GetInputFile(path string) *InputFile {
	if _, err := os.Stat(path); err == nil {
		return &InputFile{
			InputFileLocal: &InputFileLocal{Path: path},
		}
	}

	return &InputFile{
		InputFileRemote: &InputFileRemote{Id: path},
	}
}

// FromID returns the user ID or chat ID of the sender.
func (t *Message) FromID() int64 {
	if t.SenderId == nil {
		return 0
	}
	if t.SenderId.MessageSenderUser != nil {
		return t.SenderId.MessageSenderUser.UserId
	}
	if t.SenderId.MessageSenderChat != nil {
		return t.SenderId.MessageSenderChat.ChatId
	}
	return 0
}

// ChatID returns the chat ID.
func (t *Message) ChatID() int64 {
	return t.ChatId
}

// ReplyToMessageID returns the ID of the replied message.
func (t *Message) ReplyToMessageID() int64 {
	if t.ReplyTo == nil {
		return 0
	}
	if t.ReplyTo.MessageReplyToMessage != nil {
		return t.ReplyTo.MessageReplyToMessage.MessageId
	}
	return 0
}

// Text returns the text of the message.
func (t *Message) Text() string {
	if t.Content == nil {
		return ""
	}
	if t.Content.MessageText != nil {
		return t.Content.MessageText.Text.Text
	}
	return ""
}

// Entities returns the entities of the message text.
func (t *Message) Entities() []*TextEntity {
	if t.Content != nil && t.Content.MessageText != nil {
		return t.Content.MessageText.Text.Entities
	}
	return nil
}

// Caption returns the caption of the message.
func (t *Message) Caption() string {
	if t.Content == nil {
		return ""
	}
	if t.Content.MessagePhoto != nil {
		return t.Content.MessagePhoto.Caption.Text
	}
	if t.Content.MessageVideo != nil {
		return t.Content.MessageVideo.Caption.Text
	}
	if t.Content.MessageAnimation != nil {
		return t.Content.MessageAnimation.Caption.Text
	}
	if t.Content.MessageAudio != nil {
		return t.Content.MessageAudio.Caption.Text
	}
	if t.Content.MessageDocument != nil {
		return t.Content.MessageDocument.Caption.Text
	}
	if t.Content.MessageVoiceNote != nil {
		return t.Content.MessageVoiceNote.Caption.Text
	}
	return ""
}

// CaptionEntities returns the entities of the message caption.
func (t *Message) CaptionEntities() []*TextEntity {
	if t.Content == nil {
		return nil
	}
	if t.Content.MessagePhoto != nil {
		return t.Content.MessagePhoto.Caption.Entities
	}
	if t.Content.MessageVideo != nil {
		return t.Content.MessageVideo.Caption.Entities
	}
	if t.Content.MessageAnimation != nil {
		return t.Content.MessageAnimation.Caption.Entities
	}
	if t.Content.MessageAudio != nil {
		return t.Content.MessageAudio.Caption.Entities
	}
	if t.Content.MessageDocument != nil {
		return t.Content.MessageDocument.Caption.Entities
	}
	if t.Content.MessageVoiceNote != nil {
		return t.Content.MessageVoiceNote.Caption.Entities
	}
	return nil
}

// RemoteFileID returns the remote file ID.
func (t *Message) RemoteFileID() string {
	if t.Content == nil {
		return ""
	}
	getFileId := func(f *File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.Id
		}
		return ""
	}

	if t.Content.MessagePhoto != nil && len(t.Content.MessagePhoto.Photo.Sizes) > 0 {
		return getFileId(t.Content.MessagePhoto.Photo.Sizes[len(t.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	}
	if t.Content.MessageVideo != nil {
		return getFileId(t.Content.MessageVideo.Video.Video)
	}
	if t.Content.MessageSticker != nil {
		return getFileId(t.Content.MessageSticker.Sticker.Sticker)
	}
	if t.Content.MessageAnimation != nil {
		return getFileId(t.Content.MessageAnimation.Animation.Animation)
	}
	if t.Content.MessageAudio != nil {
		return getFileId(t.Content.MessageAudio.Audio.Audio)
	}
	if t.Content.MessageDocument != nil {
		return getFileId(t.Content.MessageDocument.Document.Document)
	}
	if t.Content.MessageVoiceNote != nil {
		return getFileId(t.Content.MessageVoiceNote.VoiceNote.Voice)
	}
	if t.Content.MessageVideoNote != nil {
		return getFileId(t.Content.MessageVideoNote.VideoNote.Video)
	}
	return ""
}

// RemoteUniqueFileID returns the remote unique file ID.
func (t *Message) RemoteUniqueFileID() string {
	if t.Content == nil {
		return ""
	}
	getUniqueId := func(f *File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.UniqueId
		}
		return ""
	}

	if t.Content.MessagePhoto != nil && len(t.Content.MessagePhoto.Photo.Sizes) > 0 {
		return getUniqueId(t.Content.MessagePhoto.Photo.Sizes[len(t.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	}
	if t.Content.MessageVideo != nil {
		return getUniqueId(t.Content.MessageVideo.Video.Video)
	}
	if t.Content.MessageSticker != nil {
		return getUniqueId(t.Content.MessageSticker.Sticker.Sticker)
	}
	if t.Content.MessageAnimation != nil {
		return getUniqueId(t.Content.MessageAnimation.Animation.Animation)
	}
	if t.Content.MessageAudio != nil {
		return getUniqueId(t.Content.MessageAudio.Audio.Audio)
	}
	if t.Content.MessageDocument != nil {
		return getUniqueId(t.Content.MessageDocument.Document.Document)
	}
	if t.Content.MessageVoiceNote != nil {
		return getUniqueId(t.Content.MessageVoiceNote.VoiceNote.Voice)
	}
	if t.Content.MessageVideoNote != nil {
		return getUniqueId(t.Content.MessageVideoNote.VideoNote.Video)
	}
	return ""
}

// Delete deletes the message.
func (t *Message) Delete(c *Client, revoke bool) error {
	_, err := c.DeleteMessages(t.ChatId, []int64{t.Id}, revoke)
	return err
}

// Pin pins the message.
func (t *Message) Pin(c *Client, disableNotification bool, onlyForSelf bool) error {
	_, err := c.PinChatMessage(t.ChatId, t.Id, disableNotification, onlyForSelf)
	return err
}

// Unpin unpins the message.
func (t *Message) Unpin(c *Client) error {
	_, err := c.UnpinChatMessage(t.ChatId, t.Id)
	return err
}

// GetChat returns information about the chat where the message was sent.
func (t *Message) GetChat(c *Client) (*Chat, error) {
	return c.GetChat(t.ChatId)
}

// GetUser returns information about the sender of the message.
func (t *Message) GetUser(c *Client) (*User, error) {
	userId := t.FromID()
	if userId == 0 {
		return nil, nil
	}
	return c.GetUser(userId)
}

// LeaveChat leaves the chat where the message was sent.
func (t *Message) LeaveChat(c *Client) error {
	_, err := c.LeaveChat(t.ChatId)
	return err
}

// Download downloads the media file attached to the message.
// Returns a bound *File or nil if no downloadable media is present.
func (t *Message) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	if t.Content == nil {
		return nil, nil
	}

	resolve := func(f *File) (*File, error) {
		if f == nil {
			return nil, nil
		}
		if f.Remote != nil {
			fi, err := c.GetRemoteFile(f.Remote.Id, nil)
			if err != nil {
				return nil, err
			}
			return fi, nil
		}
		return f, nil
	}

	var (
		fileInfo *File
		err      error
	)

	if t.Content.MessagePhoto != nil && len(t.Content.MessagePhoto.Photo.Sizes) > 0 {
		fileInfo, err = resolve(t.Content.MessagePhoto.Photo.Sizes[len(t.Content.MessagePhoto.Photo.Sizes)-1].Photo)
	} else if t.Content.MessageVideo != nil {
		fileInfo, err = resolve(t.Content.MessageVideo.Video.Video)
	} else if t.Content.MessageSticker != nil {
		fileInfo, err = resolve(t.Content.MessageSticker.Sticker.Sticker)
	} else if t.Content.MessageAnimation != nil {
		fileInfo, err = resolve(t.Content.MessageAnimation.Animation.Animation)
	} else if t.Content.MessageAudio != nil {
		fileInfo, err = resolve(t.Content.MessageAudio.Audio.Audio)
	} else if t.Content.MessageDocument != nil {
		fileInfo, err = resolve(t.Content.MessageDocument.Document.Document)
	} else if t.Content.MessageVoiceNote != nil {
		fileInfo, err = resolve(t.Content.MessageVoiceNote.VoiceNote.Voice)
	} else if t.Content.MessageVideoNote != nil {
		fileInfo, err = resolve(t.Content.MessageVideoNote.VideoNote.Video)
	}

	if err != nil {
		return nil, err
	}
	if fileInfo == nil {
		return nil, nil
	}

	return fileInfo.Download(c, priority, offset, limit, synchronous)
}

// CallbackData returns the callback data payload as bytes.
func (t *UpdateNewCallbackQuery) CallbackData() []byte {
	if t.Payload == nil {
		return nil
	}
	if t.Payload.CallbackQueryPayloadData != nil {
		return t.Payload.CallbackQueryPayloadData.Data
	}
	return nil
}

// GetMessage returns the message that originated the query.
func (t *UpdateNewCallbackQuery) GetMessage(c *Client) (*Message, error) {
	if t.MessageId == 0 {
		return nil, nil
	}
	return c.GetMessage(t.ChatId, t.MessageId)
}

// Answer sends an answer to the callback query.
func (t *UpdateNewCallbackQuery) Answer(c *Client, text string, showAlert bool, url string, cacheTime int32) error {
	_, err := c.AnswerCallbackQuery(t.Id, text, showAlert, url, cacheTime)
	return err
}
