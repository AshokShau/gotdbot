package gotdbot

// AddLocalMessageOpts contains optional parameters for AddLocalMessage
type AddLocalMessageOpts struct {
	ReplyTo *InputMessageReplyTo
}

// AddPendingPaidMessageReactionOpts contains optional parameters for AddPendingPaidMessageReaction
type AddPendingPaidMessageReactionOpts struct {
	TypeField *PaidReactionType
}

// AnswerInlineQueryOpts contains optional parameters for AnswerInlineQuery
type AnswerInlineQueryOpts struct {
	Button *InlineQueryResultsButton
}

// CreateNewSupergroupChatOpts contains optional parameters for CreateNewSupergroupChat
type CreateNewSupergroupChatOpts struct {
	Location *ChatLocation
}

// DecryptGroupCallDataOpts contains optional parameters for DecryptGroupCallData
type DecryptGroupCallDataOpts struct {
	DataChannel *GroupCallDataChannel
}

// DeleteCommandsOpts contains optional parameters for DeleteCommands
type DeleteCommandsOpts struct {
	Scope *BotCommandScope
}

// EditBusinessMessageCaptionOpts contains optional parameters for EditBusinessMessageCaption
type EditBusinessMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditBusinessMessageChecklistOpts contains optional parameters for EditBusinessMessageChecklist
type EditBusinessMessageChecklistOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageLiveLocationOpts contains optional parameters for EditBusinessMessageLiveLocation
type EditBusinessMessageLiveLocationOpts struct {
	ReplyMarkup *ReplyMarkup
	Location    *Location
}

// EditBusinessMessageMediaOpts contains optional parameters for EditBusinessMessageMedia
type EditBusinessMessageMediaOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageReplyMarkupOpts contains optional parameters for EditBusinessMessageReplyMarkup
type EditBusinessMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageTextOpts contains optional parameters for EditBusinessMessageText
type EditBusinessMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditInlineMessageCaptionOpts contains optional parameters for EditInlineMessageCaption
type EditInlineMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditInlineMessageLiveLocationOpts contains optional parameters for EditInlineMessageLiveLocation
type EditInlineMessageLiveLocationOpts struct {
	ReplyMarkup *ReplyMarkup
	Location    *Location
}

// EditInlineMessageMediaOpts contains optional parameters for EditInlineMessageMedia
type EditInlineMessageMediaOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditInlineMessageReplyMarkupOpts contains optional parameters for EditInlineMessageReplyMarkup
type EditInlineMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditInlineMessageTextOpts contains optional parameters for EditInlineMessageText
type EditInlineMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageCaptionOpts contains optional parameters for EditMessageCaption
type EditMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditMessageChecklistOpts contains optional parameters for EditMessageChecklist
type EditMessageChecklistOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageLiveLocationOpts contains optional parameters for EditMessageLiveLocation
type EditMessageLiveLocationOpts struct {
	ReplyMarkup *ReplyMarkup
	Location    *Location
}

// EditMessageMediaOpts contains optional parameters for EditMessageMedia
type EditMessageMediaOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageReplyMarkupOpts contains optional parameters for EditMessageReplyMarkup
type EditMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageSchedulingStateOpts contains optional parameters for EditMessageSchedulingState
type EditMessageSchedulingStateOpts struct {
	SchedulingState *MessageSchedulingState
}

// EditMessageTextOpts contains optional parameters for EditMessageText
type EditMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditStoryOpts contains optional parameters for EditStory
type EditStoryOpts struct {
	Content *InputStoryContent
	Areas   *InputStoryAreas
	Caption *FormattedText
}

// FinishFileGenerationOpts contains optional parameters for FinishFileGeneration
type FinishFileGenerationOpts struct {
	Error *Error
}

// ForwardMessagesOpts contains optional parameters for ForwardMessages
type ForwardMessagesOpts struct {
	TopicId *MessageTopic
	Options *MessageSendOptions
}

// GetChatEventLogOpts contains optional parameters for GetChatEventLog
type GetChatEventLogOpts struct {
	Filters *ChatEventLogFilters
}

// GetChatInviteLinkMembersOpts contains optional parameters for GetChatInviteLinkMembers
type GetChatInviteLinkMembersOpts struct {
	OffsetMember *ChatInviteLinkMember
}

// GetChatJoinRequestsOpts contains optional parameters for GetChatJoinRequests
type GetChatJoinRequestsOpts struct {
	OffsetRequest *ChatJoinRequest
}

// GetChatMessageCalendarOpts contains optional parameters for GetChatMessageCalendar
type GetChatMessageCalendarOpts struct {
	TopicId *MessageTopic
}

// GetChatMessageCountOpts contains optional parameters for GetChatMessageCount
type GetChatMessageCountOpts struct {
	TopicId *MessageTopic
}

// GetChatMessagePositionOpts contains optional parameters for GetChatMessagePosition
type GetChatMessagePositionOpts struct {
	TopicId *MessageTopic
}

