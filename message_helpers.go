package gotdbot

import (
	"strings"
)

// FromID returns the user ID or chat ID of the sender.
func (m *Message) FromID() int64 {
	if m.SenderId == nil {
		return 0
	}
	if u, ok := m.SenderId.(*MessageSenderUser); ok {
		return u.UserId
	}
	if c, ok := m.SenderId.(*MessageSenderChat); ok {
		return c.ChatId
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
	if r, ok := m.ReplyTo.(*MessageReplyToMessage); ok {
		return r.MessageId
	}
	return 0
}

// Text returns the text of the message.
func (m *Message) Text() string {
	if m.Content == nil {
		return ""
	}
	if c, ok := m.Content.(*MessageText); ok {
		return c.Text.Text
	}
	return ""
}

// Entities returns the entities of the message text.
func (m *Message) Entities() []*TextEntity {
	if m.Content == nil {
		return nil
	}
	if c, ok := m.Content.(*MessageText); ok {
		return c.Text.Entities
	}
	return nil
}

// Caption returns the caption of the message.
func (m *Message) Caption() string {
	if m.Content == nil {
		return ""
	}
	if c, ok := m.Content.(*MessagePhoto); ok {
		return c.Caption.Text
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		return c.Caption.Text
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		return c.Caption.Text
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		return c.Caption.Text
	}
	if c, ok := m.Content.(*MessageDocument); ok {
		return c.Caption.Text
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		return c.Caption.Text
	}
	return ""
}

// CaptionEntities returns the entities of the message caption.
func (m *Message) CaptionEntities() []*TextEntity {
	if m.Content == nil {
		return nil
	}
	if c, ok := m.Content.(*MessagePhoto); ok {
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageDocument); ok {
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		return c.Caption.Entities
	}
	return nil
}

// Delete deletes the message.
func (m *Message) Delete(c *Client, revoke bool) error {
	_, err := c.DeleteMessages(m.ChatId, []int64{m.Id}, revoke)
	return err
}

// Pin pins the message.
func (m *Message) Pin(c *Client, disableNotification bool, onlyForSelf bool) error {
	_, err := c.PinChatMessage(m.ChatId, m.Id, disableNotification, onlyForSelf)
	return err
}

// Unpin unpins the message.
func (m *Message) Unpin(c *Client) error {
	_, err := c.UnpinChatMessage(m.ChatId, m.Id)
	return err
}

// GetChat returns information about the chat where the message was sent.
func (m *Message) GetChat(c *Client) (*Chat, error) {
	return c.GetChat(m.ChatId)
}

// GetUser returns information about the sender of the message.
func (m *Message) GetUser(c *Client) (*User, error) {
	userId := m.FromID()
	if userId == 0 {
		return nil, nil
	}
	return c.GetUser(userId)
}

// LeaveChat leaves the chat where the message was sent.
func (m *Message) LeaveChat(c *Client) error {
	_, err := c.LeaveChat(m.ChatId)
	return err
}

// RemoteFileID returns the remote file ID.
func (m *Message) RemoteFileID() string {
	if m.Content == nil {
		return ""
	}
	getFileId := func(f *File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.Id
		}
		return ""
	}

	if c, ok := m.Content.(*MessagePhoto); ok && len(c.Photo.Sizes) > 0 {
		return getFileId(c.Photo.Sizes[len(c.Photo.Sizes)-1].Photo)
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		return getFileId(c.Video.Video)
	}
	if c, ok := m.Content.(*MessageSticker); ok {
		return getFileId(c.Sticker.Sticker)
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		return getFileId(c.Animation.Animation)
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		return getFileId(c.Audio.Audio)
	}
	if c, ok := m.Content.(*MessageDocument); ok {
		return getFileId(c.Document.Document)
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		return getFileId(c.VoiceNote.Voice)
	}
	if c, ok := m.Content.(*MessageVideoNote); ok {
		return getFileId(c.VideoNote.Video)
	}
	return ""
}

