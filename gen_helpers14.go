package gotdbot

// EditForumTopic is a helper method for Client.EditForumTopic
func (t TonTransactionTypeSuggestedPostPayment) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(t.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (t TonTransactionTypeSuggestedPostPayment) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(t.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (t TonTransactionTypeSuggestedPostPayment) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(t.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (t TonTransactionTypeSuggestedPostPayment) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(t.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (t TonTransactionTypeSuggestedPostPayment) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(t.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (t TonTransactionTypeSuggestedPostPayment) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(t.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (t TonTransactionTypeSuggestedPostPayment) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(t.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (t TonTransactionTypeSuggestedPostPayment) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(t.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (t TonTransactionTypeSuggestedPostPayment) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(t.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (t TonTransactionTypeSuggestedPostPayment) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, t.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (t TonTransactionTypeSuggestedPostPayment) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(t.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (t TonTransactionTypeSuggestedPostPayment) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(t.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (t TonTransactionTypeSuggestedPostPayment) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(t.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (t TonTransactionTypeSuggestedPostPayment) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(t.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (t TonTransactionTypeSuggestedPostPayment) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(t.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (t TonTransactionTypeSuggestedPostPayment) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(t.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (t TonTransactionTypeSuggestedPostPayment) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(t.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (t TonTransactionTypeSuggestedPostPayment) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(t.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (t TonTransactionTypeSuggestedPostPayment) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(t.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (t TonTransactionTypeSuggestedPostPayment) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(t.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (t TonTransactionTypeSuggestedPostPayment) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(t.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (t TonTransactionTypeSuggestedPostPayment) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(t.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (t TonTransactionTypeSuggestedPostPayment) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(t.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(t.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (t TonTransactionTypeSuggestedPostPayment) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(t.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (t TonTransactionTypeSuggestedPostPayment) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(t.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (t TonTransactionTypeSuggestedPostPayment) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(t.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (t TonTransactionTypeSuggestedPostPayment) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(t.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (t TonTransactionTypeSuggestedPostPayment) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(t.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (t TonTransactionTypeSuggestedPostPayment) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(t.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (t TonTransactionTypeSuggestedPostPayment) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(t.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (t TonTransactionTypeSuggestedPostPayment) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(t.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (t TonTransactionTypeSuggestedPostPayment) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(t.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (t TonTransactionTypeSuggestedPostPayment) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(t.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (t TonTransactionTypeSuggestedPostPayment) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(t.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (t TonTransactionTypeSuggestedPostPayment) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(t.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (t TonTransactionTypeSuggestedPostPayment) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(t.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (t TonTransactionTypeSuggestedPostPayment) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(t.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (t TonTransactionTypeSuggestedPostPayment) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(t.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (t TonTransactionTypeSuggestedPostPayment) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(t.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (t TonTransactionTypeSuggestedPostPayment) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(t.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (t TonTransactionTypeSuggestedPostPayment) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(t.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (t TonTransactionTypeSuggestedPostPayment) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(t.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (t TonTransactionTypeSuggestedPostPayment) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(t.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (t TonTransactionTypeSuggestedPostPayment) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(t.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (t TonTransactionTypeSuggestedPostPayment) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(t.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (t TonTransactionTypeSuggestedPostPayment) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(t.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (t TonTransactionTypeSuggestedPostPayment) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(t.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (t TonTransactionTypeSuggestedPostPayment) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(t.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (t TonTransactionTypeSuggestedPostPayment) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(t.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (t TonTransactionTypeSuggestedPostPayment) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(t.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (t TonTransactionTypeSuggestedPostPayment) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(t.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (t TonTransactionTypeSuggestedPostPayment) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(t.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (t TonTransactionTypeSuggestedPostPayment) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(t.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TonTransactionTypeSuggestedPostPayment) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(t.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (t TonTransactionTypeSuggestedPostPayment) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(t.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (t TonTransactionTypeSuggestedPostPayment) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, t.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (t TonTransactionTypeSuggestedPostPayment) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(t.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (t TonTransactionTypeSuggestedPostPayment) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(t.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (t TonTransactionTypeSuggestedPostPayment) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(t.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (t TonTransactionTypeSuggestedPostPayment) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(t.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (t TonTransactionTypeSuggestedPostPayment) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, t.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (t TonTransactionTypeSuggestedPostPayment) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(t.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (t TonTransactionTypeSuggestedPostPayment) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(t.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (t TonTransactionTypeSuggestedPostPayment) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(t.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (t TonTransactionTypeSuggestedPostPayment) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(t.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (t TonTransactionTypeSuggestedPostPayment) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(t.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (t TonTransactionTypeSuggestedPostPayment) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(t.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (t TonTransactionTypeSuggestedPostPayment) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(t.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (t TonTransactionTypeSuggestedPostPayment) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(t.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (t TonTransactionTypeSuggestedPostPayment) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(t.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (t TonTransactionTypeSuggestedPostPayment) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(t.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (t TonTransactionTypeSuggestedPostPayment) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(t.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (t TonTransactionTypeSuggestedPostPayment) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(t.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (t TonTransactionTypeSuggestedPostPayment) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(t.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (t TonTransactionTypeSuggestedPostPayment) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(t.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (t TonTransactionTypeSuggestedPostPayment) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(t.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (t TonTransactionTypeSuggestedPostPayment) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(t.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (t TonTransactionTypeSuggestedPostPayment) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(t.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (t TonTransactionTypeSuggestedPostPayment) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(t.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (t TonTransactionTypeSuggestedPostPayment) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(t.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (t TonTransactionTypeSuggestedPostPayment) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(t.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (t TonTransactionTypeSuggestedPostPayment) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, t.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (t TonTransactionTypeSuggestedPostPayment) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(t.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (t TonTransactionTypeSuggestedPostPayment) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(t.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TonTransactionTypeSuggestedPostPayment) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(t.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (t TonTransactionTypeSuggestedPostPayment) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(t.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (t TonTransactionTypeSuggestedPostPayment) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(t.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (t TonTransactionTypeSuggestedPostPayment) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(t.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (t TonTransactionTypeSuggestedPostPayment) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(t.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (t TonTransactionTypeSuggestedPostPayment) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(t.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (t TonTransactionTypeSuggestedPostPayment) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(t.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (t TonTransactionTypeSuggestedPostPayment) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(t.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (t TonTransactionTypeSuggestedPostPayment) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(t.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (t TonTransactionTypeSuggestedPostPayment) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(t.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (t TonTransactionTypeSuggestedPostPayment) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(t.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (t TonTransactionTypeSuggestedPostPayment) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(t.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (t TonTransactionTypeSuggestedPostPayment) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(t.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (t TonTransactionTypeSuggestedPostPayment) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(t.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (t TonTransactionTypeSuggestedPostPayment) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(t.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (t TonTransactionTypeSuggestedPostPayment) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(t.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TonTransactionTypeSuggestedPostPayment) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(t.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (t TonTransactionTypeSuggestedPostPayment) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(t.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (t TonTransactionTypeSuggestedPostPayment) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(t.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (t TonTransactionTypeSuggestedPostPayment) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(t.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (t TonTransactionTypeSuggestedPostPayment) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(t.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (t TonTransactionTypeSuggestedPostPayment) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(t.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (t TonTransactionTypeSuggestedPostPayment) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(t.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (t TonTransactionTypeSuggestedPostPayment) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(t.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (t TonTransactionTypeSuggestedPostPayment) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, t.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (t TonTransactionTypeSuggestedPostPayment) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(t.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (t TonTransactionTypeSuggestedPostPayment) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(t.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (t TonTransactionTypeSuggestedPostPayment) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(t.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (t TonTransactionTypeSuggestedPostPayment) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(t.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (t TonTransactionTypeSuggestedPostPayment) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(t.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (t TonTransactionTypeSuggestedPostPayment) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(t.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (t TonTransactionTypeSuggestedPostPayment) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (t TonTransactionTypeSuggestedPostPayment) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, t.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (t TonTransactionTypeSuggestedPostPayment) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(t.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (t TonTransactionTypeSuggestedPostPayment) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (t TonTransactionTypeSuggestedPostPayment) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(t.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(t.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (t TonTransactionTypeSuggestedPostPayment) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(t.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (t TonTransactionTypeSuggestedPostPayment) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(t.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (t TonTransactionTypeSuggestedPostPayment) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(t.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (t TonTransactionTypeSuggestedPostPayment) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(t.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (t TonTransactionTypeSuggestedPostPayment) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(t.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (t TonTransactionTypeSuggestedPostPayment) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(t.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(t.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (t TonTransactionTypeSuggestedPostPayment) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, t.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (t TonTransactionTypeSuggestedPostPayment) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(t.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (t TonTransactionTypeSuggestedPostPayment) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(t.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (t TonTransactionTypeSuggestedPostPayment) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(t.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (t TonTransactionTypeSuggestedPostPayment) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(t.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (t TonTransactionTypeSuggestedPostPayment) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, t.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (t TonTransactionTypeSuggestedPostPayment) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, t.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (t TonTransactionTypeSuggestedPostPayment) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, t.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (t TonTransactionTypeSuggestedPostPayment) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(t.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (t TonTransactionTypeSuggestedPostPayment) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(t.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (t TonTransactionTypeSuggestedPostPayment) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(t.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (t TonTransactionTypeSuggestedPostPayment) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(t.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (t TonTransactionTypeSuggestedPostPayment) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(t.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TonTransactionTypeSuggestedPostPayment) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(t.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (t TonTransactionTypeSuggestedPostPayment) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, t.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (t TonTransactionTypeSuggestedPostPayment) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(t.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (t TonTransactionTypeSuggestedPostPayment) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(t.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (t TonTransactionTypeSuggestedPostPayment) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(t.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (t TonTransactionTypeSuggestedPostPayment) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(t.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (t TonTransactionTypeSuggestedPostPayment) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(t.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (t TonTransactionTypeSuggestedPostPayment) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(t.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (t TonTransactionTypeSuggestedPostPayment) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(t.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (t TonTransactionTypeSuggestedPostPayment) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(t.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (t TonTransactionTypeSuggestedPostPayment) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(t.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (t TonTransactionTypeSuggestedPostPayment) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(t.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (t TonTransactionTypeSuggestedPostPayment) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(t.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (t TonTransactionTypeSuggestedPostPayment) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(t.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (t TonTransactionTypeSuggestedPostPayment) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(t.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (t TonTransactionTypeSuggestedPostPayment) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(t.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (t TonTransactionTypeSuggestedPostPayment) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(t.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (t TonTransactionTypeSuggestedPostPayment) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(t.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (t TonTransactionTypeSuggestedPostPayment) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(t.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (t TonTransactionTypeSuggestedPostPayment) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(t.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (t TonTransactionTypeSuggestedPostPayment) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(t.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (t TonTransactionTypeSuggestedPostPayment) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(t.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (t TonTransactionTypeSuggestedPostPayment) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(t.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (t TonTransactionTypeSuggestedPostPayment) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(t.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (t TonTransactionTypeSuggestedPostPayment) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(t.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (t TonTransactionTypeSuggestedPostPayment) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(t.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (t TonTransactionTypeSuggestedPostPayment) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(t.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (t TonTransactionTypeSuggestedPostPayment) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(t.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TonTransactionTypeSuggestedPostPayment) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(t.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (t TonTransactionTypeSuggestedPostPayment) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(t.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (t TonTransactionTypeSuggestedPostPayment) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(t.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (t TonTransactionTypeSuggestedPostPayment) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(t.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (t TonTransactionTypeSuggestedPostPayment) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(t.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (t TonTransactionTypeSuggestedPostPayment) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(t.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (t TonTransactionTypeSuggestedPostPayment) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(t.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (t TonTransactionTypeSuggestedPostPayment) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(t.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (t TonTransactionTypeSuggestedPostPayment) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(t.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (t TonTransactionTypeSuggestedPostPayment) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(t.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (t TonTransactionTypeSuggestedPostPayment) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(t.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (t TonTransactionTypeSuggestedPostPayment) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(t.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (t TonTransactionTypeSuggestedPostPayment) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, t.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (t TonTransactionTypeSuggestedPostPayment) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(t.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (t TonTransactionTypeSuggestedPostPayment) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(t.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (t TonTransactionTypeSuggestedPostPayment) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(t.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(t.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(t.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(t.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(t.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, t.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(t.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (t TonTransactionTypeSuggestedPostPayment) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(t.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (t TonTransactionTypeSuggestedPostPayment) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(t.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (t TonTransactionTypeSuggestedPostPayment) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(t.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (t TonTransactionTypeSuggestedPostPayment) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(t.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (t TonTransactionTypeSuggestedPostPayment) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(t.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TonTransactionTypeSuggestedPostPayment) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(t.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (t TonTransactionTypeSuggestedPostPayment) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(t.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (t TonTransactionTypeSuggestedPostPayment) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(t.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (t TonTransactionTypeSuggestedPostPayment) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(t.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (t TonTransactionTypeSuggestedPostPayment) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(t.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (t TonTransactionTypeSuggestedPostPayment) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(t.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (t TonTransactionTypeSuggestedPostPayment) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(t.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (t TonTransactionTypeSuggestedPostPayment) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(t.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TonTransactionTypeUpgradedGiftPurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TonTransactionTypeUpgradedGiftPurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TonTransactionTypeUpgradedGiftPurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TonTransactionTypeUpgradedGiftPurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TonTransactionTypeUpgradedGiftPurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (t TonTransactionTypeUpgradedGiftPurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TonTransactionTypeUpgradedGiftPurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TonTransactionTypeUpgradedGiftPurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TonTransactionTypeUpgradedGiftPurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TonTransactionTypeUpgradedGiftPurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TonTransactionTypeUpgradedGiftPurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TonTransactionTypeUpgradedGiftPurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TonTransactionTypeUpgradedGiftPurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TonTransactionTypeUpgradedGiftPurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TonTransactionTypeUpgradedGiftPurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TonTransactionTypeUpgradedGiftPurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetUser is a helper method for Client.GetUser
func (t TonTransactionTypeUpgradedGiftPurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TonTransactionTypeUpgradedGiftPurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TonTransactionTypeUpgradedGiftPurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TonTransactionTypeUpgradedGiftPurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TonTransactionTypeUpgradedGiftPurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TonTransactionTypeUpgradedGiftPurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TonTransactionTypeUpgradedGiftPurchase) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TonTransactionTypeUpgradedGiftPurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TonTransactionTypeUpgradedGiftPurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, t.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TonTransactionTypeUpgradedGiftPurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TonTransactionTypeUpgradedGiftPurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TonTransactionTypeUpgradedGiftPurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TonTransactionTypeUpgradedGiftPurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TonTransactionTypeUpgradedGiftPurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TonTransactionTypeUpgradedGiftPurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TonTransactionTypeUpgradedGiftPurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TonTransactionTypeUpgradedGiftPurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TonTransactionTypeUpgradedGiftPurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TonTransactionTypeUpgradedGiftPurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TonTransactionTypeUpgradedGiftPurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TonTransactionTypeUpgradedGiftPurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TonTransactionTypeUpgradedGiftPurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TonTransactionTypeUpgradedGiftPurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TonTransactionTypeUpgradedGiftPurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TonTransactionTypeUpgradedGiftPurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TonTransactionTypeUpgradedGiftPurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TonTransactionTypeUpgradedGiftPurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TonTransactionTypeUpgradedGiftSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TonTransactionTypeUpgradedGiftSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TonTransactionTypeUpgradedGiftSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TonTransactionTypeUpgradedGiftSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TonTransactionTypeUpgradedGiftSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (t TonTransactionTypeUpgradedGiftSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TonTransactionTypeUpgradedGiftSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TonTransactionTypeUpgradedGiftSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TonTransactionTypeUpgradedGiftSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TonTransactionTypeUpgradedGiftSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TonTransactionTypeUpgradedGiftSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TonTransactionTypeUpgradedGiftSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TonTransactionTypeUpgradedGiftSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TonTransactionTypeUpgradedGiftSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TonTransactionTypeUpgradedGiftSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TonTransactionTypeUpgradedGiftSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetUser is a helper method for Client.GetUser
func (t TonTransactionTypeUpgradedGiftSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TonTransactionTypeUpgradedGiftSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TonTransactionTypeUpgradedGiftSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TonTransactionTypeUpgradedGiftSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TonTransactionTypeUpgradedGiftSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TonTransactionTypeUpgradedGiftSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TonTransactionTypeUpgradedGiftSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TonTransactionTypeUpgradedGiftSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TonTransactionTypeUpgradedGiftSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, t.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TonTransactionTypeUpgradedGiftSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TonTransactionTypeUpgradedGiftSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TonTransactionTypeUpgradedGiftSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TonTransactionTypeUpgradedGiftSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TonTransactionTypeUpgradedGiftSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TonTransactionTypeUpgradedGiftSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TonTransactionTypeUpgradedGiftSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TonTransactionTypeUpgradedGiftSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TonTransactionTypeUpgradedGiftSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TonTransactionTypeUpgradedGiftSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TonTransactionTypeUpgradedGiftSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TonTransactionTypeUpgradedGiftSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TonTransactionTypeUpgradedGiftSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TonTransactionTypeUpgradedGiftSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TonTransactionTypeUpgradedGiftSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TonTransactionTypeUpgradedGiftSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TonTransactionTypeUpgradedGiftSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TonTransactionTypeUpgradedGiftSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// SetTdlibParameters is a helper method for Client.SetTdlibParameters
func (u UnconfirmedSession) SetTdlibParameters(client *Client, useTestDc bool, databaseDirectory string, filesDirectory string, databaseEncryptionKey string, useFileDatabase bool, useChatInfoDatabase bool, useMessageDatabase bool, useSecretChats bool, apiId int32, apiHash string, systemLanguageCode string, systemVersion string, applicationVersion string) (*Ok, error) {
	return client.SetTdlibParameters(useTestDc, databaseDirectory, filesDirectory, databaseEncryptionKey, useFileDatabase, useChatInfoDatabase, useMessageDatabase, useSecretChats, apiId, apiHash, systemLanguageCode, u.DeviceModel, systemVersion, applicationVersion)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UnreadReaction) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, u.SenderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UnreadReaction) AddMessageReaction(client *Client, chatId int64, messageId int64, reactionType *ReactionType, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, messageId, reactionType, u.IsBig, updateRecentReactions)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UnreadReaction) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, u.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UnreadReaction) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, u.SenderId, reportSpam)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UnreadReaction) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, u.SenderId)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UnreadReaction) SetMessageReactions(client *Client, chatId int64, messageId int64, reactionTypes []*ReactionType) (*Ok, error) {
	return client.SetMessageReactions(chatId, messageId, reactionTypes, u.IsBig)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (u UnreadReaction) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(u.SenderId, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateAnimatedEmojiMessageClicked) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateAnimatedEmojiMessageClicked) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateAnimatedEmojiMessageClicked) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateAnimatedEmojiMessageClicked) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateAnimatedEmojiMessageClicked) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateAnimatedEmojiMessageClicked) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateAnimatedEmojiMessageClicked) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateAnimatedEmojiMessageClicked) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateAnimatedEmojiMessageClicked) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateAnimatedEmojiMessageClicked) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateAnimatedEmojiMessageClicked) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateAnimatedEmojiMessageClicked) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateAnimatedEmojiMessageClicked) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateAnimatedEmojiMessageClicked) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateAnimatedEmojiMessageClicked) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateAnimatedEmojiMessageClicked) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateAnimatedEmojiMessageClicked) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateAnimatedEmojiMessageClicked) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateAnimatedEmojiMessageClicked) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateAnimatedEmojiMessageClicked) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateAnimatedEmojiMessageClicked) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateAnimatedEmojiMessageClicked) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateAnimatedEmojiMessageClicked) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateAnimatedEmojiMessageClicked) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateAnimatedEmojiMessageClicked) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateAnimatedEmojiMessageClicked) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateAnimatedEmojiMessageClicked) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateAnimatedEmojiMessageClicked) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateAnimatedEmojiMessageClicked) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateAnimatedEmojiMessageClicked) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateAnimatedEmojiMessageClicked) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateAnimatedEmojiMessageClicked) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateAnimatedEmojiMessageClicked) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateAnimatedEmojiMessageClicked) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateAnimatedEmojiMessageClicked) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateAnimatedEmojiMessageClicked) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateAnimatedEmojiMessageClicked) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateAnimatedEmojiMessageClicked) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateAnimatedEmojiMessageClicked) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateAnimatedEmojiMessageClicked) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateAnimatedEmojiMessageClicked) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateAnimatedEmojiMessageClicked) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateAnimatedEmojiMessageClicked) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateAnimatedEmojiMessageClicked) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateAnimatedEmojiMessageClicked) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateAnimatedEmojiMessageClicked) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateAnimatedEmojiMessageClicked) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateAnimatedEmojiMessageClicked) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateAnimatedEmojiMessageClicked) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateAnimatedEmojiMessageClicked) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateAnimatedEmojiMessageClicked) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateAnimatedEmojiMessageClicked) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateAnimatedEmojiMessageClicked) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateAnimatedEmojiMessageClicked) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateAnimatedEmojiMessageClicked) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateAnimatedEmojiMessageClicked) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateAnimatedEmojiMessageClicked) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateAnimatedEmojiMessageClicked) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateAnimatedEmojiMessageClicked) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateAnimatedEmojiMessageClicked) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateAnimatedEmojiMessageClicked) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateAnimatedEmojiMessageClicked) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateAnimatedEmojiMessageClicked) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateAnimatedEmojiMessageClicked) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateAnimatedEmojiMessageClicked) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateAnimatedEmojiMessageClicked) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateAnimatedEmojiMessageClicked) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateAnimatedEmojiMessageClicked) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateAnimatedEmojiMessageClicked) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateAnimatedEmojiMessageClicked) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateAnimatedEmojiMessageClicked) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateAnimatedEmojiMessageClicked) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateAnimatedEmojiMessageClicked) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateAnimatedEmojiMessageClicked) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateAnimatedEmojiMessageClicked) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateAnimatedEmojiMessageClicked) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateAnimatedEmojiMessageClicked) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateAnimatedEmojiMessageClicked) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateAnimatedEmojiMessageClicked) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateAnimatedEmojiMessageClicked) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateAnimatedEmojiMessageClicked) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateAnimatedEmojiMessageClicked) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateAnimatedEmojiMessageClicked) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateAnimatedEmojiMessageClicked) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateAnimatedEmojiMessageClicked) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateAnimatedEmojiMessageClicked) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateAnimatedEmojiMessageClicked) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateAnimatedEmojiMessageClicked) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateAnimatedEmojiMessageClicked) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateAnimatedEmojiMessageClicked) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateAnimatedEmojiMessageClicked) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateAnimatedEmojiMessageClicked) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateAnimatedEmojiMessageClicked) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateAnimatedEmojiMessageClicked) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateAnimatedEmojiMessageClicked) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateAnimatedEmojiMessageClicked) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateAnimatedEmojiMessageClicked) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateAnimatedEmojiMessageClicked) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateAnimatedEmojiMessageClicked) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateAnimatedEmojiMessageClicked) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateAnimatedEmojiMessageClicked) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateAnimatedEmojiMessageClicked) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateAnimatedEmojiMessageClicked) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateAnimatedEmojiMessageClicked) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateAnimatedEmojiMessageClicked) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateAnimatedEmojiMessageClicked) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateAnimatedEmojiMessageClicked) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateAnimatedEmojiMessageClicked) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateAnimatedEmojiMessageClicked) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateAnimatedEmojiMessageClicked) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateAnimatedEmojiMessageClicked) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateAnimatedEmojiMessageClicked) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateAnimatedEmojiMessageClicked) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateAnimatedEmojiMessageClicked) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateAnimatedEmojiMessageClicked) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateAnimatedEmojiMessageClicked) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateAnimatedEmojiMessageClicked) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateAnimatedEmojiMessageClicked) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateAnimatedEmojiMessageClicked) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateAnimatedEmojiMessageClicked) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateAnimatedEmojiMessageClicked) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateAnimatedEmojiMessageClicked) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateAnimatedEmojiMessageClicked) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateAnimatedEmojiMessageClicked) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateAnimatedEmojiMessageClicked) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateAnimatedEmojiMessageClicked) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateAnimatedEmojiMessageClicked) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateAnimatedEmojiMessageClicked) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateAnimatedEmojiMessageClicked) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateAnimatedEmojiMessageClicked) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateAnimatedEmojiMessageClicked) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateAnimatedEmojiMessageClicked) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateAnimatedEmojiMessageClicked) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateAnimatedEmojiMessageClicked) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateAnimatedEmojiMessageClicked) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateAnimatedEmojiMessageClicked) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateAnimatedEmojiMessageClicked) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateAnimatedEmojiMessageClicked) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateAnimatedEmojiMessageClicked) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateAnimatedEmojiMessageClicked) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateAnimatedEmojiMessageClicked) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateAnimatedEmojiMessageClicked) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateAnimatedEmojiMessageClicked) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateAnimatedEmojiMessageClicked) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateAnimatedEmojiMessageClicked) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateAnimatedEmojiMessageClicked) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateAnimatedEmojiMessageClicked) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateAnimatedEmojiMessageClicked) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateAnimatedEmojiMessageClicked) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateAnimatedEmojiMessageClicked) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateAnimatedEmojiMessageClicked) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateAnimatedEmojiMessageClicked) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateAnimatedEmojiMessageClicked) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateAnimatedEmojiMessageClicked) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateAnimatedEmojiMessageClicked) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateAnimatedEmojiMessageClicked) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateAnimatedEmojiMessageClicked) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateAnimatedEmojiMessageClicked) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateAnimatedEmojiMessageClicked) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateAnimatedEmojiMessageClicked) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateAnimatedEmojiMessageClicked) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateAnimatedEmojiMessageClicked) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateAnimatedEmojiMessageClicked) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateAnimatedEmojiMessageClicked) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateAnimatedEmojiMessageClicked) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateAnimatedEmojiMessageClicked) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateAnimatedEmojiMessageClicked) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateAnimatedEmojiMessageClicked) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateAnimatedEmojiMessageClicked) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateAnimatedEmojiMessageClicked) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateAnimatedEmojiMessageClicked) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateAnimatedEmojiMessageClicked) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateAnimatedEmojiMessageClicked) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateAnimatedEmojiMessageClicked) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateAnimatedEmojiMessageClicked) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateAnimatedEmojiMessageClicked) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateAnimatedEmojiMessageClicked) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateAnimatedEmojiMessageClicked) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateAnimatedEmojiMessageClicked) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateAnimatedEmojiMessageClicked) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateAnimatedEmojiMessageClicked) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateAnimatedEmojiMessageClicked) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateAnimatedEmojiMessageClicked) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateAnimatedEmojiMessageClicked) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateAnimatedEmojiMessageClicked) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateAnimatedEmojiMessageClicked) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateAnimatedEmojiMessageClicked) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateAnimatedEmojiMessageClicked) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateAnimatedEmojiMessageClicked) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateAnimatedEmojiMessageClicked) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateAnimatedEmojiMessageClicked) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateAnimatedEmojiMessageClicked) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateAnimatedEmojiMessageClicked) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateAnimatedEmojiMessageClicked) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateAnimatedEmojiMessageClicked) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateAnimatedEmojiMessageClicked) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateAnimatedEmojiMessageClicked) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateAnimatedEmojiMessageClicked) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateAnimatedEmojiMessageClicked) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateAnimatedEmojiMessageClicked) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateAnimatedEmojiMessageClicked) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateAnimatedEmojiMessageClicked) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateAnimatedEmojiMessageClicked) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateAnimatedEmojiMessageClicked) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateAnimatedEmojiMessageClicked) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateAnimatedEmojiMessageClicked) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateAnimatedEmojiMessageClicked) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateAnimatedEmojiMessageClicked) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateAnimatedEmojiMessageClicked) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateAnimatedEmojiMessageClicked) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateAnimatedEmojiMessageClicked) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateAnimatedEmojiMessageClicked) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateAnimatedEmojiMessageClicked) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateAnimatedEmojiMessageClicked) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateAnimatedEmojiMessageClicked) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateAnimatedEmojiMessageClicked) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateAnimatedEmojiMessageClicked) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateAnimatedEmojiMessageClicked) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateAnimatedEmojiMessageClicked) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateAnimatedEmojiMessageClicked) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateAnimatedEmojiMessageClicked) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateAnimatedEmojiMessageClicked) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateAnimatedEmojiMessageClicked) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateAnimatedEmojiMessageClicked) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateAnimatedEmojiMessageClicked) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, u.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateAnimatedEmojiMessageClicked) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, u.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateAnimatedEmojiMessageClicked) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateAnimatedEmojiMessageClicked) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateAnimatedEmojiMessageClicked) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateAnimatedEmojiMessageClicked) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateAnimatedEmojiMessageClicked) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateAnimatedEmojiMessageClicked) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateAnimatedEmojiMessageClicked) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateAnimatedEmojiMessageClicked) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, u.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateAnimatedEmojiMessageClicked) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateAnimatedEmojiMessageClicked) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateAnimatedEmojiMessageClicked) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateAnimatedEmojiMessageClicked) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, u.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateAnimatedEmojiMessageClicked) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateAnimatedEmojiMessageClicked) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (u UpdateApplicationRecaptchaVerificationRequired) SetApplicationVerificationToken(client *Client, token string) (*Ok, error) {
	return client.SetApplicationVerificationToken(u.VerificationId, token)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (u UpdateApplicationVerificationRequired) GetPassportAuthorizationForm(client *Client, botUserId int64, scope string, publicKey string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(botUserId, scope, publicKey, u.Nonce)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (u UpdateApplicationVerificationRequired) SetApplicationVerificationToken(client *Client, token string) (*Ok, error) {
	return client.SetApplicationVerificationToken(u.VerificationId, token)
}

// SetAutosaveSettings is a helper method for Client.SetAutosaveSettings
func (u UpdateAutosaveSettings) SetAutosaveSettings(client *Client) (*Ok, error) {
	return client.SetAutosaveSettings(u.Scope, u.Settings)
}

// CreateBasicGroupChat is a helper method for Client.CreateBasicGroupChat
func (u UpdateBasicGroupFullInfo) CreateBasicGroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateBasicGroupChat(u.BasicGroupId, force)
}

// GetBasicGroup is a helper method for Client.GetBasicGroup
func (u UpdateBasicGroupFullInfo) GetBasicGroup(client *Client) (*BasicGroup, error) {
	return client.GetBasicGroup(u.BasicGroupId)
}

// GetBasicGroupFullInfo is a helper method for Client.GetBasicGroupFullInfo
func (u UpdateBasicGroupFullInfo) GetBasicGroupFullInfo(client *Client) (*BasicGroupFullInfo, error) {
	return client.GetBasicGroupFullInfo(u.BasicGroupId)
}

// GetBusinessConnection is a helper method for Client.GetBusinessConnection
func (u UpdateBusinessMessageEdited) GetBusinessConnection(client *Client) (*BusinessConnection, error) {
	return client.GetBusinessConnection(u.ConnectionId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateBusinessMessagesDeleted) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateBusinessMessagesDeleted) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateBusinessMessagesDeleted) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateBusinessMessagesDeleted) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateBusinessMessagesDeleted) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateBusinessMessagesDeleted) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateBusinessMessagesDeleted) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateBusinessMessagesDeleted) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateBusinessMessagesDeleted) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateBusinessMessagesDeleted) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateBusinessMessagesDeleted) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateBusinessMessagesDeleted) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateBusinessMessagesDeleted) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateBusinessMessagesDeleted) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateBusinessMessagesDeleted) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateBusinessMessagesDeleted) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateBusinessMessagesDeleted) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateBusinessMessagesDeleted) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateBusinessMessagesDeleted) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateBusinessMessagesDeleted) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateBusinessMessagesDeleted) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateBusinessMessagesDeleted) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateBusinessMessagesDeleted) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateBusinessMessagesDeleted) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateBusinessMessagesDeleted) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateBusinessMessagesDeleted) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateBusinessMessagesDeleted) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteBusinessMessages is a helper method for Client.DeleteBusinessMessages
