package gotdbot

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypePaidMessageSend) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypePaidMessageSend) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypePaidMessageSend) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypePaidMessageSend) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypePaidMessageSend) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypePaidMessageSend) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypePaidMessageSend) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypePaidMessageSend) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypePaidMessageSend) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypePaidMessageSend) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypePaidMessageSend) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypePaidMessageSend) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypePaidMessageSend) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypePaidMessageSend) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypePaidMessageSend) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypePaidMessageSend) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypePaidMessageSend) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypePaidMessageSend) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypePaidMessageSend) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypePaidMessageSend) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypePaidMessageSend) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypePaidMessageSend) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypePaidMessageSend) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypePaidMessageSend) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypePaidMessageSend) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypePaidMessageSend) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypePaidMessageSend) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypePaidMessageSend) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypePaidMessageSend) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypePaidMessageSend) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypePaidMessageSend) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypePaidMessageSend) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypePaidMessageSend) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypePaidMessageSend) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidMessageSend) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypePaidMessageSend) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypePaidMessageSend) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypePaidMessageSend) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypePaidMessageSend) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypePaidMessageSend) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypePaidMessageSend) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypePaidMessageSend) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypePaidMessageSend) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypePaidMessageSend) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypePaidMessageSend) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypePaidMessageSend) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypePaidMessageSend) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypePaidMessageSend) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypePaidMessageSend) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypePaidMessageSend) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypePaidMessageSend) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypePaidMessageSend) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypePaidMessageSend) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypePaidMessageSend) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypePaidMessageSend) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypePaidMessageSend) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypePaidMessageSend) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypePaidMessageSend) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypePaidMessageSend) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypePaidMessageSend) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypePaidMessageSend) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypePaidMessageSend) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypePaidMessageSend) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypePaidMessageSend) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypePaidMessageSend) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypePaidMessageSend) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypePaidMessageSend) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypePaidMessageSend) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypePaidMessageSend) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypePaidMessageSend) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypePaidMessageSend) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypePaidMessageSend) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypePaidMessageSend) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypePaidMessageSend) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypePaidMessageSend) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypePaidMessageSend) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypePaidMessageSend) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypePaidMessageSend) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypePaidMessageSend) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypePaidMessageSend) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypePaidMessageSend) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypePaidMessageSend) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypePaidMessageSend) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypePaidMessageSend) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypePaidMessageSend) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypePaidMessageSend) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypePaidMessageSend) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypePaidMessageSend) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypePaidMessageSend) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypePaidMessageSend) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypePaidMessageSend) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypePaidMessageSend) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypePaidMessageSend) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypePaidMessageSend) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypePaidMessageSend) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypePaidMessageSend) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypePaidMessageSend) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypePaidMessageSend) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypePaidMessageSend) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypePaidMessageSend) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypePaidMessageSend) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypePaidMessageSend) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypePaidMessageSend) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypePaidMessageSend) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypePaidMessageSend) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypePaidMessageSend) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypePaidMessageSend) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypePaidMessageSend) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypePaidMessageSend) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypePaidMessageSend) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypePaidMessageSend) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypePaidMessageSend) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypePaidMessageSend) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypePremiumPurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypePremiumPurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypePremiumPurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypePremiumPurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypePremiumPurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypePremiumPurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypePremiumPurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypePremiumPurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypePremiumPurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypePremiumPurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypePremiumPurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypePremiumPurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypePremiumPurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypePremiumPurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypePremiumPurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (s StarTransactionTypePremiumPurchase) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(s.MonthCount)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypePremiumPurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypePremiumPurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypePremiumPurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypePremiumPurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypePremiumPurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypePremiumPurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypePremiumPurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypePremiumPurchase) GiftPremiumWithStars(client *Client, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, s.MonthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypePremiumPurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypePremiumPurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypePremiumPurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypePremiumPurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypePremiumPurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypePremiumPurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypePremiumPurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypePremiumPurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypePremiumPurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypePremiumPurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypePremiumPurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypePremiumPurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypePremiumPurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypePremiumPurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypePremiumPurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypePremiumPurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypePremiumPurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypePremiumPurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypePremiumPurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypePremiumPurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeSuggestedPostPaymentReceive) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeSuggestedPostPaymentReceive) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeSuggestedPostPaymentReceive) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeSuggestedPostPaymentReceive) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeSuggestedPostPaymentReceive) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeSuggestedPostPaymentReceive) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeSuggestedPostPaymentReceive) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeSuggestedPostPaymentReceive) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeSuggestedPostPaymentReceive) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeSuggestedPostPaymentReceive) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeSuggestedPostPaymentReceive) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeSuggestedPostPaymentReceive) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeSuggestedPostPaymentReceive) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeSuggestedPostPaymentReceive) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeSuggestedPostPaymentReceive) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeSuggestedPostPaymentReceive) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeSuggestedPostPaymentReceive) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeSuggestedPostPaymentReceive) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeSuggestedPostPaymentReceive) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeSuggestedPostPaymentReceive) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeSuggestedPostPaymentReceive) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeSuggestedPostPaymentReceive) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeSuggestedPostPaymentReceive) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeSuggestedPostPaymentReceive) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeSuggestedPostPaymentSend) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeSuggestedPostPaymentSend) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeSuggestedPostPaymentSend) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeSuggestedPostPaymentSend) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeSuggestedPostPaymentSend) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeSuggestedPostPaymentSend) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeSuggestedPostPaymentSend) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeSuggestedPostPaymentSend) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeSuggestedPostPaymentSend) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeSuggestedPostPaymentSend) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeSuggestedPostPaymentSend) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeSuggestedPostPaymentSend) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeSuggestedPostPaymentSend) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeSuggestedPostPaymentSend) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeSuggestedPostPaymentSend) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeSuggestedPostPaymentSend) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeSuggestedPostPaymentSend) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeSuggestedPostPaymentSend) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeSuggestedPostPaymentSend) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeSuggestedPostPaymentSend) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeSuggestedPostPaymentSend) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeSuggestedPostPaymentSend) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeSuggestedPostPaymentSend) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeSuggestedPostPaymentSend) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeSuggestedPostPaymentSend) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeSuggestedPostPaymentSend) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeSuggestedPostPaymentSend) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeSuggestedPostPaymentSend) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeSuggestedPostPaymentSend) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeSuggestedPostPaymentSend) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeSuggestedPostPaymentSend) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeSuggestedPostPaymentSend) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeSuggestedPostPaymentSend) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeSuggestedPostPaymentSend) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeSuggestedPostPaymentSend) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeSuggestedPostPaymentSend) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeSuggestedPostPaymentSend) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeSuggestedPostPaymentSend) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeSuggestedPostPaymentSend) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeSuggestedPostPaymentSend) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeSuggestedPostPaymentSend) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeSuggestedPostPaymentSend) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeSuggestedPostPaymentSend) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeSuggestedPostPaymentSend) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeSuggestedPostPaymentSend) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeSuggestedPostPaymentSend) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeSuggestedPostPaymentSend) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeSuggestedPostPaymentSend) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeSuggestedPostPaymentSend) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeSuggestedPostPaymentSend) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeSuggestedPostPaymentSend) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeSuggestedPostPaymentSend) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeSuggestedPostPaymentSend) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeSuggestedPostPaymentSend) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeSuggestedPostPaymentSend) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeSuggestedPostPaymentSend) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeSuggestedPostPaymentSend) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeSuggestedPostPaymentSend) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeSuggestedPostPaymentSend) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeSuggestedPostPaymentSend) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeSuggestedPostPaymentSend) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeSuggestedPostPaymentSend) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeSuggestedPostPaymentSend) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeSuggestedPostPaymentSend) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeSuggestedPostPaymentSend) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeSuggestedPostPaymentSend) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeSuggestedPostPaymentSend) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeSuggestedPostPaymentSend) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeSuggestedPostPaymentSend) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeSuggestedPostPaymentSend) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeSuggestedPostPaymentSend) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeSuggestedPostPaymentSend) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeSuggestedPostPaymentSend) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeSuggestedPostPaymentSend) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeSuggestedPostPaymentSend) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeSuggestedPostPaymentSend) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeSuggestedPostPaymentSend) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeSuggestedPostPaymentSend) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeSuggestedPostPaymentSend) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeSuggestedPostPaymentSend) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeSuggestedPostPaymentSend) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeSuggestedPostPaymentSend) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeSuggestedPostPaymentSend) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeSuggestedPostPaymentSend) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeSuggestedPostPaymentSend) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeSuggestedPostPaymentSend) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeSuggestedPostPaymentSend) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeSuggestedPostPaymentSend) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeSuggestedPostPaymentSend) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeSuggestedPostPaymentSend) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeSuggestedPostPaymentSend) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeSuggestedPostPaymentSend) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeSuggestedPostPaymentSend) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeSuggestedPostPaymentSend) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeSuggestedPostPaymentSend) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeSuggestedPostPaymentSend) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeSuggestedPostPaymentSend) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeUpgradedGiftPurchase) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeUpgradedGiftPurchase) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeUpgradedGiftPurchase) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeUpgradedGiftPurchase) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeUpgradedGiftPurchase) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeUpgradedGiftPurchase) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeUpgradedGiftPurchase) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeUpgradedGiftPurchase) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeUpgradedGiftPurchase) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeUpgradedGiftPurchase) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeUpgradedGiftPurchase) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeUpgradedGiftPurchase) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeUpgradedGiftPurchase) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeUpgradedGiftPurchase) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeUpgradedGiftPurchase) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeUpgradedGiftPurchase) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeUpgradedGiftPurchase) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeUpgradedGiftPurchase) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeUpgradedGiftPurchase) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeUpgradedGiftPurchase) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeUpgradedGiftPurchase) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeUpgradedGiftPurchase) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeUpgradedGiftPurchase) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeUpgradedGiftPurchase) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeUpgradedGiftPurchase) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeUpgradedGiftPurchase) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeUpgradedGiftPurchase) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeUpgradedGiftPurchase) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeUpgradedGiftPurchase) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeUpgradedGiftPurchase) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeUpgradedGiftPurchase) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeUpgradedGiftPurchase) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeUpgradedGiftPurchase) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeUpgradedGiftPurchase) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeUpgradedGiftPurchase) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeUpgradedGiftPurchase) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeUpgradedGiftPurchase) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeUpgradedGiftPurchase) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeUpgradedGiftPurchase) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeUpgradedGiftPurchase) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeUpgradedGiftPurchase) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeUpgradedGiftPurchase) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeUpgradedGiftPurchase) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeUpgradedGiftSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeUpgradedGiftSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeUpgradedGiftSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeUpgradedGiftSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeUpgradedGiftSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeUpgradedGiftSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeUpgradedGiftSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeUpgradedGiftSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeUpgradedGiftSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeUpgradedGiftSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeUpgradedGiftSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeUpgradedGiftSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeUpgradedGiftSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeUpgradedGiftSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeUpgradedGiftSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeUpgradedGiftSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeUpgradedGiftSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeUpgradedGiftSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeUpgradedGiftSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeUpgradedGiftSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeUpgradedGiftSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeUpgradedGiftSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeUpgradedGiftSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeUpgradedGiftSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeUpgradedGiftSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeUpgradedGiftSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeUpgradedGiftSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeUpgradedGiftSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeUpgradedGiftSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeUpgradedGiftSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeUpgradedGiftSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeUpgradedGiftSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeUpgradedGiftSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeUpgradedGiftSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeUpgradedGiftSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeUpgradedGiftSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeUpgradedGiftSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeUpgradedGiftSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeUpgradedGiftSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeUpgradedGiftSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeUpgradedGiftSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeUpgradedGiftSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeUpgradedGiftSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeUserDeposit) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeUserDeposit) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeUserDeposit) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeUserDeposit) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeUserDeposit) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeUserDeposit) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeUserDeposit) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeUserDeposit) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeUserDeposit) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeUserDeposit) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeUserDeposit) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeUserDeposit) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeUserDeposit) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeUserDeposit) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeUserDeposit) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeUserDeposit) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeUserDeposit) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeUserDeposit) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeUserDeposit) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeUserDeposit) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeUserDeposit) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeUserDeposit) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeUserDeposit) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeUserDeposit) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeUserDeposit) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeUserDeposit) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeUserDeposit) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeUserDeposit) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeUserDeposit) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeUserDeposit) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeUserDeposit) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeUserDeposit) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeUserDeposit) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeUserDeposit) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeUserDeposit) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeUserDeposit) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeUserDeposit) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeUserDeposit) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeUserDeposit) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeUserDeposit) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeUserDeposit) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeUserDeposit) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeUserDeposit) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// CheckAuthenticationBotToken is a helper method for Client.CheckAuthenticationBotToken
func (s StatisticalGraphAsync) CheckAuthenticationBotToken(client *Client) (*Ok, error) {
	return client.CheckAuthenticationBotToken(s.Token)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StatisticalGraphAsync) GetStatisticalGraph(client *Client, chatId int64, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(chatId, s.Token, x)
}

