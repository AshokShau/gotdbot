package gotdbot

// ImportMessages is a helper method for Client.ImportMessages
func (c ChatActiveStories) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(c.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (c ChatActiveStories) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(c.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (c ChatActiveStories) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(c.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (c ChatActiveStories) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(c.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c ChatActiveStories) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(c.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (c ChatActiveStories) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(c.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (c ChatActiveStories) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(c.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c ChatActiveStories) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(c.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c ChatActiveStories) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(c.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (c ChatActiveStories) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(c.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (c ChatActiveStories) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(c.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatActiveStories) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(c.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c ChatActiveStories) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(c.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c ChatActiveStories) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(c.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (c ChatActiveStories) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(c.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (c ChatActiveStories) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(c.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (c ChatActiveStories) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(c.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (c ChatActiveStories) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(c.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (c ChatActiveStories) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(c.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c ChatActiveStories) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, c.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c ChatActiveStories) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(c.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (c ChatActiveStories) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(c.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (c ChatActiveStories) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(c.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c ChatActiveStories) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(c.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c ChatActiveStories) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(c.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (c ChatActiveStories) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(c.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (c ChatActiveStories) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (c ChatActiveStories) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, c.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (c ChatActiveStories) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(c.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (c ChatActiveStories) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (c ChatActiveStories) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(c.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (c ChatActiveStories) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(c.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (c ChatActiveStories) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(c.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (c ChatActiveStories) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(c.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c ChatActiveStories) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(c.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c ChatActiveStories) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(c.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c ChatActiveStories) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(c.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (c ChatActiveStories) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(c.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c ChatActiveStories) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(c.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c ChatActiveStories) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, c.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (c ChatActiveStories) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(c.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (c ChatActiveStories) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(c.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (c ChatActiveStories) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(c.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (c ChatActiveStories) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(c.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c ChatActiveStories) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, c.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (c ChatActiveStories) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (c ChatActiveStories) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (c ChatActiveStories) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(c.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (c ChatActiveStories) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(c.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (c ChatActiveStories) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(c.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (c ChatActiveStories) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(c.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (c ChatActiveStories) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(c.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c ChatActiveStories) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(c.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c ChatActiveStories) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, c.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (c ChatActiveStories) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(c.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetList is a helper method for Client.SetChatActiveStoriesList
func (c ChatActiveStories) SetList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(c.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (c ChatActiveStories) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(c.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (c ChatActiveStories) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(c.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (c ChatActiveStories) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(c.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (c ChatActiveStories) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(c.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (c ChatActiveStories) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(c.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (c ChatActiveStories) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(c.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (c ChatActiveStories) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(c.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (c ChatActiveStories) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(c.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (c ChatActiveStories) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(c.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (c ChatActiveStories) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(c.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatActiveStories) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(c.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (c ChatActiveStories) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(c.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (c ChatActiveStories) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(c.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (c ChatActiveStories) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(c.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (c ChatActiveStories) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(c.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (c ChatActiveStories) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(c.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (c ChatActiveStories) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(c.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (c ChatActiveStories) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(c.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (c ChatActiveStories) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(c.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (c ChatActiveStories) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(c.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (c ChatActiveStories) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(c.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (c ChatActiveStories) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(c.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (c ChatActiveStories) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(c.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (c ChatActiveStories) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(c.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatActiveStories) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(c.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c ChatActiveStories) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(c.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c ChatActiveStories) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(c.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c ChatActiveStories) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(c.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (c ChatActiveStories) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(c.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (c ChatActiveStories) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(c.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c ChatActiveStories) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(c.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatActiveStories) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(c.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (c ChatActiveStories) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(c.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (c ChatActiveStories) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(c.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c ChatActiveStories) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(c.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (c ChatActiveStories) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(c.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c ChatActiveStories) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, c.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c ChatActiveStories) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(c.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c ChatActiveStories) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(c.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (c ChatActiveStories) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(c.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (c ChatActiveStories) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(c.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (c ChatActiveStories) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(c.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c ChatActiveStories) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(c.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (c ChatActiveStories) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(c.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (c ChatActiveStories) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, c.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (c ChatActiveStories) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(c.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (c ChatActiveStories) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(c.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (c ChatActiveStories) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(c.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (c ChatActiveStories) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(c.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c ChatActiveStories) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(c.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c ChatActiveStories) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(c.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatActiveStories) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(c.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c ChatActiveStories) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(c.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (c ChatActiveStories) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(c.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (c ChatActiveStories) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(c.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (c ChatActiveStories) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(c.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (c ChatActiveStories) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(c.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (c ChatActiveStories) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(c.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (c ChatActiveStories) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(c.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatAdministrator) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatAdministrator) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatAdministrator) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatAdministrator) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatAdministrator) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatAdministrator) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatAdministrator) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatAdministrator) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatAdministrator) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatAdministrator) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatAdministrator) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatAdministrator) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatAdministrator) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatAdministrator) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatAdministrator) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatAdministrator) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatAdministrator) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatAdministrator) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatAdministrator) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatAdministrator) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatAdministrator) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatAdministrator) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatAdministrator) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatAdministrator) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatAdministrator) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatAdministrator) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatAdministrator) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatAdministrator) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatAdministrator) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatAdministrator) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatAdministrator) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatAdministrator) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatAdministrator) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatAdministrator) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatAdministrator) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatAdministrator) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatAdministrator) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatAdministrator) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatAdministrator) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatAdministrator) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatAdministrator) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatAdministrator) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatAdministrator) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// Set is a helper method for Client.SetChatBackground
func (c ChatBackground) Set(client *Client, chatId int64, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(chatId, c.DarkThemeDimming, onlyForSelf, opts)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatBoost) CreateChatInviteLink(client *Client, chatId int64, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, c.ExpirationDate, memberLimit, createsJoinRequest)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatBoost) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, c.StartDate, isRtmpStream)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatBoost) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, c.ExpirationDate, memberLimit, createsJoinRequest)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (c ChatBoost) OptimizeStorage(client *Client, size int64, ttl int32, immunityDelay int32, fileTypes []*FileType, chatIds []int64, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, c.Count, immunityDelay, fileTypes, chatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// Get is a helper method for Client.GetChatBoostLevelFeatures
func (c ChatBoostLevelFeatures) Get(client *Client, isChannel bool) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(isChannel, c.Level)
}

// ConfirmQrCodeAuthentication is a helper method for Client.ConfirmQrCodeAuthentication
func (c ChatBoostLink) ConfirmQrCodeAuthentication(client *Client) (*Session, error) {
	return client.ConfirmQrCodeAuthentication(c.Link)
}

// DeleteBusinessChatLink is a helper method for Client.DeleteBusinessChatLink
func (c ChatBoostLink) DeleteBusinessChatLink(client *Client) (*Ok, error) {
	return client.DeleteBusinessChatLink(c.Link)
}

// EditBusinessChatLink is a helper method for Client.EditBusinessChatLink
func (c ChatBoostLink) EditBusinessChatLink(client *Client, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	return client.EditBusinessChatLink(c.Link, linkInfo)
}

// GetDeepLinkInfo is a helper method for Client.GetDeepLinkInfo
func (c ChatBoostLink) GetDeepLinkInfo(client *Client) (*DeepLinkInfo, error) {
	return client.GetDeepLinkInfo(c.Link)
}

// GetExternalLink is a helper method for Client.GetExternalLink
func (c ChatBoostLink) GetExternalLink(client *Client, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetExternalLink(c.Link, allowWriteAccess)
}

// GetExternalLinkInfo is a helper method for Client.GetExternalLinkInfo
func (c ChatBoostLink) GetExternalLinkInfo(client *Client) (*LoginUrlInfo, error) {
	return client.GetExternalLinkInfo(c.Link)
}

// GetInternalLinkType is a helper method for Client.GetInternalLinkType
func (c ChatBoostLink) GetInternalLinkType(client *Client) (*InternalLinkType, error) {
	return client.GetInternalLinkType(c.Link)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (c ChatBoostLink) SetBusinessAccountProfilePhoto(client *Client, businessConnectionId string, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(businessConnectionId, c.IsPublic, opts)
}

// SetProfilePhoto is a helper method for Client.SetProfilePhoto
func (c ChatBoostLink) SetProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetProfilePhoto(photo, c.IsPublic)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatBoostLinkInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(c.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (c ChatBoostLinkInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(c.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (c ChatBoostLinkInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(c.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c ChatBoostLinkInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(c.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c ChatBoostLinkInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, c.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (c ChatBoostLinkInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(c.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c ChatBoostLinkInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(c.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c ChatBoostLinkInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(c.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c ChatBoostLinkInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(c.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (c ChatBoostLinkInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(c.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (c ChatBoostLinkInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c ChatBoostLinkInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(c.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatBoostLinkInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(c.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (c ChatBoostLinkInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(c.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (c ChatBoostLinkInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(c.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c ChatBoostLinkInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(c.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c ChatBoostLinkInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(c.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c ChatBoostLinkInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(c.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (c ChatBoostLinkInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(c.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c ChatBoostLinkInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(c.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatBoostLinkInfo) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(c.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatBoostLinkInfo) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(c.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatBoostLinkInfo) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(c.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatBoostLinkInfo) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(c.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c ChatBoostLinkInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(c.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c ChatBoostLinkInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(c.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c ChatBoostLinkInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(c.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (c ChatBoostLinkInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(c.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (c ChatBoostLinkInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(c.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (c ChatBoostLinkInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(c.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (c ChatBoostLinkInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(c.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (c ChatBoostLinkInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(c.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c ChatBoostLinkInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(c.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (c ChatBoostLinkInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(c.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (c ChatBoostLinkInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(c.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (c ChatBoostLinkInfo) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(c.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (c ChatBoostLinkInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(c.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c ChatBoostLinkInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(c.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (c ChatBoostLinkInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(c.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c ChatBoostLinkInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c ChatBoostLinkInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, c.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c ChatBoostLinkInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c ChatBoostLinkInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c ChatBoostLinkInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, c.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c ChatBoostLinkInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatBoostLinkInfo) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(c.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatBoostLinkInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(c.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatBoostLinkInfo) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(c.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c ChatBoostLinkInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c ChatBoostLinkInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(c.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c ChatBoostLinkInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c ChatBoostLinkInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(c.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c ChatBoostLinkInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(c.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c ChatBoostLinkInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(c.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c ChatBoostLinkInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(c.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (c ChatBoostLinkInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(c.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (c ChatBoostLinkInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, c.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c ChatBoostLinkInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(c.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c ChatBoostLinkInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(c.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (c ChatBoostLinkInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(c.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (c ChatBoostLinkInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(c.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (c ChatBoostLinkInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(c.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (c ChatBoostLinkInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(c.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (c ChatBoostLinkInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(c.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (c ChatBoostLinkInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(c.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (c ChatBoostLinkInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(c.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (c ChatBoostLinkInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(c.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (c ChatBoostLinkInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(c.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (c ChatBoostLinkInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(c.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (c ChatBoostLinkInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(c.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c ChatBoostLinkInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(c.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (c ChatBoostLinkInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(c.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c ChatBoostLinkInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(c.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (c ChatBoostLinkInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(c.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c ChatBoostLinkInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(c.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (c ChatBoostLinkInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(c.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c ChatBoostLinkInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(c.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatBoostLinkInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(c.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (c ChatBoostLinkInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(c.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (c ChatBoostLinkInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(c.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (c ChatBoostLinkInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(c.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (c ChatBoostLinkInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(c.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (c ChatBoostLinkInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(c.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (c ChatBoostLinkInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(c.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (c ChatBoostLinkInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(c.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c ChatBoostLinkInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(c.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (c ChatBoostLinkInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(c.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (c ChatBoostLinkInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(c.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (c ChatBoostLinkInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(c.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (c ChatBoostLinkInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(c.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (c ChatBoostLinkInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(c.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (c ChatBoostLinkInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(c.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (c ChatBoostLinkInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(c.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (c ChatBoostLinkInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(c.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (c ChatBoostLinkInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(c.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatBoostLinkInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(c.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (c ChatBoostLinkInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(c.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (c ChatBoostLinkInfo) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(c.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (c ChatBoostLinkInfo) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(c.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (c ChatBoostLinkInfo) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(c.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (c ChatBoostLinkInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(c.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatBoostLinkInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(c.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c ChatBoostLinkInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(c.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c ChatBoostLinkInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, c.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (c ChatBoostLinkInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(c.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c ChatBoostLinkInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(c.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c ChatBoostLinkInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(c.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c ChatBoostLinkInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(c.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c ChatBoostLinkInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, c.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (c ChatBoostLinkInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(c.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c ChatBoostLinkInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(c.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c ChatBoostLinkInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(c.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c ChatBoostLinkInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(c.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c ChatBoostLinkInfo) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(c.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (c ChatBoostLinkInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(c.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c ChatBoostLinkInfo) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(c.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c ChatBoostLinkInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(c.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c ChatBoostLinkInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(c.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c ChatBoostLinkInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(c.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c ChatBoostLinkInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(c.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (c ChatBoostLinkInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(c.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c ChatBoostLinkInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(c.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c ChatBoostLinkInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(c.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c ChatBoostLinkInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(c.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c ChatBoostLinkInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(c.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c ChatBoostLinkInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(c.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c ChatBoostLinkInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(c.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c ChatBoostLinkInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(c.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (c ChatBoostLinkInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(c.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (c ChatBoostLinkInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, c.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (c ChatBoostLinkInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(c.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c ChatBoostLinkInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(c.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatBoostLinkInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(c.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (c ChatBoostLinkInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(c.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (c ChatBoostLinkInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(c.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c ChatBoostLinkInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(c.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c ChatBoostLinkInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(c.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (c ChatBoostLinkInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(c.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (c ChatBoostLinkInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(c.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (c ChatBoostLinkInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(c.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (c ChatBoostLinkInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(c.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c ChatBoostLinkInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(c.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (c ChatBoostLinkInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(c.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (c ChatBoostLinkInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(c.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c ChatBoostLinkInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(c.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c ChatBoostLinkInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(c.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (c ChatBoostLinkInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(c.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (c ChatBoostLinkInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(c.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatBoostLinkInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(c.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c ChatBoostLinkInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(c.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c ChatBoostLinkInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(c.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (c ChatBoostLinkInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(c.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (c ChatBoostLinkInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(c.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (c ChatBoostLinkInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(c.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (c ChatBoostLinkInfo) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(c.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (c ChatBoostLinkInfo) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(c.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c ChatBoostLinkInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, c.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c ChatBoostLinkInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(c.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (c ChatBoostLinkInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(c.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (c ChatBoostLinkInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(c.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c ChatBoostLinkInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(c.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c ChatBoostLinkInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(c.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (c ChatBoostLinkInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(c.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (c ChatBoostLinkInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (c ChatBoostLinkInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, c.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (c ChatBoostLinkInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(c.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (c ChatBoostLinkInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (c ChatBoostLinkInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(c.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (c ChatBoostLinkInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(c.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (c ChatBoostLinkInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(c.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (c ChatBoostLinkInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(c.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c ChatBoostLinkInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(c.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c ChatBoostLinkInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(c.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c ChatBoostLinkInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(c.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (c ChatBoostLinkInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(c.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c ChatBoostLinkInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(c.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c ChatBoostLinkInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, c.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (c ChatBoostLinkInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(c.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (c ChatBoostLinkInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(c.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (c ChatBoostLinkInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(c.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (c ChatBoostLinkInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(c.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c ChatBoostLinkInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, c.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (c ChatBoostLinkInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (c ChatBoostLinkInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (c ChatBoostLinkInfo) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(c.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (c ChatBoostLinkInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(c.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (c ChatBoostLinkInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(c.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (c ChatBoostLinkInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(c.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (c ChatBoostLinkInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(c.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c ChatBoostLinkInfo) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(c.ChatId, forumTopicId, draftId, text)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (c ChatBoostLinkInfo) SetBusinessAccountProfilePhoto(client *Client, businessConnectionId string, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(businessConnectionId, c.IsPublic, opts)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c ChatBoostLinkInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, c.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (c ChatBoostLinkInfo) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(c.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (c ChatBoostLinkInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(c.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (c ChatBoostLinkInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(c.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (c ChatBoostLinkInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(c.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (c ChatBoostLinkInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(c.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (c ChatBoostLinkInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(c.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (c ChatBoostLinkInfo) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(c.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (c ChatBoostLinkInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(c.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (c ChatBoostLinkInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(c.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (c ChatBoostLinkInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(c.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (c ChatBoostLinkInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(c.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (c ChatBoostLinkInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(c.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatBoostLinkInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(c.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (c ChatBoostLinkInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(c.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (c ChatBoostLinkInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(c.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (c ChatBoostLinkInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(c.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (c ChatBoostLinkInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(c.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (c ChatBoostLinkInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(c.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (c ChatBoostLinkInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(c.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (c ChatBoostLinkInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(c.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (c ChatBoostLinkInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(c.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (c ChatBoostLinkInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(c.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (c ChatBoostLinkInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(c.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (c ChatBoostLinkInfo) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(c.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (c ChatBoostLinkInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(c.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (c ChatBoostLinkInfo) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(c.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatBoostLinkInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(c.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c ChatBoostLinkInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(c.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c ChatBoostLinkInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(c.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c ChatBoostLinkInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(c.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (c ChatBoostLinkInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(c.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (c ChatBoostLinkInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(c.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c ChatBoostLinkInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(c.ChatId, messageId, optionIds)
}

// SetProfilePhoto is a helper method for Client.SetProfilePhoto
func (c ChatBoostLinkInfo) SetProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetProfilePhoto(photo, c.IsPublic)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatBoostLinkInfo) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(c.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (c ChatBoostLinkInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(c.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (c ChatBoostLinkInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(c.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c ChatBoostLinkInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(c.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (c ChatBoostLinkInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(c.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c ChatBoostLinkInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, c.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c ChatBoostLinkInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(c.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c ChatBoostLinkInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(c.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (c ChatBoostLinkInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(c.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (c ChatBoostLinkInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(c.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (c ChatBoostLinkInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(c.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c ChatBoostLinkInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(c.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (c ChatBoostLinkInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(c.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (c ChatBoostLinkInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, c.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (c ChatBoostLinkInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(c.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (c ChatBoostLinkInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(c.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (c ChatBoostLinkInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(c.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (c ChatBoostLinkInfo) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(c.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c ChatBoostLinkInfo) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(c.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c ChatBoostLinkInfo) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(c.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatBoostLinkInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(c.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c ChatBoostLinkInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(c.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (c ChatBoostLinkInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(c.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (c ChatBoostLinkInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(c.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (c ChatBoostLinkInfo) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(c.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (c ChatBoostLinkInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(c.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (c ChatBoostLinkInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(c.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (c ChatBoostLinkInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(c.ChatId, messageIds, forceRead, opts)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatBoostSlot) CreateChatInviteLink(client *Client, chatId int64, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, name, c.ExpirationDate, memberLimit, createsJoinRequest)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatBoostSlot) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, c.StartDate, isRtmpStream)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatBoostSlot) EditChatInviteLink(client *Client, chatId int64, inviteLink string, name string, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, name, c.ExpirationDate, memberLimit, createsJoinRequest)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatBoostSourceGiftCode) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatBoostSourceGiftCode) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatBoostSourceGiftCode) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatBoostSourceGiftCode) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatBoostSourceGiftCode) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatBoostSourceGiftCode) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatBoostSourceGiftCode) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatBoostSourceGiftCode) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatBoostSourceGiftCode) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatBoostSourceGiftCode) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatBoostSourceGiftCode) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatBoostSourceGiftCode) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatBoostSourceGiftCode) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatBoostSourceGiftCode) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatBoostSourceGiftCode) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatBoostSourceGiftCode) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatBoostSourceGiftCode) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatBoostSourceGiftCode) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatBoostSourceGiftCode) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatBoostSourceGiftCode) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatBoostSourceGiftCode) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatBoostSourceGiftCode) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatBoostSourceGiftCode) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatBoostSourceGiftCode) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatBoostSourceGiftCode) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatBoostSourceGiftCode) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatBoostSourceGiftCode) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatBoostSourceGiftCode) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatBoostSourceGiftCode) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatBoostSourceGiftCode) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatBoostSourceGiftCode) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatBoostSourceGiftCode) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatBoostSourceGiftCode) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatBoostSourceGiftCode) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatBoostSourceGiftCode) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatBoostSourceGiftCode) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatBoostSourceGiftCode) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatBoostSourceGiftCode) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatBoostSourceGiftCode) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatBoostSourceGiftCode) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatBoostSourceGiftCode) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatBoostSourceGiftCode) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatBoostSourceGiftCode) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatBoostSourceGiveaway) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatBoostSourceGiveaway) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (c ChatBoostSourceGiveaway) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, c.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c ChatBoostSourceGiveaway) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, c.StarCount, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatBoostSourceGiveaway) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatBoostSourceGiveaway) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (c ChatBoostSourceGiveaway) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, c.StarCount)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatBoostSourceGiveaway) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatBoostSourceGiveaway) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatBoostSourceGiveaway) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatBoostSourceGiveaway) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatBoostSourceGiveaway) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (c ChatBoostSourceGiveaway) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, c.StarCount)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatBoostSourceGiveaway) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatBoostSourceGiveaway) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatBoostSourceGiveaway) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatBoostSourceGiveaway) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatBoostSourceGiveaway) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatBoostSourceGiveaway) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatBoostSourceGiveaway) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (c ChatBoostSourceGiveaway) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, c.StarCount, password)
}

// GetUser is a helper method for Client.GetUser
func (c ChatBoostSourceGiveaway) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatBoostSourceGiveaway) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatBoostSourceGiveaway) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatBoostSourceGiveaway) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatBoostSourceGiveaway) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatBoostSourceGiveaway) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatBoostSourceGiveaway) GiftPremiumWithStars(client *Client, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, c.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (c ChatBoostSourceGiveaway) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, c.StarCount)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatBoostSourceGiveaway) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (c ChatBoostSourceGiveaway) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, c.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatBoostSourceGiveaway) PlaceGiftAuctionBid(client *Client, giftId string, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, c.StarCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatBoostSourceGiveaway) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatBoostSourceGiveaway) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatBoostSourceGiveaway) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatBoostSourceGiveaway) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (c ChatBoostSourceGiveaway) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, c.StarCount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatBoostSourceGiveaway) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatBoostSourceGiveaway) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatBoostSourceGiveaway) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatBoostSourceGiveaway) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatBoostSourceGiveaway) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatBoostSourceGiveaway) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatBoostSourceGiveaway) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatBoostSourceGiveaway) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatBoostSourceGiveaway) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatBoostSourceGiveaway) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatBoostSourceGiveaway) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatBoostSourceGiveaway) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (c ChatBoostSourceGiveaway) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, c.StarCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatBoostSourceGiveaway) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (c ChatBoostSourceGiveaway) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, c.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (c ChatBoostSourceGiveaway) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, c.StarCount)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatBoostSourceGiveaway) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatBoostSourcePremium) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatBoostSourcePremium) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatBoostSourcePremium) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatBoostSourcePremium) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatBoostSourcePremium) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatBoostSourcePremium) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatBoostSourcePremium) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatBoostSourcePremium) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatBoostSourcePremium) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatBoostSourcePremium) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatBoostSourcePremium) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatBoostSourcePremium) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatBoostSourcePremium) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatBoostSourcePremium) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatBoostSourcePremium) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatBoostSourcePremium) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatBoostSourcePremium) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatBoostSourcePremium) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatBoostSourcePremium) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatBoostSourcePremium) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatBoostSourcePremium) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatBoostSourcePremium) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatBoostSourcePremium) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatBoostSourcePremium) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatBoostSourcePremium) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatBoostSourcePremium) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatBoostSourcePremium) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatBoostSourcePremium) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatBoostSourcePremium) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatBoostSourcePremium) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatBoostSourcePremium) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatBoostSourcePremium) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatBoostSourcePremium) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatBoostSourcePremium) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatBoostSourcePremium) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatBoostSourcePremium) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatBoostSourcePremium) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatBoostSourcePremium) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatBoostSourcePremium) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatBoostSourcePremium) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatBoostSourcePremium) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatBoostSourcePremium) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatBoostSourcePremium) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// GetChatBoostLevelFeatures is a helper method for Client.GetChatBoostLevelFeatures
func (c ChatBoostStatus) GetChatBoostLevelFeatures(client *Client, isChannel bool) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(isChannel, c.Level)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatEvent) BanChatMember(client *Client, chatId int64, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(chatId, c.MemberId, bannedUntilDate, revokeMessages)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c ChatEvent) GetChatMember(client *Client, chatId int64) (*ChatMember, error) {
	return client.GetChatMember(chatId, c.MemberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatEvent) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, c.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatEvent) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, c.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (c ChatEvent) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, c.Date)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatEvent) SetChatMemberStatus(client *Client, chatId int64, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(chatId, c.MemberId, status)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (c ChatEventAutomaticTranslationToggled) ToggleSupergroupHasAutomaticTranslation(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(supergroupId, c.HasAutomaticTranslation)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (c ChatEventHasAggressiveAntiSpamEnabledToggled) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(supergroupId, c.HasAggressiveAntiSpamEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c ChatEventHasProtectedContentToggled) ToggleChatHasProtectedContent(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(chatId, c.HasProtectedContent)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (c ChatEventIsAllHistoryAvailableToggled) ToggleSupergroupIsAllHistoryAvailable(client *Client, supergroupId int64) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(supergroupId, c.IsAllHistoryAvailable)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (c ChatEventIsForumToggled) CreateNewSupergroupChat(client *Client, title string, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, c.IsForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (c ChatEventIsForumToggled) ToggleSupergroupIsForum(client *Client, supergroupId int64, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(supergroupId, c.IsForum, hasForumTabs)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatEventMemberInvited) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatEventMemberInvited) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatEventMemberInvited) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatEventMemberInvited) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatEventMemberInvited) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatEventMemberInvited) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatEventMemberInvited) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatEventMemberInvited) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatEventMemberInvited) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatEventMemberInvited) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatEventMemberInvited) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatEventMemberInvited) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatEventMemberInvited) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatEventMemberInvited) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatEventMemberInvited) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatEventMemberInvited) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatEventMemberInvited) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatEventMemberInvited) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatEventMemberInvited) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatEventMemberInvited) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatEventMemberInvited) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatEventMemberInvited) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatEventMemberInvited) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatEventMemberInvited) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatEventMemberInvited) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatEventMemberInvited) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatEventMemberInvited) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatEventMemberInvited) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatEventMemberInvited) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatEventMemberInvited) SetChatMemberStatus(client *Client, chatId int64, memberId *MessageSender) (*Ok, error) {
	return client.SetChatMemberStatus(chatId, memberId, c.Status)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatEventMemberInvited) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatEventMemberInvited) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatEventMemberInvited) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatEventMemberInvited) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatEventMemberInvited) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatEventMemberInvited) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatEventMemberInvited) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatEventMemberInvited) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatEventMemberInvited) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatEventMemberInvited) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatEventMemberInvited) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatEventMemberInvited) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatEventMemberInvited) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatEventMemberInvited) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatEventMemberPromoted) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatEventMemberPromoted) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatEventMemberPromoted) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatEventMemberPromoted) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatEventMemberPromoted) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatEventMemberPromoted) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatEventMemberPromoted) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatEventMemberPromoted) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatEventMemberPromoted) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatEventMemberPromoted) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatEventMemberPromoted) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatEventMemberPromoted) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatEventMemberPromoted) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatEventMemberPromoted) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatEventMemberPromoted) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatEventMemberPromoted) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatEventMemberPromoted) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatEventMemberPromoted) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatEventMemberPromoted) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatEventMemberPromoted) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatEventMemberPromoted) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatEventMemberPromoted) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatEventMemberPromoted) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatEventMemberPromoted) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatEventMemberPromoted) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatEventMemberPromoted) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatEventMemberPromoted) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatEventMemberPromoted) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatEventMemberPromoted) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatEventMemberPromoted) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatEventMemberPromoted) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatEventMemberPromoted) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatEventMemberPromoted) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatEventMemberPromoted) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatEventMemberPromoted) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatEventMemberPromoted) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatEventMemberPromoted) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatEventMemberPromoted) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatEventMemberPromoted) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatEventMemberPromoted) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatEventMemberPromoted) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatEventMemberPromoted) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatEventMemberPromoted) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatEventMemberRestricted) BanChatMember(client *Client, chatId int64, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(chatId, c.MemberId, bannedUntilDate, revokeMessages)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c ChatEventMemberRestricted) GetChatMember(client *Client, chatId int64) (*ChatMember, error) {
	return client.GetChatMember(chatId, c.MemberId)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatEventMemberRestricted) SetChatMemberStatus(client *Client, chatId int64, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(chatId, c.MemberId, status)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatEventMemberSubscriptionExtended) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatEventMemberSubscriptionExtended) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatEventMemberSubscriptionExtended) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatEventMemberSubscriptionExtended) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatEventMemberSubscriptionExtended) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatEventMemberSubscriptionExtended) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatEventMemberSubscriptionExtended) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatEventMemberSubscriptionExtended) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatEventMemberSubscriptionExtended) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatEventMemberSubscriptionExtended) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatEventMemberSubscriptionExtended) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatEventMemberSubscriptionExtended) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatEventMemberSubscriptionExtended) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatEventMemberSubscriptionExtended) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatEventMemberSubscriptionExtended) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatEventMemberSubscriptionExtended) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatEventMemberSubscriptionExtended) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatEventMemberSubscriptionExtended) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatEventMemberSubscriptionExtended) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatEventMemberSubscriptionExtended) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatEventMemberSubscriptionExtended) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatEventMemberSubscriptionExtended) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatEventMemberSubscriptionExtended) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatEventMemberSubscriptionExtended) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatEventMemberSubscriptionExtended) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatEventMemberSubscriptionExtended) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatEventMemberSubscriptionExtended) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatEventMemberSubscriptionExtended) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatEventMemberSubscriptionExtended) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatEventMemberSubscriptionExtended) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatEventMemberSubscriptionExtended) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatEventMemberSubscriptionExtended) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatEventMemberSubscriptionExtended) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatEventMemberSubscriptionExtended) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatEventMemberSubscriptionExtended) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatEventMemberSubscriptionExtended) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatEventMemberSubscriptionExtended) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatEventMemberSubscriptionExtended) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatEventMemberSubscriptionExtended) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatEventMemberSubscriptionExtended) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatEventMemberSubscriptionExtended) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatEventMemberSubscriptionExtended) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatEventMemberSubscriptionExtended) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (c ChatEventShowMessageSenderToggled) ToggleSupergroupSignMessages(client *Client, supergroupId int64, signMessages bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(supergroupId, signMessages, c.ShowMessageSender)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (c ChatEventSignMessagesToggled) ToggleSupergroupSignMessages(client *Client, supergroupId int64, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(supergroupId, c.SignMessages, showMessageSender)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (c ChatEventVideoChatCreated) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(c.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (c ChatEventVideoChatCreated) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(c.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (c ChatEventVideoChatCreated) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(c.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c ChatEventVideoChatCreated) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(c.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (c ChatEventVideoChatCreated) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(c.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (c ChatEventVideoChatCreated) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(c.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (c ChatEventVideoChatCreated) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(c.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (c ChatEventVideoChatCreated) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(c.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (c ChatEventVideoChatCreated) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(c.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (c ChatEventVideoChatCreated) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(c.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (c ChatEventVideoChatCreated) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(c.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (c ChatEventVideoChatCreated) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(c.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (c ChatEventVideoChatCreated) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(c.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (c ChatEventVideoChatCreated) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(c.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (c ChatEventVideoChatCreated) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(c.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (c ChatEventVideoChatCreated) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(c.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (c ChatEventVideoChatCreated) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(c.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatEventVideoChatCreated) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(c.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (c ChatEventVideoChatCreated) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(c.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (c ChatEventVideoChatCreated) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(c.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (c ChatEventVideoChatCreated) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(c.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (c ChatEventVideoChatCreated) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(c.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (c ChatEventVideoChatCreated) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(c.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (c ChatEventVideoChatCreated) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(c.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (c ChatEventVideoChatCreated) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(c.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (c ChatEventVideoChatCreated) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(c.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (c ChatEventVideoChatCreated) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(c.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (c ChatEventVideoChatCreated) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(c.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (c ChatEventVideoChatCreated) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(c.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (c ChatEventVideoChatCreated) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(c.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (c ChatEventVideoChatCreated) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(c.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (c ChatEventVideoChatCreated) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(c.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (c ChatEventVideoChatCreated) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(c.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (c ChatEventVideoChatCreated) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(c.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (c ChatEventVideoChatCreated) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(c.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (c ChatEventVideoChatCreated) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(c.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (c ChatEventVideoChatCreated) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(c.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (c ChatEventVideoChatCreated) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(c.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (c ChatEventVideoChatCreated) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(c.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (c ChatEventVideoChatCreated) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(c.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (c ChatEventVideoChatCreated) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(c.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (c ChatEventVideoChatCreated) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(c.GroupCallId, muteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (c ChatEventVideoChatEnded) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(c.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (c ChatEventVideoChatEnded) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(c.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (c ChatEventVideoChatEnded) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(c.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c ChatEventVideoChatEnded) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(c.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (c ChatEventVideoChatEnded) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(c.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (c ChatEventVideoChatEnded) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(c.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (c ChatEventVideoChatEnded) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(c.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (c ChatEventVideoChatEnded) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(c.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (c ChatEventVideoChatEnded) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(c.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (c ChatEventVideoChatEnded) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(c.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (c ChatEventVideoChatEnded) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(c.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (c ChatEventVideoChatEnded) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(c.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (c ChatEventVideoChatEnded) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(c.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (c ChatEventVideoChatEnded) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(c.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (c ChatEventVideoChatEnded) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(c.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (c ChatEventVideoChatEnded) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(c.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (c ChatEventVideoChatEnded) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(c.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatEventVideoChatEnded) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(c.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (c ChatEventVideoChatEnded) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(c.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (c ChatEventVideoChatEnded) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(c.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (c ChatEventVideoChatEnded) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(c.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (c ChatEventVideoChatEnded) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(c.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (c ChatEventVideoChatEnded) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(c.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (c ChatEventVideoChatEnded) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(c.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (c ChatEventVideoChatEnded) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(c.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (c ChatEventVideoChatEnded) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(c.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (c ChatEventVideoChatEnded) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(c.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (c ChatEventVideoChatEnded) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(c.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (c ChatEventVideoChatEnded) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(c.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (c ChatEventVideoChatEnded) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(c.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (c ChatEventVideoChatEnded) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(c.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (c ChatEventVideoChatEnded) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(c.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (c ChatEventVideoChatEnded) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(c.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (c ChatEventVideoChatEnded) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(c.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (c ChatEventVideoChatEnded) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(c.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (c ChatEventVideoChatEnded) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(c.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (c ChatEventVideoChatEnded) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(c.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (c ChatEventVideoChatEnded) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(c.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (c ChatEventVideoChatEnded) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(c.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (c ChatEventVideoChatEnded) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(c.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (c ChatEventVideoChatEnded) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(c.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (c ChatEventVideoChatEnded) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(c.GroupCallId, muteNewParticipants)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (c ChatEventVideoChatMuteNewParticipantsToggled) ToggleVideoChatMuteNewParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(groupCallId, c.MuteNewParticipants)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c ChatEventVideoChatParticipantIsMutedToggled) DecryptGroupCallData(client *Client, groupCallId int32, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, c.ParticipantId, data, opts)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (c ChatEventVideoChatParticipantIsMutedToggled) SetGroupCallParticipantVolumeLevel(client *Client, groupCallId int32, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(groupCallId, c.ParticipantId, volumeLevel)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (c ChatEventVideoChatParticipantIsMutedToggled) ToggleGroupCallParticipantIsHandRaised(client *Client, groupCallId int32, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(groupCallId, c.ParticipantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (c ChatEventVideoChatParticipantIsMutedToggled) ToggleGroupCallParticipantIsMuted(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(groupCallId, c.ParticipantId, c.IsMuted)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c ChatEventVideoChatParticipantVolumeLevelChanged) DecryptGroupCallData(client *Client, groupCallId int32, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, c.ParticipantId, data, opts)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (c ChatEventVideoChatParticipantVolumeLevelChanged) SetGroupCallParticipantVolumeLevel(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(groupCallId, c.ParticipantId, c.VolumeLevel)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (c ChatEventVideoChatParticipantVolumeLevelChanged) ToggleGroupCallParticipantIsHandRaised(client *Client, groupCallId int32, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(groupCallId, c.ParticipantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (c ChatEventVideoChatParticipantVolumeLevelChanged) ToggleGroupCallParticipantIsMuted(client *Client, groupCallId int32, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(groupCallId, c.ParticipantId, isMuted)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatFolderIcon) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, c.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (c ChatFolderIcon) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(c.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (c ChatFolderIcon) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(c.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c ChatFolderIcon) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, c.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatFolderIcon) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatFolderIcon) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, c.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatFolderIcon) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, c.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (c ChatFolderIcon) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, c.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatFolderIcon) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, c.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (c ChatFolderIcon) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, c.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (c ChatFolderIcon) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(c.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c ChatFolderIcon) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, c.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatFolderIcon) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatFolderIcon) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, c.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatFolderIcon) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, c.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (c ChatFolderIcon) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(c.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (c ChatFolderIcon) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(c.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (c ChatFolderIcon) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(c.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (c ChatFolderIcon) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(c.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatFolderIcon) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, c.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (c ChatFolderIcon) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(c.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (c ChatFolderIcon) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(c.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (c ChatFolderIcon) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, c.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c ChatFolderIcon) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(c.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (c ChatFolderIcon) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, c.Name)
}

// SetOption is a helper method for Client.SetOption
func (c ChatFolderIcon) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(c.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (c ChatFolderIcon) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, c.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatFolderIcon) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, c.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatFolderIcon) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(c.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatFolderIcon) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, c.Name)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (c ChatFolderInviteLink) AddChatFolderByInviteLink(client *Client) (*Ok, error) {
	return client.AddChatFolderByInviteLink(c.InviteLink, c.ChatIds)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatFolderInviteLink) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, c.Name, sticker)
}

// Check is a helper method for Client.CheckChatFolderInviteLink
func (c ChatFolderInviteLink) Check(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(c.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (c ChatFolderInviteLink) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(c.InviteLink)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (c ChatFolderInviteLink) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(c.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (c ChatFolderInviteLink) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(c.Name)
}

// Create is a helper method for Client.CreateChatFolderInviteLink
func (c ChatFolderInviteLink) Create(client *Client, chatFolderId int32) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, c.Name, c.ChatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatFolderInviteLink) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatFolderInviteLink) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, c.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatFolderInviteLink) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, c.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (c ChatFolderInviteLink) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, c.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatFolderInviteLink) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, c.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (c ChatFolderInviteLink) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, c.Name, storyIds)
}

// Delete is a helper method for Client.DeleteChatFolderInviteLink
func (c ChatFolderInviteLink) Delete(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, c.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c ChatFolderInviteLink) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, c.InviteLink)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (c ChatFolderInviteLink) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(c.Name)
}

// DiscardCall is a helper method for Client.DiscardCall
func (c ChatFolderInviteLink) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, c.InviteLink, duration, isVideo, connectionId)
}

// Edit is a helper method for Client.EditChatFolderInviteLink
func (c ChatFolderInviteLink) Edit(client *Client, chatFolderId int32) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, c.InviteLink, c.Name, c.ChatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatFolderInviteLink) EditChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, c.InviteLink, c.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatFolderInviteLink) EditChatSubscriptionInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, c.InviteLink, c.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatFolderInviteLink) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, c.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (c ChatFolderInviteLink) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(c.Name, typeField)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c ChatFolderInviteLink) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, c.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c ChatFolderInviteLink) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, c.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c ChatFolderInviteLink) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, c.InviteLink, query, limit, opts)
}

// GetOption is a helper method for Client.GetOption
func (c ChatFolderInviteLink) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(c.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (c ChatFolderInviteLink) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(c.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (c ChatFolderInviteLink) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(c.Name)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (c ChatFolderInviteLink) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(c.InviteLink)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (c ChatFolderInviteLink) OptimizeStorage(client *Client, size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, count, immunityDelay, fileTypes, c.ChatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c ChatFolderInviteLink) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, c.InviteLink, approve)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatFolderInviteLink) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, c.Name, oldSticker, newSticker)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c ChatFolderInviteLink) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, c.InviteLink)
}

// SearchBackground is a helper method for Client.SearchBackground
func (c ChatFolderInviteLink) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(c.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (c ChatFolderInviteLink) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(c.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (c ChatFolderInviteLink) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, c.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c ChatFolderInviteLink) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(c.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (c ChatFolderInviteLink) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, c.Name)
}

// SetOption is a helper method for Client.SetOption
func (c ChatFolderInviteLink) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(c.Name, opts)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (c ChatFolderInviteLink) SetPinnedChats(client *Client, chatList *ChatList) (*Ok, error) {
	return client.SetPinnedChats(chatList, c.ChatIds)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (c ChatFolderInviteLink) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, c.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatFolderInviteLink) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, c.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatFolderInviteLink) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(c.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatFolderInviteLink) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, c.Name)
}

// ProcessChatFolderNewChats is a helper method for Client.ProcessChatFolderNewChats
func (c ChatFolderInviteLinkInfo) ProcessChatFolderNewChats(client *Client, chatFolderId int32) (*Ok, error) {
	return client.ProcessChatFolderNewChats(chatFolderId, c.AddedChatIds)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (c ChatFolderName) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(c.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (c ChatFolderName) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(c.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatFolderName) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, c.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (c ChatFolderName) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(c.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatFolderName) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, c.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (c ChatFolderName) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(c.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (c ChatFolderName) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, c.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (c ChatFolderName) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, c.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c ChatFolderName) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, c.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (c ChatFolderName) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(c.Text, toLanguageCode)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (c ChatInviteLink) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(c.InviteLink, chatIds)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatInviteLink) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, c.Name, sticker)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (c ChatInviteLink) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(c.InviteLink)
}

// Check is a helper method for Client.CheckChatInviteLink
func (c ChatInviteLink) Check(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(c.InviteLink)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (c ChatInviteLink) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(c.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (c ChatInviteLink) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(c.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c ChatInviteLink) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, c.Name, chatIds)
}

// Create is a helper method for Client.CreateChatInviteLink
func (c ChatInviteLink) Create(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, c.Name, c.ExpirationDate, c.MemberLimit, c.CreatesJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatInviteLink) CreateChatSubscriptionInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, c.Name, c.SubscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatInviteLink) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, c.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (c ChatInviteLink) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, c.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatInviteLink) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, c.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (c ChatInviteLink) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, c.Name, storyIds)
}

// DeleteAllRevokeds is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c ChatInviteLink) DeleteAllRevokeds(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(chatId, c.CreatorUserId)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (c ChatInviteLink) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, c.InviteLink)
}

// DeleteRevoked is a helper method for Client.DeleteRevokedChatInviteLink
func (c ChatInviteLink) DeleteRevoked(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, c.InviteLink)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (c ChatInviteLink) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(c.Name)
}

// DiscardCall is a helper method for Client.DiscardCall
func (c ChatInviteLink) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, c.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c ChatInviteLink) EditChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, c.InviteLink, c.Name, chatIds)
}

// Edit is a helper method for Client.EditChatInviteLink
func (c ChatInviteLink) Edit(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, c.InviteLink, c.Name, c.ExpirationDate, c.MemberLimit, c.CreatesJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatInviteLink) EditChatSubscriptionInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, c.InviteLink, c.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatInviteLink) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, c.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (c ChatInviteLink) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(c.Name, typeField)
}

// Get is a helper method for Client.GetChatInviteLink
func (c ChatInviteLink) Get(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, c.InviteLink)
}

// GetMembers is a helper method for Client.GetChatInviteLinkMembers
func (c ChatInviteLink) GetMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, c.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// Gets is a helper method for Client.GetChatInviteLinks
func (c ChatInviteLink) Gets(client *Client, chatId int64, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(chatId, c.CreatorUserId, c.IsRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c ChatInviteLink) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, c.InviteLink, query, limit, opts)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatInviteLink) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, c.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatInviteLink) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, c.Date)
}

// GetOption is a helper method for Client.GetOption
func (c ChatInviteLink) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(c.Name)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (c ChatInviteLink) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, c.Date)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (c ChatInviteLink) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(c.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (c ChatInviteLink) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(c.Name)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (c ChatInviteLink) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(c.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c ChatInviteLink) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, c.InviteLink, approve)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatInviteLink) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, c.Name, oldSticker, newSticker)
}

// Revoke is a helper method for Client.RevokeChatInviteLink
func (c ChatInviteLink) Revoke(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, c.InviteLink)
}

// SearchBackground is a helper method for Client.SearchBackground
func (c ChatInviteLink) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(c.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (c ChatInviteLink) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(c.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (c ChatInviteLink) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, c.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c ChatInviteLink) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(c.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (c ChatInviteLink) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, c.Name)
}

// SetOption is a helper method for Client.SetOption
func (c ChatInviteLink) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(c.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (c ChatInviteLink) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, c.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatInviteLink) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, c.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatInviteLink) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(c.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatInviteLink) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, c.Name)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatInviteLinkCount) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatInviteLinkCount) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatInviteLinkCount) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatInviteLinkCount) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatInviteLinkCount) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatInviteLinkCount) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatInviteLinkCount) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatInviteLinkCount) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatInviteLinkCount) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatInviteLinkCount) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatInviteLinkCount) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatInviteLinkCount) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatInviteLinkCount) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatInviteLinkCount) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatInviteLinkCount) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatInviteLinkCount) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatInviteLinkCount) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatInviteLinkCount) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatInviteLinkCount) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatInviteLinkCount) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatInviteLinkCount) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatInviteLinkCount) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatInviteLinkCount) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatInviteLinkCount) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatInviteLinkCount) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatInviteLinkCount) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatInviteLinkCount) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatInviteLinkCount) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatInviteLinkCount) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatInviteLinkCount) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatInviteLinkCount) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatInviteLinkCount) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatInviteLinkCount) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatInviteLinkCount) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatInviteLinkCount) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatInviteLinkCount) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatInviteLinkCount) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatInviteLinkCount) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatInviteLinkCount) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatInviteLinkCount) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatInviteLinkCount) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatInviteLinkCount) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatInviteLinkCount) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatInviteLinkInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(c.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (c ChatInviteLinkInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(c.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (c ChatInviteLinkInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(c.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c ChatInviteLinkInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(c.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c ChatInviteLinkInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, c.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (c ChatInviteLinkInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(c.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c ChatInviteLinkInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(c.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c ChatInviteLinkInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(c.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c ChatInviteLinkInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(c.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (c ChatInviteLinkInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(c.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (c ChatInviteLinkInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c ChatInviteLinkInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(c.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatInviteLinkInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(c.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (c ChatInviteLinkInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(c.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (c ChatInviteLinkInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(c.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c ChatInviteLinkInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(c.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c ChatInviteLinkInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(c.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c ChatInviteLinkInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(c.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (c ChatInviteLinkInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(c.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c ChatInviteLinkInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(c.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatInviteLinkInfo) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(c.ChatId, name, expirationDate, memberLimit, c.CreatesJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatInviteLinkInfo) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(c.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatInviteLinkInfo) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(c.ChatId, name, isNameImplicit, icon)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (c ChatInviteLinkInfo) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, c.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatInviteLinkInfo) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, c.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (c ChatInviteLinkInfo) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(c.Title, isForum, isChannel, c.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatInviteLinkInfo) CreateVideoChat(client *Client, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(c.ChatId, c.Title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c ChatInviteLinkInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(c.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c ChatInviteLinkInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(c.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c ChatInviteLinkInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(c.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (c ChatInviteLinkInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(c.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (c ChatInviteLinkInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(c.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (c ChatInviteLinkInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(c.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (c ChatInviteLinkInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(c.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (c ChatInviteLinkInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(c.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c ChatInviteLinkInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(c.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (c ChatInviteLinkInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(c.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (c ChatInviteLinkInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(c.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (c ChatInviteLinkInfo) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(c.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (c ChatInviteLinkInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(c.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c ChatInviteLinkInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(c.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (c ChatInviteLinkInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(c.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c ChatInviteLinkInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c ChatInviteLinkInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, c.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c ChatInviteLinkInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c ChatInviteLinkInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c ChatInviteLinkInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, c.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c ChatInviteLinkInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatInviteLinkInfo) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(c.ChatId, inviteLink, name, expirationDate, memberLimit, c.CreatesJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatInviteLinkInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(c.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatInviteLinkInfo) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(c.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c ChatInviteLinkInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c ChatInviteLinkInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(c.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c ChatInviteLinkInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c ChatInviteLinkInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(c.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c ChatInviteLinkInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(c.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c ChatInviteLinkInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(c.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c ChatInviteLinkInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(c.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (c ChatInviteLinkInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(c.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (c ChatInviteLinkInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, c.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c ChatInviteLinkInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(c.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c ChatInviteLinkInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(c.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (c ChatInviteLinkInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(c.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (c ChatInviteLinkInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(c.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (c ChatInviteLinkInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(c.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (c ChatInviteLinkInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(c.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (c ChatInviteLinkInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(c.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (c ChatInviteLinkInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(c.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (c ChatInviteLinkInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(c.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (c ChatInviteLinkInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(c.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (c ChatInviteLinkInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(c.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (c ChatInviteLinkInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(c.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (c ChatInviteLinkInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(c.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c ChatInviteLinkInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(c.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (c ChatInviteLinkInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(c.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c ChatInviteLinkInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(c.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (c ChatInviteLinkInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(c.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c ChatInviteLinkInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(c.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (c ChatInviteLinkInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(c.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c ChatInviteLinkInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(c.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatInviteLinkInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(c.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (c ChatInviteLinkInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(c.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (c ChatInviteLinkInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(c.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (c ChatInviteLinkInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(c.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (c ChatInviteLinkInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(c.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (c ChatInviteLinkInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(c.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (c ChatInviteLinkInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(c.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (c ChatInviteLinkInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(c.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c ChatInviteLinkInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(c.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (c ChatInviteLinkInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(c.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (c ChatInviteLinkInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(c.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (c ChatInviteLinkInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(c.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (c ChatInviteLinkInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(c.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (c ChatInviteLinkInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(c.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (c ChatInviteLinkInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(c.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (c ChatInviteLinkInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(c.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (c ChatInviteLinkInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(c.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (c ChatInviteLinkInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(c.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatInviteLinkInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(c.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (c ChatInviteLinkInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(c.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (c ChatInviteLinkInfo) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(c.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (c ChatInviteLinkInfo) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(c.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (c ChatInviteLinkInfo) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(c.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (c ChatInviteLinkInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(c.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatInviteLinkInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(c.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c ChatInviteLinkInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(c.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c ChatInviteLinkInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, c.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (c ChatInviteLinkInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(c.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c ChatInviteLinkInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(c.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c ChatInviteLinkInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(c.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c ChatInviteLinkInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(c.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c ChatInviteLinkInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, c.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (c ChatInviteLinkInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(c.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c ChatInviteLinkInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(c.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c ChatInviteLinkInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(c.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c ChatInviteLinkInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(c.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c ChatInviteLinkInfo) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(c.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (c ChatInviteLinkInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(c.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c ChatInviteLinkInfo) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(c.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c ChatInviteLinkInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(c.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c ChatInviteLinkInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(c.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c ChatInviteLinkInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(c.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c ChatInviteLinkInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(c.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (c ChatInviteLinkInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(c.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c ChatInviteLinkInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(c.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c ChatInviteLinkInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(c.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c ChatInviteLinkInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(c.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c ChatInviteLinkInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(c.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c ChatInviteLinkInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(c.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c ChatInviteLinkInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(c.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c ChatInviteLinkInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(c.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (c ChatInviteLinkInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(c.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (c ChatInviteLinkInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, c.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (c ChatInviteLinkInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(c.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c ChatInviteLinkInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(c.ChatId, storyId, isDark)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (c ChatInviteLinkInfo) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(c.Title)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatInviteLinkInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(c.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (c ChatInviteLinkInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(c.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (c ChatInviteLinkInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(c.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c ChatInviteLinkInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(c.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c ChatInviteLinkInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(c.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (c ChatInviteLinkInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(c.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (c ChatInviteLinkInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(c.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (c ChatInviteLinkInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(c.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (c ChatInviteLinkInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(c.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c ChatInviteLinkInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(c.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (c ChatInviteLinkInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(c.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (c ChatInviteLinkInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(c.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c ChatInviteLinkInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(c.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c ChatInviteLinkInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(c.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (c ChatInviteLinkInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(c.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (c ChatInviteLinkInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(c.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatInviteLinkInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(c.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c ChatInviteLinkInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(c.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c ChatInviteLinkInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(c.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (c ChatInviteLinkInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(c.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (c ChatInviteLinkInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(c.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (c ChatInviteLinkInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(c.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (c ChatInviteLinkInfo) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(c.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (c ChatInviteLinkInfo) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(c.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c ChatInviteLinkInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, c.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c ChatInviteLinkInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(c.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (c ChatInviteLinkInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(c.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (c ChatInviteLinkInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(c.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c ChatInviteLinkInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(c.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c ChatInviteLinkInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(c.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (c ChatInviteLinkInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(c.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (c ChatInviteLinkInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (c ChatInviteLinkInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, c.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (c ChatInviteLinkInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(c.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (c ChatInviteLinkInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (c ChatInviteLinkInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(c.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (c ChatInviteLinkInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(c.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (c ChatInviteLinkInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(c.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (c ChatInviteLinkInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(c.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c ChatInviteLinkInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(c.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c ChatInviteLinkInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(c.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c ChatInviteLinkInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(c.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (c ChatInviteLinkInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(c.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c ChatInviteLinkInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(c.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c ChatInviteLinkInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, c.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (c ChatInviteLinkInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(c.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (c ChatInviteLinkInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(c.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (c ChatInviteLinkInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(c.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (c ChatInviteLinkInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(c.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c ChatInviteLinkInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, c.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (c ChatInviteLinkInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (c ChatInviteLinkInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, c.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (c ChatInviteLinkInfo) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(c.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (c ChatInviteLinkInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(c.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (c ChatInviteLinkInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(c.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (c ChatInviteLinkInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(c.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (c ChatInviteLinkInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(c.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c ChatInviteLinkInfo) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(c.ChatId, forumTopicId, draftId, text)
}

// SetAccentColor is a helper method for Client.SetAccentColor
func (c ChatInviteLinkInfo) SetAccentColor(client *Client, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetAccentColor(c.AccentColorId, backgroundCustomEmojiId)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (c ChatInviteLinkInfo) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, c.Description)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (c ChatInviteLinkInfo) SetBusinessAccountProfilePhoto(client *Client, businessConnectionId string, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(businessConnectionId, c.IsPublic, opts)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c ChatInviteLinkInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, c.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (c ChatInviteLinkInfo) SetChatAccentColor(client *Client, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(c.ChatId, c.AccentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (c ChatInviteLinkInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(c.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (c ChatInviteLinkInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(c.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (c ChatInviteLinkInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(c.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (c ChatInviteLinkInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(c.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (c ChatInviteLinkInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(c.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (c ChatInviteLinkInfo) SetChatDescription(client *Client) (*Ok, error) {
	return client.SetChatDescription(c.ChatId, c.Description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (c ChatInviteLinkInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(c.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (c ChatInviteLinkInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(c.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (c ChatInviteLinkInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(c.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (c ChatInviteLinkInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(c.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (c ChatInviteLinkInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(c.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (c ChatInviteLinkInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(c.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (c ChatInviteLinkInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(c.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (c ChatInviteLinkInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(c.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (c ChatInviteLinkInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(c.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (c ChatInviteLinkInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(c.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (c ChatInviteLinkInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(c.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (c ChatInviteLinkInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(c.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (c ChatInviteLinkInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(c.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (c ChatInviteLinkInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(c.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (c ChatInviteLinkInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(c.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (c ChatInviteLinkInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(c.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (c ChatInviteLinkInfo) SetChatTitle(client *Client) (*Ok, error) {
	return client.SetChatTitle(c.ChatId, c.Title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (c ChatInviteLinkInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(c.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (c ChatInviteLinkInfo) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(c.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatInviteLinkInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(c.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c ChatInviteLinkInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(c.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c ChatInviteLinkInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(c.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c ChatInviteLinkInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(c.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (c ChatInviteLinkInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(c.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (c ChatInviteLinkInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(c.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c ChatInviteLinkInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(c.ChatId, messageId, optionIds)
}

// SetProfilePhoto is a helper method for Client.SetProfilePhoto
func (c ChatInviteLinkInfo) SetProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetProfilePhoto(photo, c.IsPublic)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatInviteLinkInfo) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, c.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c ChatInviteLinkInfo) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(c.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (c ChatInviteLinkInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(c.ChatId, defaultParticipantId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (c ChatInviteLinkInfo) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, c.Title)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (c ChatInviteLinkInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(c.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c ChatInviteLinkInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(c.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (c ChatInviteLinkInfo) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, c.Title, recordVideo, usePortraitOrientation)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (c ChatInviteLinkInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(c.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c ChatInviteLinkInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, c.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c ChatInviteLinkInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(c.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c ChatInviteLinkInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(c.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (c ChatInviteLinkInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(c.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (c ChatInviteLinkInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(c.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (c ChatInviteLinkInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(c.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c ChatInviteLinkInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(c.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (c ChatInviteLinkInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(c.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (c ChatInviteLinkInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, c.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (c ChatInviteLinkInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(c.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (c ChatInviteLinkInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(c.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (c ChatInviteLinkInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(c.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (c ChatInviteLinkInfo) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(c.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c ChatInviteLinkInfo) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(c.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c ChatInviteLinkInfo) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(c.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatInviteLinkInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(c.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c ChatInviteLinkInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(c.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (c ChatInviteLinkInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(c.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (c ChatInviteLinkInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(c.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (c ChatInviteLinkInfo) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(c.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (c ChatInviteLinkInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(c.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (c ChatInviteLinkInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(c.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (c ChatInviteLinkInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(c.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatInviteLinkMember) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatInviteLinkMember) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatInviteLinkMember) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatInviteLinkMember) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatInviteLinkMember) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatInviteLinkMember) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatInviteLinkMember) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatInviteLinkMember) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatInviteLinkMember) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatInviteLinkMember) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatInviteLinkMember) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatInviteLinkMember) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatInviteLinkMember) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatInviteLinkMember) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatInviteLinkMember) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatInviteLinkMember) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatInviteLinkMember) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatInviteLinkMember) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatInviteLinkMember) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatInviteLinkMember) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatInviteLinkMember) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatInviteLinkMember) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatInviteLinkMember) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatInviteLinkMember) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatInviteLinkMember) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatInviteLinkMember) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatInviteLinkMember) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatInviteLinkMember) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatInviteLinkMember) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatInviteLinkMember) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatInviteLinkMember) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatInviteLinkMember) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatInviteLinkMember) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatInviteLinkMember) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatInviteLinkMember) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatInviteLinkMember) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatInviteLinkMember) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatInviteLinkMember) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatInviteLinkMember) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatInviteLinkMember) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatInviteLinkMember) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatInviteLinkMember) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatInviteLinkMember) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatJoinRequest) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatJoinRequest) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatJoinRequest) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatJoinRequest) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatJoinRequest) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatJoinRequest) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatJoinRequest) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatJoinRequest) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatJoinRequest) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatJoinRequest) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatJoinRequest) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, c.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatJoinRequest) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, c.Date)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatJoinRequest) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatJoinRequest) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatJoinRequest) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatJoinRequest) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatJoinRequest) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (c ChatJoinRequest) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, c.Date)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatJoinRequest) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatJoinRequest) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatJoinRequest) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatJoinRequest) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatJoinRequest) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatJoinRequest) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatJoinRequest) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatJoinRequest) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatJoinRequest) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatJoinRequest) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// Process is a helper method for Client.ProcessChatJoinRequest
func (c ChatJoinRequest) Process(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatJoinRequest) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatJoinRequest) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatJoinRequest) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetBio is a helper method for Client.SetBio
func (c ChatJoinRequest) SetBio(client *Client) (*Ok, error) {
	return client.SetBio(c.Bio)
}

// SetBusinessAccountBio is a helper method for Client.SetBusinessAccountBio
func (c ChatJoinRequest) SetBusinessAccountBio(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountBio(businessConnectionId, c.Bio)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatJoinRequest) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatJoinRequest) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatJoinRequest) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatJoinRequest) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatJoinRequest) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatJoinRequest) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatJoinRequest) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatJoinRequest) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatJoinRequest) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatJoinRequest) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatJoinRequest) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatJoinRequest) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatJoinRequest) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatJoinRequest) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (c ChatJoinRequestsInfo) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, c.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (c ChatJoinRequestsInfo) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(c.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (c ChatJoinRequestsInfo) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, c.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (c ChatJoinRequestsInfo) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, c.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (c ChatJoinRequestsInfo) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(c.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (c ChatJoinRequestsInfo) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(c.UserIds)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c ChatListFolder) CreateChatFolderInviteLink(client *Client, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(c.ChatFolderId, name, chatIds)
}

// DeleteChatFolder is a helper method for Client.DeleteChatFolder
func (c ChatListFolder) DeleteChatFolder(client *Client, leaveChatIds []int64) (*Ok, error) {
	return client.DeleteChatFolder(c.ChatFolderId, leaveChatIds)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (c ChatListFolder) DeleteChatFolderInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(c.ChatFolderId, inviteLink)
}

// EditChatFolder is a helper method for Client.EditChatFolder
func (c ChatListFolder) EditChatFolder(client *Client, folder *ChatFolder) (*ChatFolderInfo, error) {
	return client.EditChatFolder(c.ChatFolderId, folder)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c ChatListFolder) EditChatFolderInviteLink(client *Client, inviteLink string, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(c.ChatFolderId, inviteLink, name, chatIds)
}

// GetChatFolder is a helper method for Client.GetChatFolder
func (c ChatListFolder) GetChatFolder(client *Client) (*ChatFolder, error) {
	return client.GetChatFolder(c.ChatFolderId)
}

// GetChatFolderChatsToLeave is a helper method for Client.GetChatFolderChatsToLeave
func (c ChatListFolder) GetChatFolderChatsToLeave(client *Client) (*Chats, error) {
	return client.GetChatFolderChatsToLeave(c.ChatFolderId)
}

// GetChatFolderInviteLinks is a helper method for Client.GetChatFolderInviteLinks
func (c ChatListFolder) GetChatFolderInviteLinks(client *Client) (*ChatFolderInviteLinks, error) {
	return client.GetChatFolderInviteLinks(c.ChatFolderId)
}

// GetChatFolderNewChats is a helper method for Client.GetChatFolderNewChats
func (c ChatListFolder) GetChatFolderNewChats(client *Client) (*Chats, error) {
	return client.GetChatFolderNewChats(c.ChatFolderId)
}

// GetChatsForChatFolderInviteLink is a helper method for Client.GetChatsForChatFolderInviteLink
func (c ChatListFolder) GetChatsForChatFolderInviteLink(client *Client) (*Chats, error) {
	return client.GetChatsForChatFolderInviteLink(c.ChatFolderId)
}

// ProcessChatFolderNewChats is a helper method for Client.ProcessChatFolderNewChats
func (c ChatListFolder) ProcessChatFolderNewChats(client *Client, addedChatIds []int64) (*Ok, error) {
	return client.ProcessChatFolderNewChats(c.ChatFolderId, addedChatIds)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (c ChatLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(c.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c ChatLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(c.Location, zoom, width, height, scale, chatId)
}

// Ban is a helper method for Client.BanChatMember
func (c ChatMember) Ban(client *Client, chatId int64, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(chatId, c.MemberId, bannedUntilDate, revokeMessages)
}

// Get is a helper method for Client.GetChatMember
func (c ChatMember) Get(client *Client, chatId int64) (*ChatMember, error) {
	return client.GetChatMember(chatId, c.MemberId)
}

// SetStatus is a helper method for Client.SetChatMemberStatus
func (c ChatMember) SetStatus(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatMemberStatus(chatId, c.MemberId, c.Status)
}

// SendChatAction is a helper method for Client.SendChatAction
func (c ChatMembersFilterMention) SendChatAction(client *Client, chatId int64, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(chatId, c.TopicId, businessConnectionId, opts)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatMemberStatusBanned) BanChatMember(client *Client, chatId int64, memberId *MessageSender, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(chatId, memberId, c.BannedUntilDate, revokeMessages)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (c ChatMemberStatusRestricted) SetChatPermissions(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPermissions(chatId, c.Permissions)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (c ChatPhotoInfo) AnswerInlineQuery(client *Client, inlineQueryId string, results []*InputInlineQueryResult, cacheTime int32, nextOffset string, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, c.IsPersonal, results, cacheTime, nextOffset, opts)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (c ChatPhotoStickerTypeCustomEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, name string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(name, c.CustomEmojiId)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (c ChatPhotoStickerTypeRegularOrMask) SetSupergroupStickerSet(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupStickerSet(supergroupId, c.StickerSetId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c ChatPosition) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, messageId int64) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, messageId, c.IsPinned)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (c ChatPosition) ToggleChatIsPinned(client *Client, chatList *ChatList, chatId int64) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, chatId, c.IsPinned)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c ChatPosition) ToggleForumTopicIsPinned(client *Client, chatId int64, forumTopicId int32) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(chatId, forumTopicId, c.IsPinned)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (c ChatPosition) ToggleSavedMessagesTopicIsPinned(client *Client, savedMessagesTopicId int64) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(savedMessagesTopicId, c.IsPinned)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (c ChatRevenueTransactions) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, c.NextOffset, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatRevenueTransactionTypeSponsoredMessageEarnings) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, c.StartDate, isRtmpStream)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatRevenueTransactionTypeSuggestedPostEarnings) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (c Chats) AddChatFolderByInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.AddChatFolderByInviteLink(inviteLink, c.ChatIds)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (c Chats) CreateChatFolderInviteLink(client *Client, chatFolderId int32, name string) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, name, c.ChatIds)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c Chats) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, name string) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, name, c.ChatIds)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (c Chats) OptimizeStorage(client *Client, size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, count, immunityDelay, fileTypes, c.ChatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// SetPinned is a helper method for Client.SetPinnedChats
func (c Chats) SetPinned(client *Client, chatList *ChatList) (*Ok, error) {
	return client.SetPinnedChats(chatList, c.ChatIds)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (c ChatSourcePublicServiceAnnouncement) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, c.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (c ChatSourcePublicServiceAnnouncement) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, c.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (c ChatSourcePublicServiceAnnouncement) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(c.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (c ChatSourcePublicServiceAnnouncement) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(c.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (c ChatSourcePublicServiceAnnouncement) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(c.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (c ChatSourcePublicServiceAnnouncement) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, c.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c ChatSourcePublicServiceAnnouncement) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, c.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (c ChatSourcePublicServiceAnnouncement) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, c.Text)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c ChatSourcePublicServiceAnnouncement) SaveApplicationLogEvent(client *Client, chatId int64, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(c.Type, chatId, data)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (c ChatSourcePublicServiceAnnouncement) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(c.Text, inputLanguageCodes)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatStatisticsAdministratorActionsInfo) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatStatisticsAdministratorActionsInfo) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatStatisticsAdministratorActionsInfo) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatStatisticsAdministratorActionsInfo) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatStatisticsAdministratorActionsInfo) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatStatisticsAdministratorActionsInfo) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatStatisticsAdministratorActionsInfo) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatStatisticsAdministratorActionsInfo) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatStatisticsAdministratorActionsInfo) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatStatisticsAdministratorActionsInfo) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatStatisticsAdministratorActionsInfo) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatStatisticsAdministratorActionsInfo) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatStatisticsAdministratorActionsInfo) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatStatisticsAdministratorActionsInfo) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatStatisticsAdministratorActionsInfo) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatStatisticsAdministratorActionsInfo) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatStatisticsAdministratorActionsInfo) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatStatisticsAdministratorActionsInfo) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatStatisticsAdministratorActionsInfo) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatStatisticsAdministratorActionsInfo) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatStatisticsAdministratorActionsInfo) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatStatisticsAdministratorActionsInfo) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatStatisticsAdministratorActionsInfo) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatStatisticsAdministratorActionsInfo) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatStatisticsAdministratorActionsInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c ChatStatisticsAdministratorActionsInfo) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c ChatStatisticsAdministratorActionsInfo) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c ChatStatisticsAdministratorActionsInfo) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c ChatStatisticsAdministratorActionsInfo) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c ChatStatisticsAdministratorActionsInfo) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c ChatStatisticsAdministratorActionsInfo) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c ChatStatisticsAdministratorActionsInfo) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c ChatStatisticsAdministratorActionsInfo) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c ChatStatisticsAdministratorActionsInfo) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c ChatStatisticsAdministratorActionsInfo) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c ChatStatisticsAdministratorActionsInfo) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c ChatStatisticsAdministratorActionsInfo) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c ChatStatisticsAdministratorActionsInfo) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c ChatStatisticsAdministratorActionsInfo) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c ChatStatisticsAdministratorActionsInfo) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c ChatStatisticsAdministratorActionsInfo) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c ChatStatisticsAdministratorActionsInfo) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c ChatStatisticsAdministratorActionsInfo) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatStatisticsInviterInfo) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c ChatStatisticsInviterInfo) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c ChatStatisticsInviterInfo) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c ChatStatisticsInviterInfo) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c ChatStatisticsInviterInfo) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (c ChatStatisticsInviterInfo) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c ChatStatisticsInviterInfo) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatStatisticsInviterInfo) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c ChatStatisticsInviterInfo) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c ChatStatisticsInviterInfo) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatStatisticsInviterInfo) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c ChatStatisticsInviterInfo) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c ChatStatisticsInviterInfo) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c ChatStatisticsInviterInfo) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c ChatStatisticsInviterInfo) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c ChatStatisticsInviterInfo) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c ChatStatisticsInviterInfo) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatStatisticsInviterInfo) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c ChatStatisticsInviterInfo) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c ChatStatisticsInviterInfo) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c ChatStatisticsInviterInfo) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c ChatStatisticsInviterInfo) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c ChatStatisticsInviterInfo) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (c ChatStatisticsInviterInfo) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c ChatStatisticsInviterInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}
