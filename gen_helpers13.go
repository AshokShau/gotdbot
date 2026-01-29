package gotdbot

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StoryAreaTypeMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StoryAreaTypeMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StoryAreaTypeMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StoryAreaTypeMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StoryAreaTypeMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StoryAreaTypeMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StoryAreaTypeMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StoryAreaTypeMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StoryAreaTypeMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StoryAreaTypeMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StoryAreaTypeMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, s.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StoryAreaTypeMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StoryAreaTypeMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StoryAreaTypeMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StoryAreaTypeMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StoryAreaTypeMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StoryAreaTypeMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StoryAreaTypeMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StoryAreaTypeMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StoryAreaTypeMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StoryAreaTypeMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StoryAreaTypeMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StoryAreaTypeMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StoryAreaTypeMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StoryAreaTypeMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StoryAreaTypeMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StoryAreaTypeMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StoryAreaTypeMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StoryAreaTypeMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StoryAreaTypeMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StoryAreaTypeMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StoryAreaTypeMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, s.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StoryAreaTypeMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, s.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StoryAreaTypeMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StoryAreaTypeMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StoryAreaTypeMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, s.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StoryAreaTypeMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, s.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StoryAreaTypeMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StoryAreaTypeMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StoryAreaTypeMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(s.ChatId, s.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StoryAreaTypeMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, s.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StoryAreaTypeMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, s.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StoryAreaTypeMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, s.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StoryAreaTypeMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, s.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StoryAreaTypeMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StoryAreaTypeMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, s.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StoryAreaTypeMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, s.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StoryAreaTypeMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, s.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StoryAreaTypeMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, s.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StoryAreaTypeMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, s.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StoryAreaTypeMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StoryAreaTypeMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, s.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StoryAreaTypeMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, s.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StoryAreaTypeMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, s.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StoryAreaTypeMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, s.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StoryAreaTypeMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, s.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StoryAreaTypeMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, s.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StoryAreaTypeMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, s.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StoryAreaTypeMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StoryAreaTypeMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StoryAreaTypeMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StoryAreaTypeMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StoryAreaTypeMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StoryAreaTypeMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StoryAreaTypeMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StoryAreaTypeMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, s.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StoryAreaTypeMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StoryAreaTypeMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StoryAreaTypeMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StoryAreaTypeMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StoryAreaTypeMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StoryAreaTypeMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, s.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StoryAreaTypeMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StoryAreaTypeMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StoryAreaTypeMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, s.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StoryAreaTypeMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StoryAreaTypeMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, s.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StoryAreaTypeMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StoryAreaTypeMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StoryAreaTypeMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (s StoryAreaTypeMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(s.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StoryAreaTypeMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, s.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StoryAreaTypeMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StoryAreaTypeMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StoryAreaTypeMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StoryAreaTypeMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StoryAreaTypeMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StoryAreaTypeMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, s.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StoryAreaTypeMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, s.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StoryAreaTypeMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StoryAreaTypeMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StoryAreaTypeMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, s.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StoryAreaTypeMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StoryAreaTypeMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StoryAreaTypeMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StoryAreaTypeMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StoryAreaTypeMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StoryAreaTypeMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StoryAreaTypeMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StoryAreaTypeMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StoryAreaTypeMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StoryAreaTypeMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StoryAreaTypeMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StoryAreaTypeMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, s.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StoryAreaTypeMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, s.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s StoryAreaTypeMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, s.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StoryAreaTypeMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StoryAreaTypeMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StoryAreaTypeMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StoryAreaTypeMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StoryAreaTypeMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StoryAreaTypeMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StoryAreaTypeMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StoryAreaTypeMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StoryAreaTypeMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StoryAreaTypeMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StoryAreaTypeMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StoryAreaTypeMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StoryAreaTypeMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StoryAreaTypeMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StoryAreaTypeMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StoryAreaTypeMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StoryAreaTypeMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, s.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StoryAreaTypeMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StoryAreaTypeMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StoryAreaTypeMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StoryAreaTypeMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StoryAreaTypeMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StoryAreaTypeMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StoryAreaTypeMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StoryAreaTypeMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StoryAreaTypeMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StoryAreaTypeMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StoryAreaTypeMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StoryAreaTypeMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StoryAreaTypeMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StoryAreaTypeMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StoryAreaTypeMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StoryAreaTypeMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StoryAreaTypeMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StoryAreaTypeMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StoryAreaTypeMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StoryAreaTypeMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StoryAreaTypeMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StoryAreaTypeMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StoryAreaTypeMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StoryAreaTypeMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StoryAreaTypeMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StoryAreaTypeMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StoryAreaTypeMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, s.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StoryAreaTypeMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, s.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StoryAreaTypeMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, s.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StoryAreaTypeMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, s.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StoryAreaTypeMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StoryAreaTypeMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StoryAreaTypeMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, s.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StoryAreaTypeMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StoryAreaTypeMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StoryAreaTypeMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, s.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StoryAreaTypeMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, s.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StoryAreaTypeMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StoryAreaTypeMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StoryAreaTypeMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, s.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StoryAreaTypeMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, s.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StoryAreaTypeMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StoryAreaTypeMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StoryAreaTypeMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StoryAreaTypeMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StoryAreaTypeMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StoryAreaTypeMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StoryAreaTypeMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StoryAreaTypeMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StoryAreaTypeMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StoryAreaTypeMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StoryAreaTypeMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StoryAreaTypeMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StoryAreaTypeMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StoryAreaTypeMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, s.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StoryAreaTypeMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StoryAreaTypeMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StoryAreaTypeMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StoryAreaTypeMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, s.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StoryAreaTypeMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StoryAreaTypeMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StoryAreaTypeSuggestedReaction) AddMessageReaction(client *Client, chatId int64, messageId int64, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, messageId, s.ReactionType, isBig, updateRecentReactions)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StoryAreaTypeSuggestedReaction) GetChatRevenueStatistics(client *Client, chatId int64) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(chatId, s.IsDark)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StoryAreaTypeSuggestedReaction) GetChatStatistics(client *Client, chatId int64) (*ChatStatistics, error) {
	return client.GetChatStatistics(chatId, s.IsDark)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StoryAreaTypeSuggestedReaction) GetMessageStatistics(client *Client, chatId int64, messageId int64) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, messageId, s.IsDark)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StoryAreaTypeSuggestedReaction) GetStarRevenueStatistics(client *Client, ownerId *MessageSender) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(ownerId, s.IsDark)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StoryAreaTypeSuggestedReaction) GetStoryStatistics(client *Client, chatId int64, storyId int32) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, storyId, s.IsDark)
}

// GetTonRevenueStatistics is a helper method for Client.GetTonRevenueStatistics
func (s StoryAreaTypeSuggestedReaction) GetTonRevenueStatistics(client *Client) (*TonRevenueStatistics, error) {
	return client.GetTonRevenueStatistics(s.IsDark)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StoryAreaTypeSuggestedReaction) RemoveMessageReaction(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, messageId, s.ReactionType)
}