func (u UpdateBusinessMessagesDeleted) DeleteBusinessMessages(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessMessages(businessConnectionId, u.MessageIds)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateBusinessMessagesDeleted) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateBusinessMessagesDeleted) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateBusinessMessagesDeleted) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateBusinessMessagesDeleted) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateBusinessMessagesDeleted) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateBusinessMessagesDeleted) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateBusinessMessagesDeleted) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateBusinessMessagesDeleted) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateBusinessMessagesDeleted) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateBusinessMessagesDeleted) DeleteMessages(client *Client, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, u.MessageIds, revoke)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (u UpdateBusinessMessagesDeleted) DeleteQuickReplyShortcutMessages(client *Client, shortcutId int32) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(shortcutId, u.MessageIds)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateBusinessMessagesDeleted) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateBusinessMessagesDeleted) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateBusinessMessagesDeleted) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateBusinessMessagesDeleted) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateBusinessMessagesDeleted) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateBusinessMessagesDeleted) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateBusinessMessagesDeleted) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateBusinessMessagesDeleted) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateBusinessMessagesDeleted) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateBusinessMessagesDeleted) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateBusinessMessagesDeleted) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateBusinessMessagesDeleted) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateBusinessMessagesDeleted) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateBusinessMessagesDeleted) ForwardMessages(client *Client, fromChatId int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, u.MessageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateBusinessMessagesDeleted) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetBusinessConnection is a helper method for Client.GetBusinessConnection
func (u UpdateBusinessMessagesDeleted) GetBusinessConnection(client *Client) (*BusinessConnection, error) {
	return client.GetBusinessConnection(u.ConnectionId)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateBusinessMessagesDeleted) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateBusinessMessagesDeleted) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateBusinessMessagesDeleted) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateBusinessMessagesDeleted) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateBusinessMessagesDeleted) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateBusinessMessagesDeleted) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateBusinessMessagesDeleted) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateBusinessMessagesDeleted) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateBusinessMessagesDeleted) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateBusinessMessagesDeleted) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateBusinessMessagesDeleted) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateBusinessMessagesDeleted) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateBusinessMessagesDeleted) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateBusinessMessagesDeleted) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateBusinessMessagesDeleted) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateBusinessMessagesDeleted) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateBusinessMessagesDeleted) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateBusinessMessagesDeleted) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateBusinessMessagesDeleted) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateBusinessMessagesDeleted) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateBusinessMessagesDeleted) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateBusinessMessagesDeleted) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateBusinessMessagesDeleted) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateBusinessMessagesDeleted) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateBusinessMessagesDeleted) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateBusinessMessagesDeleted) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateBusinessMessagesDeleted) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateBusinessMessagesDeleted) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateBusinessMessagesDeleted) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateBusinessMessagesDeleted) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateBusinessMessagesDeleted) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateBusinessMessagesDeleted) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateBusinessMessagesDeleted) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateBusinessMessagesDeleted) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateBusinessMessagesDeleted) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateBusinessMessagesDeleted) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateBusinessMessagesDeleted) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateBusinessMessagesDeleted) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateBusinessMessagesDeleted) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateBusinessMessagesDeleted) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateBusinessMessagesDeleted) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateBusinessMessagesDeleted) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateBusinessMessagesDeleted) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateBusinessMessagesDeleted) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateBusinessMessagesDeleted) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateBusinessMessagesDeleted) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateBusinessMessagesDeleted) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateBusinessMessagesDeleted) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateBusinessMessagesDeleted) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateBusinessMessagesDeleted) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateBusinessMessagesDeleted) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateBusinessMessagesDeleted) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateBusinessMessagesDeleted) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateBusinessMessagesDeleted) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateBusinessMessagesDeleted) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateBusinessMessagesDeleted) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateBusinessMessagesDeleted) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateBusinessMessagesDeleted) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateBusinessMessagesDeleted) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateBusinessMessagesDeleted) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateBusinessMessagesDeleted) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateBusinessMessagesDeleted) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateBusinessMessagesDeleted) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateBusinessMessagesDeleted) GetMessages(client *Client) (*Messages, error) {
	return client.GetMessages(u.ChatId, u.MessageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateBusinessMessagesDeleted) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateBusinessMessagesDeleted) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateBusinessMessagesDeleted) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateBusinessMessagesDeleted) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateBusinessMessagesDeleted) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateBusinessMessagesDeleted) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateBusinessMessagesDeleted) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateBusinessMessagesDeleted) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateBusinessMessagesDeleted) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateBusinessMessagesDeleted) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateBusinessMessagesDeleted) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateBusinessMessagesDeleted) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateBusinessMessagesDeleted) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateBusinessMessagesDeleted) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateBusinessMessagesDeleted) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateBusinessMessagesDeleted) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateBusinessMessagesDeleted) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateBusinessMessagesDeleted) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateBusinessMessagesDeleted) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateBusinessMessagesDeleted) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateBusinessMessagesDeleted) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateBusinessMessagesDeleted) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateBusinessMessagesDeleted) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateBusinessMessagesDeleted) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateBusinessMessagesDeleted) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateBusinessMessagesDeleted) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateBusinessMessagesDeleted) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateBusinessMessagesDeleted) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateBusinessMessagesDeleted) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateBusinessMessagesDeleted) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateBusinessMessagesDeleted) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateBusinessMessagesDeleted) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateBusinessMessagesDeleted) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateBusinessMessagesDeleted) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateBusinessMessagesDeleted) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateBusinessMessagesDeleted) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// ReaddQuickReplyShortcutMessages is a helper method for Client.ReaddQuickReplyShortcutMessages
