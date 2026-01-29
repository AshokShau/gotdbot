package gotdbot

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputStoryAreaTypeMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(i.ChatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputStoryAreaTypeMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(i.ChatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputStoryAreaTypeMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(i.ChatId, i.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (i InputStoryAreaTypeMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(i.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputStoryAreaTypeMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(i.ChatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputStoryAreaTypeMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(i.ChatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputStoryAreaTypeMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(i.ChatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputStoryAreaTypeMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(i.ChatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputStoryAreaTypeMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(i.ChatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputStoryAreaTypeMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(i.ChatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputStoryAreaTypeMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(i.ChatId, i.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InputStoryAreaTypeMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(i.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (i InputStoryAreaTypeMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, i.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InputStoryAreaTypeMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(i.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputStoryAreaTypeMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(i.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InputStoryAreaTypeMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(i.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (i InputStoryAreaTypeMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(i.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (i InputStoryAreaTypeMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(i.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputStoryAreaTypeMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(i.ChatId, i.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InputStoryAreaTypeMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(i.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (i InputStoryAreaTypeMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(i.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (i InputStoryAreaTypeMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(i.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (i InputStoryAreaTypeMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(i.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (i InputStoryAreaTypeMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(i.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputStoryAreaTypeMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(i.ChatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (i InputStoryAreaTypeMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(i.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (i InputStoryAreaTypeMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(i.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputStoryAreaTypeMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(i.ChatId, i.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InputStoryAreaTypeMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(i.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputStoryAreaTypeMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(i.ChatId, i.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (i InputStoryAreaTypeMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(i.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InputStoryAreaTypeMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(i.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InputStoryAreaTypeMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(i.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputStoryAreaTypeMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputStoryAreaTypeMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(i.ChatId, i.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (i InputStoryAreaTypeMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(i.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (i InputStoryAreaTypeMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(i.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (i InputStoryAreaTypeMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(i.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (i InputStoryAreaTypeMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(i.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (i InputStoryAreaTypeMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(i.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputStoryAreaTypeMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, i.ChatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputStoryAreaTypeMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(i.ChatId, i.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (i InputStoryAreaTypeMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(i.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (i InputStoryAreaTypeMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(i.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputStoryAreaTypeMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(i.ChatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputStoryAreaTypeMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (i InputStoryAreaTypeMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(i.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InputStoryAreaTypeMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (i InputStoryAreaTypeMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, i.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (i InputStoryAreaTypeMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(i.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InputStoryAreaTypeMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (i InputStoryAreaTypeMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(i.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (i InputStoryAreaTypeMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(i.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (i InputStoryAreaTypeMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(i.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (i InputStoryAreaTypeMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(i.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InputStoryAreaTypeMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(i.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputStoryAreaTypeMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(i.ChatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputStoryAreaTypeMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(i.ChatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputStoryAreaTypeMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (i InputStoryAreaTypeMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(i.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InputStoryAreaTypeMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(i.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (i InputStoryAreaTypeMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, i.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InputStoryAreaTypeMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(i.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InputStoryAreaTypeMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(i.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (i InputStoryAreaTypeMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(i.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InputStoryAreaTypeMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(i.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InputStoryAreaTypeMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, i.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InputStoryAreaTypeMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (i InputStoryAreaTypeMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (i InputStoryAreaTypeMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(i.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InputStoryAreaTypeMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(i.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InputStoryAreaTypeMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(i.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (i InputStoryAreaTypeMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(i.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (i InputStoryAreaTypeMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(i.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputStoryAreaTypeMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(i.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputStoryAreaTypeMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, i.ChatId, i.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (i InputStoryAreaTypeMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(i.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (i InputStoryAreaTypeMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(i.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (i InputStoryAreaTypeMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(i.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (i InputStoryAreaTypeMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(i.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (i InputStoryAreaTypeMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(i.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (i InputStoryAreaTypeMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(i.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InputStoryAreaTypeMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(i.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (i InputStoryAreaTypeMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(i.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (i InputStoryAreaTypeMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(i.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (i InputStoryAreaTypeMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(i.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (i InputStoryAreaTypeMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(i.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (i InputStoryAreaTypeMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(i.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (i InputStoryAreaTypeMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(i.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (i InputStoryAreaTypeMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(i.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (i InputStoryAreaTypeMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(i.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (i InputStoryAreaTypeMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(i.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (i InputStoryAreaTypeMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(i.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (i InputStoryAreaTypeMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(i.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (i InputStoryAreaTypeMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(i.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (i InputStoryAreaTypeMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(i.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (i InputStoryAreaTypeMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(i.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (i InputStoryAreaTypeMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(i.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (i InputStoryAreaTypeMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(i.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputStoryAreaTypeMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(i.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (i InputStoryAreaTypeMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(i.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (i InputStoryAreaTypeMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(i.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputStoryAreaTypeMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(i.ChatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputStoryAreaTypeMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(i.ChatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputStoryAreaTypeMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(i.ChatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputStoryAreaTypeMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(i.ChatId, i.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (i InputStoryAreaTypeMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(i.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (i InputStoryAreaTypeMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(i.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputStoryAreaTypeMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(i.ChatId, i.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputStoryAreaTypeMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(i.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (i InputStoryAreaTypeMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(i.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputStoryAreaTypeMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(i.ChatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputStoryAreaTypeMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(i.ChatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (i InputStoryAreaTypeMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(i.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputStoryAreaTypeMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputStoryAreaTypeMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(i.ChatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputStoryAreaTypeMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(i.ChatId, i.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (i InputStoryAreaTypeMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(i.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (i InputStoryAreaTypeMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(i.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (i InputStoryAreaTypeMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(i.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (i InputStoryAreaTypeMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(i.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (i InputStoryAreaTypeMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(i.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (i InputStoryAreaTypeMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, i.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (i InputStoryAreaTypeMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(i.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (i InputStoryAreaTypeMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(i.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (i InputStoryAreaTypeMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(i.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InputStoryAreaTypeMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(i.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (i InputStoryAreaTypeMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(i.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (i InputStoryAreaTypeMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(i.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InputStoryAreaTypeMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(i.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputStoryAreaTypeMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(i.ChatId, i.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (i InputStoryAreaTypeMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(i.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (i InputStoryAreaTypeMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(i.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (i InputStoryAreaTypeMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(i.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputStoryAreaTypeMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(i.ChatId, i.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (i InputStoryAreaTypeMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(i.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (i InputStoryAreaTypeMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(i.ChatId, messageIds, forceRead, opts)
}

// SearchPublicStoriesByVenue is a helper method for Client.SearchPublicStoriesByVenue
func (i InputStoryAreaTypePreviousVenue) SearchPublicStoriesByVenue(client *Client, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByVenue(i.VenueProvider, i.VenueId, offset, limit)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputStoryAreaTypeSuggestedReaction) AddMessageReaction(client *Client, chatId int64, messageId int64, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, messageId, i.ReactionType, isBig, updateRecentReactions)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetChatRevenueStatistics(client *Client, chatId int64) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(chatId, i.IsDark)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetChatStatistics(client *Client, chatId int64) (*ChatStatistics, error) {
	return client.GetChatStatistics(chatId, i.IsDark)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetMessageStatistics(client *Client, chatId int64, messageId int64) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, messageId, i.IsDark)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetStarRevenueStatistics(client *Client, ownerId *MessageSender) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(ownerId, i.IsDark)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetStoryStatistics(client *Client, chatId int64, storyId int32) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, storyId, i.IsDark)
}

// GetTonRevenueStatistics is a helper method for Client.GetTonRevenueStatistics
func (i InputStoryAreaTypeSuggestedReaction) GetTonRevenueStatistics(client *Client) (*TonRevenueStatistics, error) {
	return client.GetTonRevenueStatistics(i.IsDark)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputStoryAreaTypeSuggestedReaction) RemoveMessageReaction(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, messageId, i.ReactionType)
}

// SetDefaultReactionType is a helper method for Client.SetDefaultReactionType
func (i InputStoryAreaTypeSuggestedReaction) SetDefaultReactionType(client *Client) (*Ok, error) {
	return client.SetDefaultReactionType(i.ReactionType)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (i InputStoryAreaTypeUpgradedGift) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, i.GiftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (i InputStoryAreaTypeUpgradedGift) SendResoldGift(client *Client, ownerId *MessageSender, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(i.GiftName, ownerId, price)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (i InputStoryAreaTypeWeather) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(i.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (i InputStoryAreaTypeWeather) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(i.Emoji)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (i InputStoryContentVideo) EditStoryCover(client *Client, storyPosterChatId int64, storyId int32) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, storyId, i.CoverFrameTimestamp)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputSuggestedPostInfo) ApproveSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, messageId, i.SendDate)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (i InputTextQuote) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(i.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (i InputTextQuote) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(i.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InputTextQuote) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, i.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (i InputTextQuote) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(i.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InputTextQuote) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, i.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (i InputTextQuote) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(i.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (i InputTextQuote) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, i.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (i InputTextQuote) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, i.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputTextQuote) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, i.Text)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (i InputTextQuote) SetStickerPositionInSet(client *Client, sticker *InputFile) (*Ok, error) {
	return client.SetStickerPositionInSet(sticker, i.Position)
}

// TranslateText is a helper method for Client.TranslateText
func (i InputTextQuote) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(i.Text, toLanguageCode)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InputThumbnail) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, i.Width, i.Height, scale, chatId)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InternalLinkTypeAttachmentMenuBot) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InternalLinkTypeAttachmentMenuBot) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InternalLinkTypeAttachmentMenuBot) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InternalLinkTypeAttachmentMenuBot) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InternalLinkTypeAttachmentMenuBot) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InternalLinkTypeAttachmentMenuBot) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InternalLinkTypeAttachmentMenuBot) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InternalLinkTypeAttachmentMenuBot) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// ApplyPremiumGiftCode is a helper method for Client.ApplyPremiumGiftCode
func (i InternalLinkTypeAuthenticationCode) ApplyPremiumGiftCode(client *Client) (*Ok, error) {
	return client.ApplyPremiumGiftCode(i.Code)
}

// CheckAuthenticationCode is a helper method for Client.CheckAuthenticationCode
func (i InternalLinkTypeAuthenticationCode) CheckAuthenticationCode(client *Client) (*Ok, error) {
	return client.CheckAuthenticationCode(i.Code)
}

// CheckEmailAddressVerificationCode is a helper method for Client.CheckEmailAddressVerificationCode
func (i InternalLinkTypeAuthenticationCode) CheckEmailAddressVerificationCode(client *Client) (*Ok, error) {
	return client.CheckEmailAddressVerificationCode(i.Code)
}

// CheckPhoneNumberCode is a helper method for Client.CheckPhoneNumberCode
func (i InternalLinkTypeAuthenticationCode) CheckPhoneNumberCode(client *Client) (*Ok, error) {
	return client.CheckPhoneNumberCode(i.Code)
}

// CheckPremiumGiftCode is a helper method for Client.CheckPremiumGiftCode
func (i InternalLinkTypeAuthenticationCode) CheckPremiumGiftCode(client *Client) (*PremiumGiftCodeInfo, error) {
	return client.CheckPremiumGiftCode(i.Code)
}

// CheckRecoveryEmailAddressCode is a helper method for Client.CheckRecoveryEmailAddressCode
func (i InternalLinkTypeAuthenticationCode) CheckRecoveryEmailAddressCode(client *Client) (*PasswordState, error) {
	return client.CheckRecoveryEmailAddressCode(i.Code)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InternalLinkTypeBotStart) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, i.StartParameter, parameters)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InternalLinkTypeBotStart) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, webAppShortName string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, webAppShortName, i.StartParameter, allowWriteAccess, parameters)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InternalLinkTypeBotStartInGroup) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, i.StartParameter, parameters)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InternalLinkTypeBotStartInGroup) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, webAppShortName string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, webAppShortName, i.StartParameter, allowWriteAccess, parameters)
}

// GetBusinessChatLinkInfo is a helper method for Client.GetBusinessChatLinkInfo
func (i InternalLinkTypeBusinessChat) GetBusinessChatLinkInfo(client *Client) (*BusinessChatLinkInfo, error) {
	return client.GetBusinessChatLinkInfo(i.LinkName)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (i InternalLinkTypeBuyStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, i.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InternalLinkTypeBuyStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, i.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (i InternalLinkTypeBuyStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, i.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (i InternalLinkTypeBuyStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, i.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (i InternalLinkTypeBuyStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, i.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InternalLinkTypeBuyStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, i.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (i InternalLinkTypeBuyStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, i.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (i InternalLinkTypeBuyStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, i.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InternalLinkTypeBuyStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, i.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (i InternalLinkTypeBuyStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, i.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (i InternalLinkTypeBuyStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, i.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (i InternalLinkTypeBuyStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, i.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (i InternalLinkTypeBuyStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, i.StarCount)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InternalLinkTypeChatAffiliateProgram) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, i.Username)
}

// GetRecentlyVisitedTMeUrls is a helper method for Client.GetRecentlyVisitedTMeUrls
func (i InternalLinkTypeChatAffiliateProgram) GetRecentlyVisitedTMeUrls(client *Client) (*TMeUrls, error) {
	return client.GetRecentlyVisitedTMeUrls(i.Referrer)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (i InternalLinkTypeChatAffiliateProgram) SearchChatAffiliateProgram(client *Client) (*Chat, error) {
	return client.SearchChatAffiliateProgram(i.Username, i.Referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (i InternalLinkTypeChatAffiliateProgram) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(i.Username)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (i InternalLinkTypeChatAffiliateProgram) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, i.Username)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (i InternalLinkTypeChatAffiliateProgram) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, i.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (i InternalLinkTypeChatAffiliateProgram) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(i.Username)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (i InternalLinkTypeChatAffiliateProgram) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, i.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (i InternalLinkTypeChatAffiliateProgram) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, i.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (i InternalLinkTypeChatAffiliateProgram) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(i.Username, isActive)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InternalLinkTypeChatBoost) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InternalLinkTypeChatBoost) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InternalLinkTypeChatBoost) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InternalLinkTypeChatBoost) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InternalLinkTypeChatBoost) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InternalLinkTypeChatBoost) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InternalLinkTypeChatBoost) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InternalLinkTypeChatBoost) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (i InternalLinkTypeChatFolderInvite) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(i.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (i InternalLinkTypeChatFolderInvite) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(i.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (i InternalLinkTypeChatFolderInvite) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(i.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (i InternalLinkTypeChatFolderInvite) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, i.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InternalLinkTypeChatFolderInvite) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, i.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InternalLinkTypeChatFolderInvite) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, i.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InternalLinkTypeChatFolderInvite) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, i.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InternalLinkTypeChatFolderInvite) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, i.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InternalLinkTypeChatFolderInvite) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, i.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InternalLinkTypeChatFolderInvite) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, i.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InternalLinkTypeChatFolderInvite) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, i.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InternalLinkTypeChatFolderInvite) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, i.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (i InternalLinkTypeChatFolderInvite) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(i.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InternalLinkTypeChatFolderInvite) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, i.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InternalLinkTypeChatFolderInvite) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, i.InviteLink)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (i InternalLinkTypeChatInvite) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(i.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (i InternalLinkTypeChatInvite) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(i.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (i InternalLinkTypeChatInvite) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(i.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (i InternalLinkTypeChatInvite) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, i.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InternalLinkTypeChatInvite) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, i.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InternalLinkTypeChatInvite) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, i.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InternalLinkTypeChatInvite) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, i.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InternalLinkTypeChatInvite) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, i.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InternalLinkTypeChatInvite) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, i.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InternalLinkTypeChatInvite) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, i.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InternalLinkTypeChatInvite) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, i.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InternalLinkTypeChatInvite) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, i.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (i InternalLinkTypeChatInvite) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(i.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InternalLinkTypeChatInvite) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, i.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InternalLinkTypeChatInvite) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, i.InviteLink)
}

// GetGiftAuctionState is a helper method for Client.GetGiftAuctionState
func (i InternalLinkTypeGiftAuction) GetGiftAuctionState(client *Client) (*GiftAuctionState, error) {
	return client.GetGiftAuctionState(i.AuctionId)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (i InternalLinkTypeGiftCollection) AddGiftCollectionGifts(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(ownerId, i.CollectionId, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (i InternalLinkTypeGiftCollection) DeleteGiftCollection(client *Client, ownerId *MessageSender) (*Ok, error) {
	return client.DeleteGiftCollection(ownerId, i.CollectionId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (i InternalLinkTypeGiftCollection) GetReceivedGifts(client *Client, businessConnectionId string, ownerId *MessageSender, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, ownerId, i.CollectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (i InternalLinkTypeGiftCollection) RemoveGiftCollectionGifts(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(ownerId, i.CollectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (i InternalLinkTypeGiftCollection) ReorderGiftCollectionGifts(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(ownerId, i.CollectionId, receivedGiftIds)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (i InternalLinkTypeGiftCollection) SetGiftCollectionName(client *Client, ownerId *MessageSender, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, i.CollectionId, name)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (i InternalLinkTypeGroupCall) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(i.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (i InternalLinkTypeGroupCall) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(i.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (i InternalLinkTypeGroupCall) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(i.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (i InternalLinkTypeGroupCall) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, i.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InternalLinkTypeGroupCall) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, i.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (i InternalLinkTypeGroupCall) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, i.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InternalLinkTypeGroupCall) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, i.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InternalLinkTypeGroupCall) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, i.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InternalLinkTypeGroupCall) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, i.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InternalLinkTypeGroupCall) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, i.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InternalLinkTypeGroupCall) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, i.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InternalLinkTypeGroupCall) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, i.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (i InternalLinkTypeGroupCall) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(i.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InternalLinkTypeGroupCall) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, i.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InternalLinkTypeGroupCall) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, i.InviteLink)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InternalLinkTypeInstantView) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InternalLinkTypeInstantView) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InternalLinkTypeInstantView) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InternalLinkTypeInstantView) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InternalLinkTypeInstantView) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InternalLinkTypeInstantView) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InternalLinkTypeInstantView) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InternalLinkTypeInstantView) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// AddCustomServerLanguagePack is a helper method for Client.AddCustomServerLanguagePack
func (i InternalLinkTypeLanguagePack) AddCustomServerLanguagePack(client *Client) (*Ok, error) {
	return client.AddCustomServerLanguagePack(i.LanguagePackId)
}

// DeleteLanguagePack is a helper method for Client.DeleteLanguagePack
func (i InternalLinkTypeLanguagePack) DeleteLanguagePack(client *Client) (*Ok, error) {
	return client.DeleteLanguagePack(i.LanguagePackId)
}

// GetLanguagePackInfo is a helper method for Client.GetLanguagePackInfo
func (i InternalLinkTypeLanguagePack) GetLanguagePackInfo(client *Client) (*LanguagePackInfo, error) {
	return client.GetLanguagePackInfo(i.LanguagePackId)
}

// GetLanguagePackString is a helper method for Client.GetLanguagePackString
func (i InternalLinkTypeLanguagePack) GetLanguagePackString(client *Client, languagePackDatabasePath string, localizationTarget string, key string) (*LanguagePackStringValue, error) {
	return client.GetLanguagePackString(languagePackDatabasePath, localizationTarget, i.LanguagePackId, key)
}

// GetLanguagePackStrings is a helper method for Client.GetLanguagePackStrings
func (i InternalLinkTypeLanguagePack) GetLanguagePackStrings(client *Client, keys []string) (*LanguagePackStrings, error) {
	return client.GetLanguagePackStrings(i.LanguagePackId, keys)
}

// SetCustomLanguagePackString is a helper method for Client.SetCustomLanguagePackString
func (i InternalLinkTypeLanguagePack) SetCustomLanguagePackString(client *Client, newString *LanguagePackString) (*Ok, error) {
	return client.SetCustomLanguagePackString(i.LanguagePackId, newString)
}

// SynchronizeLanguagePack is a helper method for Client.SynchronizeLanguagePack
func (i InternalLinkTypeLanguagePack) SynchronizeLanguagePack(client *Client) (*Ok, error) {
	return client.SynchronizeLanguagePack(i.LanguagePackId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InternalLinkTypeMainWebApp) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, i.StartParameter, parameters)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InternalLinkTypeMainWebApp) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, webAppShortName string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, webAppShortName, i.StartParameter, allowWriteAccess, parameters)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InternalLinkTypeMessage) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InternalLinkTypeMessage) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InternalLinkTypeMessage) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InternalLinkTypeMessage) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InternalLinkTypeMessage) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InternalLinkTypeMessage) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InternalLinkTypeMessage) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InternalLinkTypeMessage) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (i InternalLinkTypeMessageDraft) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(i.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (i InternalLinkTypeMessageDraft) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(i.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InternalLinkTypeMessageDraft) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, i.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (i InternalLinkTypeMessageDraft) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(i.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InternalLinkTypeMessageDraft) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, i.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (i InternalLinkTypeMessageDraft) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(i.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (i InternalLinkTypeMessageDraft) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, i.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (i InternalLinkTypeMessageDraft) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, i.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InternalLinkTypeMessageDraft) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, i.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (i InternalLinkTypeMessageDraft) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(i.Text, toLanguageCode)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (i InternalLinkTypePassportDataRequest) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(i.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (i InternalLinkTypePassportDataRequest) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(i.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (i InternalLinkTypePassportDataRequest) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(i.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InternalLinkTypePassportDataRequest) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(i.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (i InternalLinkTypePassportDataRequest) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, i.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (i InternalLinkTypePassportDataRequest) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(i.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (i InternalLinkTypePassportDataRequest) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(i.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (i InternalLinkTypePassportDataRequest) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(i.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (i InternalLinkTypePassportDataRequest) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(i.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (i InternalLinkTypePassportDataRequest) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(i.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (i InternalLinkTypePassportDataRequest) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(i.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (i InternalLinkTypePassportDataRequest) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(i.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (i InternalLinkTypePassportDataRequest) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(i.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (i InternalLinkTypePassportDataRequest) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(i.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (i InternalLinkTypePassportDataRequest) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(i.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (i InternalLinkTypePassportDataRequest) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(i.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (i InternalLinkTypePassportDataRequest) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, i.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InternalLinkTypePassportDataRequest) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(i.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InternalLinkTypePassportDataRequest) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, i.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (i InternalLinkTypePassportDataRequest) GetPassportAuthorizationForm(client *Client) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(i.BotUserId, i.Scope, i.PublicKey, i.Nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (i InternalLinkTypePassportDataRequest) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(i.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InternalLinkTypePassportDataRequest) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, i.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (i InternalLinkTypePassportDataRequest) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(i.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InternalLinkTypePassportDataRequest) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(i.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (i InternalLinkTypePassportDataRequest) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(i.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InternalLinkTypePassportDataRequest) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, i.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (i InternalLinkTypePassportDataRequest) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(i.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (i InternalLinkTypePassportDataRequest) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(i.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (i InternalLinkTypePassportDataRequest) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(i.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (i InternalLinkTypePassportDataRequest) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(i.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InternalLinkTypePassportDataRequest) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(i.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (i InternalLinkTypePassportDataRequest) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(i.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (i InternalLinkTypePassportDataRequest) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(i.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InternalLinkTypePassportDataRequest) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(i.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (i InternalLinkTypePassportDataRequest) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(i.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (i InternalLinkTypePassportDataRequest) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(i.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (i InternalLinkTypePassportDataRequest) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(i.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (i InternalLinkTypePassportDataRequest) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(i.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (i InternalLinkTypePassportDataRequest) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(i.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (i InternalLinkTypePassportDataRequest) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(i.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (i InternalLinkTypePassportDataRequest) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(i.BotUserId, username, isActive)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (i InternalLinkTypePhoneNumberConfirmation) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(i.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (i InternalLinkTypePhoneNumberConfirmation) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(i.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (i InternalLinkTypePhoneNumberConfirmation) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(i.PhoneNumber, opts)
}

// GetRecentlyVisitedTMeUrls is a helper method for Client.GetRecentlyVisitedTMeUrls
func (i InternalLinkTypePremiumFeatures) GetRecentlyVisitedTMeUrls(client *Client) (*TMeUrls, error) {
	return client.GetRecentlyVisitedTMeUrls(i.Referrer)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (i InternalLinkTypePremiumFeatures) SearchChatAffiliateProgram(client *Client, username string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(username, i.Referrer)
}

// GetRecentlyVisitedTMeUrls is a helper method for Client.GetRecentlyVisitedTMeUrls
func (i InternalLinkTypePremiumGift) GetRecentlyVisitedTMeUrls(client *Client) (*TMeUrls, error) {
	return client.GetRecentlyVisitedTMeUrls(i.Referrer)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (i InternalLinkTypePremiumGift) SearchChatAffiliateProgram(client *Client, username string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(username, i.Referrer)
}

// ApplyPremiumGiftCode is a helper method for Client.ApplyPremiumGiftCode
func (i InternalLinkTypePremiumGiftCode) ApplyPremiumGiftCode(client *Client) (*Ok, error) {
	return client.ApplyPremiumGiftCode(i.Code)
}

// CheckAuthenticationCode is a helper method for Client.CheckAuthenticationCode
func (i InternalLinkTypePremiumGiftCode) CheckAuthenticationCode(client *Client) (*Ok, error) {
	return client.CheckAuthenticationCode(i.Code)
}

// CheckEmailAddressVerificationCode is a helper method for Client.CheckEmailAddressVerificationCode
func (i InternalLinkTypePremiumGiftCode) CheckEmailAddressVerificationCode(client *Client) (*Ok, error) {
	return client.CheckEmailAddressVerificationCode(i.Code)
}

// CheckPhoneNumberCode is a helper method for Client.CheckPhoneNumberCode
func (i InternalLinkTypePremiumGiftCode) CheckPhoneNumberCode(client *Client) (*Ok, error) {
	return client.CheckPhoneNumberCode(i.Code)
}

// CheckPremiumGiftCode is a helper method for Client.CheckPremiumGiftCode
func (i InternalLinkTypePremiumGiftCode) CheckPremiumGiftCode(client *Client) (*PremiumGiftCodeInfo, error) {
	return client.CheckPremiumGiftCode(i.Code)
}

// CheckRecoveryEmailAddressCode is a helper method for Client.CheckRecoveryEmailAddressCode
func (i InternalLinkTypePremiumGiftCode) CheckRecoveryEmailAddressCode(client *Client) (*PasswordState, error) {
	return client.CheckRecoveryEmailAddressCode(i.Code)
}

// AddProxy is a helper method for Client.AddProxy
func (i InternalLinkTypeProxy) AddProxy(client *Client, enable bool) (*Proxy, error) {
	return client.AddProxy(i.Server, i.Port, enable, i.Type)
}

// EditProxy is a helper method for Client.EditProxy
func (i InternalLinkTypeProxy) EditProxy(client *Client, proxyId int32, enable bool) (*Proxy, error) {
	return client.EditProxy(proxyId, i.Server, i.Port, enable, i.Type)
}

// TestProxy is a helper method for Client.TestProxy
func (i InternalLinkTypeProxy) TestProxy(client *Client, dcId int32, timeout float64) (*Ok, error) {
	return client.TestProxy(i.Server, i.Port, i.Type, dcId, timeout)
}

// CloseStory is a helper method for Client.CloseStory
func (i InternalLinkTypeStory) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, i.StoryId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (i InternalLinkTypeStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, i.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (i InternalLinkTypeStory) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, i.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (i InternalLinkTypeStory) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, i.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (i InternalLinkTypeStory) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, i.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (i InternalLinkTypeStory) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, i.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (i InternalLinkTypeStory) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, i.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (i InternalLinkTypeStory) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, i.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (i InternalLinkTypeStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(i.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (i InternalLinkTypeStory) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, i.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InternalLinkTypeStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, i.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (i InternalLinkTypeStory) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, i.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (i InternalLinkTypeStory) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, i.StoryId, optionId, text)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (i InternalLinkTypeStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(i.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (i InternalLinkTypeStory) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, i.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (i InternalLinkTypeStory) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, i.StoryId, isPostedToChatPage)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InternalLinkTypeStoryAlbum) AddStoryAlbumStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(chatId, i.StoryAlbumId, storyIds)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InternalLinkTypeStoryAlbum) DeleteStoryAlbum(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteStoryAlbum(chatId, i.StoryAlbumId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InternalLinkTypeStoryAlbum) GetStoryAlbumStories(client *Client, chatId int64, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(chatId, i.StoryAlbumId, offset, limit)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InternalLinkTypeStoryAlbum) RemoveStoryAlbumStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(chatId, i.StoryAlbumId, storyIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InternalLinkTypeStoryAlbum) ReorderStoryAlbumStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(chatId, i.StoryAlbumId, storyIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InternalLinkTypeStoryAlbum) SetStoryAlbumName(client *Client, chatId int64, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, i.StoryAlbumId, name)
}

// ConfirmQrCodeAuthentication is a helper method for Client.ConfirmQrCodeAuthentication
func (i InternalLinkTypeUnknownDeepLink) ConfirmQrCodeAuthentication(client *Client) (*Session, error) {
	return client.ConfirmQrCodeAuthentication(i.Link)
}

// DeleteBusinessChatLink is a helper method for Client.DeleteBusinessChatLink
func (i InternalLinkTypeUnknownDeepLink) DeleteBusinessChatLink(client *Client) (*Ok, error) {
	return client.DeleteBusinessChatLink(i.Link)
}

// EditBusinessChatLink is a helper method for Client.EditBusinessChatLink
func (i InternalLinkTypeUnknownDeepLink) EditBusinessChatLink(client *Client, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	return client.EditBusinessChatLink(i.Link, linkInfo)
}

// GetDeepLinkInfo is a helper method for Client.GetDeepLinkInfo
func (i InternalLinkTypeUnknownDeepLink) GetDeepLinkInfo(client *Client) (*DeepLinkInfo, error) {
	return client.GetDeepLinkInfo(i.Link)
}

// GetExternalLink is a helper method for Client.GetExternalLink
func (i InternalLinkTypeUnknownDeepLink) GetExternalLink(client *Client, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetExternalLink(i.Link, allowWriteAccess)
}

// GetExternalLinkInfo is a helper method for Client.GetExternalLinkInfo
func (i InternalLinkTypeUnknownDeepLink) GetExternalLinkInfo(client *Client) (*LoginUrlInfo, error) {
	return client.GetExternalLinkInfo(i.Link)
}

// GetInternalLinkType is a helper method for Client.GetInternalLinkType
func (i InternalLinkTypeUnknownDeepLink) GetInternalLinkType(client *Client) (*InternalLinkType, error) {
	return client.GetInternalLinkType(i.Link)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (i InternalLinkTypeUpgradedGift) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, i.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (i InternalLinkTypeUpgradedGift) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(i.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (i InternalLinkTypeUpgradedGift) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(i.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (i InternalLinkTypeUpgradedGift) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, i.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InternalLinkTypeUpgradedGift) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InternalLinkTypeUpgradedGift) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, i.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InternalLinkTypeUpgradedGift) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, i.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (i InternalLinkTypeUpgradedGift) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, i.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InternalLinkTypeUpgradedGift) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, i.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InternalLinkTypeUpgradedGift) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, i.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (i InternalLinkTypeUpgradedGift) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(i.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InternalLinkTypeUpgradedGift) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, i.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InternalLinkTypeUpgradedGift) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InternalLinkTypeUpgradedGift) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, i.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InternalLinkTypeUpgradedGift) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, i.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (i InternalLinkTypeUpgradedGift) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(i.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (i InternalLinkTypeUpgradedGift) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(i.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (i InternalLinkTypeUpgradedGift) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(i.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (i InternalLinkTypeUpgradedGift) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(i.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (i InternalLinkTypeUpgradedGift) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, i.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (i InternalLinkTypeUpgradedGift) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(i.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (i InternalLinkTypeUpgradedGift) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(i.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (i InternalLinkTypeUpgradedGift) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, i.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (i InternalLinkTypeUpgradedGift) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(i.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (i InternalLinkTypeUpgradedGift) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, i.Name)
}

// SetOption is a helper method for Client.SetOption
func (i InternalLinkTypeUpgradedGift) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(i.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (i InternalLinkTypeUpgradedGift) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, i.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (i InternalLinkTypeUpgradedGift) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, i.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InternalLinkTypeUpgradedGift) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(i.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InternalLinkTypeUpgradedGift) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, i.Name)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (i InternalLinkTypeUserPhoneNumber) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(i.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (i InternalLinkTypeUserPhoneNumber) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(i.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (i InternalLinkTypeUserPhoneNumber) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(i.PhoneNumber, opts)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (i InternalLinkTypeUserToken) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(i.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InternalLinkTypeUserToken) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, i.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (i InternalLinkTypeUserToken) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(i.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (i InternalLinkTypeUserToken) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(i.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (i InternalLinkTypeUserToken) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(i.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (i InternalLinkTypeUserToken) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, i.Token)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (i InternalLinkTypeVideoChat) JoinVideoChat(client *Client, groupCallId int32, joinParameters *GroupCallJoinParameters, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(groupCallId, joinParameters, i.InviteHash, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InternalLinkTypeWebApp) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, i.StartParameter, parameters)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InternalLinkTypeWebApp) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, i.WebAppShortName, i.StartParameter, allowWriteAccess, parameters)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (i InternalLinkTypeWebApp) SearchWebApp(client *Client, botUserId int64) (*FoundWebApp, error) {
	return client.SearchWebApp(botUserId, i.WebAppShortName)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InviteGroupCallParticipantResultSuccess) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(i.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i InviteGroupCallParticipantResultSuccess) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(i.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (i InviteGroupCallParticipantResultSuccess) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(i.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InviteGroupCallParticipantResultSuccess) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(i.ChatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InviteGroupCallParticipantResultSuccess) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, i.ChatId, i.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InviteGroupCallParticipantResultSuccess) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(i.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InviteGroupCallParticipantResultSuccess) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(i.ChatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InviteGroupCallParticipantResultSuccess) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(i.ChatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InviteGroupCallParticipantResultSuccess) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(i.ChatId, i.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (i InviteGroupCallParticipantResultSuccess) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(i.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (i InviteGroupCallParticipantResultSuccess) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InviteGroupCallParticipantResultSuccess) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(i.ChatId, i.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (i InviteGroupCallParticipantResultSuccess) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(i.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InviteGroupCallParticipantResultSuccess) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (i InviteGroupCallParticipantResultSuccess) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(i.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (i InviteGroupCallParticipantResultSuccess) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(i.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (i InviteGroupCallParticipantResultSuccess) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(i.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InviteGroupCallParticipantResultSuccess) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(i.ChatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InviteGroupCallParticipantResultSuccess) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(i.ChatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (i InviteGroupCallParticipantResultSuccess) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(i.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InviteGroupCallParticipantResultSuccess) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(i.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InviteGroupCallParticipantResultSuccess) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(i.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InviteGroupCallParticipantResultSuccess) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(i.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InviteGroupCallParticipantResultSuccess) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(i.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InviteGroupCallParticipantResultSuccess) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(i.ChatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InviteGroupCallParticipantResultSuccess) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(i.ChatId, i.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (i InviteGroupCallParticipantResultSuccess) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(i.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (i InviteGroupCallParticipantResultSuccess) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(i.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (i InviteGroupCallParticipantResultSuccess) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(i.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (i InviteGroupCallParticipantResultSuccess) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(i.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (i InviteGroupCallParticipantResultSuccess) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(i.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (i InviteGroupCallParticipantResultSuccess) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(i.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InviteGroupCallParticipantResultSuccess) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(i.ChatId, i.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (i InviteGroupCallParticipantResultSuccess) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(i.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (i InviteGroupCallParticipantResultSuccess) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(i.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (i InviteGroupCallParticipantResultSuccess) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(i.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (i InviteGroupCallParticipantResultSuccess) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(i.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(i.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (i InviteGroupCallParticipantResultSuccess) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(i.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, i.ChatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InviteGroupCallParticipantResultSuccess) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(i.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InviteGroupCallParticipantResultSuccess) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(i.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InviteGroupCallParticipantResultSuccess) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(i.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InviteGroupCallParticipantResultSuccess) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(i.ChatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InviteGroupCallParticipantResultSuccess) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(i.ChatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InviteGroupCallParticipantResultSuccess) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(i.ChatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InviteGroupCallParticipantResultSuccess) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InviteGroupCallParticipantResultSuccess) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(i.ChatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InviteGroupCallParticipantResultSuccess) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(i.ChatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InviteGroupCallParticipantResultSuccess) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(i.ChatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InviteGroupCallParticipantResultSuccess) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (i InviteGroupCallParticipantResultSuccess) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(i.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InviteGroupCallParticipantResultSuccess) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, i.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InviteGroupCallParticipantResultSuccess) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(i.ChatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InviteGroupCallParticipantResultSuccess) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(i.ChatId, i.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (i InviteGroupCallParticipantResultSuccess) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(i.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (i InviteGroupCallParticipantResultSuccess) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(i.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (i InviteGroupCallParticipantResultSuccess) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(i.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (i InviteGroupCallParticipantResultSuccess) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(i.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (i InviteGroupCallParticipantResultSuccess) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(i.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (i InviteGroupCallParticipantResultSuccess) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(i.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (i InviteGroupCallParticipantResultSuccess) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(i.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (i InviteGroupCallParticipantResultSuccess) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(i.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (i InviteGroupCallParticipantResultSuccess) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(i.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InviteGroupCallParticipantResultSuccess) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(i.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (i InviteGroupCallParticipantResultSuccess) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(i.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(i.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (i InviteGroupCallParticipantResultSuccess) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(i.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (i InviteGroupCallParticipantResultSuccess) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(i.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (i InviteGroupCallParticipantResultSuccess) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(i.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InviteGroupCallParticipantResultSuccess) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(i.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (i InviteGroupCallParticipantResultSuccess) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(i.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (i InviteGroupCallParticipantResultSuccess) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(i.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (i InviteGroupCallParticipantResultSuccess) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(i.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (i InviteGroupCallParticipantResultSuccess) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(i.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (i InviteGroupCallParticipantResultSuccess) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(i.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InviteGroupCallParticipantResultSuccess) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(i.ChatId, filter, i.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (i InviteGroupCallParticipantResultSuccess) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(i.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (i InviteGroupCallParticipantResultSuccess) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(i.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (i InviteGroupCallParticipantResultSuccess) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(i.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (i InviteGroupCallParticipantResultSuccess) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(i.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (i InviteGroupCallParticipantResultSuccess) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(i.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (i InviteGroupCallParticipantResultSuccess) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(i.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (i InviteGroupCallParticipantResultSuccess) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(i.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (i InviteGroupCallParticipantResultSuccess) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(i.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (i InviteGroupCallParticipantResultSuccess) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(i.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (i InviteGroupCallParticipantResultSuccess) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(i.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (i InviteGroupCallParticipantResultSuccess) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(i.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (i InviteGroupCallParticipantResultSuccess) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(i.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (i InviteGroupCallParticipantResultSuccess) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(i.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (i InviteGroupCallParticipantResultSuccess) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(i.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (i InviteGroupCallParticipantResultSuccess) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(i.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (i InviteGroupCallParticipantResultSuccess) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(i.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (i InviteGroupCallParticipantResultSuccess) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(i.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (i InviteGroupCallParticipantResultSuccess) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(i.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (i InviteGroupCallParticipantResultSuccess) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(i.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InviteGroupCallParticipantResultSuccess) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(i.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InviteGroupCallParticipantResultSuccess) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(i.ChatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InviteGroupCallParticipantResultSuccess) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(i.ChatId, i.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InviteGroupCallParticipantResultSuccess) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, i.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (i InviteGroupCallParticipantResultSuccess) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(i.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InviteGroupCallParticipantResultSuccess) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(i.ChatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InviteGroupCallParticipantResultSuccess) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(i.ChatId, i.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (i InviteGroupCallParticipantResultSuccess) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(i.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InviteGroupCallParticipantResultSuccess) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, i.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InviteGroupCallParticipantResultSuccess) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(i.ChatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InviteGroupCallParticipantResultSuccess) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(i.ChatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InviteGroupCallParticipantResultSuccess) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(i.ChatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InviteGroupCallParticipantResultSuccess) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(i.ChatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InviteGroupCallParticipantResultSuccess) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(i.ChatId, i.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (i InviteGroupCallParticipantResultSuccess) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(i.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InviteGroupCallParticipantResultSuccess) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(i.ChatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InviteGroupCallParticipantResultSuccess) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(i.ChatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InviteGroupCallParticipantResultSuccess) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(i.ChatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InviteGroupCallParticipantResultSuccess) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(i.ChatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InviteGroupCallParticipantResultSuccess) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(i.ChatId, i.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (i InviteGroupCallParticipantResultSuccess) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(i.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InviteGroupCallParticipantResultSuccess) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(i.ChatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InviteGroupCallParticipantResultSuccess) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(i.ChatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InviteGroupCallParticipantResultSuccess) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(i.ChatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InviteGroupCallParticipantResultSuccess) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(i.ChatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InviteGroupCallParticipantResultSuccess) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(i.ChatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InviteGroupCallParticipantResultSuccess) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(i.ChatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InviteGroupCallParticipantResultSuccess) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(i.ChatId, i.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (i InviteGroupCallParticipantResultSuccess) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(i.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (i InviteGroupCallParticipantResultSuccess) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, i.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (i InviteGroupCallParticipantResultSuccess) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(i.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (i InviteGroupCallParticipantResultSuccess) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(i.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InviteGroupCallParticipantResultSuccess) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(i.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (i InviteGroupCallParticipantResultSuccess) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(i.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (i InviteGroupCallParticipantResultSuccess) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(i.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InviteGroupCallParticipantResultSuccess) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(i.ChatId, i.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (i InviteGroupCallParticipantResultSuccess) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(i.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (i InviteGroupCallParticipantResultSuccess) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(i.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (i InviteGroupCallParticipantResultSuccess) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(i.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (i InviteGroupCallParticipantResultSuccess) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(i.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (i InviteGroupCallParticipantResultSuccess) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(i.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InviteGroupCallParticipantResultSuccess) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(i.ChatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (i InviteGroupCallParticipantResultSuccess) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(i.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (i InviteGroupCallParticipantResultSuccess) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(i.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InviteGroupCallParticipantResultSuccess) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(i.ChatId, i.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InviteGroupCallParticipantResultSuccess) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(i.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InviteGroupCallParticipantResultSuccess) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(i.ChatId, i.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (i InviteGroupCallParticipantResultSuccess) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(i.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InviteGroupCallParticipantResultSuccess) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(i.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (i InviteGroupCallParticipantResultSuccess) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(i.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InviteGroupCallParticipantResultSuccess) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InviteGroupCallParticipantResultSuccess) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(i.ChatId, i.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (i InviteGroupCallParticipantResultSuccess) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(i.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (i InviteGroupCallParticipantResultSuccess) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(i.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (i InviteGroupCallParticipantResultSuccess) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(i.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (i InviteGroupCallParticipantResultSuccess) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(i.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (i InviteGroupCallParticipantResultSuccess) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(i.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InviteGroupCallParticipantResultSuccess) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, i.ChatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InviteGroupCallParticipantResultSuccess) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(i.ChatId, i.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (i InviteGroupCallParticipantResultSuccess) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(i.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (i InviteGroupCallParticipantResultSuccess) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(i.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InviteGroupCallParticipantResultSuccess) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(i.ChatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InviteGroupCallParticipantResultSuccess) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(i.ChatId, i.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (i InviteGroupCallParticipantResultSuccess) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(i.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (i InviteGroupCallParticipantResultSuccess) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (i InviteGroupCallParticipantResultSuccess) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, i.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (i InviteGroupCallParticipantResultSuccess) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(i.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (i InviteGroupCallParticipantResultSuccess) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(i.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (i InviteGroupCallParticipantResultSuccess) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(i.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(i.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (i InviteGroupCallParticipantResultSuccess) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(i.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (i InviteGroupCallParticipantResultSuccess) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(i.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InviteGroupCallParticipantResultSuccess) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(i.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InviteGroupCallParticipantResultSuccess) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(i.ChatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InviteGroupCallParticipantResultSuccess) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(i.ChatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InviteGroupCallParticipantResultSuccess) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (i InviteGroupCallParticipantResultSuccess) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(i.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (i InviteGroupCallParticipantResultSuccess) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(i.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (i InviteGroupCallParticipantResultSuccess) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, i.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InviteGroupCallParticipantResultSuccess) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(i.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InviteGroupCallParticipantResultSuccess) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(i.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (i InviteGroupCallParticipantResultSuccess) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(i.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InviteGroupCallParticipantResultSuccess) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(i.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InviteGroupCallParticipantResultSuccess) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, i.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (i InviteGroupCallParticipantResultSuccess) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (i InviteGroupCallParticipantResultSuccess) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, i.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (i InviteGroupCallParticipantResultSuccess) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(i.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (i InviteGroupCallParticipantResultSuccess) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(i.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (i InviteGroupCallParticipantResultSuccess) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(i.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (i InviteGroupCallParticipantResultSuccess) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(i.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (i InviteGroupCallParticipantResultSuccess) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(i.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InviteGroupCallParticipantResultSuccess) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(i.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InviteGroupCallParticipantResultSuccess) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, i.ChatId, i.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (i InviteGroupCallParticipantResultSuccess) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(i.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (i InviteGroupCallParticipantResultSuccess) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(i.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (i InviteGroupCallParticipantResultSuccess) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(i.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (i InviteGroupCallParticipantResultSuccess) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(i.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (i InviteGroupCallParticipantResultSuccess) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(i.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (i InviteGroupCallParticipantResultSuccess) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(i.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InviteGroupCallParticipantResultSuccess) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(i.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (i InviteGroupCallParticipantResultSuccess) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(i.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (i InviteGroupCallParticipantResultSuccess) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(i.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (i InviteGroupCallParticipantResultSuccess) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(i.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (i InviteGroupCallParticipantResultSuccess) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(i.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (i InviteGroupCallParticipantResultSuccess) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(i.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (i InviteGroupCallParticipantResultSuccess) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(i.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (i InviteGroupCallParticipantResultSuccess) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(i.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (i InviteGroupCallParticipantResultSuccess) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(i.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (i InviteGroupCallParticipantResultSuccess) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(i.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (i InviteGroupCallParticipantResultSuccess) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(i.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (i InviteGroupCallParticipantResultSuccess) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(i.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (i InviteGroupCallParticipantResultSuccess) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(i.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (i InviteGroupCallParticipantResultSuccess) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(i.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (i InviteGroupCallParticipantResultSuccess) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(i.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (i InviteGroupCallParticipantResultSuccess) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(i.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (i InviteGroupCallParticipantResultSuccess) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(i.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InviteGroupCallParticipantResultSuccess) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(i.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (i InviteGroupCallParticipantResultSuccess) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(i.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (i InviteGroupCallParticipantResultSuccess) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(i.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InviteGroupCallParticipantResultSuccess) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(i.ChatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InviteGroupCallParticipantResultSuccess) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(i.ChatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InviteGroupCallParticipantResultSuccess) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(i.ChatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InviteGroupCallParticipantResultSuccess) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(i.ChatId, i.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (i InviteGroupCallParticipantResultSuccess) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(i.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (i InviteGroupCallParticipantResultSuccess) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(i.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InviteGroupCallParticipantResultSuccess) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(i.ChatId, i.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InviteGroupCallParticipantResultSuccess) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(i.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (i InviteGroupCallParticipantResultSuccess) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(i.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InviteGroupCallParticipantResultSuccess) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(i.ChatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InviteGroupCallParticipantResultSuccess) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(i.ChatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (i InviteGroupCallParticipantResultSuccess) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(i.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InviteGroupCallParticipantResultSuccess) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, i.ChatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InviteGroupCallParticipantResultSuccess) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(i.ChatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InviteGroupCallParticipantResultSuccess) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(i.ChatId, i.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (i InviteGroupCallParticipantResultSuccess) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(i.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (i InviteGroupCallParticipantResultSuccess) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(i.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (i InviteGroupCallParticipantResultSuccess) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(i.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (i InviteGroupCallParticipantResultSuccess) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(i.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (i InviteGroupCallParticipantResultSuccess) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(i.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (i InviteGroupCallParticipantResultSuccess) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, i.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (i InviteGroupCallParticipantResultSuccess) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(i.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (i InviteGroupCallParticipantResultSuccess) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(i.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (i InviteGroupCallParticipantResultSuccess) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(i.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (i InviteGroupCallParticipantResultSuccess) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(i.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (i InviteGroupCallParticipantResultSuccess) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(i.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (i InviteGroupCallParticipantResultSuccess) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(i.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InviteGroupCallParticipantResultSuccess) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(i.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InviteGroupCallParticipantResultSuccess) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(i.ChatId, i.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (i InviteGroupCallParticipantResultSuccess) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(i.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (i InviteGroupCallParticipantResultSuccess) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(i.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (i InviteGroupCallParticipantResultSuccess) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(i.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InviteGroupCallParticipantResultSuccess) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(i.ChatId, i.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (i InviteGroupCallParticipantResultSuccess) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(i.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (i InviteGroupCallParticipantResultSuccess) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(i.ChatId, messageIds, forceRead, opts)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (i Invoice) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(i.Currency, amount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (i Invoice) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, i.Currency, amount)
}

// GetLanguagePackString is a helper method for Client.GetLanguagePackString
func (j JsonObjectMember) GetLanguagePackString(client *Client, languagePackDatabasePath string, localizationTarget string, languagePackId string) (*LanguagePackStringValue, error) {
	return client.GetLanguagePackString(languagePackDatabasePath, localizationTarget, languagePackId, j.Key)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (k KeyboardButton) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, k.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (k KeyboardButton) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, k.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (k KeyboardButton) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(k.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (k KeyboardButton) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(k.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (k KeyboardButton) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(k.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (k KeyboardButton) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, k.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (k KeyboardButton) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, k.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (k KeyboardButton) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, k.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (k KeyboardButton) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(k.Text, inputLanguageCodes)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (k KeyboardButtonTypeWebApp) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, k.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (k KeyboardButtonTypeWebApp) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, k.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (k KeyboardButtonTypeWebApp) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, k.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (k KeyboardButtonTypeWebApp) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(k.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (k KeyboardButtonTypeWebApp) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(k.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (k KeyboardButtonTypeWebApp) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, k.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (k KeyboardButtonTypeWebApp) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(k.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (k KeyboardButtonTypeWebApp) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, k.Url, parameters, opts)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (l LabeledPricePart) CheckAuthenticationPremiumPurchase(client *Client, currency string) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(currency, l.Amount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (l LabeledPricePart) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, currency string) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, currency, l.Amount)
}

// SetSavedMessagesTagLabel is a helper method for Client.SetSavedMessagesTagLabel
func (l LabeledPricePart) SetSavedMessagesTagLabel(client *Client, tag *ReactionType) (*Ok, error) {
	return client.SetSavedMessagesTagLabel(tag, l.Label)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (l LanguagePackInfo) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, l.Name, sticker)
}

// ChangeStickerSet is a helper method for Client.ChangeStickerSet
func (l LanguagePackInfo) ChangeStickerSet(client *Client, setId string, isArchived bool) (*Ok, error) {
	return client.ChangeStickerSet(setId, l.IsInstalled, isArchived)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (l LanguagePackInfo) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(l.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (l LanguagePackInfo) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(l.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (l LanguagePackInfo) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, l.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (l LanguagePackInfo) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, l.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (l LanguagePackInfo) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, l.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (l LanguagePackInfo) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, l.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (l LanguagePackInfo) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, l.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (l LanguagePackInfo) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, l.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (l LanguagePackInfo) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, l.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (l LanguagePackInfo) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(l.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (l LanguagePackInfo) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, l.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (l LanguagePackInfo) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, l.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (l LanguagePackInfo) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, l.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (l LanguagePackInfo) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, l.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (l LanguagePackInfo) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(l.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (l LanguagePackInfo) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(l.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (l LanguagePackInfo) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(l.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (l LanguagePackInfo) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(l.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (l LanguagePackInfo) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, l.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (l LanguagePackInfo) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(l.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (l LanguagePackInfo) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(l.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (l LanguagePackInfo) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, l.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (l LanguagePackInfo) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(l.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (l LanguagePackInfo) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, l.Name)
}

// SetOption is a helper method for Client.SetOption
func (l LanguagePackInfo) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(l.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (l LanguagePackInfo) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, l.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (l LanguagePackInfo) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, l.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (l LanguagePackInfo) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(l.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (l LanguagePackInfo) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, l.Name)
}

// Get is a helper method for Client.GetLanguagePackString
func (l LanguagePackString) Get(client *Client, languagePackDatabasePath string, localizationTarget string, languagePackId string) (*LanguagePackStringValue, error) {
	return client.GetLanguagePackString(languagePackDatabasePath, localizationTarget, languagePackId, l.Key)
}

// SetCustomLanguagePack is a helper method for Client.SetCustomLanguagePack
func (l LanguagePackStrings) SetCustomLanguagePack(client *Client, info *LanguagePackInfo) (*Ok, error) {
	return client.SetCustomLanguagePack(info, l.Strings)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreview) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreview) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (l LinkPreview) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, l.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (l LinkPreview) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, l.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (l LinkPreview) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(l.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (l LinkPreview) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, l.Title, startDate, isRtmpStream)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreview) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreview) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreview) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (l LinkPreview) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(l.Title)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreview) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreview) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreview) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (l LinkPreview) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, l.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (l LinkPreview) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, l.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (l LinkPreview) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, l.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (l LinkPreview) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, l.Title, recordVideo, usePortraitOrientation)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewOptions) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewOptions) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewOptions) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewOptions) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewOptions) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewOptions) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewOptions) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewOptions) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (l LinkPreviewTypeChat) CreateChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, expirationDate, memberLimit, l.CreatesJoinRequest)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (l LinkPreviewTypeChat) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, expirationDate int32, memberLimit int32) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, expirationDate, memberLimit, l.CreatesJoinRequest)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewTypeEmbeddedAnimationPlayer) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewTypeEmbeddedAnimationPlayer) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DiscardCall is a helper method for Client.DiscardCall
func (l LinkPreviewTypeEmbeddedAnimationPlayer) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, l.Duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewTypeEmbeddedAnimationPlayer) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewTypeEmbeddedAnimationPlayer) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (l LinkPreviewTypeEmbeddedAnimationPlayer) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, l.Width, l.Height, scale, chatId)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewTypeEmbeddedAnimationPlayer) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewTypeEmbeddedAnimationPlayer) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewTypeEmbeddedAnimationPlayer) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewTypeEmbeddedAnimationPlayer) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (l LinkPreviewTypeEmbeddedAnimationPlayer) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, l.Duration, paidMessageStarCount)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewTypeEmbeddedAudioPlayer) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewTypeEmbeddedAudioPlayer) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DiscardCall is a helper method for Client.DiscardCall
func (l LinkPreviewTypeEmbeddedAudioPlayer) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, l.Duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewTypeEmbeddedAudioPlayer) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewTypeEmbeddedAudioPlayer) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (l LinkPreviewTypeEmbeddedAudioPlayer) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, l.Width, l.Height, scale, chatId)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewTypeEmbeddedAudioPlayer) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewTypeEmbeddedAudioPlayer) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewTypeEmbeddedAudioPlayer) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewTypeEmbeddedAudioPlayer) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (l LinkPreviewTypeEmbeddedAudioPlayer) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, l.Duration, paidMessageStarCount)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewTypeEmbeddedVideoPlayer) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewTypeEmbeddedVideoPlayer) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DiscardCall is a helper method for Client.DiscardCall
func (l LinkPreviewTypeEmbeddedVideoPlayer) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, l.Duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewTypeEmbeddedVideoPlayer) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewTypeEmbeddedVideoPlayer) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (l LinkPreviewTypeEmbeddedVideoPlayer) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, l.Width, l.Height, scale, chatId)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewTypeEmbeddedVideoPlayer) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewTypeEmbeddedVideoPlayer) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewTypeEmbeddedVideoPlayer) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewTypeEmbeddedVideoPlayer) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (l LinkPreviewTypeEmbeddedVideoPlayer) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, l.Duration, paidMessageStarCount)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewTypeExternalAudio) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewTypeExternalAudio) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DiscardCall is a helper method for Client.DiscardCall
func (l LinkPreviewTypeExternalAudio) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, l.Duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewTypeExternalAudio) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewTypeExternalAudio) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (l LinkPreviewTypeExternalAudio) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(l.MimeType)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewTypeExternalAudio) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewTypeExternalAudio) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewTypeExternalAudio) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewTypeExternalAudio) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (l LinkPreviewTypeExternalAudio) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, l.Duration, paidMessageStarCount)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LinkPreviewTypeExternalVideo) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LinkPreviewTypeExternalVideo) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DiscardCall is a helper method for Client.DiscardCall
func (l LinkPreviewTypeExternalVideo) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, l.Duration, isVideo, connectionId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LinkPreviewTypeExternalVideo) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LinkPreviewTypeExternalVideo) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetFileExtension is a helper method for Client.GetFileExtension
func (l LinkPreviewTypeExternalVideo) GetFileExtension(client *Client) (*Text, error) {
	return client.GetFileExtension(l.MimeType)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (l LinkPreviewTypeExternalVideo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, l.Width, l.Height, scale, chatId)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LinkPreviewTypeExternalVideo) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LinkPreviewTypeExternalVideo) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LinkPreviewTypeExternalVideo) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LinkPreviewTypeExternalVideo) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (l LinkPreviewTypeExternalVideo) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, l.Duration, paidMessageStarCount)
}

// CloseStory is a helper method for Client.CloseStory
func (l LinkPreviewTypeLiveStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(l.StoryPosterChatId, l.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (l LinkPreviewTypeLiveStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(l.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (l LinkPreviewTypeLiveStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, l.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (l LinkPreviewTypeLiveStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(l.StoryPosterChatId, l.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (l LinkPreviewTypeLiveStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(l.StoryPosterChatId, l.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (l LinkPreviewTypeLiveStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(l.StoryPosterChatId, l.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (l LinkPreviewTypeLiveStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(l.StoryPosterChatId, l.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (l LinkPreviewTypeLiveStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(l.StoryPosterChatId, l.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (l LinkPreviewTypeLiveStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(l.StoryPosterChatId, l.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (l LinkPreviewTypeLiveStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(l.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (l LinkPreviewTypeLiveStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(l.StoryPosterChatId, l.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (l LinkPreviewTypeLiveStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, l.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (l LinkPreviewTypeLiveStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(l.StoryPosterChatId, l.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (l LinkPreviewTypeLiveStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(l.StoryPosterChatId, l.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (l LinkPreviewTypeLiveStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(l.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (l LinkPreviewTypeLiveStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(l.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (l LinkPreviewTypeLiveStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(l.StoryPosterChatId, l.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (l LinkPreviewTypeLiveStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(l.StoryPosterChatId, l.StoryId, isPostedToChatPage)
}

// CloseStory is a helper method for Client.CloseStory
func (l LinkPreviewTypeStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(l.StoryPosterChatId, l.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (l LinkPreviewTypeStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(l.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (l LinkPreviewTypeStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, l.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (l LinkPreviewTypeStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(l.StoryPosterChatId, l.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (l LinkPreviewTypeStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(l.StoryPosterChatId, l.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (l LinkPreviewTypeStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(l.StoryPosterChatId, l.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (l LinkPreviewTypeStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(l.StoryPosterChatId, l.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (l LinkPreviewTypeStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(l.StoryPosterChatId, l.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (l LinkPreviewTypeStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(l.StoryPosterChatId, l.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (l LinkPreviewTypeStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(l.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (l LinkPreviewTypeStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(l.StoryPosterChatId, l.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (l LinkPreviewTypeStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, l.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (l LinkPreviewTypeStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(l.StoryPosterChatId, l.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (l LinkPreviewTypeStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(l.StoryPosterChatId, l.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (l LinkPreviewTypeStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(l.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (l LinkPreviewTypeStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(l.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (l LinkPreviewTypeStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(l.StoryPosterChatId, l.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (l LinkPreviewTypeStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(l.StoryPosterChatId, l.StoryId, isPostedToChatPage)
}

// GetCountryFlagEmoji is a helper method for Client.GetCountryFlagEmoji
func (l LocationAddress) GetCountryFlagEmoji(client *Client) (*Text, error) {
	return client.GetCountryFlagEmoji(l.CountryCode)
}

// GetPreferredCountryLanguage is a helper method for Client.GetPreferredCountryLanguage
func (l LocationAddress) GetPreferredCountryLanguage(client *Client) (*Text, error) {
	return client.GetPreferredCountryLanguage(l.CountryCode)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LoginUrlInfoOpen) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LoginUrlInfoOpen) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, l.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LoginUrlInfoOpen) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LoginUrlInfoOpen) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LoginUrlInfoOpen) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LoginUrlInfoOpen) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LoginUrlInfoOpen) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LoginUrlInfoOpen) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, l.Url, parameters, opts)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (l LoginUrlInfoRequestConfirmation) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(l.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (l LoginUrlInfoRequestConfirmation) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(l.BotUserId)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (l LoginUrlInfoRequestConfirmation) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, l.Url, cacheTime)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (l LoginUrlInfoRequestConfirmation) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(l.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (l LoginUrlInfoRequestConfirmation) CheckWebAppFileDownload(client *Client, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(l.BotUserId, fileName, l.Url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (l LoginUrlInfoRequestConfirmation) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, l.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (l LoginUrlInfoRequestConfirmation) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(l.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (l LoginUrlInfoRequestConfirmation) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(l.BotUserId)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (l LoginUrlInfoRequestConfirmation) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, l.Url)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (l LoginUrlInfoRequestConfirmation) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(l.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (l LoginUrlInfoRequestConfirmation) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(l.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (l LoginUrlInfoRequestConfirmation) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(l.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (l LoginUrlInfoRequestConfirmation) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(l.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (l LoginUrlInfoRequestConfirmation) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(l.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (l LoginUrlInfoRequestConfirmation) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(l.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (l LoginUrlInfoRequestConfirmation) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(l.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (l LoginUrlInfoRequestConfirmation) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(l.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (l LoginUrlInfoRequestConfirmation) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(l.BotUserId)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (l LoginUrlInfoRequestConfirmation) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(l.Url)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (l LoginUrlInfoRequestConfirmation) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, l.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (l LoginUrlInfoRequestConfirmation) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(l.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (l LoginUrlInfoRequestConfirmation) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, l.BotUserId, startParameter, parameters)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (l LoginUrlInfoRequestConfirmation) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(l.Url)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (l LoginUrlInfoRequestConfirmation) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(l.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (l LoginUrlInfoRequestConfirmation) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(l.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (l LoginUrlInfoRequestConfirmation) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, l.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (l LoginUrlInfoRequestConfirmation) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(l.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (l LoginUrlInfoRequestConfirmation) GetWebAppUrl(client *Client, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(l.BotUserId, l.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (l LoginUrlInfoRequestConfirmation) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(l.Url, onlyLocal)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (l LoginUrlInfoRequestConfirmation) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(l.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (l LoginUrlInfoRequestConfirmation) OpenWebApp(client *Client, chatId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, l.BotUserId, l.Url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (l LoginUrlInfoRequestConfirmation) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(l.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (l LoginUrlInfoRequestConfirmation) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(l.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (l LoginUrlInfoRequestConfirmation) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(l.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (l LoginUrlInfoRequestConfirmation) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(l.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (l LoginUrlInfoRequestConfirmation) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(l.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (l LoginUrlInfoRequestConfirmation) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(l.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (l LoginUrlInfoRequestConfirmation) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(l.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (l LoginUrlInfoRequestConfirmation) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(l.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (l LoginUrlInfoRequestConfirmation) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(l.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (l LoginUrlInfoRequestConfirmation) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(l.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (l LoginUrlInfoRequestConfirmation) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(l.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (l LoginUrlInfoRequestConfirmation) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(l.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (l LoginUrlInfoRequestConfirmation) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(l.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (l LoginUrlInfoRequestConfirmation) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(l.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (l LoginUrlInfoRequestConfirmation) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(l.BotUserId, username, isActive)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (l LogVerbosityLevel) AddLogMessage(client *Client, text string) (*Ok, error) {
	return client.AddLogMessage(l.VerbosityLevel, text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (m MainWebApp) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, m.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (m MainWebApp) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, m.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (m MainWebApp) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, m.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (m MainWebApp) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(m.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (m MainWebApp) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(m.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (m MainWebApp) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, m.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (m MainWebApp) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(m.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MainWebApp) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, m.Url, parameters, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m Message) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m Message) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m Message) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m Message) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, m.Id, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m Message) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, m.Id, priority)
}

// AddLocal is a helper method for Client.AddLocalMessage
func (m Message) AddLocal(client *Client, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, m.SenderId, disableNotification, inputMessageContent, opts)
}

// AddReaction is a helper method for Client.AddMessageReaction
func (m Message) AddReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, m.Id, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m Message) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, m.Id, options)
}

// AddPendingPaidReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m Message) AddPendingPaidReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, m.Id, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m Message) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m Message) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m Message) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, m.Id, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m Message) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (m Message) BlockSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(m.Id, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (m Message) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (m Message) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m Message) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmoji is a helper method for Client.ClickAnimatedEmojiMessage
func (m Message) ClickAnimatedEmoji(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, m.Id)
}

// ClickChatSponsored is a helper method for Client.ClickChatSponsoredMessage
func (m Message) ClickChatSponsored(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, m.Id, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m Message) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m Message) CommitPendingPaidReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, m.Id)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m Message) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m Message) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m Message) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m Message) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m Message) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, m.Id)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m Message) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, m.Id, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m Message) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m Message) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m Message) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m Message) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatsByDate is a helper method for Client.DeleteChatMessagesByDate
func (m Message) DeleteChatsByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatsBySender is a helper method for Client.DeleteChatMessagesBySender
func (m Message) DeleteChatsBySender(client *Client) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, m.SenderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m Message) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, m.Id)
}

// DeleteDirectsChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m Message) DeleteDirectsChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectsChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m Message) DeleteDirectsChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m Message) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteGroupCallsBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m Message) DeleteGroupCallsBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, m.SenderId, reportSpam)
}

// Deletes is a helper method for Client.DeleteMessages
func (m Message) Deletes(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m Message) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m Message) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessCaption is a helper method for Client.EditBusinessMessageCaption
func (m Message) EditBusinessCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, m.Id, showCaptionAboveMedia, opts)
}

// EditBusinessChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m Message) EditBusinessChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, m.Id, checklist, opts)
}

// EditBusinessLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m Message) EditBusinessLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, m.Id, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMedia is a helper method for Client.EditBusinessMessageMedia
func (m Message) EditBusinessMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, m.Id, inputMessageContent, opts)
}

// EditBusinessReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m Message) EditBusinessReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, m.Id, opts)
}

// EditBusinessText is a helper method for Client.EditBusinessMessageText
func (m Message) EditBusinessText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, m.Id, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m Message) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m Message) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m Message) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditCaption is a helper method for Client.EditMessageCaption
func (m Message) EditCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, m.Id, showCaptionAboveMedia, opts)
}

// EditChecklist is a helper method for Client.EditMessageChecklist
func (m Message) EditChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, m.Id, checklist, opts)
}

// EditLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m Message) EditLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, m.Id, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMedia is a helper method for Client.EditMessageMedia
func (m Message) EditMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, m.Id, inputMessageContent, opts)
}

// EditReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m Message) EditReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, m.Id, opts)
}

// EditSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m Message) EditSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, m.Id, opts)
}

// EditText is a helper method for Client.EditMessageText
func (m Message) EditText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, m.Id, inputMessageContent, opts)
}

// EditQuickReply is a helper method for Client.EditQuickReplyMessage
func (m Message) EditQuickReply(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, m.Id, inputMessageContent)
}

// Forwards is a helper method for Client.ForwardMessages
func (m Message) Forwards(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m Message) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m Message) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, m.Id, payload)
}

// GetCallbackQuery is a helper method for Client.GetCallbackQueryMessage
func (m Message) GetCallbackQuery(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, m.Id, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m Message) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m Message) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m Message) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m Message) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m Message) GetChatAvailableSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m Message) GetChatAvailablePaidReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m Message) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m Message) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m Message) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m Message) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m Message) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m Message) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m Message) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m Message) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m Message) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m Message) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m Message) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m Message) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatByDate is a helper method for Client.GetChatMessageByDate
func (m Message) GetChatByDate(client *Client) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, m.Date)
}

// GetChatCalendar is a helper method for Client.GetChatMessageCalendar
func (m Message) GetChatCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatCount is a helper method for Client.GetChatMessageCount
func (m Message) GetChatCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatPosition is a helper method for Client.GetChatMessagePosition
func (m Message) GetChatPosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, m.Id, opts)
}

// GetChatPinned is a helper method for Client.GetChatPinnedMessage
func (m Message) GetChatPinned(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m Message) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m Message) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m Message) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m Message) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduleds is a helper method for Client.GetChatScheduledMessages
func (m Message) GetChatScheduleds(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m Message) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m Message) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparsePositions is a helper method for Client.GetChatSparseMessagePositions
func (m Message) GetChatSparsePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoreds is a helper method for Client.GetChatSponsoredMessages
func (m Message) GetChatSponsoreds(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m Message) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m Message) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectsChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m Message) GetDirectsChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectsChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m Message) GetDirectsChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectsChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m Message) GetDirectsChatTopicMessageByDate(client *Client, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, m.Date)
}

// GetDirectsChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m Message) GetDirectsChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m Message) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m Message) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m Message) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m Message) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m Message) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, m.Id, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m Message) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, m.Id)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m Message) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m Message) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m Message) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, m.Id, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m Message) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, m.Id, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m Message) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m Message) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// Get is a helper method for Client.GetMessage
func (m Message) Get(client *Client) (*Message, error) {
	return client.GetMessage(m.ChatId, m.Id)
}

// GetAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m Message) GetAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, m.Id, offset, limit, opts)
}

// GetAuthor is a helper method for Client.GetMessageAuthor
func (m Message) GetAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, m.Id)
}

// GetAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m Message) GetAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, m.Id, rowSize)
}

// GetEffect is a helper method for Client.GetMessageEffect
func (m Message) GetEffect(client *Client) (*MessageEffect, error) {
	return client.GetMessageEffect(m.EffectId)
}

// GetEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m Message) GetEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, m.Id, forAlbum)
}

// GetImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m Message) GetImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetLink is a helper method for Client.GetMessageLink
func (m Message) GetLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, m.Id, mediaTimestamp, forAlbum, inMessageThread)
}

// GetLocally is a helper method for Client.GetMessageLocally
func (m Message) GetLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, m.Id)
}

// GetProperties is a helper method for Client.GetMessageProperties
func (m Message) GetProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, m.Id)
}

// GetPublicForwards is a helper method for Client.GetMessagePublicForwards
func (m Message) GetPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, m.Id, offset, limit)
}

// GetReadDate is a helper method for Client.GetMessageReadDate
func (m Message) GetReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, m.Id)
}

// Gets is a helper method for Client.GetMessages
func (m Message) Gets(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetStatistics is a helper method for Client.GetMessageStatistics
func (m Message) GetStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, m.Id, isDark)
}

// GetThread is a helper method for Client.GetMessageThread
func (m Message) GetThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, m.Id)
}

// GetThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m Message) GetThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, m.Id, fromMessageId, offset, limit)
}

// GetViewers is a helper method for Client.GetMessageViewers
func (m Message) GetViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, m.Id)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m Message) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, m.Id)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m Message) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, m.Id, optionId, offset, limit)
}

// GetReplied is a helper method for Client.GetRepliedMessage
func (m Message) GetReplied(client *Client) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, m.Id)
}

// GetSavedsTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (m Message) GetSavedsTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, m.Date)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m Message) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m Message) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m Message) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m Message) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m Message) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m Message) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m Message) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m Message) GetVideoAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, m.Id)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m Message) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// Imports is a helper method for Client.ImportMessages
func (m Message) Imports(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m Message) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m Message) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectsChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m Message) LoadDirectsChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m Message) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, m.Id, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m Message) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m Message) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenContent is a helper method for Client.OpenMessageContent
func (m Message) OpenContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, m.Id)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m Message) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChat is a helper method for Client.PinChatMessage
func (m Message) PinChat(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, m.Id, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m Message) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m Message) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m Message) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (m Message) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(m.Id, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m Message) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, m.Id, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m Message) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m Message) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectsChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m Message) ReadAllDirectsChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m Message) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m Message) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusiness is a helper method for Client.ReadBusinessMessage
func (m Message) ReadBusiness(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, m.Id)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m Message) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, m.Id)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m Message) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m Message) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveReaction is a helper method for Client.RemoveMessageReaction
func (m Message) RemoveReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, m.Id, reactionType)
}

// RemovePendingPaidReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m Message) RemovePendingPaidReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, m.Id)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m Message) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m Message) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m Message) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m Message) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m Message) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m Message) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m Message) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m Message) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m Message) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m Message) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsored is a helper method for Client.ReportChatSponsoredMessage
func (m Message) ReportChatSponsored(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, m.Id, optionId)
}

// ReportReactions is a helper method for Client.ReportMessageReactions
func (m Message) ReportReactions(client *Client) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, m.Id, m.SenderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m Message) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, m.Id)
}

// Resends is a helper method for Client.ResendMessages
func (m Message) Resends(client *Client, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, m.PaidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m Message) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m Message) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m Message) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChats is a helper method for Client.SearchChatMessages
func (m Message) SearchChats(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocations is a helper method for Client.SearchChatRecentLocationMessages
func (m Message) SearchChatRecentLocations(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecrets is a helper method for Client.SearchSecretMessages
func (m Message) SearchSecrets(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStart is a helper method for Client.SendBotStartMessage
func (m Message) SendBotStart(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusiness is a helper method for Client.SendBusinessMessage
func (m Message) SendBusiness(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, m.EffectId, inputMessageContent, opts)
}

// SendBusinessAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m Message) SendBusinessAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, m.EffectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m Message) SendChatAction(client *Client, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, m.TopicId, businessConnectionId, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m Message) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, m.PaidMessageStarCount)
}

// SendGroupCall is a helper method for Client.SendGroupCallMessage
func (m Message) SendGroupCall(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, m.PaidMessageStarCount)
}

// SendInlineQueryResult is a helper method for Client.SendInlineQueryResultMessage
func (m Message) SendInlineQueryResult(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// Send is a helper method for Client.SendMessage
func (m Message) Send(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendAlbum is a helper method for Client.SendMessageAlbum
func (m Message) SendAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcuts is a helper method for Client.SendQuickReplyShortcutMessages
func (m Message) SendQuickReplyShortcuts(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextDraft is a helper method for Client.SendTextMessageDraft
func (m Message) SendTextDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m Message) SetBusinessIsPinned(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, m.Id, m.IsPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m Message) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m Message) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m Message) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m Message) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m Message) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m Message) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m Message) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectsGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m Message) SetChatDirectsGroup(client *Client, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, m.PaidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m Message) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraft is a helper method for Client.SetChatDraftMessage
func (m Message) SetChatDraft(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m Message) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m Message) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m Message) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m Message) SetChatAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatSender is a helper method for Client.SetChatMessageSender
func (m Message) SetChatSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m Message) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m Message) SetChatPaidStarCount(client *Client) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, m.PaidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m Message) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m Message) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m Message) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m Message) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m Message) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m Message) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m Message) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectsChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m Message) SetDirectsChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m Message) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m Message) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, m.Id, editMessage, userId, score, force)
}

// SetGroupCallPaidStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m Message) SetGroupCallPaidStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, m.PaidMessageStarCount)
}

// SetFactCheck is a helper method for Client.SetMessageFactCheck
func (m Message) SetFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, m.Id, opts)
}

// SetReactions is a helper method for Client.SetMessageReactions
func (m Message) SetReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, m.Id, reactionTypes, isBig)
}

// SetSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (m Message) SetSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(m.SenderId, opts)
}

// SetPaidReactionType is a helper method for Client.SetPaidMessageReactionType
func (m Message) SetPaidReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, m.Id, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m Message) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m Message) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m Message) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, m.Id, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m Message) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m Message) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m Message) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, m.Id, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m Message) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, m.Id, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m Message) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, m.PaidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m Message) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, m.Id, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m Message) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, m.Id, opts)
}

// Summarize is a helper method for Client.SummarizeMessage
func (m Message) Summarize(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, m.Id, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m Message) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m Message) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m Message) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m Message) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m Message) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m Message) ToggleChatIsPinned(client *Client, chatList *ChatList) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, m.IsPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m Message) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m Message) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectsChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m Message) ToggleDirectsChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m Message) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m Message) ToggleForumTopicIsPinned(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, m.IsPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m Message) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// ToggleSavedsTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (m Message) ToggleSavedsTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, m.IsPinned)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m Message) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateText is a helper method for Client.TranslateMessageText
func (m Message) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, m.Id, toLanguageCode)
}

// UnpinAllChats is a helper method for Client.UnpinAllChatMessages
func (m Message) UnpinAllChats(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectsChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m Message) UnpinAllDirectsChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopics is a helper method for Client.UnpinAllForumTopicMessages
func (m Message) UnpinAllForumTopics(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChat is a helper method for Client.UnpinChatMessage
func (m Message) UnpinChat(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, m.Id)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m Message) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// Views is a helper method for Client.ViewMessages
func (m Message) Views(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (m MessageAnimatedEmoji) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(m.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (m MessageAnimatedEmoji) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(m.Emoji)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageAnimation) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageAnimation) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (m MessageAnimation) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, m.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageAnimation) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageAudio) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageBasicGroupChatCreate) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, m.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageBasicGroupChatCreate) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, m.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageBasicGroupChatCreate) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(m.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageBasicGroupChatCreate) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, m.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (m MessageBasicGroupChatCreate) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(m.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageBasicGroupChatCreate) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, m.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageBasicGroupChatCreate) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, m.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageBasicGroupChatCreate) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, m.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageBasicGroupChatCreate) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, m.Title, recordVideo, usePortraitOrientation)
}

// CreateCall is a helper method for Client.CreateCall
func (m MessageCall) CreateCall(client *Client, userId int64, protocol *CallProtocol) (*CallId, error) {
	return client.CreateCall(userId, protocol, m.IsVideo)
}

// DiscardCall is a helper method for Client.DiscardCall
func (m MessageCall) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, m.Duration, m.IsVideo, connectionId)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageCall) InviteGroupCallParticipant(client *Client, groupCallId int32, userId int64) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, userId, m.IsVideo)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageCall) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, m.Duration, paidMessageStarCount)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageChatChangeTitle) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, m.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageChatChangeTitle) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, m.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageChatChangeTitle) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(m.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageChatChangeTitle) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, m.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (m MessageChatChangeTitle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(m.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageChatChangeTitle) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, m.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageChatChangeTitle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, m.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageChatChangeTitle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, m.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageChatChangeTitle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, m.Title, recordVideo, usePortraitOrientation)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageChatDeleteMember) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, m.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (m MessageChatDeleteMember) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(m.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageChatDeleteMember) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(m.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (m MessageChatDeleteMember) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(m.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (m MessageChatDeleteMember) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(m.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (m MessageChatDeleteMember) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(m.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (m MessageChatDeleteMember) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(m.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageChatDeleteMember) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(m.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (m MessageChatDeleteMember) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(m.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (m MessageChatDeleteMember) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(m.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageChatDeleteMember) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, m.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (m MessageChatDeleteMember) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(m.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (m MessageChatDeleteMember) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, m.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (m MessageChatDeleteMember) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(m.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (m MessageChatDeleteMember) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(m.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (m MessageChatDeleteMember) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(m.UserId)
}

// GetUser is a helper method for Client.GetUser
func (m MessageChatDeleteMember) GetUser(client *Client) (*User, error) {
	return client.GetUser(m.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageChatDeleteMember) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, m.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (m MessageChatDeleteMember) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(m.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (m MessageChatDeleteMember) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(m.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (m MessageChatDeleteMember) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(m.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (m MessageChatDeleteMember) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(m.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageChatDeleteMember) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(m.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageChatDeleteMember) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, m.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageChatDeleteMember) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, m.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageChatDeleteMember) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, m.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (m MessageChatDeleteMember) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(m.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageChatDeleteMember) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(m.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (m MessageChatDeleteMember) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(m.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageChatDeleteMember) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, m.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (m MessageChatDeleteMember) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, m.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (m MessageChatDeleteMember) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(m.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (m MessageChatDeleteMember) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(m.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageChatDeleteMember) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(m.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (m MessageChatDeleteMember) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(m.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (m MessageChatDeleteMember) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(m.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (m MessageChatDeleteMember) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(m.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (m MessageChatDeleteMember) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(m.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (m MessageChatDeleteMember) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(m.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (m MessageChatDeleteMember) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(m.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (m MessageChatDeleteMember) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(m.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageChatDeleteMember) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, m.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (m MessageChatDeleteMember) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(m.UserId, stickerFormat, sticker)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageChatSetBackground) PinChatMessage(client *Client, chatId int64, messageId int64, disableNotification bool) (*Ok, error) {
	return client.PinChatMessage(chatId, messageId, disableNotification, m.OnlyForSelf)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageChatSetBackground) SetChatBackground(client *Client, chatId int64, darkThemeDimming int32, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(chatId, darkThemeDimming, m.OnlyForSelf, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageChatSetMessageAutoDeleteTime) CreateNewBasicGroupChat(client *Client, userIds []int64, title string) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, title, m.MessageAutoDeleteTime)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageChatSetMessageAutoDeleteTime) CreateNewSupergroupChat(client *Client, title string, isForum bool, isChannel bool, description string, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, isForum, isChannel, description, m.MessageAutoDeleteTime, forImport, opts)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageChatSetMessageAutoDeleteTime) SetChatMessageAutoDeleteTime(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(chatId, m.MessageAutoDeleteTime)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageChatShared) ShareChatWithBot(client *Client, chatId int64, messageId int64, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, messageId, m.ButtonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageChatShared) ShareUsersWithBot(client *Client, chatId int64, messageId int64, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, messageId, m.ButtonId, sharedUserIds, onlyCheck)
}

// CreateBasicGroupChat is a helper method for Client.CreateBasicGroupChat
func (m MessageChatUpgradeFrom) CreateBasicGroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateBasicGroupChat(m.BasicGroupId, force)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageChatUpgradeFrom) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, m.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageChatUpgradeFrom) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, m.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageChatUpgradeFrom) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(m.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageChatUpgradeFrom) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, m.Title, startDate, isRtmpStream)
}

// GetBasicGroup is a helper method for Client.GetBasicGroup
func (m MessageChatUpgradeFrom) GetBasicGroup(client *Client) (*BasicGroup, error) {
	return client.GetBasicGroup(m.BasicGroupId)
}

// GetBasicGroupFullInfo is a helper method for Client.GetBasicGroupFullInfo
func (m MessageChatUpgradeFrom) GetBasicGroupFullInfo(client *Client) (*BasicGroupFullInfo, error) {
	return client.GetBasicGroupFullInfo(m.BasicGroupId)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (m MessageChatUpgradeFrom) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(m.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageChatUpgradeFrom) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, m.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageChatUpgradeFrom) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, m.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageChatUpgradeFrom) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, m.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageChatUpgradeFrom) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, m.Title, recordVideo, usePortraitOrientation)
}

// CreateSupergroupChat is a helper method for Client.CreateSupergroupChat
func (m MessageChatUpgradeTo) CreateSupergroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(m.SupergroupId, force)
}

// DisableAllSupergroupUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (m MessageChatUpgradeTo) DisableAllSupergroupUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(m.SupergroupId)
}

// GetSupergroup is a helper method for Client.GetSupergroup
func (m MessageChatUpgradeTo) GetSupergroup(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(m.SupergroupId)
}

// GetSupergroupFullInfo is a helper method for Client.GetSupergroupFullInfo
func (m MessageChatUpgradeTo) GetSupergroupFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(m.SupergroupId)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (m MessageChatUpgradeTo) GetSupergroupMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(m.SupergroupId, offset, limit, opts)
}

// ReorderSupergroupActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (m MessageChatUpgradeTo) ReorderSupergroupActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(m.SupergroupId, usernames)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m MessageChatUpgradeTo) ReportSupergroupAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(m.SupergroupId, messageId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (m MessageChatUpgradeTo) ReportSupergroupSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(m.SupergroupId, messageIds)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (m MessageChatUpgradeTo) SetSupergroupCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(m.SupergroupId, customEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (m MessageChatUpgradeTo) SetSupergroupMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(m.SupergroupId, mainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (m MessageChatUpgradeTo) SetSupergroupStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(m.SupergroupId, stickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (m MessageChatUpgradeTo) SetSupergroupUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(m.SupergroupId, unrestrictBoostCount)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (m MessageChatUpgradeTo) SetSupergroupUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(m.SupergroupId, username)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (m MessageChatUpgradeTo) ToggleSupergroupCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(m.SupergroupId, canHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (m MessageChatUpgradeTo) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(m.SupergroupId, hasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (m MessageChatUpgradeTo) ToggleSupergroupHasAutomaticTranslation(client *Client, hasAutomaticTranslation bool) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(m.SupergroupId, hasAutomaticTranslation)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (m MessageChatUpgradeTo) ToggleSupergroupHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(m.SupergroupId, hasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (m MessageChatUpgradeTo) ToggleSupergroupIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(m.SupergroupId, isAllHistoryAvailable)
}

// ToggleSupergroupIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (m MessageChatUpgradeTo) ToggleSupergroupIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(m.SupergroupId)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (m MessageChatUpgradeTo) ToggleSupergroupIsForum(client *Client, isForum bool, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(m.SupergroupId, isForum, hasForumTabs)
}

// ToggleSupergroupJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (m MessageChatUpgradeTo) ToggleSupergroupJoinByRequest(client *Client, joinByRequest bool) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(m.SupergroupId, joinByRequest)
}

// ToggleSupergroupJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (m MessageChatUpgradeTo) ToggleSupergroupJoinToSendMessages(client *Client, joinToSendMessages bool) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(m.SupergroupId, joinToSendMessages)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (m MessageChatUpgradeTo) ToggleSupergroupSignMessages(client *Client, signMessages bool, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(m.SupergroupId, signMessages, showMessageSender)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (m MessageChatUpgradeTo) ToggleSupergroupUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(m.SupergroupId, username, isActive)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageChecklistTasksDone) MarkChecklistTasksAsDone(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, messageId, m.MarkedAsDoneTaskIds, m.MarkedAsNotDoneTaskIds)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageCopyOptions) ForwardMessages(client *Client, chatId int64, fromChatId int64, messageIds []int64, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(chatId, fromChatId, messageIds, m.SendCopy, removeCaption, opts)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (m MessageCustomServiceAction) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, m.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (m MessageCustomServiceAction) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, m.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (m MessageCustomServiceAction) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(m.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (m MessageCustomServiceAction) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(m.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (m MessageCustomServiceAction) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(m.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageCustomServiceAction) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, m.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageCustomServiceAction) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, m.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (m MessageCustomServiceAction) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, m.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (m MessageCustomServiceAction) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(m.Text, inputLanguageCodes)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (m MessageDice) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(m.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (m MessageDice) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(m.Emoji)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageDirectMessagePriceChanged) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, m.PaidMessageStarCount, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageDirectMessagePriceChanged) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, m.PaidMessageStarCount)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageDirectMessagePriceChanged) SendGroupCallMessage(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, m.PaidMessageStarCount)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageDirectMessagePriceChanged) SetChatDirectMessagesGroup(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, m.IsEnabled, m.PaidMessageStarCount)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageDirectMessagePriceChanged) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, m.PaidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessageDirectMessagePriceChanged) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, m.PaidMessageStarCount)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageDirectMessagePriceChanged) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, isRtmpStream, enableMessages, m.PaidMessageStarCount)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageDocument) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (m MessageEffect) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(m.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (m MessageEffect) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(m.Emoji)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageFileTypeGroup) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, m.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageFileTypeGroup) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, m.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageFileTypeGroup) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(m.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageFileTypeGroup) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, m.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (m MessageFileTypeGroup) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(m.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageFileTypeGroup) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, m.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageFileTypeGroup) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, m.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageFileTypeGroup) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, m.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageFileTypeGroup) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, m.Title, recordVideo, usePortraitOrientation)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageFileTypePrivate) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, m.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (m MessageFileTypePrivate) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(m.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (m MessageFileTypePrivate) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(m.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (m MessageFileTypePrivate) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, m.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageFileTypePrivate) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageFileTypePrivate) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, m.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageFileTypePrivate) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, m.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (m MessageFileTypePrivate) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, m.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageFileTypePrivate) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, m.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (m MessageFileTypePrivate) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, m.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (m MessageFileTypePrivate) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(m.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (m MessageFileTypePrivate) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, m.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageFileTypePrivate) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageFileTypePrivate) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, m.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageFileTypePrivate) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, m.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (m MessageFileTypePrivate) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(m.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (m MessageFileTypePrivate) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(m.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (m MessageFileTypePrivate) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(m.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (m MessageFileTypePrivate) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(m.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageFileTypePrivate) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, m.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (m MessageFileTypePrivate) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(m.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (m MessageFileTypePrivate) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(m.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (m MessageFileTypePrivate) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, m.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (m MessageFileTypePrivate) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(m.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (m MessageFileTypePrivate) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, m.Name)
}

// SetOption is a helper method for Client.SetOption
func (m MessageFileTypePrivate) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(m.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (m MessageFileTypePrivate) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, m.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageFileTypePrivate) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, m.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageFileTypePrivate) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(m.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageFileTypePrivate) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, m.Name)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageForumTopicCreated) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, m.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (m MessageForumTopicCreated) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(m.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (m MessageForumTopicCreated) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(m.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (m MessageForumTopicCreated) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, m.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageForumTopicCreated) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageForumTopicCreated) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, m.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageForumTopicCreated) CreateForumTopic(client *Client, chatId int64) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, m.Name, m.IsNameImplicit, m.Icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (m MessageForumTopicCreated) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, m.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageForumTopicCreated) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, m.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (m MessageForumTopicCreated) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, m.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (m MessageForumTopicCreated) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(m.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (m MessageForumTopicCreated) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, m.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageForumTopicCreated) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageForumTopicCreated) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, m.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageForumTopicCreated) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, m.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (m MessageForumTopicCreated) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(m.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (m MessageForumTopicCreated) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(m.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (m MessageForumTopicCreated) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(m.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (m MessageForumTopicCreated) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(m.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageForumTopicCreated) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, m.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (m MessageForumTopicCreated) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(m.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (m MessageForumTopicCreated) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(m.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (m MessageForumTopicCreated) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, m.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (m MessageForumTopicCreated) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(m.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (m MessageForumTopicCreated) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, m.Name)
}

// SetOption is a helper method for Client.SetOption
func (m MessageForumTopicCreated) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(m.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (m MessageForumTopicCreated) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, m.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageForumTopicCreated) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, m.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageForumTopicCreated) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(m.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageForumTopicCreated) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, m.Name)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageForumTopicEdited) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, m.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (m MessageForumTopicEdited) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(m.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (m MessageForumTopicEdited) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(m.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (m MessageForumTopicEdited) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, m.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageForumTopicEdited) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageForumTopicEdited) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, m.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageForumTopicEdited) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, m.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (m MessageForumTopicEdited) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, m.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageForumTopicEdited) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, m.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (m MessageForumTopicEdited) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, m.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (m MessageForumTopicEdited) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(m.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (m MessageForumTopicEdited) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, m.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageForumTopicEdited) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, m.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageForumTopicEdited) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, m.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageForumTopicEdited) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, m.Name, editIconCustomEmoji, m.IconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (m MessageForumTopicEdited) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(m.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (m MessageForumTopicEdited) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(m.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (m MessageForumTopicEdited) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(m.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (m MessageForumTopicEdited) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(m.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageForumTopicEdited) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, m.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (m MessageForumTopicEdited) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(m.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (m MessageForumTopicEdited) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(m.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (m MessageForumTopicEdited) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, m.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (m MessageForumTopicEdited) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(m.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (m MessageForumTopicEdited) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, m.Name)
}

// SetOption is a helper method for Client.SetOption
func (m MessageForumTopicEdited) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(m.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (m MessageForumTopicEdited) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, m.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageForumTopicEdited) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, m.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageForumTopicEdited) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(m.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageForumTopicEdited) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, m.Name)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageForumTopicIsClosedToggled) ToggleForumTopicIsClosed(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(chatId, forumTopicId, m.IsClosed)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageForumTopicIsHiddenToggled) ToggleGeneralForumTopicIsHidden(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(chatId, m.IsHidden)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageForwardInfo) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, m.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageForwardInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, m.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (m MessageForwardInfo) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, m.Date)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageGameScore) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, userId int64, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, userId, m.Score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (m MessageGameScore) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, userId int64, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, userId, m.Score, force)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageGift) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, m.SenderId, disableNotification, inputMessageContent, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessageGift) BuyGiftUpgrade(client *Client, ownerId *MessageSender, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, m.PrepaidUpgradeHash, starCount)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageGift) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, m.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageGift) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, m.SenderId, reportSpam)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessageGift) DropGiftOriginalDetails(client *Client, starCount int64) (*Ok, error) {
	return client.DropGiftOriginalDetails(m.ReceivedGiftId, starCount)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (m MessageGift) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(m.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (m MessageGift) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(m.Text)
}

// GetReceivedGift is a helper method for Client.GetReceivedGift
func (m MessageGift) GetReceivedGift(client *Client) (*ReceivedGift, error) {
	return client.GetReceivedGift(m.ReceivedGiftId)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (m MessageGift) GetUpgradedGiftWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(m.ReceivedGiftId, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageGift) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, m.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (m MessageGift) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(m.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageGift) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, m.Text, m.IsPrivate)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageGift) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, m.SenderId)
}

// SearchQuote is a helper method for Client.SearchQuote
func (m MessageGift) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(m.Text, quote, quotePosition)
}

// SellGift is a helper method for Client.SellGift
func (m MessageGift) SellGift(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SellGift(businessConnectionId, m.ReceivedGiftId)
}

// SendGift is a helper method for Client.SendGift
func (m MessageGift) SendGift(client *Client, giftId string, ownerId *MessageSender, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, m.Text, m.IsPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageGift) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, m.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageGift) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, m.Text)
}

// SetGiftResalePrice is a helper method for Client.SetGiftResalePrice
func (m MessageGift) SetGiftResalePrice(client *Client, opts *SetGiftResalePriceOpts) (*Ok, error) {
	return client.SetGiftResalePrice(m.ReceivedGiftId, opts)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (m MessageGift) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(m.SenderId, opts)
}

// ToggleGiftIsSaved is a helper method for Client.ToggleGiftIsSaved
func (m MessageGift) ToggleGiftIsSaved(client *Client) (*Ok, error) {
	return client.ToggleGiftIsSaved(m.ReceivedGiftId, m.IsSaved)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessageGift) TransferGift(client *Client, businessConnectionId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	return client.TransferGift(businessConnectionId, m.ReceivedGiftId, newOwnerId, starCount)
}

// TranslateText is a helper method for Client.TranslateText
func (m MessageGift) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(m.Text, toLanguageCode)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessageGift) UpgradeGift(client *Client, businessConnectionId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, m.ReceivedGiftId, keepOriginalDetails, starCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessageGiftedPremium) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, m.Amount)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (m MessageGiftedPremium) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(m.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (m MessageGiftedPremium) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(m.Text)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (m MessageGiftedPremium) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(m.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageGiftedPremium) GiftPremiumWithStars(client *Client, userId int64, starCount int64) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, m.MonthCount, m.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (m MessageGiftedPremium) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(m.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageGiftedPremium) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, m.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (m MessageGiftedPremium) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(m.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (m MessageGiftedPremium) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, m.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageGiftedPremium) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, m.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageGiftedPremium) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, m.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessageGiftedPremium) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, m.Amount)
}

// TranslateText is a helper method for Client.TranslateText
func (m MessageGiftedPremium) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(m.Text, toLanguageCode)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageGiftedStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, m.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageGiftedStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, m.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessageGiftedStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, m.StarCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessageGiftedStars) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, m.Amount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessageGiftedStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, m.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessageGiftedStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, m.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageGiftedStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, m.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (m MessageGiftedStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiftedStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, m.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageGiftedStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, m.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (m MessageGiftedStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, m.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessageGiftedStars) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, m.Amount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (m MessageGiftedStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, m.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessageGiftedStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, m.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessageGiftedStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, m.Parameters, m.WinnerCount, starCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiveawayCompleted) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, m.WinnerCount, starCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageGiveawayCreated) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, m.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageGiveawayCreated) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, m.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessageGiveawayCreated) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, m.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessageGiveawayCreated) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, m.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessageGiveawayCreated) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, m.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageGiveawayCreated) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, m.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (m MessageGiveawayCreated) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiveawayCreated) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, m.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageGiveawayCreated) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, m.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (m MessageGiveawayCreated) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, m.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (m MessageGiveawayCreated) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, m.StarCount)
}