// SearchUserByToken is a helper method for Client.SearchUserByToken
func (s StatisticalGraphAsync) SearchUserByToken(client *Client) (*User, error) {
	return client.SearchUserByToken(s.Token)
}

// SendAuthenticationFirebaseSms is a helper method for Client.SendAuthenticationFirebaseSms
func (s StatisticalGraphAsync) SendAuthenticationFirebaseSms(client *Client) (*Ok, error) {
	return client.SendAuthenticationFirebaseSms(s.Token)
}

// SendPhoneNumberFirebaseSms is a helper method for Client.SendPhoneNumberFirebaseSms
func (s StatisticalGraphAsync) SendPhoneNumberFirebaseSms(client *Client) (*Ok, error) {
	return client.SendPhoneNumberFirebaseSms(s.Token)
}

// SetApplicationVerificationToken is a helper method for Client.SetApplicationVerificationToken
func (s StatisticalGraphAsync) SetApplicationVerificationToken(client *Client, verificationId int64) (*Ok, error) {
	return client.SetApplicationVerificationToken(verificationId, s.Token)
}

// AnswerPreCheckoutQuery is a helper method for Client.AnswerPreCheckoutQuery
func (s StatisticalGraphError) AnswerPreCheckoutQuery(client *Client, preCheckoutQueryId string) (*Ok, error) {
	return client.AnswerPreCheckoutQuery(preCheckoutQueryId, s.ErrorMessage)
}

