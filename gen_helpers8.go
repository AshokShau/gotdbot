package gotdbot

// CanPostStory is a helper method for Client.CanPostStory
func (m MessageThreadInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m MessageThreadInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessageThreadInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessageThreadInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m MessageThreadInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessageThreadInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageThreadInfo) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageThreadInfo) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageThreadInfo) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageThreadInfo) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessageThreadInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageThreadInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m MessageThreadInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m MessageThreadInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m MessageThreadInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m MessageThreadInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (m MessageThreadInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageThreadInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessageThreadInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m MessageThreadInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m MessageThreadInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageThreadInfo) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (m MessageThreadInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m MessageThreadInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m MessageThreadInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageThreadInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessageThreadInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageThreadInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessageThreadInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessageThreadInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessageThreadInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageThreadInfo) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageThreadInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageThreadInfo) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageThreadInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessageThreadInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageThreadInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessageThreadInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessageThreadInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessageThreadInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessageThreadInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageThreadInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m MessageThreadInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessageThreadInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessageThreadInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m MessageThreadInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m MessageThreadInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m MessageThreadInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m MessageThreadInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m MessageThreadInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m MessageThreadInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m MessageThreadInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m MessageThreadInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m MessageThreadInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageThreadInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m MessageThreadInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m MessageThreadInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m MessageThreadInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m MessageThreadInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m MessageThreadInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m MessageThreadInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m MessageThreadInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m MessageThreadInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageThreadInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (m MessageThreadInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (m MessageThreadInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessageThreadInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (m MessageThreadInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m MessageThreadInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m MessageThreadInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m MessageThreadInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m MessageThreadInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (m MessageThreadInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m MessageThreadInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m MessageThreadInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageThreadInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (m MessageThreadInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m MessageThreadInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m MessageThreadInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m MessageThreadInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m MessageThreadInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageThreadInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m MessageThreadInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageThreadInfo) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageThreadInfo) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageThreadInfo) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m MessageThreadInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageThreadInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessageThreadInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m MessageThreadInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m MessageThreadInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessageThreadInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessageThreadInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageThreadInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageThreadInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessageThreadInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(m.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessageThreadInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessageThreadInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessageThreadInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessageThreadInfo) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m MessageThreadInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessageThreadInfo) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessageThreadInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessageThreadInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessageThreadInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessageThreadInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (m MessageThreadInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessageThreadInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessageThreadInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessageThreadInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessageThreadInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessageThreadInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessageThreadInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessageThreadInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m MessageThreadInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m MessageThreadInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m MessageThreadInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageThreadInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageThreadInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m MessageThreadInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m MessageThreadInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessageThreadInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageThreadInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (m MessageThreadInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m MessageThreadInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m MessageThreadInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m MessageThreadInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageThreadInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m MessageThreadInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m MessageThreadInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessageThreadInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MessageThreadInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageThreadInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageThreadInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageThreadInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m MessageThreadInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessageThreadInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m MessageThreadInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m MessageThreadInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m MessageThreadInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageThreadInfo) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageThreadInfo) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessageThreadInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessageThreadInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m MessageThreadInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m MessageThreadInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessageThreadInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessageThreadInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m MessageThreadInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m MessageThreadInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m MessageThreadInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m MessageThreadInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m MessageThreadInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m MessageThreadInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m MessageThreadInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m MessageThreadInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageThreadInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageThreadInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessageThreadInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageThreadInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageThreadInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m MessageThreadInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m MessageThreadInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m MessageThreadInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (m MessageThreadInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (m MessageThreadInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (m MessageThreadInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (m MessageThreadInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageThreadInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageThreadInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m MessageThreadInfo) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (m MessageThreadInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (m MessageThreadInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (m MessageThreadInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageThreadInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageThreadInfo) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessageThreadInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m MessageThreadInfo) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m MessageThreadInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m MessageThreadInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m MessageThreadInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageThreadInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m MessageThreadInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m MessageThreadInfo) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageThreadInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m MessageThreadInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (m MessageThreadInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m MessageThreadInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m MessageThreadInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m MessageThreadInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageThreadInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (m MessageThreadInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m MessageThreadInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageThreadInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m MessageThreadInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m MessageThreadInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m MessageThreadInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m MessageThreadInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m MessageThreadInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m MessageThreadInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageThreadInfo) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m MessageThreadInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageThreadInfo) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageThreadInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessageThreadInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessageThreadInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessageThreadInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m MessageThreadInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m MessageThreadInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessageThreadInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageThreadInfo) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m MessageThreadInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageThreadInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageThreadInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageThreadInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessageThreadInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessageThreadInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessageThreadInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m MessageThreadInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m MessageThreadInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m MessageThreadInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m MessageThreadInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m MessageThreadInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m MessageThreadInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m MessageThreadInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m MessageThreadInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m MessageThreadInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageThreadInfo) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageThreadInfo) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageThreadInfo) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageThreadInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessageThreadInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (m MessageThreadInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m MessageThreadInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageThreadInfo) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessageThreadInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m MessageThreadInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (m MessageThreadInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageTopicForum) DeleteForumTopic(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteForumTopic(chatId, m.ForumTopicId)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageTopicForum) EditForumTopic(client *Client, chatId int64, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, m.ForumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageTopicForum) GetForumTopic(client *Client, chatId int64) (*ForumTopic, error) {
	return client.GetForumTopic(chatId, m.ForumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageTopicForum) GetForumTopicHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(chatId, m.ForumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageTopicForum) GetForumTopicLink(client *Client, chatId int64) (*MessageLink, error) {
	return client.GetForumTopicLink(chatId, m.ForumTopicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageTopicForum) ReadAllForumTopicMentions(client *Client, chatId int64) (*Ok, error) {
	return client.ReadAllForumTopicMentions(chatId, m.ForumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageTopicForum) ReadAllForumTopicReactions(client *Client, chatId int64) (*Ok, error) {
	return client.ReadAllForumTopicReactions(chatId, m.ForumTopicId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageTopicForum) SendTextMessageDraft(client *Client, chatId int64, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, m.ForumTopicId, draftId, text)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageTopicForum) SetForumTopicNotificationSettings(client *Client, chatId int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(chatId, m.ForumTopicId, notificationSettings)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageTopicForum) ToggleForumTopicIsClosed(client *Client, chatId int64, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(chatId, m.ForumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageTopicForum) ToggleForumTopicIsPinned(client *Client, chatId int64, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(chatId, m.ForumTopicId, isPinned)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageTopicForum) UnpinAllForumTopicMessages(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(chatId, m.ForumTopicId)
}

// DeleteSavedMessagesTopicHistory is a helper method for Client.DeleteSavedMessagesTopicHistory
func (m MessageTopicSavedMessages) DeleteSavedMessagesTopicHistory(client *Client) (*Ok, error) {
	return client.DeleteSavedMessagesTopicHistory(m.SavedMessagesTopicId)
}

// DeleteSavedMessagesTopicMessagesByDate is a helper method for Client.DeleteSavedMessagesTopicMessagesByDate
func (m MessageTopicSavedMessages) DeleteSavedMessagesTopicMessagesByDate(client *Client, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteSavedMessagesTopicMessagesByDate(m.SavedMessagesTopicId, minDate, maxDate)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageTopicSavedMessages) GetChatSparseMessagePositions(client *Client, chatId int64, filter *SearchMessagesFilter, fromMessageId int64, limit int32) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(chatId, filter, fromMessageId, limit, m.SavedMessagesTopicId)
}

// GetSavedMessagesTags is a helper method for Client.GetSavedMessagesTags
func (m MessageTopicSavedMessages) GetSavedMessagesTags(client *Client) (*SavedMessagesTags, error) {
	return client.GetSavedMessagesTags(m.SavedMessagesTopicId)
}

// GetSavedMessagesTopicHistory is a helper method for Client.GetSavedMessagesTopicHistory
func (m MessageTopicSavedMessages) GetSavedMessagesTopicHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetSavedMessagesTopicHistory(m.SavedMessagesTopicId, fromMessageId, offset, limit)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (m MessageTopicSavedMessages) GetSavedMessagesTopicMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(m.SavedMessagesTopicId, date)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (m MessageTopicSavedMessages) SearchSavedMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(m.SavedMessagesTopicId, query, fromMessageId, offset, limit, opts)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (m MessageTopicSavedMessages) ToggleSavedMessagesTopicIsPinned(client *Client, isPinned bool) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(m.SavedMessagesTopicId, isPinned)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageUpgradedGift) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, m.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageUpgradedGift) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, m.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageUpgradedGift) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, m.SenderId, reportSpam)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessageUpgradedGift) DropGiftOriginalDetails(client *Client, starCount int64) (*Ok, error) {
	return client.DropGiftOriginalDetails(m.ReceivedGiftId, starCount)
}

// GetReceivedGift is a helper method for Client.GetReceivedGift
func (m MessageUpgradedGift) GetReceivedGift(client *Client) (*ReceivedGift, error) {
	return client.GetReceivedGift(m.ReceivedGiftId)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (m MessageUpgradedGift) GetUpgradedGiftWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(m.ReceivedGiftId, password)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageUpgradedGift) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, m.SenderId)
}

// SellGift is a helper method for Client.SellGift
func (m MessageUpgradedGift) SellGift(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SellGift(businessConnectionId, m.ReceivedGiftId)
}

// SetGiftResalePrice is a helper method for Client.SetGiftResalePrice
func (m MessageUpgradedGift) SetGiftResalePrice(client *Client, opts *SetGiftResalePriceOpts) (*Ok, error) {
	return client.SetGiftResalePrice(m.ReceivedGiftId, opts)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (m MessageUpgradedGift) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(m.SenderId, opts)
}

