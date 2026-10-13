package gotdbot

import (
	"fmt"
)

// SendMessageOpts contains optional parameters for SendMessage.
type SendMessageOpts struct {
	ParseMode                         string
	Entities                          []TextEntity
	DisableWebPagePreview             bool
	Url                               string
	ForceSmallMedia                   bool
	ForceLargeMedia                   bool
	ShowAboveText                     bool
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	ClearDraft                        bool
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendMessage sends a text message to chat.
func (c *Client) SendMessage(chatId int64, text string, opts *SendMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}

	formattedText, err := c.GetFormattedText(text, opts.Entities, opts.ParseMode)
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

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendPhotoOpts contains optional parameters for SendPhoto.
type SendPhotoOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	AddedStickerFileIds               []int32
	Width                             int32
	Height                            int32
	SelfDestructType                  MessageSelfDestructType
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	HasSpoiler                        bool
	ShowCaptionAboveMedia             bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
	// Video of a live photo; pass nil to send a regular photo.
	Video InputFile
}

// SendPhoto sends a photo to chat.
func (c *Client) SendPhoto(chatId int64, photo InputFile, opts *SendPhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPhotoOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessagePhoto{
		Photo: &InputPhoto{
			Photo:               photo,
			Video:               opts.Video,
			Thumbnail:           opts.Thumbnail,
			AddedStickerFileIds: opts.AddedStickerFileIds,
			Width:               opts.Width,
			Height:              opts.Height,
		},
		Caption:               caption,
		SelfDestructType:      opts.SelfDestructType,
		HasSpoiler:            opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendVideoOpts contains optional parameters for SendVideo.
type SendVideoOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	AddedStickerFileIds               []int32
	SupportsStreaming                 bool
	Duration                          int32
	Width                             int32
	Height                            int32
	SelfDestructType                  MessageSelfDestructType
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	HasSpoiler                        bool
	ShowCaptionAboveMedia             bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	StartTimestamp                    int32
	Cover                             InputFile
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendVideo sends a video to chat.
func (c *Client) SendVideo(chatId int64, video InputFile, opts *SendVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageVideo{
		Video: &InputVideo{
			Video:               video,
			Thumbnail:           opts.Thumbnail,
			AddedStickerFileIds: opts.AddedStickerFileIds,
			Duration:            opts.Duration,
			Width:               opts.Width,
			Height:              opts.Height,
			SupportsStreaming:   opts.SupportsStreaming,
			StartTimestamp:      opts.StartTimestamp,
			Cover:               opts.Cover,
		},
		Caption:               caption,
		SelfDestructType:      opts.SelfDestructType,
		HasSpoiler:            opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendAnimationOpts contains optional parameters for SendAnimation.
type SendAnimationOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	AddedStickerFileIds               []int32
	Duration                          int32
	Width                             int32
	Height                            int32
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	HasSpoiler                        bool
	ShowCaptionAboveMedia             bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendAnimation sends an animation to chat.
func (c *Client) SendAnimation(chatId int64, animation InputFile, opts *SendAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAnimationOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageAnimation{
		Animation: &InputAnimation{
			Animation:           animation,
			Thumbnail:           opts.Thumbnail,
			AddedStickerFileIds: opts.AddedStickerFileIds,
			Duration:            opts.Duration,
			Width:               opts.Width,
			Height:              opts.Height,
		},
		Caption:               caption,
		HasSpoiler:            opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendAudioOpts contains optional parameters for SendAudio.
type SendAudioOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	Title                             string
	Performer                         string
	Duration                          int32
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	AlbumCoverThumbnail               *InputThumbnail
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendAudio sends an audio to chat.
func (c *Client) SendAudio(chatId int64, audio InputFile, opts *SendAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &SendAudioOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageAudio{
		Audio: &InputAudio{
			Audio:               audio,
			AlbumCoverThumbnail: opts.AlbumCoverThumbnail,
			Title:               opts.Title,
			Performer:           opts.Performer,
			Duration:            opts.Duration,
		},
		Caption: caption,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendDocumentOpts contains optional parameters for SendDocument.
type SendDocumentOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	DisableContentTypeDetection       bool
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendDocument sends a document to chat.
func (c *Client) SendDocument(chatId int64, document InputFile, opts *SendDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDocumentOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageDocument{
		Document: &InputDocument{
			Document:                    document,
			Thumbnail:                   opts.Thumbnail,
			DisableContentTypeDetection: opts.DisableContentTypeDetection,
		},
		Caption: caption,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendVoiceOpts contains optional parameters for SendVoice.
type SendVoiceOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	Duration                          int32
	Waveform                          []byte
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	SelfDestructType                  MessageSelfDestructType
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendVoice sends a voice note to chat.
func (c *Client) SendVoice(chatId int64, voice InputFile, opts *SendVoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVoiceOpts{}
	}

	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	content := &InputMessageVoiceNote{
		VoiceNote: &InputVoiceNote{
			VoiceNote: voice,
			Waveform:  opts.Waveform,
			Duration:  opts.Duration,
		},
		Caption:          caption,
		SelfDestructType: opts.SelfDestructType,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendVideoNoteOpts contains optional parameters for SendVideoNote.
type SendVideoNoteOpts struct {
	Duration                          int32
	Length                            int32
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	SelfDestructType                  MessageSelfDestructType
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendVideoNote sends a video note to chat.
func (c *Client) SendVideoNote(chatId int64, videoNote InputFile, opts *SendVideoNoteOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVideoNoteOpts{}
	}

	content := &InputMessageVideoNote{
		VideoNote: &InputVideoNote{
			VideoNote: videoNote,
			Thumbnail: opts.Thumbnail,
			Duration:  opts.Duration,
			Length:    opts.Length,
		},
		SelfDestructType: opts.SelfDestructType,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendStickerOpts contains optional parameters for SendSticker.
type SendStickerOpts struct {
	Emoji                             string
	Width                             int32
	Height                            int32
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	Thumbnail                         *InputThumbnail
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendSticker sends a sticker to chat.
func (c *Client) SendSticker(chatId int64, sticker InputFile, opts *SendStickerOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStickerOpts{}
	}

	content := &InputMessageSticker{
		Sticker: &InputSticker{
			Sticker:   sticker,
			Thumbnail: opts.Thumbnail,
			Width:     opts.Width,
			Height:    opts.Height,
		},
		Emoji: opts.Emoji,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendMediaGroup sends a native TDLib album of 2-10 photo/video, audio, or
// document contents. Audio and document albums must each contain only that type.
func (c *Client) SendMediaGroup(chatId int64, media []InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	if len(media) < 2 || len(media) > 10 {
		return nil, fmt.Errorf("media group requires 2-10 items")
	}
	kind := ""
	var above *bool
	for i, item := range media {
		category := ""
		var captionAbove bool
		switch v := item.(type) {
		case *InputMessagePhoto:
			if v == nil || v.Photo == nil || v.Photo.Photo == nil {
				return nil, fmt.Errorf("media[%d]: photo required", i)
			}
			category = "visual"
			captionAbove = v.ShowCaptionAboveMedia
		case *InputMessageVideo:
			if v == nil || v.Video == nil || v.Video.Video == nil {
				return nil, fmt.Errorf("media[%d]: video required", i)
			}
			category = "visual"
			captionAbove = v.ShowCaptionAboveMedia
		case *InputMessageAudio:
			if v == nil || v.Audio == nil || v.Audio.Audio == nil {
				return nil, fmt.Errorf("media[%d]: audio required", i)
			}
			category = "audio"
		case *InputMessageDocument:
			if v == nil || v.Document == nil || v.Document.Document == nil {
				return nil, fmt.Errorf("media[%d]: document required", i)
			}
			category = "document"
		default:
			return nil, fmt.Errorf("media[%d]: unsupported album content %T", i, item)
		}
		if kind != "" && kind != category {
			return nil, fmt.Errorf("audio and document albums require items of the same type")
		}
		if above != nil && *above != captionAbove {
			return nil, fmt.Errorf("album items must agree on ShowCaptionAboveMedia")
		}
		kind = category
		value := captionAbove
		above = &value
	}
	return c.SendMessageAlbum(chatId, media, opts)
}

// SendContactOpts contains optional parameters for SendContact.
type SendContactOpts struct {
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendContact sends a contact to chat.
func (c *Client) SendContact(chatId int64, contact *Contact, opts *SendContactOpts) (*Message, error) {
	if opts == nil {
		opts = &SendContactOpts{}
	}
	content := &InputMessageContact{
		Contact: contact,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendLocationOpts contains optional parameters for SendLocation.
type SendLocationOpts struct {
	Heading                           int32
	LivePeriod                        int32
	ProximityAlertRadius              int32
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendLocation sends a location to chat.
func (c *Client) SendLocation(chatId int64, location *Location, opts *SendLocationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendLocationOpts{}
	}

	var content InputMessageContent
	if opts.LivePeriod > 0 {
		content = &InputMessageLiveLocation{
			Location: &LiveLocation{
				Location:             location,
				LivePeriod:           opts.LivePeriod,
				Heading:              opts.Heading,
				ProximityAlertRadius: opts.ProximityAlertRadius,
			},
		}
	} else {
		content = &InputMessageLocation{
			Location: location,
		}
	}

	return c.sendLocationContent(chatId, content, opts)
}

func (c *Client) sendLocationContent(chatId int64, content InputMessageContent, opts *SendLocationOpts) (*Message, error) {
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendVenueOpts contains optional parameters for SendVenue.
type SendVenueOpts struct {
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool
	CallbackQueryID                   int64
	ReceiverUserID                    int64
	ReplaceCallbackQueryMessage       bool
}

// SendVenue sends a venue to chat.
func (c *Client) SendVenue(chatId int64, venue *Venue, opts *SendVenueOpts) (*Message, error) {
	if opts == nil {
		opts = &SendVenueOpts{}
	}
	content := &InputMessageVenue{
		Venue: venue,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendChecklistOpts contains optional parameters for SendChecklist.
type SendChecklistOpts struct {
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendChecklist sends a checklist to chat.
func (c *Client) SendChecklist(chatId int64, checklist *InputChecklist, opts *SendChecklistOpts) (*Message, error) {
	if opts == nil {
		opts = &SendChecklistOpts{}
	}
	content := &InputMessageChecklist{
		Checklist: checklist,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendDiceOpts contains optional parameters for SendDice.
type SendDiceOpts struct {
	ClearDraft                        bool
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendDice sends a dice to chat.
func (c *Client) SendDice(chatId int64, emoji string, opts *SendDiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendDiceOpts{}
	}
	content := &InputMessageDice{
		Emoji:      emoji,
		ClearDraft: opts.ClearDraft,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendGameOpts contains optional parameters for SendGame.
type SendGameOpts struct {
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendGame sends a game to chat.
func (c *Client) SendGame(chatId int64, botUserId int64, gameShortName string, opts *SendGameOpts) (*Message, error) {
	if opts == nil {
		opts = &SendGameOpts{}
	}
	content := &InputMessageGame{
		BotUserId:     botUserId,
		GameShortName: gameShortName,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendPollOpts contains optional parameters for SendPoll.
type SendPollOpts struct {
	// AllowAddingOptions permits adding answers to a regular poll after creation.
	AllowAddingOptions                bool
	AllowsMultipleAnswers             bool
	AllowsRevoting                    bool
	CloseDate                         int32
	Description                       string
	DescriptionEntities               []TextEntity
	ParseMode                         string
	HideResultsUntilCloses            bool
	IsAnonymous                       bool
	IsClosed                          bool
	OpenPeriod                        int32
	QuestionEntities                  []TextEntity
	ShuffleOptions                    bool
	Type                              InputPollType
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	Media                             InputPollMedia
	MembersOnly                       bool
	CountryCodes                      []string
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendPoll sends a poll to chat.
func (c *Client) SendPoll(chatId int64, question string, options []InputPollOption, opts *SendPollOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPollOpts{}
	}
	formattedQuestion, err := c.GetFormattedText(question, opts.QuestionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	description, err := c.GetFormattedText(opts.Description, opts.DescriptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	pollType := opts.Type
	if pollType == nil {
		pollType = &InputPollTypeRegular{AllowAddingOptions: opts.AllowAddingOptions}
	}
	content := &InputMessagePoll{
		AllowsMultipleAnswers:  opts.AllowsMultipleAnswers,
		AllowsRevoting:         opts.AllowsRevoting,
		CloseDate:              opts.CloseDate,
		Description:            description,
		HideResultsUntilCloses: opts.HideResultsUntilCloses,
		IsAnonymous:            opts.IsAnonymous,
		IsClosed:               opts.IsClosed,
		OpenPeriod:             opts.OpenPeriod,
		Options:                options,
		Question:               formattedQuestion,
		ShuffleOptions:         opts.ShuffleOptions,
		Type:                   pollType,
		Media:                  opts.Media,
		MembersOnly:            opts.MembersOnly,
		CountryCodes:           opts.CountryCodes,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendQuizOpts controls quiz formatting, explanations, and delivery.
type SendQuizOpts struct {
	// SendPollOpts configures the question, answers, and delivery. Type is ignored.
	SendPollOpts
	Explanation         string
	ExplanationEntities []TextEntity
	ExplanationMedia    InputPollMedia
}

// SendQuiz sends a quiz with increasing, zero-based correct answer IDs.
func (c *Client) SendQuiz(chatId int64, question string, options []InputPollOption, correctOptionIds []int32, opts *SendQuizOpts) (*Message, error) {
	if len(correctOptionIds) == 0 {
		return nil, fmt.Errorf("quiz requires at least one correct answer")
	}
	for index, id := range correctOptionIds {
		if id < 0 || int(id) >= len(options) || (index > 0 && id <= correctOptionIds[index-1]) {
			return nil, fmt.Errorf("quiz answer IDs must be increasing and within the answer list")
		}
	}
	if opts == nil {
		opts = &SendQuizOpts{}
	}
	explanation, err := c.GetFormattedText(opts.Explanation, opts.ExplanationEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	copyOpts := opts.SendPollOpts
	copyOpts.Type = &InputPollTypeQuiz{CorrectOptionIds: correctOptionIds, Explanation: explanation, ExplanationMedia: opts.ExplanationMedia}
	return c.SendPoll(chatId, question, options, &copyOpts)
}

// SendStakeDiceOpts contains optional parameters for SendStakeDice.
type SendStakeDiceOpts struct {
	ClearDraft                        bool
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendStakeDice sends a stake dice to chat.
func (c *Client) SendStakeDice(chatId int64, stakeGramAmount int64, stateHash string, opts *SendStakeDiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStakeDiceOpts{}
	}
	content := &InputMessageStakeDice{
		ClearDraft:      opts.ClearDraft,
		StakeGramAmount: stakeGramAmount,
		StateHash:       stateHash,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendStoryOpts contains optional parameters for SendStory.
type SendStoryOpts struct {
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendStory sends a story to chat.
func (c *Client) SendStory(chatId int64, storyPosterChatId int64, storyId int32, opts *SendStoryOpts) (*Message, error) {
	if opts == nil {
		opts = &SendStoryOpts{}
	}
	content := &InputMessageStory{
		StoryId:           storyId,
		StoryPosterChatId: storyPosterChatId,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendInvoiceOpts contains optional parameters for SendInvoice.
type SendInvoiceOpts struct {
	PaidMedia                         *InputPaidMedia
	PaidMediaCaption                  string
	PaidMediaEntities                 []TextEntity
	ParseMode                         string
	PhotoHeight                       int32
	PhotoSize                         int32
	PhotoUrl                          string
	PhotoWidth                        int32
	ProviderData                      string
	ProviderToken                     string
	StartParameter                    string
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendInvoice sends an invoice to chat.
func (c *Client) SendInvoice(chatId int64, invoice *Invoice, title string, description string, payload []byte, opts *SendInvoiceOpts) (*Message, error) {
	if opts == nil {
		opts = &SendInvoiceOpts{}
	}
	paidMediaCaption, err := c.GetFormattedText(opts.PaidMediaCaption, opts.PaidMediaEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	content := &InputMessageInvoice{
		Description:      description,
		Invoice:          invoice,
		PaidMedia:        opts.PaidMedia,
		PaidMediaCaption: paidMediaCaption,
		Payload:          payload,
		PhotoHeight:      opts.PhotoHeight,
		PhotoSize:        opts.PhotoSize,
		PhotoUrl:         opts.PhotoUrl,
		PhotoWidth:       opts.PhotoWidth,
		ProviderData:     opts.ProviderData,
		ProviderToken:    opts.ProviderToken,
		StartParameter:   opts.StartParameter,
		Title:            title,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendPaidMediaOpts contains optional parameters for SendPaidMedia.
type SendPaidMediaOpts struct {
	Caption                           string
	CaptionEntities                   []TextEntity
	ParseMode                         string
	Payload                           string
	ShowCaptionAboveMedia             bool
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup
	EffectId                          int64
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	SuggestedPostInfo                 *InputSuggestedPostInfo
	UpdateOrderOfInstalledStickerSets bool

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

// SendPaidMedia sends paid media to chat.
func (c *Client) SendPaidMedia(chatId int64, starCount int64, paidMedia []InputPaidMedia, opts *SendPaidMediaOpts) (*Message, error) {
	if opts == nil {
		opts = &SendPaidMediaOpts{}
	}
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	content := &InputMessagePaidMedia{
		Caption:               caption,
		PaidMedia:             paidMedia,
		Payload:               opts.Payload,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
		StarCount:             starCount,
	}
	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendForwardedOpts controls forwarding, reply destinations, and message delivery.
type SendForwardedOpts struct {
	InGameShare                       bool
	DisableNotification               bool
	EffectId                          int64
	ReplaceVideoStartTimestamp        bool
	NewVideoStartTimestamp            int32
	FromBackground                    bool
	OnlyPreview                       bool
	PaidMessageStarCount              int64
	ProtectContent                    bool
	SchedulingState                   MessageSchedulingState
	SendingId                         int32
	UpdateOrderOfInstalledStickerSets bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyToMessageID                  int64
	ReplyMarkup                       ReplyMarkup

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
	AllowPaidBroadcast          bool
	SuggestedPostInfo           *InputSuggestedPostInfo
}

// SendForwarded sends a forwarded message to chat.
func (c *Client) SendForwarded(chatId int64, fromChatId int64, messageId int64, opts *SendForwardedOpts) (*Message, error) {
	if opts == nil {
		opts = &SendForwardedOpts{}
	}

	content := &InputMessageForwarded{
		FromChatId:                 fromChatId,
		MessageId:                  messageId,
		InGameShare:                opts.InGameShare,
		ReplaceVideoStartTimestamp: opts.ReplaceVideoStartTimestamp,
		NewVideoStartTimestamp:     opts.NewVideoStartTimestamp,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			ProtectContent:                    opts.ProtectContent,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendMessageDraftOpts controls text draft formatting and generation controls.
type SendMessageDraftOpts struct {
	ParseMode    string
	Entities     []TextEntity
	ForumTopicId int32
	CanStop      bool
	KeepOnStop   bool
}

// SendMessageDraft streams a partial text message; for bots only. Reuse draftId
// to update a draft, then send the final message with SendMessage.
func (c *Client) SendMessageDraft(chatId int64, draftId int64, text string, opts *SendMessageDraftOpts) error {
	if opts == nil {
		opts = &SendMessageDraftOpts{}
	}
	formatted, err := c.GetFormattedText(text, opts.Entities, opts.ParseMode)
	if err != nil {
		return err
	}
	return c.SendTextMessageDraft(chatId, draftId, &SendTextMessageDraftOpts{
		Text: formatted, ForumTopicId: opts.ForumTopicId, CanStop: opts.CanStop, KeepOnStop: opts.KeepOnStop,
	})
}

// SendRichMessageDraftTextOpts controls a rich Markdown or HTML draft. Media
// must already be uploaded; draft requests cannot upload new files.
type SendRichMessageDraftTextOpts struct {
	SendRichMessageDraftOpts
	Media                 []InputRichMessageMedia
	IsRtl                 bool
	DetectAutomaticBlocks bool
}

// SendRichMessageDraftMarkdown streams a partial rich Markdown message; for bots only.
func (c *Client) SendRichMessageDraftMarkdown(chatId int64, draftId int64, text string, opts *SendRichMessageDraftTextOpts) error {
	if opts == nil {
		opts = &SendRichMessageDraftTextOpts{}
	}
	return c.SendRichMessageDraft(chatId, draftId, &InputRichMessage{
		Source: &RichMessageSourceMarkdown{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, &opts.SendRichMessageDraftOpts)
}

// SendRichMessageDraftHTML streams a partial rich HTML message; for bots only.
func (c *Client) SendRichMessageDraftHTML(chatId int64, draftId int64, text string, opts *SendRichMessageDraftTextOpts) error {
	if opts == nil {
		opts = &SendRichMessageDraftTextOpts{}
	}
	return c.SendRichMessageDraft(chatId, draftId, &InputRichMessage{
		Source: &RichMessageSourceHtml{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, &opts.SendRichMessageDraftOpts)
}

// SendRichMessage sends a rich message to chat.
func (c *Client) SendRichMessage(chatId int64, richMessage *InputRichMessage, opts *SendMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &SendMessageOpts{}
	}

	content := &InputMessageRichMessage{
		ClearDraft: opts.ClearDraft,
		Message:    richMessage,
	}

	return c.sendMessageWithContent(chatId, content, &sendMessageContentOptions{
		Options: &MessageSendOptions{
			DisableNotification:               opts.DisableNotification,
			ProtectContent:                    opts.ProtectContent,
			AllowPaidBroadcast:                opts.AllowPaidBroadcast,
			EffectId:                          opts.EffectId,
			FromBackground:                    opts.FromBackground,
			OnlyPreview:                       opts.OnlyPreview,
			PaidMessageStarCount:              opts.PaidMessageStarCount,
			SchedulingState:                   opts.SchedulingState,
			SendingId:                         opts.SendingId,
			SuggestedPostInfo:                 opts.SuggestedPostInfo,
			UpdateOrderOfInstalledStickerSets: opts.UpdateOrderOfInstalledStickerSets,
		},
		TopicId:                     opts.TopicId,
		Quote:                       opts.Quote,
		ReplyTo:                     opts.ReplyTo,
		ReplyToMessageID:            opts.ReplyToMessageID,
		ReplyMarkup:                 opts.ReplyMarkup,
		CallbackQueryID:             opts.CallbackQueryID,
		ReceiverUserID:              opts.ReceiverUserID,
		ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
	})
}

// SendRichTextOpts contains options for SendRichMessageMarkdown and
// SendRichMessageHTML. ParseMode, Entities, and link preview options in
// SendMessageOpts are unused: TDLib parses the rich source itself.
type SendRichTextOpts struct {
	SendMessageOpts
	// Media referenced by identifier in the rich source.
	Media []InputRichMessageMedia
	// Display the message from right to left.
	IsRtl bool
	// Detect URLs, email addresses, and other automatic blocks.
	DetectAutomaticBlocks bool
}

// SendRichMessageMarkdown sends a rich Markdown message; for bots only.
// Rich Markdown is separate from the Markdown/MarkdownV2 text parse modes.
func (c *Client) SendRichMessageMarkdown(chatId int64, text string, opts *SendRichTextOpts) (*Message, error) {
	if opts == nil {
		opts = &SendRichTextOpts{}
	}
	return c.SendRichMessage(chatId, &InputRichMessage{
		Source: &RichMessageSourceMarkdown{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, &opts.SendMessageOpts)
}

// SendRichMessageHTML sends a rich HTML message; for bots only.
// Rich HTML is separate from the HTML text parse mode.
func (c *Client) SendRichMessageHTML(chatId int64, text string, opts *SendRichTextOpts) (*Message, error) {
	if opts == nil {
		opts = &SendRichTextOpts{}
	}
	return c.SendRichMessage(chatId, &InputRichMessage{
		Source: &RichMessageSourceHtml{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, &opts.SendMessageOpts)
}

// SendLivePhoto sends a photo with an attached live-photo video. The video argument
// supplies the motion component and takes precedence over opts.Video.
func (c *Client) SendLivePhoto(chatId int64, photo, video InputFile, opts *SendPhotoOpts) (*Message, error) {
	copyOpts := SendPhotoOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	copyOpts.Video = video
	return c.SendPhoto(chatId, photo, &copyOpts)
}

// SendLiveLocation sends a native live location. Its period, heading, and
// proximity radius come from location; corresponding fields in opts are unused.
func (c *Client) SendLiveLocation(chatId int64, location *LiveLocation, opts *SendLocationOpts) (*Message, error) {
	if opts == nil {
		opts = &SendLocationOpts{}
	}
	return c.sendLocationContent(chatId, &InputMessageLiveLocation{Location: location}, opts)
}

type sendMessageContentOptions struct {
	Options                     *MessageSendOptions
	TopicId                     MessageTopic
	Quote                       *InputTextQuote
	ReplyTo                     InputMessageReplyTo
	ReplyToMessageID            int64
	ReplyMarkup                 ReplyMarkup
	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
}

func (c *Client) sendMessageWithContent(chatId int64, content InputMessageContent, opts *sendMessageContentOptions) (*Message, error) {
	if opts == nil {
		opts = &sendMessageContentOptions{}
	}
	replyTo := opts.ReplyTo
	if replyTo == nil && opts.ReplyToMessageID > 0 {
		replyTo = &InputMessageReplyToMessage{
			MessageId: opts.ReplyToMessageID,
			Quote:     opts.Quote,
		}
	}

	if c.config != nil && c.config.LoadMessagesBeforeReply {
		var replyID int64
		switch reply := replyTo.(type) {
		case *InputMessageReplyToMessage:
			if reply != nil {
				replyID = reply.MessageId
			}
		case InputMessageReplyToMessage:
			replyID = reply.MessageId
		}
		if replyID > 0 {
			if _, err := c.GetMessage(chatId, replyID); err != nil {
				return nil, fmt.Errorf("load reply message: %w", err)
			}
		}
	}

	if opts.ReceiverUserID != 0 || opts.CallbackQueryID > 0 {
		var sendingId int32
		var onlyPreview bool

		var protectContent bool
		if opts.Options != nil {
			sendingId = opts.Options.SendingId
			onlyPreview = opts.Options.OnlyPreview
			protectContent = opts.Options.ProtectContent
		}

		return c.SendEphemeralMessage(opts.CallbackQueryID, chatId, content, opts.ReceiverUserID, sendingId, &SendEphemeralMessageOpts{
			OnlyPreview:                 onlyPreview,
			ProtectContent:              protectContent,
			ReplaceCallbackQueryMessage: opts.ReplaceCallbackQueryMessage,
			ReplyMarkup:                 opts.ReplyMarkup,
			ReplyTo:                     replyTo,
			TopicId:                     opts.TopicId,
		})
	}

	return c.SendMessageContent(chatId, content, &SendMessageContentOpts{
		TopicId:     opts.TopicId,
		ReplyTo:     replyTo,
		Options:     opts.Options,
		ReplyMarkup: opts.ReplyMarkup,
	})
}