// AnswerShippingQuery is a helper method for Client.AnswerShippingQuery
func (s StatisticalGraphError) AnswerShippingQuery(client *Client, shippingQueryId string, shippingOptions []*ShippingOption) (*Ok, error) {
	return client.AnswerShippingQuery(shippingQueryId, shippingOptions, s.ErrorMessage)
}

// SetBotUpdatesStatus is a helper method for Client.SetBotUpdatesStatus
func (s StatisticalGraphError) SetBotUpdatesStatus(client *Client, pendingUpdateCount int32) (*Ok, error) {
	return client.SetBotUpdatesStatus(pendingUpdateCount, s.ErrorMessage)
}

// ChangeSet is a helper method for Client.ChangeStickerSet
func (s Sticker) ChangeSet(client *Client, isInstalled bool, isArchived bool) (*Ok, error) {
	return client.ChangeStickerSet(s.SetId, isInstalled, isArchived)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (s Sticker) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(s.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (s Sticker) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(s.Emoji)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s Sticker) GetMapThumbnailFile(client *Client, location *Location, zoom int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, s.Width, s.Height, scale, chatId)
}

// GetSet is a helper method for Client.GetStickerSet
func (s Sticker) GetSet(client *Client) (*StickerSet, error) {
	return client.GetStickerSet(s.SetId)
}

// GetSetName is a helper method for Client.GetStickerSetName
func (s Sticker) GetSetName(client *Client) (*Text, error) {
	return client.GetStickerSetName(s.SetId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StickerFullTypeCustomEmoji) CreateNewStickerSet(client *Client, userId int64, title string, name string, stickerType *StickerType, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, name, stickerType, s.NeedsRepainting, stickers, source)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (s StickerFullTypeCustomEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, name string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(name, s.CustomEmojiId)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StickerSet) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, s.Name, sticker)
}

// Change is a helper method for Client.ChangeStickerSet
func (s StickerSet) Change(client *Client, setId string) (*Ok, error) {
	return client.ChangeStickerSet(setId, s.IsInstalled, s.IsArchived)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (s StickerSet) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(s.Name)
}

// CheckName is a helper method for Client.CheckStickerSetName
func (s StickerSet) CheckName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(s.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (s StickerSet) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, s.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StickerSet) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StickerSet) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, s.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StickerSet) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, s.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StickerSet) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, s.Name, receivedGiftIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (s StickerSet) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, s.Title, messageAutoDeleteTime)
}

// CreateNew is a helper method for Client.CreateNewStickerSet
func (s StickerSet) CreateNew(client *Client, userId int64, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, s.Title, s.Name, s.StickerType, s.NeedsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (s StickerSet) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(s.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (s StickerSet) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, s.Name, storyIds)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StickerSet) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, s.Title, startDate, isRtmpStream)
}

// Delete is a helper method for Client.DeleteStickerSet
func (s StickerSet) Delete(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(s.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (s StickerSet) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, s.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StickerSet) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StickerSet) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, s.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StickerSet) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, s.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StickerSet) GetAllStickerEmojis(client *Client, query string, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(s.StickerType, query, chatId, returnOnlyMainEmoji)
}

// GetArchiveds is a helper method for Client.GetArchivedStickerSets
func (s StickerSet) GetArchiveds(client *Client, offsetStickerSetId string, limit int32) (*StickerSets, error) {
	return client.GetArchivedStickerSets(s.StickerType, offsetStickerSetId, limit)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (s StickerSet) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(s.Name, typeField)
}

// GetInstalleds is a helper method for Client.GetInstalledStickerSets
func (s StickerSet) GetInstalleds(client *Client) (*StickerSets, error) {
	return client.GetInstalledStickerSets(s.StickerType)
}

// GetOption is a helper method for Client.GetOption
func (s StickerSet) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(s.Name)
}

// GetStickers is a helper method for Client.GetStickers
func (s StickerSet) GetStickers(client *Client, query string, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(s.StickerType, query, limit, chatId)
}

// GetSuggestedName is a helper method for Client.GetSuggestedStickerSetName
func (s StickerSet) GetSuggestedName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(s.Title)
}

// GetTrendings is a helper method for Client.GetTrendingStickerSets
func (s StickerSet) GetTrendings(client *Client, offset int32, limit int32) (*TrendingStickerSets, error) {
	return client.GetTrendingStickerSets(s.StickerType, offset, limit)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (s StickerSet) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(s.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (s StickerSet) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(s.Name)
}

// ReorderInstalleds is a helper method for Client.ReorderInstalledStickerSets
func (s StickerSet) ReorderInstalleds(client *Client, stickerSetIds []string) (*Ok, error) {
	return client.ReorderInstalledStickerSets(s.StickerType, stickerSetIds)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StickerSet) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, s.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (s StickerSet) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(s.Name)
}

// SearchInstalleds is a helper method for Client.SearchInstalledStickerSets
func (s StickerSet) SearchInstalleds(client *Client, query string, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(s.StickerType, query, limit)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s StickerSet) SearchStickers(client *Client, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(s.StickerType, emojis, query, inputLanguageCodes, offset, limit)
}

// Search is a helper method for Client.SearchStickerSet
func (s StickerSet) Search(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(s.Name, ignoreCache)
}

// Searchs is a helper method for Client.SearchStickerSets
func (s StickerSet) Searchs(client *Client, query string) (*StickerSets, error) {
	return client.SearchStickerSets(s.StickerType, query)
}

// SetBotName is a helper method for Client.SetBotName
func (s StickerSet) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, s.Name)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StickerSet) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, s.Title)
}

// SetCustomEmojiThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (s StickerSet) SetCustomEmojiThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(s.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StickerSet) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, s.Name)
}

// SetOption is a helper method for Client.SetOption
func (s StickerSet) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(s.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (s StickerSet) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, s.Name)
}

// SetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StickerSet) SetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, s.Name, opts)
}

// SetTitle is a helper method for Client.SetStickerSetTitle
func (s StickerSet) SetTitle(client *Client) (*Ok, error) {
	return client.SetStickerSetTitle(s.Name, s.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StickerSet) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, s.Name)
}

// SetSupergroup is a helper method for Client.SetSupergroupStickerSet
func (s StickerSet) SetSupergroup(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupStickerSet(supergroupId, s.Id)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (s StickerSet) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, s.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (s StickerSet) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, s.Title, recordVideo, usePortraitOrientation)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StickerSetInfo) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, s.Name, sticker)
}

// ChangeStickerSet is a helper method for Client.ChangeStickerSet
func (s StickerSetInfo) ChangeStickerSet(client *Client, setId string) (*Ok, error) {
	return client.ChangeStickerSet(setId, s.IsInstalled, s.IsArchived)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (s StickerSetInfo) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(s.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (s StickerSetInfo) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(s.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (s StickerSetInfo) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, s.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StickerSetInfo) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StickerSetInfo) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, s.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StickerSetInfo) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, s.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StickerSetInfo) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, s.Name, receivedGiftIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (s StickerSetInfo) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, s.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StickerSetInfo) CreateNewStickerSet(client *Client, userId int64, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, s.Title, s.Name, s.StickerType, s.NeedsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (s StickerSetInfo) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(s.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (s StickerSetInfo) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, s.Name, storyIds)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StickerSetInfo) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, s.Title, startDate, isRtmpStream)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (s StickerSetInfo) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(s.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (s StickerSetInfo) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, s.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StickerSetInfo) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StickerSetInfo) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, s.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StickerSetInfo) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, s.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StickerSetInfo) GetAllStickerEmojis(client *Client, query string, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(s.StickerType, query, chatId, returnOnlyMainEmoji)
}