// RemoteUniqueFileID returns the remote unique file ID.
func (m *Message) RemoteUniqueFileID() string {
	if m.Content == nil {
		return ""
	}
	getUniqueId := func(f *File) string {
		if f != nil && f.Remote != nil {
			return f.Remote.UniqueId
		}
		return ""
	}

	if c, ok := m.Content.(*MessagePhoto); ok && len(c.Photo.Sizes) > 0 {
		return getUniqueId(c.Photo.Sizes[len(c.Photo.Sizes)-1].Photo)
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		return getUniqueId(c.Video.Video)
	}
	if c, ok := m.Content.(*MessageSticker); ok {
		return getUniqueId(c.Sticker.Sticker)
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		return getUniqueId(c.Animation.Animation)
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		return getUniqueId(c.Audio.Audio)
	}
	if c, ok := m.Content.(*MessageDocument); ok {
		return getUniqueId(c.Document.Document)
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		return getUniqueId(c.VoiceNote.Voice)
	}
	if c, ok := m.Content.(*MessageVideoNote); ok {
		return getUniqueId(c.VideoNote.Video)
	}
	return ""
}

// Download downloads the media file attached to the message.
// Returns a bound *File or nil if no downloadable media is present.
func (m *Message) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	if m.Content == nil {
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

	if c, ok := m.Content.(*MessagePhoto); ok && len(c.Photo.Sizes) > 0 {
		fileInfo, err = resolve(c.Photo.Sizes[len(c.Photo.Sizes)-1].Photo)
	} else if c, ok := m.Content.(*MessageVideo); ok {
		fileInfo, err = resolve(c.Video.Video)
	} else if c, ok := m.Content.(*MessageSticker); ok {
		fileInfo, err = resolve(c.Sticker.Sticker)
	} else if c, ok := m.Content.(*MessageAnimation); ok {
		fileInfo, err = resolve(c.Animation.Animation)
	} else if c, ok := m.Content.(*MessageAudio); ok {
		fileInfo, err = resolve(c.Audio.Audio)
	} else if c, ok := m.Content.(*MessageDocument); ok {
		fileInfo, err = resolve(c.Document.Document)
	} else if c, ok := m.Content.(*MessageVoiceNote); ok {
		fileInfo, err = resolve(c.VoiceNote.Voice)
	} else if c, ok := m.Content.(*MessageVideoNote); ok {
		fileInfo, err = resolve(c.VideoNote.Video)
	}

	if err != nil {
		return nil, err
	}
	if fileInfo == nil {
		return nil, nil
	}

	return fileInfo.Download(c, priority, offset, limit, synchronous)
}

// Mention returns the text mention of the message sender.
func (m *Message) Mention(c *Client, parseMode string) (string, error) {
	chat, err := c.GetChat(m.FromID())
	if err != nil {
		return "", err
	}
	html := strings.ToLower(parseMode) == "html"
	return Mention(chat.Title, m.FromID(), html, true), nil
}

// GetMessageProperties returns the message properties.
func (m *Message) GetMessageProperties(c *Client) (*MessageProperties, error) {
	return c.GetMessageProperties(m.ChatId, m.Id)
}

// GetMessageLink returns the message link.
func (m *Message) GetMessageLink(c *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return c.GetMessageLink(m.ChatId, m.Id, mediaTimestamp, forAlbum, inMessageThread)
}

// GetRepliedMessage returns the replied message.
func (m *Message) GetRepliedMessage(c *Client) (*Message, error) {
	return c.GetRepliedMessage(m.ChatId, m.Id)
}

// GetChatMember returns member info in the current chat.
func (m *Message) GetChatMember(c *Client) (*ChatMember, error) {
	return c.GetChatMember(m.ChatId, m.SenderId)
}

// SetChatMemberStatus sets chat member status.
func (m *Message) SetChatMemberStatus(c *Client, status ChatMemberStatus) (*Ok, error) {
	return c.SetChatMemberStatus(m.ChatId, m.SenderId, status)
}