// SetDefaultReactionType is a helper method for Client.SetDefaultReactionType
func (s StoryAreaTypeSuggestedReaction) SetDefaultReactionType(client *Client) (*Ok, error) {
	return client.SetDefaultReactionType(s.ReactionType)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StoryAreaTypeUpgradedGift) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, s.GiftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StoryAreaTypeUpgradedGift) SendResoldGift(client *Client, ownerId *MessageSender, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(s.GiftName, ownerId, price)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (s StoryAreaTypeWeather) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(s.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (s StoryAreaTypeWeather) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(s.Emoji)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s StoryContentLive) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(s.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (s StoryContentLive) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(s.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (s StoryContentLive) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(s.GroupCallId)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StoryContentLive) CreateVideoChat(client *Client, chatId int64, title string, startDate int32) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, startDate, s.IsRtmpStream)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (s StoryContentLive) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(s.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (s StoryContentLive) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(s.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (s StoryContentLive) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(s.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (s StoryContentLive) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(s.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (s StoryContentLive) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(s.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (s StoryContentLive) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(s.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (s StoryContentLive) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(s.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (s StoryContentLive) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(s.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (s StoryContentLive) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(s.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (s StoryContentLive) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(s.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (s StoryContentLive) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(s.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (s StoryContentLive) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(s.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (s StoryContentLive) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(s.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (s StoryContentLive) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(s.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StoryContentLive) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(s.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (s StoryContentLive) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(s.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (s StoryContentLive) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(s.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (s StoryContentLive) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(s.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (s StoryContentLive) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(s.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (s StoryContentLive) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(s.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (s StoryContentLive) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(s.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (s StoryContentLive) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(s.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (s StoryContentLive) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(s.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (s StoryContentLive) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(s.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (s StoryContentLive) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(s.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (s StoryContentLive) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(s.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (s StoryContentLive) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(s.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (s StoryContentLive) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(s.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (s StoryContentLive) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(s.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (s StoryContentLive) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(s.GroupCallId, audioSourceId, payload)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StoryContentLive) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, s.IsRtmpStream, enableMessages, paidMessageStarCount)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (s StoryContentLive) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(s.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (s StoryContentLive) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(s.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (s StoryContentLive) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(s.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (s StoryContentLive) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(s.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (s StoryContentLive) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(s.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (s StoryContentLive) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(s.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (s StoryContentLive) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(s.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (s StoryContentLive) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(s.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (s StoryContentLive) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(s.GroupCallId, muteNewParticipants)
}

// CloseStory is a helper method for Client.CloseStory
func (s StoryFullId) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, s.StoryId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (s StoryFullId) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, s.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (s StoryFullId) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, s.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (s StoryFullId) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, s.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (s StoryFullId) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, s.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (s StoryFullId) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, s.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (s StoryFullId) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, s.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (s StoryFullId) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, s.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s StoryFullId) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(s.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (s StoryFullId) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, s.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StoryFullId) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, s.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (s StoryFullId) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, s.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (s StoryFullId) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, s.StoryId, optionId, text)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (s StoryFullId) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(s.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (s StoryFullId) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, s.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (s StoryFullId) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, s.StoryId, isPostedToChatPage)
}

// CloseStory is a helper method for Client.CloseStory
func (s StoryInfo) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, s.StoryId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (s StoryInfo) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, s.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (s StoryInfo) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, s.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (s StoryInfo) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, s.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (s StoryInfo) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, s.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (s StoryInfo) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, s.StoryId, coverFrameTimestamp)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StoryInfo) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, s.Date)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (s StoryInfo) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, s.StoryId, preferForwards, offset, limit, opts)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StoryInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, s.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (s StoryInfo) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, s.Date)
}

// GetStory is a helper method for Client.GetStory
func (s StoryInfo) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, s.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s StoryInfo) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(s.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (s StoryInfo) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, s.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StoryInfo) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, s.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (s StoryInfo) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, s.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (s StoryInfo) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, s.StoryId, optionId, text)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (s StoryInfo) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(s.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (s StoryInfo) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, s.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (s StoryInfo) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, s.StoryId, isPostedToChatPage)
}

// GetBlockedMessageSenders is a helper method for Client.GetBlockedMessageSenders
func (s StoryInteraction) GetBlockedMessageSenders(client *Client, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetBlockedMessageSenders(s.BlockList, offset, limit)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (s StoryInteractions) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, s.NextOffset, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StoryOriginPublicStory) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StoryOriginPublicStory) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StoryOriginPublicStory) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StoryOriginPublicStory) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StoryOriginPublicStory) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StoryOriginPublicStory) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StoryOriginPublicStory) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StoryOriginPublicStory) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StoryOriginPublicStory) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StoryOriginPublicStory) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StoryOriginPublicStory) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StoryOriginPublicStory) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StoryOriginPublicStory) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StoryOriginPublicStory) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StoryOriginPublicStory) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StoryOriginPublicStory) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StoryOriginPublicStory) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StoryOriginPublicStory) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StoryOriginPublicStory) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CloseStory is a helper method for Client.CloseStory
func (s StoryOriginPublicStory) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, s.StoryId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StoryOriginPublicStory) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StoryOriginPublicStory) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StoryOriginPublicStory) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StoryOriginPublicStory) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StoryOriginPublicStory) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StoryOriginPublicStory) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StoryOriginPublicStory) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StoryOriginPublicStory) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (s StoryOriginPublicStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, s.StoryId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StoryOriginPublicStory) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StoryOriginPublicStory) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StoryOriginPublicStory) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StoryOriginPublicStory) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StoryOriginPublicStory) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StoryOriginPublicStory) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StoryOriginPublicStory) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StoryOriginPublicStory) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StoryOriginPublicStory) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StoryOriginPublicStory) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StoryOriginPublicStory) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStory is a helper method for Client.DeleteStory
func (s StoryOriginPublicStory) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, s.StoryId)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StoryOriginPublicStory) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StoryOriginPublicStory) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StoryOriginPublicStory) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StoryOriginPublicStory) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StoryOriginPublicStory) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StoryOriginPublicStory) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StoryOriginPublicStory) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (s StoryOriginPublicStory) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, s.StoryId, content, areas, caption, privacySettings)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StoryOriginPublicStory) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StoryOriginPublicStory) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StoryOriginPublicStory) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StoryOriginPublicStory) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StoryOriginPublicStory) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StoryOriginPublicStory) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StoryOriginPublicStory) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StoryOriginPublicStory) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StoryOriginPublicStory) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StoryOriginPublicStory) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// EditStory is a helper method for Client.EditStory
func (s StoryOriginPublicStory) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, s.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (s StoryOriginPublicStory) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, s.StoryId, coverFrameTimestamp)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StoryOriginPublicStory) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StoryOriginPublicStory) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StoryOriginPublicStory) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StoryOriginPublicStory) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StoryOriginPublicStory) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StoryOriginPublicStory) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StoryOriginPublicStory) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StoryOriginPublicStory) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StoryOriginPublicStory) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StoryOriginPublicStory) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StoryOriginPublicStory) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StoryOriginPublicStory) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StoryOriginPublicStory) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StoryOriginPublicStory) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StoryOriginPublicStory) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StoryOriginPublicStory) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StoryOriginPublicStory) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StoryOriginPublicStory) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StoryOriginPublicStory) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StoryOriginPublicStory) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StoryOriginPublicStory) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StoryOriginPublicStory) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StoryOriginPublicStory) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StoryOriginPublicStory) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StoryOriginPublicStory) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StoryOriginPublicStory) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StoryOriginPublicStory) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StoryOriginPublicStory) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StoryOriginPublicStory) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StoryOriginPublicStory) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StoryOriginPublicStory) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StoryOriginPublicStory) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StoryOriginPublicStory) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StoryOriginPublicStory) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StoryOriginPublicStory) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StoryOriginPublicStory) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StoryOriginPublicStory) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StoryOriginPublicStory) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (s StoryOriginPublicStory) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, s.StoryId, preferForwards, offset, limit, opts)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StoryOriginPublicStory) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StoryOriginPublicStory) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StoryOriginPublicStory) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StoryOriginPublicStory) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StoryOriginPublicStory) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StoryOriginPublicStory) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StoryOriginPublicStory) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StoryOriginPublicStory) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StoryOriginPublicStory) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StoryOriginPublicStory) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StoryOriginPublicStory) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StoryOriginPublicStory) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StoryOriginPublicStory) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StoryOriginPublicStory) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StoryOriginPublicStory) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StoryOriginPublicStory) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StoryOriginPublicStory) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StoryOriginPublicStory) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StoryOriginPublicStory) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StoryOriginPublicStory) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StoryOriginPublicStory) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StoryOriginPublicStory) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StoryOriginPublicStory) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StoryOriginPublicStory) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StoryOriginPublicStory) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StoryOriginPublicStory) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StoryOriginPublicStory) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StoryOriginPublicStory) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StoryOriginPublicStory) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StoryOriginPublicStory) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StoryOriginPublicStory) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StoryOriginPublicStory) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StoryOriginPublicStory) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StoryOriginPublicStory) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StoryOriginPublicStory) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StoryOriginPublicStory) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StoryOriginPublicStory) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStory is a helper method for Client.GetStory
func (s StoryOriginPublicStory) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, s.StoryId, onlyLocal)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StoryOriginPublicStory) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s StoryOriginPublicStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(s.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (s StoryOriginPublicStory) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, s.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StoryOriginPublicStory) GetStoryStatistics(client *Client, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, s.StoryId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StoryOriginPublicStory) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StoryOriginPublicStory) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StoryOriginPublicStory) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StoryOriginPublicStory) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StoryOriginPublicStory) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StoryOriginPublicStory) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StoryOriginPublicStory) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StoryOriginPublicStory) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StoryOriginPublicStory) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StoryOriginPublicStory) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StoryOriginPublicStory) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StoryOriginPublicStory) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StoryOriginPublicStory) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenStory is a helper method for Client.OpenStory
func (s StoryOriginPublicStory) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, s.StoryId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StoryOriginPublicStory) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StoryOriginPublicStory) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StoryOriginPublicStory) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StoryOriginPublicStory) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StoryOriginPublicStory) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StoryOriginPublicStory) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StoryOriginPublicStory) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StoryOriginPublicStory) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StoryOriginPublicStory) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StoryOriginPublicStory) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StoryOriginPublicStory) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StoryOriginPublicStory) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StoryOriginPublicStory) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StoryOriginPublicStory) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StoryOriginPublicStory) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StoryOriginPublicStory) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StoryOriginPublicStory) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StoryOriginPublicStory) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StoryOriginPublicStory) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StoryOriginPublicStory) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StoryOriginPublicStory) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StoryOriginPublicStory) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StoryOriginPublicStory) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StoryOriginPublicStory) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StoryOriginPublicStory) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StoryOriginPublicStory) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StoryOriginPublicStory) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StoryOriginPublicStory) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StoryOriginPublicStory) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ReportStory is a helper method for Client.ReportStory
func (s StoryOriginPublicStory) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, s.StoryId, optionId, text)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StoryOriginPublicStory) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StoryOriginPublicStory) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StoryOriginPublicStory) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StoryOriginPublicStory) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StoryOriginPublicStory) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StoryOriginPublicStory) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StoryOriginPublicStory) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StoryOriginPublicStory) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StoryOriginPublicStory) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StoryOriginPublicStory) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StoryOriginPublicStory) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StoryOriginPublicStory) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StoryOriginPublicStory) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StoryOriginPublicStory) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StoryOriginPublicStory) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StoryOriginPublicStory) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StoryOriginPublicStory) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StoryOriginPublicStory) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StoryOriginPublicStory) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StoryOriginPublicStory) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StoryOriginPublicStory) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StoryOriginPublicStory) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StoryOriginPublicStory) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StoryOriginPublicStory) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StoryOriginPublicStory) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StoryOriginPublicStory) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StoryOriginPublicStory) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StoryOriginPublicStory) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StoryOriginPublicStory) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StoryOriginPublicStory) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StoryOriginPublicStory) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StoryOriginPublicStory) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StoryOriginPublicStory) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StoryOriginPublicStory) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StoryOriginPublicStory) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StoryOriginPublicStory) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StoryOriginPublicStory) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StoryOriginPublicStory) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StoryOriginPublicStory) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StoryOriginPublicStory) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StoryOriginPublicStory) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StoryOriginPublicStory) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StoryOriginPublicStory) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StoryOriginPublicStory) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StoryOriginPublicStory) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StoryOriginPublicStory) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StoryOriginPublicStory) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StoryOriginPublicStory) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StoryOriginPublicStory) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StoryOriginPublicStory) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StoryOriginPublicStory) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (s StoryOriginPublicStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(s.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (s StoryOriginPublicStory) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, s.StoryId, updateRecentReactions, opts)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StoryOriginPublicStory) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StoryOriginPublicStory) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StoryOriginPublicStory) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StoryOriginPublicStory) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StoryOriginPublicStory) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StoryOriginPublicStory) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StoryOriginPublicStory) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StoryOriginPublicStory) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StoryOriginPublicStory) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StoryOriginPublicStory) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StoryOriginPublicStory) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StoryOriginPublicStory) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StoryOriginPublicStory) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StoryOriginPublicStory) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StoryOriginPublicStory) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StoryOriginPublicStory) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StoryOriginPublicStory) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StoryOriginPublicStory) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StoryOriginPublicStory) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (s StoryOriginPublicStory) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, s.StoryId, isPostedToChatPage)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StoryOriginPublicStory) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StoryOriginPublicStory) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StoryOriginPublicStory) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StoryOriginPublicStory) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StoryOriginPublicStory) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StoryOriginPublicStory) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StoryOriginPublicStory) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StoryOriginPublicStory) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StoryPrivacySettingsSelectedUsers) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, s.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (s StoryPrivacySettingsSelectedUsers) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(s.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StoryPrivacySettingsSelectedUsers) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, s.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (s StoryPrivacySettingsSelectedUsers) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, s.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (s StoryPrivacySettingsSelectedUsers) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(s.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (s StoryPrivacySettingsSelectedUsers) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(s.UserIds)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (s StoryVideo) EditStoryCover(client *Client, storyPosterChatId int64, storyId int32) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, storyId, s.CoverFrameTimestamp)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StoryVideo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, s.Width, s.Height, scale, chatId)
}

