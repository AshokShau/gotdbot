package gotdbot

import (
	"strings"
)

// GetLink returns the message link.
func (m *Message) GetLink(c *Client) (*MessageLink, error) {
	return c.GetMessageLink(m.ChatId, 0, 0, m.Id, "", &GetMessageLinkOpts{})
}

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
	return m.IsOutgoing
}

// IsEditedMessage returns true if the message is edited.
func (m *Message) IsEditedMessage() bool {
	return m.EditDate > 0
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
	if m == nil || m.Content == nil {
		return nil, ErrNoFormattedText
	}
	switch content := m.Content.(type) {
	case *MessageText:
		return content.Text, nil
	default:
		return nil, ErrNoFormattedText
	}
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
func (m *Message) Entities() []TextEntity {
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

// GetFormattedCaption extracts the *FormattedText from the message content if it has a caption.
func (m *Message) GetFormattedCaption() (*FormattedText, error) {
	if m == nil || m.Content == nil {
		return nil, ErrNoFormattedText
	}
	switch content := m.Content.(type) {
	case *MessageAnimation:
		return content.Caption, nil
	case *MessageAudio:
		return content.Caption, nil
	case *MessageDocument:
		return content.Caption, nil
	case *MessagePhoto:
		return content.Caption, nil
	case *MessageVideo:
		return content.Caption, nil
	case *MessageVoiceNote:
		return content.Caption, nil
	default:
		return nil, ErrNoFormattedText
	}
}

// CaptionEntities returns the entities of the message caption.
func (m *Message) CaptionEntities() []TextEntity {
	if m.Content == nil {
		return nil
	}
	if c, ok := m.Content.(*MessagePhoto); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageDocument); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		if c.Caption == nil {
			return nil
		}
		return c.Caption.Entities
	}
	return nil
}

// Delete deletes the message.
func (m *Message) Delete(c *Client, revoke bool) error {
	err := c.DeleteMessages(m.ChatId, []int64{m.Id}, &DeleteMessagesOpts{Revoke: revoke})
	return err
}

// Pin pins the message.
func (m *Message) Pin(c *Client, disableNotification bool, onlyForSelf bool) error {
	err := c.PinChatMessage(m.ChatId, m.Id, &PinChatMessageOpts{DisableNotification: disableNotification, OnlyForSelf: onlyForSelf})
	return err
}

// Unpin unpins the message.
func (m *Message) Unpin(c *Client) error {
	err := c.UnpinChatMessage(m.ChatId, m.Id)
	return err
}

// GetChat returns information about the chat where the message was sent.
func (m *Message) GetChat(c *Client) (*Chat, error) {
	return c.GetChat(m.ChatId)
}

// GetUser returns information about the sender of the message.
func (m *Message) GetUser(c *Client) (*User, error) {
	userId := m.SenderID()
	if userId == 0 {
		return nil, nil
	}
	return c.GetUser(userId)
}

// LeaveChat leaves the chat where the message was sent.
func (m *Message) LeaveChat(c *Client) error {
	err := c.LeaveChat(m.ChatId)
	return err
}

// mediaFile extracts the primary file without duplicating media-type switches.
func (m *Message) mediaFile() *File {
	if m == nil {
		return nil
	}
	switch content := m.Content.(type) {
	case *MessagePhoto:
		if content != nil && content.Photo != nil && len(content.Photo.Sizes) > 0 {
			return content.Photo.Sizes[len(content.Photo.Sizes)-1].Photo
		}
	case *MessageVideo:
		if content != nil && content.Video != nil {
			return content.Video.Video
		}
	case *MessageSticker:
		if content != nil && content.Sticker != nil {
			return content.Sticker.Sticker
		}
	case *MessageAnimation:
		if content != nil && content.Animation != nil {
			return content.Animation.Animation
		}
	case *MessageAudio:
		if content != nil && content.Audio != nil {
			return content.Audio.Audio
		}
	case *MessageDocument:
		if content != nil && content.Document != nil {
			return content.Document.Document
		}
	case *MessageVoiceNote:
		if content != nil && content.VoiceNote != nil {
			return content.VoiceNote.Voice
		}
	case *MessageVideoNote:
		if content != nil && content.VideoNote != nil {
			return content.VideoNote.Video
		}
	}
	return nil
}

// RemoteFileID returns the remote file ID.
func (m *Message) RemoteFileID() string {
	if file := m.mediaFile(); file != nil && file.Remote != nil {
		return file.Remote.Id
	}
	return ""
}

