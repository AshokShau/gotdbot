package gotdbot

import (
	"fmt"
	"html"
	"os"
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
	EffectId              int64
}

// SendTextMessage sends a text message to chat
func (c *Client) SendTextMessage(chatId int64, text string, opts *SendTextMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendTextMessageOpts{}
	}

	formattedText, err := GetFormattedText(c, text, opts.Entities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

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
	EffectId            int64
}

// SendPhoto sends a photo to chat
func (c *Client) SendPhoto(chatId int64, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPhotoOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessagePhoto{
		Photo:               photo,
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
	EffectId            int64
}

// SendVideo sends a video to chat
func (c *Client) SendVideo(chatId int64, video InputFile, opts *SendVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageVideo{
		Video:               video,
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
	EffectId            int64
}

// SendAnimation sends an animation to chat
func (c *Client) SendAnimation(chatId int64, animation InputFile, opts *SendAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAnimationOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageAnimation{
		Animation:           animation,
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
	EffectId            int64
}

// SendAudio sends an audio to chat
func (c *Client) SendAudio(chatId int64, audio InputFile, opts *SendAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAudioOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageAudio{
		Audio:               audio,
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
	EffectId                    int64
}

// SendDocument sends a document to chat
func (c *Client) SendDocument(chatId int64, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDocumentOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageDocument{
		Document:                    document,
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
	EffectId            int64
}

// SendVoice sends a voice note to chat
func (c *Client) SendVoice(chatId int64, voice InputFile, opts *SendVoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVoiceOpts{}
	}

	caption, err := GetFormattedText(c, opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageVoiceNote{
		VoiceNote: voice,
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
	EffectId            int64
}

// SendVideoNote sends a video note to chat
func (c *Client) SendVideoNote(chatId int64, videoNote InputFile, opts *SendVideoNoteOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoNoteOpts{}
	}

	content := &InputMessageVideoNote{
		VideoNote: videoNote,
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
	EffectId            int64
}

// SendSticker sends a sticker to chat
func (c *Client) SendSticker(chatId int64, sticker InputFile, opts *SendStickerOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStickerOpts{}
	}

	content := &InputMessageSticker{
		Sticker:   sticker,
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
	ReplyMarkup         ReplyMarkup
	ReplyToMessageID    int64
	EffectId            int64
}

// SendCopy copies a message to chat
func (c *Client) SendCopy(chatId int64, fromChatId int64, messageId int64, opts *SendCopyOpts) (*Message, error) {
	if opts == nil {
		opts = &SendCopyOpts{}
	}

	caption, err := GetFormattedText(c, opts.NewCaption, opts.NewCaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
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
	}, opts.TopicId, opts.Quote, opts.ReplyTo, opts.ReplyToMessageID, opts.ReplyMarkup)
}

// ForwardMessageOpts contains optional parameters for ForwardMessage
type ForwardMessageOpts struct {
	InGameShare         bool
	DisableNotification bool
	EffectId            int64
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

	formattedText, err := GetFormattedText(c, text, opts.Entities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

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

	return c.EditMessageText(chatId, content, messageId, &EditMessageTextOpts{
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

	return c.ParseTextEntities(mode, text)
}

func GetFormattedText(c *Client, text string, entities []TextEntity, parseMode string) (*FormattedText, error) {
	if len(entities) > 0 {
		return &FormattedText{
			Text:     text,
			Entities: entities,
		}, nil
	} else if parseMode != "" {
		ft, err := c.ParseText(text, parseMode)
		if err == nil {
			return ft, nil
		}
		return nil, err
	}
	return &FormattedText{Text: text}, nil
}

func GetInputFile(path string) InputFile {
	if _, err := os.Stat(path); err == nil {
		return InputFileLocal{Path: path}
	}

	return InputFileRemote{Id: path}
}

// EscapeHTML escapes HTML characters in the given text.
func EscapeHTML(text string) string {
	return html.EscapeString(text)
}

// EscapeMarkdown escapes Markdown characters in the given text.
func EscapeMarkdown(text string, version int) string {
	var chars string
	if version == 1 {
		chars = "_*`[\\"
	} else {
		chars = "_*[]()~`>#+-=|{}.!\\"
	}
	var b strings.Builder
	for _, c := range text {
		if strings.ContainsRune(chars, c) {
			b.WriteRune('\\')
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Mention returns a text mention for the given user ID.
func Mention(text string, userId int64, isHtml bool, escape bool) string {
	if escape {
		if isHtml {
			text = EscapeHTML(text)
		} else {
			text = EscapeMarkdown(text, 2)
		}
	}
	if isHtml {
		return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", userId, text)
	}
	return fmt.Sprintf("[%s](tg://user?id=%d)", text, userId)
}