// GetChatNotificationSettingsExceptionsOpts contains optional parameters for GetChatNotificationSettingsExceptions
type GetChatNotificationSettingsExceptionsOpts struct {
	Scope *NotificationSettingsScope
}

// GetChatsOpts contains optional parameters for GetChats
type GetChatsOpts struct {
	ChatList *ChatList
}

// GetChatStoryInteractionsOpts contains optional parameters for GetChatStoryInteractions
type GetChatStoryInteractionsOpts struct {
	ReactionType *ReactionType
}

// GetCommandsOpts contains optional parameters for GetCommands
type GetCommandsOpts struct {
	Scope *BotCommandScope
}

// GetGroupCallStreamSegmentOpts contains optional parameters for GetGroupCallStreamSegment
type GetGroupCallStreamSegmentOpts struct {
	VideoQuality *GroupCallVideoQuality
}

// GetInlineQueryResultsOpts contains optional parameters for GetInlineQueryResults
type GetInlineQueryResultsOpts struct {
	UserLocation *Location
}

// GetLinkPreviewOpts contains optional parameters for GetLinkPreview
type GetLinkPreviewOpts struct {
	LinkPreviewOptions *LinkPreviewOptions
}

// GetMessageAddedReactionsOpts contains optional parameters for GetMessageAddedReactions
type GetMessageAddedReactionsOpts struct {
	ReactionType *ReactionType
}

// GetPaymentFormOpts contains optional parameters for GetPaymentForm
type GetPaymentFormOpts struct {
	Theme *ThemeParameters
}

// GetRemoteFileOpts contains optional parameters for GetRemoteFile
type GetRemoteFileOpts struct {
	FileType *FileType
}

// GetStarTransactionsOpts contains optional parameters for GetStarTransactions
type GetStarTransactionsOpts struct {
	Direction *TransactionDirection
}

// GetSupergroupMembersOpts contains optional parameters for GetSupergroupMembers
type GetSupergroupMembersOpts struct {
	Filter *SupergroupMembersFilter
}

// GetTonTransactionsOpts contains optional parameters for GetTonTransactions
type GetTonTransactionsOpts struct {
	Direction *TransactionDirection
}

// JoinVideoChatOpts contains optional parameters for JoinVideoChat
type JoinVideoChatOpts struct {
	ParticipantId *MessageSender
}

// LoadChatsOpts contains optional parameters for LoadChats
type LoadChatsOpts struct {
	ChatList *ChatList
}

// OpenWebAppOpts contains optional parameters for OpenWebApp
type OpenWebAppOpts struct {
	TopicId *MessageTopic
	ReplyTo *InputMessageReplyTo
}

// PostStoryOpts contains optional parameters for PostStory
type PostStoryOpts struct {
	Areas           *InputStoryAreas
	Caption         *FormattedText
	FromStoryFullId *StoryFullId
}

// PreliminaryUploadFileOpts contains optional parameters for PreliminaryUploadFile
type PreliminaryUploadFileOpts struct {
	FileType *FileType
}

// ResendAuthenticationCodeOpts contains optional parameters for ResendAuthenticationCode
type ResendAuthenticationCodeOpts struct {
	Reason *ResendCodeReason
}

// ResendMessagesOpts contains optional parameters for ResendMessages
type ResendMessagesOpts struct {
	Quote *InputTextQuote
}

// ResendPhoneNumberCodeOpts contains optional parameters for ResendPhoneNumberCode
type ResendPhoneNumberCodeOpts struct {
	Reason *ResendCodeReason
}

// SearchChatMembersOpts contains optional parameters for SearchChatMembers
type SearchChatMembersOpts struct {
	Filter *ChatMembersFilter
}

// SearchChatMessagesOpts contains optional parameters for SearchChatMessages
type SearchChatMessagesOpts struct {
	TopicId  *MessageTopic
	SenderId *MessageSender
	Filter   *SearchMessagesFilter
}

// SearchMessagesOpts contains optional parameters for SearchMessages
type SearchMessagesOpts struct {
	ChatList       *ChatList
	Filter         *SearchMessagesFilter
	ChatTypeFilter *SearchMessagesChatTypeFilter
}

// SearchSavedMessagesOpts contains optional parameters for SearchSavedMessages
type SearchSavedMessagesOpts struct {
	Tag *ReactionType
}

// SearchSecretMessagesOpts contains optional parameters for SearchSecretMessages
type SearchSecretMessagesOpts struct {
	Filter *SearchMessagesFilter
}

// SendBusinessMessageOpts contains optional parameters for SendBusinessMessage
type SendBusinessMessageOpts struct {
	ReplyTo     *InputMessageReplyTo
	ReplyMarkup *ReplyMarkup
}

// SendBusinessMessageAlbumOpts contains optional parameters for SendBusinessMessageAlbum
type SendBusinessMessageAlbumOpts struct {
	ReplyTo *InputMessageReplyTo
}

// SendChatActionOpts contains optional parameters for SendChatAction
type SendChatActionOpts struct {
	Action *ChatAction
}

