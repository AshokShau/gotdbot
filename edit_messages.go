package gotdbot

import (
	"fmt"
	"reflect"
)

// EditRichMessageOpts contains options for editing a rich message.
type EditRichMessageOpts struct {
	ReplyMarkup ReplyMarkup
}

// EditRichTextOpts contains options for editing rich Markdown or HTML.
type EditRichTextOpts struct {
	Media                 []InputRichMessageMedia
	IsRtl                 bool
	DetectAutomaticBlocks bool
	ReplyMarkup           ReplyMarkup
}

// EditRichMessage replaces a message's text or rich content with a rich
// message. Source may contain blocks, Markdown, or HTML.
func (c *Client) EditRichMessage(chatId, messageId int64, message *InputRichMessage, opts *EditRichMessageOpts) (*Message, error) {
	if err := c.loadMessageBeforeEdit(chatId, messageId); err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditRichMessageOpts{}
	}
	return c.EditMessageTextContent(chatId, &InputMessageRichMessage{Message: message}, messageId, &EditMessageTextContentOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditRichMessageMarkdown replaces a message with rich Markdown; for bots only.
func (c *Client) EditRichMessageMarkdown(chatId, messageId int64, text string, opts *EditRichTextOpts) (*Message, error) {
	message, markup := richMarkdownForEdit(text, opts)
	return c.EditRichMessage(chatId, messageId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditRichMessageHTML replaces a message with rich HTML; for bots only.
func (c *Client) EditRichMessageHTML(chatId, messageId int64, text string, opts *EditRichTextOpts) (*Message, error) {
	message, markup := richHTMLForEdit(text, opts)
	return c.EditRichMessage(chatId, messageId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditInlineRichMessage replaces an inline message with rich content;
// for bots only. Inline edits cannot upload new files.
func (c *Client) EditInlineRichMessage(inlineMessageId string, message *InputRichMessage, opts *EditRichMessageOpts) error {
	if opts == nil {
		opts = &EditRichMessageOpts{}
	}
	return c.EditInlineMessageText(inlineMessageId, &InputMessageRichMessage{Message: message}, &EditInlineMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineRichMessageMarkdown replaces an inline message with rich Markdown.
func (c *Client) EditInlineRichMessageMarkdown(inlineMessageId, text string, opts *EditRichTextOpts) error {
	message, markup := richMarkdownForEdit(text, opts)
	return c.EditInlineRichMessage(inlineMessageId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditInlineRichMessageHTML replaces an inline message with rich HTML.
func (c *Client) EditInlineRichMessageHTML(inlineMessageId, text string, opts *EditRichTextOpts) error {
	message, markup := richHTMLForEdit(text, opts)
	return c.EditInlineRichMessage(inlineMessageId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditEphemeralRichMessage replaces an ephemeral message with rich content;
// for bots only. receiverUserId identifies the user who can see the message.
func (c *Client) EditEphemeralRichMessage(chatId int64, ephemeralMessageId int32, receiverUserId int64, message *InputRichMessage, opts *EditRichMessageOpts) error {
	if opts == nil {
		opts = &EditRichMessageOpts{}
	}
	return c.EditEphemeralMessage(chatId, ephemeralMessageId, receiverUserId, &EditEphemeralMessageOpts{
		InputMessageContent: &InputMessageRichMessage{Message: message}, ReplyMarkup: opts.ReplyMarkup,
	})
}

// EditEphemeralRichMessageMarkdown replaces an ephemeral message with rich Markdown.
func (c *Client) EditEphemeralRichMessageMarkdown(chatId int64, ephemeralMessageId int32, receiverUserId int64, text string, opts *EditRichTextOpts) error {
	message, markup := richMarkdownForEdit(text, opts)
	return c.EditEphemeralRichMessage(chatId, ephemeralMessageId, receiverUserId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

// EditEphemeralRichMessageHTML replaces an ephemeral message with rich HTML.
func (c *Client) EditEphemeralRichMessageHTML(chatId int64, ephemeralMessageId int32, receiverUserId int64, text string, opts *EditRichTextOpts) error {
	message, markup := richHTMLForEdit(text, opts)
	return c.EditEphemeralRichMessage(chatId, ephemeralMessageId, receiverUserId, message, &EditRichMessageOpts{ReplyMarkup: markup})
}

func richMarkdownForEdit(text string, opts *EditRichTextOpts) (*InputRichMessage, ReplyMarkup) {
	if opts == nil {
		opts = &EditRichTextOpts{}
	}
	return &InputRichMessage{
		Source: &RichMessageSourceMarkdown{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, opts.ReplyMarkup
}

func richHTMLForEdit(text string, opts *EditRichTextOpts) (*InputRichMessage, ReplyMarkup) {
	if opts == nil {
		opts = &EditRichTextOpts{}
	}
	return &InputRichMessage{
		Source: &RichMessageSourceHtml{Text: text, Media: opts.Media},
		IsRtl:  opts.IsRtl, DetectAutomaticBlocks: opts.DetectAutomaticBlocks,
	}, opts.ReplyMarkup
}

// EditRichMessage replaces this message with rich content.
func (m *Message) EditRichMessage(c *Client, message *InputRichMessage, opts *EditRichMessageOpts) (*Message, error) {
	return c.EditRichMessage(m.ChatId, m.Id, message, opts)
}

// EditRichMessageMarkdown replaces this message with rich Markdown.
func (m *Message) EditRichMessageMarkdown(c *Client, text string, opts *EditRichTextOpts) (*Message, error) {
	return c.EditRichMessageMarkdown(m.ChatId, m.Id, text, opts)
}

// EditRichMessageHTML replaces this message with rich HTML.
func (m *Message) EditRichMessageHTML(c *Client, text string, opts *EditRichTextOpts) (*Message, error) {
	return c.EditRichMessageHTML(m.ChatId, m.Id, text, opts)
}

// EditEphemeralMessageTextOpts contains optional parameters for EditEphemeralMessageText.
type EditEphemeralMessageTextOpts struct {
	ParseMode             string
	Entities              []TextEntity
	DisableWebPagePreview bool
	Url                   string
	ForceSmallMedia       bool
	ForceLargeMedia       bool
	ShowAboveText         bool
	ReplyMarkup           ReplyMarkup
}

// EditEphemeralMessageText edits the text of an ephemeral message.
func (c *Client) EditEphemeralMessageText(chatId int64, ephemeralMessageId int32, receiverUserId int64, text string, opts *EditEphemeralMessageTextOpts) error {
	if opts == nil {
		opts = &EditEphemeralMessageTextOpts{}
	}

	formattedText, err := c.GetFormattedText(text, opts.Entities, opts.ParseMode)
	if err != nil {
		return err
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

	return c.EditEphemeralMessage(chatId, ephemeralMessageId, receiverUserId, &EditEphemeralMessageOpts{
		InputMessageContent: content,
		ReplyMarkup:         opts.ReplyMarkup,
	})
}

// EditEphemeralMessageMediaOpts contains optional parameters for EditEphemeralMessageMedia.
type EditEphemeralMessageMediaOpts struct {
	ReplyMarkup ReplyMarkup
}

// EditEphemeralMessageMedia edits the media content of an ephemeral message.
func (c *Client) EditEphemeralMessageMedia(chatId int64, ephemeralMessageId int32, receiverUserId int64, content InputMessageContent, opts *EditEphemeralMessageMediaOpts) error {
	if opts == nil {
		opts = &EditEphemeralMessageMediaOpts{}
	}

	return c.EditEphemeralMessage(chatId, ephemeralMessageId, receiverUserId, &EditEphemeralMessageOpts{
		InputMessageContent: content,
		ReplyMarkup:         opts.ReplyMarkup,
	})
}

// EditEphemeralMessageReplyMarkup edits the reply markup of an ephemeral message.
func (c *Client) EditEphemeralMessageReplyMarkup(chatId int64, ephemeralMessageId int32, receiverUserId int64, opts *EditEphemeralMessageOpts) error {
	copyOpts := EditEphemeralMessageOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	copyOpts.InputMessageContent = nil
	return c.EditEphemeralMessage(chatId, ephemeralMessageId, receiverUserId, &copyOpts)
}

// EditMessageTextOpts contains optional parameters for EditMessageText.
type EditMessageTextOpts struct {
	ParseMode             string
	Entities              []TextEntity
	DisableWebPagePreview bool
	Url                   string
	ForceSmallMedia       bool
	ForceLargeMedia       bool
	ShowAboveText         bool
	ReplyMarkup           ReplyMarkup
}

func (c *Client) loadMessageBeforeEdit(chatId, messageId int64) error {
	if c.config != nil && c.config.UseMessageDatabase != nil && !*c.config.UseMessageDatabase {
		_, err := c.GetMessage(chatId, messageId)
		return err
	}
	return nil
}

func (c *Client) textContentForEdit(text string, opts *EditMessageTextOpts) (*InputMessageText, error) {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}
	formatted, err := c.GetFormattedText(text, opts.Entities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessageText{Text: formatted, LinkPreviewOptions: &LinkPreviewOptions{
		IsDisabled: opts.DisableWebPagePreview, Url: opts.Url,
		ForceSmallMedia: opts.ForceSmallMedia, ForceLargeMedia: opts.ForceLargeMedia,
		ShowAboveText: opts.ShowAboveText,
	}}, nil
}

// EditInlineText replaces an inline message's text with formatted text.
func (c *Client) EditInlineText(inlineMessageId, text string, opts *EditMessageTextOpts) error {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}
	content, err := c.textContentForEdit(text, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageText(inlineMessageId, content, &EditInlineMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineCaption replaces or removes an inline message's formatted caption.
func (c *Client) EditInlineCaption(inlineMessageId, caption string, opts *EditMessageCaptionOpts) error {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}
	formatted, err := c.GetFormattedText(caption, opts.Entities, opts.ParseMode)
	if err != nil {
		return err
	}
	return c.EditInlineMessageCaption(inlineMessageId, &EditInlineMessageCaptionOpts{
		Caption: formatted, ReplyMarkup: opts.ReplyMarkup, ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	})
}

// EditMessageText edits a text message.
func (c *Client) EditMessageText(chatId int64, messageId int64, text string, opts *EditMessageTextOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}

	if err := c.loadMessageBeforeEdit(chatId, messageId); err != nil {
		return nil, err
	}

	content, err := c.textContentForEdit(text, opts)
	if err != nil {
		return nil, err
	}

	return c.EditMessageTextContent(chatId, content, messageId, &EditMessageTextContentOpts{
		ReplyMarkup: opts.ReplyMarkup,
	})
}

// EditMessageCaptionOpts contains optional parameters for EditMessageCaption.
type EditMessageCaptionOpts struct {
	ParseMode             string
	Entities              []TextEntity
	ShowCaptionAboveMedia bool
	ReplyMarkup           ReplyMarkup
}

// EditMessageCaption edits the caption of a message.
func (c *Client) EditMessageCaption(chatId int64, messageId int64, caption string, opts *EditMessageCaptionOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}

	if err := c.loadMessageBeforeEdit(chatId, messageId); err != nil {
		return nil, err
	}

	formattedText, err := c.GetFormattedText(caption, opts.Entities, opts.ParseMode)
	if err != nil {
		return nil, err
	}

	return c.EditMessageCaptionText(chatId, messageId, &EditMessageCaptionTextOpts{
		Caption:               formattedText,
		ReplyMarkup:           opts.ReplyMarkup,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
	})
}

// EditContent replaces text, rich content, media, a live location, or a checklist.
// Both value and pointer content types are accepted. Pass nil opts for no markup.
func (c *Client) EditContent(chatId int64, messageId int64, content InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	content, err := normalizeMessageContent(content)
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditMessageMediaOpts{}
	}
	switch content.(type) {
	case *InputMessageText, *InputMessageRichMessage, *InputMessageAnimation, *InputMessageAudio, *InputMessageDocument, *InputMessagePhoto, *InputMessageVideo, *InputMessageLiveLocation, *InputMessageChecklist:
	default:
		return nil, fmt.Errorf("unsupported content type for editing: %T", content)
	}
	if err := c.loadMessageBeforeEdit(chatId, messageId); err != nil {
		return nil, err
	}
	switch t := content.(type) {
	case *InputMessageText, *InputMessageRichMessage:
		return c.EditMessageTextContent(chatId, content, messageId, &EditMessageTextContentOpts{
			ReplyMarkup: opts.ReplyMarkup,
		})
	case *InputMessageAnimation, *InputMessageAudio, *InputMessageDocument, *InputMessagePhoto, *InputMessageVideo:
		return c.EditMessageMedia(chatId, content, messageId, opts)
	case *InputMessageLiveLocation:
		return c.EditMessageLiveLocation(chatId, messageId, &EditMessageLiveLocationOpts{
			Location:    t.Location,
			ReplyMarkup: opts.ReplyMarkup,
		})
	case *InputMessageChecklist:
		return c.EditMessageChecklist(chatId, t.Checklist, messageId, &EditMessageChecklistOpts{
			ReplyMarkup: opts.ReplyMarkup,
		})
	default:
		return nil, fmt.Errorf("unsupported content type for editing: %T", content)
	}
}

func normalizeMessageContent(content InputMessageContent) (InputMessageContent, error) {
	if content == nil {
		return nil, fmt.Errorf("message content required")
	}
	value := reflect.ValueOf(content)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil, fmt.Errorf("message content required")
		}
		return content, nil
	}
	pointer := reflect.New(value.Type())
	pointer.Elem().Set(value)
	return pointer.Interface().(InputMessageContent), nil
}

// EditInlineContent replaces editable content in an inline message.
// New files cannot be uploaded when editing an inline message.
func (c *Client) EditInlineContent(inlineMessageId string, content InputMessageContent, opts *EditInlineMessageMediaOpts) error {
	content, err := normalizeMessageContent(content)
	if err != nil {
		return err
	}
	if opts == nil {
		opts = &EditInlineMessageMediaOpts{}
	}
	switch value := content.(type) {
	case *InputMessageText, *InputMessageRichMessage:
		return c.EditInlineMessageText(inlineMessageId, content, &EditInlineMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
	case *InputMessageAnimation, *InputMessageAudio, *InputMessageDocument, *InputMessagePhoto, *InputMessageVideo:
		return c.EditInlineMessageMedia(inlineMessageId, content, opts)
	case *InputMessageLiveLocation:
		return c.EditInlineMessageLiveLocation(inlineMessageId, &EditInlineMessageLiveLocationOpts{Location: value.Location, ReplyMarkup: opts.ReplyMarkup})
	default:
		return fmt.Errorf("unsupported content type for inline editing: %T", content)
	}
}

// EditBusinessContent replaces editable content in a business message.
func (c *Client) EditBusinessContent(connectionId string, chatId, messageId int64, content InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	content, err := normalizeMessageContent(content)
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = &EditBusinessMessageMediaOpts{}
	}
	switch value := content.(type) {
	case *InputMessageText, *InputMessageRichMessage:
		return c.EditBusinessMessageText(connectionId, chatId, content, messageId, &EditBusinessMessageTextOpts{ReplyMarkup: opts.ReplyMarkup})
	case *InputMessageAnimation, *InputMessageAudio, *InputMessageDocument, *InputMessagePhoto, *InputMessageVideo:
		return c.EditBusinessMessageMedia(connectionId, chatId, content, messageId, opts)
	case *InputMessageLiveLocation:
		return c.EditBusinessMessageLiveLocation(connectionId, chatId, messageId, &EditBusinessMessageLiveLocationOpts{Location: value.Location, ReplyMarkup: opts.ReplyMarkup})
	case *InputMessageChecklist:
		return c.EditBusinessMessageChecklist(connectionId, chatId, value.Checklist, messageId, &EditBusinessMessageChecklistOpts{ReplyMarkup: opts.ReplyMarkup})
	default:
		return nil, fmt.Errorf("unsupported content type for business editing: %T", content)
	}
}

// StopInlineMessageLiveLocation stops sharing an inline message's live location.
// Location in opts is ignored; ReplyMarkup is applied to the stopped message.
func (c *Client) StopInlineMessageLiveLocation(inlineMessageId string, opts *EditInlineMessageLiveLocationOpts) error {
	copyOpts := EditInlineMessageLiveLocationOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	copyOpts.Location = nil
	return c.EditInlineMessageLiveLocation(inlineMessageId, &copyOpts)
}

// DeleteMessage deletes a message. Set Revoke to delete it for all chat members.
func (c *Client) DeleteMessage(chatId, messageId int64, opts *DeleteMessagesOpts) error {
	return c.DeleteMessages(chatId, []int64{messageId}, opts)
}

// StopMessageLiveLocation stops an active live location.
// Location in opts is ignored; ReplyMarkup is applied to the stopped message.
func (c *Client) StopMessageLiveLocation(chatId, messageId int64, opts *EditMessageLiveLocationOpts) (*Message, error) {
	copyOpts := EditMessageLiveLocationOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	copyOpts.Location = nil
	return c.EditMessageLiveLocation(chatId, messageId, &copyOpts)
}

// EditGeneralForumTopic edits the General topic (native topic ID 1).
func (c *Client) EditGeneralForumTopic(chatId int64, name string) error {
	return c.EditForumTopic(chatId, 1, &EditForumTopicOpts{Name: name})
}

// EditMessagePhotoOpts controls the replacement photo, its caption, and reply markup.
type EditMessagePhotoOpts struct {
	// Caption replaces the current caption; an empty string removes it.
	Caption string
	// CaptionEntities takes precedence over ParseMode.
	CaptionEntities []TextEntity
	// ParseMode defaults to the client's configured parse mode.
	ParseMode             string
	ReplyMarkup           ReplyMarkup
	AddedStickerFileIds   []int32
	Width                 int32
	Height                int32
	Thumbnail             *InputThumbnail
	Video                 InputFile
	HasSpoiler            bool
	ShowCaptionAboveMedia bool
	SelfDestructType      MessageSelfDestructType
}

func (c *Client) photoContentForEdit(photo InputFile, opts *EditMessagePhotoOpts) (*InputMessagePhoto, error) {
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessagePhoto{Photo: &InputPhoto{Photo: photo,
		Video:               opts.Video,
		Thumbnail:           opts.Thumbnail,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Width:               opts.Width,
		Height:              opts.Height}, Caption: caption, HasSpoiler: opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
		SelfDestructType:      opts.SelfDestructType}, nil
}

// EditMessagePhoto replaces the photo and caption of a chat message.
func (c *Client) EditMessagePhoto(chatId, messageId int64, photo InputFile, opts *EditMessagePhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessagePhotoOpts{}
	}
	content, err := c.photoContentForEdit(photo, opts)
	if err != nil {
		return nil, err
	}
	return c.EditContent(chatId, messageId, content, &EditMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineMessagePhoto replaces the photo and caption of an inline message.
// Files must already be uploaded; inline edits cannot upload new files.
func (c *Client) EditInlineMessagePhoto(inlineMessageId string, photo InputFile, opts *EditMessagePhotoOpts) error {
	if opts == nil {
		opts = &EditMessagePhotoOpts{}
	}
	content, err := c.photoContentForEdit(photo, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageMedia(inlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageVideoOpts controls the replacement video, its caption, and reply markup.
type EditMessageVideoOpts struct {
	// Caption replaces the current caption; an empty string removes it.
	Caption string
	// CaptionEntities takes precedence over ParseMode.
	CaptionEntities []TextEntity
	// ParseMode defaults to the client's configured parse mode.
	ParseMode             string
	ReplyMarkup           ReplyMarkup
	AddedStickerFileIds   []int32
	Width                 int32
	Height                int32
	Duration              int32
	Thumbnail             *InputThumbnail
	Cover                 InputFile
	StartTimestamp        int32
	SupportsStreaming     bool
	HasSpoiler            bool
	ShowCaptionAboveMedia bool
	SelfDestructType      MessageSelfDestructType
}

func (c *Client) videoContentForEdit(video InputFile, opts *EditMessageVideoOpts) (*InputMessageVideo, error) {
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessageVideo{Video: &InputVideo{Video: video,
		Thumbnail:           opts.Thumbnail,
		Cover:               opts.Cover,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Width:               opts.Width,
		Height:              opts.Height,
		Duration:            opts.Duration,
		StartTimestamp:      opts.StartTimestamp,
		SupportsStreaming:   opts.SupportsStreaming}, Caption: caption, HasSpoiler: opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia,
		SelfDestructType:      opts.SelfDestructType}, nil
}

// EditMessageVideo replaces the video and caption of a chat message.
func (c *Client) EditMessageVideo(chatId, messageId int64, video InputFile, opts *EditMessageVideoOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageVideoOpts{}
	}
	content, err := c.videoContentForEdit(video, opts)
	if err != nil {
		return nil, err
	}
	return c.EditContent(chatId, messageId, content, &EditMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineMessageVideo replaces the video and caption of an inline message.
// Files must already be uploaded; inline edits cannot upload new files.
func (c *Client) EditInlineMessageVideo(inlineMessageId string, video InputFile, opts *EditMessageVideoOpts) error {
	if opts == nil {
		opts = &EditMessageVideoOpts{}
	}
	content, err := c.videoContentForEdit(video, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageMedia(inlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageAnimationOpts controls the replacement animation, its caption, and reply markup.
type EditMessageAnimationOpts struct {
	// Caption replaces the current caption; an empty string removes it.
	Caption string
	// CaptionEntities takes precedence over ParseMode.
	CaptionEntities []TextEntity
	// ParseMode defaults to the client's configured parse mode.
	ParseMode             string
	ReplyMarkup           ReplyMarkup
	AddedStickerFileIds   []int32
	Width                 int32
	Height                int32
	Duration              int32
	Thumbnail             *InputThumbnail
	HasSpoiler            bool
	ShowCaptionAboveMedia bool
}

func (c *Client) animationContentForEdit(animation InputFile, opts *EditMessageAnimationOpts) (*InputMessageAnimation, error) {
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessageAnimation{Animation: &InputAnimation{Animation: animation,
		Thumbnail:           opts.Thumbnail,
		AddedStickerFileIds: opts.AddedStickerFileIds,
		Width:               opts.Width,
		Height:              opts.Height,
		Duration:            opts.Duration}, Caption: caption, HasSpoiler: opts.HasSpoiler,
		ShowCaptionAboveMedia: opts.ShowCaptionAboveMedia}, nil
}

// EditMessageAnimation replaces the animation and caption of a chat message.
func (c *Client) EditMessageAnimation(chatId, messageId int64, animation InputFile, opts *EditMessageAnimationOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageAnimationOpts{}
	}
	content, err := c.animationContentForEdit(animation, opts)
	if err != nil {
		return nil, err
	}
	return c.EditContent(chatId, messageId, content, &EditMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineMessageAnimation replaces the animation and caption of an inline message.
// Files must already be uploaded; inline edits cannot upload new files.
func (c *Client) EditInlineMessageAnimation(inlineMessageId string, animation InputFile, opts *EditMessageAnimationOpts) error {
	if opts == nil {
		opts = &EditMessageAnimationOpts{}
	}
	content, err := c.animationContentForEdit(animation, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageMedia(inlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageAudioOpts controls the replacement audio, its caption, and reply markup.
type EditMessageAudioOpts struct {
	// Caption replaces the current caption; an empty string removes it.
	Caption string
	// CaptionEntities takes precedence over ParseMode.
	CaptionEntities []TextEntity
	// ParseMode defaults to the client's configured parse mode.
	ParseMode           string
	ReplyMarkup         ReplyMarkup
	Duration            int32
	AlbumCoverThumbnail *InputThumbnail
	Title               string
	Performer           string
}

func (c *Client) audioContentForEdit(audio InputFile, opts *EditMessageAudioOpts) (*InputMessageAudio, error) {
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessageAudio{Audio: &InputAudio{Audio: audio,
		AlbumCoverThumbnail: opts.AlbumCoverThumbnail,
		Title:               opts.Title,
		Performer:           opts.Performer,
		Duration:            opts.Duration}, Caption: caption}, nil
}

// EditMessageAudio replaces the audio and caption of a chat message.
func (c *Client) EditMessageAudio(chatId, messageId int64, audio InputFile, opts *EditMessageAudioOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageAudioOpts{}
	}
	content, err := c.audioContentForEdit(audio, opts)
	if err != nil {
		return nil, err
	}
	return c.EditContent(chatId, messageId, content, &EditMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineMessageAudio replaces the audio and caption of an inline message.
// Files must already be uploaded; inline edits cannot upload new files.
func (c *Client) EditInlineMessageAudio(inlineMessageId string, audio InputFile, opts *EditMessageAudioOpts) error {
	if opts == nil {
		opts = &EditMessageAudioOpts{}
	}
	content, err := c.audioContentForEdit(audio, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageMedia(inlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditMessageDocumentOpts controls the replacement document, its caption, and reply markup.
type EditMessageDocumentOpts struct {
	// Caption replaces the current caption; an empty string removes it.
	Caption string
	// CaptionEntities takes precedence over ParseMode.
	CaptionEntities []TextEntity
	// ParseMode defaults to the client's configured parse mode.
	ParseMode                   string
	ReplyMarkup                 ReplyMarkup
	Thumbnail                   *InputThumbnail
	DisableContentTypeDetection bool
}

func (c *Client) documentContentForEdit(document InputFile, opts *EditMessageDocumentOpts) (*InputMessageDocument, error) {
	caption, err := c.GetFormattedText(opts.Caption, opts.CaptionEntities, opts.ParseMode)
	if err != nil {
		return nil, err
	}
	return &InputMessageDocument{Document: &InputDocument{Document: document,
		Thumbnail:                   opts.Thumbnail,
		DisableContentTypeDetection: opts.DisableContentTypeDetection}, Caption: caption}, nil
}

// EditMessageDocument replaces the document and caption of a chat message.
func (c *Client) EditMessageDocument(chatId, messageId int64, document InputFile, opts *EditMessageDocumentOpts) (*Message, error) {
	if opts == nil {
		opts = &EditMessageDocumentOpts{}
	}
	content, err := c.documentContentForEdit(document, opts)
	if err != nil {
		return nil, err
	}
	return c.EditContent(chatId, messageId, content, &EditMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}

// EditInlineMessageDocument replaces the document and caption of an inline message.
// Files must already be uploaded; inline edits cannot upload new files.
func (c *Client) EditInlineMessageDocument(inlineMessageId string, document InputFile, opts *EditMessageDocumentOpts) error {
	if opts == nil {
		opts = &EditMessageDocumentOpts{}
	}
	content, err := c.documentContentForEdit(document, opts)
	if err != nil {
		return err
	}
	return c.EditInlineMessageMedia(inlineMessageId, content, &EditInlineMessageMediaOpts{ReplyMarkup: opts.ReplyMarkup})
}
