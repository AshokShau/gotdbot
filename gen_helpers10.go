package gotdbot

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarSubscription) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarSubscription) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarSubscription) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarSubscription) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarSubscription) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarSubscription) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarSubscription) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarSubscription) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarSubscription) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarSubscription) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarSubscription) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarSubscription) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarSubscription) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarSubscription) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarSubscription) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarSubscription) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarSubscription) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarSubscription) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarSubscription) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarSubscription) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarSubscription) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarSubscription) EditChatInviteLink(client *Client, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, s.ExpirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarSubscription) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarSubscription) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarSubscription) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarSubscription) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarSubscription) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarSubscription) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarSubscription) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarSubscription) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarSubscription) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// Edit is a helper method for Client.EditStarSubscription
func (s StarSubscription) Edit(client *Client, subscriptionId string) (*Ok, error) {
	return client.EditStarSubscription(subscriptionId, s.IsCanceled)
}

// EditUser is a helper method for Client.EditUserStarSubscription
func (s StarSubscription) EditUser(client *Client, userId int64, telegramPaymentChargeId string) (*Ok, error) {
	return client.EditUserStarSubscription(userId, telegramPaymentChargeId, s.IsCanceled)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarSubscription) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarSubscription) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarSubscription) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarSubscription) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarSubscription) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarSubscription) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarSubscription) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarSubscription) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarSubscription) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarSubscription) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarSubscription) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarSubscription) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarSubscription) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarSubscription) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarSubscription) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarSubscription) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarSubscription) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarSubscription) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarSubscription) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarSubscription) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarSubscription) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarSubscription) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarSubscription) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarSubscription) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarSubscription) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarSubscription) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarSubscription) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarSubscription) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarSubscription) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarSubscription) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarSubscription) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarSubscription) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarSubscription) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarSubscription) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarSubscription) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarSubscription) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarSubscription) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarSubscription) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarSubscription) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarSubscription) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarSubscription) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarSubscription) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarSubscription) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarSubscription) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarSubscription) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarSubscription) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarSubscription) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarSubscription) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarSubscription) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarSubscription) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarSubscription) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarSubscription) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarSubscription) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarSubscription) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarSubscription) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarSubscription) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarSubscription) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarSubscription) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarSubscription) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarSubscription) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarSubscription) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarSubscription) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarSubscription) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarSubscription) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarSubscription) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarSubscription) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarSubscription) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarSubscription) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarSubscription) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarSubscription) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarSubscription) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarSubscription) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarSubscription) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarSubscription) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarSubscription) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarSubscription) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarSubscription) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarSubscription) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarSubscription) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarSubscription) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarSubscription) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarSubscription) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarSubscription) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarSubscription) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarSubscription) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarSubscription) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarSubscription) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarSubscription) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarSubscription) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarSubscription) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarSubscription) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarSubscription) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarSubscription) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarSubscription) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarSubscription) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarSubscription) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarSubscription) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarSubscription) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarSubscription) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarSubscription) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarSubscription) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarSubscription) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarSubscription) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarSubscription) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarSubscription) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarSubscription) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarSubscription) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarSubscription) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarSubscription) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarSubscription) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarSubscription) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarSubscription) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarSubscription) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarSubscription) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarSubscription) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarSubscription) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarSubscription) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarSubscription) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarSubscription) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarSubscription) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarSubscription) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarSubscription) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarSubscription) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarSubscription) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarSubscription) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarSubscription) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarSubscription) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarSubscription) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarSubscription) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarSubscription) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarSubscription) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarSubscription) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarSubscription) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarSubscription) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarSubscription) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarSubscription) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarSubscription) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarSubscription) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarSubscription) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarSubscription) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarSubscription) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarSubscription) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarSubscription) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarSubscription) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarSubscription) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarSubscription) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarSubscription) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarSubscription) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarSubscription) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarSubscription) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarSubscription) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarSubscription) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarSubscription) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarSubscription) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarSubscription) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarSubscription) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarSubscription) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarSubscription) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarSubscription) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarSubscription) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarSubscription) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarSubscription) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarSubscription) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarSubscription) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarSubscription) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarSubscription) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarSubscription) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarSubscription) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarSubscription) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarSubscription) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarSubscription) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarSubscription) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarSubscription) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarSubscription) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarSubscription) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarSubscription) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarSubscription) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarSubscription) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarSubscription) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarSubscription) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarSubscription) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarSubscription) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarSubscription) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarSubscription) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarSubscription) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarSubscription) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarSubscription) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarSubscription) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarSubscription) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarSubscription) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarSubscription) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarSubscription) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarSubscription) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarSubscription) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarSubscription) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarSubscription) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarSubscription) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s StarSubscriptionPricing) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, s.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarSubscriptionPricing) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, s.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarSubscriptionPricing) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, s.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (s StarSubscriptionPricing) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, s.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarSubscriptionPricing) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, s.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarSubscriptionPricing) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, s.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (s StarSubscriptionPricing) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, s.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s StarSubscriptionPricing) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, s.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarSubscriptionPricing) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, s.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s StarSubscriptionPricing) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, s.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (s StarSubscriptionPricing) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, s.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (s StarSubscriptionPricing) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, s.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (s StarSubscriptionPricing) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, s.StarCount)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (s StarSubscriptions) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, s.NextOffset, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (s StarSubscriptionTypeBot) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, s.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarSubscriptionTypeBot) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, s.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (s StarSubscriptionTypeBot) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(s.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarSubscriptionTypeBot) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, s.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (s StarSubscriptionTypeBot) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(s.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarSubscriptionTypeBot) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, s.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (s StarSubscriptionTypeBot) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, s.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (s StarSubscriptionTypeBot) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, s.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (s StarSubscriptionTypeBot) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, s.Title, recordVideo, usePortraitOrientation)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (s StarSubscriptionTypeChannel) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(s.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (s StarSubscriptionTypeChannel) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(s.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (s StarSubscriptionTypeChannel) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(s.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (s StarSubscriptionTypeChannel) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, s.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarSubscriptionTypeChannel) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, s.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (s StarSubscriptionTypeChannel) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, s.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (s StarSubscriptionTypeChannel) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, s.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarSubscriptionTypeChannel) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, s.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarSubscriptionTypeChannel) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, s.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarSubscriptionTypeChannel) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, s.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarSubscriptionTypeChannel) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, s.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarSubscriptionTypeChannel) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, s.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (s StarSubscriptionTypeChannel) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(s.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarSubscriptionTypeChannel) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, s.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarSubscriptionTypeChannel) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, s.InviteLink)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransaction) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, s.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransaction) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, s.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (s StarTransaction) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, s.Date)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (s StarTransactions) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, s.NextOffset, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeAffiliateProgramCommission) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeAffiliateProgramCommission) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeAffiliateProgramCommission) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeAffiliateProgramCommission) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeAffiliateProgramCommission) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeAffiliateProgramCommission) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeAffiliateProgramCommission) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeAffiliateProgramCommission) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeAffiliateProgramCommission) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeAffiliateProgramCommission) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeAffiliateProgramCommission) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeAffiliateProgramCommission) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeAffiliateProgramCommission) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeAffiliateProgramCommission) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeAffiliateProgramCommission) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeAffiliateProgramCommission) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeAffiliateProgramCommission) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeAffiliateProgramCommission) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeAffiliateProgramCommission) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeAffiliateProgramCommission) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeAffiliateProgramCommission) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeAffiliateProgramCommission) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeAffiliateProgramCommission) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeAffiliateProgramCommission) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeAffiliateProgramCommission) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeAffiliateProgramCommission) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeAffiliateProgramCommission) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeAffiliateProgramCommission) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeAffiliateProgramCommission) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeAffiliateProgramCommission) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeAffiliateProgramCommission) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeAffiliateProgramCommission) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeAffiliateProgramCommission) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeAffiliateProgramCommission) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeAffiliateProgramCommission) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeAffiliateProgramCommission) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeAffiliateProgramCommission) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeAffiliateProgramCommission) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeAffiliateProgramCommission) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeAffiliateProgramCommission) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeAffiliateProgramCommission) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeAffiliateProgramCommission) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeAffiliateProgramCommission) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeAffiliateProgramCommission) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeAffiliateProgramCommission) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeAffiliateProgramCommission) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeAffiliateProgramCommission) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeAffiliateProgramCommission) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeAffiliateProgramCommission) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeAffiliateProgramCommission) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeAffiliateProgramCommission) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeAffiliateProgramCommission) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeAffiliateProgramCommission) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeAffiliateProgramCommission) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeAffiliateProgramCommission) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeAffiliateProgramCommission) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeAffiliateProgramCommission) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeAffiliateProgramCommission) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeAffiliateProgramCommission) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeAffiliateProgramCommission) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeAffiliateProgramCommission) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeAffiliateProgramCommission) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeAffiliateProgramCommission) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeAffiliateProgramCommission) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeAffiliateProgramCommission) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeAffiliateProgramCommission) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeAffiliateProgramCommission) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeAffiliateProgramCommission) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeAffiliateProgramCommission) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeAffiliateProgramCommission) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeAffiliateProgramCommission) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeAffiliateProgramCommission) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeAffiliateProgramCommission) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeAffiliateProgramCommission) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeAffiliateProgramCommission) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeAffiliateProgramCommission) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeAffiliateProgramCommission) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeAffiliateProgramCommission) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeAffiliateProgramCommission) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeAffiliateProgramCommission) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeAffiliateProgramCommission) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeAffiliateProgramCommission) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeAffiliateProgramCommission) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeAffiliateProgramCommission) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeAffiliateProgramCommission) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeAffiliateProgramCommission) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeAffiliateProgramCommission) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeAffiliateProgramCommission) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeAffiliateProgramCommission) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeAffiliateProgramCommission) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeAffiliateProgramCommission) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeAffiliateProgramCommission) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeAffiliateProgramCommission) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeAffiliateProgramCommission) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeAffiliateProgramCommission) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeAffiliateProgramCommission) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeAffiliateProgramCommission) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeAffiliateProgramCommission) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeAffiliateProgramCommission) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeAffiliateProgramCommission) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeAffiliateProgramCommission) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeAffiliateProgramCommission) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeAffiliateProgramCommission) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeAffiliateProgramCommission) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeAffiliateProgramCommission) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeAffiliateProgramCommission) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeAffiliateProgramCommission) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeAffiliateProgramCommission) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeAffiliateProgramCommission) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeAffiliateProgramCommission) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeAffiliateProgramCommission) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeAffiliateProgramCommission) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeAffiliateProgramCommission) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeAffiliateProgramCommission) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeAffiliateProgramCommission) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeAffiliateProgramCommission) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeAffiliateProgramCommission) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeAffiliateProgramCommission) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeAffiliateProgramCommission) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeAffiliateProgramCommission) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeAffiliateProgramCommission) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeAffiliateProgramCommission) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeAffiliateProgramCommission) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeAffiliateProgramCommission) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeAffiliateProgramCommission) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeAffiliateProgramCommission) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeAffiliateProgramCommission) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeAffiliateProgramCommission) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeAffiliateProgramCommission) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeAffiliateProgramCommission) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeAffiliateProgramCommission) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeAffiliateProgramCommission) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeAffiliateProgramCommission) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeAffiliateProgramCommission) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeAffiliateProgramCommission) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeAffiliateProgramCommission) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeAffiliateProgramCommission) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeAffiliateProgramCommission) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeAffiliateProgramCommission) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeAffiliateProgramCommission) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeAffiliateProgramCommission) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeAffiliateProgramCommission) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeAffiliateProgramCommission) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeAffiliateProgramCommission) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeAffiliateProgramCommission) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeAffiliateProgramCommission) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeAffiliateProgramCommission) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeAffiliateProgramCommission) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeAffiliateProgramCommission) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeAffiliateProgramCommission) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeAffiliateProgramCommission) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeAffiliateProgramCommission) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeAffiliateProgramCommission) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeAffiliateProgramCommission) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeAffiliateProgramCommission) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeAffiliateProgramCommission) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeAffiliateProgramCommission) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeAffiliateProgramCommission) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeAffiliateProgramCommission) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeAffiliateProgramCommission) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeAffiliateProgramCommission) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeAffiliateProgramCommission) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeAffiliateProgramCommission) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeAffiliateProgramCommission) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeAffiliateProgramCommission) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeAffiliateProgramCommission) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeAffiliateProgramCommission) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeAffiliateProgramCommission) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeAffiliateProgramCommission) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeAffiliateProgramCommission) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeAffiliateProgramCommission) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeAffiliateProgramCommission) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeAffiliateProgramCommission) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeAffiliateProgramCommission) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeAffiliateProgramCommission) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeAffiliateProgramCommission) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeAffiliateProgramCommission) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeAffiliateProgramCommission) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeAffiliateProgramCommission) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeAffiliateProgramCommission) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeAffiliateProgramCommission) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeAffiliateProgramCommission) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeAffiliateProgramCommission) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeAffiliateProgramCommission) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeAffiliateProgramCommission) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeAffiliateProgramCommission) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeAffiliateProgramCommission) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeAffiliateProgramCommission) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeAffiliateProgramCommission) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeAffiliateProgramCommission) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeAffiliateProgramCommission) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeAffiliateProgramCommission) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeAffiliateProgramCommission) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeAffiliateProgramCommission) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeAffiliateProgramCommission) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeAffiliateProgramCommission) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeAffiliateProgramCommission) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeAffiliateProgramCommission) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeAffiliateProgramCommission) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeAffiliateProgramCommission) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeAffiliateProgramCommission) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeAffiliateProgramCommission) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeAffiliateProgramCommission) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeAffiliateProgramCommission) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeAffiliateProgramCommission) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeAffiliateProgramCommission) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeAffiliateProgramCommission) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeAffiliateProgramCommission) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeAffiliateProgramCommission) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeAffiliateProgramCommission) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeAffiliateProgramCommission) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotInvoicePurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotInvoicePurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotInvoicePurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotInvoicePurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotInvoicePurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotInvoicePurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotInvoicePurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotInvoicePurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotInvoicePurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotInvoicePurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotInvoicePurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotInvoicePurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotInvoicePurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotInvoicePurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotInvoicePurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotInvoicePurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotInvoicePurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotInvoicePurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotInvoicePurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotInvoicePurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotInvoicePurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotInvoicePurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotInvoicePurchase) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotInvoicePurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotInvoicePurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotInvoicePurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotInvoicePurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotInvoicePurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotInvoicePurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotInvoicePurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotInvoicePurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotInvoicePurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotInvoicePurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotInvoicePurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotInvoicePurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotInvoicePurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotInvoicePurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotInvoicePurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotInvoicePurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotInvoicePurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotInvoicePurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotInvoicePurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotInvoicePurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotInvoiceSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotInvoiceSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotInvoiceSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotInvoiceSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotInvoiceSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotInvoiceSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotInvoiceSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotInvoiceSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotInvoiceSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotInvoiceSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotInvoiceSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotInvoiceSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotInvoiceSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotInvoiceSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotInvoiceSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotInvoiceSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotInvoiceSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotInvoiceSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotInvoiceSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotInvoiceSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotInvoiceSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotInvoiceSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotInvoiceSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotInvoiceSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotInvoiceSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotInvoiceSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotInvoiceSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotInvoiceSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotInvoiceSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotInvoiceSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotInvoiceSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotInvoiceSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotInvoiceSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotInvoiceSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotInvoiceSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotInvoiceSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotInvoiceSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotInvoiceSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotInvoiceSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotInvoiceSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotInvoiceSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotInvoiceSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotInvoiceSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotPaidMediaPurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotPaidMediaPurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotPaidMediaPurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotPaidMediaPurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotPaidMediaPurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotPaidMediaPurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotPaidMediaPurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotPaidMediaPurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotPaidMediaPurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotPaidMediaPurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotPaidMediaPurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotPaidMediaPurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotPaidMediaPurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotPaidMediaPurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotPaidMediaPurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotPaidMediaPurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotPaidMediaPurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotPaidMediaPurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotPaidMediaPurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotPaidMediaPurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotPaidMediaPurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotPaidMediaPurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotPaidMediaPurchase) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotPaidMediaPurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotPaidMediaPurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotPaidMediaPurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotPaidMediaPurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotPaidMediaPurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotPaidMediaPurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotPaidMediaPurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotPaidMediaPurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotPaidMediaPurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotPaidMediaPurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotPaidMediaPurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotPaidMediaPurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotPaidMediaPurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotPaidMediaPurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotPaidMediaPurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotPaidMediaPurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotPaidMediaPurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotPaidMediaPurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotPaidMediaPurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotPaidMediaPurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotPaidMediaSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotPaidMediaSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotPaidMediaSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotPaidMediaSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotPaidMediaSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotPaidMediaSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotPaidMediaSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotPaidMediaSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotPaidMediaSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotPaidMediaSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotPaidMediaSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotPaidMediaSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotPaidMediaSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotPaidMediaSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotPaidMediaSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetPushReceiverId is a helper method for Client.GetPushReceiverId
