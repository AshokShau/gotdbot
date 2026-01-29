package types

// SetAuthenticationPhoneNumberOpts contains optional parameters for SetAuthenticationPhoneNumber
type SetAuthenticationPhoneNumberOpts struct {
	Settings *PhoneNumberAuthenticationSettings
}

// ResendAuthenticationCodeOpts contains optional parameters for ResendAuthenticationCode
type ResendAuthenticationCodeOpts struct {
	Reason *ResendCodeReason
}

// GetRemoteFileOpts contains optional parameters for GetRemoteFile
type GetRemoteFileOpts struct {
	FileType *FileType
}

// LoadChatsOpts contains optional parameters for LoadChats
type LoadChatsOpts struct {
	ChatList *ChatList
}

// GetChatsOpts contains optional parameters for GetChats
type GetChatsOpts struct {
	ChatList *ChatList
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

// SearchSecretMessagesOpts contains optional parameters for SearchSecretMessages
type SearchSecretMessagesOpts struct {
	Filter *SearchMessagesFilter
}

// SearchSavedMessagesOpts contains optional parameters for SearchSavedMessages
type SearchSavedMessagesOpts struct {
	Tag *ReactionType
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

// SendInlineQueryResultMessageOpts contains optional parameters for SendInlineQueryResultMessage
type SendInlineQueryResultMessageOpts struct {
	TopicId *MessageTopic
	ReplyTo *InputMessageReplyTo
	Options *MessageSendOptions
}

// ForwardMessagesOpts contains optional parameters for ForwardMessages
type ForwardMessagesOpts struct {
	TopicId *MessageTopic
	Options *MessageSendOptions
}

// ResendMessagesOpts contains optional parameters for ResendMessages
type ResendMessagesOpts struct {
	Quote *InputTextQuote
}

// AddLocalMessageOpts contains optional parameters for AddLocalMessage
type AddLocalMessageOpts struct {
	ReplyTo *InputMessageReplyTo
}

// EditMessageTextOpts contains optional parameters for EditMessageText
type EditMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageLiveLocationOpts contains optional parameters for EditMessageLiveLocation
type EditMessageLiveLocationOpts struct {
	ReplyMarkup *ReplyMarkup
	Location    *Location
}

// EditMessageChecklistOpts contains optional parameters for EditMessageChecklist
type EditMessageChecklistOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageMediaOpts contains optional parameters for EditMessageMedia
type EditMessageMediaOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageCaptionOpts contains optional parameters for EditMessageCaption
type EditMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditMessageReplyMarkupOpts contains optional parameters for EditMessageReplyMarkup
type EditMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditInlineMessageTextOpts contains optional parameters for EditInlineMessageText
type EditInlineMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
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

// EditInlineMessageCaptionOpts contains optional parameters for EditInlineMessageCaption
type EditInlineMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditInlineMessageReplyMarkupOpts contains optional parameters for EditInlineMessageReplyMarkup
type EditInlineMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditMessageSchedulingStateOpts contains optional parameters for EditMessageSchedulingState
type EditMessageSchedulingStateOpts struct {
	SchedulingState *MessageSchedulingState
}

// SetMessageFactCheckOpts contains optional parameters for SetMessageFactCheck
type SetMessageFactCheckOpts struct {
	Text *FormattedText
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

// EditBusinessMessageTextOpts contains optional parameters for EditBusinessMessageText
type EditBusinessMessageTextOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageLiveLocationOpts contains optional parameters for EditBusinessMessageLiveLocation
type EditBusinessMessageLiveLocationOpts struct {
	ReplyMarkup *ReplyMarkup
	Location    *Location
}

// EditBusinessMessageChecklistOpts contains optional parameters for EditBusinessMessageChecklist
type EditBusinessMessageChecklistOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageMediaOpts contains optional parameters for EditBusinessMessageMedia
type EditBusinessMessageMediaOpts struct {
	ReplyMarkup *ReplyMarkup
}

// EditBusinessMessageCaptionOpts contains optional parameters for EditBusinessMessageCaption
type EditBusinessMessageCaptionOpts struct {
	ReplyMarkup *ReplyMarkup
	Caption     *FormattedText
}

// EditBusinessMessageReplyMarkupOpts contains optional parameters for EditBusinessMessageReplyMarkup
type EditBusinessMessageReplyMarkupOpts struct {
	ReplyMarkup *ReplyMarkup
}

// StopBusinessPollOpts contains optional parameters for StopBusinessPoll
type StopBusinessPollOpts struct {
	ReplyMarkup *ReplyMarkup
}

// SetBusinessAccountProfilePhotoOpts contains optional parameters for SetBusinessAccountProfilePhoto
type SetBusinessAccountProfilePhotoOpts struct {
	Photo *InputChatPhoto
}

// AddPendingPaidMessageReactionOpts contains optional parameters for AddPendingPaidMessageReaction
type AddPendingPaidMessageReactionOpts struct {
	TypeField *PaidReactionType
}

// GetMessageAddedReactionsOpts contains optional parameters for GetMessageAddedReactions
type GetMessageAddedReactionsOpts struct {
	ReactionType *ReactionType
}

// StopPollOpts contains optional parameters for StopPoll
type StopPollOpts struct {
	ReplyMarkup *ReplyMarkup
}

// GetInlineQueryResultsOpts contains optional parameters for GetInlineQueryResults
type GetInlineQueryResultsOpts struct {
	UserLocation *Location
}

// AnswerInlineQueryOpts contains optional parameters for AnswerInlineQuery
type AnswerInlineQueryOpts struct {
	Button *InlineQueryResultsButton
}

// OpenWebAppOpts contains optional parameters for OpenWebApp
type OpenWebAppOpts struct {
	TopicId *MessageTopic
	ReplyTo *InputMessageReplyTo
}

// SendChatActionOpts contains optional parameters for SendChatAction
type SendChatActionOpts struct {
	Action *ChatAction
}

// ViewMessagesOpts contains optional parameters for ViewMessages
type ViewMessagesOpts struct {
	Source *MessageSource
}

// CreateNewSupergroupChatOpts contains optional parameters for CreateNewSupergroupChat
type CreateNewSupergroupChatOpts struct {
	Location *ChatLocation
}

// SetChatPhotoOpts contains optional parameters for SetChatPhoto
type SetChatPhotoOpts struct {
	Photo *InputChatPhoto
}

// SetChatEmojiStatusOpts contains optional parameters for SetChatEmojiStatus
type SetChatEmojiStatusOpts struct {
	EmojiStatus *EmojiStatus
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

// SearchChatMembersOpts contains optional parameters for SearchChatMembers
type SearchChatMembersOpts struct {
	Filter *ChatMembersFilter
}

// GetChatNotificationSettingsExceptionsOpts contains optional parameters for GetChatNotificationSettingsExceptions
type GetChatNotificationSettingsExceptionsOpts struct {
	Scope *NotificationSettingsScope
}

// PostStoryOpts contains optional parameters for PostStory
type PostStoryOpts struct {
	Areas           *InputStoryAreas
	Caption         *FormattedText
	FromStoryFullId *StoryFullId
}

// EditStoryOpts contains optional parameters for EditStory
type EditStoryOpts struct {
	Content *InputStoryContent
	Areas   *InputStoryAreas
	Caption *FormattedText
}

// SetStoryReactionOpts contains optional parameters for SetStoryReaction
type SetStoryReactionOpts struct {
	ReactionType *ReactionType
}

// GetChatStoryInteractionsOpts contains optional parameters for GetChatStoryInteractions
type GetChatStoryInteractionsOpts struct {
	ReactionType *ReactionType
}

// PreliminaryUploadFileOpts contains optional parameters for PreliminaryUploadFile
type PreliminaryUploadFileOpts struct {
	FileType *FileType
}

// FinishFileGenerationOpts contains optional parameters for FinishFileGeneration
type FinishFileGenerationOpts struct {
	Error *Error
}

// GetChatInviteLinkMembersOpts contains optional parameters for GetChatInviteLinkMembers
type GetChatInviteLinkMembersOpts struct {
	OffsetMember *ChatInviteLinkMember
}

// GetChatJoinRequestsOpts contains optional parameters for GetChatJoinRequests
type GetChatJoinRequestsOpts struct {
	OffsetRequest *ChatJoinRequest
}

// JoinVideoChatOpts contains optional parameters for JoinVideoChat
type JoinVideoChatOpts struct {
	ParticipantId *MessageSender
}

// GetGroupCallStreamSegmentOpts contains optional parameters for GetGroupCallStreamSegment
type GetGroupCallStreamSegmentOpts struct {
	VideoQuality *GroupCallVideoQuality
}

// DecryptGroupCallDataOpts contains optional parameters for DecryptGroupCallData
type DecryptGroupCallDataOpts struct {
	DataChannel *GroupCallDataChannel
}

// SetMessageSenderBlockListOpts contains optional parameters for SetMessageSenderBlockList
type SetMessageSenderBlockListOpts struct {
	BlockList *BlockList
}

// GetLinkPreviewOpts contains optional parameters for GetLinkPreview
type GetLinkPreviewOpts struct {
	LinkPreviewOptions *LinkPreviewOptions
}

// SetBusinessOpeningHoursOpts contains optional parameters for SetBusinessOpeningHours
type SetBusinessOpeningHoursOpts struct {
	OpeningHours *BusinessOpeningHours
}

// SendPhoneNumberCodeOpts contains optional parameters for SendPhoneNumberCode
type SendPhoneNumberCodeOpts struct {
	Settings *PhoneNumberAuthenticationSettings
}

// ResendPhoneNumberCodeOpts contains optional parameters for ResendPhoneNumberCode
type ResendPhoneNumberCodeOpts struct {
	Reason *ResendCodeReason
}

// SetCommandsOpts contains optional parameters for SetCommands
type SetCommandsOpts struct {
	Scope *BotCommandScope
}

// DeleteCommandsOpts contains optional parameters for DeleteCommands
type DeleteCommandsOpts struct {
	Scope *BotCommandScope
}

// GetCommandsOpts contains optional parameters for GetCommands
type GetCommandsOpts struct {
	Scope *BotCommandScope
}

// GetSupergroupMembersOpts contains optional parameters for GetSupergroupMembers
type GetSupergroupMembersOpts struct {
	Filter *SupergroupMembersFilter
}

// GetChatEventLogOpts contains optional parameters for GetChatEventLog
type GetChatEventLogOpts struct {
	Filters *ChatEventLogFilters
}

// GetPaymentFormOpts contains optional parameters for GetPaymentForm
type GetPaymentFormOpts struct {
	Theme *ThemeParameters
}

// ValidateOrderInfoOpts contains optional parameters for ValidateOrderInfo
type ValidateOrderInfoOpts struct {
	OrderInfo *OrderInfo
}

// SendPaymentFormOpts contains optional parameters for SendPaymentForm
type SendPaymentFormOpts struct {
	Credentials *InputCredentials
}

// SetGiftResalePriceOpts contains optional parameters for SetGiftResalePrice
type SetGiftResalePriceOpts struct {
	Price *GiftResalePrice
}

// SetDefaultBackgroundOpts contains optional parameters for SetDefaultBackground
type SetDefaultBackgroundOpts struct {
	Background *InputBackground
	TypeField  *BackgroundType
}

// SetOptionOpts contains optional parameters for SetOption
type SetOptionOpts struct {
	Value *OptionValue
}

// GetTonTransactionsOpts contains optional parameters for GetTonTransactions
type GetTonTransactionsOpts struct {
	Direction *TransactionDirection
}

// SetNetworkTypeOpts contains optional parameters for SetNetworkType
type SetNetworkTypeOpts struct {
	TypeField *NetworkType
}

// SetStickerSetThumbnailOpts contains optional parameters for SetStickerSetThumbnail
type SetStickerSetThumbnailOpts struct {
	Thumbnail *InputFile
	Format    *StickerFormat
}

// SetStickerMaskPositionOpts contains optional parameters for SetStickerMaskPosition
type SetStickerMaskPositionOpts struct {
	MaskPosition *MaskPosition
}

// GetStarTransactionsOpts contains optional parameters for GetStarTransactions
type GetStarTransactionsOpts struct {
	Direction *TransactionDirection
}

// SetChatAffiliateProgramOpts contains optional parameters for SetChatAffiliateProgram
type SetChatAffiliateProgramOpts struct {
	Parameters *AffiliateProgramParameters
}
