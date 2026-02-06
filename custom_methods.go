package gotdbot

import (
	"strings"
)

// SendTextMessageOpts contains optional parameters for SendTextMessage
type SendTextMessageOpts struct {
	ParseMode             string
	Entities              []TextEntity
	DisableWebPagePreview bool
	Url                   string
	ForceSmallMedia       bool
	ForceLargeMedia       bool
	ShowAboveText         bool
	DisableNotification   bool
	ProtectContent        bool
	AllowPaidBroadcast    bool
	TopicId               MessageTopic
	Quote                 *InputTextQuote
	ReplyTo               InputMessageReplyTo
	ReplyToMessageID      int64
	ReplyMarkup           ReplyMarkup
	ClearDraft            bool
	EffectId              string
}

// SendTextMessage sends a text message to chat
func (c *Client) SendTextMessage(chatId int64, text string, opts *SendTextMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendTextMessageOpts{}
	}

	formattedText := GetFormattedText(c, text, opts.Entities, opts.ParseMode)

	linkPreviewOptions := &LinkPreviewOptions{
		IsDisabled:      opts.DisableWebPagePreview,
		Url:             opts.Url,
		ForceSmallMedia: opts.ForceSmallMedia,
		ForceLargeMedia: opts.ForceLargeMedia,
		ShowAboveText:   opts.ShowAboveText,
	}

	content := &InputMessageText{
		Text:               formattedText,
		LinkPreviewOptions: linkPreviewOptions,
		ClearDraft:         opts.ClearDraft,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendPhotoOpts contains optional parameters for SendPhoto
type SendPhotoOpts struct {
	Caption             string
	CaptionEntities     []TextEntity
	ParseMode           string
	AddedStickerFileIds []int32
	Width               int32
	Height              int32
	SelfDestructType    MessageSelfDestructType
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	HasSpoiler          bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	Thumbnail           *InputThumbnail
	EffectId            string
}

// SendPhoto sends a photo to chat
func (c *Client) SendPhoto(chatId int64, photo string, opts *SendPhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPhotoOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(photo)

	content := &InputMessagePhoto{
		Photo:               inputFile,
		Thumbnail:           opts.Thumbnail,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Width:               opts.Width,
		Height:              opts.Height,
		Caption:             caption,
		SelfDestructType:    opts.SelfDestructType,
		HasSpoiler:          opts.HasSpoiler,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendVideoOpts contains optional parameters for SendVideo
type SendVideoOpts struct {
	Caption             string
	CaptionEntities     []TextEntity
	ParseMode           string
	AddedStickerFileIds []int32
	SupportsStreaming   bool
	Duration            int32
	Width               int32
	Height              int32
	SelfDestructType    MessageSelfDestructType
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	HasSpoiler          bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	Thumbnail           *InputThumbnail
	EffectId            string
}

// SendVideo sends a video to chat
func (c *Client) SendVideo(chatId int64, video string, opts *SendVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(video)

	content := &InputMessageVideo{
		Video:               inputFile,
		Thumbnail:           opts.Thumbnail,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Duration:            opts.Duration,
		Width:               opts.Width,
		Height:              opts.Height,
		SupportsStreaming:   opts.SupportsStreaming,
		Caption:             caption,
		SelfDestructType:    opts.SelfDestructType,
		HasSpoiler:          opts.HasSpoiler,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendAnimationOpts contains optional parameters for SendAnimation
type SendAnimationOpts struct {
	Caption             string
	CaptionEntities     []TextEntity
	ParseMode           string
	AddedStickerFileIds []int32
	Duration            int32
	Width               int32
	Height              int32
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	HasSpoiler          bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	Thumbnail           *InputThumbnail
	EffectId            string
}

// SendAnimation sends an animation to chat
func (c *Client) SendAnimation(chatId int64, animation string, opts *SendAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAnimationOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(animation)

	content := &InputMessageAnimation{
		Animation:           inputFile,
		Thumbnail:           opts.Thumbnail,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Duration:            opts.Duration,
		Width:               opts.Width,
		Height:              opts.Height,
		Caption:             caption,
		HasSpoiler:          opts.HasSpoiler,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendAudioOpts contains optional parameters for SendAudio
type SendAudioOpts struct {
	Caption             string
	CaptionEntities     []TextEntity
	ParseMode           string
	Title               string
	Performer           string
	Duration            int32
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	AlbumCoverThumbnail *InputThumbnail
	EffectId            string
}

// SendAudio sends an audio to chat
func (c *Client) SendAudio(chatId int64, audio string, opts *SendAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAudioOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(audio)

	content := &InputMessageAudio{
		Audio:               inputFile,
		AlbumCoverThumbnail: opts.AlbumCoverThumbnail,
		Title:               opts.Title,
		Performer:           opts.Performer,
		Duration:            opts.Duration,
		Caption:             caption,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendDocumentOpts contains optional parameters for SendDocument
type SendDocumentOpts struct {
	Caption                     string
	CaptionEntities             []TextEntity
	ParseMode                   string
	DisableContentTypeDetection bool
	DisableNotification         bool
	ProtectContent              bool
	AllowPaidBroadcast          bool
	TopicId                     MessageTopic
	Quote                       *InputTextQuote
	ReplyTo                     InputMessageReplyTo
	ReplyToMessageID            int64
	ReplyMarkup                 ReplyMarkup
	Thumbnail                   *InputThumbnail
	EffectId                    string
}

// SendDocument sends a document to chat
func (c *Client) SendDocument(chatId int64, document string, opts *SendDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDocumentOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(document)

	content := &InputMessageDocument{
		Document:                    inputFile,
		Thumbnail:                   opts.Thumbnail,
		DisableContentTypeDetection: opts.DisableContentTypeDetection,
		Caption:                     caption,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendVoiceOpts contains optional parameters for SendVoice
type SendVoiceOpts struct {
	Caption             string
	CaptionEntities     []TextEntity
	ParseMode           string
	Duration            int32
	Waveform            []byte
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	EffectId            string
}

// SendVoice sends a voice note to chat
func (c *Client) SendVoice(chatId int64, voice string, opts *SendVoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVoiceOpts{}
	}

	caption := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	inputFile := GetInputFile(voice)

	content := &InputMessageVoiceNote{
		VoiceNote: inputFile,
		Waveform:  opts.Waveform,
		Duration:  opts.Duration,
		Caption:   caption,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendVideoNoteOpts contains optional parameters for SendVideoNote
type SendVideoNoteOpts struct {
	Duration            int32
	Length              int32
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	Thumbnail           *InputThumbnail
	EffectId            string
}

// SendVideoNote sends a video note to chat
func (c *Client) SendVideoNote(chatId int64, videoNote string, opts *SendVideoNoteOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoNoteOpts{}
	}

	inputFile := GetInputFile(videoNote)

	content := &InputMessageVideoNote{
		VideoNote: inputFile,
		Thumbnail: opts.Thumbnail,
		Duration:  opts.Duration,
		Length:    opts.Length,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendStickerOpts contains optional parameters for SendSticker
type SendStickerOpts struct {
	Emoji               string
	Width               int32
	Height              int32
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	ReplyMarkup         ReplyMarkup
	Thumbnail           *InputThumbnail
	EffectId            string
}

// SendSticker sends a sticker to chat
func (c *Client) SendSticker(chatId int64, sticker string, opts *SendStickerOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStickerOpts{}
	}

	inputFile := GetInputFile(sticker)

	content := &InputMessageSticker{
		Sticker:   inputFile,
		Thumbnail: opts.Thumbnail,
		Width:     opts.Width,
		Height:    opts.Height,
		Emoji:     opts.Emoji,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// SendCopyOpts contains optional parameters for SendCopy
type SendCopyOpts struct {
	InGameShare         bool
	ReplaceCaption      bool
	NewCaption          string
	NewCaptionEntities  []TextEntity
	ParseMode           string
	DisableNotification bool
	ProtectContent      bool
	AllowPaidBroadcast  bool
	TopicId             MessageTopic
	Quote               *InputTextQuote
	ReplyTo             InputMessageReplyTo
	ReplyToMessageID    int64
	EffectId            string
}

// SendCopy copies a message to chat
func (c *Client) SendCopy(chatId int64, fromChatId int64, messageId int64, opts *SendCopyOpts) (*Message, error) {
	if opts == nil {
		opts = &SendCopyOpts{}
	}

	caption := GetFormattedText(c, opts.NewCaption, opts.NewCaptionEntities, opts.ParseMode)

	content := &InputMessageForwarded{
		FromChatId:  fromChatId,
		MessageId:   messageId,
		InGameShare: opts.InGameShare,
		CopyOptions: &MessageCopyOptions{
			SendCopy:       true,
			ReplaceCaption: opts.ReplaceCaption,
			NewCaption:     caption,
		},
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		ProtectContent:      opts.ProtectContent,
		AllowPaidBroadcast:  opts.AllowPaidBroadcast,
		EffectId:            opts.EffectId,
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, nil)
}

// ForwardMessageOpts contains optional parameters for ForwardMessage
type ForwardMessageOpts struct {
	InGameShare         bool
	DisableNotification bool
	EffectId            string
}

// ForwardMessage forwards a message to chat
func (c *Client) ForwardMessage(chatId int64, fromChatId int64, messageId int64, opts *ForwardMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &ForwardMessageOpts{}
	}

	content := &InputMessageForwarded{
		FromChatId:  fromChatId,
		MessageId:   messageId,
		InGameShare: opts.InGameShare,
	}

	return c.sendMessageWithContent(chatId, content, &MessageSendOptions{
		DisableNotification: opts.DisableNotification,
		EffectId:            opts.EffectId,
	}, nil, nil, nil, 0, nil)
}

// EditTextMessageOpts contains optional parameters for EditTextMessage
type EditTextMessageOpts struct {
	ParseMode             string
	Entities              []TextEntity
	DisableWebPagePreview bool
	Url                   string
	ForceSmallMedia       bool
	ForceLargeMedia       bool
	ShowAboveText         bool
	ReplyMarkup           ReplyMarkup
}

// EditTextMessage edits a text message
func (c *Client) EditTextMessage(chatId int64, messageId int64, text string, opts *EditTextMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &EditTextMessageOpts{}
	}

	if !*c.config.UseMessageDatabase {
		if _, err := c.GetMessage(chatId, messageId); err != nil {
			return nil, err
		}
	}

	formattedText := GetFormattedText(c, text, opts.Entities, opts.ParseMode)

	linkPreviewOptions := &LinkPreviewOptions{
		IsDisabled:      opts.DisableWebPagePreview,
		Url:             opts.Url,
		ForceSmallMedia: opts.ForceSmallMedia,
		ForceLargeMedia: opts.ForceLargeMedia,
		ShowAboveText:   opts.ShowAboveText,
	}

	content := &InputMessageText{
		Text:               formattedText,
		LinkPreviewOptions: linkPreviewOptions,
	}

	return c.EditMessageText(chatId, messageId, content, &EditMessageTextOpts{
		ReplyMarkup: opts.ReplyMarkup,
	})
}

// GetSupergroupId returns the supergroup ID from a chat ID
func (c *Client) GetSupergroupId(chatId int64) (int64, error) {
	chat, err := c.GetChat(chatId)
	if err != nil {
		return 0, err
	}

	if chat.TypeField == nil {
		return 0, nil
	}

	if ct, ok := chat.TypeField.(*ChatTypeSupergroup); ok {
		return ct.SupergroupId, nil
	}

	return 0, nil
}

// sendMessageWithContent is a helper function to send a message with content
func (c *Client) sendMessageWithContent(
	chatId int64,
	content InputMessageContent,
	options *MessageSendOptions,
	topicId MessageTopic,
	quote *InputTextQuote,
	replyTo InputMessageReplyTo,
	replyToMessageId int64,
	replyMarkup ReplyMarkup,
) (*Message, error) {
	if replyToMessageId > 0 {
		replyTo = &InputMessageReplyToMessage{
			MessageId: replyToMessageId,
			Quote:     quote,
		}
	}

	if c.config.LoadMessagesBeforeReply && replyToMessageId > 0 {
		_, _ = c.GetMessage(chatId, replyToMessageId)
	}

	return c.SendMessage(chatId, content, &SendMessageOpts{
		TopicId:     topicId,
		ReplyTo:     replyTo,
		Options:     options,
		ReplyMarkup: replyMarkup,
	})
}

// ParseText parses the text using the specified parse mode.
func (c *Client) ParseText(text string, parseMode string) (*FormattedText, error) {
	var mode TextParseMode

	switch strings.ToLower(parseMode) {
	case "markdown":
		mode = &TextParseModeMarkdown{Version: 1}
	case "markdownv2":
		mode = &TextParseModeMarkdown{Version: 2}
	case "html":
		mode = &TextParseModeHTML{}
	default:
		return &FormattedText{Text: text}, nil
	}

	return c.ParseTextEntities(text, mode)
}