func (s StarTransactionTypeBotPaidMediaSale) GetPushReceiverId(client *Client) (*PushReceiverId, error) {
	return client.GetPushReceiverId(s.Payload)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotPaidMediaSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotPaidMediaSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotPaidMediaSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotPaidMediaSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotPaidMediaSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotPaidMediaSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotPaidMediaSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotPaidMediaSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotPaidMediaSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotPaidMediaSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotPaidMediaSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// ProcessPushNotification is a helper method for Client.ProcessPushNotification
func (s StarTransactionTypeBotPaidMediaSale) ProcessPushNotification(client *Client) (*Ok, error) {
	return client.ProcessPushNotification(s.Payload)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotPaidMediaSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotPaidMediaSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotPaidMediaSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotPaidMediaSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotPaidMediaSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotPaidMediaSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotPaidMediaSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotPaidMediaSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotPaidMediaSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotPaidMediaSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotPaidMediaSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotPaidMediaSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotPaidMediaSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (s StarTransactionTypeBotPaidMediaSale) StartGroupCallScreenSharing(client *Client, groupCallId int32, audioSourceId int32) (*Text, error) {
	return client.StartGroupCallScreenSharing(groupCallId, audioSourceId, s.Payload)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotPaidMediaSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotPaidMediaSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotPaidMediaSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotPaidMediaSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotSubscriptionPurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotSubscriptionPurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotSubscriptionPurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotSubscriptionPurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotSubscriptionPurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotSubscriptionPurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotSubscriptionPurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotSubscriptionPurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotSubscriptionPurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotSubscriptionPurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotSubscriptionPurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotSubscriptionPurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotSubscriptionPurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotSubscriptionPurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotSubscriptionPurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotSubscriptionPurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotSubscriptionPurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotSubscriptionPurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotSubscriptionPurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotSubscriptionPurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotSubscriptionPurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotSubscriptionPurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotSubscriptionPurchase) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotSubscriptionPurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotSubscriptionPurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotSubscriptionPurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotSubscriptionPurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotSubscriptionPurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotSubscriptionPurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotSubscriptionPurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotSubscriptionPurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotSubscriptionPurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotSubscriptionPurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotSubscriptionPurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotSubscriptionPurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotSubscriptionPurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotSubscriptionPurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotSubscriptionPurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotSubscriptionPurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotSubscriptionPurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotSubscriptionPurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotSubscriptionPurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotSubscriptionPurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBotSubscriptionSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBotSubscriptionSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBotSubscriptionSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBotSubscriptionSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBotSubscriptionSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBotSubscriptionSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBotSubscriptionSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBotSubscriptionSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBotSubscriptionSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBotSubscriptionSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBotSubscriptionSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBotSubscriptionSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBotSubscriptionSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBotSubscriptionSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBotSubscriptionSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBotSubscriptionSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBotSubscriptionSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBotSubscriptionSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBotSubscriptionSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBotSubscriptionSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBotSubscriptionSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBotSubscriptionSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBotSubscriptionSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBotSubscriptionSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBotSubscriptionSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBotSubscriptionSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBotSubscriptionSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBotSubscriptionSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBotSubscriptionSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBotSubscriptionSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBotSubscriptionSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBotSubscriptionSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBotSubscriptionSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBotSubscriptionSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBotSubscriptionSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBotSubscriptionSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBotSubscriptionSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBotSubscriptionSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBotSubscriptionSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBotSubscriptionSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBotSubscriptionSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBotSubscriptionSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBotSubscriptionSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBusinessBotTransferReceive) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBusinessBotTransferReceive) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBusinessBotTransferReceive) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBusinessBotTransferReceive) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBusinessBotTransferReceive) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBusinessBotTransferReceive) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBusinessBotTransferReceive) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBusinessBotTransferReceive) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBusinessBotTransferReceive) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBusinessBotTransferReceive) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBusinessBotTransferReceive) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBusinessBotTransferReceive) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBusinessBotTransferReceive) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBusinessBotTransferReceive) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBusinessBotTransferReceive) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBusinessBotTransferReceive) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBusinessBotTransferReceive) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBusinessBotTransferReceive) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBusinessBotTransferReceive) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBusinessBotTransferReceive) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBusinessBotTransferReceive) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBusinessBotTransferReceive) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBusinessBotTransferReceive) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBusinessBotTransferReceive) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBusinessBotTransferReceive) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBusinessBotTransferReceive) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBusinessBotTransferReceive) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBusinessBotTransferReceive) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBusinessBotTransferReceive) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBusinessBotTransferReceive) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBusinessBotTransferReceive) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBusinessBotTransferReceive) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBusinessBotTransferReceive) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBusinessBotTransferReceive) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBusinessBotTransferReceive) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBusinessBotTransferReceive) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBusinessBotTransferReceive) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBusinessBotTransferReceive) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBusinessBotTransferReceive) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBusinessBotTransferReceive) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBusinessBotTransferReceive) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBusinessBotTransferReceive) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBusinessBotTransferReceive) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeBusinessBotTransferSend) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeBusinessBotTransferSend) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeBusinessBotTransferSend) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeBusinessBotTransferSend) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeBusinessBotTransferSend) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeBusinessBotTransferSend) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeBusinessBotTransferSend) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeBusinessBotTransferSend) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeBusinessBotTransferSend) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeBusinessBotTransferSend) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeBusinessBotTransferSend) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeBusinessBotTransferSend) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeBusinessBotTransferSend) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeBusinessBotTransferSend) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeBusinessBotTransferSend) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeBusinessBotTransferSend) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeBusinessBotTransferSend) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeBusinessBotTransferSend) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeBusinessBotTransferSend) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeBusinessBotTransferSend) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeBusinessBotTransferSend) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeBusinessBotTransferSend) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeBusinessBotTransferSend) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeBusinessBotTransferSend) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeBusinessBotTransferSend) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeBusinessBotTransferSend) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeBusinessBotTransferSend) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeBusinessBotTransferSend) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeBusinessBotTransferSend) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeBusinessBotTransferSend) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeBusinessBotTransferSend) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeBusinessBotTransferSend) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeBusinessBotTransferSend) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeBusinessBotTransferSend) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeBusinessBotTransferSend) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeBusinessBotTransferSend) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeBusinessBotTransferSend) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeBusinessBotTransferSend) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeBusinessBotTransferSend) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeBusinessBotTransferSend) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeBusinessBotTransferSend) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeBusinessBotTransferSend) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeBusinessBotTransferSend) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelPaidMediaPurchase) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeChannelPaidMediaPurchase) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeChannelPaidMediaPurchase) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeChannelPaidMediaPurchase) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, s.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeChannelPaidMediaPurchase) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, s.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeChannelPaidMediaPurchase) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, s.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeChannelPaidMediaPurchase) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, s.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeChannelPaidMediaPurchase) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, s.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeChannelPaidMediaPurchase) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeChannelPaidMediaPurchase) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeChannelPaidMediaPurchase) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, s.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeChannelPaidMediaPurchase) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (s StarTransactionTypeChannelPaidMediaPurchase) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(s.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeChannelPaidMediaPurchase) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeChannelPaidMediaPurchase) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeChannelPaidMediaPurchase) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, s.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, s.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeChannelPaidMediaPurchase) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeChannelPaidMediaPurchase) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeChannelPaidMediaPurchase) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeChannelPaidMediaPurchase) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, s.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeChannelPaidMediaPurchase) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, s.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, s.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeChannelPaidMediaPurchase) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, s.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeChannelPaidMediaPurchase) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeChannelPaidMediaPurchase) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, s.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, s.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, s.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeChannelPaidMediaPurchase) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, s.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeChannelPaidMediaPurchase) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeChannelPaidMediaPurchase) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, s.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, s.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, s.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeChannelPaidMediaPurchase) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeChannelPaidMediaPurchase) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeChannelPaidMediaPurchase) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeChannelPaidMediaPurchase) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeChannelPaidMediaPurchase) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeChannelPaidMediaPurchase) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeChannelPaidMediaPurchase) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelPaidMediaPurchase) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, s.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeChannelPaidMediaPurchase) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, s.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeChannelPaidMediaPurchase) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, s.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeChannelPaidMediaPurchase) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, s.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(s.ChatId, s.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, s.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, s.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, s.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, s.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, s.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, s.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, s.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, s.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, s.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, s.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, s.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, s.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeChannelPaidMediaPurchase) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, s.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeChannelPaidMediaPurchase) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, s.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeChannelPaidMediaPurchase) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, s.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, s.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeChannelPaidMediaPurchase) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeChannelPaidMediaPurchase) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeChannelPaidMediaPurchase) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeChannelPaidMediaPurchase) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelPaidMediaPurchase) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeChannelPaidMediaPurchase) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeChannelPaidMediaPurchase) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, s.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeChannelPaidMediaPurchase) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeChannelPaidMediaPurchase) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeChannelPaidMediaPurchase) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeChannelPaidMediaPurchase) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, s.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeChannelPaidMediaPurchase) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeChannelPaidMediaPurchase) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeChannelPaidMediaPurchase) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, s.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeChannelPaidMediaPurchase) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, s.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeChannelPaidMediaPurchase) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelPaidMediaPurchase) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeChannelPaidMediaPurchase) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (s StarTransactionTypeChannelPaidMediaPurchase) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(s.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeChannelPaidMediaPurchase) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, s.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, s.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeChannelPaidMediaPurchase) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, s.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, s.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeChannelPaidMediaPurchase) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeChannelPaidMediaPurchase) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeChannelPaidMediaPurchase) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeChannelPaidMediaPurchase) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeChannelPaidMediaPurchase) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeChannelPaidMediaPurchase) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, s.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, s.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s StarTransactionTypeChannelPaidMediaPurchase) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, s.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeChannelPaidMediaPurchase) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeChannelPaidMediaPurchase) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeChannelPaidMediaPurchase) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeChannelPaidMediaPurchase) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeChannelPaidMediaPurchase) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeChannelPaidMediaPurchase) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeChannelPaidMediaPurchase) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeChannelPaidMediaPurchase) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, s.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeChannelPaidMediaPurchase) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeChannelPaidMediaPurchase) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeChannelPaidMediaPurchase) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelPaidMediaPurchase) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, s.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeChannelPaidMediaPurchase) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, s.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeChannelPaidMediaPurchase) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, s.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeChannelPaidMediaPurchase) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, s.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeChannelPaidMediaPurchase) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeChannelPaidMediaPurchase) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeChannelPaidMediaPurchase) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, s.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeChannelPaidMediaPurchase) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeChannelPaidMediaPurchase) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeChannelPaidMediaPurchase) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, s.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeChannelPaidMediaPurchase) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, s.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeChannelPaidMediaPurchase) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeChannelPaidMediaPurchase) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeChannelPaidMediaPurchase) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, s.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, s.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeChannelPaidMediaPurchase) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelPaidMediaPurchase) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeChannelPaidMediaPurchase) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, s.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeChannelPaidMediaPurchase) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, s.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeChannelPaidMediaPurchase) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeChannelPaidMediaPurchase) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelPaidMediaSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeChannelPaidMediaSale) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, s.MessageId, tasks)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeChannelPaidMediaSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeChannelPaidMediaSale) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, s.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeChannelPaidMediaSale) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, s.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeChannelPaidMediaSale) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, s.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeChannelPaidMediaSale) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, s.MessageId, starCount, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeChannelPaidMediaSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeChannelPaidMediaSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeChannelPaidMediaSale) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, s.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (s StarTransactionTypeChannelPaidMediaSale) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(s.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeChannelPaidMediaSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeChannelPaidMediaSale) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, s.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeChannelPaidMediaSale) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, s.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidMediaSale) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, s.MessageId)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeChannelPaidMediaSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeChannelPaidMediaSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeChannelPaidMediaSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeChannelPaidMediaSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeChannelPaidMediaSale) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, s.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeChannelPaidMediaSale) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, s.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeChannelPaidMediaSale) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, s.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, s.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, s.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, s.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeChannelPaidMediaSale) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, s.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, s.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, s.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeChannelPaidMediaSale) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, s.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (s StarTransactionTypeChannelPaidMediaSale) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, s.MessageId, inputMessageContent)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeChannelPaidMediaSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeChannelPaidMediaSale) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, s.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeChannelPaidMediaSale) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, s.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeChannelPaidMediaSale) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, s.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelPaidMediaSale) GetGameHighScores(client *Client, chatId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, s.MessageId, s.UserId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeChannelPaidMediaSale) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, s.MessageId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeChannelPaidMediaSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeChannelPaidMediaSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeChannelPaidMediaSale) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, s.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeChannelPaidMediaSale) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, s.MessageId, buttonId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeChannelPaidMediaSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeChannelPaidMediaSale) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, s.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, s.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, s.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, s.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, s.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, s.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, s.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, s.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeChannelPaidMediaSale) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, s.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, s.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, s.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, s.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, s.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeChannelPaidMediaSale) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, s.MessageId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeChannelPaidMediaSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeChannelPaidMediaSale) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, s.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeChannelPaidMediaSale) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, s.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeChannelPaidMediaSale) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, s.MessageId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeChannelPaidMediaSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeChannelPaidMediaSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelPaidMediaSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeChannelPaidMediaSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeChannelPaidMediaSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeChannelPaidMediaSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeChannelPaidMediaSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeChannelPaidMediaSale) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, s.MessageId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeChannelPaidMediaSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeChannelPaidMediaSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeChannelPaidMediaSale) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, s.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeChannelPaidMediaSale) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, s.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeChannelPaidMediaSale) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, s.MessageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeChannelPaidMediaSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelPaidMediaSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (s StarTransactionTypeChannelPaidMediaSale) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(s.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeChannelPaidMediaSale) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, s.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeChannelPaidMediaSale) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, s.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeChannelPaidMediaSale) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, s.MessageId)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeChannelPaidMediaSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeChannelPaidMediaSale) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, s.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidMediaSale) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, s.MessageId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeChannelPaidMediaSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeChannelPaidMediaSale) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, s.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeChannelPaidMediaSale) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, s.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s StarTransactionTypeChannelPaidMediaSale) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, s.MessageId)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeChannelPaidMediaSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeChannelPaidMediaSale) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, s.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelPaidMediaSale) SetGameScore(client *Client, chatId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, s.MessageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeChannelPaidMediaSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeChannelPaidMediaSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeChannelPaidMediaSale) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, s.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeChannelPaidMediaSale) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, s.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeChannelPaidMediaSale) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, s.MessageId, typeField)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeChannelPaidMediaSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeChannelPaidMediaSale) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, s.MessageId, optionIds)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeChannelPaidMediaSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeChannelPaidMediaSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeChannelPaidMediaSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeChannelPaidMediaSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeChannelPaidMediaSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeChannelPaidMediaSale) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, s.MessageId, buttonId, sharedChatId, onlyCheck)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeChannelPaidMediaSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeChannelPaidMediaSale) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, s.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeChannelPaidMediaSale) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, s.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeChannelPaidMediaSale) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, s.MessageId, opts)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeChannelPaidMediaSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeChannelPaidMediaSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeChannelPaidMediaSale) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, s.MessageId, translateToLanguageCode)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelPaidMediaSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeChannelPaidMediaSale) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, s.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeChannelPaidMediaSale) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, s.MessageId)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeChannelPaidMediaSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelPaidReactionReceive) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeChannelPaidReactionReceive) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, s.MessageId, tasks)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeChannelPaidReactionReceive) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeChannelPaidReactionReceive) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, s.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeChannelPaidReactionReceive) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, s.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeChannelPaidReactionReceive) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, s.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeChannelPaidReactionReceive) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, s.MessageId, starCount, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeChannelPaidReactionReceive) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeChannelPaidReactionReceive) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeChannelPaidReactionReceive) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, s.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (s StarTransactionTypeChannelPaidReactionReceive) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(s.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeChannelPaidReactionReceive) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeChannelPaidReactionReceive) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, s.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeChannelPaidReactionReceive) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, s.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidReactionReceive) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, s.MessageId)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeChannelPaidReactionReceive) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeChannelPaidReactionReceive) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeChannelPaidReactionReceive) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeChannelPaidReactionReceive) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeChannelPaidReactionReceive) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, s.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeChannelPaidReactionReceive) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, s.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeChannelPaidReactionReceive) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, s.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, s.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, s.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, s.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeChannelPaidReactionReceive) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, s.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, s.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, s.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeChannelPaidReactionReceive) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, s.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (s StarTransactionTypeChannelPaidReactionReceive) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, s.MessageId, inputMessageContent)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeChannelPaidReactionReceive) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeChannelPaidReactionReceive) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, s.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeChannelPaidReactionReceive) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, s.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeChannelPaidReactionReceive) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, s.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelPaidReactionReceive) GetGameHighScores(client *Client, chatId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, s.MessageId, s.UserId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeChannelPaidReactionReceive) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, s.MessageId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeChannelPaidReactionReceive) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeChannelPaidReactionReceive) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeChannelPaidReactionReceive) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, s.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeChannelPaidReactionReceive) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, s.MessageId, buttonId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeChannelPaidReactionReceive) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, s.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, s.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, s.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, s.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, s.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, s.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, s.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, s.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, s.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, s.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, s.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, s.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, s.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeChannelPaidReactionReceive) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, s.MessageId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeChannelPaidReactionReceive) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeChannelPaidReactionReceive) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, s.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeChannelPaidReactionReceive) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, s.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeChannelPaidReactionReceive) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, s.MessageId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeChannelPaidReactionReceive) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeChannelPaidReactionReceive) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelPaidReactionReceive) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeChannelPaidReactionReceive) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeChannelPaidReactionReceive) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeChannelPaidReactionReceive) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeChannelPaidReactionReceive) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeChannelPaidReactionReceive) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, s.MessageId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeChannelPaidReactionReceive) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeChannelPaidReactionReceive) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeChannelPaidReactionReceive) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, s.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeChannelPaidReactionReceive) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, s.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeChannelPaidReactionReceive) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, s.MessageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeChannelPaidReactionReceive) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelPaidReactionReceive) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (s StarTransactionTypeChannelPaidReactionReceive) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(s.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeChannelPaidReactionReceive) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, s.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeChannelPaidReactionReceive) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, s.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeChannelPaidReactionReceive) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, s.MessageId)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeChannelPaidReactionReceive) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeChannelPaidReactionReceive) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, s.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidReactionReceive) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, s.MessageId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeChannelPaidReactionReceive) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeChannelPaidReactionReceive) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, s.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeChannelPaidReactionReceive) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, s.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s StarTransactionTypeChannelPaidReactionReceive) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, s.MessageId)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeChannelPaidReactionReceive) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeChannelPaidReactionReceive) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, s.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelPaidReactionReceive) SetGameScore(client *Client, chatId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, s.MessageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeChannelPaidReactionReceive) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeChannelPaidReactionReceive) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeChannelPaidReactionReceive) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, s.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeChannelPaidReactionReceive) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, s.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeChannelPaidReactionReceive) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, s.MessageId, typeField)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeChannelPaidReactionReceive) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeChannelPaidReactionReceive) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, s.MessageId, optionIds)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeChannelPaidReactionReceive) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeChannelPaidReactionReceive) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeChannelPaidReactionReceive) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeChannelPaidReactionReceive) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeChannelPaidReactionReceive) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeChannelPaidReactionReceive) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, s.MessageId, buttonId, sharedChatId, onlyCheck)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeChannelPaidReactionReceive) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeChannelPaidReactionReceive) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, s.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeChannelPaidReactionReceive) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, s.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeChannelPaidReactionReceive) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, s.MessageId, opts)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeChannelPaidReactionReceive) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeChannelPaidReactionReceive) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeChannelPaidReactionReceive) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, s.MessageId, translateToLanguageCode)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelPaidReactionReceive) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeChannelPaidReactionReceive) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, s.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeChannelPaidReactionReceive) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, s.MessageId)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeChannelPaidReactionReceive) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelPaidReactionSend) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeChannelPaidReactionSend) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeChannelPaidReactionSend) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeChannelPaidReactionSend) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, s.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeChannelPaidReactionSend) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, s.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeChannelPaidReactionSend) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeChannelPaidReactionSend) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, s.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeChannelPaidReactionSend) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, s.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeChannelPaidReactionSend) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, s.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeChannelPaidReactionSend) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeChannelPaidReactionSend) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeChannelPaidReactionSend) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, s.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeChannelPaidReactionSend) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (s StarTransactionTypeChannelPaidReactionSend) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(s.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeChannelPaidReactionSend) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeChannelPaidReactionSend) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeChannelPaidReactionSend) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeChannelPaidReactionSend) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, s.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeChannelPaidReactionSend) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, s.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeChannelPaidReactionSend) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidReactionSend) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeChannelPaidReactionSend) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeChannelPaidReactionSend) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeChannelPaidReactionSend) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, s.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeChannelPaidReactionSend) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, s.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeChannelPaidReactionSend) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeChannelPaidReactionSend) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, s.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelPaidReactionSend) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeChannelPaidReactionSend) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeChannelPaidReactionSend) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeChannelPaidReactionSend) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeChannelPaidReactionSend) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, s.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeChannelPaidReactionSend) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeChannelPaidReactionSend) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, s.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, s.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, s.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeChannelPaidReactionSend) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (s StarTransactionTypeChannelPaidReactionSend) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, s.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeChannelPaidReactionSend) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeChannelPaidReactionSend) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeChannelPaidReactionSend) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, s.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeChannelPaidReactionSend) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, s.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeChannelPaidReactionSend) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeChannelPaidReactionSend) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeChannelPaidReactionSend) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeChannelPaidReactionSend) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeChannelPaidReactionSend) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeChannelPaidReactionSend) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeChannelPaidReactionSend) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeChannelPaidReactionSend) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeChannelPaidReactionSend) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeChannelPaidReactionSend) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeChannelPaidReactionSend) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeChannelPaidReactionSend) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeChannelPaidReactionSend) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeChannelPaidReactionSend) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeChannelPaidReactionSend) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeChannelPaidReactionSend) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeChannelPaidReactionSend) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeChannelPaidReactionSend) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeChannelPaidReactionSend) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeChannelPaidReactionSend) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeChannelPaidReactionSend) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, s.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeChannelPaidReactionSend) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeChannelPaidReactionSend) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeChannelPaidReactionSend) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeChannelPaidReactionSend) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeChannelPaidReactionSend) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeChannelPaidReactionSend) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeChannelPaidReactionSend) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeChannelPaidReactionSend) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeChannelPaidReactionSend) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeChannelPaidReactionSend) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeChannelPaidReactionSend) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeChannelPaidReactionSend) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeChannelPaidReactionSend) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelPaidReactionSend) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeChannelPaidReactionSend) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeChannelPaidReactionSend) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeChannelPaidReactionSend) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeChannelPaidReactionSend) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeChannelPaidReactionSend) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeChannelPaidReactionSend) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelPaidReactionSend) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, s.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeChannelPaidReactionSend) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, s.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeChannelPaidReactionSend) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeChannelPaidReactionSend) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeChannelPaidReactionSend) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, s.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeChannelPaidReactionSend) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, s.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeChannelPaidReactionSend) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeChannelPaidReactionSend) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeChannelPaidReactionSend) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(s.ChatId, s.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, s.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, s.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, s.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, s.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, s.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, s.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, s.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeChannelPaidReactionSend) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, s.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, s.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeChannelPaidReactionSend) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, s.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, s.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, s.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeChannelPaidReactionSend) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, s.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeChannelPaidReactionSend) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, s.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeChannelPaidReactionSend) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, s.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeChannelPaidReactionSend) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, s.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeChannelPaidReactionSend) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeChannelPaidReactionSend) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeChannelPaidReactionSend) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeChannelPaidReactionSend) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelPaidReactionSend) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeChannelPaidReactionSend) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeChannelPaidReactionSend) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}