// GetArchivedStickerSets is a helper method for Client.GetArchivedStickerSets
func (s StickerSetInfo) GetArchivedStickerSets(client *Client, offsetStickerSetId string, limit int32) (*StickerSets, error) {
	return client.GetArchivedStickerSets(s.StickerType, offsetStickerSetId, limit)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (s StickerSetInfo) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(s.Name, typeField)
}

// GetInstalledStickerSets is a helper method for Client.GetInstalledStickerSets
func (s StickerSetInfo) GetInstalledStickerSets(client *Client) (*StickerSets, error) {
	return client.GetInstalledStickerSets(s.StickerType)
}

// GetOption is a helper method for Client.GetOption
func (s StickerSetInfo) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(s.Name)
}

// GetStickers is a helper method for Client.GetStickers
func (s StickerSetInfo) GetStickers(client *Client, query string, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(s.StickerType, query, limit, chatId)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (s StickerSetInfo) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(s.Title)
}

// GetTrendingStickerSets is a helper method for Client.GetTrendingStickerSets
func (s StickerSetInfo) GetTrendingStickerSets(client *Client, offset int32, limit int32) (*TrendingStickerSets, error) {
	return client.GetTrendingStickerSets(s.StickerType, offset, limit)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (s StickerSetInfo) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(s.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (s StickerSetInfo) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(s.Name)
}

// ReorderInstalledStickerSets is a helper method for Client.ReorderInstalledStickerSets
func (s StickerSetInfo) ReorderInstalledStickerSets(client *Client, stickerSetIds []string) (*Ok, error) {
	return client.ReorderInstalledStickerSets(s.StickerType, stickerSetIds)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StickerSetInfo) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, s.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (s StickerSetInfo) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(s.Name)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (s StickerSetInfo) SearchInstalledStickerSets(client *Client, query string, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(s.StickerType, query, limit)
}

// SearchStickers is a helper method for Client.SearchStickers
func (s StickerSetInfo) SearchStickers(client *Client, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(s.StickerType, emojis, query, inputLanguageCodes, offset, limit)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (s StickerSetInfo) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(s.Name, ignoreCache)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (s StickerSetInfo) SearchStickerSets(client *Client, query string) (*StickerSets, error) {
	return client.SearchStickerSets(s.StickerType, query)
}

// SetBotName is a helper method for Client.SetBotName
func (s StickerSetInfo) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, s.Name)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StickerSetInfo) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, s.Title)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (s StickerSetInfo) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(s.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StickerSetInfo) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, s.Name)
}

// SetOption is a helper method for Client.SetOption
func (s StickerSetInfo) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(s.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (s StickerSetInfo) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, s.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StickerSetInfo) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, s.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (s StickerSetInfo) SetStickerSetTitle(client *Client) (*Ok, error) {
	return client.SetStickerSetTitle(s.Name, s.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StickerSetInfo) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, s.Name)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (s StickerSetInfo) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, s.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (s StickerSetInfo) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, s.Title, recordVideo, usePortraitOrientation)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (s StorageStatistics) OptimizeStorage(client *Client, ttl int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(s.Size, ttl, s.Count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StorageStatisticsByChat) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StorageStatisticsByChat) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StorageStatisticsByChat) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StorageStatisticsByChat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StorageStatisticsByChat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StorageStatisticsByChat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StorageStatisticsByChat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StorageStatisticsByChat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StorageStatisticsByChat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StorageStatisticsByChat) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StorageStatisticsByChat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StorageStatisticsByChat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StorageStatisticsByChat) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StorageStatisticsByChat) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StorageStatisticsByChat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StorageStatisticsByChat) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StorageStatisticsByChat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StorageStatisticsByChat) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StorageStatisticsByChat) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StorageStatisticsByChat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StorageStatisticsByChat) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StorageStatisticsByChat) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StorageStatisticsByChat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StorageStatisticsByChat) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StorageStatisticsByChat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StorageStatisticsByChat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StorageStatisticsByChat) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StorageStatisticsByChat) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StorageStatisticsByChat) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StorageStatisticsByChat) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StorageStatisticsByChat) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StorageStatisticsByChat) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StorageStatisticsByChat) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StorageStatisticsByChat) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StorageStatisticsByChat) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StorageStatisticsByChat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StorageStatisticsByChat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StorageStatisticsByChat) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StorageStatisticsByChat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StorageStatisticsByChat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StorageStatisticsByChat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StorageStatisticsByChat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StorageStatisticsByChat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StorageStatisticsByChat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StorageStatisticsByChat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StorageStatisticsByChat) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StorageStatisticsByChat) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StorageStatisticsByChat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StorageStatisticsByChat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StorageStatisticsByChat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StorageStatisticsByChat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StorageStatisticsByChat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StorageStatisticsByChat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StorageStatisticsByChat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StorageStatisticsByChat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StorageStatisticsByChat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StorageStatisticsByChat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StorageStatisticsByChat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StorageStatisticsByChat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StorageStatisticsByChat) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StorageStatisticsByChat) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StorageStatisticsByChat) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StorageStatisticsByChat) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StorageStatisticsByChat) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StorageStatisticsByChat) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StorageStatisticsByChat) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StorageStatisticsByChat) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StorageStatisticsByChat) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StorageStatisticsByChat) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StorageStatisticsByChat) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StorageStatisticsByChat) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StorageStatisticsByChat) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StorageStatisticsByChat) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StorageStatisticsByChat) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StorageStatisticsByChat) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StorageStatisticsByChat) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StorageStatisticsByChat) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StorageStatisticsByChat) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StorageStatisticsByChat) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StorageStatisticsByChat) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StorageStatisticsByChat) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StorageStatisticsByChat) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StorageStatisticsByChat) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StorageStatisticsByChat) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StorageStatisticsByChat) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StorageStatisticsByChat) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StorageStatisticsByChat) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StorageStatisticsByChat) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StorageStatisticsByChat) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StorageStatisticsByChat) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StorageStatisticsByChat) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StorageStatisticsByChat) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StorageStatisticsByChat) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StorageStatisticsByChat) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StorageStatisticsByChat) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StorageStatisticsByChat) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StorageStatisticsByChat) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StorageStatisticsByChat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StorageStatisticsByChat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StorageStatisticsByChat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StorageStatisticsByChat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StorageStatisticsByChat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StorageStatisticsByChat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StorageStatisticsByChat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StorageStatisticsByChat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StorageStatisticsByChat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StorageStatisticsByChat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StorageStatisticsByChat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StorageStatisticsByChat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StorageStatisticsByChat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StorageStatisticsByChat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StorageStatisticsByChat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StorageStatisticsByChat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StorageStatisticsByChat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StorageStatisticsByChat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StorageStatisticsByChat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StorageStatisticsByChat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StorageStatisticsByChat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StorageStatisticsByChat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StorageStatisticsByChat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StorageStatisticsByChat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StorageStatisticsByChat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StorageStatisticsByChat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StorageStatisticsByChat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StorageStatisticsByChat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StorageStatisticsByChat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StorageStatisticsByChat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StorageStatisticsByChat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StorageStatisticsByChat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StorageStatisticsByChat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StorageStatisticsByChat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StorageStatisticsByChat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StorageStatisticsByChat) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StorageStatisticsByChat) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StorageStatisticsByChat) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StorageStatisticsByChat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StorageStatisticsByChat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StorageStatisticsByChat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StorageStatisticsByChat) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StorageStatisticsByChat) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StorageStatisticsByChat) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StorageStatisticsByChat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StorageStatisticsByChat) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StorageStatisticsByChat) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StorageStatisticsByChat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StorageStatisticsByChat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (s StorageStatisticsByChat) OptimizeStorage(client *Client, ttl int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(s.Size, ttl, s.Count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StorageStatisticsByChat) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StorageStatisticsByChat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StorageStatisticsByChat) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StorageStatisticsByChat) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StorageStatisticsByChat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StorageStatisticsByChat) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StorageStatisticsByChat) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StorageStatisticsByChat) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StorageStatisticsByChat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StorageStatisticsByChat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StorageStatisticsByChat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StorageStatisticsByChat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StorageStatisticsByChat) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StorageStatisticsByChat) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StorageStatisticsByChat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StorageStatisticsByChat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StorageStatisticsByChat) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StorageStatisticsByChat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StorageStatisticsByChat) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StorageStatisticsByChat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StorageStatisticsByChat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StorageStatisticsByChat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StorageStatisticsByChat) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StorageStatisticsByChat) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StorageStatisticsByChat) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StorageStatisticsByChat) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StorageStatisticsByChat) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StorageStatisticsByChat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StorageStatisticsByChat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StorageStatisticsByChat) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StorageStatisticsByChat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StorageStatisticsByChat) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StorageStatisticsByChat) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StorageStatisticsByChat) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StorageStatisticsByChat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StorageStatisticsByChat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StorageStatisticsByChat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StorageStatisticsByChat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StorageStatisticsByChat) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StorageStatisticsByChat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StorageStatisticsByChat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StorageStatisticsByChat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StorageStatisticsByChat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StorageStatisticsByChat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StorageStatisticsByChat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StorageStatisticsByChat) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StorageStatisticsByChat) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StorageStatisticsByChat) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StorageStatisticsByChat) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StorageStatisticsByChat) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StorageStatisticsByChat) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StorageStatisticsByChat) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StorageStatisticsByChat) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StorageStatisticsByChat) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StorageStatisticsByChat) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StorageStatisticsByChat) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StorageStatisticsByChat) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StorageStatisticsByChat) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StorageStatisticsByChat) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StorageStatisticsByChat) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StorageStatisticsByChat) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StorageStatisticsByChat) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StorageStatisticsByChat) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StorageStatisticsByChat) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StorageStatisticsByChat) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StorageStatisticsByChat) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StorageStatisticsByChat) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StorageStatisticsByChat) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StorageStatisticsByChat) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StorageStatisticsByChat) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StorageStatisticsByChat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StorageStatisticsByChat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StorageStatisticsByChat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StorageStatisticsByChat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StorageStatisticsByChat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StorageStatisticsByChat) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StorageStatisticsByChat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StorageStatisticsByChat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StorageStatisticsByChat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StorageStatisticsByChat) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StorageStatisticsByChat) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StorageStatisticsByChat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StorageStatisticsByChat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StorageStatisticsByChat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StorageStatisticsByChat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StorageStatisticsByChat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StorageStatisticsByChat) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StorageStatisticsByChat) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StorageStatisticsByChat) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StorageStatisticsByChat) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StorageStatisticsByChat) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StorageStatisticsByChat) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StorageStatisticsByChat) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StorageStatisticsByChat) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StorageStatisticsByChat) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StorageStatisticsByChat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StorageStatisticsByChat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StorageStatisticsByChat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StorageStatisticsByChat) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StorageStatisticsByChat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StorageStatisticsByChat) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StorageStatisticsByChat) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StorageStatisticsByChat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StorageStatisticsByChat) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StorageStatisticsByChat) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StorageStatisticsByChat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (s StorageStatisticsByFileType) OptimizeStorage(client *Client, ttl int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(s.Size, ttl, s.Count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StorePaymentPurposeGiftedStars) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StorePaymentPurposeGiftedStars) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s StorePaymentPurposeGiftedStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, s.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StorePaymentPurposeGiftedStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, s.StarCount, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StorePaymentPurposeGiftedStars) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StorePaymentPurposeGiftedStars) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StorePaymentPurposeGiftedStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, s.StarCount)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StorePaymentPurposeGiftedStars) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposeGiftedStars) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// CreateCall is a helper method for Client.CreateCall