// CreateSupergroupChat is a helper method for Client.CreateSupergroupChat
func (s SuggestedActionConvertToBroadcastGroup) CreateSupergroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(s.SupergroupId, force)
}

// DisableAllSupergroupUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (s SuggestedActionConvertToBroadcastGroup) DisableAllSupergroupUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(s.SupergroupId)
}

// GetSupergroup is a helper method for Client.GetSupergroup
func (s SuggestedActionConvertToBroadcastGroup) GetSupergroup(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(s.SupergroupId)
}

// GetSupergroupFullInfo is a helper method for Client.GetSupergroupFullInfo
func (s SuggestedActionConvertToBroadcastGroup) GetSupergroupFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(s.SupergroupId)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (s SuggestedActionConvertToBroadcastGroup) GetSupergroupMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(s.SupergroupId, offset, limit, opts)
}

// ReorderSupergroupActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (s SuggestedActionConvertToBroadcastGroup) ReorderSupergroupActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(s.SupergroupId, usernames)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s SuggestedActionConvertToBroadcastGroup) ReportSupergroupAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(s.SupergroupId, messageId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (s SuggestedActionConvertToBroadcastGroup) ReportSupergroupSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(s.SupergroupId, messageIds)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (s SuggestedActionConvertToBroadcastGroup) SetSupergroupCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(s.SupergroupId, customEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (s SuggestedActionConvertToBroadcastGroup) SetSupergroupMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(s.SupergroupId, mainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (s SuggestedActionConvertToBroadcastGroup) SetSupergroupStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(s.SupergroupId, stickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (s SuggestedActionConvertToBroadcastGroup) SetSupergroupUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(s.SupergroupId, unrestrictBoostCount)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (s SuggestedActionConvertToBroadcastGroup) SetSupergroupUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(s.SupergroupId, username)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(s.SupergroupId, canHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(s.SupergroupId, hasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupHasAutomaticTranslation(client *Client, hasAutomaticTranslation bool) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(s.SupergroupId, hasAutomaticTranslation)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(s.SupergroupId, hasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(s.SupergroupId, isAllHistoryAvailable)
}

// ToggleSupergroupIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(s.SupergroupId)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupIsForum(client *Client, isForum bool, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(s.SupergroupId, isForum, hasForumTabs)
}

// ToggleSupergroupJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupJoinByRequest(client *Client, joinByRequest bool) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(s.SupergroupId, joinByRequest)
}

// ToggleSupergroupJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupJoinToSendMessages(client *Client, joinToSendMessages bool) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(s.SupergroupId, joinToSendMessages)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupSignMessages(client *Client, signMessages bool, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(s.SupergroupId, signMessages, showMessageSender)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (s SuggestedActionConvertToBroadcastGroup) ToggleSupergroupUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(s.SupergroupId, username, isActive)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s SuggestedActionCustom) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, s.Name, sticker)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (s SuggestedActionCustom) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, s.Url, cacheTime)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (s SuggestedActionCustom) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(s.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (s SuggestedActionCustom) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(s.Name)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (s SuggestedActionCustom) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, s.Url)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (s SuggestedActionCustom) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, s.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s SuggestedActionCustom) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s SuggestedActionCustom) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, s.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s SuggestedActionCustom) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, s.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s SuggestedActionCustom) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, s.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s SuggestedActionCustom) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, s.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (s SuggestedActionCustom) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, s.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (s SuggestedActionCustom) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(s.Name)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (s SuggestedActionCustom) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, s.Url)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (s SuggestedActionCustom) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, s.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s SuggestedActionCustom) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s SuggestedActionCustom) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, s.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s SuggestedActionCustom) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, s.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (s SuggestedActionCustom) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(s.Name, typeField)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (s SuggestedActionCustom) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(s.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (s SuggestedActionCustom) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(s.Url)
}

// GetOption is a helper method for Client.GetOption
func (s SuggestedActionCustom) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(s.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (s SuggestedActionCustom) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(s.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (s SuggestedActionCustom) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(s.Name)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (s SuggestedActionCustom) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, s.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (s SuggestedActionCustom) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(s.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s SuggestedActionCustom) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, s.Url, parameters, opts)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s SuggestedActionCustom) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, s.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (s SuggestedActionCustom) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(s.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (s SuggestedActionCustom) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(s.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (s SuggestedActionCustom) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, s.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (s SuggestedActionCustom) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(s.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s SuggestedActionCustom) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, s.Name)
}

// SetOption is a helper method for Client.SetOption
func (s SuggestedActionCustom) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(s.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (s SuggestedActionCustom) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, s.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s SuggestedActionCustom) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, s.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (s SuggestedActionCustom) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(s.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s SuggestedActionCustom) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, s.Name)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s SuggestedPostInfo) ApproveSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, messageId, s.SendDate)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s SuggestedPostPriceStar) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, s.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s SuggestedPostPriceStar) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, s.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s SuggestedPostPriceStar) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, s.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (s SuggestedPostPriceStar) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, s.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s SuggestedPostPriceStar) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, s.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s SuggestedPostPriceStar) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, s.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (s SuggestedPostPriceStar) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, s.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s SuggestedPostPriceStar) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, s.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s SuggestedPostPriceStar) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, s.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SuggestedPostPriceStar) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, s.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (s SuggestedPostPriceStar) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, s.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (s SuggestedPostPriceStar) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, s.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (s SuggestedPostPriceStar) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, s.StarCount)
}

// CreateNewChat is a helper method for Client.CreateNewSupergroupChat
func (s Supergroup) CreateNewChat(client *Client, title string, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, s.IsForum, s.IsChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateChat is a helper method for Client.CreateSupergroupChat
func (s Supergroup) CreateChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(s.Id, force)
}

// DisableAllUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (s Supergroup) DisableAllUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(s.Id)
}

// GetChatBoostFeatures is a helper method for Client.GetChatBoostFeatures
func (s Supergroup) GetChatBoostFeatures(client *Client) (*ChatBoostFeatures, error) {
	return client.GetChatBoostFeatures(s.IsChannel)
}

// GetChatBoostLevelFeatures is a helper method for Client.GetChatBoostLevelFeatures
func (s Supergroup) GetChatBoostLevelFeatures(client *Client, level int32) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(s.IsChannel, level)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s Supergroup) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, s.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s Supergroup) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, s.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (s Supergroup) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, s.Date)
}

// Get is a helper method for Client.GetSupergroup
func (s Supergroup) Get(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(s.Id)
}

// GetFullInfo is a helper method for Client.GetSupergroupFullInfo
func (s Supergroup) GetFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(s.Id)
}

// GetMembers is a helper method for Client.GetSupergroupMembers
func (s Supergroup) GetMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(s.Id, offset, limit, opts)
}

// ReorderActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (s Supergroup) ReorderActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(s.Id, usernames)
}

// ReportAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s Supergroup) ReportAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(s.Id, messageId)
}

// ReportSpam is a helper method for Client.ReportSupergroupSpam
func (s Supergroup) ReportSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(s.Id, messageIds)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s Supergroup) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, s.PaidMessageStarCount, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s Supergroup) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, s.PaidMessageStarCount)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (s Supergroup) SendGroupCallMessage(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, s.PaidMessageStarCount)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s Supergroup) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, s.PaidMessageStarCount)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s Supergroup) SetChatMemberStatus(client *Client, chatId int64, memberId *MessageSender) (*Ok, error) {
	return client.SetChatMemberStatus(chatId, memberId, s.Status)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s Supergroup) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, s.PaidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (s Supergroup) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, s.PaidMessageStarCount)
}

// SetCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (s Supergroup) SetCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(s.Id, customEmojiStickerSetId)
}

// SetMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (s Supergroup) SetMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(s.Id, mainProfileTab)
}

// SetStickerSet is a helper method for Client.SetSupergroupStickerSet
func (s Supergroup) SetStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(s.Id, stickerSetId)
}

// SetUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (s Supergroup) SetUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(s.Id, unrestrictBoostCount)
}

// SetUsername is a helper method for Client.SetSupergroupUsername
func (s Supergroup) SetUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(s.Id, username)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s Supergroup) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, isRtmpStream, enableMessages, s.PaidMessageStarCount)
}

// ToggleCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (s Supergroup) ToggleCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(s.Id, canHaveSponsoredMessages)
}

// ToggleHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (s Supergroup) ToggleHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(s.Id, hasAggressiveAntiSpamEnabled)
}

// ToggleHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (s Supergroup) ToggleHasAutomaticTranslation(client *Client) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(s.Id, s.HasAutomaticTranslation)
}

// ToggleHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (s Supergroup) ToggleHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(s.Id, hasHiddenMembers)
}

// ToggleIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (s Supergroup) ToggleIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(s.Id, isAllHistoryAvailable)
}

// ToggleIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (s Supergroup) ToggleIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(s.Id)
}

// ToggleIsForum is a helper method for Client.ToggleSupergroupIsForum
func (s Supergroup) ToggleIsForum(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsForum(s.Id, s.IsForum, s.HasForumTabs)
}

// ToggleJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (s Supergroup) ToggleJoinByRequest(client *Client) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(s.Id, s.JoinByRequest)
}

// ToggleJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (s Supergroup) ToggleJoinToSendMessages(client *Client) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(s.Id, s.JoinToSendMessages)
}

// ToggleSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (s Supergroup) ToggleSignMessages(client *Client) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(s.Id, s.SignMessages, s.ShowMessageSender)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (s Supergroup) ToggleUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(s.Id, username, isActive)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (s SupergroupFullInfo) CreateNewSupergroupChat(client *Client, title string, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, isForum, isChannel, s.Description, messageAutoDeleteTime, forImport, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (s SupergroupFullInfo) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, s.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s SupergroupFullInfo) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, s.Description)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s SupergroupFullInfo) SetChatLocation(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatLocation(chatId, s.Location)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s SupergroupFullInfo) SetChatSlowModeDelay(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatSlowModeDelay(chatId, s.SlowModeDelay)
}

// SetMainProfileTab is a helper method for Client.SetMainProfileTab
func (s SupergroupFullInfo) SetMainProfileTab(client *Client) (*Ok, error) {
	return client.SetMainProfileTab(s.MainProfileTab)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (s SupergroupFullInfo) SetSupergroupCustomEmojiStickerSet(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(supergroupId, s.CustomEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (s SupergroupFullInfo) SetSupergroupMainProfileTab(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(supergroupId, s.MainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (s SupergroupFullInfo) SetSupergroupStickerSet(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupStickerSet(supergroupId, s.StickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (s SupergroupFullInfo) SetSupergroupUnrestrictBoostCount(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(supergroupId, s.UnrestrictBoostCount)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (s SupergroupFullInfo) ToggleSupergroupCanHaveSponsoredMessages(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(supergroupId, s.CanHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (s SupergroupFullInfo) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(supergroupId, s.HasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (s SupergroupFullInfo) ToggleSupergroupHasHiddenMembers(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(supergroupId, s.HasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (s SupergroupFullInfo) ToggleSupergroupIsAllHistoryAvailable(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(supergroupId, s.IsAllHistoryAvailable)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s SupergroupMembersFilterBanned) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, s.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s SupergroupMembersFilterBanned) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, s.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s SupergroupMembersFilterBanned) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, s.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s SupergroupMembersFilterBanned) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, s.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s SupergroupMembersFilterBanned) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, s.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (s SupergroupMembersFilterBanned) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(s.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (s SupergroupMembersFilterBanned) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(s.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (s SupergroupMembersFilterBanned) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, s.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s SupergroupMembersFilterBanned) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, s.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s SupergroupMembersFilterBanned) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, s.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s SupergroupMembersFilterBanned) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (s SupergroupMembersFilterBanned) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(s.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (s SupergroupMembersFilterBanned) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(s.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (s SupergroupMembersFilterBanned) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(s.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (s SupergroupMembersFilterBanned) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(s.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s SupergroupMembersFilterBanned) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, s.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (s SupergroupMembersFilterBanned) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(s.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (s SupergroupMembersFilterBanned) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(s.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (s SupergroupMembersFilterBanned) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(s.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SupergroupMembersFilterBanned) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(s.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (s SupergroupMembersFilterBanned) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(s.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (s SupergroupMembersFilterBanned) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s SupergroupMembersFilterBanned) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, s.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s SupergroupMembersFilterBanned) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, s.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s SupergroupMembersFilterBanned) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, s.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (s SupergroupMembersFilterBanned) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, s.Query, limit, returnNoneForEmptyQuery)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s SupergroupMembersFilterContacts) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, s.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s SupergroupMembersFilterContacts) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, s.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s SupergroupMembersFilterContacts) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, s.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s SupergroupMembersFilterContacts) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, s.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s SupergroupMembersFilterContacts) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, s.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (s SupergroupMembersFilterContacts) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(s.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (s SupergroupMembersFilterContacts) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(s.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (s SupergroupMembersFilterContacts) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, s.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s SupergroupMembersFilterContacts) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, s.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s SupergroupMembersFilterContacts) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, s.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s SupergroupMembersFilterContacts) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (s SupergroupMembersFilterContacts) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(s.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (s SupergroupMembersFilterContacts) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(s.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (s SupergroupMembersFilterContacts) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(s.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (s SupergroupMembersFilterContacts) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(s.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s SupergroupMembersFilterContacts) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, s.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (s SupergroupMembersFilterContacts) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(s.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (s SupergroupMembersFilterContacts) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(s.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (s SupergroupMembersFilterContacts) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(s.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SupergroupMembersFilterContacts) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(s.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (s SupergroupMembersFilterContacts) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(s.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (s SupergroupMembersFilterContacts) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s SupergroupMembersFilterContacts) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, s.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s SupergroupMembersFilterContacts) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, s.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s SupergroupMembersFilterContacts) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, s.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (s SupergroupMembersFilterContacts) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, s.Query, limit, returnNoneForEmptyQuery)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s SupergroupMembersFilterMention) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, s.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s SupergroupMembersFilterMention) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, s.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s SupergroupMembersFilterMention) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, s.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s SupergroupMembersFilterMention) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, s.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s SupergroupMembersFilterMention) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, s.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (s SupergroupMembersFilterMention) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(s.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (s SupergroupMembersFilterMention) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(s.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (s SupergroupMembersFilterMention) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, s.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s SupergroupMembersFilterMention) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, s.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s SupergroupMembersFilterMention) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, s.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s SupergroupMembersFilterMention) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (s SupergroupMembersFilterMention) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(s.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (s SupergroupMembersFilterMention) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(s.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (s SupergroupMembersFilterMention) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(s.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (s SupergroupMembersFilterMention) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(s.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s SupergroupMembersFilterMention) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, s.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (s SupergroupMembersFilterMention) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(s.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (s SupergroupMembersFilterMention) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(s.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (s SupergroupMembersFilterMention) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(s.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SupergroupMembersFilterMention) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(s.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (s SupergroupMembersFilterMention) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(s.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (s SupergroupMembersFilterMention) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s SupergroupMembersFilterMention) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, s.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s SupergroupMembersFilterMention) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, s.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s SupergroupMembersFilterMention) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, s.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (s SupergroupMembersFilterMention) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, s.Query, limit, returnNoneForEmptyQuery)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s SupergroupMembersFilterMention) SendChatAction(client *Client, chatId int64, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(chatId, s.TopicId, businessConnectionId, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s SupergroupMembersFilterRestricted) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, s.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s SupergroupMembersFilterRestricted) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, s.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s SupergroupMembersFilterRestricted) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, s.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s SupergroupMembersFilterRestricted) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, s.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s SupergroupMembersFilterRestricted) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, s.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (s SupergroupMembersFilterRestricted) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(s.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (s SupergroupMembersFilterRestricted) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(s.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (s SupergroupMembersFilterRestricted) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, s.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s SupergroupMembersFilterRestricted) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, s.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s SupergroupMembersFilterRestricted) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, s.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s SupergroupMembersFilterRestricted) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (s SupergroupMembersFilterRestricted) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(s.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (s SupergroupMembersFilterRestricted) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(s.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (s SupergroupMembersFilterRestricted) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(s.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (s SupergroupMembersFilterRestricted) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(s.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s SupergroupMembersFilterRestricted) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, s.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (s SupergroupMembersFilterRestricted) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(s.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (s SupergroupMembersFilterRestricted) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(s.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (s SupergroupMembersFilterRestricted) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(s.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SupergroupMembersFilterRestricted) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(s.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (s SupergroupMembersFilterRestricted) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(s.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (s SupergroupMembersFilterRestricted) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s SupergroupMembersFilterRestricted) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, s.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s SupergroupMembersFilterRestricted) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, s.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s SupergroupMembersFilterRestricted) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, s.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (s SupergroupMembersFilterRestricted) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, s.Query, limit, returnNoneForEmptyQuery)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s SupergroupMembersFilterSearch) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, s.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s SupergroupMembersFilterSearch) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, s.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s SupergroupMembersFilterSearch) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, s.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s SupergroupMembersFilterSearch) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, s.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s SupergroupMembersFilterSearch) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, s.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (s SupergroupMembersFilterSearch) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(s.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (s SupergroupMembersFilterSearch) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(s.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (s SupergroupMembersFilterSearch) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, s.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (s SupergroupMembersFilterSearch) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, s.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s SupergroupMembersFilterSearch) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, s.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s SupergroupMembersFilterSearch) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (s SupergroupMembersFilterSearch) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(s.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (s SupergroupMembersFilterSearch) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(s.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (s SupergroupMembersFilterSearch) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(s.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (s SupergroupMembersFilterSearch) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(s.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s SupergroupMembersFilterSearch) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, s.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (s SupergroupMembersFilterSearch) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(s.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (s SupergroupMembersFilterSearch) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(s.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (s SupergroupMembersFilterSearch) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(s.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s SupergroupMembersFilterSearch) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(s.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (s SupergroupMembersFilterSearch) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(s.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (s SupergroupMembersFilterSearch) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, s.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s SupergroupMembersFilterSearch) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, s.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s SupergroupMembersFilterSearch) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, s.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s SupergroupMembersFilterSearch) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, s.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (s SupergroupMembersFilterSearch) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, s.Query, limit, returnNoneForEmptyQuery)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TelegramPaymentPurposeGiftedStars) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TelegramPaymentPurposeGiftedStars) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (t TelegramPaymentPurposeGiftedStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, t.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (t TelegramPaymentPurposeGiftedStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, t.StarCount, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TelegramPaymentPurposeGiftedStars) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TelegramPaymentPurposeGiftedStars) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (t TelegramPaymentPurposeGiftedStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, t.StarCount)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TelegramPaymentPurposeGiftedStars) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposeGiftedStars) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// CreateCall is a helper method for Client.CreateCall
func (t TelegramPaymentPurposeGiftedStars) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TelegramPaymentPurposeGiftedStars) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TelegramPaymentPurposeGiftedStars) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TelegramPaymentPurposeGiftedStars) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (t TelegramPaymentPurposeGiftedStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, t.StarCount)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TelegramPaymentPurposeGiftedStars) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TelegramPaymentPurposeGiftedStars) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TelegramPaymentPurposeGiftedStars) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TelegramPaymentPurposeGiftedStars) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TelegramPaymentPurposeGiftedStars) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TelegramPaymentPurposeGiftedStars) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TelegramPaymentPurposeGiftedStars) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (t TelegramPaymentPurposeGiftedStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, t.StarCount, password)
}