func (u UpdateBusinessMessagesDeleted) ReaddQuickReplyShortcutMessages(client *Client, shortcutName string) (*QuickReplyMessages, error) {
	return client.ReaddQuickReplyShortcutMessages(shortcutName, u.MessageIds)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateBusinessMessagesDeleted) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateBusinessMessagesDeleted) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateBusinessMessagesDeleted) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateBusinessMessagesDeleted) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateBusinessMessagesDeleted) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateBusinessMessagesDeleted) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateBusinessMessagesDeleted) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateBusinessMessagesDeleted) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateBusinessMessagesDeleted) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateBusinessMessagesDeleted) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateBusinessMessagesDeleted) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateBusinessMessagesDeleted) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateBusinessMessagesDeleted) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateBusinessMessagesDeleted) ReportChat(client *Client, optionId string, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, u.MessageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateBusinessMessagesDeleted) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateBusinessMessagesDeleted) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateBusinessMessagesDeleted) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (u UpdateBusinessMessagesDeleted) ReportSupergroupSpam(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupSpam(supergroupId, u.MessageIds)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateBusinessMessagesDeleted) ResendMessages(client *Client, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, u.MessageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateBusinessMessagesDeleted) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateBusinessMessagesDeleted) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateBusinessMessagesDeleted) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateBusinessMessagesDeleted) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateBusinessMessagesDeleted) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateBusinessMessagesDeleted) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateBusinessMessagesDeleted) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateBusinessMessagesDeleted) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateBusinessMessagesDeleted) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateBusinessMessagesDeleted) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateBusinessMessagesDeleted) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateBusinessMessagesDeleted) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateBusinessMessagesDeleted) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateBusinessMessagesDeleted) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateBusinessMessagesDeleted) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateBusinessMessagesDeleted) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateBusinessMessagesDeleted) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateBusinessMessagesDeleted) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateBusinessMessagesDeleted) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateBusinessMessagesDeleted) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateBusinessMessagesDeleted) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateBusinessMessagesDeleted) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateBusinessMessagesDeleted) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateBusinessMessagesDeleted) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateBusinessMessagesDeleted) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateBusinessMessagesDeleted) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateBusinessMessagesDeleted) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateBusinessMessagesDeleted) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateBusinessMessagesDeleted) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateBusinessMessagesDeleted) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateBusinessMessagesDeleted) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateBusinessMessagesDeleted) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateBusinessMessagesDeleted) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateBusinessMessagesDeleted) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateBusinessMessagesDeleted) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateBusinessMessagesDeleted) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateBusinessMessagesDeleted) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateBusinessMessagesDeleted) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateBusinessMessagesDeleted) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateBusinessMessagesDeleted) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateBusinessMessagesDeleted) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateBusinessMessagesDeleted) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateBusinessMessagesDeleted) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateBusinessMessagesDeleted) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateBusinessMessagesDeleted) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateBusinessMessagesDeleted) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateBusinessMessagesDeleted) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateBusinessMessagesDeleted) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateBusinessMessagesDeleted) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateBusinessMessagesDeleted) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateBusinessMessagesDeleted) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateBusinessMessagesDeleted) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateBusinessMessagesDeleted) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateBusinessMessagesDeleted) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateBusinessMessagesDeleted) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateBusinessMessagesDeleted) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateBusinessMessagesDeleted) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateBusinessMessagesDeleted) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateBusinessMessagesDeleted) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateBusinessMessagesDeleted) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateBusinessMessagesDeleted) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateBusinessMessagesDeleted) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateBusinessMessagesDeleted) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateBusinessMessagesDeleted) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateBusinessMessagesDeleted) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateBusinessMessagesDeleted) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateBusinessMessagesDeleted) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateBusinessMessagesDeleted) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateBusinessMessagesDeleted) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateBusinessMessagesDeleted) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateBusinessMessagesDeleted) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateBusinessMessagesDeleted) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateBusinessMessagesDeleted) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateBusinessMessagesDeleted) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateBusinessMessagesDeleted) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateBusinessMessagesDeleted) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateBusinessMessagesDeleted) ViewMessages(client *Client, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, u.MessageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatAccentColors) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatAccentColors) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatAccentColors) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatAccentColors) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatAccentColors) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatAccentColors) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatAccentColors) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatAccentColors) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatAccentColors) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatAccentColors) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatAccentColors) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatAccentColors) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatAccentColors) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatAccentColors) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatAccentColors) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatAccentColors) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatAccentColors) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatAccentColors) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatAccentColors) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatAccentColors) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatAccentColors) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatAccentColors) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatAccentColors) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatAccentColors) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatAccentColors) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatAccentColors) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatAccentColors) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatAccentColors) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatAccentColors) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatAccentColors) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatAccentColors) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatAccentColors) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatAccentColors) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatAccentColors) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatAccentColors) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatAccentColors) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatAccentColors) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatAccentColors) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatAccentColors) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatAccentColors) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatAccentColors) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatAccentColors) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatAccentColors) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatAccentColors) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatAccentColors) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatAccentColors) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatAccentColors) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatAccentColors) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatAccentColors) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatAccentColors) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatAccentColors) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatAccentColors) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatAccentColors) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatAccentColors) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatAccentColors) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatAccentColors) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatAccentColors) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatAccentColors) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatAccentColors) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatAccentColors) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatAccentColors) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatAccentColors) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatAccentColors) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatAccentColors) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatAccentColors) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatAccentColors) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatAccentColors) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatAccentColors) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatAccentColors) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatAccentColors) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatAccentColors) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatAccentColors) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatAccentColors) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatAccentColors) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatAccentColors) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatAccentColors) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatAccentColors) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatAccentColors) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatAccentColors) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatAccentColors) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatAccentColors) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatAccentColors) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatAccentColors) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatAccentColors) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatAccentColors) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatAccentColors) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatAccentColors) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatAccentColors) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatAccentColors) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatAccentColors) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatAccentColors) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatAccentColors) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatAccentColors) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatAccentColors) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatAccentColors) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatAccentColors) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatAccentColors) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatAccentColors) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatAccentColors) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatAccentColors) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatAccentColors) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatAccentColors) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatAccentColors) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatAccentColors) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatAccentColors) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatAccentColors) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatAccentColors) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatAccentColors) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatAccentColors) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatAccentColors) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatAccentColors) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatAccentColors) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatAccentColors) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatAccentColors) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatAccentColors) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatAccentColors) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatAccentColors) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatAccentColors) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatAccentColors) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatAccentColors) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatAccentColors) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatAccentColors) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatAccentColors) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatAccentColors) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatAccentColors) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatAccentColors) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatAccentColors) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatAccentColors) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatAccentColors) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatAccentColors) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatAccentColors) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatAccentColors) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatAccentColors) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatAccentColors) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatAccentColors) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatAccentColors) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatAccentColors) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatAccentColors) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatAccentColors) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatAccentColors) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatAccentColors) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatAccentColors) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatAccentColors) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatAccentColors) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatAccentColors) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatAccentColors) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatAccentColors) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatAccentColors) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatAccentColors) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatAccentColors) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatAccentColors) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatAccentColors) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatAccentColors) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatAccentColors) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatAccentColors) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatAccentColors) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatAccentColors) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatAccentColors) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatAccentColors) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatAccentColors) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatAccentColors) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatAccentColors) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatAccentColors) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatAccentColors) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatAccentColors) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatAccentColors) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatAccentColors) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatAccentColors) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatAccentColors) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatAccentColors) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatAccentColors) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatAccentColors) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatAccentColors) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatAccentColors) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatAccentColors) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatAccentColors) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatAccentColors) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatAccentColors) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatAccentColors) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatAccentColors) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatAccentColors) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatAccentColors) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatAccentColors) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatAccentColors) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatAccentColors) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatAccentColors) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatAccentColors) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatAccentColors) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatAccentColors) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatAccentColors) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetAccentColor is a helper method for Client.SetAccentColor