func (s StorePaymentPurposeGiftedStars) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StorePaymentPurposeGiftedStars) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StorePaymentPurposeGiftedStars) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StorePaymentPurposeGiftedStars) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (s StorePaymentPurposeGiftedStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, s.StarCount)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StorePaymentPurposeGiftedStars) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StorePaymentPurposeGiftedStars) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StorePaymentPurposeGiftedStars) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StorePaymentPurposeGiftedStars) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StorePaymentPurposeGiftedStars) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StorePaymentPurposeGiftedStars) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StorePaymentPurposeGiftedStars) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StorePaymentPurposeGiftedStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, s.StarCount, password)
}

// GetUser is a helper method for Client.GetUser
func (s StorePaymentPurposeGiftedStars) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StorePaymentPurposeGiftedStars) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StorePaymentPurposeGiftedStars) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StorePaymentPurposeGiftedStars) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StorePaymentPurposeGiftedStars) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StorePaymentPurposeGiftedStars) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StorePaymentPurposeGiftedStars) GiftPremiumWithStars(client *Client, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, s.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (s StorePaymentPurposeGiftedStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, s.StarCount)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StorePaymentPurposeGiftedStars) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s StorePaymentPurposeGiftedStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, s.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StorePaymentPurposeGiftedStars) PlaceGiftAuctionBid(client *Client, giftId string, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, s.StarCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StorePaymentPurposeGiftedStars) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StorePaymentPurposeGiftedStars) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StorePaymentPurposeGiftedStars) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StorePaymentPurposeGiftedStars) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s StorePaymentPurposeGiftedStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, s.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposeGiftedStars) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StorePaymentPurposeGiftedStars) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StorePaymentPurposeGiftedStars) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StorePaymentPurposeGiftedStars) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StorePaymentPurposeGiftedStars) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StorePaymentPurposeGiftedStars) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StorePaymentPurposeGiftedStars) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StorePaymentPurposeGiftedStars) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StorePaymentPurposeGiftedStars) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StorePaymentPurposeGiftedStars) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StorePaymentPurposeGiftedStars) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StorePaymentPurposeGiftedStars) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StorePaymentPurposeGiftedStars) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (s StorePaymentPurposeGiftedStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, s.StarCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StorePaymentPurposeGiftedStars) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (s StorePaymentPurposeGiftedStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, s.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (s StorePaymentPurposeGiftedStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, s.StarCount)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StorePaymentPurposeGiftedStars) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StorePaymentPurposePremiumGift) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StorePaymentPurposePremiumGift) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StorePaymentPurposePremiumGift) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StorePaymentPurposePremiumGift) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StorePaymentPurposePremiumGift) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposePremiumGift) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// CreateCall is a helper method for Client.CreateCall
func (s StorePaymentPurposePremiumGift) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StorePaymentPurposePremiumGift) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StorePaymentPurposePremiumGift) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StorePaymentPurposePremiumGift) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StorePaymentPurposePremiumGift) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StorePaymentPurposePremiumGift) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StorePaymentPurposePremiumGift) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StorePaymentPurposePremiumGift) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (s StorePaymentPurposePremiumGift) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(s.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (s StorePaymentPurposePremiumGift) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(s.Text)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StorePaymentPurposePremiumGift) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StorePaymentPurposePremiumGift) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StorePaymentPurposePremiumGift) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StorePaymentPurposePremiumGift) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StorePaymentPurposePremiumGift) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StorePaymentPurposePremiumGift) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StorePaymentPurposePremiumGift) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StorePaymentPurposePremiumGift) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StorePaymentPurposePremiumGift) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StorePaymentPurposePremiumGift) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, s.Text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StorePaymentPurposePremiumGift) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (s StorePaymentPurposePremiumGift) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(s.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StorePaymentPurposePremiumGift) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, s.Text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StorePaymentPurposePremiumGift) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StorePaymentPurposePremiumGift) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StorePaymentPurposePremiumGift) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StorePaymentPurposePremiumGift) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SearchQuote is a helper method for Client.SearchQuote
func (s StorePaymentPurposePremiumGift) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(s.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (s StorePaymentPurposePremiumGift) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, s.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (s StorePaymentPurposePremiumGift) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, s.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StorePaymentPurposePremiumGift) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, s.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposePremiumGift) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StorePaymentPurposePremiumGift) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StorePaymentPurposePremiumGift) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StorePaymentPurposePremiumGift) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StorePaymentPurposePremiumGift) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StorePaymentPurposePremiumGift) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StorePaymentPurposePremiumGift) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StorePaymentPurposePremiumGift) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StorePaymentPurposePremiumGift) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StorePaymentPurposePremiumGift) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StorePaymentPurposePremiumGift) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StorePaymentPurposePremiumGift) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StorePaymentPurposePremiumGift) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StorePaymentPurposePremiumGift) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// TranslateText is a helper method for Client.TranslateText
func (s StorePaymentPurposePremiumGift) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(s.Text, toLanguageCode)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StorePaymentPurposePremiumGift) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StorePaymentPurposePremiumGiftCodes) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, s.UserIds)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposePremiumGiftCodes) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (s StorePaymentPurposePremiumGiftCodes) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(s.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StorePaymentPurposePremiumGiftCodes) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, s.UserIds, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (s StorePaymentPurposePremiumGiftCodes) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(s.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (s StorePaymentPurposePremiumGiftCodes) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(s.Text)
}