// GetUser is a helper method for Client.GetUser
func (t TelegramPaymentPurposeGiftedStars) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TelegramPaymentPurposeGiftedStars) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TelegramPaymentPurposeGiftedStars) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TelegramPaymentPurposeGiftedStars) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TelegramPaymentPurposeGiftedStars) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TelegramPaymentPurposeGiftedStars) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposeGiftedStars) GiftPremiumWithStars(client *Client, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, t.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (t TelegramPaymentPurposeGiftedStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, t.StarCount)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TelegramPaymentPurposeGiftedStars) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (t TelegramPaymentPurposeGiftedStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, t.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TelegramPaymentPurposeGiftedStars) PlaceGiftAuctionBid(client *Client, giftId string, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, t.StarCount, t.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TelegramPaymentPurposeGiftedStars) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TelegramPaymentPurposeGiftedStars) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TelegramPaymentPurposeGiftedStars) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TelegramPaymentPurposeGiftedStars) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (t TelegramPaymentPurposeGiftedStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, t.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposeGiftedStars) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TelegramPaymentPurposeGiftedStars) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TelegramPaymentPurposeGiftedStars) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TelegramPaymentPurposeGiftedStars) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TelegramPaymentPurposeGiftedStars) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TelegramPaymentPurposeGiftedStars) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TelegramPaymentPurposeGiftedStars) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TelegramPaymentPurposeGiftedStars) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TelegramPaymentPurposeGiftedStars) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TelegramPaymentPurposeGiftedStars) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TelegramPaymentPurposeGiftedStars) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TelegramPaymentPurposeGiftedStars) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TelegramPaymentPurposeGiftedStars) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (t TelegramPaymentPurposeGiftedStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, t.StarCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TelegramPaymentPurposeGiftedStars) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (t TelegramPaymentPurposeGiftedStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, t.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (t TelegramPaymentPurposeGiftedStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, t.StarCount)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TelegramPaymentPurposeGiftedStars) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (t TelegramPaymentPurposeJoinChat) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(t.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (t TelegramPaymentPurposeJoinChat) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(t.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (t TelegramPaymentPurposeJoinChat) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(t.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (t TelegramPaymentPurposeJoinChat) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, t.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (t TelegramPaymentPurposeJoinChat) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, t.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (t TelegramPaymentPurposeJoinChat) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, t.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (t TelegramPaymentPurposeJoinChat) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, t.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (t TelegramPaymentPurposeJoinChat) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, t.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (t TelegramPaymentPurposeJoinChat) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, t.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (t TelegramPaymentPurposeJoinChat) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, t.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (t TelegramPaymentPurposeJoinChat) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, t.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (t TelegramPaymentPurposeJoinChat) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, t.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (t TelegramPaymentPurposeJoinChat) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(t.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (t TelegramPaymentPurposeJoinChat) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, t.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (t TelegramPaymentPurposeJoinChat) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, t.InviteLink)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TelegramPaymentPurposePremiumGift) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TelegramPaymentPurposePremiumGift) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TelegramPaymentPurposePremiumGift) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TelegramPaymentPurposePremiumGift) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TelegramPaymentPurposePremiumGift) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposePremiumGift) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// CreateCall is a helper method for Client.CreateCall
func (t TelegramPaymentPurposePremiumGift) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TelegramPaymentPurposePremiumGift) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TelegramPaymentPurposePremiumGift) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TelegramPaymentPurposePremiumGift) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TelegramPaymentPurposePremiumGift) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TelegramPaymentPurposePremiumGift) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TelegramPaymentPurposePremiumGift) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TelegramPaymentPurposePremiumGift) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (t TelegramPaymentPurposePremiumGift) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(t.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (t TelegramPaymentPurposePremiumGift) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(t.Text)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TelegramPaymentPurposePremiumGift) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TelegramPaymentPurposePremiumGift) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (t TelegramPaymentPurposePremiumGift) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(t.MonthCount)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TelegramPaymentPurposePremiumGift) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetUser is a helper method for Client.GetUser
func (t TelegramPaymentPurposePremiumGift) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TelegramPaymentPurposePremiumGift) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TelegramPaymentPurposePremiumGift) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TelegramPaymentPurposePremiumGift) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TelegramPaymentPurposePremiumGift) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TelegramPaymentPurposePremiumGift) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposePremiumGift) GiftPremiumWithStars(client *Client, starCount int64) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, starCount, t.MonthCount, t.Text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TelegramPaymentPurposePremiumGift) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (t TelegramPaymentPurposePremiumGift) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(t.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TelegramPaymentPurposePremiumGift) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, t.UserId, t.Text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TelegramPaymentPurposePremiumGift) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TelegramPaymentPurposePremiumGift) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TelegramPaymentPurposePremiumGift) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TelegramPaymentPurposePremiumGift) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SearchQuote is a helper method for Client.SearchQuote
func (t TelegramPaymentPurposePremiumGift) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(t.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (t TelegramPaymentPurposePremiumGift) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, t.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (t TelegramPaymentPurposePremiumGift) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, t.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TelegramPaymentPurposePremiumGift) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, t.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposePremiumGift) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TelegramPaymentPurposePremiumGift) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TelegramPaymentPurposePremiumGift) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TelegramPaymentPurposePremiumGift) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TelegramPaymentPurposePremiumGift) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TelegramPaymentPurposePremiumGift) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TelegramPaymentPurposePremiumGift) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TelegramPaymentPurposePremiumGift) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TelegramPaymentPurposePremiumGift) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TelegramPaymentPurposePremiumGift) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TelegramPaymentPurposePremiumGift) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TelegramPaymentPurposePremiumGift) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TelegramPaymentPurposePremiumGift) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TelegramPaymentPurposePremiumGift) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// TranslateText is a helper method for Client.TranslateText
func (t TelegramPaymentPurposePremiumGift) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(t.Text, toLanguageCode)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TelegramPaymentPurposePremiumGift) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (t TelegramPaymentPurposePremiumGiftCodes) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, t.UserIds)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposePremiumGiftCodes) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (t TelegramPaymentPurposePremiumGiftCodes) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(t.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (t TelegramPaymentPurposePremiumGiftCodes) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, t.UserIds, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (t TelegramPaymentPurposePremiumGiftCodes) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(t.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (t TelegramPaymentPurposePremiumGiftCodes) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(t.Text)
}

// GetPremiumGiveawayPaymentOptions is a helper method for Client.GetPremiumGiveawayPaymentOptions
func (t TelegramPaymentPurposePremiumGiftCodes) GetPremiumGiveawayPaymentOptions(client *Client) (*PremiumGiveawayPaymentOptions, error) {
	return client.GetPremiumGiveawayPaymentOptions(t.BoostedChatId)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (t TelegramPaymentPurposePremiumGiftCodes) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(t.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposePremiumGiftCodes) GiftPremiumWithStars(client *Client, userId int64, starCount int64) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, t.MonthCount, t.Text)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (t TelegramPaymentPurposePremiumGiftCodes) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, t.UserIds)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (t TelegramPaymentPurposePremiumGiftCodes) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(t.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TelegramPaymentPurposePremiumGiftCodes) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, t.Text, isPrivate)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (t TelegramPaymentPurposePremiumGiftCodes) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(t.UserIds)
}

// SearchQuote is a helper method for Client.SearchQuote
func (t TelegramPaymentPurposePremiumGiftCodes) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(t.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (t TelegramPaymentPurposePremiumGiftCodes) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, t.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (t TelegramPaymentPurposePremiumGiftCodes) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, t.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TelegramPaymentPurposePremiumGiftCodes) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, t.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposePremiumGiftCodes) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (t TelegramPaymentPurposePremiumGiftCodes) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(t.UserIds)
}