// RemoteDuration returns the remote media duration.
func (m *Message) RemoteDuration() int32 {
	if m.Content == nil {
		return 0
	}
	if c, ok := m.Content.(*MessageVideo); ok {
		return c.Video.Duration
	}
	if c, ok := m.Content.(*MessageAnimation); ok {
		return c.Animation.Duration
	}
	if c, ok := m.Content.(*MessageAudio); ok {
		return c.Audio.Duration
	}
	if c, ok := m.Content.(*MessageVoiceNote); ok {
		return c.VoiceNote.Duration
	}
	if c, ok := m.Content.(*MessageVideoNote); ok {
		return c.VideoNote.Duration
	}
	return 0
}

// RemoteFileSize returns the size, or the expected size if not known yet.
func (m *Message) RemoteFileSize() int64 {
	if file := m.mediaFile(); file != nil {
		if file.Size > 0 {
			return file.Size
		}
		return file.ExpectedSize
	}
	return 0
}

// RemoteUniqueFileID returns the remote unique file ID.
func (m *Message) RemoteUniqueFileID() string {
	if file := m.mediaFile(); file != nil && file.Remote != nil {
		return file.Remote.UniqueId
	}
	return ""
}

// Download downloads the media file attached to the message.
// Returns a native *File or nil if no downloadable media is present.
func (m *Message) Download(c *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	fileInfo := m.mediaFile()
	if fileInfo == nil {
		return nil, nil
	}
	// A remote reference can resolve the file for the supplied client too.
	if fileInfo.Remote != nil && fileInfo.Remote.Id != "" {
		var err error
		fileInfo, err = c.GetRemoteFile(fileInfo.Remote.Id, nil)
		if err != nil {
			return nil, err
		}
	}
	return c.DownloadFile(fileInfo.Id, limit, offset, priority, &DownloadFileOpts{Synchronous: synchronous})
}

// Mention returns the text mention of the message sender.
func (m *Message) Mention(c *Client, parseMode string) (string, error) {
	chat, err := c.GetChat(m.SenderID())
	if err != nil {
		return "", err
	}
	html := strings.ToLower(parseMode) == "html"
	return Mention(chat.Title, m.SenderID(), html, true), nil
}

// GetMessageProperties returns the message properties.
func (m *Message) GetMessageProperties(c *Client) (*MessageProperties, error) {
	return c.GetMessageProperties(m.ChatId, m.Id)
}

// GetMessageLink returns the message link.
func (m *Message) GetMessageLink(c *Client, checklistTaskId int32, mediaTimestamp int32, pollOptionId string, opts *GetMessageLinkOpts) (*MessageLink, error) {
	return c.GetMessageLink(m.ChatId, checklistTaskId, mediaTimestamp, m.Id, pollOptionId, opts)
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
func (m *Message) SetChatMemberStatus(c *Client, status ChatMemberStatus) error {
	return c.SetChatMemberStatus(m.ChatId, m.SenderId, status)
}

// Ban bans the message sender.
func (m *Message) Ban(c *Client, bannedUntilDate int32) error {
	return m.SetChatMemberStatus(c, &ChatMemberStatusBanned{
		BannedUntilDate: bannedUntilDate,
	})
}

// Kick kicks the message sender.
func (m *Message) Kick(c *Client) error {
	return m.SetChatMemberStatus(c, &ChatMemberStatusLeft{})
}

// Restrict restricts the message sender.
func (m *Message) Restrict(c *Client, permissions *ChatPermissions, restrictedUntilDate int32) error {
	return m.SetChatMemberStatus(c, &ChatMemberStatusRestricted{
		IsMember:            true,
		Permissions:         permissions,
		RestrictedUntilDate: restrictedUntilDate,
	})
}

// React reacts to the current message.
func (m *Message) React(c *Client, reactionTypes []ReactionType, opts *SetMessageReactionsOpts) error {
	return c.SetMessageReactions(m.ChatId, m.Id, reactionTypes, opts)
}

// Action sends a chat action to a specific chat.
func (m *Message) Action(c *Client, opts *SendChatActionOpts) error {
	if opts == nil {
		opts = &SendChatActionOpts{Action: ChatActionTyping{}}
	}
	return c.SendChatAction("", m.ChatId, opts)
}

// ReplyText replies to the message with text.
func (m *Message) ReplyText(c *Client, text string, opts *SendMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}

	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendMessage(m.ChatId, text, opts)
}