// GetPremiumGiveawayPaymentOptions is a helper method for Client.GetPremiumGiveawayPaymentOptions
func (s StorePaymentPurposePremiumGiftCodes) GetPremiumGiveawayPaymentOptions(client *Client) (*PremiumGiveawayPaymentOptions, error) {
	return client.GetPremiumGiveawayPaymentOptions(s.BoostedChatId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StorePaymentPurposePremiumGiftCodes) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, s.Text)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (s StorePaymentPurposePremiumGiftCodes) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, s.UserIds)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (s StorePaymentPurposePremiumGiftCodes) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(s.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StorePaymentPurposePremiumGiftCodes) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, s.Text, isPrivate)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (s StorePaymentPurposePremiumGiftCodes) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(s.UserIds)
}

// SearchQuote is a helper method for Client.SearchQuote
func (s StorePaymentPurposePremiumGiftCodes) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(s.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (s StorePaymentPurposePremiumGiftCodes) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, s.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (s StorePaymentPurposePremiumGiftCodes) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, s.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StorePaymentPurposePremiumGiftCodes) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, s.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposePremiumGiftCodes) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (s StorePaymentPurposePremiumGiftCodes) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(s.UserIds)
}

// TranslateText is a helper method for Client.TranslateText
func (s StorePaymentPurposePremiumGiftCodes) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(s.Text, toLanguageCode)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposePremiumGiveaway) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s StorePaymentPurposePremiumGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string, winnerCount int32, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, s.Parameters, winnerCount, starCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposePremiumGiveaway) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposePremiumSubscription) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, currency string, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, s.IsRestore, currency, amount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s StorePaymentPurposeStarGiveaway) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, s.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StorePaymentPurposeStarGiveaway) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, s.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StorePaymentPurposeStarGiveaway) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, s.StarCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposeStarGiveaway) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (s StorePaymentPurposeStarGiveaway) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, s.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StorePaymentPurposeStarGiveaway) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, s.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StorePaymentPurposeStarGiveaway) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, s.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (s StorePaymentPurposeStarGiveaway) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, s.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s StorePaymentPurposeStarGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, s.Parameters, s.WinnerCount, s.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StorePaymentPurposeStarGiveaway) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, s.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s StorePaymentPurposeStarGiveaway) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, s.StarCount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposeStarGiveaway) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (s StorePaymentPurposeStarGiveaway) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, s.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (s StorePaymentPurposeStarGiveaway) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, s.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (s StorePaymentPurposeStarGiveaway) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, s.StarCount)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StorePaymentPurposeStars) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StorePaymentPurposeStars) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StorePaymentPurposeStars) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StorePaymentPurposeStars) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StorePaymentPurposeStars) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StorePaymentPurposeStars) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StorePaymentPurposeStars) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StorePaymentPurposeStars) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (s StorePaymentPurposeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, s.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StorePaymentPurposeStars) AddPendingPaidMessageReaction(client *Client, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, s.StarCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StorePaymentPurposeStars) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StorePaymentPurposeStars) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StorePaymentPurposeStars) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StorePaymentPurposeStars) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StorePaymentPurposeStars) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StorePaymentPurposeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, s.StarCount)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StorePaymentPurposeStars) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (s StorePaymentPurposeStars) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(s.Currency, s.Amount)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StorePaymentPurposeStars) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StorePaymentPurposeStars) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StorePaymentPurposeStars) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StorePaymentPurposeStars) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StorePaymentPurposeStars) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StorePaymentPurposeStars) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StorePaymentPurposeStars) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StorePaymentPurposeStars) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StorePaymentPurposeStars) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StorePaymentPurposeStars) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StorePaymentPurposeStars) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StorePaymentPurposeStars) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StorePaymentPurposeStars) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StorePaymentPurposeStars) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StorePaymentPurposeStars) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StorePaymentPurposeStars) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StorePaymentPurposeStars) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StorePaymentPurposeStars) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StorePaymentPurposeStars) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StorePaymentPurposeStars) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StorePaymentPurposeStars) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StorePaymentPurposeStars) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StorePaymentPurposeStars) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StorePaymentPurposeStars) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (s StorePaymentPurposeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, s.StarCount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StorePaymentPurposeStars) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StorePaymentPurposeStars) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StorePaymentPurposeStars) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StorePaymentPurposeStars) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StorePaymentPurposeStars) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StorePaymentPurposeStars) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StorePaymentPurposeStars) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StorePaymentPurposeStars) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StorePaymentPurposeStars) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StorePaymentPurposeStars) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StorePaymentPurposeStars) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StorePaymentPurposeStars) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StorePaymentPurposeStars) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StorePaymentPurposeStars) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StorePaymentPurposeStars) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StorePaymentPurposeStars) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StorePaymentPurposeStars) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StorePaymentPurposeStars) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StorePaymentPurposeStars) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StorePaymentPurposeStars) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StorePaymentPurposeStars) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StorePaymentPurposeStars) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StorePaymentPurposeStars) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StorePaymentPurposeStars) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StorePaymentPurposeStars) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StorePaymentPurposeStars) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StorePaymentPurposeStars) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StorePaymentPurposeStars) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StorePaymentPurposeStars) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StorePaymentPurposeStars) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StorePaymentPurposeStars) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StorePaymentPurposeStars) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StorePaymentPurposeStars) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StorePaymentPurposeStars) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StorePaymentPurposeStars) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StorePaymentPurposeStars) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StorePaymentPurposeStars) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StorePaymentPurposeStars) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StorePaymentPurposeStars) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StorePaymentPurposeStars) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StorePaymentPurposeStars) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StorePaymentPurposeStars) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StorePaymentPurposeStars) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StorePaymentPurposeStars) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StorePaymentPurposeStars) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StorePaymentPurposeStars) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StorePaymentPurposeStars) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StorePaymentPurposeStars) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StorePaymentPurposeStars) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StorePaymentPurposeStars) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StorePaymentPurposeStars) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StorePaymentPurposeStars) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StorePaymentPurposeStars) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StorePaymentPurposeStars) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StorePaymentPurposeStars) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StorePaymentPurposeStars) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StorePaymentPurposeStars) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StorePaymentPurposeStars) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StorePaymentPurposeStars) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StorePaymentPurposeStars) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StorePaymentPurposeStars) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StorePaymentPurposeStars) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StorePaymentPurposeStars) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StorePaymentPurposeStars) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StorePaymentPurposeStars) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StorePaymentPurposeStars) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StorePaymentPurposeStars) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StorePaymentPurposeStars) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StorePaymentPurposeStars) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StorePaymentPurposeStars) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StorePaymentPurposeStars) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StorePaymentPurposeStars) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StorePaymentPurposeStars) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StorePaymentPurposeStars) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StorePaymentPurposeStars) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StorePaymentPurposeStars) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StorePaymentPurposeStars) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StorePaymentPurposeStars) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StorePaymentPurposeStars) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StorePaymentPurposeStars) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StorePaymentPurposeStars) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StorePaymentPurposeStars) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StorePaymentPurposeStars) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StorePaymentPurposeStars) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StorePaymentPurposeStars) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StorePaymentPurposeStars) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StorePaymentPurposeStars) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StorePaymentPurposeStars) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StorePaymentPurposeStars) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StorePaymentPurposeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, s.StarCount, password)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StorePaymentPurposeStars) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StorePaymentPurposeStars) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StorePaymentPurposeStars) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StorePaymentPurposeStars) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StorePaymentPurposeStars) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StorePaymentPurposeStars) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StorePaymentPurposeStars) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StorePaymentPurposeStars) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StorePaymentPurposeStars) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StorePaymentPurposeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, s.StarCount, monthCount, text)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StorePaymentPurposeStars) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (s StorePaymentPurposeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, s.StarCount)
}

