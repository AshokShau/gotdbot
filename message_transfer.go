package gotdbot

// CopyMessageOpts contains optional parameters for CopyMessage.
type CopyMessageOpts struct {
	InGameShare                       bool
	ReplaceCaption                    bool
	NewCaption                        string
	NewCaptionEntities                []TextEntity
	ParseMode                         string
	DisableNotification               bool
	ProtectContent                    bool
	AllowPaidBroadcast                bool
	TopicId                           MessageTopic
	Quote                             *InputTextQuote
	ReplyTo                           InputMessageReplyTo
	ReplyMarkup                       ReplyMarkup
	ReplyToMessageID                  int64
	EffectId                          int64
	NewShowCaptionAboveMedia          bool
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
	ReplaceVideoStartTimestamp  bool
	NewVideoStartTimestamp      int32
}

// CopyMessage copies a message to chat.
func (c *Client) CopyMessage(chatId int64, fromChatId int64, messageId int64, opts *CopyMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &CopyMessageOpts{}
	}

	var caption *FormattedText
	if opts.ReplaceCaption {
		var err error
		caption, err = c.GetFormattedText(opts.NewCaption, opts.NewCaptionEntities, opts.ParseMode)
		if err != nil {
			return nil, err
		}
	}
	content := &InputMessageForwarded{
		ReplaceVideoStartTimestamp: opts.ReplaceVideoStartTimestamp,
		NewVideoStartTimestamp:     opts.NewVideoStartTimestamp,
		FromChatId:                 fromChatId,
		MessageId:                  messageId,
		InGameShare:                opts.InGameShare,
		CopyOptions: &MessageCopyOptions{
			SendCopy:                 true,
			ReplaceCaption:           opts.ReplaceCaption,
			NewCaption:               caption,
			NewShowCaptionAboveMedia: opts.NewShowCaptionAboveMedia,
		},
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

// ForwardMessageOpts contains optional parameters for ForwardMessage.
type ForwardMessageOpts struct {
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

	CallbackQueryID             int64
	ReceiverUserID              int64
	ReplaceCallbackQueryMessage bool
	AllowPaidBroadcast          bool
	SuggestedPostInfo           *InputSuggestedPostInfo
	TopicId                     MessageTopic
	Quote                       *InputTextQuote
	ReplyTo                     InputMessageReplyTo
	ReplyToMessageID            int64
	ReplyMarkup                 ReplyMarkup
}

// ForwardMessage forwards a message to chat.
func (c *Client) ForwardMessage(chatId int64, fromChatId int64, messageId int64, opts *ForwardMessageOpts) (*Message, error) {
	if opts == nil {
		opts = &ForwardMessageOpts{}
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

// CopyMessages sends copies of the specified messages without forwarding attribution.
func (c *Client) CopyMessages(chatId, fromChatId int64, messageIds []int64, opts *ForwardMessagesOpts) (*Messages, error) {
	copyOpts := ForwardMessagesOpts{}
	if opts != nil {
		copyOpts = *opts
	}
	copyOpts.SendCopy = true
	return c.ForwardMessages(chatId, fromChatId, messageIds, &copyOpts)
}