// ReplyAlbum replies to the message with an album.
func (m *Message) ReplyAlbum(c *Client, inputMessageContents []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	if opts == nil {
		opts = &SendMessageAlbumOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyTo == nil {
		opts.ReplyTo = &InputMessageReplyToMessage{
			MessageId: m.Id,
		}
	}
	return c.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// ReplyAnimation replies to the message with animation.
func (m *Message) ReplyAnimation(c *Client, animation InputFile, opts *SendAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAnimationOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAnimation(m.ChatId, animation, opts)
}

// ReplyAudio replies to the message with audio.
func (m *Message) ReplyAudio(c *Client, audio InputFile, opts *SendAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAudioOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendAudio(m.ChatId, audio, opts)
}

// ReplyDocument replies to the message with a document.
func (m *Message) ReplyDocument(c *Client, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDocumentOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendDocument(m.ChatId, document, opts)
}

// ReplyPhoto replies to the message with a photo.
func (m *Message) ReplyPhoto(c *Client, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPhotoOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPhoto(m.ChatId, photo, opts)
}

// ReplyVideo replies to the message with a video.
func (m *Message) ReplyVideo(c *Client, video InputFile, opts *SendVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideo(m.ChatId, video, opts)
}

// ReplyVideoNote replies to the message with a video note.
func (m *Message) ReplyVideoNote(c *Client, videoNote InputFile, opts *SendVideoNoteOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoNoteOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVideoNote(m.ChatId, videoNote, opts)
}

// ReplyVoice replies to the message with a voice note.
func (m *Message) ReplyVoice(c *Client, voice InputFile, opts *SendVoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVoiceOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVoice(m.ChatId, voice, opts)
}

// ReplySticker replies to the message with a sticker.
func (m *Message) ReplySticker(c *Client, sticker InputFile, opts *SendStickerOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStickerOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendSticker(m.ChatId, sticker, opts)
}

// Copy copies message to chat.
func (m *Message) Copy(c *Client, chatId int64, opts *CopyMessageOpts) (*Message, error) {
	return c.CopyMessage(chatId, m.ChatId, m.Id, opts)
}

// Forward forwards message to chat.
func (m *Message) Forward(c *Client, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return c.ForwardMessage(chatId, m.ChatId, m.Id, opts)
}

// EditText edits a text message.
func (m *Message) EditText(c *Client, text string, opts *EditMessageTextOpts) (*Message, error) {
	return c.EditMessageText(m.ChatId, m.Id, text, opts)
}

// EditCaption edits the caption of a message.
func (m *Message) EditCaption(c *Client, caption string, opts *EditMessageCaptionOpts) (*Message, error) {
	return c.EditMessageCaption(m.ChatId, m.Id, caption, opts)
}

// EditReplyMarkup replaces or removes this message's reply markup.
func (m *Message) EditReplyMarkup(c *Client, replyMarkup ReplyMarkup) (*Message, error) {
	return c.EditMessageReplyMarkup(m.ChatId, m.Id, &EditMessageReplyMarkupOpts{ReplyMarkup: replyMarkup})
}

// ReplyMediaGroup replies with a validated native TDLib album.
func (m *Message) ReplyMediaGroup(c *Client, media []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	copyOpts := SendMessageAlbumOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	if copyOpts.ReplyTo == nil {
		copyOpts.ReplyTo = &InputMessageReplyToMessage{MessageId: m.Id}
	}
	return c.SendMediaGroup(m.ChatId, media, &copyOpts)
}

// ReplyChecklist replies to the message with a checklist.
func (m *Message) ReplyChecklist(c *Client, checklist *InputChecklist, opts *SendChecklistOpts) (*Message, error) {
	if opts == nil {
		opts = &SendChecklistOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendChecklist(m.ChatId, checklist, opts)
}

// ReplyContact replies to the message with a contact.
func (m *Message) ReplyContact(c *Client, contact *Contact, opts *SendContactOpts) (*Message, error) {
	if opts == nil {
		opts = &SendContactOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendContact(m.ChatId, contact, opts)
}

// ReplyDice replies to the message with a dice.
func (m *Message) ReplyDice(c *Client, emoji string, opts *SendDiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDiceOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendDice(m.ChatId, emoji, opts)
}

// ReplyGame replies to the message with a game.
func (m *Message) ReplyGame(c *Client, botUserId int64, gameShortName string, opts *SendGameOpts) (*Message, error) {
	if opts == nil {
		opts = &SendGameOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendGame(m.ChatId, botUserId, gameShortName, opts)
}

// ReplyInvoice replies to the message with an invoice.
func (m *Message) ReplyInvoice(c *Client, invoice *Invoice, title string, description string, payload []byte, opts *SendInvoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendInvoiceOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendInvoice(m.ChatId, invoice, title, description, payload, opts)
}

// ReplyLocation replies to the message with a location.
func (m *Message) ReplyLocation(c *Client, location *Location, opts *SendLocationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendLocationOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendLocation(m.ChatId, location, opts)
}

// ReplyPaidMedia replies to the message with paid media.
func (m *Message) ReplyPaidMedia(c *Client, starCount int64, paidMedia []InputPaidMedia, opts *SendPaidMediaOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPaidMediaOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPaidMedia(m.ChatId, starCount, paidMedia, opts)
}

// ReplyPoll replies to the message with a poll.
func (m *Message) ReplyPoll(c *Client, question string, options []InputPollOption, opts *SendPollOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPollOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendPoll(m.ChatId, question, options, opts)
}

// ReplyStakeDice replies to the message with a stake dice.
func (m *Message) ReplyStakeDice(c *Client, stakeGramAmount int64, stateHash string, opts *SendStakeDiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStakeDiceOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendStakeDice(m.ChatId, stakeGramAmount, stateHash, opts)
}

// ReplyStory replies to the message with a story.
func (m *Message) ReplyStory(c *Client, storyPosterChatId int64, storyId int32, opts *SendStoryOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStoryOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendStory(m.ChatId, storyPosterChatId, storyId, opts)
}

// ReplyVenue replies to the message with a venue.
func (m *Message) ReplyVenue(c *Client, venue *Venue, opts *SendVenueOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVenueOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendVenue(m.ChatId, venue, opts)
}

// ReplyCopy replies to the message by copying another message.
func (m *Message) ReplyCopy(c *Client, fromChatId int64, messageId int64, opts *CopyMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &CopyMessageOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.CopyMessage(m.ChatId, fromChatId, messageId, opts)
}

// EditContent edits the message with the given content.
func (m *Message) EditContent(c *Client, content InputMessageContent, replyMarkup ReplyMarkup) (*Message, error) {
	return c.EditContent(m.ChatId, m.Id, content, replyMarkup)
}

// ReplyContent replies to the message with the given content.
func (m *Message) ReplyContent(c *Client, content InputMessageContent, opts *SendMessageContentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageContentOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyTo == nil {
		opts.ReplyTo = &InputMessageReplyToMessage{
			MessageId: m.Id,
		}
	}
	return c.SendMessageContent(m.ChatId, content, opts)
}

// ReplyRichMessage replies to the message with a rich message.
func (m *Message) ReplyRichMessage(c *Client, richMessage *InputRichMessage, opts *SendMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendRichMessage(m.ChatId, richMessage, opts)
}

// ReplyForwarded replies to the message by forwarding another message.
func (m *Message) ReplyForwarded(c *Client, fromChatId int64, messageId int64, inGameShare bool, opts *SendForwardedOpts) (*Message, error) {
	if opts == nil {
		opts = &SendForwardedOpts{}
	} else {
		copyOpts := *opts
		opts = &copyOpts
	}
	if opts.ReplyToMessageID == 0 {
		opts.ReplyToMessageID = m.Id
	}
	return c.SendForwarded(m.ChatId, fromChatId, messageId, inGameShare, opts)
}

// SendMessage sends text to this chat using the familiar convenience signature.
func (chat *Chat) SendMessage(c *Client, text string, opts *SendMessageOpts) (*Message, error) {
	return c.SendMessage(chat.Id, text, opts)
}

// SendPhoto sends a photo to this chat.
func (chat *Chat) SendPhoto(c *Client, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	return c.SendPhoto(chat.Id, photo, opts)
}

// SendDocument sends a document to this chat.
func (chat *Chat) SendDocument(c *Client, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	return c.SendDocument(chat.Id, document, opts)
}

// SendVideo sends a video to this chat.
func (chat *Chat) SendVideo(c *Client, video InputFile, opts *SendVideoOpts) (*Message, error) {
	return c.SendVideo(chat.Id, video, opts)
}

// SendMediaGroup sends a validated media group to this chat.
func (chat *Chat) SendMediaGroup(c *Client, media []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return c.SendMediaGroup(chat.Id, media, opts)
}