// JoinChat is a helper method for Client.JoinChat
func (s StorePaymentPurposeStars) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (s StorePaymentPurposeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, s.StarCount)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StorePaymentPurposeStars) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StorePaymentPurposeStars) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StorePaymentPurposeStars) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StorePaymentPurposeStars) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StorePaymentPurposeStars) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StorePaymentPurposeStars) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StorePaymentPurposeStars) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StorePaymentPurposeStars) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StorePaymentPurposeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, s.StarCount, userId, text, isPrivate)
}

// PostStory is a helper method for Client.PostStory
func (s StorePaymentPurposeStars) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StorePaymentPurposeStars) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StorePaymentPurposeStars) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StorePaymentPurposeStars) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StorePaymentPurposeStars) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StorePaymentPurposeStars) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StorePaymentPurposeStars) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StorePaymentPurposeStars) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StorePaymentPurposeStars) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StorePaymentPurposeStars) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StorePaymentPurposeStars) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StorePaymentPurposeStars) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StorePaymentPurposeStars) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StorePaymentPurposeStars) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StorePaymentPurposeStars) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StorePaymentPurposeStars) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StorePaymentPurposeStars) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StorePaymentPurposeStars) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StorePaymentPurposeStars) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StorePaymentPurposeStars) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StorePaymentPurposeStars) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StorePaymentPurposeStars) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StorePaymentPurposeStars) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StorePaymentPurposeStars) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StorePaymentPurposeStars) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StorePaymentPurposeStars) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StorePaymentPurposeStars) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StorePaymentPurposeStars) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StorePaymentPurposeStars) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StorePaymentPurposeStars) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StorePaymentPurposeStars) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StorePaymentPurposeStars) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StorePaymentPurposeStars) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (s StorePaymentPurposeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, s.StarCount)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StorePaymentPurposeStars) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StorePaymentPurposeStars) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StorePaymentPurposeStars) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StorePaymentPurposeStars) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StorePaymentPurposeStars) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StorePaymentPurposeStars) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StorePaymentPurposeStars) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StorePaymentPurposeStars) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StorePaymentPurposeStars) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StorePaymentPurposeStars) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (s StorePaymentPurposeStars) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, s.Currency, s.Amount)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StorePaymentPurposeStars) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StorePaymentPurposeStars) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StorePaymentPurposeStars) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StorePaymentPurposeStars) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StorePaymentPurposeStars) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StorePaymentPurposeStars) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StorePaymentPurposeStars) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StorePaymentPurposeStars) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StorePaymentPurposeStars) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StorePaymentPurposeStars) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StorePaymentPurposeStars) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StorePaymentPurposeStars) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StorePaymentPurposeStars) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StorePaymentPurposeStars) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StorePaymentPurposeStars) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StorePaymentPurposeStars) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StorePaymentPurposeStars) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StorePaymentPurposeStars) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StorePaymentPurposeStars) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StorePaymentPurposeStars) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StorePaymentPurposeStars) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StorePaymentPurposeStars) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StorePaymentPurposeStars) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StorePaymentPurposeStars) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StorePaymentPurposeStars) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StorePaymentPurposeStars) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StorePaymentPurposeStars) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StorePaymentPurposeStars) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StorePaymentPurposeStars) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StorePaymentPurposeStars) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StorePaymentPurposeStars) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StorePaymentPurposeStars) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StorePaymentPurposeStars) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StorePaymentPurposeStars) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StorePaymentPurposeStars) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StorePaymentPurposeStars) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StorePaymentPurposeStars) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StorePaymentPurposeStars) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StorePaymentPurposeStars) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StorePaymentPurposeStars) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StorePaymentPurposeStars) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StorePaymentPurposeStars) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StorePaymentPurposeStars) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StorePaymentPurposeStars) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StorePaymentPurposeStars) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StorePaymentPurposeStars) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StorePaymentPurposeStars) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StorePaymentPurposeStars) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StorePaymentPurposeStars) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StorePaymentPurposeStars) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StorePaymentPurposeStars) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StorePaymentPurposeStars) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StorePaymentPurposeStars) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StorePaymentPurposeStars) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (s StorePaymentPurposeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, s.StarCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StorePaymentPurposeStars) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (s StorePaymentPurposeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, s.StarCount)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StorePaymentPurposeStars) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StorePaymentPurposeStars) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StorePaymentPurposeStars) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StorePaymentPurposeStars) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StorePaymentPurposeStars) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StorePaymentPurposeStars) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (s StorePaymentPurposeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, s.StarCount)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StorePaymentPurposeStars) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// Close is a helper method for Client.CloseStory
func (s Story) Close(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, s.Id)
}

// DeleteBusiness is a helper method for Client.DeleteBusinessStory
func (s Story) DeleteBusiness(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, s.Id)
}

