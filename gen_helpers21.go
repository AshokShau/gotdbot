package gotdbot

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateForumTopic) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateForumTopic) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateForumTopic) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateForumTopic) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateForumTopic) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateForumTopic) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateForumTopic) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateForumTopic) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateForumTopic) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateForumTopic) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateForumTopic) DeleteForumTopic(client *Client) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, u.ForumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateForumTopic) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateForumTopic) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateForumTopic) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateForumTopic) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateForumTopic) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateForumTopic) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateForumTopic) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateForumTopic) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateForumTopic) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateForumTopic) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateForumTopic) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateForumTopic) EditForumTopic(client *Client, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, u.ForumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateForumTopic) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateForumTopic) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateForumTopic) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateForumTopic) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateForumTopic) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateForumTopic) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateForumTopic) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateForumTopic) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateForumTopic) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateForumTopic) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateForumTopic) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateForumTopic) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateForumTopic) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateForumTopic) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateForumTopic) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateForumTopic) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateForumTopic) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateForumTopic) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateForumTopic) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateForumTopic) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateForumTopic) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateForumTopic) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateForumTopic) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateForumTopic) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateForumTopic) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateForumTopic) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateForumTopic) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateForumTopic) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateForumTopic) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateForumTopic) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateForumTopic) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateForumTopic) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateForumTopic) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateForumTopic) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateForumTopic) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateForumTopic) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateForumTopic) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateForumTopic) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateForumTopic) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateForumTopic) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateForumTopic) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateForumTopic) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateForumTopic) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateForumTopic) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateForumTopic) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateForumTopic) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateForumTopic) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateForumTopic) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateForumTopic) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateForumTopic) GetForumTopic(client *Client) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, u.ForumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateForumTopic) GetForumTopicHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, u.ForumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateForumTopic) GetForumTopicLink(client *Client) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, u.ForumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateForumTopic) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateForumTopic) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateForumTopic) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateForumTopic) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateForumTopic) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateForumTopic) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateForumTopic) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateForumTopic) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateForumTopic) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateForumTopic) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateForumTopic) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateForumTopic) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateForumTopic) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateForumTopic) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateForumTopic) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateForumTopic) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateForumTopic) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateForumTopic) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateForumTopic) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateForumTopic) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateForumTopic) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateForumTopic) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateForumTopic) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateForumTopic) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateForumTopic) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateForumTopic) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateForumTopic) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateForumTopic) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateForumTopic) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateForumTopic) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateForumTopic) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateForumTopic) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateForumTopic) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateForumTopic) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateForumTopic) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateForumTopic) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateForumTopic) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateForumTopic) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateForumTopic) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateForumTopic) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateForumTopic) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateForumTopic) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateForumTopic) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateForumTopic) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateForumTopic) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateForumTopic) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateForumTopic) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateForumTopic) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateForumTopic) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateForumTopic) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateForumTopic) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateForumTopic) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateForumTopic) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateForumTopic) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateForumTopic) ReadAllForumTopicMentions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, u.ForumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateForumTopic) ReadAllForumTopicReactions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, u.ForumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateForumTopic) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateForumTopic) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateForumTopic) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateForumTopic) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateForumTopic) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateForumTopic) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateForumTopic) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateForumTopic) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateForumTopic) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateForumTopic) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateForumTopic) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateForumTopic) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateForumTopic) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateForumTopic) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateForumTopic) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateForumTopic) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateForumTopic) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateForumTopic) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateForumTopic) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateForumTopic) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateForumTopic) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateForumTopic) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateForumTopic) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateForumTopic) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateForumTopic) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateForumTopic) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateForumTopic) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateForumTopic) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateForumTopic) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateForumTopic) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateForumTopic) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateForumTopic) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateForumTopic) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateForumTopic) SendTextMessageDraft(client *Client, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, u.ForumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateForumTopic) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, u.IsPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateForumTopic) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateForumTopic) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateForumTopic) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateForumTopic) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateForumTopic) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateForumTopic) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateForumTopic) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateForumTopic) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateForumTopic) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateForumTopic) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateForumTopic) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateForumTopic) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateForumTopic) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateForumTopic) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateForumTopic) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateForumTopic) SetChatNotificationSettings(client *Client) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, u.NotificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateForumTopic) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateForumTopic) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateForumTopic) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateForumTopic) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateForumTopic) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateForumTopic) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateForumTopic) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateForumTopic) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateForumTopic) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateForumTopic) SetForumTopicNotificationSettings(client *Client) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, u.ForumTopicId, u.NotificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateForumTopic) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateForumTopic) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateForumTopic) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateForumTopic) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateForumTopic) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateForumTopic) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateForumTopic) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateForumTopic) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateForumTopic) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateForumTopic) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateForumTopic) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateForumTopic) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateForumTopic) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateForumTopic) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateForumTopic) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateForumTopic) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateForumTopic) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateForumTopic) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateForumTopic) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateForumTopic) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateForumTopic) ToggleChatIsPinned(client *Client, chatList *ChatList) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, u.IsPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateForumTopic) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateForumTopic) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateForumTopic) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateForumTopic) ToggleForumTopicIsClosed(client *Client, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, u.ForumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateForumTopic) ToggleForumTopicIsPinned(client *Client) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, u.ForumTopicId, u.IsPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateForumTopic) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (u UpdateForumTopic) ToggleSavedMessagesTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, u.IsPinned)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateForumTopic) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateForumTopic) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateForumTopic) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateForumTopic) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateForumTopic) UnpinAllForumTopicMessages(client *Client) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, u.ForumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateForumTopic) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateForumTopic) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateForumTopic) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateGroupCallMessagesDeleted) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateGroupCallMessagesDeleted) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateGroupCallMessagesDeleted) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateGroupCallMessagesDeleted) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateGroupCallMessagesDeleted) DeleteGroupCallMessages(client *Client, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, u.MessageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateGroupCallMessagesDeleted) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateGroupCallMessagesDeleted) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateGroupCallMessagesDeleted) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateGroupCallMessagesDeleted) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateGroupCallMessagesDeleted) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateGroupCallMessagesDeleted) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateGroupCallMessagesDeleted) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateGroupCallMessagesDeleted) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateGroupCallMessagesDeleted) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateGroupCallMessagesDeleted) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateGroupCallMessagesDeleted) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateGroupCallMessagesDeleted) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateGroupCallMessagesDeleted) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateGroupCallMessagesDeleted) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateGroupCallMessagesDeleted) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateGroupCallMessagesDeleted) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateGroupCallMessagesDeleted) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateGroupCallMessagesDeleted) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateGroupCallMessagesDeleted) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateGroupCallMessagesDeleted) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateGroupCallMessagesDeleted) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateGroupCallMessagesDeleted) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateGroupCallMessagesDeleted) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateGroupCallMessagesDeleted) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateGroupCallMessagesDeleted) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateGroupCallMessagesDeleted) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateGroupCallMessagesDeleted) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateGroupCallMessagesDeleted) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateGroupCallMessagesDeleted) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateGroupCallMessagesDeleted) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateGroupCallMessagesDeleted) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateGroupCallMessagesDeleted) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateGroupCallMessageSendFailed) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateGroupCallMessageSendFailed) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateGroupCallMessageSendFailed) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateGroupCallMessageSendFailed) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateGroupCallMessageSendFailed) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateGroupCallMessageSendFailed) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateGroupCallMessageSendFailed) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateGroupCallMessageSendFailed) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateGroupCallMessageSendFailed) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateGroupCallMessageSendFailed) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateGroupCallMessageSendFailed) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateGroupCallMessageSendFailed) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateGroupCallMessageSendFailed) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateGroupCallMessageSendFailed) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateGroupCallMessageSendFailed) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateGroupCallMessageSendFailed) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateGroupCallMessageSendFailed) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateGroupCallMessageSendFailed) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateGroupCallMessageSendFailed) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateGroupCallMessageSendFailed) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateGroupCallMessageSendFailed) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateGroupCallMessageSendFailed) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateGroupCallMessageSendFailed) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateGroupCallMessageSendFailed) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateGroupCallMessageSendFailed) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateGroupCallMessageSendFailed) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateGroupCallMessageSendFailed) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateGroupCallMessageSendFailed) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateGroupCallMessageSendFailed) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateGroupCallMessageSendFailed) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateGroupCallMessageSendFailed) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateGroupCallMessageSendFailed) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateGroupCallMessageSendFailed) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateGroupCallMessageSendFailed) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// TestReturnError is a helper method for Client.TestReturnError
func (u UpdateGroupCallMessageSendFailed) TestReturnError(client *Client) (*Error, error) {
	return client.TestReturnError(u.Error)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateGroupCallMessageSendFailed) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateGroupCallMessageSendFailed) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateGroupCallMessageSendFailed) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateGroupCallParticipant) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateGroupCallParticipant) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateGroupCallParticipant) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateGroupCallParticipant) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateGroupCallParticipant) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateGroupCallParticipant) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateGroupCallParticipant) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateGroupCallParticipant) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateGroupCallParticipant) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateGroupCallParticipant) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateGroupCallParticipant) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateGroupCallParticipant) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateGroupCallParticipant) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateGroupCallParticipant) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateGroupCallParticipant) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateGroupCallParticipant) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateGroupCallParticipant) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateGroupCallParticipant) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateGroupCallParticipant) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateGroupCallParticipant) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateGroupCallParticipant) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateGroupCallParticipant) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateGroupCallParticipant) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateGroupCallParticipant) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateGroupCallParticipant) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateGroupCallParticipant) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateGroupCallParticipant) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateGroupCallParticipant) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateGroupCallParticipant) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateGroupCallParticipant) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateGroupCallParticipant) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateGroupCallParticipant) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateGroupCallParticipant) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateGroupCallParticipant) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateGroupCallParticipant) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateGroupCallParticipant) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateGroupCallParticipant) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateGroupCallParticipant) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateGroupCallParticipant) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateGroupCallParticipant) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateGroupCallParticipant) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateGroupCallParticipant) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateGroupCallParticipants) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateGroupCallParticipants) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateGroupCallParticipants) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateGroupCallParticipants) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateGroupCallParticipants) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateGroupCallParticipants) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateGroupCallParticipants) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateGroupCallParticipants) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateGroupCallParticipants) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateGroupCallParticipants) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateGroupCallParticipants) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateGroupCallParticipants) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateGroupCallParticipants) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateGroupCallParticipants) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateGroupCallParticipants) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateGroupCallParticipants) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateGroupCallParticipants) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateGroupCallParticipants) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateGroupCallParticipants) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateGroupCallParticipants) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateGroupCallParticipants) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateGroupCallParticipants) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateGroupCallParticipants) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateGroupCallParticipants) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateGroupCallParticipants) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateGroupCallParticipants) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateGroupCallParticipants) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateGroupCallParticipants) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateGroupCallParticipants) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateGroupCallParticipants) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateGroupCallParticipants) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateGroupCallParticipants) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateGroupCallParticipants) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateGroupCallParticipants) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateGroupCallParticipants) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateGroupCallParticipants) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateGroupCallParticipants) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateGroupCallParticipants) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateGroupCallParticipants) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateGroupCallParticipants) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateGroupCallParticipants) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateGroupCallParticipants) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateGroupCallVerificationState) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateGroupCallVerificationState) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateGroupCallVerificationState) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateGroupCallVerificationState) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateGroupCallVerificationState) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateGroupCallVerificationState) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateGroupCallVerificationState) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateGroupCallVerificationState) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateGroupCallVerificationState) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateGroupCallVerificationState) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateGroupCallVerificationState) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateGroupCallVerificationState) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateGroupCallVerificationState) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateGroupCallVerificationState) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateGroupCallVerificationState) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateGroupCallVerificationState) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateGroupCallVerificationState) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateGroupCallVerificationState) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateGroupCallVerificationState) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateGroupCallVerificationState) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateGroupCallVerificationState) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateGroupCallVerificationState) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateGroupCallVerificationState) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateGroupCallVerificationState) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateGroupCallVerificationState) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateGroupCallVerificationState) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateGroupCallVerificationState) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateGroupCallVerificationState) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateGroupCallVerificationState) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateGroupCallVerificationState) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateGroupCallVerificationState) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateGroupCallVerificationState) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateGroupCallVerificationState) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateGroupCallVerificationState) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateGroupCallVerificationState) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateGroupCallVerificationState) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateGroupCallVerificationState) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateGroupCallVerificationState) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateGroupCallVerificationState) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateGroupCallVerificationState) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateGroupCallVerificationState) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateGroupCallVerificationState) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdateInstalledStickerSets) CreateNewStickerSet(client *Client, userId int64, title string, name string, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, name, u.StickerType, needsRepainting, stickers, source)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateInstalledStickerSets) GetAllStickerEmojis(client *Client, query string, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(u.StickerType, query, chatId, returnOnlyMainEmoji)
}