// TranslateText is a helper method for Client.TranslateText
func (t TelegramPaymentPurposePremiumGiftCodes) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(t.Text, toLanguageCode)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposePremiumGiveaway) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (t TelegramPaymentPurposePremiumGiveaway) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(t.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposePremiumGiveaway) GiftPremiumWithStars(client *Client, userId int64, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, t.MonthCount, text)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (t TelegramPaymentPurposePremiumGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, t.Parameters, t.WinnerCount, starCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposePremiumGiveaway) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (t TelegramPaymentPurposeStarGiveaway) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, t.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (t TelegramPaymentPurposeStarGiveaway) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, t.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (t TelegramPaymentPurposeStarGiveaway) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, t.StarCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposeStarGiveaway) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (t TelegramPaymentPurposeStarGiveaway) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, t.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (t TelegramPaymentPurposeStarGiveaway) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, t.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposeStarGiveaway) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, t.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (t TelegramPaymentPurposeStarGiveaway) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, t.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (t TelegramPaymentPurposeStarGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, t.Parameters, t.WinnerCount, t.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TelegramPaymentPurposeStarGiveaway) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, t.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (t TelegramPaymentPurposeStarGiveaway) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, t.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposeStarGiveaway) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (t TelegramPaymentPurposeStarGiveaway) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, t.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (t TelegramPaymentPurposeStarGiveaway) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, t.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (t TelegramPaymentPurposeStarGiveaway) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, t.StarCount)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TelegramPaymentPurposeStars) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(t.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (t TelegramPaymentPurposeStars) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(t.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (t TelegramPaymentPurposeStars) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(t.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (t TelegramPaymentPurposeStars) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(t.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (t TelegramPaymentPurposeStars) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, t.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (t TelegramPaymentPurposeStars) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(t.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (t TelegramPaymentPurposeStars) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(t.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (t TelegramPaymentPurposeStars) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(t.ChatId, messageId, options)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (t TelegramPaymentPurposeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, t.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (t TelegramPaymentPurposeStars) AddPendingPaidMessageReaction(client *Client, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(t.ChatId, messageId, t.StarCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (t TelegramPaymentPurposeStars) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(t.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (t TelegramPaymentPurposeStars) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (t TelegramPaymentPurposeStars) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(t.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (t TelegramPaymentPurposeStars) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(t.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (t TelegramPaymentPurposeStars) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(t.ChatId, slotIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (t TelegramPaymentPurposeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, t.StarCount)
}

// CanPostStory is a helper method for Client.CanPostStory
func (t TelegramPaymentPurposeStars) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(t.ChatId)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (t TelegramPaymentPurposeStars) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(t.Currency, t.Amount)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (t TelegramPaymentPurposeStars) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(t.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (t TelegramPaymentPurposeStars) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(t.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (t TelegramPaymentPurposeStars) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(t.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (t TelegramPaymentPurposeStars) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(t.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (t TelegramPaymentPurposeStars) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(t.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (t TelegramPaymentPurposeStars) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(t.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (t TelegramPaymentPurposeStars) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(t.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (t TelegramPaymentPurposeStars) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(t.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (t TelegramPaymentPurposeStars) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(t.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (t TelegramPaymentPurposeStars) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(t.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (t TelegramPaymentPurposeStars) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(t.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (t TelegramPaymentPurposeStars) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(t.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (t TelegramPaymentPurposeStars) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(t.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (t TelegramPaymentPurposeStars) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(t.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (t TelegramPaymentPurposeStars) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(t.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (t TelegramPaymentPurposeStars) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(t.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (t TelegramPaymentPurposeStars) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(t.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (t TelegramPaymentPurposeStars) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(t.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (t TelegramPaymentPurposeStars) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(t.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (t TelegramPaymentPurposeStars) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(t.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (t TelegramPaymentPurposeStars) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(t.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (t TelegramPaymentPurposeStars) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(t.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (t TelegramPaymentPurposeStars) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(t.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (t TelegramPaymentPurposeStars) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(t.ChatId, storyAlbumId)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (t TelegramPaymentPurposeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, t.StarCount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (t TelegramPaymentPurposeStars) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, t.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (t TelegramPaymentPurposeStars) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, t.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (t TelegramPaymentPurposeStars) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, t.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (t TelegramPaymentPurposeStars) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, t.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (t TelegramPaymentPurposeStars) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, t.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (t TelegramPaymentPurposeStars) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, t.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (t TelegramPaymentPurposeStars) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(t.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (t TelegramPaymentPurposeStars) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(t.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (t TelegramPaymentPurposeStars) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(t.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (t TelegramPaymentPurposeStars) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(t.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (t TelegramPaymentPurposeStars) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(t.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (t TelegramPaymentPurposeStars) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(t.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (t TelegramPaymentPurposeStars) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(t.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (t TelegramPaymentPurposeStars) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(t.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (t TelegramPaymentPurposeStars) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(t.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (t TelegramPaymentPurposeStars) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(t.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (t TelegramPaymentPurposeStars) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(t.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (t TelegramPaymentPurposeStars) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, t.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (t TelegramPaymentPurposeStars) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(t.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (t TelegramPaymentPurposeStars) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(t.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (t TelegramPaymentPurposeStars) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(t.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (t TelegramPaymentPurposeStars) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(t.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (t TelegramPaymentPurposeStars) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(t.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (t TelegramPaymentPurposeStars) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(t.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (t TelegramPaymentPurposeStars) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(t.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (t TelegramPaymentPurposeStars) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(t.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (t TelegramPaymentPurposeStars) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(t.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (t TelegramPaymentPurposeStars) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(t.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (t TelegramPaymentPurposeStars) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(t.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (t TelegramPaymentPurposeStars) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(t.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (t TelegramPaymentPurposeStars) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(t.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (t TelegramPaymentPurposeStars) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(t.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (t TelegramPaymentPurposeStars) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(t.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (t TelegramPaymentPurposeStars) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(t.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (t TelegramPaymentPurposeStars) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(t.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (t TelegramPaymentPurposeStars) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(t.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (t TelegramPaymentPurposeStars) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(t.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (t TelegramPaymentPurposeStars) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(t.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (t TelegramPaymentPurposeStars) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(t.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (t TelegramPaymentPurposeStars) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(t.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (t TelegramPaymentPurposeStars) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(t.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (t TelegramPaymentPurposeStars) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(t.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (t TelegramPaymentPurposeStars) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(t.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (t TelegramPaymentPurposeStars) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(t.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (t TelegramPaymentPurposeStars) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(t.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (t TelegramPaymentPurposeStars) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(t.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (t TelegramPaymentPurposeStars) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(t.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (t TelegramPaymentPurposeStars) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(t.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (t TelegramPaymentPurposeStars) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(t.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (t TelegramPaymentPurposeStars) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(t.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (t TelegramPaymentPurposeStars) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(t.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (t TelegramPaymentPurposeStars) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(t.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (t TelegramPaymentPurposeStars) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(t.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (t TelegramPaymentPurposeStars) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(t.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (t TelegramPaymentPurposeStars) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(t.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (t TelegramPaymentPurposeStars) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(t.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (t TelegramPaymentPurposeStars) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(t.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (t TelegramPaymentPurposeStars) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(t.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (t TelegramPaymentPurposeStars) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(t.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (t TelegramPaymentPurposeStars) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(t.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (t TelegramPaymentPurposeStars) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(t.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (t TelegramPaymentPurposeStars) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(t.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TelegramPaymentPurposeStars) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(t.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (t TelegramPaymentPurposeStars) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(t.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (t TelegramPaymentPurposeStars) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, t.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (t TelegramPaymentPurposeStars) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(t.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (t TelegramPaymentPurposeStars) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(t.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (t TelegramPaymentPurposeStars) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(t.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (t TelegramPaymentPurposeStars) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(t.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (t TelegramPaymentPurposeStars) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, t.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (t TelegramPaymentPurposeStars) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(t.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (t TelegramPaymentPurposeStars) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(t.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (t TelegramPaymentPurposeStars) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(t.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (t TelegramPaymentPurposeStars) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(t.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (t TelegramPaymentPurposeStars) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(t.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (t TelegramPaymentPurposeStars) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(t.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (t TelegramPaymentPurposeStars) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(t.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (t TelegramPaymentPurposeStars) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(t.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (t TelegramPaymentPurposeStars) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(t.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (t TelegramPaymentPurposeStars) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(t.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (t TelegramPaymentPurposeStars) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(t.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (t TelegramPaymentPurposeStars) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(t.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (t TelegramPaymentPurposeStars) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(t.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (t TelegramPaymentPurposeStars) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(t.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (t TelegramPaymentPurposeStars) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(t.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (t TelegramPaymentPurposeStars) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(t.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (t TelegramPaymentPurposeStars) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(t.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (t TelegramPaymentPurposeStars) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(t.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (t TelegramPaymentPurposeStars) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(t.ChatId, messageId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (t TelegramPaymentPurposeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, t.StarCount, password)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (t TelegramPaymentPurposeStars) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(t.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (t TelegramPaymentPurposeStars) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, t.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (t TelegramPaymentPurposeStars) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(t.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (t TelegramPaymentPurposeStars) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(t.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TelegramPaymentPurposeStars) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(t.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (t TelegramPaymentPurposeStars) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(t.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (t TelegramPaymentPurposeStars) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(t.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (t TelegramPaymentPurposeStars) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(t.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (t TelegramPaymentPurposeStars) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(t.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TelegramPaymentPurposeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, t.StarCount, monthCount, text)
}

// ImportMessages is a helper method for Client.ImportMessages
func (t TelegramPaymentPurposeStars) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(t.ChatId, messageFile, attachedFiles)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (t TelegramPaymentPurposeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, t.StarCount)
}

// JoinChat is a helper method for Client.JoinChat
func (t TelegramPaymentPurposeStars) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(t.ChatId)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (t TelegramPaymentPurposeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, t.StarCount)
}

// LeaveChat is a helper method for Client.LeaveChat
func (t TelegramPaymentPurposeStars) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(t.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (t TelegramPaymentPurposeStars) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(t.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (t TelegramPaymentPurposeStars) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(t.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (t TelegramPaymentPurposeStars) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(t.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (t TelegramPaymentPurposeStars) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(t.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (t TelegramPaymentPurposeStars) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(t.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (t TelegramPaymentPurposeStars) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(t.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (t TelegramPaymentPurposeStars) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(t.ChatId, messageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TelegramPaymentPurposeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, t.StarCount, userId, text, isPrivate)
}

// PostStory is a helper method for Client.PostStory
func (t TelegramPaymentPurposeStars) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(t.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TelegramPaymentPurposeStars) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(t.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (t TelegramPaymentPurposeStars) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(t.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (t TelegramPaymentPurposeStars) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(t.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (t TelegramPaymentPurposeStars) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(t.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (t TelegramPaymentPurposeStars) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(t.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (t TelegramPaymentPurposeStars) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(t.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (t TelegramPaymentPurposeStars) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(t.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (t TelegramPaymentPurposeStars) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(t.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (t TelegramPaymentPurposeStars) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, t.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (t TelegramPaymentPurposeStars) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(t.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (t TelegramPaymentPurposeStars) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(t.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (t TelegramPaymentPurposeStars) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(t.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (t TelegramPaymentPurposeStars) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(t.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (t TelegramPaymentPurposeStars) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(t.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (t TelegramPaymentPurposeStars) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(t.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (t TelegramPaymentPurposeStars) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (t TelegramPaymentPurposeStars) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, t.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (t TelegramPaymentPurposeStars) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(t.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (t TelegramPaymentPurposeStars) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (t TelegramPaymentPurposeStars) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(t.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (t TelegramPaymentPurposeStars) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(t.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (t TelegramPaymentPurposeStars) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(t.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (t TelegramPaymentPurposeStars) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(t.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (t TelegramPaymentPurposeStars) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(t.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (t TelegramPaymentPurposeStars) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(t.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (t TelegramPaymentPurposeStars) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(t.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (t TelegramPaymentPurposeStars) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(t.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (t TelegramPaymentPurposeStars) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(t.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (t TelegramPaymentPurposeStars) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, t.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (t TelegramPaymentPurposeStars) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(t.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (t TelegramPaymentPurposeStars) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(t.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (t TelegramPaymentPurposeStars) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(t.ChatId, limit)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (t TelegramPaymentPurposeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, t.StarCount)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (t TelegramPaymentPurposeStars) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(t.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (t TelegramPaymentPurposeStars) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, t.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (t TelegramPaymentPurposeStars) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, t.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (t TelegramPaymentPurposeStars) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, t.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (t TelegramPaymentPurposeStars) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(t.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (t TelegramPaymentPurposeStars) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(t.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (t TelegramPaymentPurposeStars) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(t.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (t TelegramPaymentPurposeStars) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(t.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (t TelegramPaymentPurposeStars) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(t.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TelegramPaymentPurposeStars) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(t.ChatId, forumTopicId, draftId, text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (t TelegramPaymentPurposeStars) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, t.Currency, t.Amount)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (t TelegramPaymentPurposeStars) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, t.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (t TelegramPaymentPurposeStars) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(t.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (t TelegramPaymentPurposeStars) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(t.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (t TelegramPaymentPurposeStars) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(t.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (t TelegramPaymentPurposeStars) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(t.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (t TelegramPaymentPurposeStars) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(t.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (t TelegramPaymentPurposeStars) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(t.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (t TelegramPaymentPurposeStars) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(t.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (t TelegramPaymentPurposeStars) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(t.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (t TelegramPaymentPurposeStars) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(t.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (t TelegramPaymentPurposeStars) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(t.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (t TelegramPaymentPurposeStars) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(t.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (t TelegramPaymentPurposeStars) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(t.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (t TelegramPaymentPurposeStars) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(t.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (t TelegramPaymentPurposeStars) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(t.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (t TelegramPaymentPurposeStars) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(t.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (t TelegramPaymentPurposeStars) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(t.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (t TelegramPaymentPurposeStars) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(t.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (t TelegramPaymentPurposeStars) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(t.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (t TelegramPaymentPurposeStars) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(t.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (t TelegramPaymentPurposeStars) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(t.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (t TelegramPaymentPurposeStars) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(t.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (t TelegramPaymentPurposeStars) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(t.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (t TelegramPaymentPurposeStars) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(t.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (t TelegramPaymentPurposeStars) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(t.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (t TelegramPaymentPurposeStars) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(t.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (t TelegramPaymentPurposeStars) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(t.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TelegramPaymentPurposeStars) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(t.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (t TelegramPaymentPurposeStars) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(t.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (t TelegramPaymentPurposeStars) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(t.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (t TelegramPaymentPurposeStars) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(t.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (t TelegramPaymentPurposeStars) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(t.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (t TelegramPaymentPurposeStars) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(t.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (t TelegramPaymentPurposeStars) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(t.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (t TelegramPaymentPurposeStars) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(t.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (t TelegramPaymentPurposeStars) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(t.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (t TelegramPaymentPurposeStars) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(t.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (t TelegramPaymentPurposeStars) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(t.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (t TelegramPaymentPurposeStars) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(t.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (t TelegramPaymentPurposeStars) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, t.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (t TelegramPaymentPurposeStars) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(t.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (t TelegramPaymentPurposeStars) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(t.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (t TelegramPaymentPurposeStars) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(t.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (t TelegramPaymentPurposeStars) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(t.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (t TelegramPaymentPurposeStars) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(t.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (t TelegramPaymentPurposeStars) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(t.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (t TelegramPaymentPurposeStars) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(t.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (t TelegramPaymentPurposeStars) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, t.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (t TelegramPaymentPurposeStars) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(t.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (t TelegramPaymentPurposeStars) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(t.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (t TelegramPaymentPurposeStars) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(t.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (t TelegramPaymentPurposeStars) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(t.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (t TelegramPaymentPurposeStars) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(t.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (t TelegramPaymentPurposeStars) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(t.ChatId, isHidden)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (t TelegramPaymentPurposeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, t.StarCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TelegramPaymentPurposeStars) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(t.ChatId, userId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (t TelegramPaymentPurposeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, t.StarCount)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (t TelegramPaymentPurposeStars) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(t.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (t TelegramPaymentPurposeStars) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(t.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (t TelegramPaymentPurposeStars) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(t.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (t TelegramPaymentPurposeStars) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(t.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (t TelegramPaymentPurposeStars) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(t.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (t TelegramPaymentPurposeStars) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(t.ChatId)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (t TelegramPaymentPurposeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, t.StarCount)
}

// ViewMessages is a helper method for Client.ViewMessages
func (t TelegramPaymentPurposeStars) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(t.ChatId, messageIds, forceRead, opts)
}

// CreateTemporaryPassword is a helper method for Client.CreateTemporaryPassword
func (t TemporaryPasswordState) CreateTemporaryPassword(client *Client, password string) (*TemporaryPasswordState, error) {
	return client.CreateTemporaryPassword(password, t.ValidFor)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (t TermsOfService) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(t.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (t TermsOfService) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(t.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TermsOfService) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, t.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (t TermsOfService) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(t.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TermsOfService) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, t.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (t TermsOfService) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(t.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (t TermsOfService) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, t.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (t TermsOfService) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, t.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TermsOfService) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, t.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (t TermsOfService) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(t.Text, toLanguageCode)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (t Text) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, t.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (t Text) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, t.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (t Text) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(t.Text, inputLanguageCodes)
}

// GetEntities is a helper method for Client.GetTextEntities
func (t Text) GetEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(t.Text)
}

// ParseEntities is a helper method for Client.ParseTextEntities
func (t Text) ParseEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(t.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (t Text) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, t.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (t Text) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, t.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (t Text) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, t.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (t Text) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(t.Text, inputLanguageCodes)
}

// GetBlockedMessageSenders is a helper method for Client.GetBlockedMessageSenders
func (t TextEntity) GetBlockedMessageSenders(client *Client, blockList *BlockList, limit int32) (*MessageSenders, error) {
	return client.GetBlockedMessageSenders(blockList, t.Offset, limit)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (t TextEntity) GetChatHistory(client *Client, chatId int64, fromMessageId int64, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(chatId, fromMessageId, t.Offset, limit, onlyLocal)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (t TextEntity) GetDirectMessagesChatTopicHistory(client *Client, chatId int64, topicId int64, fromMessageId int64, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(chatId, topicId, fromMessageId, t.Offset, limit)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (t TextEntity) GetForumTopicHistory(client *Client, chatId int64, forumTopicId int32, fromMessageId int64, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(chatId, forumTopicId, fromMessageId, t.Offset, limit)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (t TextEntity) GetMessageThreadHistory(client *Client, chatId int64, messageId int64, fromMessageId int64, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, messageId, fromMessageId, t.Offset, limit)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (t TextEntity) GetPollVoters(client *Client, chatId int64, messageId int64, optionId int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, messageId, optionId, t.Offset, limit)
}

// GetSavedMessagesTopicHistory is a helper method for Client.GetSavedMessagesTopicHistory
func (t TextEntity) GetSavedMessagesTopicHistory(client *Client, savedMessagesTopicId int64, fromMessageId int64, limit int32) (*Messages, error) {
	return client.GetSavedMessagesTopicHistory(savedMessagesTopicId, fromMessageId, t.Offset, limit)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (t TextEntity) GetStoryAlbumStories(client *Client, chatId int64, storyAlbumId int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(chatId, storyAlbumId, t.Offset, limit)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (t TextEntity) GetSupergroupMembers(client *Client, supergroupId int64, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(supergroupId, t.Offset, limit, opts)
}

// GetTrendingStickerSets is a helper method for Client.GetTrendingStickerSets
func (t TextEntity) GetTrendingStickerSets(client *Client, stickerType *StickerType, limit int32) (*TrendingStickerSets, error) {
	return client.GetTrendingStickerSets(stickerType, t.Offset, limit)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TextEntity) GetUserProfileAudios(client *Client, userId int64, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(userId, t.Offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TextEntity) GetUserProfilePhotos(client *Client, userId int64, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(userId, t.Offset, limit)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (t TextEntity) SearchChatMessages(client *Client, chatId int64, query string, fromMessageId int64, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, query, fromMessageId, t.Offset, limit, opts)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (t TextEntity) SearchSavedMessages(client *Client, savedMessagesTopicId int64, query string, fromMessageId int64, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, query, fromMessageId, t.Offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (t TextEntity) SearchStickers(client *Client, stickerType *StickerType, emojis string, query string, inputLanguageCodes []string, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, query, inputLanguageCodes, t.Offset, limit)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (t TextEntityTypeCustomEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, name string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(name, t.CustomEmojiId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (t TextEntityTypeMediaTimestamp) GetMessageLink(client *Client, chatId int64, messageId int64, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, messageId, t.MediaTimestamp, forAlbum, inMessageThread)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TextEntityTypeMentionName) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TextEntityTypeMentionName) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TextEntityTypeMentionName) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TextEntityTypeMentionName) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TextEntityTypeMentionName) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (t TextEntityTypeMentionName) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TextEntityTypeMentionName) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TextEntityTypeMentionName) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TextEntityTypeMentionName) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TextEntityTypeMentionName) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TextEntityTypeMentionName) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TextEntityTypeMentionName) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TextEntityTypeMentionName) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TextEntityTypeMentionName) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TextEntityTypeMentionName) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TextEntityTypeMentionName) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetUser is a helper method for Client.GetUser
func (t TextEntityTypeMentionName) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TextEntityTypeMentionName) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TextEntityTypeMentionName) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TextEntityTypeMentionName) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TextEntityTypeMentionName) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TextEntityTypeMentionName) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TextEntityTypeMentionName) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TextEntityTypeMentionName) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TextEntityTypeMentionName) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, t.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TextEntityTypeMentionName) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TextEntityTypeMentionName) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TextEntityTypeMentionName) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TextEntityTypeMentionName) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TextEntityTypeMentionName) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TextEntityTypeMentionName) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TextEntityTypeMentionName) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TextEntityTypeMentionName) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TextEntityTypeMentionName) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TextEntityTypeMentionName) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TextEntityTypeMentionName) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TextEntityTypeMentionName) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TextEntityTypeMentionName) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TextEntityTypeMentionName) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TextEntityTypeMentionName) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TextEntityTypeMentionName) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TextEntityTypeMentionName) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TextEntityTypeMentionName) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (t TextEntityTypeTextUrl) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, t.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (t TextEntityTypeTextUrl) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, t.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (t TextEntityTypeTextUrl) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, t.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (t TextEntityTypeTextUrl) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(t.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (t TextEntityTypeTextUrl) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(t.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (t TextEntityTypeTextUrl) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, t.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (t TextEntityTypeTextUrl) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(t.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (t TextEntityTypeTextUrl) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, t.Url, parameters, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (t TextQuote) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(t.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (t TextQuote) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(t.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TextQuote) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, t.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (t TextQuote) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(t.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TextQuote) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, t.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (t TextQuote) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(t.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (t TextQuote) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, t.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (t TextQuote) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, t.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (t TextQuote) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, t.Text)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (t TextQuote) SetStickerPositionInSet(client *Client, sticker *InputFile) (*Ok, error) {
	return client.SetStickerPositionInSet(sticker, t.Position)
}

// TranslateText is a helper method for Client.TranslateText
func (t TextQuote) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(t.Text, toLanguageCode)
}

// GetMapFile is a helper method for Client.GetMapThumbnailFile
func (t Thumbnail) GetMapFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, t.Width, t.Height, scale, chatId)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TimeZone) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, t.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (t TimeZone) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(t.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (t TimeZone) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(t.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (t TimeZone) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, t.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (t TimeZone) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, t.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (t TimeZone) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, t.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (t TimeZone) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, t.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (t TimeZone) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, t.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TimeZone) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, t.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (t TimeZone) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, t.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (t TimeZone) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(t.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (t TimeZone) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, t.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (t TimeZone) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, t.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (t TimeZone) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, t.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (t TimeZone) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, t.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (t TimeZone) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(t.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (t TimeZone) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(t.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (t TimeZone) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(t.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (t TimeZone) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(t.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TimeZone) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, t.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (t TimeZone) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(t.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (t TimeZone) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(t.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (t TimeZone) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, t.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (t TimeZone) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(t.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (t TimeZone) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, t.Name)
}

// SetOption is a helper method for Client.SetOption
func (t TimeZone) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(t.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (t TimeZone) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, t.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TimeZone) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, t.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (t TimeZone) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(t.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (t TimeZone) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, t.Name)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (t TMeUrl) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, t.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (t TMeUrl) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, t.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (t TMeUrl) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, t.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (t TMeUrl) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(t.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (t TMeUrl) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(t.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (t TMeUrl) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, t.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (t TMeUrl) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(t.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (t TMeUrl) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, t.Url, parameters, opts)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (t TMeUrlTypeStickerSet) SetSupergroupStickerSet(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupStickerSet(supergroupId, t.StickerSetId)
}

// CreateSupergroupChat is a helper method for Client.CreateSupergroupChat
func (t TMeUrlTypeSupergroup) CreateSupergroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(t.SupergroupId, force)
}

// DisableAllSupergroupUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (t TMeUrlTypeSupergroup) DisableAllSupergroupUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(t.SupergroupId)
}

// GetSupergroup is a helper method for Client.GetSupergroup
func (t TMeUrlTypeSupergroup) GetSupergroup(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(t.SupergroupId)
}

// GetSupergroupFullInfo is a helper method for Client.GetSupergroupFullInfo
func (t TMeUrlTypeSupergroup) GetSupergroupFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(t.SupergroupId)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (t TMeUrlTypeSupergroup) GetSupergroupMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(t.SupergroupId, offset, limit, opts)
}

// ReorderSupergroupActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (t TMeUrlTypeSupergroup) ReorderSupergroupActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(t.SupergroupId, usernames)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (t TMeUrlTypeSupergroup) ReportSupergroupAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(t.SupergroupId, messageId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (t TMeUrlTypeSupergroup) ReportSupergroupSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(t.SupergroupId, messageIds)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (t TMeUrlTypeSupergroup) SetSupergroupCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(t.SupergroupId, customEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (t TMeUrlTypeSupergroup) SetSupergroupMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(t.SupergroupId, mainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (t TMeUrlTypeSupergroup) SetSupergroupStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(t.SupergroupId, stickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (t TMeUrlTypeSupergroup) SetSupergroupUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(t.SupergroupId, unrestrictBoostCount)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (t TMeUrlTypeSupergroup) SetSupergroupUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(t.SupergroupId, username)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (t TMeUrlTypeSupergroup) ToggleSupergroupCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(t.SupergroupId, canHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (t TMeUrlTypeSupergroup) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(t.SupergroupId, hasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (t TMeUrlTypeSupergroup) ToggleSupergroupHasAutomaticTranslation(client *Client, hasAutomaticTranslation bool) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(t.SupergroupId, hasAutomaticTranslation)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (t TMeUrlTypeSupergroup) ToggleSupergroupHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(t.SupergroupId, hasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (t TMeUrlTypeSupergroup) ToggleSupergroupIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(t.SupergroupId, isAllHistoryAvailable)
}

// ToggleSupergroupIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (t TMeUrlTypeSupergroup) ToggleSupergroupIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(t.SupergroupId)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (t TMeUrlTypeSupergroup) ToggleSupergroupIsForum(client *Client, isForum bool, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(t.SupergroupId, isForum, hasForumTabs)
}

// ToggleSupergroupJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (t TMeUrlTypeSupergroup) ToggleSupergroupJoinByRequest(client *Client, joinByRequest bool) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(t.SupergroupId, joinByRequest)
}

// ToggleSupergroupJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (t TMeUrlTypeSupergroup) ToggleSupergroupJoinToSendMessages(client *Client, joinToSendMessages bool) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(t.SupergroupId, joinToSendMessages)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (t TMeUrlTypeSupergroup) ToggleSupergroupSignMessages(client *Client, signMessages bool, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(t.SupergroupId, signMessages, showMessageSender)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (t TMeUrlTypeSupergroup) ToggleSupergroupUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(t.SupergroupId, username, isActive)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TMeUrlTypeUser) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, t.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (t TMeUrlTypeUser) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(t.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (t TMeUrlTypeUser) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(t.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (t TMeUrlTypeUser) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(t.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (t TMeUrlTypeUser) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(t.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (t TMeUrlTypeUser) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(t.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (t TMeUrlTypeUser) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(t.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (t TMeUrlTypeUser) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(t.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (t TMeUrlTypeUser) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(t.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (t TMeUrlTypeUser) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(t.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (t TMeUrlTypeUser) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, t.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (t TMeUrlTypeUser) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(t.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (t TMeUrlTypeUser) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, t.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (t TMeUrlTypeUser) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(t.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (t TMeUrlTypeUser) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(t.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (t TMeUrlTypeUser) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(t.UserId)
}

// GetUser is a helper method for Client.GetUser
func (t TMeUrlTypeUser) GetUser(client *Client) (*User, error) {
	return client.GetUser(t.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (t TMeUrlTypeUser) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, t.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (t TMeUrlTypeUser) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(t.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (t TMeUrlTypeUser) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(t.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (t TMeUrlTypeUser) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(t.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (t TMeUrlTypeUser) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(t.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (t TMeUrlTypeUser) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(t.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (t TMeUrlTypeUser) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, t.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (t TMeUrlTypeUser) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, t.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (t TMeUrlTypeUser) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, t.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (t TMeUrlTypeUser) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(t.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (t TMeUrlTypeUser) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(t.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (t TMeUrlTypeUser) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(t.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (t TMeUrlTypeUser) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, t.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (t TMeUrlTypeUser) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, t.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (t TMeUrlTypeUser) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(t.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (t TMeUrlTypeUser) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(t.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (t TMeUrlTypeUser) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(t.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (t TMeUrlTypeUser) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(t.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (t TMeUrlTypeUser) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(t.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (t TMeUrlTypeUser) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(t.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (t TMeUrlTypeUser) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(t.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (t TMeUrlTypeUser) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(t.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (t TMeUrlTypeUser) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(t.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (t TMeUrlTypeUser) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(t.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (t TMeUrlTypeUser) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, t.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (t TMeUrlTypeUser) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(t.UserId, stickerFormat, sticker)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (t TonTransaction) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, t.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (t TonTransaction) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, t.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (t TonTransaction) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, t.Date)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (t TonTransactions) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, t.NextOffset, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (t TonTransactionTypeSuggestedPostPayment) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(t.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (t TonTransactionTypeSuggestedPostPayment) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(t.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (t TonTransactionTypeSuggestedPostPayment) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(t.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (t TonTransactionTypeSuggestedPostPayment) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(t.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (t TonTransactionTypeSuggestedPostPayment) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, t.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (t TonTransactionTypeSuggestedPostPayment) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(t.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (t TonTransactionTypeSuggestedPostPayment) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(t.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (t TonTransactionTypeSuggestedPostPayment) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(t.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (t TonTransactionTypeSuggestedPostPayment) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(t.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (t TonTransactionTypeSuggestedPostPayment) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(t.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (t TonTransactionTypeSuggestedPostPayment) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(t.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (t TonTransactionTypeSuggestedPostPayment) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(t.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (t TonTransactionTypeSuggestedPostPayment) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(t.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (t TonTransactionTypeSuggestedPostPayment) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(t.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (t TonTransactionTypeSuggestedPostPayment) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(t.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (t TonTransactionTypeSuggestedPostPayment) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(t.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (t TonTransactionTypeSuggestedPostPayment) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(t.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (t TonTransactionTypeSuggestedPostPayment) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(t.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (t TonTransactionTypeSuggestedPostPayment) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(t.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (t TonTransactionTypeSuggestedPostPayment) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(t.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(t.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (t TonTransactionTypeSuggestedPostPayment) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(t.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (t TonTransactionTypeSuggestedPostPayment) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(t.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (t TonTransactionTypeSuggestedPostPayment) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(t.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (t TonTransactionTypeSuggestedPostPayment) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(t.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (t TonTransactionTypeSuggestedPostPayment) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(t.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (t TonTransactionTypeSuggestedPostPayment) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(t.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (t TonTransactionTypeSuggestedPostPayment) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(t.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (t TonTransactionTypeSuggestedPostPayment) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(t.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (t TonTransactionTypeSuggestedPostPayment) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(t.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (t TonTransactionTypeSuggestedPostPayment) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(t.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (t TonTransactionTypeSuggestedPostPayment) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(t.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (t TonTransactionTypeSuggestedPostPayment) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(t.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (t TonTransactionTypeSuggestedPostPayment) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(t.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (t TonTransactionTypeSuggestedPostPayment) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(t.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (t TonTransactionTypeSuggestedPostPayment) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(t.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (t TonTransactionTypeSuggestedPostPayment) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(t.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(t.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (t TonTransactionTypeSuggestedPostPayment) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(t.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, t.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, t.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, t.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, t.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, t.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (t TonTransactionTypeSuggestedPostPayment) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, t.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (t TonTransactionTypeSuggestedPostPayment) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(t.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (t TonTransactionTypeSuggestedPostPayment) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(t.ChatId, inviteLink, name)
}