// ToggleGiftIsSaved is a helper method for Client.ToggleGiftIsSaved
func (m MessageUpgradedGift) ToggleGiftIsSaved(client *Client) (*Ok, error) {
	return client.ToggleGiftIsSaved(m.ReceivedGiftId, m.IsSaved)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessageUpgradedGift) TransferGift(client *Client, businessConnectionId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	return client.TransferGift(businessConnectionId, m.ReceivedGiftId, newOwnerId, starCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessageUpgradedGift) UpgradeGift(client *Client, businessConnectionId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, m.ReceivedGiftId, keepOriginalDetails, starCount)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageUpgradedGiftPurchaseOffer) CreateChatInviteLink(client *Client, chatId int64, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, m.ExpirationDate, memberLimit, createsJoinRequest)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageUpgradedGiftPurchaseOffer) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, m.ExpirationDate, memberLimit, createsJoinRequest)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageUpgradedGiftPurchaseOffer) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, m.Price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (m MessageUpgradedGiftPurchaseOffer) SendResoldGift(client *Client, giftName string, ownerId *MessageSender) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, ownerId, m.Price)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageUpgradedGiftPurchaseOfferRejected) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, m.Price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (m MessageUpgradedGiftPurchaseOfferRejected) SendResoldGift(client *Client, giftName string, ownerId *MessageSender) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, ownerId, m.Price)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageUsersShared) ShareChatWithBot(client *Client, chatId int64, messageId int64, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, messageId, m.ButtonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageUsersShared) ShareUsersWithBot(client *Client, chatId int64, messageId int64, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, messageId, m.ButtonId, sharedUserIds, onlyCheck)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageVideo) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageVideo) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (m MessageVideo) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, m.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageVideo) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// DiscardCall is a helper method for Client.DiscardCall