func (u UpdateChatAccentColors) SetAccentColor(client *Client) (*Ok, error) {
	return client.SetAccentColor(u.AccentColorId, u.BackgroundCustomEmojiId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatAccentColors) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatAccentColors) SetChatAccentColor(client *Client) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, u.AccentColorId, u.BackgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatAccentColors) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatAccentColors) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatAccentColors) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatAccentColors) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatAccentColors) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatAccentColors) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatAccentColors) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatAccentColors) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatAccentColors) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatAccentColors) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatAccentColors) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatAccentColors) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatAccentColors) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatAccentColors) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatAccentColors) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatAccentColors) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatAccentColors) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatAccentColors) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatAccentColors) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatAccentColors) SetChatProfileAccentColor(client *Client) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, u.ProfileAccentColorId, u.ProfileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatAccentColors) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatAccentColors) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatAccentColors) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatAccentColors) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatAccentColors) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatAccentColors) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatAccentColors) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatAccentColors) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatAccentColors) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatAccentColors) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatAccentColors) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatAccentColors) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetProfileAccentColor is a helper method for Client.SetProfileAccentColor
func (u UpdateChatAccentColors) SetProfileAccentColor(client *Client) (*Ok, error) {
	return client.SetProfileAccentColor(u.ProfileAccentColorId, u.ProfileBackgroundCustomEmojiId)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatAccentColors) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatAccentColors) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatAccentColors) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatAccentColors) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatAccentColors) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatAccentColors) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatAccentColors) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatAccentColors) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatAccentColors) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatAccentColors) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatAccentColors) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatAccentColors) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatAccentColors) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatAccentColors) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatAccentColors) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatAccentColors) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatAccentColors) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatAccentColors) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatAccentColors) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatAccentColors) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatAccentColors) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatAccentColors) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatAccentColors) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatAccentColors) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatAccentColors) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatAccentColors) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatAccentColors) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatAccentColors) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatAction) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatAction) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatAction) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatAction) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatAction) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatAction) AddLocalMessage(client *Client, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, u.SenderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatAction) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatAction) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatAction) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatAction) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatAction) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatAction) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatAction) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatAction) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatAction) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatAction) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatAction) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatAction) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatAction) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatAction) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatAction) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatAction) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatAction) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatAction) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatAction) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatAction) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatAction) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatAction) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatAction) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatAction) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatAction) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatAction) DeleteChatMessagesBySender(client *Client) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, u.SenderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatAction) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatAction) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatAction) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatAction) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateChatAction) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, u.SenderId, reportSpam)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatAction) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatAction) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatAction) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatAction) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatAction) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatAction) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatAction) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatAction) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatAction) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatAction) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatAction) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatAction) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatAction) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatAction) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatAction) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatAction) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatAction) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatAction) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatAction) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatAction) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatAction) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatAction) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatAction) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatAction) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatAction) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatAction) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatAction) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatAction) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatAction) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatAction) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatAction) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatAction) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatAction) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatAction) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatAction) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatAction) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatAction) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatAction) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatAction) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatAction) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatAction) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatAction) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatAction) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatAction) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatAction) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatAction) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatAction) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatAction) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatAction) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatAction) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatAction) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatAction) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatAction) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatAction) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatAction) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatAction) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatAction) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatAction) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatAction) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatAction) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatAction) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatAction) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatAction) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatAction) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatAction) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatAction) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatAction) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatAction) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatAction) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatAction) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatAction) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatAction) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatAction) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatAction) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatAction) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatAction) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatAction) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatAction) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatAction) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatAction) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatAction) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatAction) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatAction) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatAction) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatAction) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatAction) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatAction) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatAction) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatAction) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatAction) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatAction) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatAction) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatAction) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatAction) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatAction) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatAction) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatAction) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatAction) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatAction) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatAction) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatAction) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatAction) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatAction) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatAction) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatAction) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatAction) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatAction) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatAction) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatAction) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatAction) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatAction) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatAction) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatAction) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatAction) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatAction) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatAction) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatAction) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatAction) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatAction) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatAction) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatAction) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatAction) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatAction) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatAction) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatAction) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatAction) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatAction) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatAction) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatAction) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatAction) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatAction) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatAction) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatAction) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatAction) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatAction) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatAction) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatAction) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatAction) ReportMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, u.SenderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatAction) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatAction) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatAction) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatAction) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatAction) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatAction) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatAction) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatAction) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatAction) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatAction) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatAction) SendChatAction(client *Client, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, u.TopicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatAction) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatAction) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatAction) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatAction) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatAction) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatAction) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatAction) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatAction) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatAction) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatAction) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatAction) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatAction) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatAction) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatAction) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatAction) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatAction) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatAction) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatAction) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatAction) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatAction) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatAction) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatAction) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatAction) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatAction) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatAction) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatAction) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatAction) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatAction) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatAction) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatAction) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatAction) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatAction) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatAction) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatAction) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatAction) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (u UpdateChatAction) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(u.SenderId, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatAction) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatAction) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatAction) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatAction) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatAction) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatAction) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatAction) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatAction) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatAction) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatAction) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatAction) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatAction) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatAction) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatAction) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatAction) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatAction) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatAction) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatAction) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatAction) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatAction) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatAction) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatAction) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatAction) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatAction) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatAction) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatAction) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatAction) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatAction) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatAction) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatAction) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatAction) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatAction) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatActionBar) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatActionBar) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatActionBar) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatActionBar) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatActionBar) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatActionBar) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatActionBar) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatActionBar) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatActionBar) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatActionBar) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatActionBar) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatActionBar) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatActionBar) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatActionBar) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatActionBar) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatActionBar) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatActionBar) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatActionBar) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatActionBar) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatActionBar) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatActionBar) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatActionBar) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatActionBar) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatActionBar) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatActionBar) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatActionBar) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatActionBar) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatActionBar) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatActionBar) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatActionBar) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatActionBar) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatActionBar) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatActionBar) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatActionBar) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatActionBar) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatActionBar) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatActionBar) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatActionBar) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatActionBar) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatActionBar) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatActionBar) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatActionBar) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatActionBar) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatActionBar) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatActionBar) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatActionBar) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatActionBar) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatActionBar) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatActionBar) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatActionBar) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatActionBar) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatActionBar) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatActionBar) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatActionBar) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatActionBar) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatActionBar) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatActionBar) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatActionBar) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatActionBar) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatActionBar) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatActionBar) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatActionBar) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatActionBar) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatActionBar) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatActionBar) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatActionBar) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatActionBar) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatActionBar) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatActionBar) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatActionBar) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatActionBar) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatActionBar) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatActionBar) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatActionBar) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatActionBar) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatActionBar) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatActionBar) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatActionBar) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatActionBar) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatActionBar) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatActionBar) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatActionBar) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatActionBar) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatActionBar) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatActionBar) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatActionBar) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatActionBar) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatActionBar) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatActionBar) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatActionBar) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatActionBar) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatActionBar) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatActionBar) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatActionBar) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatActionBar) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatActionBar) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatActionBar) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatActionBar) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatActionBar) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatActionBar) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatActionBar) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatActionBar) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatActionBar) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatActionBar) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatActionBar) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatActionBar) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatActionBar) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatActionBar) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatActionBar) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatActionBar) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatActionBar) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatActionBar) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatActionBar) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatActionBar) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatActionBar) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatActionBar) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatActionBar) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatActionBar) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatActionBar) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatActionBar) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatActionBar) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatActionBar) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatActionBar) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatActionBar) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatActionBar) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatActionBar) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatActionBar) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatActionBar) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatActionBar) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatActionBar) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatActionBar) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatActionBar) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatActionBar) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatActionBar) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatActionBar) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatActionBar) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatActionBar) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatActionBar) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatActionBar) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatActionBar) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatActionBar) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatActionBar) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatActionBar) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatActionBar) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatActionBar) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatActionBar) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatActionBar) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatActionBar) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatActionBar) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatActionBar) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatActionBar) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatActionBar) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatActionBar) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatActionBar) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatActionBar) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatActionBar) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatActionBar) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatActionBar) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatActionBar) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatActionBar) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatActionBar) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatActionBar) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatActionBar) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatActionBar) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatActionBar) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatActionBar) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatActionBar) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatActionBar) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatActionBar) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatActionBar) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatActionBar) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatActionBar) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}