// Delete is a helper method for Client.DeleteStory
func (s Story) Delete(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, s.Id)
}

// EditBusiness is a helper method for Client.EditBusinessStory
func (s Story) EditBusiness(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, s.Id, content, areas, s.Caption, s.PrivacySettings)
}

// Edit is a helper method for Client.EditStory
func (s Story) Edit(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, s.Id, opts)
}

// EditCover is a helper method for Client.EditStoryCover
func (s Story) EditCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, s.Id, coverFrameTimestamp)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s Story) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, s.Date)
}

// GetChatInteractions is a helper method for Client.GetChatStoryInteractions
func (s Story) GetChatInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, s.Id, preferForwards, offset, limit, opts)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s Story) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, s.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (s Story) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, s.Date)
}

// Get is a helper method for Client.GetStory
func (s Story) Get(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, s.Id, onlyLocal)
}

// GetInteractions is a helper method for Client.GetStoryInteractions
func (s Story) GetInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(s.Id, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetPublicForwards is a helper method for Client.GetStoryPublicForwards
func (s Story) GetPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, s.Id, offset, limit)
}

// GetStatistics is a helper method for Client.GetStoryStatistics
func (s Story) GetStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, s.Id, isDark)
}

// Open is a helper method for Client.OpenStory
func (s Story) Open(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, s.Id)
}

// Post is a helper method for Client.PostStory
func (s Story) Post(client *Client, chatId int64, content *InputStoryContent, activePeriod int32, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(chatId, content, s.PrivacySettings, s.AlbumIds, activePeriod, s.IsPostedToChatPage, protectContent, opts)
}

// Report is a helper method for Client.ReportStory
func (s Story) Report(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, s.Id, optionId, text)
}

// SetPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (s Story) SetPrivacySettings(client *Client) (*Ok, error) {
	return client.SetStoryPrivacySettings(s.Id, s.PrivacySettings)
}

// SetReaction is a helper method for Client.SetStoryReaction
func (s Story) SetReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, s.Id, updateRecentReactions, opts)
}

// StartLive is a helper method for Client.StartLiveStory
func (s Story) StartLive(client *Client, chatId int64, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, s.PrivacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// ToggleIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (s Story) ToggleIsPostedToChatPage(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, s.Id, s.IsPostedToChatPage)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StoryAlbum) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, s.Name, sticker)
}

// AddStories is a helper method for Client.AddStoryAlbumStories
func (s StoryAlbum) AddStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(chatId, s.Id, storyIds)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (s StoryAlbum) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(s.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (s StoryAlbum) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(s.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (s StoryAlbum) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, s.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StoryAlbum) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StoryAlbum) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, s.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StoryAlbum) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, s.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StoryAlbum) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, s.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StoryAlbum) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, s.Name, stickerType, needsRepainting, stickers, source)
}

// Create is a helper method for Client.CreateStoryAlbum
func (s StoryAlbum) Create(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, s.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (s StoryAlbum) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(s.Name)
}

// Delete is a helper method for Client.DeleteStoryAlbum
func (s StoryAlbum) Delete(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteStoryAlbum(chatId, s.Id)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (s StoryAlbum) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, s.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StoryAlbum) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, s.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StoryAlbum) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, s.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StoryAlbum) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, s.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (s StoryAlbum) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(s.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (s StoryAlbum) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(s.Name)
}

// GetStories is a helper method for Client.GetStoryAlbumStories
func (s StoryAlbum) GetStories(client *Client, chatId int64, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(chatId, s.Id, offset, limit)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (s StoryAlbum) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(s.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (s StoryAlbum) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(s.Name)
}

// RemoveStories is a helper method for Client.RemoveStoryAlbumStories
func (s StoryAlbum) RemoveStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(chatId, s.Id, storyIds)
}

// ReorderStories is a helper method for Client.ReorderStoryAlbumStories
func (s StoryAlbum) ReorderStories(client *Client, chatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(chatId, s.Id, storyIds)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StoryAlbum) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, s.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (s StoryAlbum) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(s.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (s StoryAlbum) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(s.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (s StoryAlbum) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, s.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (s StoryAlbum) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(s.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StoryAlbum) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, s.Name)
}

// SetOption is a helper method for Client.SetOption
func (s StoryAlbum) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(s.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (s StoryAlbum) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, s.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StoryAlbum) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, s.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (s StoryAlbum) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(s.Name, title)
}

// SetName is a helper method for Client.SetStoryAlbumName
func (s StoryAlbum) SetName(client *Client, chatId int64) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, s.Id, s.Name)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (s StoryAreaTypeLink) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, s.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (s StoryAreaTypeLink) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, s.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (s StoryAreaTypeLink) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, s.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (s StoryAreaTypeLink) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(s.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (s StoryAreaTypeLink) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(s.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (s StoryAreaTypeLink) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, s.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (s StoryAreaTypeLink) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(s.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StoryAreaTypeLink) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, s.Url, parameters, opts)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (s StoryAreaTypeLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(s.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StoryAreaTypeLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(s.Location, zoom, width, height, scale, chatId)
}

// SearchPublicStoriesByLocation is a helper method for Client.SearchPublicStoriesByLocation
func (s StoryAreaTypeLocation) SearchPublicStoriesByLocation(client *Client, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByLocation(s.Address, offset, limit)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StoryAreaTypeMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StoryAreaTypeMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StoryAreaTypeMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StoryAreaTypeMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, s.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StoryAreaTypeMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, s.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StoryAreaTypeMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StoryAreaTypeMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, s.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StoryAreaTypeMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, s.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StoryAreaTypeMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, s.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StoryAreaTypeMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StoryAreaTypeMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StoryAreaTypeMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, s.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StoryAreaTypeMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (s StoryAreaTypeMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(s.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (s StoryAreaTypeMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StoryAreaTypeMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StoryAreaTypeMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StoryAreaTypeMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, s.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StoryAreaTypeMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, s.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StoryAreaTypeMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StoryAreaTypeMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StoryAreaTypeMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StoryAreaTypeMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StoryAreaTypeMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StoryAreaTypeMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StoryAreaTypeMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, s.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StoryAreaTypeMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, s.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StoryAreaTypeMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StoryAreaTypeMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StoryAreaTypeMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StoryAreaTypeMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StoryAreaTypeMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StoryAreaTypeMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StoryAreaTypeMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, s.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StoryAreaTypeMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StoryAreaTypeMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StoryAreaTypeMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StoryAreaTypeMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StoryAreaTypeMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StoryAreaTypeMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StoryAreaTypeMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StoryAreaTypeMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, s.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StoryAreaTypeMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StoryAreaTypeMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StoryAreaTypeMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StoryAreaTypeMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StoryAreaTypeMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StoryAreaTypeMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StoryAreaTypeMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StoryAreaTypeMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, s.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StoryAreaTypeMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, s.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StoryAreaTypeMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, s.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StoryAreaTypeMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StoryAreaTypeMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, s.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StoryAreaTypeMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, s.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StoryAreaTypeMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, s.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (s StoryAreaTypeMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, s.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StoryAreaTypeMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StoryAreaTypeMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StoryAreaTypeMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, s.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StoryAreaTypeMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, s.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StoryAreaTypeMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StoryAreaTypeMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StoryAreaTypeMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StoryAreaTypeMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StoryAreaTypeMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StoryAreaTypeMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StoryAreaTypeMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StoryAreaTypeMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StoryAreaTypeMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StoryAreaTypeMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StoryAreaTypeMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}