func (m MessageVideoChatEnded) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, m.Duration, isVideo, connectionId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageVideoChatEnded) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, m.Duration, paidMessageStarCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageVideoChatScheduled) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(m.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (m MessageVideoChatScheduled) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(m.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (m MessageVideoChatScheduled) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(m.GroupCallId)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageVideoChatScheduled) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, m.StartDate, isRtmpStream)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (m MessageVideoChatScheduled) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(m.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (m MessageVideoChatScheduled) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(m.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageVideoChatScheduled) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(m.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (m MessageVideoChatScheduled) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(m.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (m MessageVideoChatScheduled) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(m.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (m MessageVideoChatScheduled) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(m.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (m MessageVideoChatScheduled) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(m.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (m MessageVideoChatScheduled) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(m.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (m MessageVideoChatScheduled) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(m.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (m MessageVideoChatScheduled) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(m.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (m MessageVideoChatScheduled) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(m.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (m MessageVideoChatScheduled) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(m.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (m MessageVideoChatScheduled) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(m.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (m MessageVideoChatScheduled) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(m.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageVideoChatScheduled) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(m.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (m MessageVideoChatScheduled) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(m.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (m MessageVideoChatScheduled) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(m.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (m MessageVideoChatScheduled) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(m.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (m MessageVideoChatScheduled) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(m.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (m MessageVideoChatScheduled) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(m.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (m MessageVideoChatScheduled) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(m.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (m MessageVideoChatScheduled) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(m.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageVideoChatScheduled) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(m.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessageVideoChatScheduled) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(m.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (m MessageVideoChatScheduled) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(m.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (m MessageVideoChatScheduled) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(m.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (m MessageVideoChatScheduled) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(m.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageVideoChatScheduled) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(m.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageVideoChatScheduled) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(m.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (m MessageVideoChatScheduled) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(m.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (m MessageVideoChatScheduled) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(m.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (m MessageVideoChatScheduled) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(m.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (m MessageVideoChatScheduled) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(m.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (m MessageVideoChatScheduled) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(m.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (m MessageVideoChatScheduled) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(m.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (m MessageVideoChatScheduled) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(m.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (m MessageVideoChatScheduled) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(m.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (m MessageVideoChatScheduled) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(m.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (m MessageVideoChatScheduled) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(m.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageVideoChatStarted) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(m.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (m MessageVideoChatStarted) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(m.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (m MessageVideoChatStarted) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(m.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (m MessageVideoChatStarted) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(m.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (m MessageVideoChatStarted) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(m.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageVideoChatStarted) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(m.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (m MessageVideoChatStarted) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(m.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (m MessageVideoChatStarted) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(m.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (m MessageVideoChatStarted) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(m.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (m MessageVideoChatStarted) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(m.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (m MessageVideoChatStarted) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(m.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (m MessageVideoChatStarted) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(m.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (m MessageVideoChatStarted) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(m.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (m MessageVideoChatStarted) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(m.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (m MessageVideoChatStarted) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(m.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (m MessageVideoChatStarted) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(m.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (m MessageVideoChatStarted) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(m.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageVideoChatStarted) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(m.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (m MessageVideoChatStarted) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(m.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (m MessageVideoChatStarted) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(m.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (m MessageVideoChatStarted) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(m.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (m MessageVideoChatStarted) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(m.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (m MessageVideoChatStarted) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(m.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (m MessageVideoChatStarted) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(m.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (m MessageVideoChatStarted) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(m.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageVideoChatStarted) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(m.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessageVideoChatStarted) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(m.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (m MessageVideoChatStarted) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(m.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (m MessageVideoChatStarted) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(m.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (m MessageVideoChatStarted) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(m.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageVideoChatStarted) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(m.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageVideoChatStarted) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(m.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (m MessageVideoChatStarted) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(m.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (m MessageVideoChatStarted) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(m.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (m MessageVideoChatStarted) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(m.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (m MessageVideoChatStarted) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(m.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (m MessageVideoChatStarted) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(m.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (m MessageVideoChatStarted) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(m.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (m MessageVideoChatStarted) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(m.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (m MessageVideoChatStarted) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(m.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (m MessageVideoChatStarted) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(m.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (m MessageVideoChatStarted) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(m.GroupCallId, muteNewParticipants)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageViewer) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, m.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (m MessageViewer) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(m.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageViewer) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(m.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (m MessageViewer) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(m.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (m MessageViewer) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(m.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (m MessageViewer) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(m.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (m MessageViewer) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(m.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageViewer) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(m.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (m MessageViewer) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(m.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (m MessageViewer) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(m.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageViewer) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, m.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (m MessageViewer) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(m.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (m MessageViewer) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, m.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (m MessageViewer) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(m.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (m MessageViewer) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(m.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (m MessageViewer) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(m.UserId)
}

// GetUser is a helper method for Client.GetUser
func (m MessageViewer) GetUser(client *Client) (*User, error) {
	return client.GetUser(m.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageViewer) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, m.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (m MessageViewer) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(m.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (m MessageViewer) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(m.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (m MessageViewer) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(m.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (m MessageViewer) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(m.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageViewer) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(m.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageViewer) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, m.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageViewer) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, m.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageViewer) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, m.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (m MessageViewer) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(m.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageViewer) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(m.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (m MessageViewer) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(m.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageViewer) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, m.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (m MessageViewer) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, m.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (m MessageViewer) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(m.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (m MessageViewer) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(m.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageViewer) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(m.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (m MessageViewer) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(m.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (m MessageViewer) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(m.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (m MessageViewer) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(m.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (m MessageViewer) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(m.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (m MessageViewer) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(m.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (m MessageViewer) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(m.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (m MessageViewer) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(m.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageViewer) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, m.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (m MessageViewer) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(m.UserId, stickerFormat, sticker)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageVoiceNote) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// AnswerCustomQuery is a helper method for Client.AnswerCustomQuery
func (m MessageWebAppDataReceived) AnswerCustomQuery(client *Client, customQueryId string) (*Ok, error) {
	return client.AnswerCustomQuery(customQueryId, m.Data)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (m MessageWebAppDataReceived) SendWebAppData(client *Client, botUserId int64) (*Ok, error) {
	return client.SendWebAppData(botUserId, m.ButtonText, m.Data)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (m MessageWebAppDataSent) SendWebAppData(client *Client, botUserId int64, data string) (*Ok, error) {
	return client.SendWebAppData(botUserId, m.ButtonText, data)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (m Minithumbnail) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, m.Data, opts)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (m Minithumbnail) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, m.Data, unencryptedPrefixSize)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m Minithumbnail) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, m.Width, m.Height, scale, chatId)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (m Minithumbnail) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, m.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (m Minithumbnail) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, m.Data)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (n Notification) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, n.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (n Notification) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, n.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (n Notification) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, n.Date)
}

// Remove is a helper method for Client.RemoveNotification
func (n Notification) Remove(client *Client, notificationGroupId int32) (*Ok, error) {
	return client.RemoveNotification(notificationGroupId, n.Id)
}

// AddChatMember is a helper method for Client.AddChatMember
func (n NotificationGroup) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(n.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (n NotificationGroup) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(n.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (n NotificationGroup) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(n.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (n NotificationGroup) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(n.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (n NotificationGroup) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, n.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (n NotificationGroup) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(n.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (n NotificationGroup) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(n.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (n NotificationGroup) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(n.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (n NotificationGroup) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(n.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (n NotificationGroup) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(n.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (n NotificationGroup) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(n.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (n NotificationGroup) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(n.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (n NotificationGroup) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(n.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (n NotificationGroup) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(n.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (n NotificationGroup) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(n.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (n NotificationGroup) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(n.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (n NotificationGroup) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(n.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (n NotificationGroup) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(n.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (n NotificationGroup) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(n.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (n NotificationGroup) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(n.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (n NotificationGroup) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(n.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (n NotificationGroup) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(n.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (n NotificationGroup) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(n.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (n NotificationGroup) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(n.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (n NotificationGroup) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(n.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (n NotificationGroup) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(n.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (n NotificationGroup) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(n.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (n NotificationGroup) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(n.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (n NotificationGroup) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(n.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (n NotificationGroup) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(n.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (n NotificationGroup) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(n.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (n NotificationGroup) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(n.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (n NotificationGroup) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(n.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (n NotificationGroup) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(n.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (n NotificationGroup) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(n.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (n NotificationGroup) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(n.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (n NotificationGroup) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(n.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (n NotificationGroup) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(n.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (n NotificationGroup) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(n.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (n NotificationGroup) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, n.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (n NotificationGroup) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, n.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (n NotificationGroup) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, n.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (n NotificationGroup) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, n.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (n NotificationGroup) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, n.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (n NotificationGroup) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, n.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (n NotificationGroup) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(n.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (n NotificationGroup) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(n.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (n NotificationGroup) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(n.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (n NotificationGroup) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(n.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (n NotificationGroup) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(n.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (n NotificationGroup) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(n.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (n NotificationGroup) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(n.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (n NotificationGroup) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(n.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (n NotificationGroup) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(n.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (n NotificationGroup) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(n.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (n NotificationGroup) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(n.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (n NotificationGroup) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, n.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (n NotificationGroup) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(n.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (n NotificationGroup) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(n.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (n NotificationGroup) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(n.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (n NotificationGroup) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(n.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (n NotificationGroup) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(n.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (n NotificationGroup) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(n.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (n NotificationGroup) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(n.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (n NotificationGroup) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(n.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (n NotificationGroup) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(n.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (n NotificationGroup) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(n.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (n NotificationGroup) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(n.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (n NotificationGroup) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(n.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (n NotificationGroup) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(n.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (n NotificationGroup) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(n.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (n NotificationGroup) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(n.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (n NotificationGroup) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(n.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (n NotificationGroup) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(n.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (n NotificationGroup) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(n.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (n NotificationGroup) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(n.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (n NotificationGroup) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(n.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (n NotificationGroup) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(n.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (n NotificationGroup) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(n.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (n NotificationGroup) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(n.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (n NotificationGroup) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(n.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (n NotificationGroup) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(n.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (n NotificationGroup) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(n.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (n NotificationGroup) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(n.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (n NotificationGroup) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(n.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (n NotificationGroup) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(n.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (n NotificationGroup) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(n.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (n NotificationGroup) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(n.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (n NotificationGroup) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(n.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (n NotificationGroup) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(n.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (n NotificationGroup) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(n.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (n NotificationGroup) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(n.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (n NotificationGroup) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(n.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (n NotificationGroup) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(n.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (n NotificationGroup) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(n.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (n NotificationGroup) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(n.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (n NotificationGroup) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(n.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (n NotificationGroup) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(n.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (n NotificationGroup) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(n.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (n NotificationGroup) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(n.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (n NotificationGroup) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(n.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (n NotificationGroup) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(n.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (n NotificationGroup) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(n.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (n NotificationGroup) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, n.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (n NotificationGroup) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(n.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (n NotificationGroup) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(n.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (n NotificationGroup) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(n.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (n NotificationGroup) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(n.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (n NotificationGroup) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, n.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (n NotificationGroup) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(n.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (n NotificationGroup) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(n.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (n NotificationGroup) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(n.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (n NotificationGroup) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(n.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (n NotificationGroup) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(n.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (n NotificationGroup) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(n.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (n NotificationGroup) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(n.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (n NotificationGroup) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(n.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (n NotificationGroup) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(n.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (n NotificationGroup) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(n.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (n NotificationGroup) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(n.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (n NotificationGroup) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(n.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (n NotificationGroup) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(n.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (n NotificationGroup) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(n.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (n NotificationGroup) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(n.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (n NotificationGroup) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(n.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (n NotificationGroup) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(n.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (n NotificationGroup) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(n.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (n NotificationGroup) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(n.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (n NotificationGroup) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(n.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (n NotificationGroup) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, n.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (n NotificationGroup) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(n.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (n NotificationGroup) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(n.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (n NotificationGroup) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(n.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (n NotificationGroup) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(n.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (n NotificationGroup) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(n.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (n NotificationGroup) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(n.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (n NotificationGroup) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(n.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (n NotificationGroup) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(n.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (n NotificationGroup) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(n.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (n NotificationGroup) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(n.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (n NotificationGroup) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(n.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (n NotificationGroup) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(n.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (n NotificationGroup) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(n.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (n NotificationGroup) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(n.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (n NotificationGroup) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(n.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (n NotificationGroup) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(n.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (n NotificationGroup) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(n.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (n NotificationGroup) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(n.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (n NotificationGroup) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(n.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (n NotificationGroup) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(n.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (n NotificationGroup) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(n.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (n NotificationGroup) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(n.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (n NotificationGroup) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(n.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (n NotificationGroup) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(n.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (n NotificationGroup) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(n.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (n NotificationGroup) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(n.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (n NotificationGroup) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, n.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (n NotificationGroup) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(n.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (n NotificationGroup) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(n.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (n NotificationGroup) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(n.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (n NotificationGroup) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(n.ChatId, messageId, reactionType)
}

// RemoveNotification is a helper method for Client.RemoveNotification
func (n NotificationGroup) RemoveNotification(client *Client, notificationId int32) (*Ok, error) {
	return client.RemoveNotification(n.Id, notificationId)
}

// Remove is a helper method for Client.RemoveNotificationGroup
func (n NotificationGroup) Remove(client *Client, maxNotificationId int32) (*Ok, error) {
	return client.RemoveNotificationGroup(n.Id, maxNotificationId)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (n NotificationGroup) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(n.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (n NotificationGroup) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(n.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (n NotificationGroup) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(n.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (n NotificationGroup) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, n.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (n NotificationGroup) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(n.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (n NotificationGroup) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(n.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (n NotificationGroup) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(n.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (n NotificationGroup) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(n.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (n NotificationGroup) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(n.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (n NotificationGroup) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(n.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (n NotificationGroup) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(n.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (n NotificationGroup) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(n.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (n NotificationGroup) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(n.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (n NotificationGroup) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(n.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (n NotificationGroup) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(n.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (n NotificationGroup) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, n.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (n NotificationGroup) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(n.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (n NotificationGroup) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(n.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (n NotificationGroup) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(n.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (n NotificationGroup) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(n.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (n NotificationGroup) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, n.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (n NotificationGroup) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, n.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (n NotificationGroup) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, n.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (n NotificationGroup) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(n.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (n NotificationGroup) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(n.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (n NotificationGroup) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(n.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (n NotificationGroup) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(n.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (n NotificationGroup) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(n.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (n NotificationGroup) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(n.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (n NotificationGroup) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, n.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (n NotificationGroup) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(n.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (n NotificationGroup) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(n.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (n NotificationGroup) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(n.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (n NotificationGroup) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(n.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (n NotificationGroup) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(n.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (n NotificationGroup) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(n.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (n NotificationGroup) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(n.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (n NotificationGroup) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(n.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (n NotificationGroup) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(n.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (n NotificationGroup) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(n.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (n NotificationGroup) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(n.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (n NotificationGroup) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(n.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (n NotificationGroup) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(n.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (n NotificationGroup) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(n.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (n NotificationGroup) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(n.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (n NotificationGroup) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(n.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (n NotificationGroup) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(n.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (n NotificationGroup) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(n.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (n NotificationGroup) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(n.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (n NotificationGroup) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(n.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (n NotificationGroup) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(n.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (n NotificationGroup) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(n.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (n NotificationGroup) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(n.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (n NotificationGroup) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(n.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (n NotificationGroup) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(n.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (n NotificationGroup) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(n.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (n NotificationGroup) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(n.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (n NotificationGroup) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(n.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (n NotificationGroup) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(n.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (n NotificationGroup) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(n.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (n NotificationGroup) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(n.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (n NotificationGroup) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(n.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (n NotificationGroup) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(n.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (n NotificationGroup) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(n.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (n NotificationGroup) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(n.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (n NotificationGroup) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(n.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (n NotificationGroup) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(n.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (n NotificationGroup) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(n.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (n NotificationGroup) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, n.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (n NotificationGroup) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(n.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (n NotificationGroup) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(n.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (n NotificationGroup) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(n.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (n NotificationGroup) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(n.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (n NotificationGroup) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(n.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (n NotificationGroup) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(n.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (n NotificationGroup) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(n.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (n NotificationGroup) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, n.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (n NotificationGroup) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(n.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (n NotificationGroup) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(n.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (n NotificationGroup) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(n.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (n NotificationGroup) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(n.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (n NotificationGroup) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(n.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (n NotificationGroup) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(n.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (n NotificationGroup) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(n.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (n NotificationGroup) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(n.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (n NotificationGroup) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(n.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (n NotificationGroup) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(n.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (n NotificationGroup) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(n.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (n NotificationGroup) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(n.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (n NotificationGroup) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(n.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (n NotificationGroup) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(n.ChatId, messageIds, forceRead, opts)
}

// AnswerCustomQuery is a helper method for Client.AnswerCustomQuery
func (n NotificationSound) AnswerCustomQuery(client *Client, customQueryId string) (*Ok, error) {
	return client.AnswerCustomQuery(customQueryId, n.Data)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (n NotificationSound) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, n.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (n NotificationSound) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, n.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (n NotificationSound) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(n.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (n NotificationSound) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, n.Title, startDate, isRtmpStream)
}

// DiscardCall is a helper method for Client.DiscardCall
func (n NotificationSound) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, n.Duration, isVideo, connectionId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (n NotificationSound) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, n.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (n NotificationSound) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, n.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (n NotificationSound) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, n.Date)
}

// GetSaved is a helper method for Client.GetSavedNotificationSound
func (n NotificationSound) GetSaved(client *Client) (*NotificationSounds, error) {
	return client.GetSavedNotificationSound(n.Id)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (n NotificationSound) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(n.Title)
}

// RemoveSaved is a helper method for Client.RemoveSavedNotificationSound
func (n NotificationSound) RemoveSaved(client *Client) (*Ok, error) {
	return client.RemoveSavedNotificationSound(n.Id)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (n NotificationSound) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, n.Duration, paidMessageStarCount)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (n NotificationSound) SendWebAppData(client *Client, botUserId int64, buttonText string) (*Ok, error) {
	return client.SendWebAppData(botUserId, buttonText, n.Data)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (n NotificationSound) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, n.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (n NotificationSound) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, n.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (n NotificationSound) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, n.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (n NotificationSound) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, n.Title, recordVideo, usePortraitOrientation)
}

// AcceptCall is a helper method for Client.AcceptCall
func (n NotificationTypeNewCall) AcceptCall(client *Client, protocol *CallProtocol) (*Ok, error) {
	return client.AcceptCall(n.CallId, protocol)
}

// DiscardCall is a helper method for Client.DiscardCall
func (n NotificationTypeNewCall) DiscardCall(client *Client, isDisconnected bool, inviteLink string, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(n.CallId, isDisconnected, inviteLink, duration, isVideo, connectionId)
}

// SendCallDebugInformation is a helper method for Client.SendCallDebugInformation
func (n NotificationTypeNewCall) SendCallDebugInformation(client *Client, debugInformation string) (*Ok, error) {
	return client.SendCallDebugInformation(n.CallId, debugInformation)
}

// SendCallLog is a helper method for Client.SendCallLog
func (n NotificationTypeNewCall) SendCallLog(client *Client, logFile *InputFile) (*Ok, error) {
	return client.SendCallLog(n.CallId, logFile)
}

// SendCallRating is a helper method for Client.SendCallRating
func (n NotificationTypeNewCall) SendCallRating(client *Client, rating int32, comment string, problems []*CallProblem) (*Ok, error) {
	return client.SendCallRating(n.CallId, rating, comment, problems)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (n NotificationTypeNewCall) SendCallSignalingData(client *Client, data string) (*Ok, error) {
	return client.SendCallSignalingData(n.CallId, data)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (n NotificationTypeNewPushMessage) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, n.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (n NotificationTypeNewPushMessage) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, n.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (n NotificationTypeNewPushMessage) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, n.SenderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (n NotificationTypeNewPushMessage) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, n.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (n NotificationTypeNewPushMessage) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, n.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (n NotificationTypeNewPushMessage) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, n.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (n NotificationTypeNewPushMessage) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, n.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (n NotificationTypeNewPushMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(n.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (n NotificationTypeNewPushMessage) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, n.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (n NotificationTypeNewPushMessage) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, n.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (n NotificationTypeNewPushMessage) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, n.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (n NotificationTypeNewPushMessage) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, n.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (n NotificationTypeNewPushMessage) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, n.MessageId, comment)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (n NotificationTypeNewPushMessage) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, n.SenderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (n NotificationTypeNewPushMessage) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, n.MessageId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (n NotificationTypeNewPushMessage) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, n.SenderId, reportSpam)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (n NotificationTypeNewPushMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, n.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (n NotificationTypeNewPushMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, n.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (n NotificationTypeNewPushMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, n.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (n NotificationTypeNewPushMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, n.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (n NotificationTypeNewPushMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, n.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (n NotificationTypeNewPushMessage) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, n.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (n NotificationTypeNewPushMessage) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, n.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (n NotificationTypeNewPushMessage) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, n.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (n NotificationTypeNewPushMessage) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, n.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (n NotificationTypeNewPushMessage) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, n.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (n NotificationTypeNewPushMessage) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, n.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (n NotificationTypeNewPushMessage) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, n.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (n NotificationTypeNewPushMessage) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, n.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (n NotificationTypeNewPushMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, n.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (n NotificationTypeNewPushMessage) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, n.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (n NotificationTypeNewPushMessage) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, n.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (n NotificationTypeNewPushMessage) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, n.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (n NotificationTypeNewPushMessage) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, n.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (n NotificationTypeNewPushMessage) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, n.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (n NotificationTypeNewPushMessage) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, n.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (n NotificationTypeNewPushMessage) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, n.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (n NotificationTypeNewPushMessage) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, n.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (n NotificationTypeNewPushMessage) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, n.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (n NotificationTypeNewPushMessage) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, n.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (n NotificationTypeNewPushMessage) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, n.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (n NotificationTypeNewPushMessage) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, n.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (n NotificationTypeNewPushMessage) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, n.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (n NotificationTypeNewPushMessage) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, n.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (n NotificationTypeNewPushMessage) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, n.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (n NotificationTypeNewPushMessage) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, n.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (n NotificationTypeNewPushMessage) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, n.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (n NotificationTypeNewPushMessage) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, n.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (n NotificationTypeNewPushMessage) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, n.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (n NotificationTypeNewPushMessage) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, n.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (n NotificationTypeNewPushMessage) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, n.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (n NotificationTypeNewPushMessage) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, n.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (n NotificationTypeNewPushMessage) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, n.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (n NotificationTypeNewPushMessage) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, n.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (n NotificationTypeNewPushMessage) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, n.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (n NotificationTypeNewPushMessage) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, n.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (n NotificationTypeNewPushMessage) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, n.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (n NotificationTypeNewPushMessage) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, n.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (n NotificationTypeNewPushMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(n.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (n NotificationTypeNewPushMessage) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, n.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (n NotificationTypeNewPushMessage) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, n.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (n NotificationTypeNewPushMessage) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, n.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (n NotificationTypeNewPushMessage) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, n.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (n NotificationTypeNewPushMessage) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, n.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (n NotificationTypeNewPushMessage) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, n.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (n NotificationTypeNewPushMessage) ReportMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, n.MessageId, n.SenderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (n NotificationTypeNewPushMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, n.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (n NotificationTypeNewPushMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, n.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (n NotificationTypeNewPushMessage) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, n.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (n NotificationTypeNewPushMessage) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, n.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (n NotificationTypeNewPushMessage) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, n.MessageId, reactionTypes, isBig)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (n NotificationTypeNewPushMessage) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(n.SenderId, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (n NotificationTypeNewPushMessage) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, n.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (n NotificationTypeNewPushMessage) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, n.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (n NotificationTypeNewPushMessage) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, n.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (n NotificationTypeNewPushMessage) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, n.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (n NotificationTypeNewPushMessage) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, n.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (n NotificationTypeNewPushMessage) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, n.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (n NotificationTypeNewPushMessage) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, n.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (n NotificationTypeNewPushMessage) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, n.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (n NotificationTypeNewPushMessage) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, n.MessageId)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (o OrderInfo) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, o.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (o OrderInfo) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(o.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (o OrderInfo) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(o.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (o OrderInfo) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, o.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (o OrderInfo) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, o.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (o OrderInfo) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, o.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (o OrderInfo) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, o.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (o OrderInfo) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, o.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (o OrderInfo) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, o.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (o OrderInfo) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, o.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (o OrderInfo) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(o.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (o OrderInfo) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, o.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (o OrderInfo) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, o.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (o OrderInfo) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, o.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (o OrderInfo) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, o.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (o OrderInfo) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(o.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (o OrderInfo) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(o.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (o OrderInfo) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(o.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (o OrderInfo) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(o.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (o OrderInfo) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, o.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (o OrderInfo) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(o.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (o OrderInfo) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(o.Name, ignoreCache)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (o OrderInfo) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(o.PhoneNumber, onlyLocal)
}

// SendEmailAddressVerificationCode is a helper method for Client.SendEmailAddressVerificationCode
func (o OrderInfo) SendEmailAddressVerificationCode(client *Client) (*EmailAddressAuthenticationCodeInfo, error) {
	return client.SendEmailAddressVerificationCode(o.EmailAddress)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (o OrderInfo) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(o.PhoneNumber, typeField, opts)
}

// SetAuthenticationEmailAddress is a helper method for Client.SetAuthenticationEmailAddress
func (o OrderInfo) SetAuthenticationEmailAddress(client *Client) (*Ok, error) {
	return client.SetAuthenticationEmailAddress(o.EmailAddress)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (o OrderInfo) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(o.PhoneNumber, opts)
}

// SetBotName is a helper method for Client.SetBotName
func (o OrderInfo) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, o.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (o OrderInfo) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(o.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (o OrderInfo) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, o.Name)
}

// SetOption is a helper method for Client.SetOption
func (o OrderInfo) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(o.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (o OrderInfo) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, o.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (o OrderInfo) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, o.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (o OrderInfo) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(o.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (o OrderInfo) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, o.Name)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (p PageBlockAnchor) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, p.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (p PageBlockAnchor) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(p.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (p PageBlockAnchor) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(p.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (p PageBlockAnchor) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, p.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (p PageBlockAnchor) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (p PageBlockAnchor) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, p.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (p PageBlockAnchor) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, p.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (p PageBlockAnchor) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, p.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PageBlockAnchor) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, p.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (p PageBlockAnchor) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, p.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (p PageBlockAnchor) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(p.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (p PageBlockAnchor) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, p.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (p PageBlockAnchor) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (p PageBlockAnchor) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, p.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (p PageBlockAnchor) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, p.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (p PageBlockAnchor) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(p.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (p PageBlockAnchor) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(p.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (p PageBlockAnchor) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(p.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (p PageBlockAnchor) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(p.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (p PageBlockAnchor) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, p.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (p PageBlockAnchor) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(p.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (p PageBlockAnchor) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(p.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (p PageBlockAnchor) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, p.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (p PageBlockAnchor) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(p.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (p PageBlockAnchor) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, p.Name)
}

// SetOption is a helper method for Client.SetOption
func (p PageBlockAnchor) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(p.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (p PageBlockAnchor) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, p.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (p PageBlockAnchor) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, p.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p PageBlockAnchor) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(p.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (p PageBlockAnchor) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, p.Name)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (p PageBlockChatLink) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, p.Username)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (p PageBlockChatLink) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, p.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PageBlockChatLink) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, p.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (p PageBlockChatLink) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(p.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p PageBlockChatLink) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, p.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (p PageBlockChatLink) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(p.Title)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (p PageBlockChatLink) SearchChatAffiliateProgram(client *Client, referrer string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(p.Username, referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (p PageBlockChatLink) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(p.Username)
}

// SetAccentColor is a helper method for Client.SetAccentColor
func (p PageBlockChatLink) SetAccentColor(client *Client, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetAccentColor(p.AccentColorId, backgroundCustomEmojiId)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (p PageBlockChatLink) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, p.Username)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (p PageBlockChatLink) SetChatAccentColor(client *Client, chatId int64, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(chatId, p.AccentColorId, backgroundCustomEmojiId)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p PageBlockChatLink) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, p.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p PageBlockChatLink) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, p.Title)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (p PageBlockChatLink) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, p.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (p PageBlockChatLink) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(p.Username)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (p PageBlockChatLink) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, p.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (p PageBlockChatLink) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, p.Title, recordVideo, usePortraitOrientation)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (p PageBlockChatLink) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, p.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (p PageBlockChatLink) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, p.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (p PageBlockChatLink) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(p.Username, isActive)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PageBlockEmbedded) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PageBlockEmbedded) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PageBlockEmbedded) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PageBlockEmbedded) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (p PageBlockEmbedded) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, p.Width, p.Height, scale, chatId)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PageBlockEmbedded) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PageBlockEmbedded) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PageBlockEmbedded) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PageBlockEmbedded) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PageBlockEmbeddedPost) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PageBlockEmbeddedPost) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PageBlockEmbeddedPost) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PageBlockEmbeddedPost) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (p PageBlockEmbeddedPost) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, p.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (p PageBlockEmbeddedPost) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, p.Date)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PageBlockEmbeddedPost) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (p PageBlockEmbeddedPost) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, p.Date)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PageBlockEmbeddedPost) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PageBlockEmbeddedPost) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PageBlockEmbeddedPost) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// SetSavedMessagesTagLabel is a helper method for Client.SetSavedMessagesTagLabel
func (p PageBlockListItem) SetSavedMessagesTagLabel(client *Client, tag *ReactionType) (*Ok, error) {
	return client.SetSavedMessagesTagLabel(tag, p.Label)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (p PageBlockMap) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(p.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (p PageBlockMap) GetMapThumbnailFile(client *Client, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(p.Location, p.Zoom, p.Width, p.Height, scale, chatId)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PageBlockPhoto) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PageBlockPhoto) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PageBlockPhoto) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PageBlockPhoto) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PageBlockPhoto) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PageBlockPhoto) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PageBlockPhoto) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PageBlockPhoto) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PageBlockRelatedArticle) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PageBlockRelatedArticle) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (p PageBlockRelatedArticle) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, p.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PageBlockRelatedArticle) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, p.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (p PageBlockRelatedArticle) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(p.Title, isForum, isChannel, p.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p PageBlockRelatedArticle) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, p.Title, startDate, isRtmpStream)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PageBlockRelatedArticle) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PageBlockRelatedArticle) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PageBlockRelatedArticle) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (p PageBlockRelatedArticle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(p.Title)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PageBlockRelatedArticle) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PageBlockRelatedArticle) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PageBlockRelatedArticle) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (p PageBlockRelatedArticle) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, p.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (p PageBlockRelatedArticle) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, p.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p PageBlockRelatedArticle) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, p.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p PageBlockRelatedArticle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, p.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (p PageBlockRelatedArticle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, p.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (p PageBlockRelatedArticle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, p.Title, recordVideo, usePortraitOrientation)
}

// DiscardCall is a helper method for Client.DiscardCall
func (p PaidMediaPreview) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, p.Duration, isVideo, connectionId)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (p PaidMediaPreview) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, p.Width, p.Height, scale, chatId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (p PaidMediaPreview) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, p.Duration, paidMessageStarCount)
}

// AddChatMember is a helper method for Client.AddChatMember
func (p PaidReactionTypeChat) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(p.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (p PaidReactionTypeChat) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(p.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (p PaidReactionTypeChat) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(p.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (p PaidReactionTypeChat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(p.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (p PaidReactionTypeChat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, p.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (p PaidReactionTypeChat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(p.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (p PaidReactionTypeChat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(p.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (p PaidReactionTypeChat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(p.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PaidReactionTypeChat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(p.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (p PaidReactionTypeChat) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(p.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (p PaidReactionTypeChat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(p.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (p PaidReactionTypeChat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(p.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (p PaidReactionTypeChat) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(p.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (p PaidReactionTypeChat) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(p.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (p PaidReactionTypeChat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(p.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (p PaidReactionTypeChat) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(p.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (p PaidReactionTypeChat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(p.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (p PaidReactionTypeChat) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(p.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (p PaidReactionTypeChat) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(p.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (p PaidReactionTypeChat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(p.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (p PaidReactionTypeChat) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(p.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (p PaidReactionTypeChat) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(p.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (p PaidReactionTypeChat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(p.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p PaidReactionTypeChat) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(p.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (p PaidReactionTypeChat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(p.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (p PaidReactionTypeChat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(p.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (p PaidReactionTypeChat) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(p.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (p PaidReactionTypeChat) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(p.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (p PaidReactionTypeChat) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(p.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (p PaidReactionTypeChat) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(p.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (p PaidReactionTypeChat) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(p.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (p PaidReactionTypeChat) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(p.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (p PaidReactionTypeChat) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(p.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (p PaidReactionTypeChat) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(p.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (p PaidReactionTypeChat) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(p.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (p PaidReactionTypeChat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(p.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (p PaidReactionTypeChat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(p.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (p PaidReactionTypeChat) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(p.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (p PaidReactionTypeChat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(p.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (p PaidReactionTypeChat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, p.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (p PaidReactionTypeChat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, p.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (p PaidReactionTypeChat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, p.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (p PaidReactionTypeChat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, p.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (p PaidReactionTypeChat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, p.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (p PaidReactionTypeChat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, p.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (p PaidReactionTypeChat) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(p.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (p PaidReactionTypeChat) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(p.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (p PaidReactionTypeChat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(p.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (p PaidReactionTypeChat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(p.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (p PaidReactionTypeChat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(p.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (p PaidReactionTypeChat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(p.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (p PaidReactionTypeChat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(p.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (p PaidReactionTypeChat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(p.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (p PaidReactionTypeChat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(p.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (p PaidReactionTypeChat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(p.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (p PaidReactionTypeChat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(p.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (p PaidReactionTypeChat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, p.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (p PaidReactionTypeChat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(p.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (p PaidReactionTypeChat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(p.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (p PaidReactionTypeChat) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(p.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (p PaidReactionTypeChat) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(p.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (p PaidReactionTypeChat) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(p.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (p PaidReactionTypeChat) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(p.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (p PaidReactionTypeChat) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(p.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (p PaidReactionTypeChat) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(p.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (p PaidReactionTypeChat) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(p.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (p PaidReactionTypeChat) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(p.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (p PaidReactionTypeChat) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(p.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (p PaidReactionTypeChat) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(p.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (p PaidReactionTypeChat) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(p.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (p PaidReactionTypeChat) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(p.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (p PaidReactionTypeChat) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(p.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (p PaidReactionTypeChat) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(p.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (p PaidReactionTypeChat) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(p.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (p PaidReactionTypeChat) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(p.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (p PaidReactionTypeChat) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(p.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (p PaidReactionTypeChat) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(p.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (p PaidReactionTypeChat) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(p.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (p PaidReactionTypeChat) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(p.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (p PaidReactionTypeChat) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(p.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (p PaidReactionTypeChat) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(p.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (p PaidReactionTypeChat) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(p.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (p PaidReactionTypeChat) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(p.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (p PaidReactionTypeChat) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(p.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (p PaidReactionTypeChat) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(p.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (p PaidReactionTypeChat) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(p.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (p PaidReactionTypeChat) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(p.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (p PaidReactionTypeChat) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(p.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (p PaidReactionTypeChat) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(p.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (p PaidReactionTypeChat) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(p.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (p PaidReactionTypeChat) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(p.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (p PaidReactionTypeChat) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(p.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (p PaidReactionTypeChat) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(p.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (p PaidReactionTypeChat) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(p.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (p PaidReactionTypeChat) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(p.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (p PaidReactionTypeChat) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(p.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (p PaidReactionTypeChat) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(p.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (p PaidReactionTypeChat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(p.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (p PaidReactionTypeChat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(p.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (p PaidReactionTypeChat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(p.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (p PaidReactionTypeChat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(p.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (p PaidReactionTypeChat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(p.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (p PaidReactionTypeChat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(p.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (p PaidReactionTypeChat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, p.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (p PaidReactionTypeChat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(p.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (p PaidReactionTypeChat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(p.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (p PaidReactionTypeChat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(p.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (p PaidReactionTypeChat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(p.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (p PaidReactionTypeChat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, p.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (p PaidReactionTypeChat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(p.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (p PaidReactionTypeChat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(p.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (p PaidReactionTypeChat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(p.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (p PaidReactionTypeChat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(p.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (p PaidReactionTypeChat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(p.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (p PaidReactionTypeChat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(p.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (p PaidReactionTypeChat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(p.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (p PaidReactionTypeChat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(p.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (p PaidReactionTypeChat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(p.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (p PaidReactionTypeChat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(p.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (p PaidReactionTypeChat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(p.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (p PaidReactionTypeChat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(p.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (p PaidReactionTypeChat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(p.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (p PaidReactionTypeChat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(p.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (p PaidReactionTypeChat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(p.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (p PaidReactionTypeChat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(p.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (p PaidReactionTypeChat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(p.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (p PaidReactionTypeChat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(p.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (p PaidReactionTypeChat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(p.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (p PaidReactionTypeChat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(p.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (p PaidReactionTypeChat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, p.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (p PaidReactionTypeChat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(p.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (p PaidReactionTypeChat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(p.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (p PaidReactionTypeChat) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(p.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (p PaidReactionTypeChat) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(p.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (p PaidReactionTypeChat) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(p.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (p PaidReactionTypeChat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(p.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (p PaidReactionTypeChat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(p.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (p PaidReactionTypeChat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(p.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (p PaidReactionTypeChat) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(p.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (p PaidReactionTypeChat) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(p.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (p PaidReactionTypeChat) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(p.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (p PaidReactionTypeChat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(p.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (p PaidReactionTypeChat) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(p.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (p PaidReactionTypeChat) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(p.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (p PaidReactionTypeChat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(p.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PaidReactionTypeChat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(p.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (p PaidReactionTypeChat) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(p.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (p PaidReactionTypeChat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(p.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (p PaidReactionTypeChat) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(p.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (p PaidReactionTypeChat) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(p.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (p PaidReactionTypeChat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(p.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (p PaidReactionTypeChat) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(p.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (p PaidReactionTypeChat) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(p.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (p PaidReactionTypeChat) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(p.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (p PaidReactionTypeChat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(p.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (p PaidReactionTypeChat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(p.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (p PaidReactionTypeChat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, p.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (p PaidReactionTypeChat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(p.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (p PaidReactionTypeChat) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(p.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (p PaidReactionTypeChat) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(p.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (p PaidReactionTypeChat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(p.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (p PaidReactionTypeChat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(p.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (p PaidReactionTypeChat) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(p.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (p PaidReactionTypeChat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(p.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (p PaidReactionTypeChat) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, p.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (p PaidReactionTypeChat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(p.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (p PaidReactionTypeChat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(p.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (p PaidReactionTypeChat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(p.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (p PaidReactionTypeChat) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(p.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (p PaidReactionTypeChat) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(p.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (p PaidReactionTypeChat) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(p.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (p PaidReactionTypeChat) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(p.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (p PaidReactionTypeChat) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(p.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (p PaidReactionTypeChat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(p.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (p PaidReactionTypeChat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(p.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (p PaidReactionTypeChat) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(p.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (p PaidReactionTypeChat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, p.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (p PaidReactionTypeChat) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(p.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (p PaidReactionTypeChat) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(p.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (p PaidReactionTypeChat) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(p.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (p PaidReactionTypeChat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(p.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (p PaidReactionTypeChat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, p.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (p PaidReactionTypeChat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, p.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (p PaidReactionTypeChat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, p.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (p PaidReactionTypeChat) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(p.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (p PaidReactionTypeChat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(p.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (p PaidReactionTypeChat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(p.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (p PaidReactionTypeChat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(p.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (p PaidReactionTypeChat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(p.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (p PaidReactionTypeChat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(p.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (p PaidReactionTypeChat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, p.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (p PaidReactionTypeChat) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(p.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (p PaidReactionTypeChat) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(p.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (p PaidReactionTypeChat) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(p.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (p PaidReactionTypeChat) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(p.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (p PaidReactionTypeChat) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(p.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (p PaidReactionTypeChat) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(p.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (p PaidReactionTypeChat) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(p.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (p PaidReactionTypeChat) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(p.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (p PaidReactionTypeChat) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(p.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (p PaidReactionTypeChat) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(p.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (p PaidReactionTypeChat) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(p.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (p PaidReactionTypeChat) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(p.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (p PaidReactionTypeChat) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(p.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (p PaidReactionTypeChat) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(p.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (p PaidReactionTypeChat) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(p.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (p PaidReactionTypeChat) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(p.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (p PaidReactionTypeChat) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(p.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (p PaidReactionTypeChat) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(p.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (p PaidReactionTypeChat) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(p.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (p PaidReactionTypeChat) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(p.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (p PaidReactionTypeChat) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(p.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (p PaidReactionTypeChat) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(p.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (p PaidReactionTypeChat) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(p.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p PaidReactionTypeChat) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(p.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (p PaidReactionTypeChat) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(p.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (p PaidReactionTypeChat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(p.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (p PaidReactionTypeChat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(p.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (p PaidReactionTypeChat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(p.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (p PaidReactionTypeChat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(p.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (p PaidReactionTypeChat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(p.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (p PaidReactionTypeChat) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(p.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (p PaidReactionTypeChat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(p.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (p PaidReactionTypeChat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(p.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (p PaidReactionTypeChat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(p.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (p PaidReactionTypeChat) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(p.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (p PaidReactionTypeChat) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(p.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (p PaidReactionTypeChat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(p.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (p PaidReactionTypeChat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(p.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (p PaidReactionTypeChat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, p.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (p PaidReactionTypeChat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(p.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (p PaidReactionTypeChat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(p.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (p PaidReactionTypeChat) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(p.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (p PaidReactionTypeChat) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(p.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (p PaidReactionTypeChat) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(p.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (p PaidReactionTypeChat) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(p.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (p PaidReactionTypeChat) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(p.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (p PaidReactionTypeChat) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, p.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (p PaidReactionTypeChat) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(p.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (p PaidReactionTypeChat) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(p.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (p PaidReactionTypeChat) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(p.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (p PaidReactionTypeChat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(p.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (p PaidReactionTypeChat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(p.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (p PaidReactionTypeChat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(p.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (p PaidReactionTypeChat) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(p.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (p PaidReactionTypeChat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(p.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (p PaidReactionTypeChat) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(p.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (p PaidReactionTypeChat) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(p.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (p PaidReactionTypeChat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(p.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (p PaidReactionTypeChat) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(p.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (p PaidReactionTypeChat) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(p.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (p PaidReactionTypeChat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(p.ChatId, messageIds, forceRead, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (p PaidReactor) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, p.SenderId, disableNotification, inputMessageContent, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (p PaidReactor) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, p.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PaidReactor) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, p.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (p PaidReactor) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, p.StarCount)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (p PaidReactor) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, p.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (p PaidReactor) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, p.SenderId, reportSpam)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (p PaidReactor) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, p.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p PaidReactor) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, p.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PaidReactor) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, p.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (p PaidReactor) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, p.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PaidReactor) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, p.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PaidReactor) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, p.StarCount, userId, text, isPrivate)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (p PaidReactor) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, p.SenderId)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (p PaidReactor) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, p.StarCount)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (p PaidReactor) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(p.SenderId, opts)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (p PaidReactor) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, p.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (p PaidReactor) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, p.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (p PaidReactor) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, p.StarCount)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (p Passkey) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, p.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (p Passkey) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(p.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (p Passkey) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(p.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (p Passkey) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, p.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (p Passkey) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (p Passkey) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, p.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (p Passkey) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, p.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (p Passkey) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, p.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p Passkey) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, p.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (p Passkey) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, p.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (p Passkey) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(p.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (p Passkey) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, p.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (p Passkey) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (p Passkey) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, p.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (p Passkey) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, p.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (p Passkey) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(p.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (p Passkey) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(p.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (p Passkey) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(p.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (p Passkey) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(p.Name)
}

// RemoveLogin is a helper method for Client.RemoveLoginPasskey
func (p Passkey) RemoveLogin(client *Client) (*Ok, error) {
	return client.RemoveLoginPasskey(p.Id)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (p Passkey) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, p.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (p Passkey) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(p.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (p Passkey) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(p.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (p Passkey) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, p.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (p Passkey) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(p.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (p Passkey) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, p.Name)
}

// SetOption is a helper method for Client.SetOption
func (p Passkey) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(p.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (p Passkey) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, p.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (p Passkey) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, p.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p Passkey) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(p.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (p Passkey) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, p.Name)
}

// SendEmailAddressVerificationCode is a helper method for Client.SendEmailAddressVerificationCode
func (p PassportElementEmailAddress) SendEmailAddressVerificationCode(client *Client) (*EmailAddressAuthenticationCodeInfo, error) {
	return client.SendEmailAddressVerificationCode(p.EmailAddress)
}

// SetAuthenticationEmailAddress is a helper method for Client.SetAuthenticationEmailAddress
func (p PassportElementEmailAddress) SetAuthenticationEmailAddress(client *Client) (*Ok, error) {
	return client.SetAuthenticationEmailAddress(p.EmailAddress)
}

// DeletePassportElement is a helper method for Client.DeletePassportElement
func (p PassportElementError) DeletePassportElement(client *Client) (*Ok, error) {
	return client.DeletePassportElement(p.Type)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (p PassportElementError) GetPassportElement(client *Client, password string) (*PassportElement, error) {
	return client.GetPassportElement(p.Type, password)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (p PassportElementPhoneNumber) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(p.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (p PassportElementPhoneNumber) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(p.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (p PassportElementPhoneNumber) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(p.PhoneNumber, opts)
}

// DeletePassportElement is a helper method for Client.DeletePassportElement
func (p PassportSuitableElement) DeletePassportElement(client *Client) (*Ok, error) {
	return client.DeletePassportElement(p.Type)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (p PassportSuitableElement) GetPassportElement(client *Client, password string) (*PassportElement, error) {
	return client.GetPassportElement(p.Type, password)
}

// Send is a helper method for Client.SendPaymentForm
func (p PaymentForm) Send(client *Client, inputInvoice *InputInvoice, orderInfoId string, shippingOptionId string, tipAmount int64, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	return client.SendPaymentForm(inputInvoice, p.Id, orderInfoId, shippingOptionId, tipAmount, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (p PaymentFormTypeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, p.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PaymentFormTypeStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, p.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (p PaymentFormTypeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, p.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (p PaymentFormTypeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, p.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p PaymentFormTypeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, p.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PaymentFormTypeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, p.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (p PaymentFormTypeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, p.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PaymentFormTypeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, p.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PaymentFormTypeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, p.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (p PaymentFormTypeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, p.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (p PaymentFormTypeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, p.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (p PaymentFormTypeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, p.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (p PaymentFormTypeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, p.StarCount)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PaymentOption) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PaymentOption) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (p PaymentOption) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, p.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PaymentOption) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, p.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (p PaymentOption) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(p.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p PaymentOption) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, p.Title, startDate, isRtmpStream)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PaymentOption) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PaymentOption) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PaymentOption) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (p PaymentOption) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(p.Title)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PaymentOption) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PaymentOption) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PaymentOption) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p PaymentOption) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, p.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p PaymentOption) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, p.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (p PaymentOption) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, p.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (p PaymentOption) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, p.Title, recordVideo, usePortraitOrientation)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (p PaymentProviderOther) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, p.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (p PaymentProviderOther) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, p.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (p PaymentProviderOther) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, p.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (p PaymentProviderOther) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(p.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (p PaymentProviderOther) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(p.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (p PaymentProviderOther) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, p.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (p PaymentProviderOther) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(p.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (p PaymentProviderOther) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, p.Url, parameters, opts)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (p PaymentReceipt) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, p.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (p PaymentReceipt) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, p.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (p PaymentReceipt) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, p.Date)
}

// SendPaymentForm is a helper method for Client.SendPaymentForm
func (p PaymentReceiptTypeRegular) SendPaymentForm(client *Client, inputInvoice *InputInvoice, paymentFormId string, orderInfoId string, shippingOptionId string, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	return client.SendPaymentForm(inputInvoice, paymentFormId, orderInfoId, shippingOptionId, p.TipAmount, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (p PaymentReceiptTypeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, p.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PaymentReceiptTypeStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, p.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (p PaymentReceiptTypeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, p.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (p PaymentReceiptTypeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, p.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p PaymentReceiptTypeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, p.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PaymentReceiptTypeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, p.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (p PaymentReceiptTypeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, p.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PaymentReceiptTypeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, p.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PaymentReceiptTypeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, p.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (p PaymentReceiptTypeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, p.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (p PaymentReceiptTypeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, p.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (p PaymentReceiptTypeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, p.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (p PaymentReceiptTypeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, p.StarCount)
}

// GetCountryFlagEmoji is a helper method for Client.GetCountryFlagEmoji
func (p PersonalDetails) GetCountryFlagEmoji(client *Client) (*Text, error) {
	return client.GetCountryFlagEmoji(p.CountryCode)
}

// GetPreferredCountryLanguage is a helper method for Client.GetPreferredCountryLanguage
func (p PersonalDetails) GetPreferredCountryLanguage(client *Client) (*Text, error) {
	return client.GetPreferredCountryLanguage(p.CountryCode)
}

// RegisterUser is a helper method for Client.RegisterUser
func (p PersonalDetails) RegisterUser(client *Client, disableNotification bool) (*Ok, error) {
	return client.RegisterUser(p.FirstName, p.LastName, disableNotification)
}

// SetBusinessAccountName is a helper method for Client.SetBusinessAccountName
func (p PersonalDetails) SetBusinessAccountName(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountName(businessConnectionId, p.FirstName, p.LastName)
}

// SetName is a helper method for Client.SetName
func (p PersonalDetails) SetName(client *Client) (*Ok, error) {
	return client.SetName(p.FirstName, p.LastName)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (p PhotoSize) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, p.Width, p.Height, scale, chatId)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (p PhotoSize) SaveApplicationLogEvent(client *Client, chatId int64, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(p.Type, chatId, data)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (p Poll) ToggleForumTopicIsClosed(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(chatId, forumTopicId, p.IsClosed)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (p PollOption) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(p.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (p PollOption) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(p.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PollOption) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, p.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (p PollOption) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(p.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PollOption) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, p.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (p PollOption) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(p.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (p PollOption) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, p.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (p PollOption) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, p.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (p PollOption) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, p.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (p PollOption) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(p.Text, toLanguageCode)
}

// ViewPremiumFeature is a helper method for Client.ViewPremiumFeature
func (p PremiumFeaturePromotionAnimation) ViewPremiumFeature(client *Client) (*Ok, error) {
	return client.ViewPremiumFeature(p.Feature)
}

// AddChatMember is a helper method for Client.AddChatMember
func (p PremiumGiftCodeInfo) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, p.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (p PremiumGiftCodeInfo) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(p.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (p PremiumGiftCodeInfo) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(p.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (p PremiumGiftCodeInfo) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(p.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (p PremiumGiftCodeInfo) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(p.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (p PremiumGiftCodeInfo) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(p.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (p PremiumGiftCodeInfo) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(p.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PremiumGiftCodeInfo) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(p.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (p PremiumGiftCodeInfo) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(p.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (p PremiumGiftCodeInfo) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(p.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (p PremiumGiftCodeInfo) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, p.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (p PremiumGiftCodeInfo) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(p.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (p PremiumGiftCodeInfo) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, p.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (p PremiumGiftCodeInfo) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(p.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (p PremiumGiftCodeInfo) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(p.UserId)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (p PremiumGiftCodeInfo) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(p.MonthCount)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (p PremiumGiftCodeInfo) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(p.UserId)
}

// GetUser is a helper method for Client.GetUser
func (p PremiumGiftCodeInfo) GetUser(client *Client) (*User, error) {
	return client.GetUser(p.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (p PremiumGiftCodeInfo) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, p.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (p PremiumGiftCodeInfo) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(p.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (p PremiumGiftCodeInfo) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(p.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (p PremiumGiftCodeInfo) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(p.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (p PremiumGiftCodeInfo) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(p.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PremiumGiftCodeInfo) GiftPremiumWithStars(client *Client, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(p.UserId, starCount, p.MonthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (p PremiumGiftCodeInfo) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, p.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PremiumGiftCodeInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, p.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (p PremiumGiftCodeInfo) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, p.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (p PremiumGiftCodeInfo) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(p.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (p PremiumGiftCodeInfo) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(p.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (p PremiumGiftCodeInfo) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(p.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (p PremiumGiftCodeInfo) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, p.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (p PremiumGiftCodeInfo) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, p.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (p PremiumGiftCodeInfo) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(p.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (p PremiumGiftCodeInfo) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(p.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (p PremiumGiftCodeInfo) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(p.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (p PremiumGiftCodeInfo) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(p.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (p PremiumGiftCodeInfo) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(p.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (p PremiumGiftCodeInfo) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(p.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (p PremiumGiftCodeInfo) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(p.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (p PremiumGiftCodeInfo) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(p.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (p PremiumGiftCodeInfo) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(p.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (p PremiumGiftCodeInfo) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(p.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (p PremiumGiftCodeInfo) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, p.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (p PremiumGiftCodeInfo) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(p.UserId, stickerFormat, sticker)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (p PremiumGiftPaymentOption) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, p.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PremiumGiftPaymentOption) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, p.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (p PremiumGiftPaymentOption) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, p.StarCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (p PremiumGiftPaymentOption) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(p.Currency, p.Amount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (p PremiumGiftPaymentOption) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, p.StarCount)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (p PremiumGiftPaymentOption) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(p.MonthCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p PremiumGiftPaymentOption) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, p.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PremiumGiftPaymentOption) GiftPremiumWithStars(client *Client, userId int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, p.StarCount, p.MonthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (p PremiumGiftPaymentOption) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, p.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PremiumGiftPaymentOption) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, p.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PremiumGiftPaymentOption) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, p.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (p PremiumGiftPaymentOption) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, p.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (p PremiumGiftPaymentOption) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, p.Currency, p.Amount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (p PremiumGiftPaymentOption) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, p.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (p PremiumGiftPaymentOption) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, p.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (p PremiumGiftPaymentOption) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, p.StarCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (p PremiumGiveawayPaymentOption) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(p.Currency, p.Amount)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (p PremiumGiveawayPaymentOption) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(p.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PremiumGiveawayPaymentOption) GiftPremiumWithStars(client *Client, userId int64, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, p.MonthCount, text)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PremiumGiveawayPaymentOption) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, p.WinnerCount, starCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (p PremiumGiveawayPaymentOption) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, p.Currency, p.Amount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (p PremiumPaymentOption) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(p.Currency, p.Amount)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (p PremiumPaymentOption) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(p.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PremiumPaymentOption) GiftPremiumWithStars(client *Client, userId int64, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, p.MonthCount, text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (p PremiumPaymentOption) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, p.Currency, p.Amount)
}

// ViewPremiumFeature is a helper method for Client.ViewPremiumFeature
func (p PremiumSourceFeature) ViewPremiumFeature(client *Client) (*Ok, error) {
	return client.ViewPremiumFeature(p.Feature)
}

// GetPremiumLimit is a helper method for Client.GetPremiumLimit
func (p PremiumSourceLimitExceeded) GetPremiumLimit(client *Client) (*PremiumLimit, error) {
	return client.GetPremiumLimit(p.LimitType)
}

// GetRecentlyVisitedTMeUrls is a helper method for Client.GetRecentlyVisitedTMeUrls
func (p PremiumSourceLink) GetRecentlyVisitedTMeUrls(client *Client) (*TMeUrls, error) {
	return client.GetRecentlyVisitedTMeUrls(p.Referrer)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (p PremiumSourceLink) SearchChatAffiliateProgram(client *Client, username string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(username, p.Referrer)
}

// Launch is a helper method for Client.LaunchPrepaidGiveaway
func (p PrepaidGiveaway) Launch(client *Client, giveawayId string, parameters *GiveawayParameters, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, p.WinnerCount, starCount)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (p PreparedInlineMessage) AnswerInlineQuery(client *Client, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, nextOffset string, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(p.InlineQueryId, isPersonal, results, cacheTime, nextOffset, opts)
}

// Save is a helper method for Client.SavePreparedInlineMessage
func (p PreparedInlineMessage) Save(client *Client, userId int64, result *InputInlineQueryResult) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(userId, result, p.ChatTypes)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (p PreparedInlineMessageId) CreateChatInviteLink(client *Client, chatId int64, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, p.ExpirationDate, memberLimit, createsJoinRequest)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (p PreparedInlineMessageId) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, p.ExpirationDate, memberLimit, createsJoinRequest)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (p ProductInfo) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, p.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p ProductInfo) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, p.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (p ProductInfo) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(p.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p ProductInfo) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, p.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (p ProductInfo) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(p.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p ProductInfo) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, p.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p ProductInfo) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, p.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (p ProductInfo) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, p.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (p ProductInfo) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, p.Title, recordVideo, usePortraitOrientation)
}

// SetChat is a helper method for Client.SetChatProfileAccentColor
func (p ProfileAccentColor) SetChat(client *Client, chatId int64, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(chatId, p.Id, profileBackgroundCustomEmojiId)
}

// Set is a helper method for Client.SetProfileAccentColor
func (p ProfileAccentColor) Set(client *Client, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetProfileAccentColor(p.Id, profileBackgroundCustomEmojiId)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (p ProfilePhoto) AnswerInlineQuery(client *Client, inlineQueryId string, results []*InputInlineQueryResult, cacheTime int32, nextOffset string, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, p.IsPersonal, results, cacheTime, nextOffset, opts)
}

// Delete is a helper method for Client.DeleteProfilePhoto
func (p ProfilePhoto) Delete(client *Client) (*Ok, error) {
	return client.DeleteProfilePhoto(p.Id)
}

// Add is a helper method for Client.AddProxy
func (p Proxy) Add(client *Client, enable bool) (*Proxy, error) {
	return client.AddProxy(p.Server, p.Port, enable, p.Type)
}

// Edit is a helper method for Client.EditProxy
func (p Proxy) Edit(client *Client, enable bool) (*Proxy, error) {
	return client.EditProxy(p.Id, p.Server, p.Port, enable, p.Type)
}

// Enable is a helper method for Client.EnableProxy
func (p Proxy) Enable(client *Client) (*Ok, error) {
	return client.EnableProxy(p.Id)
}

// GetLink is a helper method for Client.GetProxyLink
func (p Proxy) GetLink(client *Client) (*HttpUrl, error) {
	return client.GetProxyLink(p.Id)
}

// Ping is a helper method for Client.PingProxy
func (p Proxy) Ping(client *Client) (*Seconds, error) {
	return client.PingProxy(p.Id)
}

// Remove is a helper method for Client.RemoveProxy
func (p Proxy) Remove(client *Client) (*Ok, error) {
	return client.RemoveProxy(p.Id)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (p Proxy) SetChatDirectMessagesGroup(client *Client, chatId int64, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, p.IsEnabled, paidMessageStarCount)
}

// Test is a helper method for Client.TestProxy
func (p Proxy) Test(client *Client, dcId int32, timeout float64) (*Ok, error) {
	return client.TestProxy(p.Server, p.Port, p.Type, dcId, timeout)
}

// CheckAuthenticationPassword is a helper method for Client.CheckAuthenticationPassword
func (p ProxyTypeHttp) CheckAuthenticationPassword(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPassword(p.Password)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (p ProxyTypeHttp) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, p.Username)
}

// CreateTemporaryPassword is a helper method for Client.CreateTemporaryPassword
func (p ProxyTypeHttp) CreateTemporaryPassword(client *Client, validFor int32) (*TemporaryPasswordState, error) {
	return client.CreateTemporaryPassword(p.Password, validFor)
}

// DeleteAccount is a helper method for Client.DeleteAccount
func (p ProxyTypeHttp) DeleteAccount(client *Client, reason string) (*Ok, error) {
	return client.DeleteAccount(reason, p.Password)
}

// GetAllPassportElements is a helper method for Client.GetAllPassportElements
func (p ProxyTypeHttp) GetAllPassportElements(client *Client) (*PassportElements, error) {
	return client.GetAllPassportElements(p.Password)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (p ProxyTypeHttp) GetChatRevenueWithdrawalUrl(client *Client, chatId int64) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(chatId, p.Password)
}

// GetPassportAuthorizationFormAvailableElements is a helper method for Client.GetPassportAuthorizationFormAvailableElements
func (p ProxyTypeHttp) GetPassportAuthorizationFormAvailableElements(client *Client, authorizationFormId int32) (*PassportElementsWithErrors, error) {
	return client.GetPassportAuthorizationFormAvailableElements(authorizationFormId, p.Password)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (p ProxyTypeHttp) GetPassportElement(client *Client, typeField *PassportElementType) (*PassportElement, error) {
	return client.GetPassportElement(typeField, p.Password)
}

// GetRecoveryEmailAddress is a helper method for Client.GetRecoveryEmailAddress
func (p ProxyTypeHttp) GetRecoveryEmailAddress(client *Client) (*RecoveryEmailAddress, error) {
	return client.GetRecoveryEmailAddress(p.Password)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p ProxyTypeHttp) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, starCount int64) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, starCount, p.Password)
}

// GetTonWithdrawalUrl is a helper method for Client.GetTonWithdrawalUrl
func (p ProxyTypeHttp) GetTonWithdrawalUrl(client *Client) (*HttpUrl, error) {
	return client.GetTonWithdrawalUrl(p.Password)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (p ProxyTypeHttp) GetUpgradedGiftWithdrawalUrl(client *Client, receivedGiftId string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(receivedGiftId, p.Password)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (p ProxyTypeHttp) SearchChatAffiliateProgram(client *Client, referrer string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(p.Username, referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (p ProxyTypeHttp) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(p.Username)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (p ProxyTypeHttp) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, p.Username)
}

// SetPassportElement is a helper method for Client.SetPassportElement
func (p ProxyTypeHttp) SetPassportElement(client *Client, element *InputPassportElement) (*PassportElement, error) {
	return client.SetPassportElement(element, p.Password)
}

// SetRecoveryEmailAddress is a helper method for Client.SetRecoveryEmailAddress
func (p ProxyTypeHttp) SetRecoveryEmailAddress(client *Client, newRecoveryEmailAddress string) (*PasswordState, error) {
	return client.SetRecoveryEmailAddress(p.Password, newRecoveryEmailAddress)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (p ProxyTypeHttp) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, p.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (p ProxyTypeHttp) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(p.Username)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (p ProxyTypeHttp) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, p.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (p ProxyTypeHttp) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, p.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (p ProxyTypeHttp) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(p.Username, isActive)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (p ProxyTypeHttp) TransferChatOwnership(client *Client, chatId int64, userId int64) (*Ok, error) {
	return client.TransferChatOwnership(chatId, userId, p.Password)
}

// CheckAuthenticationPassword is a helper method for Client.CheckAuthenticationPassword
func (p ProxyTypeSocks5) CheckAuthenticationPassword(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPassword(p.Password)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (p ProxyTypeSocks5) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, p.Username)
}

// CreateTemporaryPassword is a helper method for Client.CreateTemporaryPassword
func (p ProxyTypeSocks5) CreateTemporaryPassword(client *Client, validFor int32) (*TemporaryPasswordState, error) {
	return client.CreateTemporaryPassword(p.Password, validFor)
}

// DeleteAccount is a helper method for Client.DeleteAccount
func (p ProxyTypeSocks5) DeleteAccount(client *Client, reason string) (*Ok, error) {
	return client.DeleteAccount(reason, p.Password)
}

// GetAllPassportElements is a helper method for Client.GetAllPassportElements
func (p ProxyTypeSocks5) GetAllPassportElements(client *Client) (*PassportElements, error) {
	return client.GetAllPassportElements(p.Password)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (p ProxyTypeSocks5) GetChatRevenueWithdrawalUrl(client *Client, chatId int64) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(chatId, p.Password)
}

// GetPassportAuthorizationFormAvailableElements is a helper method for Client.GetPassportAuthorizationFormAvailableElements
func (p ProxyTypeSocks5) GetPassportAuthorizationFormAvailableElements(client *Client, authorizationFormId int32) (*PassportElementsWithErrors, error) {
	return client.GetPassportAuthorizationFormAvailableElements(authorizationFormId, p.Password)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (p ProxyTypeSocks5) GetPassportElement(client *Client, typeField *PassportElementType) (*PassportElement, error) {
	return client.GetPassportElement(typeField, p.Password)
}

// GetRecoveryEmailAddress is a helper method for Client.GetRecoveryEmailAddress
func (p ProxyTypeSocks5) GetRecoveryEmailAddress(client *Client) (*RecoveryEmailAddress, error) {
	return client.GetRecoveryEmailAddress(p.Password)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p ProxyTypeSocks5) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, starCount int64) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, starCount, p.Password)
}

// GetTonWithdrawalUrl is a helper method for Client.GetTonWithdrawalUrl
func (p ProxyTypeSocks5) GetTonWithdrawalUrl(client *Client) (*HttpUrl, error) {
	return client.GetTonWithdrawalUrl(p.Password)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (p ProxyTypeSocks5) GetUpgradedGiftWithdrawalUrl(client *Client, receivedGiftId string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(receivedGiftId, p.Password)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (p ProxyTypeSocks5) SearchChatAffiliateProgram(client *Client, referrer string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(p.Username, referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (p ProxyTypeSocks5) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(p.Username)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (p ProxyTypeSocks5) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, p.Username)
}

// SetPassportElement is a helper method for Client.SetPassportElement
func (p ProxyTypeSocks5) SetPassportElement(client *Client, element *InputPassportElement) (*PassportElement, error) {
	return client.SetPassportElement(element, p.Password)
}

// SetRecoveryEmailAddress is a helper method for Client.SetRecoveryEmailAddress
func (p ProxyTypeSocks5) SetRecoveryEmailAddress(client *Client, newRecoveryEmailAddress string) (*PasswordState, error) {
	return client.SetRecoveryEmailAddress(p.Password, newRecoveryEmailAddress)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (p ProxyTypeSocks5) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, p.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (p ProxyTypeSocks5) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(p.Username)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (p ProxyTypeSocks5) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, p.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (p ProxyTypeSocks5) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, p.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (p ProxyTypeSocks5) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(p.Username, isActive)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (p ProxyTypeSocks5) TransferChatOwnership(client *Client, chatId int64, userId int64) (*Ok, error) {
	return client.TransferChatOwnership(chatId, userId, p.Password)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (p PublicForwards) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, p.NextOffset, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (p PublicPostSearchLimits) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, p.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (p PublicPostSearchLimits) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, p.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (p PublicPostSearchLimits) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, p.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (p PublicPostSearchLimits) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, p.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (p PublicPostSearchLimits) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, p.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (p PublicPostSearchLimits) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, p.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (p PublicPostSearchLimits) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, p.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (p PublicPostSearchLimits) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, p.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (p PublicPostSearchLimits) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, p.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (p PublicPostSearchLimits) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, p.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (p PublicPostSearchLimits) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, p.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (p PublicPostSearchLimits) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, p.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (p PublicPostSearchLimits) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, p.StarCount)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (p PushMessageContentAnimation) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, messageId int64) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, messageId, p.IsPinned)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (p PushMessageContentAnimation) ToggleChatIsPinned(client *Client, chatList *ChatList, chatId int64) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, chatId, p.IsPinned)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (p PushMessageContentAnimation) ToggleForumTopicIsPinned(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(chatId, forumTopicId, p.IsPinned)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (p PushMessageContentAnimation) ToggleSavedMessagesTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, p.IsPinned)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (p PushMessageContentAudio) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, messageId int64) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, messageId, p.IsPinned)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (p PushMessageContentAudio) ToggleChatIsPinned(client *Client, chatList *ChatList, chatId int64) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, chatId, p.IsPinned)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (p PushMessageContentAudio) ToggleForumTopicIsPinned(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(chatId, forumTopicId, p.IsPinned)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (p PushMessageContentAudio) ToggleSavedMessagesTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, p.IsPinned)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (p PushMessageContentChatChangeTitle) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, p.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PushMessageContentChatChangeTitle) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, p.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (p PushMessageContentChatChangeTitle) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(p.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (p PushMessageContentChatChangeTitle) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, p.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (p PushMessageContentChatChangeTitle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(p.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (p PushMessageContentChatChangeTitle) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, p.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (p PushMessageContentChatChangeTitle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, p.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (p PushMessageContentChatChangeTitle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, p.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (p PushMessageContentChatChangeTitle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, p.Title, recordVideo, usePortraitOrientation)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (p PushMessageContentChatSetTheme) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, p.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (p PushMessageContentChatSetTheme) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(p.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (p PushMessageContentChatSetTheme) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(p.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (p PushMessageContentChatSetTheme) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, p.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (p PushMessageContentChatSetTheme) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (p PushMessageContentChatSetTheme) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, p.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (p PushMessageContentChatSetTheme) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, p.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (p PushMessageContentChatSetTheme) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, p.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (p PushMessageContentChatSetTheme) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, p.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (p PushMessageContentChatSetTheme) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, p.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (p PushMessageContentChatSetTheme) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(p.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (p PushMessageContentChatSetTheme) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, p.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (p PushMessageContentChatSetTheme) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, p.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (p PushMessageContentChatSetTheme) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, p.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (p PushMessageContentChatSetTheme) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, p.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (p PushMessageContentChatSetTheme) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(p.Name, typeField)
}
