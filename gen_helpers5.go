package gotdbot

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputGroupCallMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(i.ChatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputGroupCallMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(i.ChatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputGroupCallMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(i.ChatId, i.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (i InputGroupCallMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(i.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InputGroupCallMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputGroupCallMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(i.ChatId, i.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (i InputGroupCallMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(i.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputGroupCallMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (i InputGroupCallMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(i.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (i InputGroupCallMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(i.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InputGroupCallMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(i.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputGroupCallMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(i.ChatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputGroupCallMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(i.ChatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (i InputGroupCallMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(i.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputGroupCallMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputGroupCallMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(i.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputGroupCallMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(i.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputGroupCallMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(i.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputGroupCallMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(i.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputGroupCallMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(i.ChatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputGroupCallMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(i.ChatId, i.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (i InputGroupCallMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(i.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (i InputGroupCallMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(i.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (i InputGroupCallMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(i.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (i InputGroupCallMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(i.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (i InputGroupCallMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(i.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (i InputGroupCallMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(i.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputGroupCallMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(i.ChatId, i.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (i InputGroupCallMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(i.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (i InputGroupCallMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(i.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (i InputGroupCallMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(i.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (i InputGroupCallMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(i.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InputGroupCallMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(i.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InputGroupCallMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(i.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputGroupCallMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputGroupCallMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, i.ChatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputGroupCallMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputGroupCallMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputGroupCallMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputGroupCallMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputGroupCallMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(i.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputGroupCallMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(i.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputGroupCallMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(i.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputGroupCallMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputGroupCallMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(i.ChatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputGroupCallMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputGroupCallMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputGroupCallMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(i.ChatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputGroupCallMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(i.ChatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputGroupCallMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputGroupCallMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InputGroupCallMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(i.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InputGroupCallMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, i.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputGroupCallMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(i.ChatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputGroupCallMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(i.ChatId, i.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (i InputGroupCallMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(i.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (i InputGroupCallMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(i.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (i InputGroupCallMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(i.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (i InputGroupCallMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(i.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (i InputGroupCallMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(i.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (i InputGroupCallMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(i.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (i InputGroupCallMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(i.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (i InputGroupCallMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(i.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (i InputGroupCallMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(i.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InputGroupCallMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(i.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (i InputGroupCallMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(i.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InputGroupCallMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(i.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (i InputGroupCallMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(i.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InputGroupCallMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(i.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (i InputGroupCallMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(i.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InputGroupCallMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(i.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (i InputGroupCallMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(i.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (i InputGroupCallMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(i.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (i InputGroupCallMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(i.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (i InputGroupCallMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(i.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (i InputGroupCallMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(i.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputGroupCallMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(i.ChatId, filter, i.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (i InputGroupCallMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(i.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (i InputGroupCallMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(i.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InputGroupCallMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(i.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (i InputGroupCallMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(i.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (i InputGroupCallMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(i.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (i InputGroupCallMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(i.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (i InputGroupCallMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(i.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (i InputGroupCallMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(i.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (i InputGroupCallMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(i.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (i InputGroupCallMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(i.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InputGroupCallMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(i.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (i InputGroupCallMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(i.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (i InputGroupCallMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(i.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (i InputGroupCallMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(i.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (i InputGroupCallMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(i.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (i InputGroupCallMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(i.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (i InputGroupCallMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(i.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (i InputGroupCallMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(i.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (i InputGroupCallMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(i.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InputGroupCallMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(i.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputGroupCallMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(i.ChatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputGroupCallMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(i.ChatId, i.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InputGroupCallMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, i.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (i InputGroupCallMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(i.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputGroupCallMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(i.ChatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputGroupCallMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(i.ChatId, i.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputGroupCallMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(i.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputGroupCallMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, i.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputGroupCallMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(i.ChatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputGroupCallMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(i.ChatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputGroupCallMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(i.ChatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputGroupCallMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(i.ChatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputGroupCallMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(i.ChatId, i.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (i InputGroupCallMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(i.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputGroupCallMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(i.ChatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputGroupCallMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(i.ChatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputGroupCallMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(i.ChatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputGroupCallMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(i.ChatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputGroupCallMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(i.ChatId, i.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (i InputGroupCallMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(i.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputGroupCallMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(i.ChatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputGroupCallMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(i.ChatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputGroupCallMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(i.ChatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputGroupCallMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(i.ChatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputGroupCallMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(i.ChatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputGroupCallMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(i.ChatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputGroupCallMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(i.ChatId, i.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InputGroupCallMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(i.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (i InputGroupCallMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, i.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InputGroupCallMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(i.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputGroupCallMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(i.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InputGroupCallMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(i.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (i InputGroupCallMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(i.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (i InputGroupCallMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(i.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputGroupCallMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(i.ChatId, i.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputGroupCallMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(i.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (i InputGroupCallMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(i.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (i InputGroupCallMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(i.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (i InputGroupCallMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(i.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (i InputGroupCallMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(i.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputGroupCallMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(i.ChatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (i InputGroupCallMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(i.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (i InputGroupCallMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(i.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputGroupCallMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(i.ChatId, i.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputGroupCallMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(i.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputGroupCallMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(i.ChatId, i.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (i InputGroupCallMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(i.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InputGroupCallMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(i.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InputGroupCallMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(i.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputGroupCallMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputGroupCallMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(i.ChatId, i.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (i InputGroupCallMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(i.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (i InputGroupCallMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(i.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (i InputGroupCallMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(i.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (i InputGroupCallMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(i.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (i InputGroupCallMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(i.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputGroupCallMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, i.ChatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputGroupCallMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(i.ChatId, i.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (i InputGroupCallMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(i.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (i InputGroupCallMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(i.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputGroupCallMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(i.ChatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputGroupCallMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (i InputGroupCallMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(i.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InputGroupCallMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (i InputGroupCallMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, i.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (i InputGroupCallMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(i.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InputGroupCallMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (i InputGroupCallMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(i.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (i InputGroupCallMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(i.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (i InputGroupCallMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(i.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (i InputGroupCallMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(i.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InputGroupCallMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(i.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputGroupCallMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(i.ChatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputGroupCallMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(i.ChatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputGroupCallMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (i InputGroupCallMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(i.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InputGroupCallMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(i.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (i InputGroupCallMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, i.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InputGroupCallMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(i.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InputGroupCallMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(i.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (i InputGroupCallMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(i.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InputGroupCallMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(i.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InputGroupCallMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, i.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputGroupCallMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (i InputGroupCallMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (i InputGroupCallMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(i.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InputGroupCallMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(i.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputGroupCallMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(i.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (i InputGroupCallMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(i.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (i InputGroupCallMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(i.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputGroupCallMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(i.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputGroupCallMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, i.ChatId, i.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (i InputGroupCallMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(i.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (i InputGroupCallMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(i.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (i InputGroupCallMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(i.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (i InputGroupCallMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(i.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (i InputGroupCallMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(i.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (i InputGroupCallMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(i.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputGroupCallMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(i.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (i InputGroupCallMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(i.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (i InputGroupCallMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(i.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (i InputGroupCallMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(i.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (i InputGroupCallMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(i.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (i InputGroupCallMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(i.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (i InputGroupCallMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(i.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (i InputGroupCallMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(i.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (i InputGroupCallMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(i.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (i InputGroupCallMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(i.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (i InputGroupCallMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(i.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (i InputGroupCallMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(i.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (i InputGroupCallMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(i.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (i InputGroupCallMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(i.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (i InputGroupCallMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(i.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (i InputGroupCallMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(i.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (i InputGroupCallMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(i.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputGroupCallMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(i.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (i InputGroupCallMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(i.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (i InputGroupCallMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(i.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputGroupCallMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(i.ChatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputGroupCallMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(i.ChatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputGroupCallMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(i.ChatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputGroupCallMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(i.ChatId, i.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (i InputGroupCallMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(i.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (i InputGroupCallMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(i.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputGroupCallMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(i.ChatId, i.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputGroupCallMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(i.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (i InputGroupCallMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(i.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputGroupCallMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(i.ChatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputGroupCallMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(i.ChatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (i InputGroupCallMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(i.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputGroupCallMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputGroupCallMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(i.ChatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputGroupCallMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(i.ChatId, i.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (i InputGroupCallMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(i.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (i InputGroupCallMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(i.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (i InputGroupCallMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(i.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (i InputGroupCallMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(i.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (i InputGroupCallMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(i.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (i InputGroupCallMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, i.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (i InputGroupCallMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(i.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (i InputGroupCallMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(i.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (i InputGroupCallMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(i.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InputGroupCallMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(i.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (i InputGroupCallMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(i.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (i InputGroupCallMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(i.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InputGroupCallMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(i.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputGroupCallMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(i.ChatId, i.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (i InputGroupCallMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(i.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (i InputGroupCallMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(i.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (i InputGroupCallMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(i.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputGroupCallMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(i.ChatId, i.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (i InputGroupCallMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(i.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (i InputGroupCallMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(i.ChatId, messageIds, forceRead, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultAnimation) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultAnimation) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultAnimation) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultAnimation) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultAnimation) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultAnimation) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultAnimation) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultAnimation) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultAnimation) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultAnimation) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultAnimation) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultAnimation) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultAnimation) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultAnimation) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultAnimation) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultAnimation) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultAnimation) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultAnimation) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultAnimation) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultAnimation) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultArticle) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultArticle) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InputInlineQueryResultArticle) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InputInlineQueryResultArticle) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultArticle) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultArticle) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultArticle) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultArticle) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InputInlineQueryResultArticle) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultArticle) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultArticle) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultArticle) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultArticle) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultArticle) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultArticle) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultArticle) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InputInlineQueryResultArticle) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InputInlineQueryResultArticle) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultArticle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InputInlineQueryResultArticle) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InputInlineQueryResultArticle) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputInlineQueryResultArticle) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultArticle) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultArticle) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputInlineQueryResultArticle) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputInlineQueryResultArticle) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultArticle) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultArticle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultArticle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultArticle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultAudio) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultAudio) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultAudio) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultAudio) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultAudio) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultAudio) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultAudio) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultAudio) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultAudio) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultAudio) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultAudio) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultAudio) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultAudio) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultAudio) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultAudio) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultAudio) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultAudio) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultAudio) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultAudio) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultAudio) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultContact) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultContact) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultContact) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultContact) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultContact) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultContact) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultContact) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultContact) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultContact) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultContact) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultContact) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultDocument) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultDocument) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultDocument) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultDocument) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultDocument) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultDocument) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultDocument) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultDocument) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultDocument) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultDocument) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultDocument) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultDocument) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultDocument) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (i InputInlineQueryResultDocument) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(i.MimeType)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultDocument) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultDocument) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultDocument) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputInlineQueryResultDocument) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputInlineQueryResultDocument) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultDocument) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultDocument) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultDocument) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultDocument) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultLocation) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultLocation) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultLocation) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultLocation) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultLocation) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultLocation) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputInlineQueryResultLocation) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, messageId int64, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, messageId, i.LivePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultLocation) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultLocation) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageLiveLocation is a helper method for Client.EditInlineMessageLiveLocation
func (i InputInlineQueryResultLocation) EditInlineMessageLiveLocation(client *Client, inlineMessageId string, heading int32, proximityAlertRadius int32, opts *EditInlineMessageLiveLocationOpts) (*Ok, error) {
	return client.EditInlineMessageLiveLocation(inlineMessageId, i.LivePeriod, heading, proximityAlertRadius, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultLocation) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultLocation) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputInlineQueryResultLocation) EditMessageLiveLocation(client *Client, chatId int64, messageId int64, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, messageId, i.LivePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultLocation) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultLocation) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultLocation) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (i InputInlineQueryResultLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(i.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputInlineQueryResultLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(i.Location, zoom, width, height, scale, chatId)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultLocation) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultLocation) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultLocation) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultLocation) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultLocation) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultLocation) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultLocation) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultPhoto) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultPhoto) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultPhoto) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultPhoto) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultPhoto) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultPhoto) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultPhoto) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultPhoto) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultPhoto) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultPhoto) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultPhoto) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultPhoto) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultPhoto) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultPhoto) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultPhoto) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultPhoto) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputInlineQueryResultPhoto) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputInlineQueryResultPhoto) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultPhoto) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultPhoto) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultPhoto) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultPhoto) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultSticker) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultSticker) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultSticker) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultSticker) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultSticker) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultSticker) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultSticker) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultSticker) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultSticker) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultSticker) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultSticker) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultVenue) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultVenue) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultVenue) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultVenue) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultVenue) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultVenue) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultVenue) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultVenue) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultVenue) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultVenue) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultVenue) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultVideo) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultVideo) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultVideo) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultVideo) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultVideo) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultVideo) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultVideo) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultVideo) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultVideo) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultVideo) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultVideo) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultVideo) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultVideo) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (i InputInlineQueryResultVideo) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(i.MimeType)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultVideo) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultVideo) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultVideo) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputInlineQueryResultVideo) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputInlineQueryResultVideo) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultVideo) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultVideo) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultVideo) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultVideo) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInlineQueryResultVoiceNote) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, disableNotification bool, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, disableNotification, i.InputMessageContent, opts)
}

// AddQuickReplyShortcutMessage is a helper method for Client.AddQuickReplyShortcutMessage
func (i InputInlineQueryResultVoiceNote) AddQuickReplyShortcutMessage(client *Client, shortcutName string, replyToMessageId int64) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutMessage(shortcutName, replyToMessageId, i.InputMessageContent)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputInlineQueryResultVoiceNote) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInlineQueryResultVoiceNote) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputInlineQueryResultVoiceNote) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInlineQueryResultVoiceNote) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInlineQueryResultVoiceNote) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInlineQueryResultVoiceNote) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, messageId, i.InputMessageContent, opts)
}

// EditInlineMessageMedia is a helper method for Client.EditInlineMessageMedia
func (i InputInlineQueryResultVoiceNote) EditInlineMessageMedia(client *Client, inlineMessageId string, opts *EditInlineMessageMediaOpts) (*Ok, error) {
	return client.EditInlineMessageMedia(inlineMessageId, i.InputMessageContent, opts)
}

// EditInlineMessageText is a helper method for Client.EditInlineMessageText
func (i InputInlineQueryResultVoiceNote) EditInlineMessageText(client *Client, inlineMessageId string, opts *EditInlineMessageTextOpts) (*Ok, error) {
	return client.EditInlineMessageText(inlineMessageId, i.InputMessageContent, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInlineQueryResultVoiceNote) EditMessageMedia(client *Client, chatId int64, messageId int64, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, messageId, i.InputMessageContent, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInlineQueryResultVoiceNote) EditMessageText(client *Client, chatId int64, messageId int64, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, messageId, i.InputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInlineQueryResultVoiceNote) EditQuickReplyMessage(client *Client, shortcutId int32, messageId int64) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, messageId, i.InputMessageContent)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputInlineQueryResultVoiceNote) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInlineQueryResultVoiceNote) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, disableNotification bool, protectContent bool, effectId string, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, disableNotification, protectContent, effectId, i.InputMessageContent, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInlineQueryResultVoiceNote) SendMessage(client *Client, chatId int64, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(chatId, i.InputMessageContent, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInlineQueryResultVoiceNote) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInlineQueryResultVoiceNote) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputInlineQueryResultVoiceNote) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputInlineQueryResultVoiceNote) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InputInvoiceMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(i.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i InputInvoiceMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(i.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (i InputInvoiceMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(i.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputInvoiceMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(i.ChatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputInvoiceMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, i.ChatId, i.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputInvoiceMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(i.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputInvoiceMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(i.ChatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputInvoiceMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(i.ChatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputInvoiceMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(i.ChatId, i.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (i InputInvoiceMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(i.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InputInvoiceMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputInvoiceMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(i.ChatId, i.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (i InputInvoiceMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(i.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputInvoiceMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (i InputInvoiceMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(i.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (i InputInvoiceMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(i.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InputInvoiceMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(i.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputInvoiceMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(i.ChatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputInvoiceMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(i.ChatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (i InputInvoiceMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(i.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputInvoiceMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputInvoiceMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(i.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputInvoiceMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(i.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputInvoiceMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(i.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputInvoiceMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(i.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputInvoiceMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(i.ChatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputInvoiceMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(i.ChatId, i.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (i InputInvoiceMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(i.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (i InputInvoiceMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(i.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (i InputInvoiceMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(i.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (i InputInvoiceMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(i.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (i InputInvoiceMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(i.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (i InputInvoiceMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(i.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputInvoiceMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(i.ChatId, i.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (i InputInvoiceMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(i.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (i InputInvoiceMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(i.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (i InputInvoiceMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(i.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (i InputInvoiceMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(i.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InputInvoiceMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(i.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InputInvoiceMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(i.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputInvoiceMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputInvoiceMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, i.ChatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputInvoiceMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputInvoiceMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputInvoiceMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputInvoiceMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputInvoiceMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(i.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputInvoiceMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(i.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputInvoiceMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(i.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputInvoiceMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputInvoiceMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(i.ChatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputInvoiceMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputInvoiceMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputInvoiceMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(i.ChatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputInvoiceMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(i.ChatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputInvoiceMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputInvoiceMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InputInvoiceMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(i.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InputInvoiceMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, i.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputInvoiceMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(i.ChatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputInvoiceMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(i.ChatId, i.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (i InputInvoiceMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(i.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (i InputInvoiceMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(i.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (i InputInvoiceMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(i.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (i InputInvoiceMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(i.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (i InputInvoiceMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(i.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (i InputInvoiceMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(i.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (i InputInvoiceMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(i.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (i InputInvoiceMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(i.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (i InputInvoiceMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(i.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InputInvoiceMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(i.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (i InputInvoiceMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(i.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InputInvoiceMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(i.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (i InputInvoiceMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(i.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InputInvoiceMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(i.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (i InputInvoiceMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(i.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InputInvoiceMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(i.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (i InputInvoiceMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(i.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (i InputInvoiceMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(i.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (i InputInvoiceMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(i.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (i InputInvoiceMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(i.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (i InputInvoiceMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(i.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputInvoiceMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(i.ChatId, filter, i.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (i InputInvoiceMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(i.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (i InputInvoiceMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(i.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InputInvoiceMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(i.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (i InputInvoiceMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(i.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (i InputInvoiceMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(i.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (i InputInvoiceMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(i.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (i InputInvoiceMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(i.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (i InputInvoiceMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(i.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (i InputInvoiceMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(i.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (i InputInvoiceMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(i.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InputInvoiceMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(i.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (i InputInvoiceMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(i.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (i InputInvoiceMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(i.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (i InputInvoiceMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(i.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (i InputInvoiceMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(i.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (i InputInvoiceMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(i.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (i InputInvoiceMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(i.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (i InputInvoiceMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(i.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (i InputInvoiceMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(i.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InputInvoiceMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(i.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputInvoiceMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(i.ChatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputInvoiceMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(i.ChatId, i.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InputInvoiceMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, i.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (i InputInvoiceMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(i.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputInvoiceMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(i.ChatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputInvoiceMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(i.ChatId, i.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputInvoiceMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(i.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputInvoiceMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, i.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputInvoiceMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(i.ChatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputInvoiceMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(i.ChatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputInvoiceMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(i.ChatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputInvoiceMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(i.ChatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputInvoiceMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(i.ChatId, i.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (i InputInvoiceMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(i.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputInvoiceMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(i.ChatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputInvoiceMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(i.ChatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputInvoiceMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(i.ChatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputInvoiceMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(i.ChatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputInvoiceMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(i.ChatId, i.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (i InputInvoiceMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(i.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputInvoiceMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(i.ChatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputInvoiceMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(i.ChatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputInvoiceMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(i.ChatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputInvoiceMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(i.ChatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputInvoiceMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(i.ChatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputInvoiceMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(i.ChatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputInvoiceMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(i.ChatId, i.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InputInvoiceMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(i.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (i InputInvoiceMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, i.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InputInvoiceMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(i.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputInvoiceMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(i.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InputInvoiceMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(i.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (i InputInvoiceMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(i.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (i InputInvoiceMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(i.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputInvoiceMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(i.ChatId, i.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputInvoiceMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(i.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (i InputInvoiceMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(i.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (i InputInvoiceMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(i.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (i InputInvoiceMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(i.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (i InputInvoiceMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(i.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputInvoiceMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(i.ChatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (i InputInvoiceMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(i.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (i InputInvoiceMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(i.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputInvoiceMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(i.ChatId, i.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputInvoiceMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(i.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputInvoiceMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(i.ChatId, i.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (i InputInvoiceMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(i.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InputInvoiceMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(i.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InputInvoiceMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(i.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputInvoiceMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputInvoiceMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(i.ChatId, i.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (i InputInvoiceMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(i.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (i InputInvoiceMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(i.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (i InputInvoiceMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(i.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (i InputInvoiceMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(i.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (i InputInvoiceMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(i.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputInvoiceMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, i.ChatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputInvoiceMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(i.ChatId, i.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (i InputInvoiceMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(i.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (i InputInvoiceMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(i.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputInvoiceMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(i.ChatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputInvoiceMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (i InputInvoiceMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(i.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InputInvoiceMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (i InputInvoiceMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, i.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (i InputInvoiceMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(i.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InputInvoiceMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (i InputInvoiceMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(i.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (i InputInvoiceMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(i.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (i InputInvoiceMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(i.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (i InputInvoiceMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(i.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InputInvoiceMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(i.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputInvoiceMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(i.ChatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputInvoiceMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(i.ChatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputInvoiceMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (i InputInvoiceMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(i.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InputInvoiceMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(i.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (i InputInvoiceMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, i.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InputInvoiceMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(i.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InputInvoiceMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(i.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (i InputInvoiceMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(i.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InputInvoiceMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(i.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InputInvoiceMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, i.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputInvoiceMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (i InputInvoiceMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (i InputInvoiceMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(i.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InputInvoiceMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(i.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputInvoiceMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(i.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (i InputInvoiceMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(i.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (i InputInvoiceMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(i.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputInvoiceMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(i.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputInvoiceMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, i.ChatId, i.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (i InputInvoiceMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(i.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (i InputInvoiceMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(i.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (i InputInvoiceMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(i.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (i InputInvoiceMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(i.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (i InputInvoiceMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(i.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (i InputInvoiceMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(i.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputInvoiceMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(i.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (i InputInvoiceMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(i.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (i InputInvoiceMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(i.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (i InputInvoiceMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(i.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (i InputInvoiceMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(i.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (i InputInvoiceMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(i.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (i InputInvoiceMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(i.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (i InputInvoiceMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(i.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (i InputInvoiceMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(i.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (i InputInvoiceMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(i.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (i InputInvoiceMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(i.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (i InputInvoiceMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(i.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (i InputInvoiceMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(i.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (i InputInvoiceMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(i.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (i InputInvoiceMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(i.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (i InputInvoiceMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(i.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (i InputInvoiceMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(i.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputInvoiceMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(i.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (i InputInvoiceMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(i.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (i InputInvoiceMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(i.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputInvoiceMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(i.ChatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputInvoiceMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(i.ChatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputInvoiceMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(i.ChatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputInvoiceMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(i.ChatId, i.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (i InputInvoiceMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(i.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (i InputInvoiceMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(i.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputInvoiceMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(i.ChatId, i.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputInvoiceMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(i.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (i InputInvoiceMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(i.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputInvoiceMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(i.ChatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputInvoiceMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(i.ChatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (i InputInvoiceMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(i.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputInvoiceMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputInvoiceMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(i.ChatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputInvoiceMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(i.ChatId, i.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (i InputInvoiceMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(i.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (i InputInvoiceMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(i.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (i InputInvoiceMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(i.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (i InputInvoiceMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(i.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (i InputInvoiceMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(i.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (i InputInvoiceMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, i.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (i InputInvoiceMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(i.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (i InputInvoiceMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(i.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (i InputInvoiceMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(i.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InputInvoiceMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(i.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (i InputInvoiceMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(i.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (i InputInvoiceMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(i.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InputInvoiceMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(i.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputInvoiceMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(i.ChatId, i.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (i InputInvoiceMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(i.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (i InputInvoiceMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(i.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (i InputInvoiceMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(i.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputInvoiceMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(i.ChatId, i.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (i InputInvoiceMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(i.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (i InputInvoiceMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(i.ChatId, messageIds, forceRead, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (i InputInvoiceName) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, i.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (i InputInvoiceName) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(i.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (i InputInvoiceName) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(i.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (i InputInvoiceName) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, i.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputInvoiceName) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputInvoiceName) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, i.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputInvoiceName) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, i.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (i InputInvoiceName) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, i.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputInvoiceName) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, i.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InputInvoiceName) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, i.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (i InputInvoiceName) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(i.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InputInvoiceName) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, i.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputInvoiceName) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputInvoiceName) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, i.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputInvoiceName) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, i.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (i InputInvoiceName) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(i.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (i InputInvoiceName) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(i.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (i InputInvoiceName) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(i.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (i InputInvoiceName) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(i.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (i InputInvoiceName) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, i.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (i InputInvoiceName) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(i.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (i InputInvoiceName) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(i.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (i InputInvoiceName) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, i.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (i InputInvoiceName) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(i.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (i InputInvoiceName) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, i.Name)
}

// SetOption is a helper method for Client.SetOption
func (i InputInvoiceName) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(i.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (i InputInvoiceName) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, i.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (i InputInvoiceName) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, i.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputInvoiceName) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(i.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputInvoiceName) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, i.Name)
}

// AddSavedAnimation is a helper method for Client.AddSavedAnimation
func (i InputMessageAnimation) AddSavedAnimation(client *Client) (*Ok, error) {
	return client.AddSavedAnimation(i.Animation)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputMessageAnimation) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessageAnimation) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageAnimation) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (i InputMessageAnimation) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, i.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessageAnimation) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessageAnimation) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// RemoveSavedAnimation is a helper method for Client.RemoveSavedAnimation
func (i InputMessageAnimation) RemoveSavedAnimation(client *Client) (*Ok, error) {
	return client.RemoveSavedAnimation(i.Animation)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputMessageAnimation) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputMessageAudio) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputMessageAudio) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputMessageAudio) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputMessageAudio) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputMessageAudio) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageAudio) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputMessageAudio) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputMessageAudio) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputMessageAudio) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputMessageAudio) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputMessageAudio) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputMessageAudio) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputMessageChecklist) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, messageId, i.Checklist, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputMessageChecklist) EditMessageChecklist(client *Client, chatId int64, messageId int64, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, messageId, i.Checklist, opts)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (i InputMessageDice) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(i.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (i InputMessageDice) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(i.Emoji)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageDocument) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputMessageForwarded) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputMessageForwarded) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, i.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputMessageForwarded) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputMessageForwarded) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputMessageForwarded) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, i.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputMessageForwarded) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, i.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputMessageForwarded) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputMessageForwarded) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputMessageForwarded) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputMessageForwarded) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, i.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputMessageForwarded) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputMessageForwarded) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, i.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputMessageForwarded) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, i.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessageForwarded) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputMessageForwarded) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputMessageForwarded) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputMessageForwarded) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputMessageForwarded) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputMessageForwarded) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessageForwarded) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputMessageForwarded) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputMessageForwarded) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputMessageForwarded) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputMessageForwarded) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputMessageForwarded) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputMessageForwarded) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputMessageForwarded) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InputMessageForwarded) ForwardMessages(client *Client, chatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(chatId, i.FromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputMessageForwarded) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputMessageForwarded) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, i.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputMessageForwarded) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, i.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputMessageForwarded) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputMessageForwarded) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, i.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputMessageForwarded) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputMessageForwarded) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, i.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputMessageForwarded) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputMessageForwarded) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputMessageForwarded) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputMessageForwarded) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputMessageForwarded) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, i.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputMessageForwarded) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputMessageForwarded) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputMessageForwarded) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputMessageForwarded) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputMessageForwarded) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, i.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputMessageForwarded) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputMessageForwarded) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputMessageForwarded) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputMessageForwarded) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputMessageForwarded) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputMessageForwarded) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputMessageForwarded) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, i.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputMessageForwarded) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, i.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputMessageForwarded) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputMessageForwarded) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, i.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputMessageForwarded) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, i.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputMessageForwarded) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputMessageForwarded) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, i.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputMessageForwarded) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputMessageForwarded) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, i.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputMessageForwarded) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputMessageForwarded) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, i.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputMessageForwarded) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputMessageForwarded) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputMessageForwarded) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputMessageForwarded) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, i.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputMessageForwarded) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputMessageForwarded) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputMessageForwarded) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputMessageForwarded) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, i.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputMessageForwarded) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, i.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputMessageForwarded) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputMessageForwarded) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputMessageForwarded) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputMessageForwarded) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputMessageForwarded) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, i.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputMessageForwarded) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, i.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputMessageForwarded) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, i.MessageId)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (i InputMessageGame) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(i.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (i InputMessageGame) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(i.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (i InputMessageGame) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(i.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InputMessageGame) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(i.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (i InputMessageGame) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, i.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (i InputMessageGame) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(i.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (i InputMessageGame) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(i.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (i InputMessageGame) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(i.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (i InputMessageGame) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(i.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (i InputMessageGame) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(i.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (i InputMessageGame) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(i.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (i InputMessageGame) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(i.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (i InputMessageGame) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(i.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (i InputMessageGame) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(i.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (i InputMessageGame) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(i.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (i InputMessageGame) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(i.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (i InputMessageGame) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, i.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InputMessageGame) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(i.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputMessageGame) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, i.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (i InputMessageGame) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(i.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (i InputMessageGame) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(i.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputMessageGame) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, i.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (i InputMessageGame) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(i.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InputMessageGame) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(i.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (i InputMessageGame) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(i.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputMessageGame) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, i.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (i InputMessageGame) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(i.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (i InputMessageGame) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(i.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (i InputMessageGame) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(i.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (i InputMessageGame) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(i.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InputMessageGame) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(i.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (i InputMessageGame) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(i.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (i InputMessageGame) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(i.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputMessageGame) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(i.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (i InputMessageGame) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(i.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (i InputMessageGame) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(i.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (i InputMessageGame) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(i.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (i InputMessageGame) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(i.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (i InputMessageGame) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(i.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (i InputMessageGame) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(i.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (i InputMessageGame) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(i.BotUserId, username, isActive)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputMessageInvoice) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputMessageInvoice) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputMessageInvoice) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputMessageInvoice) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputMessageInvoice) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, i.StartParameter, parameters)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputMessageInvoice) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputMessageInvoice) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, webAppShortName string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, webAppShortName, i.StartParameter, allowWriteAccess, parameters)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InputMessageInvoice) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputMessageInvoice) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputMessageInvoice) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputMessageInvoice) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputMessageInvoice) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputMessageInvoice) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputMessageLocation) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, messageId, i.LivePeriod, i.Heading, i.ProximityAlertRadius, opts)
}

// EditInlineMessageLiveLocation is a helper method for Client.EditInlineMessageLiveLocation
func (i InputMessageLocation) EditInlineMessageLiveLocation(client *Client, inlineMessageId string, opts *EditInlineMessageLiveLocationOpts) (*Ok, error) {
	return client.EditInlineMessageLiveLocation(inlineMessageId, i.LivePeriod, i.Heading, i.ProximityAlertRadius, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputMessageLocation) EditMessageLiveLocation(client *Client, chatId int64, messageId int64, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, messageId, i.LivePeriod, i.Heading, i.ProximityAlertRadius, opts)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (i InputMessageLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(i.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessageLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(i.Location, zoom, width, height, scale, chatId)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (i InputMessagePaidMedia) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, i.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputMessagePaidMedia) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, i.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (i InputMessagePaidMedia) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, i.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (i InputMessagePaidMedia) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, i.StarCount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessagePaidMedia) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessagePaidMedia) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (i InputMessagePaidMedia) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, i.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessagePaidMedia) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// GetPushReceiverId is a helper method for Client.GetPushReceiverId
func (i InputMessagePaidMedia) GetPushReceiverId(client *Client) (*PushReceiverId, error) {
	return client.GetPushReceiverId(i.Payload)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (i InputMessagePaidMedia) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, i.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InputMessagePaidMedia) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, i.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (i InputMessagePaidMedia) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, i.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (i InputMessagePaidMedia) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, i.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InputMessagePaidMedia) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, i.StarCount, userId, text, isPrivate)
}

// ProcessPushNotification is a helper method for Client.ProcessPushNotification
func (i InputMessagePaidMedia) ProcessPushNotification(client *Client) (*Ok, error) {
	return client.ProcessPushNotification(i.Payload)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (i InputMessagePaidMedia) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, i.StarCount)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (i InputMessagePaidMedia) StartGroupCallScreenSharing(client *Client, groupCallId int32, audioSourceId int32) (*Text, error) {
	return client.StartGroupCallScreenSharing(groupCallId, audioSourceId, i.Payload)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (i InputMessagePaidMedia) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, i.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (i InputMessagePaidMedia) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, i.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (i InputMessagePaidMedia) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, i.StarCount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessagePhoto) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessagePhoto) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (i InputMessagePhoto) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, i.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessagePhoto) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessagePhoto) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InputMessagePoll) ToggleForumTopicIsClosed(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(chatId, forumTopicId, i.IsClosed)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InputMessageReplyToExternalMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(i.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i InputMessageReplyToExternalMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(i.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (i InputMessageReplyToExternalMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(i.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputMessageReplyToExternalMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(i.ChatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputMessageReplyToExternalMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, i.ChatId, i.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputMessageReplyToExternalMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(i.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputMessageReplyToExternalMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(i.ChatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputMessageReplyToExternalMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(i.ChatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputMessageReplyToExternalMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(i.ChatId, i.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (i InputMessageReplyToExternalMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(i.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InputMessageReplyToExternalMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputMessageReplyToExternalMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(i.ChatId, i.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (i InputMessageReplyToExternalMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(i.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputMessageReplyToExternalMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (i InputMessageReplyToExternalMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(i.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (i InputMessageReplyToExternalMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(i.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InputMessageReplyToExternalMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(i.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputMessageReplyToExternalMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(i.ChatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputMessageReplyToExternalMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(i.ChatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (i InputMessageReplyToExternalMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(i.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputMessageReplyToExternalMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputMessageReplyToExternalMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(i.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputMessageReplyToExternalMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(i.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputMessageReplyToExternalMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(i.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputMessageReplyToExternalMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(i.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputMessageReplyToExternalMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(i.ChatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputMessageReplyToExternalMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(i.ChatId, i.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (i InputMessageReplyToExternalMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(i.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (i InputMessageReplyToExternalMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(i.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (i InputMessageReplyToExternalMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(i.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (i InputMessageReplyToExternalMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(i.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (i InputMessageReplyToExternalMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(i.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (i InputMessageReplyToExternalMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(i.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputMessageReplyToExternalMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(i.ChatId, i.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (i InputMessageReplyToExternalMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(i.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (i InputMessageReplyToExternalMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(i.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (i InputMessageReplyToExternalMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(i.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (i InputMessageReplyToExternalMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(i.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InputMessageReplyToExternalMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(i.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InputMessageReplyToExternalMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(i.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessageReplyToExternalMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputMessageReplyToExternalMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, i.ChatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputMessageReplyToExternalMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputMessageReplyToExternalMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputMessageReplyToExternalMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputMessageReplyToExternalMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputMessageReplyToExternalMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(i.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputMessageReplyToExternalMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(i.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputMessageReplyToExternalMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(i.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessageReplyToExternalMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputMessageReplyToExternalMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(i.ChatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputMessageReplyToExternalMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputMessageReplyToExternalMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputMessageReplyToExternalMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(i.ChatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputMessageReplyToExternalMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(i.ChatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputMessageReplyToExternalMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputMessageReplyToExternalMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InputMessageReplyToExternalMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(i.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InputMessageReplyToExternalMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, i.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputMessageReplyToExternalMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(i.ChatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputMessageReplyToExternalMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(i.ChatId, i.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (i InputMessageReplyToExternalMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(i.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (i InputMessageReplyToExternalMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(i.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (i InputMessageReplyToExternalMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(i.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (i InputMessageReplyToExternalMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(i.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (i InputMessageReplyToExternalMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(i.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (i InputMessageReplyToExternalMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(i.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (i InputMessageReplyToExternalMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(i.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (i InputMessageReplyToExternalMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(i.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (i InputMessageReplyToExternalMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(i.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InputMessageReplyToExternalMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(i.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (i InputMessageReplyToExternalMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(i.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InputMessageReplyToExternalMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(i.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (i InputMessageReplyToExternalMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(i.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InputMessageReplyToExternalMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(i.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (i InputMessageReplyToExternalMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(i.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InputMessageReplyToExternalMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(i.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (i InputMessageReplyToExternalMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(i.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (i InputMessageReplyToExternalMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(i.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (i InputMessageReplyToExternalMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(i.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (i InputMessageReplyToExternalMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(i.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (i InputMessageReplyToExternalMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(i.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputMessageReplyToExternalMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(i.ChatId, filter, i.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (i InputMessageReplyToExternalMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(i.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (i InputMessageReplyToExternalMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(i.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InputMessageReplyToExternalMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(i.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (i InputMessageReplyToExternalMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(i.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (i InputMessageReplyToExternalMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(i.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (i InputMessageReplyToExternalMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(i.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (i InputMessageReplyToExternalMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(i.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (i InputMessageReplyToExternalMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(i.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (i InputMessageReplyToExternalMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(i.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (i InputMessageReplyToExternalMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(i.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InputMessageReplyToExternalMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(i.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (i InputMessageReplyToExternalMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(i.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (i InputMessageReplyToExternalMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(i.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (i InputMessageReplyToExternalMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(i.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (i InputMessageReplyToExternalMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(i.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (i InputMessageReplyToExternalMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(i.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (i InputMessageReplyToExternalMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(i.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (i InputMessageReplyToExternalMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(i.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (i InputMessageReplyToExternalMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(i.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InputMessageReplyToExternalMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(i.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputMessageReplyToExternalMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(i.ChatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputMessageReplyToExternalMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(i.ChatId, i.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InputMessageReplyToExternalMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, i.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (i InputMessageReplyToExternalMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(i.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputMessageReplyToExternalMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(i.ChatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputMessageReplyToExternalMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(i.ChatId, i.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputMessageReplyToExternalMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(i.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessageReplyToExternalMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, i.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputMessageReplyToExternalMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(i.ChatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputMessageReplyToExternalMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(i.ChatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputMessageReplyToExternalMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(i.ChatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputMessageReplyToExternalMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(i.ChatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputMessageReplyToExternalMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(i.ChatId, i.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (i InputMessageReplyToExternalMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(i.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputMessageReplyToExternalMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(i.ChatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputMessageReplyToExternalMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(i.ChatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputMessageReplyToExternalMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(i.ChatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputMessageReplyToExternalMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(i.ChatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputMessageReplyToExternalMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(i.ChatId, i.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (i InputMessageReplyToExternalMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(i.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputMessageReplyToExternalMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(i.ChatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputMessageReplyToExternalMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(i.ChatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputMessageReplyToExternalMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(i.ChatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputMessageReplyToExternalMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(i.ChatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputMessageReplyToExternalMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(i.ChatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputMessageReplyToExternalMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(i.ChatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputMessageReplyToExternalMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(i.ChatId, i.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InputMessageReplyToExternalMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(i.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (i InputMessageReplyToExternalMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, i.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InputMessageReplyToExternalMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(i.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputMessageReplyToExternalMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(i.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InputMessageReplyToExternalMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(i.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (i InputMessageReplyToExternalMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(i.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (i InputMessageReplyToExternalMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(i.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputMessageReplyToExternalMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(i.ChatId, i.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputMessageReplyToExternalMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(i.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (i InputMessageReplyToExternalMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(i.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (i InputMessageReplyToExternalMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(i.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (i InputMessageReplyToExternalMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(i.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (i InputMessageReplyToExternalMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(i.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputMessageReplyToExternalMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(i.ChatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (i InputMessageReplyToExternalMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(i.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (i InputMessageReplyToExternalMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(i.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputMessageReplyToExternalMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(i.ChatId, i.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputMessageReplyToExternalMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(i.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputMessageReplyToExternalMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(i.ChatId, i.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (i InputMessageReplyToExternalMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(i.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InputMessageReplyToExternalMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(i.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InputMessageReplyToExternalMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(i.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputMessageReplyToExternalMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputMessageReplyToExternalMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(i.ChatId, i.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (i InputMessageReplyToExternalMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(i.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (i InputMessageReplyToExternalMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(i.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (i InputMessageReplyToExternalMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(i.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (i InputMessageReplyToExternalMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(i.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (i InputMessageReplyToExternalMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(i.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputMessageReplyToExternalMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, i.ChatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputMessageReplyToExternalMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(i.ChatId, i.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (i InputMessageReplyToExternalMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(i.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (i InputMessageReplyToExternalMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(i.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputMessageReplyToExternalMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(i.ChatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputMessageReplyToExternalMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (i InputMessageReplyToExternalMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(i.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InputMessageReplyToExternalMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (i InputMessageReplyToExternalMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, i.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (i InputMessageReplyToExternalMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(i.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InputMessageReplyToExternalMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (i InputMessageReplyToExternalMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(i.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (i InputMessageReplyToExternalMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(i.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (i InputMessageReplyToExternalMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(i.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (i InputMessageReplyToExternalMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(i.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InputMessageReplyToExternalMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(i.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputMessageReplyToExternalMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(i.ChatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputMessageReplyToExternalMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(i.ChatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputMessageReplyToExternalMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (i InputMessageReplyToExternalMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(i.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InputMessageReplyToExternalMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(i.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (i InputMessageReplyToExternalMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, i.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InputMessageReplyToExternalMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(i.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InputMessageReplyToExternalMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(i.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (i InputMessageReplyToExternalMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(i.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InputMessageReplyToExternalMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(i.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InputMessageReplyToExternalMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, i.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputMessageReplyToExternalMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (i InputMessageReplyToExternalMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (i InputMessageReplyToExternalMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(i.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InputMessageReplyToExternalMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(i.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputMessageReplyToExternalMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(i.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (i InputMessageReplyToExternalMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(i.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (i InputMessageReplyToExternalMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(i.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputMessageReplyToExternalMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(i.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputMessageReplyToExternalMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, i.ChatId, i.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (i InputMessageReplyToExternalMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(i.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (i InputMessageReplyToExternalMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(i.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (i InputMessageReplyToExternalMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(i.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (i InputMessageReplyToExternalMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(i.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (i InputMessageReplyToExternalMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(i.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (i InputMessageReplyToExternalMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(i.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputMessageReplyToExternalMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(i.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (i InputMessageReplyToExternalMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(i.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (i InputMessageReplyToExternalMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(i.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (i InputMessageReplyToExternalMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(i.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (i InputMessageReplyToExternalMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(i.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (i InputMessageReplyToExternalMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(i.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (i InputMessageReplyToExternalMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(i.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (i InputMessageReplyToExternalMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(i.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (i InputMessageReplyToExternalMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(i.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (i InputMessageReplyToExternalMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(i.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (i InputMessageReplyToExternalMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(i.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (i InputMessageReplyToExternalMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(i.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (i InputMessageReplyToExternalMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(i.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (i InputMessageReplyToExternalMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(i.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (i InputMessageReplyToExternalMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(i.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (i InputMessageReplyToExternalMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(i.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (i InputMessageReplyToExternalMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(i.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputMessageReplyToExternalMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(i.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (i InputMessageReplyToExternalMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(i.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (i InputMessageReplyToExternalMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(i.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputMessageReplyToExternalMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(i.ChatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputMessageReplyToExternalMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(i.ChatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputMessageReplyToExternalMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(i.ChatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputMessageReplyToExternalMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(i.ChatId, i.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (i InputMessageReplyToExternalMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(i.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (i InputMessageReplyToExternalMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(i.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputMessageReplyToExternalMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(i.ChatId, i.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputMessageReplyToExternalMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(i.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (i InputMessageReplyToExternalMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(i.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputMessageReplyToExternalMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(i.ChatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputMessageReplyToExternalMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(i.ChatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (i InputMessageReplyToExternalMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(i.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputMessageReplyToExternalMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputMessageReplyToExternalMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(i.ChatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputMessageReplyToExternalMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(i.ChatId, i.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (i InputMessageReplyToExternalMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(i.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (i InputMessageReplyToExternalMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(i.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (i InputMessageReplyToExternalMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(i.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (i InputMessageReplyToExternalMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(i.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (i InputMessageReplyToExternalMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(i.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (i InputMessageReplyToExternalMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, i.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (i InputMessageReplyToExternalMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(i.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (i InputMessageReplyToExternalMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(i.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (i InputMessageReplyToExternalMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(i.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InputMessageReplyToExternalMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(i.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (i InputMessageReplyToExternalMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(i.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (i InputMessageReplyToExternalMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(i.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InputMessageReplyToExternalMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(i.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputMessageReplyToExternalMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(i.ChatId, i.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (i InputMessageReplyToExternalMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(i.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (i InputMessageReplyToExternalMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(i.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (i InputMessageReplyToExternalMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(i.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputMessageReplyToExternalMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(i.ChatId, i.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (i InputMessageReplyToExternalMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(i.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (i InputMessageReplyToExternalMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(i.ChatId, messageIds, forceRead, opts)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputMessageReplyToMessage) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputMessageReplyToMessage) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, i.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputMessageReplyToMessage) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputMessageReplyToMessage) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputMessageReplyToMessage) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, i.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputMessageReplyToMessage) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, i.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputMessageReplyToMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputMessageReplyToMessage) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputMessageReplyToMessage) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputMessageReplyToMessage) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, i.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputMessageReplyToMessage) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputMessageReplyToMessage) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, i.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputMessageReplyToMessage) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, i.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessageReplyToMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputMessageReplyToMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputMessageReplyToMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputMessageReplyToMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputMessageReplyToMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputMessageReplyToMessage) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessageReplyToMessage) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputMessageReplyToMessage) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputMessageReplyToMessage) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputMessageReplyToMessage) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputMessageReplyToMessage) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputMessageReplyToMessage) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputMessageReplyToMessage) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputMessageReplyToMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputMessageReplyToMessage) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputMessageReplyToMessage) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, i.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputMessageReplyToMessage) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, i.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputMessageReplyToMessage) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputMessageReplyToMessage) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, i.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputMessageReplyToMessage) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputMessageReplyToMessage) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, i.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputMessageReplyToMessage) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputMessageReplyToMessage) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputMessageReplyToMessage) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputMessageReplyToMessage) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputMessageReplyToMessage) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, i.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputMessageReplyToMessage) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputMessageReplyToMessage) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputMessageReplyToMessage) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputMessageReplyToMessage) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputMessageReplyToMessage) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, i.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputMessageReplyToMessage) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputMessageReplyToMessage) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputMessageReplyToMessage) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputMessageReplyToMessage) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputMessageReplyToMessage) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputMessageReplyToMessage) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputMessageReplyToMessage) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, i.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputMessageReplyToMessage) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, i.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputMessageReplyToMessage) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputMessageReplyToMessage) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, i.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputMessageReplyToMessage) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, i.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputMessageReplyToMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputMessageReplyToMessage) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, i.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputMessageReplyToMessage) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputMessageReplyToMessage) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, i.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputMessageReplyToMessage) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputMessageReplyToMessage) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, i.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputMessageReplyToMessage) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputMessageReplyToMessage) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputMessageReplyToMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputMessageReplyToMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, i.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputMessageReplyToMessage) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputMessageReplyToMessage) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputMessageReplyToMessage) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputMessageReplyToMessage) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, i.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputMessageReplyToMessage) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, i.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputMessageReplyToMessage) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputMessageReplyToMessage) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputMessageReplyToMessage) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputMessageReplyToMessage) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputMessageReplyToMessage) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, i.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputMessageReplyToMessage) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, i.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputMessageReplyToMessage) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, i.MessageId)
}

// CloseStory is a helper method for Client.CloseStory
func (i InputMessageReplyToStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(i.StoryPosterChatId, i.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InputMessageReplyToStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(i.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (i InputMessageReplyToStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, i.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (i InputMessageReplyToStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(i.StoryPosterChatId, i.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageReplyToStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(i.StoryPosterChatId, i.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (i InputMessageReplyToStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(i.StoryPosterChatId, i.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (i InputMessageReplyToStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(i.StoryPosterChatId, i.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (i InputMessageReplyToStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(i.StoryPosterChatId, i.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (i InputMessageReplyToStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(i.StoryPosterChatId, i.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (i InputMessageReplyToStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(i.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (i InputMessageReplyToStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(i.StoryPosterChatId, i.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputMessageReplyToStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, i.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (i InputMessageReplyToStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(i.StoryPosterChatId, i.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (i InputMessageReplyToStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(i.StoryPosterChatId, i.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (i InputMessageReplyToStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(i.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (i InputMessageReplyToStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(i.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (i InputMessageReplyToStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(i.StoryPosterChatId, i.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (i InputMessageReplyToStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(i.StoryPosterChatId, i.StoryId, isPostedToChatPage)
}

// AddFavoriteSticker is a helper method for Client.AddFavoriteSticker
func (i InputMessageSticker) AddFavoriteSticker(client *Client) (*Ok, error) {
	return client.AddFavoriteSticker(i.Sticker)
}

// AddRecentSticker is a helper method for Client.AddRecentSticker
func (i InputMessageSticker) AddRecentSticker(client *Client, isAttached bool) (*Stickers, error) {
	return client.AddRecentSticker(isAttached, i.Sticker)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (i InputMessageSticker) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(i.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (i InputMessageSticker) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(i.Emoji)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessageSticker) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// GetStickerEmojis is a helper method for Client.GetStickerEmojis
func (i InputMessageSticker) GetStickerEmojis(client *Client) (*Emojis, error) {
	return client.GetStickerEmojis(i.Sticker)
}

// RemoveFavoriteSticker is a helper method for Client.RemoveFavoriteSticker
func (i InputMessageSticker) RemoveFavoriteSticker(client *Client) (*Ok, error) {
	return client.RemoveFavoriteSticker(i.Sticker)
}

// RemoveRecentSticker is a helper method for Client.RemoveRecentSticker
func (i InputMessageSticker) RemoveRecentSticker(client *Client, isAttached bool) (*Ok, error) {
	return client.RemoveRecentSticker(isAttached, i.Sticker)
}

// RemoveStickerFromSet is a helper method for Client.RemoveStickerFromSet
func (i InputMessageSticker) RemoveStickerFromSet(client *Client) (*Ok, error) {
	return client.RemoveStickerFromSet(i.Sticker)
}

// SetStickerEmojis is a helper method for Client.SetStickerEmojis
func (i InputMessageSticker) SetStickerEmojis(client *Client, emojis string) (*Ok, error) {
	return client.SetStickerEmojis(i.Sticker, emojis)
}

// SetStickerKeywords is a helper method for Client.SetStickerKeywords
func (i InputMessageSticker) SetStickerKeywords(client *Client, keywords []string) (*Ok, error) {
	return client.SetStickerKeywords(i.Sticker, keywords)
}

// SetStickerMaskPosition is a helper method for Client.SetStickerMaskPosition
func (i InputMessageSticker) SetStickerMaskPosition(client *Client, opts *SetStickerMaskPositionOpts) (*Ok, error) {
	return client.SetStickerMaskPosition(i.Sticker, opts)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (i InputMessageSticker) SetStickerPositionInSet(client *Client, position int32) (*Ok, error) {
	return client.SetStickerPositionInSet(i.Sticker, position)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (i InputMessageSticker) UploadStickerFile(client *Client, userId int64, stickerFormat *StickerFormat) (*File, error) {
	return client.UploadStickerFile(userId, stickerFormat, i.Sticker)
}

// CloseStory is a helper method for Client.CloseStory
func (i InputMessageStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(i.StoryPosterChatId, i.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InputMessageStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(i.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (i InputMessageStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, i.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (i InputMessageStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(i.StoryPosterChatId, i.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(i.StoryPosterChatId, i.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (i InputMessageStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(i.StoryPosterChatId, i.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (i InputMessageStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(i.StoryPosterChatId, i.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (i InputMessageStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(i.StoryPosterChatId, i.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (i InputMessageStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(i.StoryPosterChatId, i.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (i InputMessageStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(i.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (i InputMessageStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(i.StoryPosterChatId, i.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputMessageStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, i.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (i InputMessageStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(i.StoryPosterChatId, i.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (i InputMessageStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(i.StoryPosterChatId, i.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (i InputMessageStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(i.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (i InputMessageStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(i.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (i InputMessageStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(i.StoryPosterChatId, i.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (i InputMessageStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(i.StoryPosterChatId, i.StoryId, isPostedToChatPage)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (i InputMessageText) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(i.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (i InputMessageText) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(i.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InputMessageText) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, i.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (i InputMessageText) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(i.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InputMessageText) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, i.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (i InputMessageText) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(i.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (i InputMessageText) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, i.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (i InputMessageText) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, i.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputMessageText) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, i.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (i InputMessageText) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(i.Text, toLanguageCode)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputMessageVideo) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputMessageVideo) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageVideo) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (i InputMessageVideo) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, i.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputMessageVideo) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, i.ShowCaptionAboveMedia, opts)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputMessageVideo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputMessageVideo) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputMessageVideoNote) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputMessageVideoNote) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputMessageVoiceNote) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InputMessageVoiceNote) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, i.Caption, privacySettings)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputMessageVoiceNote) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputPaidMedia) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InputPaidMediaTypeVideo) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, i.Duration, isVideo, connectionId)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputPaidMediaTypeVideo) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, i.Duration, paidMessageStarCount)
}

// SendEmailAddressVerificationCode is a helper method for Client.SendEmailAddressVerificationCode
func (i InputPassportElementEmailAddress) SendEmailAddressVerificationCode(client *Client) (*EmailAddressAuthenticationCodeInfo, error) {
	return client.SendEmailAddressVerificationCode(i.EmailAddress)
}

// SetAuthenticationEmailAddress is a helper method for Client.SetAuthenticationEmailAddress
func (i InputPassportElementEmailAddress) SetAuthenticationEmailAddress(client *Client) (*Ok, error) {
	return client.SetAuthenticationEmailAddress(i.EmailAddress)
}

// DeletePassportElement is a helper method for Client.DeletePassportElement
func (i InputPassportElementError) DeletePassportElement(client *Client) (*Ok, error) {
	return client.DeletePassportElement(i.Type)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (i InputPassportElementError) GetPassportElement(client *Client, password string) (*PassportElement, error) {
	return client.GetPassportElement(i.Type, password)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (i InputPassportElementPhoneNumber) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(i.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (i InputPassportElementPhoneNumber) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(i.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (i InputPassportElementPhoneNumber) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(i.PhoneNumber, opts)
}

// AddFavoriteSticker is a helper method for Client.AddFavoriteSticker
func (i InputSticker) AddFavoriteSticker(client *Client) (*Ok, error) {
	return client.AddFavoriteSticker(i.Sticker)
}

// AddRecentSticker is a helper method for Client.AddRecentSticker
func (i InputSticker) AddRecentSticker(client *Client, isAttached bool) (*Stickers, error) {
	return client.AddRecentSticker(isAttached, i.Sticker)
}

// GetStickerEmojis is a helper method for Client.GetStickerEmojis
func (i InputSticker) GetStickerEmojis(client *Client) (*Emojis, error) {
	return client.GetStickerEmojis(i.Sticker)
}

// RemoveFavoriteSticker is a helper method for Client.RemoveFavoriteSticker
func (i InputSticker) RemoveFavoriteSticker(client *Client) (*Ok, error) {
	return client.RemoveFavoriteSticker(i.Sticker)
}

// RemoveRecentSticker is a helper method for Client.RemoveRecentSticker
func (i InputSticker) RemoveRecentSticker(client *Client, isAttached bool) (*Ok, error) {
	return client.RemoveRecentSticker(isAttached, i.Sticker)
}

// RemoveStickerFromSet is a helper method for Client.RemoveStickerFromSet
func (i InputSticker) RemoveStickerFromSet(client *Client) (*Ok, error) {
	return client.RemoveStickerFromSet(i.Sticker)
}

// SearchStickers is a helper method for Client.SearchStickers
func (i InputSticker) SearchStickers(client *Client, stickerType *StickerType, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, i.Emojis, query, inputLanguageCodes, offset, limit)
}

// SetStickerEmojis is a helper method for Client.SetStickerEmojis
func (i InputSticker) SetStickerEmojis(client *Client) (*Ok, error) {
	return client.SetStickerEmojis(i.Sticker, i.Emojis)
}

// SetStickerKeywords is a helper method for Client.SetStickerKeywords
func (i InputSticker) SetStickerKeywords(client *Client) (*Ok, error) {
	return client.SetStickerKeywords(i.Sticker, i.Keywords)
}

// SetStickerMaskPosition is a helper method for Client.SetStickerMaskPosition
func (i InputSticker) SetStickerMaskPosition(client *Client, opts *SetStickerMaskPositionOpts) (*Ok, error) {
	return client.SetStickerMaskPosition(i.Sticker, opts)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (i InputSticker) SetStickerPositionInSet(client *Client, position int32) (*Ok, error) {
	return client.SetStickerPositionInSet(i.Sticker, position)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (i InputSticker) UploadStickerFile(client *Client, userId int64, stickerFormat *StickerFormat) (*File, error) {
	return client.UploadStickerFile(userId, stickerFormat, i.Sticker)
}

// AddQuickReplyShortcutInlineQueryResultMessage is a helper method for Client.AddQuickReplyShortcutInlineQueryResultMessage
func (i InputStoryAreaTypeFoundVenue) AddQuickReplyShortcutInlineQueryResultMessage(client *Client, shortcutName string, replyToMessageId int64, hideViaBot bool) (*QuickReplyMessage, error) {
	return client.AddQuickReplyShortcutInlineQueryResultMessage(shortcutName, replyToMessageId, i.QueryId, i.ResultId, hideViaBot)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InputStoryAreaTypeFoundVenue) SendInlineQueryResultMessage(client *Client, chatId int64, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(chatId, i.QueryId, i.ResultId, hideViaBot, opts)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InputStoryAreaTypeLink) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InputStoryAreaTypeLink) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InputStoryAreaTypeLink) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InputStoryAreaTypeLink) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InputStoryAreaTypeLink) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InputStoryAreaTypeLink) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InputStoryAreaTypeLink) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputStoryAreaTypeLink) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (i InputStoryAreaTypeLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(i.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputStoryAreaTypeLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(i.Location, zoom, width, height, scale, chatId)
}

// SearchPublicStoriesByLocation is a helper method for Client.SearchPublicStoriesByLocation
func (i InputStoryAreaTypeLocation) SearchPublicStoriesByLocation(client *Client, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByLocation(i.Address, offset, limit)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InputStoryAreaTypeMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(i.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i InputStoryAreaTypeMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(i.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (i InputStoryAreaTypeMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(i.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputStoryAreaTypeMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(i.ChatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputStoryAreaTypeMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, i.ChatId, i.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputStoryAreaTypeMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(i.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputStoryAreaTypeMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(i.ChatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputStoryAreaTypeMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(i.ChatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputStoryAreaTypeMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(i.ChatId, i.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (i InputStoryAreaTypeMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(i.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InputStoryAreaTypeMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputStoryAreaTypeMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(i.ChatId, i.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (i InputStoryAreaTypeMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(i.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputStoryAreaTypeMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (i InputStoryAreaTypeMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(i.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (i InputStoryAreaTypeMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(i.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InputStoryAreaTypeMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(i.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputStoryAreaTypeMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(i.ChatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputStoryAreaTypeMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(i.ChatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (i InputStoryAreaTypeMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(i.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputStoryAreaTypeMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputStoryAreaTypeMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(i.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputStoryAreaTypeMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(i.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputStoryAreaTypeMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(i.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputStoryAreaTypeMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(i.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputStoryAreaTypeMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(i.ChatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputStoryAreaTypeMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(i.ChatId, i.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (i InputStoryAreaTypeMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(i.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (i InputStoryAreaTypeMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(i.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (i InputStoryAreaTypeMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(i.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (i InputStoryAreaTypeMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(i.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (i InputStoryAreaTypeMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(i.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (i InputStoryAreaTypeMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(i.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputStoryAreaTypeMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(i.ChatId, i.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (i InputStoryAreaTypeMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(i.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (i InputStoryAreaTypeMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(i.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (i InputStoryAreaTypeMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(i.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (i InputStoryAreaTypeMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(i.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InputStoryAreaTypeMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(i.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InputStoryAreaTypeMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(i.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputStoryAreaTypeMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputStoryAreaTypeMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, i.ChatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputStoryAreaTypeMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputStoryAreaTypeMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputStoryAreaTypeMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputStoryAreaTypeMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputStoryAreaTypeMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(i.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputStoryAreaTypeMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(i.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputStoryAreaTypeMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(i.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputStoryAreaTypeMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputStoryAreaTypeMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(i.ChatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputStoryAreaTypeMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputStoryAreaTypeMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputStoryAreaTypeMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(i.ChatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputStoryAreaTypeMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(i.ChatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputStoryAreaTypeMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputStoryAreaTypeMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InputStoryAreaTypeMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(i.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InputStoryAreaTypeMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, i.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputStoryAreaTypeMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(i.ChatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputStoryAreaTypeMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(i.ChatId, i.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (i InputStoryAreaTypeMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(i.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (i InputStoryAreaTypeMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(i.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (i InputStoryAreaTypeMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(i.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (i InputStoryAreaTypeMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(i.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (i InputStoryAreaTypeMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(i.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (i InputStoryAreaTypeMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(i.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (i InputStoryAreaTypeMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(i.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (i InputStoryAreaTypeMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(i.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (i InputStoryAreaTypeMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(i.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InputStoryAreaTypeMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(i.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (i InputStoryAreaTypeMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(i.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InputStoryAreaTypeMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(i.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (i InputStoryAreaTypeMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(i.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InputStoryAreaTypeMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(i.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (i InputStoryAreaTypeMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(i.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InputStoryAreaTypeMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(i.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (i InputStoryAreaTypeMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(i.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (i InputStoryAreaTypeMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(i.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (i InputStoryAreaTypeMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(i.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (i InputStoryAreaTypeMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(i.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (i InputStoryAreaTypeMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(i.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputStoryAreaTypeMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(i.ChatId, filter, i.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (i InputStoryAreaTypeMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(i.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (i InputStoryAreaTypeMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(i.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InputStoryAreaTypeMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(i.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (i InputStoryAreaTypeMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(i.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (i InputStoryAreaTypeMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(i.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (i InputStoryAreaTypeMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(i.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (i InputStoryAreaTypeMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(i.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (i InputStoryAreaTypeMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(i.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (i InputStoryAreaTypeMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(i.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (i InputStoryAreaTypeMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(i.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InputStoryAreaTypeMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(i.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (i InputStoryAreaTypeMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(i.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (i InputStoryAreaTypeMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(i.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (i InputStoryAreaTypeMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(i.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (i InputStoryAreaTypeMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(i.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (i InputStoryAreaTypeMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(i.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (i InputStoryAreaTypeMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(i.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (i InputStoryAreaTypeMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(i.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (i InputStoryAreaTypeMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(i.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InputStoryAreaTypeMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(i.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputStoryAreaTypeMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(i.ChatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputStoryAreaTypeMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(i.ChatId, i.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InputStoryAreaTypeMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, i.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (i InputStoryAreaTypeMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(i.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputStoryAreaTypeMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(i.ChatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputStoryAreaTypeMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(i.ChatId, i.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InputStoryAreaTypeMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(i.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputStoryAreaTypeMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, i.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputStoryAreaTypeMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(i.ChatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputStoryAreaTypeMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(i.ChatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputStoryAreaTypeMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(i.ChatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputStoryAreaTypeMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(i.ChatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputStoryAreaTypeMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(i.ChatId, i.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (i InputStoryAreaTypeMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(i.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputStoryAreaTypeMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(i.ChatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputStoryAreaTypeMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(i.ChatId, i.MessageId)
}