// Ban bans the message sender.
func (m *Message) Ban(c *Client, bannedUntilDate int32) (*Ok, error) {
	return m.SetChatMemberStatus(c, &ChatMemberStatusBanned{
		BannedUntilDate: bannedUntilDate,
	})
}

// Kick kicks the message sender.
func (m *Message) Kick(c *Client) (*Ok, error) {
	return m.SetChatMemberStatus(c, &ChatMemberStatusLeft{})
}

// Restrict restricts the message sender.
func (m *Message) Restrict(c *Client, permissions *ChatPermissions, restrictedUntilDate int32) (*Ok, error) {
	return m.SetChatMemberStatus(c, &ChatMemberStatusRestricted{
		IsMember:            true,
		Permissions:         permissions,
		RestrictedUntilDate: restrictedUntilDate,
	})
}

// React reacts to the current message.
func (m *Message) React(c *Client, emoji string, isBig bool) (*Ok, error) {
	var reactionTypes []ReactionType
	if emoji != "" {
		reactionTypes = []ReactionType{
			&ReactionTypeEmoji{
				Emoji: emoji,
			},
		}
	}
	return c.SetMessageReactions(m.ChatId, m.Id, reactionTypes, isBig)
}

// Action sends a chat action to a specific chat.
func (m *Message) Action(c *Client, action string, topicId MessageTopic) (*ChatActionSender, error) {
	return NewChatAction(c, m.ChatId, action, topicId)
}

// ReplyText replies to the message with text.
func (m *Message) ReplyText(c *Client, text string, opts *SendTextMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendTextMessageOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendTextMessage(m.ChatId, text, opts)
}

// ReplyAnimation replies to the message with animation.
func (m *Message) ReplyAnimation(c *Client, animation string, opts *SendAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAnimationOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAnimation(m.ChatId, animation, opts)
}

// ReplyAudio replies to the message with audio.
func (m *Message) ReplyAudio(c *Client, audio string, opts *SendAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAudioOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAudio(m.ChatId, audio, opts)
}

// ReplyDocument replies to the message with a document.
func (m *Message) ReplyDocument(c *Client, document string, opts *SendDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDocumentOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendDocument(m.ChatId, document, opts)
}

// ReplyPhoto replies to the message with a photo.
func (m *Message) ReplyPhoto(c *Client, photo string, opts *SendPhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPhotoOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPhoto(m.ChatId, photo, opts)
}

// ReplyVideo replies to the message with a video.
func (m *Message) ReplyVideo(c *Client, video string, opts *SendVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideo(m.ChatId, video, opts)
}

// ReplyVideoNote replies to the message with a video note.
func (m *Message) ReplyVideoNote(c *Client, videoNote string, opts *SendVideoNoteOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoNoteOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideoNote(m.ChatId, videoNote, opts)
}

// ReplyVoice replies to the message with a voice note.
func (m *Message) ReplyVoice(c *Client, voice string, opts *SendVoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVoiceOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVoice(m.ChatId, voice, opts)
}

// ReplySticker replies to the message with a sticker.
func (m *Message) ReplySticker(c *Client, sticker string, opts *SendStickerOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStickerOpts{}
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendSticker(m.ChatId, sticker, opts)
}

// Copy copies message to chat.
func (m *Message) Copy(c *Client, chatId int64, opts *SendCopyOpts) (*Message, error) {
	return c.SendCopy(chatId, m.ChatId, m.Id, opts)
}

// Forward forwards message to chat.
func (m *Message) Forward(c *Client, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return c.ForwardMessage(chatId, m.ChatId, m.Id, opts)
}

// EditText edits a text message.
func (m *Message) EditText(c *Client, text string, opts *EditTextMessageOpts) (*Message, error) {
	return c.EditTextMessage(m.ChatId, m.Id, text, opts)
}