// SendInlineQueryResultMessageOpts contains optional parameters for SendInlineQueryResultMessage
type SendInlineQueryResultMessageOpts struct {
	TopicId *MessageTopic
	ReplyTo *InputMessageReplyTo
	Options *MessageSendOptions
}

// SendMessageOpts contains optional parameters for SendMessage
type SendMessageOpts struct {
	TopicId     *MessageTopic
	ReplyTo     *InputMessageReplyTo
	Options     *MessageSendOptions
	ReplyMarkup *ReplyMarkup
}

// SendMessageAlbumOpts contains optional parameters for SendMessageAlbum
type SendMessageAlbumOpts struct {
	TopicId *MessageTopic
	ReplyTo *InputMessageReplyTo
	Options *MessageSendOptions
}

// SendPaymentFormOpts contains optional parameters for SendPaymentForm
type SendPaymentFormOpts struct {
	Credentials *InputCredentials
}

// SendPhoneNumberCodeOpts contains optional parameters for SendPhoneNumberCode
type SendPhoneNumberCodeOpts struct {
	Settings *PhoneNumberAuthenticationSettings
}

// SetAuthenticationPhoneNumberOpts contains optional parameters for SetAuthenticationPhoneNumber
type SetAuthenticationPhoneNumberOpts struct {
	Settings *PhoneNumberAuthenticationSettings
}

// SetBusinessAccountProfilePhotoOpts contains optional parameters for SetBusinessAccountProfilePhoto
type SetBusinessAccountProfilePhotoOpts struct {
	Photo *InputChatPhoto
}

// SetBusinessOpeningHoursOpts contains optional parameters for SetBusinessOpeningHours
type SetBusinessOpeningHoursOpts struct {
	OpeningHours *BusinessOpeningHours
}

// SetChatAffiliateProgramOpts contains optional parameters for SetChatAffiliateProgram
type SetChatAffiliateProgramOpts struct {
	Parameters *AffiliateProgramParameters
}

// SetChatBackgroundOpts contains optional parameters for SetChatBackground
type SetChatBackgroundOpts struct {
	Background *InputBackground
	TypeField  *BackgroundType
}

// SetChatDraftMessageOpts contains optional parameters for SetChatDraftMessage
type SetChatDraftMessageOpts struct {
	TopicId      *MessageTopic
	DraftMessage *DraftMessage
}

// SetChatEmojiStatusOpts contains optional parameters for SetChatEmojiStatus
type SetChatEmojiStatusOpts struct {
	EmojiStatus *EmojiStatus
}

// SetChatPhotoOpts contains optional parameters for SetChatPhoto
type SetChatPhotoOpts struct {
	Photo *InputChatPhoto
}

// SetCommandsOpts contains optional parameters for SetCommands
type SetCommandsOpts struct {
	Scope *BotCommandScope
}

// SetDefaultBackgroundOpts contains optional parameters for SetDefaultBackground
type SetDefaultBackgroundOpts struct {
	Background *InputBackground
	TypeField  *BackgroundType
}

// SetGiftResalePriceOpts contains optional parameters for SetGiftResalePrice
type SetGiftResalePriceOpts struct {
	Price *GiftResalePrice
}

// SetMessageFactCheckOpts contains optional parameters for SetMessageFactCheck
type SetMessageFactCheckOpts struct {
	Text *FormattedText
}

// SetMessageSenderBlockListOpts contains optional parameters for SetMessageSenderBlockList
type SetMessageSenderBlockListOpts struct {
	BlockList *BlockList
}

// SetNetworkTypeOpts contains optional parameters for SetNetworkType
type SetNetworkTypeOpts struct {
	TypeField *NetworkType
}

// SetOptionOpts contains optional parameters for SetOption
type SetOptionOpts struct {
	Value *OptionValue
}

// SetStickerMaskPositionOpts contains optional parameters for SetStickerMaskPosition
type SetStickerMaskPositionOpts struct {
	MaskPosition *MaskPosition
}

// SetStickerSetThumbnailOpts contains optional parameters for SetStickerSetThumbnail
type SetStickerSetThumbnailOpts struct {
	Thumbnail *InputFile
	Format    *StickerFormat
}

// SetStoryReactionOpts contains optional parameters for SetStoryReaction
type SetStoryReactionOpts struct {
	ReactionType *ReactionType
}

// StopBusinessPollOpts contains optional parameters for StopBusinessPoll
type StopBusinessPollOpts struct {
	ReplyMarkup *ReplyMarkup
}

// StopPollOpts contains optional parameters for StopPoll
type StopPollOpts struct {
	ReplyMarkup *ReplyMarkup
}

// ValidateOrderInfoOpts contains optional parameters for ValidateOrderInfo
type ValidateOrderInfoOpts struct {
	OrderInfo *OrderInfo
}

// ViewMessagesOpts contains optional parameters for ViewMessages
type ViewMessagesOpts struct {
	Source *MessageSource
}