// GetArchivedStickerSets is a helper method for Client.GetArchivedStickerSets
func (u UpdateInstalledStickerSets) GetArchivedStickerSets(client *Client, offsetStickerSetId string, limit int32) (*StickerSets, error) {
	return client.GetArchivedStickerSets(u.StickerType, offsetStickerSetId, limit)
}

// GetInstalledStickerSets is a helper method for Client.GetInstalledStickerSets
func (u UpdateInstalledStickerSets) GetInstalledStickerSets(client *Client) (*StickerSets, error) {
	return client.GetInstalledStickerSets(u.StickerType)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateInstalledStickerSets) GetStickers(client *Client, query string, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(u.StickerType, query, limit, chatId)
}

// GetTrendingStickerSets is a helper method for Client.GetTrendingStickerSets
func (u UpdateInstalledStickerSets) GetTrendingStickerSets(client *Client, offset int32, limit int32) (*TrendingStickerSets, error) {
	return client.GetTrendingStickerSets(u.StickerType, offset, limit)
}

// ReorderInstalledStickerSets is a helper method for Client.ReorderInstalledStickerSets
func (u UpdateInstalledStickerSets) ReorderInstalledStickerSets(client *Client) (*Ok, error) {
	return client.ReorderInstalledStickerSets(u.StickerType, u.StickerSetIds)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (u UpdateInstalledStickerSets) SearchInstalledStickerSets(client *Client, query string, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(u.StickerType, query, limit)
}

// SearchStickers is a helper method for Client.SearchStickers
func (u UpdateInstalledStickerSets) SearchStickers(client *Client, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(u.StickerType, emojis, query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (u UpdateInstalledStickerSets) SearchStickerSets(client *Client, query string) (*StickerSets, error) {
	return client.SearchStickerSets(u.StickerType, query)
}

// ViewTrendingStickerSets is a helper method for Client.ViewTrendingStickerSets
func (u UpdateInstalledStickerSets) ViewTrendingStickerSets(client *Client) (*Ok, error) {
	return client.ViewTrendingStickerSets(u.StickerSetIds)
}

// AddCustomServerLanguagePack is a helper method for Client.AddCustomServerLanguagePack
func (u UpdateLanguagePackStrings) AddCustomServerLanguagePack(client *Client) (*Ok, error) {
	return client.AddCustomServerLanguagePack(u.LanguagePackId)
}

// DeleteLanguagePack is a helper method for Client.DeleteLanguagePack
func (u UpdateLanguagePackStrings) DeleteLanguagePack(client *Client) (*Ok, error) {
	return client.DeleteLanguagePack(u.LanguagePackId)
}

// GetLanguagePackInfo is a helper method for Client.GetLanguagePackInfo
func (u UpdateLanguagePackStrings) GetLanguagePackInfo(client *Client) (*LanguagePackInfo, error) {
	return client.GetLanguagePackInfo(u.LanguagePackId)
}

// GetLanguagePackString is a helper method for Client.GetLanguagePackString
func (u UpdateLanguagePackStrings) GetLanguagePackString(client *Client, languagePackDatabasePath string, key string) (*LanguagePackStringValue, error) {
	return client.GetLanguagePackString(languagePackDatabasePath, u.LocalizationTarget, u.LanguagePackId, key)
}

// GetLanguagePackStrings is a helper method for Client.GetLanguagePackStrings
func (u UpdateLanguagePackStrings) GetLanguagePackStrings(client *Client, keys []string) (*LanguagePackStrings, error) {
	return client.GetLanguagePackStrings(u.LanguagePackId, keys)
}

// SetCustomLanguagePack is a helper method for Client.SetCustomLanguagePack
func (u UpdateLanguagePackStrings) SetCustomLanguagePack(client *Client, info *LanguagePackInfo) (*Ok, error) {
	return client.SetCustomLanguagePack(info, u.Strings)
}

// SetCustomLanguagePackString is a helper method for Client.SetCustomLanguagePackString
func (u UpdateLanguagePackStrings) SetCustomLanguagePackString(client *Client, newString *LanguagePackString) (*Ok, error) {
	return client.SetCustomLanguagePackString(u.LanguagePackId, newString)
}

// SynchronizeLanguagePack is a helper method for Client.SynchronizeLanguagePack
func (u UpdateLanguagePackStrings) SynchronizeLanguagePack(client *Client) (*Ok, error) {
	return client.SynchronizeLanguagePack(u.LanguagePackId)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (u UpdateLiveStoryTopDonors) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(u.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (u UpdateLiveStoryTopDonors) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(u.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (u UpdateLiveStoryTopDonors) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(u.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (u UpdateLiveStoryTopDonors) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(u.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (u UpdateLiveStoryTopDonors) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(u.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpdateLiveStoryTopDonors) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(u.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (u UpdateLiveStoryTopDonors) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(u.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (u UpdateLiveStoryTopDonors) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(u.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (u UpdateLiveStoryTopDonors) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(u.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (u UpdateLiveStoryTopDonors) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(u.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (u UpdateLiveStoryTopDonors) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(u.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (u UpdateLiveStoryTopDonors) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(u.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (u UpdateLiveStoryTopDonors) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(u.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (u UpdateLiveStoryTopDonors) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(u.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (u UpdateLiveStoryTopDonors) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(u.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (u UpdateLiveStoryTopDonors) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(u.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (u UpdateLiveStoryTopDonors) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(u.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateLiveStoryTopDonors) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(u.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (u UpdateLiveStoryTopDonors) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(u.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (u UpdateLiveStoryTopDonors) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(u.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (u UpdateLiveStoryTopDonors) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(u.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (u UpdateLiveStoryTopDonors) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(u.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (u UpdateLiveStoryTopDonors) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(u.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (u UpdateLiveStoryTopDonors) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(u.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (u UpdateLiveStoryTopDonors) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(u.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdateLiveStoryTopDonors) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(u.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (u UpdateLiveStoryTopDonors) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(u.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (u UpdateLiveStoryTopDonors) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(u.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (u UpdateLiveStoryTopDonors) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(u.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (u UpdateLiveStoryTopDonors) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(u.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateLiveStoryTopDonors) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(u.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateLiveStoryTopDonors) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(u.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdateLiveStoryTopDonors) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(u.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (u UpdateLiveStoryTopDonors) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(u.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (u UpdateLiveStoryTopDonors) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(u.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (u UpdateLiveStoryTopDonors) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(u.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (u UpdateLiveStoryTopDonors) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(u.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (u UpdateLiveStoryTopDonors) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(u.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (u UpdateLiveStoryTopDonors) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(u.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateLiveStoryTopDonors) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(u.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (u UpdateLiveStoryTopDonors) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(u.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (u UpdateLiveStoryTopDonors) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(u.GroupCallId, muteNewParticipants)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateMessageContent) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateMessageContent) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateMessageContent) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateMessageContent) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateMessageContent) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateMessageContent) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateMessageContent) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateMessageContent) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateMessageContent) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateMessageContent) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateMessageContent) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateMessageContent) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateMessageContent) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateMessageContent) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateMessageContent) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateMessageContent) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateMessageContent) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateMessageContent) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateMessageContent) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateMessageContent) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateMessageContent) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateMessageContent) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateMessageContent) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateMessageContent) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateMessageContent) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateMessageContent) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateMessageContent) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateMessageContent) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateMessageContent) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateMessageContent) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateMessageContent) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateMessageContent) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateMessageContent) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateMessageContent) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateMessageContent) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateMessageContent) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateMessageContent) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateMessageContent) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateMessageContent) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateMessageContent) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateMessageContent) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateMessageContent) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateMessageContent) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateMessageContent) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateMessageContent) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateMessageContent) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateMessageContent) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateMessageContent) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateMessageContent) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateMessageContent) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateMessageContent) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateMessageContent) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateMessageContent) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateMessageContent) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateMessageContent) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateMessageContent) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateMessageContent) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateMessageContent) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateMessageContent) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateMessageContent) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateMessageContent) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateMessageContent) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateMessageContent) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateMessageContent) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateMessageContent) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateMessageContent) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateMessageContent) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateMessageContent) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateMessageContent) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateMessageContent) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateMessageContent) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateMessageContent) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateMessageContent) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateMessageContent) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateMessageContent) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateMessageContent) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateMessageContent) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateMessageContent) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateMessageContent) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateMessageContent) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateMessageContent) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateMessageContent) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateMessageContent) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateMessageContent) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateMessageContent) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateMessageContent) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateMessageContent) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateMessageContent) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateMessageContent) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateMessageContent) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateMessageContent) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateMessageContent) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateMessageContent) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateMessageContent) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateMessageContent) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateMessageContent) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateMessageContent) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateMessageContent) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateMessageContent) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateMessageContent) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateMessageContent) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateMessageContent) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateMessageContent) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateMessageContent) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateMessageContent) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateMessageContent) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateMessageContent) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateMessageContent) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateMessageContent) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateMessageContent) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateMessageContent) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateMessageContent) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateMessageContent) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateMessageContent) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateMessageContent) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateMessageContent) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateMessageContent) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateMessageContent) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateMessageContent) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateMessageContent) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateMessageContent) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateMessageContent) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateMessageContent) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateMessageContent) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateMessageContent) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateMessageContent) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateMessageContent) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateMessageContent) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateMessageContent) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateMessageContent) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateMessageContent) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateMessageContent) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateMessageContent) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateMessageContent) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateMessageContent) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateMessageContent) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateMessageContent) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateMessageContent) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateMessageContent) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateMessageContent) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateMessageContent) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateMessageContent) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateMessageContent) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateMessageContent) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateMessageContent) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateMessageContent) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateMessageContent) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateMessageContent) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateMessageContent) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateMessageContent) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateMessageContent) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateMessageContent) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateMessageContent) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateMessageContent) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateMessageContent) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateMessageContent) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateMessageContent) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateMessageContent) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateMessageContent) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateMessageContent) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateMessageContent) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateMessageContent) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateMessageContent) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateMessageContent) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateMessageContent) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateMessageContent) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateMessageContent) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateMessageContent) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateMessageContent) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateMessageContent) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateMessageContent) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateMessageContent) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateMessageContent) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateMessageContent) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateMessageContent) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateMessageContent) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateMessageContent) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateMessageContent) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateMessageContent) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateMessageContent) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateMessageContent) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateMessageContent) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateMessageContent) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateMessageContent) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateMessageContent) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateMessageContent) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateMessageContent) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateMessageContent) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateMessageContent) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateMessageContent) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateMessageContent) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateMessageContent) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateMessageContent) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateMessageContent) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateMessageContent) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateMessageContent) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateMessageContent) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateMessageContent) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateMessageContent) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateMessageContent) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateMessageContent) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateMessageContent) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateMessageContent) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateMessageContent) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateMessageContent) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateMessageContent) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateMessageContent) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateMessageContent) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateMessageContent) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateMessageContent) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateMessageContent) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateMessageContent) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateMessageContent) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateMessageContent) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateMessageContent) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateMessageContent) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateMessageContent) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateMessageContent) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateMessageContent) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateMessageContent) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateMessageContent) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateMessageContent) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateMessageContent) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateMessageContent) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateMessageContent) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateMessageContent) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateMessageContent) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateMessageContent) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateMessageContent) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateMessageContent) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateMessageContent) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateMessageContent) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateMessageContent) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateMessageContent) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateMessageContent) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, u.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateMessageContent) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, u.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateMessageContent) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateMessageContent) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateMessageContent) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateMessageContent) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateMessageContent) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateMessageContent) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateMessageContent) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateMessageContent) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateMessageContent) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateMessageContent) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateMessageContent) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateMessageContent) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateMessageContent) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateMessageContent) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, u.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateMessageContent) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateMessageContent) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateMessageContent) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateMessageContent) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, u.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateMessageContent) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateMessageContent) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateMessageContentOpened) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateMessageContentOpened) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateMessageContentOpened) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateMessageContentOpened) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateMessageContentOpened) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateMessageContentOpened) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateMessageContentOpened) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateMessageContentOpened) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateMessageContentOpened) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateMessageContentOpened) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateMessageContentOpened) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateMessageContentOpened) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateMessageContentOpened) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateMessageContentOpened) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateMessageContentOpened) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateMessageContentOpened) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateMessageContentOpened) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateMessageContentOpened) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateMessageContentOpened) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateMessageContentOpened) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateMessageContentOpened) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateMessageContentOpened) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateMessageContentOpened) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateMessageContentOpened) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateMessageContentOpened) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateMessageContentOpened) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateMessageContentOpened) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateMessageContentOpened) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateMessageContentOpened) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateMessageContentOpened) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateMessageContentOpened) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateMessageContentOpened) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateMessageContentOpened) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateMessageContentOpened) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateMessageContentOpened) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateMessageContentOpened) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateMessageContentOpened) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateMessageContentOpened) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateMessageContentOpened) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateMessageContentOpened) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateMessageContentOpened) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateMessageContentOpened) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateMessageContentOpened) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateMessageContentOpened) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateMessageContentOpened) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateMessageContentOpened) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateMessageContentOpened) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateMessageContentOpened) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateMessageContentOpened) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateMessageContentOpened) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateMessageContentOpened) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateMessageContentOpened) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateMessageContentOpened) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateMessageContentOpened) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateMessageContentOpened) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateMessageContentOpened) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateMessageContentOpened) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateMessageContentOpened) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateMessageContentOpened) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateMessageContentOpened) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateMessageContentOpened) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateMessageContentOpened) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateMessageContentOpened) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateMessageContentOpened) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateMessageContentOpened) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateMessageContentOpened) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateMessageContentOpened) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateMessageContentOpened) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateMessageContentOpened) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateMessageContentOpened) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateMessageContentOpened) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateMessageContentOpened) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateMessageContentOpened) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateMessageContentOpened) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateMessageContentOpened) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateMessageContentOpened) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateMessageContentOpened) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateMessageContentOpened) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateMessageContentOpened) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateMessageContentOpened) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateMessageContentOpened) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateMessageContentOpened) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateMessageContentOpened) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateMessageContentOpened) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateMessageContentOpened) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateMessageContentOpened) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateMessageContentOpened) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateMessageContentOpened) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateMessageContentOpened) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateMessageContentOpened) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateMessageContentOpened) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateMessageContentOpened) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateMessageContentOpened) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateMessageContentOpened) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateMessageContentOpened) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateMessageContentOpened) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateMessageContentOpened) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateMessageContentOpened) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateMessageContentOpened) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateMessageContentOpened) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateMessageContentOpened) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateMessageContentOpened) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateMessageContentOpened) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateMessageContentOpened) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateMessageContentOpened) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateMessageContentOpened) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateMessageContentOpened) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateMessageContentOpened) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateMessageContentOpened) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateMessageContentOpened) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateMessageContentOpened) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateMessageContentOpened) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateMessageContentOpened) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateMessageContentOpened) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateMessageContentOpened) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateMessageContentOpened) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateMessageContentOpened) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateMessageContentOpened) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateMessageContentOpened) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateMessageContentOpened) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateMessageContentOpened) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateMessageContentOpened) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateMessageContentOpened) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateMessageContentOpened) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateMessageContentOpened) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateMessageContentOpened) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateMessageContentOpened) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateMessageContentOpened) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateMessageContentOpened) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateMessageContentOpened) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateMessageContentOpened) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateMessageContentOpened) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateMessageContentOpened) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateMessageContentOpened) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateMessageContentOpened) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateMessageContentOpened) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateMessageContentOpened) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateMessageContentOpened) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateMessageContentOpened) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateMessageContentOpened) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateMessageContentOpened) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateMessageContentOpened) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateMessageContentOpened) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateMessageContentOpened) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateMessageContentOpened) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateMessageContentOpened) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateMessageContentOpened) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateMessageContentOpened) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateMessageContentOpened) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateMessageContentOpened) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateMessageContentOpened) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateMessageContentOpened) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateMessageContentOpened) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateMessageContentOpened) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateMessageContentOpened) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateMessageContentOpened) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateMessageContentOpened) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateMessageContentOpened) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateMessageContentOpened) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateMessageContentOpened) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateMessageContentOpened) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateMessageContentOpened) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateMessageContentOpened) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateMessageContentOpened) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateMessageContentOpened) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateMessageContentOpened) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateMessageContentOpened) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateMessageContentOpened) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateMessageContentOpened) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateMessageContentOpened) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateMessageContentOpened) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateMessageContentOpened) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateMessageContentOpened) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateMessageContentOpened) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateMessageContentOpened) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateMessageContentOpened) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateMessageContentOpened) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateMessageContentOpened) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateMessageContentOpened) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateMessageContentOpened) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateMessageContentOpened) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateMessageContentOpened) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateMessageContentOpened) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateMessageContentOpened) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateMessageContentOpened) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateMessageContentOpened) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateMessageContentOpened) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateMessageContentOpened) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateMessageContentOpened) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateMessageContentOpened) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateMessageContentOpened) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateMessageContentOpened) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateMessageContentOpened) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateMessageContentOpened) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateMessageContentOpened) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateMessageContentOpened) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateMessageContentOpened) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateMessageContentOpened) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateMessageContentOpened) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateMessageContentOpened) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateMessageContentOpened) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateMessageContentOpened) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateMessageContentOpened) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateMessageContentOpened) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateMessageContentOpened) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateMessageContentOpened) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateMessageContentOpened) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateMessageContentOpened) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateMessageContentOpened) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateMessageContentOpened) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateMessageContentOpened) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateMessageContentOpened) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateMessageContentOpened) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateMessageContentOpened) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateMessageContentOpened) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateMessageContentOpened) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateMessageContentOpened) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateMessageContentOpened) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateMessageContentOpened) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateMessageContentOpened) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateMessageContentOpened) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateMessageContentOpened) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateMessageContentOpened) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateMessageContentOpened) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateMessageContentOpened) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateMessageContentOpened) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateMessageContentOpened) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateMessageContentOpened) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateMessageContentOpened) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateMessageContentOpened) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateMessageContentOpened) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateMessageContentOpened) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateMessageContentOpened) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateMessageContentOpened) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateMessageContentOpened) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, u.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateMessageContentOpened) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, u.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateMessageContentOpened) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateMessageContentOpened) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateMessageContentOpened) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateMessageContentOpened) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateMessageContentOpened) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateMessageContentOpened) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateMessageContentOpened) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateMessageContentOpened) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateMessageContentOpened) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateMessageContentOpened) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateMessageContentOpened) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateMessageContentOpened) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateMessageContentOpened) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateMessageContentOpened) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, u.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateMessageContentOpened) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateMessageContentOpened) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateMessageContentOpened) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateMessageContentOpened) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, u.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateMessageContentOpened) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateMessageContentOpened) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateMessageEdited) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateMessageEdited) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateMessageEdited) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateMessageEdited) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateMessageEdited) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateMessageEdited) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateMessageEdited) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateMessageEdited) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateMessageEdited) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateMessageEdited) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateMessageEdited) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateMessageEdited) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateMessageEdited) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateMessageEdited) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateMessageEdited) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateMessageEdited) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateMessageEdited) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateMessageEdited) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateMessageEdited) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateMessageEdited) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateMessageEdited) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateMessageEdited) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateMessageEdited) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateMessageEdited) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateMessageEdited) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateMessageEdited) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateMessageEdited) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateMessageEdited) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateMessageEdited) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateMessageEdited) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateMessageEdited) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateMessageEdited) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateMessageEdited) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateMessageEdited) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateMessageEdited) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateMessageEdited) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateMessageEdited) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateMessageEdited) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateMessageEdited) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateMessageEdited) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateMessageEdited) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateMessageEdited) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateMessageEdited) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateMessageEdited) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateMessageEdited) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateMessageEdited) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateMessageEdited) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateMessageEdited) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateMessageEdited) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateMessageEdited) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateMessageEdited) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateMessageEdited) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateMessageEdited) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateMessageEdited) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateMessageEdited) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateMessageEdited) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateMessageEdited) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateMessageEdited) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateMessageEdited) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateMessageEdited) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateMessageEdited) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateMessageEdited) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateMessageEdited) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateMessageEdited) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateMessageEdited) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateMessageEdited) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateMessageEdited) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateMessageEdited) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateMessageEdited) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateMessageEdited) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateMessageEdited) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateMessageEdited) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateMessageEdited) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateMessageEdited) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateMessageEdited) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateMessageEdited) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateMessageEdited) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateMessageEdited) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateMessageEdited) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateMessageEdited) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateMessageEdited) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateMessageEdited) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateMessageEdited) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateMessageEdited) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateMessageEdited) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateMessageEdited) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateMessageEdited) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateMessageEdited) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateMessageEdited) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateMessageEdited) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateMessageEdited) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateMessageEdited) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateMessageEdited) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateMessageEdited) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateMessageEdited) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateMessageEdited) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateMessageEdited) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateMessageEdited) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateMessageEdited) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateMessageEdited) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateMessageEdited) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateMessageEdited) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateMessageEdited) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateMessageEdited) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateMessageEdited) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateMessageEdited) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateMessageEdited) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateMessageEdited) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateMessageEdited) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateMessageEdited) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateMessageEdited) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateMessageEdited) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateMessageEdited) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateMessageEdited) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateMessageEdited) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateMessageEdited) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateMessageEdited) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateMessageEdited) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateMessageEdited) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateMessageEdited) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateMessageEdited) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateMessageEdited) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateMessageEdited) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateMessageEdited) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateMessageEdited) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateMessageEdited) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateMessageEdited) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateMessageEdited) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateMessageEdited) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateMessageEdited) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateMessageEdited) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateMessageEdited) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateMessageEdited) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateMessageEdited) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateMessageEdited) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateMessageEdited) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateMessageEdited) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateMessageEdited) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateMessageEdited) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateMessageEdited) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateMessageEdited) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateMessageEdited) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateMessageEdited) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateMessageEdited) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateMessageEdited) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateMessageEdited) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateMessageEdited) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateMessageEdited) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateMessageEdited) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateMessageEdited) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateMessageEdited) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateMessageEdited) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateMessageEdited) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateMessageEdited) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateMessageEdited) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateMessageEdited) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateMessageEdited) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateMessageEdited) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateMessageEdited) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateMessageEdited) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateMessageEdited) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateMessageEdited) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateMessageEdited) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateMessageEdited) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateMessageEdited) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateMessageEdited) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateMessageEdited) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateMessageEdited) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateMessageEdited) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateMessageEdited) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateMessageEdited) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateMessageEdited) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateMessageEdited) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateMessageEdited) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateMessageEdited) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateMessageEdited) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateMessageEdited) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateMessageEdited) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateMessageEdited) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateMessageEdited) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateMessageEdited) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateMessageEdited) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateMessageEdited) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateMessageEdited) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateMessageEdited) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateMessageEdited) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateMessageEdited) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateMessageEdited) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateMessageEdited) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateMessageEdited) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateMessageEdited) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateMessageEdited) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateMessageEdited) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateMessageEdited) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateMessageEdited) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateMessageEdited) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateMessageEdited) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateMessageEdited) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateMessageEdited) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateMessageEdited) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateMessageEdited) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateMessageEdited) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateMessageEdited) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateMessageEdited) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateMessageEdited) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateMessageEdited) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateMessageEdited) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateMessageEdited) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateMessageEdited) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateMessageEdited) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateMessageEdited) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateMessageEdited) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateMessageEdited) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateMessageEdited) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateMessageEdited) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateMessageEdited) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateMessageEdited) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateMessageEdited) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateMessageEdited) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateMessageEdited) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateMessageEdited) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateMessageEdited) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateMessageEdited) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateMessageEdited) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateMessageEdited) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateMessageEdited) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateMessageEdited) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateMessageEdited) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateMessageEdited) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateMessageEdited) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateMessageEdited) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateMessageEdited) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateMessageEdited) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateMessageEdited) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateMessageEdited) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, u.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateMessageEdited) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, u.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateMessageEdited) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateMessageEdited) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateMessageEdited) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateMessageEdited) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateMessageEdited) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateMessageEdited) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateMessageEdited) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateMessageEdited) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateMessageEdited) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateMessageEdited) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateMessageEdited) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateMessageEdited) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateMessageEdited) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateMessageEdited) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, u.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateMessageEdited) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateMessageEdited) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateMessageEdited) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateMessageEdited) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, u.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateMessageEdited) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateMessageEdited) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateMessageFactCheck) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateMessageFactCheck) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateMessageFactCheck) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateMessageFactCheck) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateMessageFactCheck) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateMessageFactCheck) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateMessageFactCheck) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateMessageFactCheck) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateMessageFactCheck) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateMessageFactCheck) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateMessageFactCheck) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateMessageFactCheck) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateMessageFactCheck) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateMessageFactCheck) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateMessageFactCheck) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateMessageFactCheck) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateMessageFactCheck) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateMessageFactCheck) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateMessageFactCheck) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateMessageFactCheck) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateMessageFactCheck) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateMessageFactCheck) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateMessageFactCheck) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateMessageFactCheck) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateMessageFactCheck) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateMessageFactCheck) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateMessageFactCheck) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateMessageFactCheck) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateMessageFactCheck) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateMessageFactCheck) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateMessageFactCheck) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateMessageFactCheck) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateMessageFactCheck) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateMessageFactCheck) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateMessageFactCheck) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateMessageFactCheck) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateMessageFactCheck) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateMessageFactCheck) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateMessageFactCheck) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateMessageFactCheck) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateMessageFactCheck) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateMessageFactCheck) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateMessageFactCheck) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateMessageFactCheck) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateMessageFactCheck) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateMessageFactCheck) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateMessageFactCheck) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateMessageFactCheck) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateMessageFactCheck) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateMessageFactCheck) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateMessageFactCheck) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateMessageFactCheck) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateMessageFactCheck) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateMessageFactCheck) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateMessageFactCheck) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateMessageFactCheck) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateMessageFactCheck) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateMessageFactCheck) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateMessageFactCheck) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateMessageFactCheck) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateMessageFactCheck) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateMessageFactCheck) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateMessageFactCheck) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateMessageFactCheck) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateMessageFactCheck) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateMessageFactCheck) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateMessageFactCheck) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateMessageFactCheck) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateMessageFactCheck) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateMessageFactCheck) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateMessageFactCheck) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateMessageFactCheck) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateMessageFactCheck) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateMessageFactCheck) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateMessageFactCheck) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateMessageFactCheck) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateMessageFactCheck) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateMessageFactCheck) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateMessageFactCheck) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateMessageFactCheck) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateMessageFactCheck) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateMessageFactCheck) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateMessageFactCheck) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateMessageFactCheck) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateMessageFactCheck) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateMessageFactCheck) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateMessageFactCheck) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateMessageFactCheck) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateMessageFactCheck) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateMessageFactCheck) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateMessageFactCheck) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateMessageFactCheck) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateMessageFactCheck) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateMessageFactCheck) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateMessageFactCheck) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateMessageFactCheck) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateMessageFactCheck) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateMessageFactCheck) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateMessageFactCheck) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateMessageFactCheck) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateMessageFactCheck) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateMessageFactCheck) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateMessageFactCheck) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateMessageFactCheck) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateMessageFactCheck) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateMessageFactCheck) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateMessageFactCheck) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateMessageFactCheck) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateMessageFactCheck) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateMessageFactCheck) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateMessageFactCheck) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateMessageFactCheck) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateMessageFactCheck) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateMessageFactCheck) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateMessageFactCheck) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateMessageFactCheck) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateMessageFactCheck) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateMessageFactCheck) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateMessageFactCheck) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateMessageFactCheck) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateMessageFactCheck) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateMessageFactCheck) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateMessageFactCheck) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateMessageFactCheck) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateMessageFactCheck) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateMessageFactCheck) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateMessageFactCheck) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateMessageFactCheck) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateMessageFactCheck) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateMessageFactCheck) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateMessageFactCheck) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateMessageFactCheck) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateMessageFactCheck) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateMessageFactCheck) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateMessageFactCheck) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateMessageFactCheck) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateMessageFactCheck) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateMessageFactCheck) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateMessageFactCheck) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateMessageFactCheck) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateMessageFactCheck) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateMessageFactCheck) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateMessageFactCheck) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateMessageFactCheck) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateMessageFactCheck) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateMessageFactCheck) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateMessageFactCheck) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateMessageFactCheck) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateMessageFactCheck) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateMessageFactCheck) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateMessageFactCheck) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateMessageFactCheck) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateMessageFactCheck) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateMessageFactCheck) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateMessageFactCheck) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateMessageFactCheck) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateMessageFactCheck) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateMessageFactCheck) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateMessageFactCheck) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateMessageFactCheck) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateMessageFactCheck) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateMessageFactCheck) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateMessageFactCheck) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateMessageFactCheck) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateMessageFactCheck) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateMessageFactCheck) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateMessageFactCheck) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateMessageFactCheck) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateMessageFactCheck) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateMessageFactCheck) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateMessageFactCheck) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateMessageFactCheck) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateMessageFactCheck) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateMessageFactCheck) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateMessageFactCheck) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateMessageFactCheck) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateMessageFactCheck) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateMessageFactCheck) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateMessageFactCheck) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateMessageFactCheck) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateMessageFactCheck) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateMessageFactCheck) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateMessageFactCheck) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateMessageFactCheck) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateMessageFactCheck) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateMessageFactCheck) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateMessageFactCheck) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateMessageFactCheck) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateMessageFactCheck) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateMessageFactCheck) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateMessageFactCheck) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateMessageFactCheck) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateMessageFactCheck) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateMessageFactCheck) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateMessageFactCheck) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateMessageFactCheck) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateMessageFactCheck) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateMessageFactCheck) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateMessageFactCheck) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateMessageFactCheck) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateMessageFactCheck) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateMessageFactCheck) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateMessageFactCheck) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateMessageFactCheck) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateMessageFactCheck) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateMessageFactCheck) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateMessageFactCheck) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateMessageFactCheck) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateMessageFactCheck) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateMessageFactCheck) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateMessageFactCheck) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateMessageFactCheck) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateMessageFactCheck) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateMessageFactCheck) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateMessageFactCheck) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateMessageFactCheck) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateMessageFactCheck) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateMessageFactCheck) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateMessageFactCheck) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateMessageFactCheck) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateMessageFactCheck) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateMessageFactCheck) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateMessageFactCheck) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateMessageFactCheck) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateMessageFactCheck) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateMessageFactCheck) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateMessageFactCheck) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateMessageFactCheck) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateMessageFactCheck) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateMessageFactCheck) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateMessageFactCheck) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateMessageFactCheck) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}
