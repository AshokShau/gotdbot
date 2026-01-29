package gotdbot

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (b BotCommandScopeChatAdministrators) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(b.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (b BotCommandScopeChatAdministrators) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(b.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (b BotCommandScopeChatAdministrators) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(b.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (b BotCommandScopeChatAdministrators) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(b.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (b BotCommandScopeChatAdministrators) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(b.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (b BotCommandScopeChatAdministrators) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(b.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (b BotCommandScopeChatAdministrators) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, b.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (b BotCommandScopeChatAdministrators) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(b.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (b BotCommandScopeChatAdministrators) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(b.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (b BotCommandScopeChatAdministrators) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(b.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (b BotCommandScopeChatAdministrators) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(b.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (b BotCommandScopeChatAdministrators) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(b.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (b BotCommandScopeChatAdministrators) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(b.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (b BotCommandScopeChatAdministrators) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (b BotCommandScopeChatAdministrators) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, b.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (b BotCommandScopeChatAdministrators) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(b.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (b BotCommandScopeChatAdministrators) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (b BotCommandScopeChatAdministrators) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(b.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (b BotCommandScopeChatAdministrators) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(b.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (b BotCommandScopeChatAdministrators) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(b.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (b BotCommandScopeChatAdministrators) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(b.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (b BotCommandScopeChatAdministrators) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(b.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (b BotCommandScopeChatAdministrators) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(b.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (b BotCommandScopeChatAdministrators) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(b.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (b BotCommandScopeChatAdministrators) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(b.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (b BotCommandScopeChatAdministrators) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(b.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (b BotCommandScopeChatAdministrators) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, b.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (b BotCommandScopeChatAdministrators) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(b.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (b BotCommandScopeChatAdministrators) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(b.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (b BotCommandScopeChatAdministrators) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(b.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (b BotCommandScopeChatAdministrators) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(b.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BotCommandScopeChatAdministrators) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, b.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (b BotCommandScopeChatAdministrators) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (b BotCommandScopeChatAdministrators) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (b BotCommandScopeChatAdministrators) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(b.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (b BotCommandScopeChatAdministrators) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(b.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (b BotCommandScopeChatAdministrators) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(b.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (b BotCommandScopeChatAdministrators) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(b.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (b BotCommandScopeChatAdministrators) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(b.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (b BotCommandScopeChatAdministrators) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(b.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (b BotCommandScopeChatAdministrators) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, b.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (b BotCommandScopeChatAdministrators) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(b.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (b BotCommandScopeChatAdministrators) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(b.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (b BotCommandScopeChatAdministrators) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(b.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (b BotCommandScopeChatAdministrators) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(b.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (b BotCommandScopeChatAdministrators) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(b.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (b BotCommandScopeChatAdministrators) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(b.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (b BotCommandScopeChatAdministrators) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(b.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (b BotCommandScopeChatAdministrators) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(b.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (b BotCommandScopeChatAdministrators) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(b.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (b BotCommandScopeChatAdministrators) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(b.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (b BotCommandScopeChatAdministrators) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(b.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (b BotCommandScopeChatAdministrators) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(b.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (b BotCommandScopeChatAdministrators) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(b.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (b BotCommandScopeChatAdministrators) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(b.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (b BotCommandScopeChatAdministrators) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(b.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (b BotCommandScopeChatAdministrators) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(b.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (b BotCommandScopeChatAdministrators) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(b.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (b BotCommandScopeChatAdministrators) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(b.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (b BotCommandScopeChatAdministrators) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(b.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (b BotCommandScopeChatAdministrators) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(b.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (b BotCommandScopeChatAdministrators) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(b.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (b BotCommandScopeChatAdministrators) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(b.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (b BotCommandScopeChatAdministrators) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(b.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (b BotCommandScopeChatAdministrators) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(b.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (b BotCommandScopeChatAdministrators) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(b.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (b BotCommandScopeChatAdministrators) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(b.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (b BotCommandScopeChatAdministrators) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(b.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (b BotCommandScopeChatAdministrators) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(b.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (b BotCommandScopeChatAdministrators) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(b.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (b BotCommandScopeChatAdministrators) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(b.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (b BotCommandScopeChatAdministrators) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(b.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (b BotCommandScopeChatAdministrators) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(b.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (b BotCommandScopeChatAdministrators) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(b.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (b BotCommandScopeChatAdministrators) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(b.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (b BotCommandScopeChatAdministrators) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(b.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (b BotCommandScopeChatAdministrators) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(b.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (b BotCommandScopeChatAdministrators) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(b.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (b BotCommandScopeChatAdministrators) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(b.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (b BotCommandScopeChatAdministrators) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, b.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (b BotCommandScopeChatAdministrators) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(b.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (b BotCommandScopeChatAdministrators) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(b.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (b BotCommandScopeChatAdministrators) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(b.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (b BotCommandScopeChatAdministrators) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(b.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (b BotCommandScopeChatAdministrators) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(b.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (b BotCommandScopeChatAdministrators) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(b.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (b BotCommandScopeChatAdministrators) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(b.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (b BotCommandScopeChatAdministrators) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, b.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (b BotCommandScopeChatAdministrators) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(b.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (b BotCommandScopeChatAdministrators) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(b.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (b BotCommandScopeChatAdministrators) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(b.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (b BotCommandScopeChatAdministrators) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(b.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (b BotCommandScopeChatAdministrators) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(b.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (b BotCommandScopeChatAdministrators) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(b.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (b BotCommandScopeChatAdministrators) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(b.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (b BotCommandScopeChatAdministrators) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(b.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (b BotCommandScopeChatAdministrators) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(b.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (b BotCommandScopeChatAdministrators) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(b.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (b BotCommandScopeChatAdministrators) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(b.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (b BotCommandScopeChatAdministrators) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(b.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (b BotCommandScopeChatAdministrators) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(b.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (b BotCommandScopeChatAdministrators) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(b.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (b BotCommandScopeChatMember) AddChatMember(client *Client, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(b.ChatId, b.UserId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (b BotCommandScopeChatMember) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(b.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (b BotCommandScopeChatMember) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(b.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (b BotCommandScopeChatMember) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(b.ChatId, messageId, tasks)
}

// AddContact is a helper method for Client.AddContact
func (b BotCommandScopeChatMember) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(b.UserId, contact, sharePhoneNumber)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (b BotCommandScopeChatMember) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, b.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (b BotCommandScopeChatMember) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(b.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (b BotCommandScopeChatMember) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(b.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (b BotCommandScopeChatMember) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(b.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (b BotCommandScopeChatMember) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(b.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (b BotCommandScopeChatMember) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(b.ChatId)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (b BotCommandScopeChatMember) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(b.UserId, name, sticker)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (b BotCommandScopeChatMember) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (b BotCommandScopeChatMember) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(b.UserId, refundPayments)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (b BotCommandScopeChatMember) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(b.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (b BotCommandScopeChatMember) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(b.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (b BotCommandScopeChatMember) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(b.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (b BotCommandScopeChatMember) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(b.ChatId)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (b BotCommandScopeChatMember) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(b.UserId, onlyLocal)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (b BotCommandScopeChatMember) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(b.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (b BotCommandScopeChatMember) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(b.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (b BotCommandScopeChatMember) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(b.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (b BotCommandScopeChatMember) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(b.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (b BotCommandScopeChatMember) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(b.ChatId, messageId)
}

// CreateCall is a helper method for Client.CreateCall
func (b BotCommandScopeChatMember) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(b.UserId, protocol, isVideo)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (b BotCommandScopeChatMember) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(b.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (b BotCommandScopeChatMember) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(b.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (b BotCommandScopeChatMember) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(b.ChatId, name, isNameImplicit, icon)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (b BotCommandScopeChatMember) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(b.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (b BotCommandScopeChatMember) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(b.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (b BotCommandScopeChatMember) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(b.UserId, force)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (b BotCommandScopeChatMember) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(b.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (b BotCommandScopeChatMember) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(b.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (b BotCommandScopeChatMember) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(b.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (b BotCommandScopeChatMember) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(b.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (b BotCommandScopeChatMember) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(b.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (b BotCommandScopeChatMember) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(b.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (b BotCommandScopeChatMember) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(b.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (b BotCommandScopeChatMember) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(b.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (b BotCommandScopeChatMember) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(b.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (b BotCommandScopeChatMember) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(b.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (b BotCommandScopeChatMember) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(b.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (b BotCommandScopeChatMember) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(b.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (b BotCommandScopeChatMember) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(b.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (b BotCommandScopeChatMember) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(b.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (b BotCommandScopeChatMember) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(b.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (b BotCommandScopeChatMember) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(b.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (b BotCommandScopeChatMember) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, b.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (b BotCommandScopeChatMember) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, b.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (b BotCommandScopeChatMember) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, b.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (b BotCommandScopeChatMember) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, b.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (b BotCommandScopeChatMember) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, b.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (b BotCommandScopeChatMember) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, b.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (b BotCommandScopeChatMember) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(b.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (b BotCommandScopeChatMember) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(b.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (b BotCommandScopeChatMember) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(b.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (b BotCommandScopeChatMember) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(b.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (b BotCommandScopeChatMember) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(b.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (b BotCommandScopeChatMember) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(b.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (b BotCommandScopeChatMember) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(b.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (b BotCommandScopeChatMember) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(b.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (b BotCommandScopeChatMember) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(b.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (b BotCommandScopeChatMember) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(b.ChatId, messageId, inputMessageContent, opts)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (b BotCommandScopeChatMember) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(b.UserId, telegramPaymentChargeId, isCanceled)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (b BotCommandScopeChatMember) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(b.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (b BotCommandScopeChatMember) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, b.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (b BotCommandScopeChatMember) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(b.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (b BotCommandScopeChatMember) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(b.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (b BotCommandScopeChatMember) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(b.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (b BotCommandScopeChatMember) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(b.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (b BotCommandScopeChatMember) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(b.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (b BotCommandScopeChatMember) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(b.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (b BotCommandScopeChatMember) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(b.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (b BotCommandScopeChatMember) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(b.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (b BotCommandScopeChatMember) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(b.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (b BotCommandScopeChatMember) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(b.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (b BotCommandScopeChatMember) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(b.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (b BotCommandScopeChatMember) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(b.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (b BotCommandScopeChatMember) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(b.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (b BotCommandScopeChatMember) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(b.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (b BotCommandScopeChatMember) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(b.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (b BotCommandScopeChatMember) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(b.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (b BotCommandScopeChatMember) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(b.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (b BotCommandScopeChatMember) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(b.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (b BotCommandScopeChatMember) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(b.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (b BotCommandScopeChatMember) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(b.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (b BotCommandScopeChatMember) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(b.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (b BotCommandScopeChatMember) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(b.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (b BotCommandScopeChatMember) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(b.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (b BotCommandScopeChatMember) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(b.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (b BotCommandScopeChatMember) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(b.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (b BotCommandScopeChatMember) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(b.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (b BotCommandScopeChatMember) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(b.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (b BotCommandScopeChatMember) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(b.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (b BotCommandScopeChatMember) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(b.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (b BotCommandScopeChatMember) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(b.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (b BotCommandScopeChatMember) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(b.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (b BotCommandScopeChatMember) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(b.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (b BotCommandScopeChatMember) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(b.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (b BotCommandScopeChatMember) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(b.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (b BotCommandScopeChatMember) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(b.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (b BotCommandScopeChatMember) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(b.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (b BotCommandScopeChatMember) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(b.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (b BotCommandScopeChatMember) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(b.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (b BotCommandScopeChatMember) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(b.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (b BotCommandScopeChatMember) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(b.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (b BotCommandScopeChatMember) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(b.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (b BotCommandScopeChatMember) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(b.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (b BotCommandScopeChatMember) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(b.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (b BotCommandScopeChatMember) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(b.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (b BotCommandScopeChatMember) GetGameHighScores(client *Client, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(b.ChatId, messageId, b.UserId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (b BotCommandScopeChatMember) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(b.ChatId, messageId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (b BotCommandScopeChatMember) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(b.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (b BotCommandScopeChatMember) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, b.UserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (b BotCommandScopeChatMember) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, b.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (b BotCommandScopeChatMember) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(b.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (b BotCommandScopeChatMember) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(b.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (b BotCommandScopeChatMember) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(b.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (b BotCommandScopeChatMember) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(b.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (b BotCommandScopeChatMember) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, b.ChatId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (b BotCommandScopeChatMember) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(b.UserId)
}

// GetMessage is a helper method for Client.GetMessage
func (b BotCommandScopeChatMember) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(b.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (b BotCommandScopeChatMember) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(b.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (b BotCommandScopeChatMember) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(b.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (b BotCommandScopeChatMember) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(b.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (b BotCommandScopeChatMember) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(b.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (b BotCommandScopeChatMember) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(b.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (b BotCommandScopeChatMember) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(b.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (b BotCommandScopeChatMember) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(b.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (b BotCommandScopeChatMember) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(b.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (b BotCommandScopeChatMember) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(b.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (b BotCommandScopeChatMember) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(b.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (b BotCommandScopeChatMember) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(b.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (b BotCommandScopeChatMember) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(b.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (b BotCommandScopeChatMember) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(b.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (b BotCommandScopeChatMember) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(b.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (b BotCommandScopeChatMember) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(b.ChatId, messageId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (b BotCommandScopeChatMember) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(b.UserId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (b BotCommandScopeChatMember) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(b.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (b BotCommandScopeChatMember) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(b.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (b BotCommandScopeChatMember) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(b.ChatId, messageId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (b BotCommandScopeChatMember) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(b.UserId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (b BotCommandScopeChatMember) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(b.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (b BotCommandScopeChatMember) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, b.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (b BotCommandScopeChatMember) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(b.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (b BotCommandScopeChatMember) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(b.ChatId, storyId, isDark)
}

// GetUser is a helper method for Client.GetUser
func (b BotCommandScopeChatMember) GetUser(client *Client) (*User, error) {
	return client.GetUser(b.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (b BotCommandScopeChatMember) GetUserChatBoosts(client *Client) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(b.ChatId, b.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (b BotCommandScopeChatMember) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(b.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (b BotCommandScopeChatMember) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(b.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (b BotCommandScopeChatMember) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(b.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (b BotCommandScopeChatMember) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(b.UserId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (b BotCommandScopeChatMember) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(b.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (b BotCommandScopeChatMember) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(b.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (b BotCommandScopeChatMember) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(b.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (b BotCommandScopeChatMember) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(b.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (b BotCommandScopeChatMember) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(b.UserId, starCount, monthCount, text)
}

// ImportMessages is a helper method for Client.ImportMessages
func (b BotCommandScopeChatMember) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(b.ChatId, messageFile, attachedFiles)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (b BotCommandScopeChatMember) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, b.UserId, isVideo)
}

// JoinChat is a helper method for Client.JoinChat
func (b BotCommandScopeChatMember) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(b.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (b BotCommandScopeChatMember) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(b.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (b BotCommandScopeChatMember) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(b.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (b BotCommandScopeChatMember) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(b.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (b BotCommandScopeChatMember) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(b.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (b BotCommandScopeChatMember) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(b.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (b BotCommandScopeChatMember) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(b.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BotCommandScopeChatMember) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(b.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (b BotCommandScopeChatMember) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(b.ChatId, messageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (b BotCommandScopeChatMember) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, b.UserId, text, isPrivate)
}

// PostStory is a helper method for Client.PostStory
func (b BotCommandScopeChatMember) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(b.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (b BotCommandScopeChatMember) ProcessChatJoinRequest(client *Client, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(b.ChatId, b.UserId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (b BotCommandScopeChatMember) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(b.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (b BotCommandScopeChatMember) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(b.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (b BotCommandScopeChatMember) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(b.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (b BotCommandScopeChatMember) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(b.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (b BotCommandScopeChatMember) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(b.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (b BotCommandScopeChatMember) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(b.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (b BotCommandScopeChatMember) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(b.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (b BotCommandScopeChatMember) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, b.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (b BotCommandScopeChatMember) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(b.ChatId, messageId)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (b BotCommandScopeChatMember) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(b.UserId, telegramPaymentChargeId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (b BotCommandScopeChatMember) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(b.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (b BotCommandScopeChatMember) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(b.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (b BotCommandScopeChatMember) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(b.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (b BotCommandScopeChatMember) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(b.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (b BotCommandScopeChatMember) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(b.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (b BotCommandScopeChatMember) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (b BotCommandScopeChatMember) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, b.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (b BotCommandScopeChatMember) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(b.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (b BotCommandScopeChatMember) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (b BotCommandScopeChatMember) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(b.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (b BotCommandScopeChatMember) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(b.ChatId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (b BotCommandScopeChatMember) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(b.UserId, name, oldSticker, newSticker)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (b BotCommandScopeChatMember) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(b.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (b BotCommandScopeChatMember) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(b.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (b BotCommandScopeChatMember) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(b.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (b BotCommandScopeChatMember) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(b.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (b BotCommandScopeChatMember) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(b.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (b BotCommandScopeChatMember) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(b.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (b BotCommandScopeChatMember) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(b.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (b BotCommandScopeChatMember) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, b.ChatId, data)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (b BotCommandScopeChatMember) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(b.UserId, result, chatTypes)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (b BotCommandScopeChatMember) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(b.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (b BotCommandScopeChatMember) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(b.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (b BotCommandScopeChatMember) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(b.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (b BotCommandScopeChatMember) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(b.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BotCommandScopeChatMember) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, b.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (b BotCommandScopeChatMember) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (b BotCommandScopeChatMember) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (b BotCommandScopeChatMember) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(b.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (b BotCommandScopeChatMember) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(b.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (b BotCommandScopeChatMember) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(b.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (b BotCommandScopeChatMember) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(b.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (b BotCommandScopeChatMember) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(b.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (b BotCommandScopeChatMember) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(b.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (b BotCommandScopeChatMember) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, b.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (b BotCommandScopeChatMember) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(b.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (b BotCommandScopeChatMember) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(b.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (b BotCommandScopeChatMember) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(b.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (b BotCommandScopeChatMember) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(b.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (b BotCommandScopeChatMember) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(b.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (b BotCommandScopeChatMember) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(b.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (b BotCommandScopeChatMember) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(b.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (b BotCommandScopeChatMember) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(b.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (b BotCommandScopeChatMember) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(b.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (b BotCommandScopeChatMember) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(b.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (b BotCommandScopeChatMember) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(b.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (b BotCommandScopeChatMember) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(b.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (b BotCommandScopeChatMember) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(b.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (b BotCommandScopeChatMember) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(b.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (b BotCommandScopeChatMember) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(b.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (b BotCommandScopeChatMember) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(b.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (b BotCommandScopeChatMember) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(b.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (b BotCommandScopeChatMember) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(b.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (b BotCommandScopeChatMember) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(b.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (b BotCommandScopeChatMember) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(b.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (b BotCommandScopeChatMember) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(b.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (b BotCommandScopeChatMember) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(b.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (b BotCommandScopeChatMember) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(b.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (b BotCommandScopeChatMember) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(b.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (b BotCommandScopeChatMember) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(b.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (b BotCommandScopeChatMember) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(b.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (b BotCommandScopeChatMember) SetGameScore(client *Client, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(b.ChatId, messageId, editMessage, b.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (b BotCommandScopeChatMember) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, b.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (b BotCommandScopeChatMember) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(b.UserId, menuButton)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (b BotCommandScopeChatMember) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(b.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (b BotCommandScopeChatMember) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(b.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (b BotCommandScopeChatMember) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(b.ChatId, messageId, typeField)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (b BotCommandScopeChatMember) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(b.UserId, errors)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (b BotCommandScopeChatMember) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(b.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (b BotCommandScopeChatMember) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(b.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (b BotCommandScopeChatMember) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(b.ChatId, messageId, optionIds)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (b BotCommandScopeChatMember) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(b.UserId, name, opts)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (b BotCommandScopeChatMember) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(b.ChatId, storyAlbumId, name)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (b BotCommandScopeChatMember) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(b.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (b BotCommandScopeChatMember) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(b.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (b BotCommandScopeChatMember) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(b.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (b BotCommandScopeChatMember) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(b.UserId, message)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (b BotCommandScopeChatMember) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(b.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (b BotCommandScopeChatMember) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(b.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (b BotCommandScopeChatMember) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(b.UserId)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (b BotCommandScopeChatMember) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(b.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (b BotCommandScopeChatMember) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(b.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (b BotCommandScopeChatMember) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, b.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (b BotCommandScopeChatMember) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(b.ChatId, messageId, opts)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (b BotCommandScopeChatMember) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(b.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (b BotCommandScopeChatMember) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(b.UserId, photo)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (b BotCommandScopeChatMember) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(b.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (b BotCommandScopeChatMember) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(b.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (b BotCommandScopeChatMember) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(b.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (b BotCommandScopeChatMember) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(b.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (b BotCommandScopeChatMember) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(b.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (b BotCommandScopeChatMember) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(b.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (b BotCommandScopeChatMember) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, b.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (b BotCommandScopeChatMember) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(b.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (b BotCommandScopeChatMember) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(b.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (b BotCommandScopeChatMember) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(b.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (b BotCommandScopeChatMember) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(b.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (b BotCommandScopeChatMember) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(b.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (b BotCommandScopeChatMember) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(b.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (b BotCommandScopeChatMember) TransferChatOwnership(client *Client, password string) (*Ok, error) {
	return client.TransferChatOwnership(b.ChatId, b.UserId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (b BotCommandScopeChatMember) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(b.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (b BotCommandScopeChatMember) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(b.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (b BotCommandScopeChatMember) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(b.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (b BotCommandScopeChatMember) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(b.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (b BotCommandScopeChatMember) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(b.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (b BotCommandScopeChatMember) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(b.ChatId)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (b BotCommandScopeChatMember) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(b.UserId, stickerFormat, sticker)
}

// ViewMessages is a helper method for Client.ViewMessages
func (b BotCommandScopeChatMember) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(b.ChatId, messageIds, forceRead, opts)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (b BotInfo) CreateNewSupergroupChat(client *Client, title string, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(title, isForum, isChannel, b.Description, messageAutoDeleteTime, forImport, opts)
}

// SetDescription is a helper method for Client.SetBotInfoDescription
func (b BotInfo) SetDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, b.Description)
}

// SetShortDescription is a helper method for Client.SetBotInfoShortDescription
func (b BotInfo) SetShortDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoShortDescription(botUserId, languageCode, b.ShortDescription)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (b BotInfo) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, b.Description)
}

// SetCommands is a helper method for Client.SetCommands
func (b BotInfo) SetCommands(client *Client, languageCode string, opts *SetCommandsOpts) (*Ok, error) {
	return client.SetCommands(languageCode, b.Commands, opts)
}

// SetDefaultChannelAdministratorRights is a helper method for Client.SetDefaultChannelAdministratorRights
func (b BotInfo) SetDefaultChannelAdministratorRights(client *Client) (*Ok, error) {
	return client.SetDefaultChannelAdministratorRights(b.DefaultChannelAdministratorRights)
}

// SetDefaultGroupAdministratorRights is a helper method for Client.SetDefaultGroupAdministratorRights
func (b BotInfo) SetDefaultGroupAdministratorRights(client *Client) (*Ok, error) {
	return client.SetDefaultGroupAdministratorRights(b.DefaultGroupAdministratorRights)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (b BotInfo) SetMenuButton(client *Client, userId int64) (*Ok, error) {
	return client.SetMenuButton(userId, b.MenuButton)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (b BotInfo) ToggleBotCanManageEmojiStatus(client *Client, botUserId int64) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(botUserId, b.CanManageEmojiStatus)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (b BotMediaPreview) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, b.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (b BotMediaPreview) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, b.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (b BotMediaPreview) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, b.Date)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (b BotMenuButton) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, b.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (b BotMenuButton) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, b.Text, showAlert, b.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (b BotMenuButton) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, b.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (b BotMenuButton) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, b.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (b BotMenuButton) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(b.Url)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (b BotMenuButton) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(b.Text, inputLanguageCodes)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (b BotMenuButton) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(b.Url)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (b BotMenuButton) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(b.Text)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (b BotMenuButton) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, b.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (b BotMenuButton) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(b.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BotMenuButton) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, b.Url, parameters, opts)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (b BotMenuButton) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(b.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (b BotMenuButton) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, b.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (b BotMenuButton) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, b.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (b BotMenuButton) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, b.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (b BotMenuButton) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(b.Text, inputLanguageCodes)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (b BotVerification) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(b.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (b BotVerification) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(b.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (b BotVerification) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(b.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (b BotVerification) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(b.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (b BotVerification) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, b.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (b BotVerification) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (b BotVerification) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(b.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (b BotVerification) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(b.BotUserId, languageCode, fileId, content)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (b BotVerification) EditForumTopic(client *Client, chatId int64, forumTopicId int32, name string, editIconCustomEmoji bool) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, name, editIconCustomEmoji, b.IconCustomEmojiId)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (b BotVerification) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(b.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (b BotVerification) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(b.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (b BotVerification) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(b.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (b BotVerification) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(b.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (b BotVerification) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(b.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (b BotVerification) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(b.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (b BotVerification) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(b.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (b BotVerification) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(b.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (b BotVerification) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, b.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (b BotVerification) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(b.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (b BotVerification) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, b.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (b BotVerification) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(b.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (b BotVerification) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(b.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (b BotVerification) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, b.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (b BotVerification) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(b.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (b BotVerification) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(b.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (b BotVerification) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(b.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BotVerification) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, b.BotUserId, url, parameters, opts)
}

// RemoveMessageSender is a helper method for Client.RemoveMessageSenderBotVerification
func (b BotVerification) RemoveMessageSender(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(b.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (b BotVerification) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(b.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (b BotVerification) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (b BotVerification) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(b.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BotVerification) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(b.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (b BotVerification) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(b.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (b BotVerification) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(b.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (b BotVerification) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(b.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (b BotVerification) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(b.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (b BotVerification) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(b.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (b BotVerification) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(b.BotUserId, photo)
}

// SetMessageSender is a helper method for Client.SetMessageSenderBotVerification
func (b BotVerification) SetMessageSender(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(b.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (b BotVerification) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(b.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (b BotVerification) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(b.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (b BotVerification) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(b.BotUserId, username, isActive)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (b BotVerificationParameters) EditForumTopic(client *Client, chatId int64, forumTopicId int32, name string, editIconCustomEmoji bool) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, name, editIconCustomEmoji, b.IconCustomEmojiId)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (b BusinessAwayMessageScheduleCustom) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, b.StartDate, isRtmpStream)
}

// DeleteQuickReplyShortcut is a helper method for Client.DeleteQuickReplyShortcut
func (b BusinessAwayMessageSettings) DeleteQuickReplyShortcut(client *Client) (*Ok, error) {
	return client.DeleteQuickReplyShortcut(b.ShortcutId)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (b BusinessAwayMessageSettings) DeleteQuickReplyShortcutMessages(client *Client, messageIds []int64) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(b.ShortcutId, messageIds)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (b BusinessAwayMessageSettings) EditQuickReplyMessage(client *Client, messageId int64, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(b.ShortcutId, messageId, inputMessageContent)
}

// LoadQuickReplyShortcutMessages is a helper method for Client.LoadQuickReplyShortcutMessages
func (b BusinessAwayMessageSettings) LoadQuickReplyShortcutMessages(client *Client) (*Ok, error) {
	return client.LoadQuickReplyShortcutMessages(b.ShortcutId)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (b BusinessAwayMessageSettings) SendQuickReplyShortcutMessages(client *Client, chatId int64, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, b.ShortcutId, sendingId)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (b BusinessAwayMessageSettings) SetQuickReplyShortcutName(client *Client, name string) (*Ok, error) {
	return client.SetQuickReplyShortcutName(b.ShortcutId, name)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (b BusinessBotManageBar) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(b.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (b BusinessBotManageBar) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(b.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (b BusinessBotManageBar) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(b.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (b BusinessBotManageBar) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(b.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (b BusinessBotManageBar) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, b.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (b BusinessBotManageBar) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (b BusinessBotManageBar) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(b.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (b BusinessBotManageBar) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(b.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (b BusinessBotManageBar) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(b.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (b BusinessBotManageBar) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(b.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (b BusinessBotManageBar) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(b.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (b BusinessBotManageBar) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(b.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (b BusinessBotManageBar) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(b.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (b BusinessBotManageBar) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(b.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (b BusinessBotManageBar) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(b.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (b BusinessBotManageBar) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(b.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (b BusinessBotManageBar) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, b.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (b BusinessBotManageBar) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(b.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (b BusinessBotManageBar) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, b.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (b BusinessBotManageBar) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(b.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (b BusinessBotManageBar) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(b.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (b BusinessBotManageBar) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, b.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (b BusinessBotManageBar) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(b.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (b BusinessBotManageBar) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(b.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (b BusinessBotManageBar) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(b.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BusinessBotManageBar) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, b.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (b BusinessBotManageBar) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(b.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (b BusinessBotManageBar) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(b.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (b BusinessBotManageBar) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (b BusinessBotManageBar) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(b.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BusinessBotManageBar) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(b.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (b BusinessBotManageBar) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(b.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (b BusinessBotManageBar) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(b.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (b BusinessBotManageBar) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(b.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (b BusinessBotManageBar) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(b.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (b BusinessBotManageBar) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(b.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (b BusinessBotManageBar) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(b.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (b BusinessBotManageBar) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(b.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (b BusinessBotManageBar) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(b.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (b BusinessBotManageBar) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(b.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (b BusinessBotManageBar) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(b.BotUserId, username, isActive)
}

// ConfirmQrCodeAuthentication is a helper method for Client.ConfirmQrCodeAuthentication
func (b BusinessChatLink) ConfirmQrCodeAuthentication(client *Client) (*Session, error) {
	return client.ConfirmQrCodeAuthentication(b.Link)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (b BusinessChatLink) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, b.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (b BusinessChatLink) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, b.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (b BusinessChatLink) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(b.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (b BusinessChatLink) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, b.Title, startDate, isRtmpStream)
}

// Delete is a helper method for Client.DeleteBusinessChatLink
func (b BusinessChatLink) Delete(client *Client) (*Ok, error) {
	return client.DeleteBusinessChatLink(b.Link)
}

// Edit is a helper method for Client.EditBusinessChatLink
func (b BusinessChatLink) Edit(client *Client, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	return client.EditBusinessChatLink(b.Link, linkInfo)
}

// GetDeepLinkInfo is a helper method for Client.GetDeepLinkInfo
func (b BusinessChatLink) GetDeepLinkInfo(client *Client) (*DeepLinkInfo, error) {
	return client.GetDeepLinkInfo(b.Link)
}

// GetExternalLink is a helper method for Client.GetExternalLink
func (b BusinessChatLink) GetExternalLink(client *Client, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetExternalLink(b.Link, allowWriteAccess)
}

// GetExternalLinkInfo is a helper method for Client.GetExternalLinkInfo
func (b BusinessChatLink) GetExternalLinkInfo(client *Client) (*LoginUrlInfo, error) {
	return client.GetExternalLinkInfo(b.Link)
}

// GetInternalLinkType is a helper method for Client.GetInternalLinkType
func (b BusinessChatLink) GetInternalLinkType(client *Client) (*InternalLinkType, error) {
	return client.GetInternalLinkType(b.Link)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (b BusinessChatLink) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(b.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (b BusinessChatLink) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(b.Text)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (b BusinessChatLink) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(b.Title)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (b BusinessChatLink) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, b.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (b BusinessChatLink) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(b.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (b BusinessChatLink) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, b.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (b BusinessChatLink) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(b.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (b BusinessChatLink) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, b.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (b BusinessChatLink) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, b.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (b BusinessChatLink) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, b.Text)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (b BusinessChatLink) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, b.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (b BusinessChatLink) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, b.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (b BusinessChatLink) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, b.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (b BusinessChatLink) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, b.Title, recordVideo, usePortraitOrientation)
}

// TranslateText is a helper method for Client.TranslateText
func (b BusinessChatLink) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(b.Text, toLanguageCode)
}

// AddChatMember is a helper method for Client.AddChatMember
func (b BusinessChatLinkInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(b.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (b BusinessChatLinkInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(b.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (b BusinessChatLinkInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(b.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (b BusinessChatLinkInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(b.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (b BusinessChatLinkInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, b.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (b BusinessChatLinkInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(b.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (b BusinessChatLinkInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(b.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (b BusinessChatLinkInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(b.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (b BusinessChatLinkInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(b.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (b BusinessChatLinkInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(b.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (b BusinessChatLinkInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (b BusinessChatLinkInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(b.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (b BusinessChatLinkInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(b.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (b BusinessChatLinkInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(b.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (b BusinessChatLinkInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(b.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (b BusinessChatLinkInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(b.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (b BusinessChatLinkInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(b.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (b BusinessChatLinkInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(b.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (b BusinessChatLinkInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(b.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (b BusinessChatLinkInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(b.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (b BusinessChatLinkInfo) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(b.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (b BusinessChatLinkInfo) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(b.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (b BusinessChatLinkInfo) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(b.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (b BusinessChatLinkInfo) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(b.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (b BusinessChatLinkInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(b.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (b BusinessChatLinkInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(b.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (b BusinessChatLinkInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(b.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (b BusinessChatLinkInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(b.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (b BusinessChatLinkInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(b.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (b BusinessChatLinkInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(b.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (b BusinessChatLinkInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(b.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (b BusinessChatLinkInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(b.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (b BusinessChatLinkInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(b.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (b BusinessChatLinkInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(b.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (b BusinessChatLinkInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(b.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (b BusinessChatLinkInfo) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(b.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (b BusinessChatLinkInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(b.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (b BusinessChatLinkInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(b.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (b BusinessChatLinkInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(b.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (b BusinessChatLinkInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, b.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (b BusinessChatLinkInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, b.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (b BusinessChatLinkInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, b.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (b BusinessChatLinkInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, b.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (b BusinessChatLinkInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, b.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (b BusinessChatLinkInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, b.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (b BusinessChatLinkInfo) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(b.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (b BusinessChatLinkInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(b.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (b BusinessChatLinkInfo) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(b.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (b BusinessChatLinkInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(b.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (b BusinessChatLinkInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(b.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (b BusinessChatLinkInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(b.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (b BusinessChatLinkInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(b.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (b BusinessChatLinkInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(b.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (b BusinessChatLinkInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(b.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (b BusinessChatLinkInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(b.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (b BusinessChatLinkInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(b.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (b BusinessChatLinkInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, b.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (b BusinessChatLinkInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(b.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (b BusinessChatLinkInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(b.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (b BusinessChatLinkInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(b.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (b BusinessChatLinkInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(b.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (b BusinessChatLinkInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(b.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (b BusinessChatLinkInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(b.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (b BusinessChatLinkInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(b.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (b BusinessChatLinkInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(b.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (b BusinessChatLinkInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(b.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (b BusinessChatLinkInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(b.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (b BusinessChatLinkInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(b.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (b BusinessChatLinkInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(b.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (b BusinessChatLinkInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(b.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (b BusinessChatLinkInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(b.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (b BusinessChatLinkInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(b.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (b BusinessChatLinkInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(b.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (b BusinessChatLinkInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(b.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (b BusinessChatLinkInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(b.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (b BusinessChatLinkInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(b.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (b BusinessChatLinkInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(b.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (b BusinessChatLinkInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(b.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (b BusinessChatLinkInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(b.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (b BusinessChatLinkInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(b.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (b BusinessChatLinkInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(b.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (b BusinessChatLinkInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(b.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (b BusinessChatLinkInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(b.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (b BusinessChatLinkInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(b.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (b BusinessChatLinkInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(b.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (b BusinessChatLinkInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(b.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (b BusinessChatLinkInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(b.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (b BusinessChatLinkInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(b.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (b BusinessChatLinkInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(b.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (b BusinessChatLinkInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(b.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (b BusinessChatLinkInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(b.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (b BusinessChatLinkInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(b.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (b BusinessChatLinkInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(b.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (b BusinessChatLinkInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(b.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (b BusinessChatLinkInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(b.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (b BusinessChatLinkInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(b.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (b BusinessChatLinkInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(b.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (b BusinessChatLinkInfo) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(b.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (b BusinessChatLinkInfo) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(b.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (b BusinessChatLinkInfo) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(b.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (b BusinessChatLinkInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(b.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (b BusinessChatLinkInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(b.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (b BusinessChatLinkInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(b.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (b BusinessChatLinkInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, b.ChatId, query, offset, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (b BusinessChatLinkInfo) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(b.Text, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (b BusinessChatLinkInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(b.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (b BusinessChatLinkInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(b.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (b BusinessChatLinkInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(b.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (b BusinessChatLinkInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(b.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (b BusinessChatLinkInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, b.ChatId)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (b BusinessChatLinkInfo) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(b.Text)
}

// GetMessage is a helper method for Client.GetMessage
func (b BusinessChatLinkInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(b.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (b BusinessChatLinkInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(b.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (b BusinessChatLinkInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(b.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (b BusinessChatLinkInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(b.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (b BusinessChatLinkInfo) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(b.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (b BusinessChatLinkInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(b.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (b BusinessChatLinkInfo) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(b.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (b BusinessChatLinkInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(b.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (b BusinessChatLinkInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(b.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (b BusinessChatLinkInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(b.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (b BusinessChatLinkInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(b.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (b BusinessChatLinkInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(b.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (b BusinessChatLinkInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(b.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (b BusinessChatLinkInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(b.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (b BusinessChatLinkInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(b.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (b BusinessChatLinkInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(b.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (b BusinessChatLinkInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(b.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (b BusinessChatLinkInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(b.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (b BusinessChatLinkInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(b.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (b BusinessChatLinkInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(b.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (b BusinessChatLinkInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, b.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (b BusinessChatLinkInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(b.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (b BusinessChatLinkInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(b.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (b BusinessChatLinkInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(b.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (b BusinessChatLinkInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(b.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (b BusinessChatLinkInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(b.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (b BusinessChatLinkInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(b.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (b BusinessChatLinkInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(b.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (b BusinessChatLinkInfo) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, b.Text)
}

// ImportMessages is a helper method for Client.ImportMessages
func (b BusinessChatLinkInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(b.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (b BusinessChatLinkInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(b.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (b BusinessChatLinkInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(b.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (b BusinessChatLinkInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(b.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (b BusinessChatLinkInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(b.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (b BusinessChatLinkInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(b.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (b BusinessChatLinkInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(b.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (b BusinessChatLinkInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(b.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BusinessChatLinkInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(b.ChatId, botUserId, url, parameters, opts)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (b BusinessChatLinkInfo) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(b.Text)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (b BusinessChatLinkInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(b.ChatId, messageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (b BusinessChatLinkInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, b.Text, isPrivate)
}

// PostStory is a helper method for Client.PostStory
func (b BusinessChatLinkInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(b.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (b BusinessChatLinkInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(b.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (b BusinessChatLinkInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(b.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (b BusinessChatLinkInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(b.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (b BusinessChatLinkInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(b.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (b BusinessChatLinkInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(b.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (b BusinessChatLinkInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(b.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (b BusinessChatLinkInfo) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(b.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (b BusinessChatLinkInfo) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(b.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (b BusinessChatLinkInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, b.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (b BusinessChatLinkInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(b.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (b BusinessChatLinkInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(b.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (b BusinessChatLinkInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(b.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (b BusinessChatLinkInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(b.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (b BusinessChatLinkInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(b.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (b BusinessChatLinkInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(b.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (b BusinessChatLinkInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (b BusinessChatLinkInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, b.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (b BusinessChatLinkInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(b.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (b BusinessChatLinkInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(b.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (b BusinessChatLinkInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(b.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (b BusinessChatLinkInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(b.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (b BusinessChatLinkInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(b.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (b BusinessChatLinkInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(b.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (b BusinessChatLinkInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(b.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (b BusinessChatLinkInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(b.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (b BusinessChatLinkInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(b.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (b BusinessChatLinkInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(b.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (b BusinessChatLinkInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(b.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (b BusinessChatLinkInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, b.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (b BusinessChatLinkInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(b.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (b BusinessChatLinkInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(b.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (b BusinessChatLinkInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(b.ChatId, limit)
}

// SearchQuote is a helper method for Client.SearchQuote
func (b BusinessChatLinkInfo) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(b.Text, quote, quotePosition)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (b BusinessChatLinkInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(b.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BusinessChatLinkInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, b.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (b BusinessChatLinkInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (b BusinessChatLinkInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, b.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (b BusinessChatLinkInfo) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(b.ChatId, topicId, businessConnectionId, opts)
}

// SendGift is a helper method for Client.SendGift
func (b BusinessChatLinkInfo) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, b.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (b BusinessChatLinkInfo) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, b.Text, paidMessageStarCount)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (b BusinessChatLinkInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(b.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (b BusinessChatLinkInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(b.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (b BusinessChatLinkInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(b.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (b BusinessChatLinkInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(b.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (b BusinessChatLinkInfo) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(b.ChatId, forumTopicId, draftId, b.Text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (b BusinessChatLinkInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, b.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (b BusinessChatLinkInfo) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(b.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (b BusinessChatLinkInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(b.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (b BusinessChatLinkInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(b.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (b BusinessChatLinkInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(b.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (b BusinessChatLinkInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(b.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (b BusinessChatLinkInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(b.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (b BusinessChatLinkInfo) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(b.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (b BusinessChatLinkInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(b.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (b BusinessChatLinkInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(b.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (b BusinessChatLinkInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(b.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (b BusinessChatLinkInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(b.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (b BusinessChatLinkInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(b.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (b BusinessChatLinkInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(b.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (b BusinessChatLinkInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(b.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (b BusinessChatLinkInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(b.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (b BusinessChatLinkInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(b.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (b BusinessChatLinkInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(b.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (b BusinessChatLinkInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(b.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (b BusinessChatLinkInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(b.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (b BusinessChatLinkInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(b.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (b BusinessChatLinkInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(b.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (b BusinessChatLinkInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(b.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (b BusinessChatLinkInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(b.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (b BusinessChatLinkInfo) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(b.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (b BusinessChatLinkInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(b.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (b BusinessChatLinkInfo) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(b.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (b BusinessChatLinkInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(b.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (b BusinessChatLinkInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(b.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (b BusinessChatLinkInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(b.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (b BusinessChatLinkInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(b.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (b BusinessChatLinkInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(b.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (b BusinessChatLinkInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(b.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (b BusinessChatLinkInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(b.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (b BusinessChatLinkInfo) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(b.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (b BusinessChatLinkInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(b.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (b BusinessChatLinkInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(b.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (b BusinessChatLinkInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(b.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (b BusinessChatLinkInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(b.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (b BusinessChatLinkInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, b.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (b BusinessChatLinkInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(b.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (b BusinessChatLinkInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(b.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (b BusinessChatLinkInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(b.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (b BusinessChatLinkInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(b.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (b BusinessChatLinkInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(b.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (b BusinessChatLinkInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(b.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (b BusinessChatLinkInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(b.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (b BusinessChatLinkInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, b.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (b BusinessChatLinkInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(b.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (b BusinessChatLinkInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(b.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (b BusinessChatLinkInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(b.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (b BusinessChatLinkInfo) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(b.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (b BusinessChatLinkInfo) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(b.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (b BusinessChatLinkInfo) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(b.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (b BusinessChatLinkInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(b.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (b BusinessChatLinkInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(b.ChatId, messageId, toLanguageCode)
}

// TranslateText is a helper method for Client.TranslateText
func (b BusinessChatLinkInfo) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(b.Text, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (b BusinessChatLinkInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(b.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (b BusinessChatLinkInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(b.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (b BusinessChatLinkInfo) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(b.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (b BusinessChatLinkInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(b.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (b BusinessChatLinkInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(b.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (b BusinessChatLinkInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(b.ChatId, messageIds, forceRead, opts)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (b BusinessConnectedBot) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(b.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (b BusinessConnectedBot) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(b.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (b BusinessConnectedBot) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(b.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (b BusinessConnectedBot) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(b.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (b BusinessConnectedBot) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, b.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (b BusinessConnectedBot) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// Delete is a helper method for Client.DeleteBusinessConnectedBot
func (b BusinessConnectedBot) Delete(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(b.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (b BusinessConnectedBot) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(b.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (b BusinessConnectedBot) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(b.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (b BusinessConnectedBot) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(b.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (b BusinessConnectedBot) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(b.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (b BusinessConnectedBot) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(b.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (b BusinessConnectedBot) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(b.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (b BusinessConnectedBot) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(b.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (b BusinessConnectedBot) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(b.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (b BusinessConnectedBot) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(b.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (b BusinessConnectedBot) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, b.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (b BusinessConnectedBot) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(b.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (b BusinessConnectedBot) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, b.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (b BusinessConnectedBot) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(b.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (b BusinessConnectedBot) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(b.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (b BusinessConnectedBot) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, b.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (b BusinessConnectedBot) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(b.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (b BusinessConnectedBot) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(b.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (b BusinessConnectedBot) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(b.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (b BusinessConnectedBot) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, b.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (b BusinessConnectedBot) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(b.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (b BusinessConnectedBot) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(b.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (b BusinessConnectedBot) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(b.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (b BusinessConnectedBot) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(b.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (b BusinessConnectedBot) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(b.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (b BusinessConnectedBot) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(b.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (b BusinessConnectedBot) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(b.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (b BusinessConnectedBot) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(b.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (b BusinessConnectedBot) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(b.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (b BusinessConnectedBot) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(b.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (b BusinessConnectedBot) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(b.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (b BusinessConnectedBot) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(b.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (b BusinessConnectedBot) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(b.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (b BusinessConnectedBot) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(b.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (b BusinessConnectedBot) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(b.BotUserId, username, isActive)
}

// AddChatMember is a helper method for Client.AddChatMember
func (b BusinessConnection) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, b.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (b BusinessConnection) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(b.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (b BusinessConnection) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(b.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (b BusinessConnection) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(b.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (b BusinessConnection) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(b.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (b BusinessConnection) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(b.UserId, protocol, isVideo)
}

// CreateInvoiceLink is a helper method for Client.CreateInvoiceLink
func (b BusinessConnection) CreateInvoiceLink(client *Client, invoice *InputMessageContent) (*HttpUrl, error) {
	return client.CreateInvoiceLink(b.Id, invoice)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (b BusinessConnection) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(b.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (b BusinessConnection) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(b.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (b BusinessConnection) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(b.UserId, force)
}

// DeleteBusinessMessages is a helper method for Client.DeleteBusinessMessages
func (b BusinessConnection) DeleteBusinessMessages(client *Client, messageIds []int64) (*Ok, error) {
	return client.DeleteBusinessMessages(b.Id, messageIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (b BusinessConnection) DeleteBusinessStory(client *Client, storyId int32) (*Ok, error) {
	return client.DeleteBusinessStory(b.Id, storyId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (b BusinessConnection) EditBusinessMessageCaption(client *Client, chatId int64, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(b.Id, chatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (b BusinessConnection) EditBusinessMessageChecklist(client *Client, chatId int64, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(b.Id, chatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (b BusinessConnection) EditBusinessMessageLiveLocation(client *Client, chatId int64, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(b.Id, chatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (b BusinessConnection) EditBusinessMessageMedia(client *Client, chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(b.Id, chatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (b BusinessConnection) EditBusinessMessageReplyMarkup(client *Client, chatId int64, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(b.Id, chatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (b BusinessConnection) EditBusinessMessageText(client *Client, chatId int64, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(b.Id, chatId, messageId, inputMessageContent, opts)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (b BusinessConnection) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(b.UserId, telegramPaymentChargeId, isCanceled)
}

// GetBusinessAccountStarAmount is a helper method for Client.GetBusinessAccountStarAmount
func (b BusinessConnection) GetBusinessAccountStarAmount(client *Client) (*StarAmount, error) {
	return client.GetBusinessAccountStarAmount(b.Id)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (b BusinessConnection) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, b.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (b BusinessConnection) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, b.Date)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (b BusinessConnection) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, b.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (b BusinessConnection) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(b.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (b BusinessConnection) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, b.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (b BusinessConnection) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(b.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (b BusinessConnection) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(b.UserId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (b BusinessConnection) GetReceivedGifts(client *Client, ownerId *MessageSender, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(b.Id, ownerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (b BusinessConnection) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, b.Date)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (b BusinessConnection) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(b.UserId)
}

// GetUser is a helper method for Client.GetUser
func (b BusinessConnection) GetUser(client *Client) (*User, error) {
	return client.GetUser(b.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (b BusinessConnection) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, b.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (b BusinessConnection) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(b.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (b BusinessConnection) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(b.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (b BusinessConnection) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(b.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (b BusinessConnection) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(b.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (b BusinessConnection) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(b.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (b BusinessConnection) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, b.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (b BusinessConnection) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, b.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (b BusinessConnection) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, b.UserId, approve)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (b BusinessConnection) ReadBusinessMessage(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(b.Id, chatId, messageId)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (b BusinessConnection) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(b.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (b BusinessConnection) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(b.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (b BusinessConnection) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(b.UserId, result, chatTypes)
}

// SellGift is a helper method for Client.SellGift
func (b BusinessConnection) SellGift(client *Client, receivedGiftId string) (*Ok, error) {
	return client.SellGift(b.Id, receivedGiftId)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (b BusinessConnection) SendBusinessMessage(client *Client, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(b.Id, chatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (b BusinessConnection) SendBusinessMessageAlbum(client *Client, chatId int64, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(b.Id, chatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (b BusinessConnection) SendChatAction(client *Client, chatId int64, topicId *MessageTopic, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(chatId, topicId, b.Id, opts)
}

// SetBusinessAccountBio is a helper method for Client.SetBusinessAccountBio
func (b BusinessConnection) SetBusinessAccountBio(client *Client, bio string) (*Ok, error) {
	return client.SetBusinessAccountBio(b.Id, bio)
}

// SetBusinessAccountGiftSettings is a helper method for Client.SetBusinessAccountGiftSettings
func (b BusinessConnection) SetBusinessAccountGiftSettings(client *Client, settings *GiftSettings) (*Ok, error) {
	return client.SetBusinessAccountGiftSettings(b.Id, settings)
}

// SetBusinessAccountName is a helper method for Client.SetBusinessAccountName
func (b BusinessConnection) SetBusinessAccountName(client *Client, firstName string, lastName string) (*Ok, error) {
	return client.SetBusinessAccountName(b.Id, firstName, lastName)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (b BusinessConnection) SetBusinessAccountProfilePhoto(client *Client, isPublic bool, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(b.Id, isPublic, opts)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (b BusinessConnection) SetBusinessAccountUsername(client *Client, username string) (*Ok, error) {
	return client.SetBusinessAccountUsername(b.Id, username)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (b BusinessConnection) SetBusinessMessageIsPinned(client *Client, chatId int64, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(b.Id, chatId, messageId, isPinned)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (b BusinessConnection) SetChatDirectMessagesGroup(client *Client, chatId int64, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, b.IsEnabled, paidMessageStarCount)
}

// SetGameScore is a helper method for Client.SetGameScore
func (b BusinessConnection) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, b.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (b BusinessConnection) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, b.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (b BusinessConnection) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(b.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (b BusinessConnection) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(b.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (b BusinessConnection) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(b.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (b BusinessConnection) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(b.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (b BusinessConnection) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(b.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (b BusinessConnection) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(b.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (b BusinessConnection) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(b.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (b BusinessConnection) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(b.UserId)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (b BusinessConnection) StopBusinessPoll(client *Client, chatId int64, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(b.Id, chatId, messageId, opts)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (b BusinessConnection) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(b.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (b BusinessConnection) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(b.UserId, photo)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (b BusinessConnection) TransferBusinessAccountStars(client *Client, starCount int64) (*Ok, error) {
	return client.TransferBusinessAccountStars(b.Id, starCount)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (b BusinessConnection) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, b.UserId, password)
}

// TransferGift is a helper method for Client.TransferGift
func (b BusinessConnection) TransferGift(client *Client, receivedGiftId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	return client.TransferGift(b.Id, receivedGiftId, newOwnerId, starCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (b BusinessConnection) UpgradeGift(client *Client, receivedGiftId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(b.Id, receivedGiftId, keepOriginalDetails, starCount)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (b BusinessConnection) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(b.UserId, stickerFormat, sticker)
}

// DeleteQuickReplyShortcut is a helper method for Client.DeleteQuickReplyShortcut
func (b BusinessGreetingMessageSettings) DeleteQuickReplyShortcut(client *Client) (*Ok, error) {
	return client.DeleteQuickReplyShortcut(b.ShortcutId)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (b BusinessGreetingMessageSettings) DeleteQuickReplyShortcutMessages(client *Client, messageIds []int64) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(b.ShortcutId, messageIds)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (b BusinessGreetingMessageSettings) EditQuickReplyMessage(client *Client, messageId int64, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(b.ShortcutId, messageId, inputMessageContent)
}

// LoadQuickReplyShortcutMessages is a helper method for Client.LoadQuickReplyShortcutMessages
func (b BusinessGreetingMessageSettings) LoadQuickReplyShortcutMessages(client *Client) (*Ok, error) {
	return client.LoadQuickReplyShortcutMessages(b.ShortcutId)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (b BusinessGreetingMessageSettings) SendQuickReplyShortcutMessages(client *Client, chatId int64, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, b.ShortcutId, sendingId)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (b BusinessGreetingMessageSettings) SetQuickReplyShortcutName(client *Client, name string) (*Ok, error) {
	return client.SetQuickReplyShortcutName(b.ShortcutId, name)
}

// SetBusinessAwayMessageSettings is a helper method for Client.SetBusinessAwayMessageSettings
func (b BusinessInfo) SetBusinessAwayMessageSettings(client *Client) (*Ok, error) {
	return client.SetBusinessAwayMessageSettings(b.AwayMessageSettings)
}

// SetBusinessGreetingMessageSettings is a helper method for Client.SetBusinessGreetingMessageSettings
func (b BusinessInfo) SetBusinessGreetingMessageSettings(client *Client) (*Ok, error) {
	return client.SetBusinessGreetingMessageSettings(b.GreetingMessageSettings)
}

// SetBusinessLocation is a helper method for Client.SetBusinessLocation
func (b BusinessInfo) SetBusinessLocation(client *Client) (*Ok, error) {
	return client.SetBusinessLocation(b.Location)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (b BusinessLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(b.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (b BusinessLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(b.Location, zoom, width, height, scale, chatId)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (b BusinessRecipients) AddChatFolderByInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.AddChatFolderByInviteLink(inviteLink, b.ChatIds)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (b BusinessRecipients) CreateChatFolderInviteLink(client *Client, chatFolderId int32, name string) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, name, b.ChatIds)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (b BusinessRecipients) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, name string) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, name, b.ChatIds)
}

// OptimizeStorage is a helper method for Client.OptimizeStorage
func (b BusinessRecipients) OptimizeStorage(client *Client, size int64, ttl int32, count int32, immunityDelay int32, fileTypes []*FileType, excludeChatIds []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	return client.OptimizeStorage(size, ttl, count, immunityDelay, fileTypes, b.ChatIds, excludeChatIds, returnDeletedFileStatistics, chatLimit)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (b BusinessRecipients) SetPinnedChats(client *Client, chatList *ChatList) (*Ok, error) {
	return client.SetPinnedChats(chatList, b.ChatIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (b BusinessStartPage) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, b.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (b BusinessStartPage) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, b.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (b BusinessStartPage) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(b.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (b BusinessStartPage) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, b.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (b BusinessStartPage) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(b.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (b BusinessStartPage) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, b.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (b BusinessStartPage) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, b.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (b BusinessStartPage) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, b.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (b BusinessStartPage) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, b.Title, recordVideo, usePortraitOrientation)
}

// Accept is a helper method for Client.AcceptCall
func (c Call) Accept(client *Client, protocol *CallProtocol) (*Ok, error) {
	return client.AcceptCall(c.Id, protocol)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c Call) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, c.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (c Call) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(c.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (c Call) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(c.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (c Call) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(c.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (c Call) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(c.UserId, onlyLocal)
}

// Create is a helper method for Client.CreateCall
func (c Call) Create(client *Client, protocol *CallProtocol) (*CallId, error) {
	return client.CreateCall(c.UserId, protocol, c.IsVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (c Call) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(c.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c Call) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(c.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (c Call) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(c.UserId, force)
}

// Discard is a helper method for Client.DiscardCall
func (c Call) Discard(client *Client, isDisconnected bool, inviteLink string, duration int32, connectionId string) (*Ok, error) {
	return client.DiscardCall(c.Id, isDisconnected, inviteLink, duration, c.IsVideo, connectionId)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (c Call) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(c.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c Call) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, c.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (c Call) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(c.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (c Call) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, c.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (c Call) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(c.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (c Call) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(c.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (c Call) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(c.UserId)
}

// GetUser is a helper method for Client.GetUser
func (c Call) GetUser(client *Client) (*User, error) {
	return client.GetUser(c.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c Call) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, c.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (c Call) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(c.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (c Call) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(c.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (c Call) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(c.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (c Call) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(c.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (c Call) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(c.UserId, starCount, monthCount, text)
}

// InviteGroupParticipant is a helper method for Client.InviteGroupCallParticipant
func (c Call) InviteGroupParticipant(client *Client, groupCallId int32) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, c.UserId, c.IsVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (c Call) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, c.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c Call) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, c.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (c Call) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(c.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (c Call) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(c.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (c Call) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(c.UserId, result, chatTypes)
}

// SendDebugInformation is a helper method for Client.SendCallDebugInformation
func (c Call) SendDebugInformation(client *Client, debugInformation string) (*Ok, error) {
	return client.SendCallDebugInformation(c.Id, debugInformation)
}

// SendLog is a helper method for Client.SendCallLog
func (c Call) SendLog(client *Client, logFile *InputFile) (*Ok, error) {
	return client.SendCallLog(c.Id, logFile)
}

// SendRating is a helper method for Client.SendCallRating
func (c Call) SendRating(client *Client, rating int32, comment string, problems []*CallProblem) (*Ok, error) {
	return client.SendCallRating(c.Id, rating, comment, problems)
}

// SendSignalingData is a helper method for Client.SendCallSignalingData
func (c Call) SendSignalingData(client *Client, data string) (*Ok, error) {
	return client.SendCallSignalingData(c.Id, data)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c Call) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, c.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (c Call) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, c.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (c Call) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(c.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (c Call) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(c.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (c Call) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(c.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c Call) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(c.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (c Call) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(c.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (c Call) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(c.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (c Call) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(c.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (c Call) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(c.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (c Call) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(c.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (c Call) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(c.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c Call) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, c.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (c Call) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(c.UserId, stickerFormat, sticker)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (c CallbackQueryAnswer) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, c.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (c CallbackQueryAnswer) AnswerCallbackQuery(client *Client, callbackQueryId string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, c.Text, c.ShowAlert, c.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (c CallbackQueryAnswer) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, c.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (c CallbackQueryAnswer) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, c.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (c CallbackQueryAnswer) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(c.Url)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (c CallbackQueryAnswer) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(c.Text, inputLanguageCodes)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (c CallbackQueryAnswer) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(c.Url)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (c CallbackQueryAnswer) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(c.Text)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (c CallbackQueryAnswer) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, c.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (c CallbackQueryAnswer) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(c.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c CallbackQueryAnswer) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, c.Url, parameters, opts)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (c CallbackQueryAnswer) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(c.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (c CallbackQueryAnswer) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, c.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (c CallbackQueryAnswer) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, c.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (c CallbackQueryAnswer) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, c.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (c CallbackQueryAnswer) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(c.Text, inputLanguageCodes)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c CallbackQueryPayloadData) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, c.Data, opts)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (c CallbackQueryPayloadData) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, c.Data, unencryptedPrefixSize)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (c CallbackQueryPayloadData) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, c.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (c CallbackQueryPayloadData) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, c.Data)
}

// CheckAuthenticationPassword is a helper method for Client.CheckAuthenticationPassword
func (c CallbackQueryPayloadDataWithPassword) CheckAuthenticationPassword(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPassword(c.Password)
}

// CreateTemporaryPassword is a helper method for Client.CreateTemporaryPassword
func (c CallbackQueryPayloadDataWithPassword) CreateTemporaryPassword(client *Client, validFor int32) (*TemporaryPasswordState, error) {
	return client.CreateTemporaryPassword(c.Password, validFor)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (c CallbackQueryPayloadDataWithPassword) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, c.Data, opts)
}

// DeleteAccount is a helper method for Client.DeleteAccount
func (c CallbackQueryPayloadDataWithPassword) DeleteAccount(client *Client, reason string) (*Ok, error) {
	return client.DeleteAccount(reason, c.Password)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (c CallbackQueryPayloadDataWithPassword) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, c.Data, unencryptedPrefixSize)
}

// GetAllPassportElements is a helper method for Client.GetAllPassportElements
func (c CallbackQueryPayloadDataWithPassword) GetAllPassportElements(client *Client) (*PassportElements, error) {
	return client.GetAllPassportElements(c.Password)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c CallbackQueryPayloadDataWithPassword) GetChatRevenueWithdrawalUrl(client *Client, chatId int64) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(chatId, c.Password)
}

// GetPassportAuthorizationFormAvailableElements is a helper method for Client.GetPassportAuthorizationFormAvailableElements
func (c CallbackQueryPayloadDataWithPassword) GetPassportAuthorizationFormAvailableElements(client *Client, authorizationFormId int32) (*PassportElementsWithErrors, error) {
	return client.GetPassportAuthorizationFormAvailableElements(authorizationFormId, c.Password)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (c CallbackQueryPayloadDataWithPassword) GetPassportElement(client *Client, typeField *PassportElementType) (*PassportElement, error) {
	return client.GetPassportElement(typeField, c.Password)
}

// GetRecoveryEmailAddress is a helper method for Client.GetRecoveryEmailAddress
func (c CallbackQueryPayloadDataWithPassword) GetRecoveryEmailAddress(client *Client) (*RecoveryEmailAddress, error) {
	return client.GetRecoveryEmailAddress(c.Password)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (c CallbackQueryPayloadDataWithPassword) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, starCount int64) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, starCount, c.Password)
}

// GetTonWithdrawalUrl is a helper method for Client.GetTonWithdrawalUrl
func (c CallbackQueryPayloadDataWithPassword) GetTonWithdrawalUrl(client *Client) (*HttpUrl, error) {
	return client.GetTonWithdrawalUrl(c.Password)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (c CallbackQueryPayloadDataWithPassword) GetUpgradedGiftWithdrawalUrl(client *Client, receivedGiftId string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(receivedGiftId, c.Password)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (c CallbackQueryPayloadDataWithPassword) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, c.Data)
}

// SetPassportElement is a helper method for Client.SetPassportElement
func (c CallbackQueryPayloadDataWithPassword) SetPassportElement(client *Client, element *InputPassportElement) (*PassportElement, error) {
	return client.SetPassportElement(element, c.Password)
}

// SetRecoveryEmailAddress is a helper method for Client.SetRecoveryEmailAddress
func (c CallbackQueryPayloadDataWithPassword) SetRecoveryEmailAddress(client *Client, newRecoveryEmailAddress string) (*PasswordState, error) {
	return client.SetRecoveryEmailAddress(c.Password, newRecoveryEmailAddress)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c CallbackQueryPayloadDataWithPassword) TransferChatOwnership(client *Client, chatId int64, userId int64) (*Ok, error) {
	return client.TransferChatOwnership(chatId, userId, c.Password)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (c CallbackQueryPayloadDataWithPassword) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, c.Data)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(c.InviteLink, chatIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(c.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(c.InviteLink)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, c.InviteLink)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, c.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (c CallDiscardReasonUpgradeToGroupCall) DiscardCall(client *Client, callId int32, isDisconnected bool, duration int32, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, c.InviteLink, duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, c.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, c.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, c.InviteLink, name)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, c.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c CallDiscardReasonUpgradeToGroupCall) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, c.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c CallDiscardReasonUpgradeToGroupCall) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, c.InviteLink, query, limit, opts)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(c.InviteLink)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c CallDiscardReasonUpgradeToGroupCall) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, c.InviteLink, approve)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (c CallDiscardReasonUpgradeToGroupCall) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, c.InviteLink)
}

// AddProxy is a helper method for Client.AddProxy
func (c CallServer) AddProxy(client *Client, server string, enable bool, typeField *ProxyType) (*Proxy, error) {
	return client.AddProxy(server, c.Port, enable, typeField)
}

// EditProxy is a helper method for Client.EditProxy
func (c CallServer) EditProxy(client *Client, proxyId int32, server string, enable bool, typeField *ProxyType) (*Proxy, error) {
	return client.EditProxy(proxyId, server, c.Port, enable, typeField)
}

// TestProxy is a helper method for Client.TestProxy
func (c CallServer) TestProxy(client *Client, server string, typeField *ProxyType, dcId int32, timeout float64) (*Ok, error) {
	return client.TestProxy(server, c.Port, typeField, dcId, timeout)
}

// CheckAuthenticationPassword is a helper method for Client.CheckAuthenticationPassword
func (c CallServerTypeWebrtc) CheckAuthenticationPassword(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPassword(c.Password)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c CallServerTypeWebrtc) CheckChatUsername(client *Client, chatId int64) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(chatId, c.Username)
}

// CreateTemporaryPassword is a helper method for Client.CreateTemporaryPassword
func (c CallServerTypeWebrtc) CreateTemporaryPassword(client *Client, validFor int32) (*TemporaryPasswordState, error) {
	return client.CreateTemporaryPassword(c.Password, validFor)
}

// DeleteAccount is a helper method for Client.DeleteAccount
func (c CallServerTypeWebrtc) DeleteAccount(client *Client, reason string) (*Ok, error) {
	return client.DeleteAccount(reason, c.Password)
}

// GetAllPassportElements is a helper method for Client.GetAllPassportElements
func (c CallServerTypeWebrtc) GetAllPassportElements(client *Client) (*PassportElements, error) {
	return client.GetAllPassportElements(c.Password)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c CallServerTypeWebrtc) GetChatRevenueWithdrawalUrl(client *Client, chatId int64) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(chatId, c.Password)
}

// GetPassportAuthorizationFormAvailableElements is a helper method for Client.GetPassportAuthorizationFormAvailableElements
func (c CallServerTypeWebrtc) GetPassportAuthorizationFormAvailableElements(client *Client, authorizationFormId int32) (*PassportElementsWithErrors, error) {
	return client.GetPassportAuthorizationFormAvailableElements(authorizationFormId, c.Password)
}

// GetPassportElement is a helper method for Client.GetPassportElement
func (c CallServerTypeWebrtc) GetPassportElement(client *Client, typeField *PassportElementType) (*PassportElement, error) {
	return client.GetPassportElement(typeField, c.Password)
}

// GetRecoveryEmailAddress is a helper method for Client.GetRecoveryEmailAddress
func (c CallServerTypeWebrtc) GetRecoveryEmailAddress(client *Client) (*RecoveryEmailAddress, error) {
	return client.GetRecoveryEmailAddress(c.Password)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (c CallServerTypeWebrtc) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, starCount int64) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, starCount, c.Password)
}

// GetTonWithdrawalUrl is a helper method for Client.GetTonWithdrawalUrl
func (c CallServerTypeWebrtc) GetTonWithdrawalUrl(client *Client) (*HttpUrl, error) {
	return client.GetTonWithdrawalUrl(c.Password)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (c CallServerTypeWebrtc) GetUpgradedGiftWithdrawalUrl(client *Client, receivedGiftId string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(receivedGiftId, c.Password)
}

// SearchChatAffiliateProgram is a helper method for Client.SearchChatAffiliateProgram
func (c CallServerTypeWebrtc) SearchChatAffiliateProgram(client *Client, referrer string) (*Chat, error) {
	return client.SearchChatAffiliateProgram(c.Username, referrer)
}

// SearchPublicChat is a helper method for Client.SearchPublicChat
func (c CallServerTypeWebrtc) SearchPublicChat(client *Client) (*Chat, error) {
	return client.SearchPublicChat(c.Username)
}

// SetBusinessAccountUsername is a helper method for Client.SetBusinessAccountUsername
func (c CallServerTypeWebrtc) SetBusinessAccountUsername(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountUsername(businessConnectionId, c.Username)
}

// SetPassportElement is a helper method for Client.SetPassportElement
func (c CallServerTypeWebrtc) SetPassportElement(client *Client, element *InputPassportElement) (*PassportElement, error) {
	return client.SetPassportElement(element, c.Password)
}

// SetRecoveryEmailAddress is a helper method for Client.SetRecoveryEmailAddress
func (c CallServerTypeWebrtc) SetRecoveryEmailAddress(client *Client, newRecoveryEmailAddress string) (*PasswordState, error) {
	return client.SetRecoveryEmailAddress(c.Password, newRecoveryEmailAddress)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (c CallServerTypeWebrtc) SetSupergroupUsername(client *Client, supergroupId int64) (*Ok, error) {
	return client.SetSupergroupUsername(supergroupId, c.Username)
}

// SetUsername is a helper method for Client.SetUsername
func (c CallServerTypeWebrtc) SetUsername(client *Client) (*Ok, error) {
	return client.SetUsername(c.Username)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (c CallServerTypeWebrtc) ToggleBotUsernameIsActive(client *Client, botUserId int64, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, c.Username, isActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (c CallServerTypeWebrtc) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, c.Username, isActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (c CallServerTypeWebrtc) ToggleUsernameIsActive(client *Client, isActive bool) (*Ok, error) {
	return client.ToggleUsernameIsActive(c.Username, isActive)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (c CallServerTypeWebrtc) TransferChatOwnership(client *Client, chatId int64, userId int64) (*Ok, error) {
	return client.TransferChatOwnership(chatId, userId, c.Password)
}

// TestReturnError is a helper method for Client.TestReturnError
func (c CallStateError) TestReturnError(client *Client) (*Error, error) {
	return client.TestReturnError(c.Error)
}

// AcceptCall is a helper method for Client.AcceptCall
func (c CallStateReady) AcceptCall(client *Client, callId int32) (*Ok, error) {
	return client.AcceptCall(callId, c.Protocol)
}

// CreateCall is a helper method for Client.CreateCall
func (c CallStateReady) CreateCall(client *Client, userId int64, isVideo bool) (*CallId, error) {
	return client.CreateCall(userId, c.Protocol, isVideo)
}

// CloseStory is a helper method for Client.CloseStory
func (c CanPostStoryResultLiveStoryIsActive) CloseStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.CloseStory(storyPosterChatId, c.StoryId)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (c CanPostStoryResultLiveStoryIsActive) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, c.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (c CanPostStoryResultLiveStoryIsActive) DeleteStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.DeleteStory(storyPosterChatId, c.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (c CanPostStoryResultLiveStoryIsActive) EditBusinessStory(client *Client, storyPosterChatId int64, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, c.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (c CanPostStoryResultLiveStoryIsActive) EditStory(client *Client, storyPosterChatId int64, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(storyPosterChatId, c.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (c CanPostStoryResultLiveStoryIsActive) EditStoryCover(client *Client, storyPosterChatId int64, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(storyPosterChatId, c.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (c CanPostStoryResultLiveStoryIsActive) GetChatStoryInteractions(client *Client, storyPosterChatId int64, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, c.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (c CanPostStoryResultLiveStoryIsActive) GetStory(client *Client, storyPosterChatId int64, onlyLocal bool) (*Story, error) {
	return client.GetStory(storyPosterChatId, c.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (c CanPostStoryResultLiveStoryIsActive) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(c.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (c CanPostStoryResultLiveStoryIsActive) GetStoryPublicForwards(client *Client, storyPosterChatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, c.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c CanPostStoryResultLiveStoryIsActive) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, c.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (c CanPostStoryResultLiveStoryIsActive) OpenStory(client *Client, storyPosterChatId int64) (*Ok, error) {
	return client.OpenStory(storyPosterChatId, c.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (c CanPostStoryResultLiveStoryIsActive) ReportStory(client *Client, storyPosterChatId int64, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, c.StoryId, optionId, text)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (c CanPostStoryResultLiveStoryIsActive) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(c.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (c CanPostStoryResultLiveStoryIsActive) SetStoryReaction(client *Client, storyPosterChatId int64, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(storyPosterChatId, c.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (c CanPostStoryResultLiveStoryIsActive) ToggleStoryIsPostedToChatPage(client *Client, storyPosterChatId int64, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(storyPosterChatId, c.StoryId, isPostedToChatPage)
}

// AddMember is a helper method for Client.AddChatMember
func (c Chat) AddMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(c.Id, userId, forwardLimit)
}

// AddMembers is a helper method for Client.AddChatMembers
func (c Chat) AddMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(c.Id, userIds)
}

// AddToList is a helper method for Client.AddChatToList
func (c Chat) AddToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(c.Id, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c Chat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(c.Id, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c Chat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, c.Id, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (c Chat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(c.Id, senderId, disableNotification, inputMessageContent, opts)
}

// AddLoginPasskey is a helper method for Client.AddLoginPasskey
func (c Chat) AddLoginPasskey(client *Client, attestationObject string) (*Passkey, error) {
	return client.AddLoginPasskey(c.ClientData, attestationObject)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c Chat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(c.Id, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c Chat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(c.Id, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c Chat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(c.Id, messageId, starCount, opts)
}

// AddRecentlyFound is a helper method for Client.AddRecentlyFoundChat
func (c Chat) AddRecentlyFound(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(c.Id)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (c Chat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(c.Id, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c Chat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(c.Id, messageId, sendDate)
}

// BanMember is a helper method for Client.BanChatMember
func (c Chat) BanMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(c.Id, memberId, bannedUntilDate, revokeMessages)
}

// Boost is a helper method for Client.BoostChat
func (c Chat) Boost(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(c.Id, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (c Chat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(c.Id)
}

// CheckAuthenticationPasskey is a helper method for Client.CheckAuthenticationPasskey
func (c Chat) CheckAuthenticationPasskey(client *Client, credentialId string, authenticatorData string, signature string, userHandle string) (*Ok, error) {
	return client.CheckAuthenticationPasskey(credentialId, c.ClientData, authenticatorData, signature, userHandle)
}

// CheckUsername is a helper method for Client.CheckChatUsername
func (c Chat) CheckUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(c.Id, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c Chat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(c.Id, messageId)
}

// ClickSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c Chat) ClickSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(c.Id, messageId, isMediaClick, fromFullscreen)
}

// Close is a helper method for Client.CloseChat
func (c Chat) Close(client *Client) (*Ok, error) {
	return client.CloseChat(c.Id)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c Chat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(c.Id, messageId)
}

// CreateInviteLink is a helper method for Client.CreateChatInviteLink
func (c Chat) CreateInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(c.Id, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c Chat) CreateSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(c.Id, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c Chat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(c.Id, name, isNameImplicit, icon)
}

// CreateNewBasicGroup is a helper method for Client.CreateNewBasicGroupChat
func (c Chat) CreateNewBasicGroup(client *Client, userIds []int64) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, c.Title, c.MessageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c Chat) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, c.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroup is a helper method for Client.CreateNewSupergroupChat
func (c Chat) CreateNewSupergroup(client *Client, isForum bool, isChannel bool, description string, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(c.Title, isForum, isChannel, description, c.MessageAutoDeleteTime, forImport, opts)
}

// CreateVideo is a helper method for Client.CreateVideoChat
func (c Chat) CreateVideo(client *Client, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(c.Id, c.Title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c Chat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(c.Id, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c Chat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(c.Id, messageId, comment)
}

// DeleteAllRevokedInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c Chat) DeleteAllRevokedInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(c.Id, creatorUserId)
}

// Delete is a helper method for Client.DeleteChat
func (c Chat) Delete(client *Client) (*Ok, error) {
	return client.DeleteChat(c.Id)
}

// DeleteBackground is a helper method for Client.DeleteChatBackground
func (c Chat) DeleteBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(c.Id, restorePrevious)
}

// DeleteHistory is a helper method for Client.DeleteChatHistory
func (c Chat) DeleteHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(c.Id, removeFromChatList, revoke)
}

// DeleteMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (c Chat) DeleteMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(c.Id, minDate, maxDate, revoke)
}

// DeleteMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (c Chat) DeleteMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(c.Id, senderId)
}

// DeleteReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c Chat) DeleteReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(c.Id, messageId)
}

// DeleteDirectMessagesTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (c Chat) DeleteDirectMessagesTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(c.Id, topicId)
}

// DeleteDirectMessagesTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (c Chat) DeleteDirectMessagesTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(c.Id, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (c Chat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(c.Id, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (c Chat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(c.Id, messageIds, revoke)
}

// DeleteRevokedInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c Chat) DeleteRevokedInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(c.Id, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (c Chat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(c.Id, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c Chat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, c.Id, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c Chat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, c.Id, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c Chat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, c.Id, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c Chat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, c.Id, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c Chat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, c.Id, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c Chat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, c.Id, messageId, inputMessageContent, opts)
}

// EditInviteLink is a helper method for Client.EditChatInviteLink
func (c Chat) EditInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(c.Id, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c Chat) EditSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(c.Id, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c Chat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(c.Id, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c Chat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(c.Id, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c Chat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(c.Id, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c Chat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(c.Id, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c Chat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(c.Id, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c Chat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(c.Id, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c Chat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(c.Id, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c Chat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(c.Id, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (c Chat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(c.Id, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (c Chat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, c.Id, returnOnlyMainEmoji)
}

// GetBlockedMessageSenders is a helper method for Client.GetBlockedMessageSenders
func (c Chat) GetBlockedMessageSenders(client *Client, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetBlockedMessageSenders(c.BlockList, offset, limit)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c Chat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(c.Id, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c Chat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(c.Id, messageId, callbackQueryId)
}

// Get is a helper method for Client.GetChat
func (c Chat) Get(client *Client) (*Chat, error) {
	return client.GetChat(c.Id)
}

// GetActiveStories is a helper method for Client.GetChatActiveStories
func (c Chat) GetActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(c.Id)
}

// GetAdministrators is a helper method for Client.GetChatAdministrators
func (c Chat) GetAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(c.Id)
}

// GetArchivedStories is a helper method for Client.GetChatArchivedStories
func (c Chat) GetArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(c.Id, fromStoryId, limit)
}

// GetAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (c Chat) GetAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(c.Id)
}

// GetAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (c Chat) GetAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(c.Id)
}

// GetBoostLink is a helper method for Client.GetChatBoostLink
func (c Chat) GetBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(c.Id)
}

// GetBoosts is a helper method for Client.GetChatBoosts
func (c Chat) GetBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(c.Id, onlyGiftCodes, offset, limit)
}

// GetBoostStatus is a helper method for Client.GetChatBoostStatus
func (c Chat) GetBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(c.Id)
}

// GetEventLog is a helper method for Client.GetChatEventLog
func (c Chat) GetEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(c.Id, query, fromEventId, limit, userIds, opts)
}

// GetHistory is a helper method for Client.GetChatHistory
func (c Chat) GetHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(c.Id, fromMessageId, offset, limit, onlyLocal)
}

// GetInviteLink is a helper method for Client.GetChatInviteLink
func (c Chat) GetInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(c.Id, inviteLink)
}

// GetInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (c Chat) GetInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(c.Id)
}

// GetInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c Chat) GetInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(c.Id, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetInviteLinks is a helper method for Client.GetChatInviteLinks
func (c Chat) GetInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(c.Id, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetJoinRequests is a helper method for Client.GetChatJoinRequests
func (c Chat) GetJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(c.Id, inviteLink, query, limit, opts)
}

// GetListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (c Chat) GetListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(c.Id)
}

// GetMember is a helper method for Client.GetChatMember
func (c Chat) GetMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(c.Id, memberId)
}

// GetMessageByDate is a helper method for Client.GetChatMessageByDate
func (c Chat) GetMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(c.Id, date)
}

// GetMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (c Chat) GetMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(c.Id, filter, fromMessageId, opts)
}

// GetMessageCount is a helper method for Client.GetChatMessageCount
func (c Chat) GetMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(c.Id, filter, returnLocal, opts)
}

// GetMessagePosition is a helper method for Client.GetChatMessagePosition
func (c Chat) GetMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(c.Id, filter, messageId, opts)
}

// GetPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (c Chat) GetPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(c.Id)
}

// GetPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (c Chat) GetPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(c.Id, fromStoryId, limit)
}

// GetRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (c Chat) GetRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(c.Id, isDark)
}

// GetRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (c Chat) GetRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(c.Id, offset, limit)
}

// GetRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c Chat) GetRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(c.Id, password)
}

// GetScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (c Chat) GetScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(c.Id)
}

// GetSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (c Chat) GetSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(c.Id, returnLocal)
}

// GetSimilarChats is a helper method for Client.GetChatSimilarChats
func (c Chat) GetSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(c.Id)
}

// GetSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (c Chat) GetSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(c.Id, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (c Chat) GetSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(c.Id)
}

// GetStatistics is a helper method for Client.GetChatStatistics
func (c Chat) GetStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(c.Id, isDark)
}

// GetStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (c Chat) GetStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(c.Id)
}

// GetDirectMessagesTopic is a helper method for Client.GetDirectMessagesChatTopic
func (c Chat) GetDirectMessagesTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(c.Id, topicId)
}

// GetDirectMessagesTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (c Chat) GetDirectMessagesTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(c.Id, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c Chat) GetDirectMessagesTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(c.Id, topicId, date)
}

// GetDirectMessagesTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (c Chat) GetDirectMessagesTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(c.Id, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (c Chat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(c.Id, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (c Chat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(c.Id, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (c Chat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(c.Id, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (c Chat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(c.Id, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c Chat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(c.Id, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c Chat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(c.Id, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c Chat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, c.Id, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (c Chat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(c.Id)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c Chat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(c.Id, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c Chat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(c.Id, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c Chat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(c.Id, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c Chat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, c.Id)
}

// GetMessage is a helper method for Client.GetMessage
func (c Chat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(c.Id, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c Chat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(c.Id, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c Chat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(c.Id, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c Chat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(c.Id, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c Chat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(c.Id, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (c Chat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(c.Id)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c Chat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(c.Id, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c Chat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(c.Id, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c Chat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(c.Id, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c Chat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(c.Id, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c Chat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(c.Id, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (c Chat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(c.Id, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c Chat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(c.Id, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c Chat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(c.Id, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c Chat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(c.Id, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c Chat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(c.Id, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c Chat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(c.Id, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c Chat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(c.Id, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c Chat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(c.Id, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (c Chat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(c.Id, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (c Chat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, c.Id)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (c Chat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(c.Id, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c Chat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(c.Id, storyId, isDark)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (c Chat) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(c.Title)
}

// GetUserBoosts is a helper method for Client.GetUserChatBoosts
func (c Chat) GetUserBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(c.Id, userId)
}

// GetVideoAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (c Chat) GetVideoAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(c.Id)
}

// GetVideoRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (c Chat) GetVideoRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(c.Id)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c Chat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(c.Id, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c Chat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(c.Id, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (c Chat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(c.Id, messageFile, attachedFiles)
}

// Join is a helper method for Client.JoinChat
func (c Chat) Join(client *Client) (*Ok, error) {
	return client.JoinChat(c.Id)
}

// Leave is a helper method for Client.LeaveChat
func (c Chat) Leave(client *Client) (*Ok, error) {
	return client.LeaveChat(c.Id)
}

// LoadDirectMessagesTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (c Chat) LoadDirectMessagesTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(c.Id, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (c Chat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(c.Id, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// Open is a helper method for Client.OpenChat
func (c Chat) Open(client *Client) (*Ok, error) {
	return client.OpenChat(c.Id)
}

// OpenSimilarChat is a helper method for Client.OpenChatSimilarChat
func (c Chat) OpenSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(c.Id, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (c Chat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(c.Id, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (c Chat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(c.Id, botUserId, url, parameters, opts)
}

// PinMessage is a helper method for Client.PinChatMessage
func (c Chat) PinMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(c.Id, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (c Chat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(c.Id, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (c Chat) ProcessJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(c.Id, userId, approve)
}

// ProcessJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (c Chat) ProcessJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(c.Id, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (c Chat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(c.Id, messageId, isGood)
}

// ReadAllMentions is a helper method for Client.ReadAllChatMentions
func (c Chat) ReadAllMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(c.Id)
}

// ReadAllReactions is a helper method for Client.ReadAllChatReactions
func (c Chat) ReadAllReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(c.Id)
}

// ReadAllDirectMessagesTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (c Chat) ReadAllDirectMessagesTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(c.Id, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (c Chat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(c.Id, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (c Chat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(c.Id, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (c Chat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, c.Id, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (c Chat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(c.Id, messageId)
}

// RemoveBusinessConnectedBotFrom is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (c Chat) RemoveBusinessConnectedBotFrom(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(c.Id)
}

// RemoveActionBar is a helper method for Client.RemoveChatActionBar
func (c Chat) RemoveActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(c.Id)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (c Chat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(c.Id, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (c Chat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(c.Id, messageId)
}

// RemoveRecentlyFound is a helper method for Client.RemoveRecentlyFoundChat
func (c Chat) RemoveRecentlyFound(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(c.Id)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (c Chat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(c.Id, storyAlbumId, storyIds)
}

// RemoveTop is a helper method for Client.RemoveTopChat
func (c Chat) RemoveTop(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, c.Id)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (c Chat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(c.Id, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (c Chat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(c.Id, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (c Chat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(c.Id)
}

// ReplacePrimaryInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (c Chat) ReplacePrimaryInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(c.Id)
}

// ReplaceVideoRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (c Chat) ReplaceVideoRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(c.Id)
}

// Report is a helper method for Client.ReportChat
func (c Chat) Report(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(c.Id, optionId, messageIds, text)
}

// ReportPhoto is a helper method for Client.ReportChatPhoto
func (c Chat) ReportPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(c.Id, fileId, reason, text)
}

// ReportSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (c Chat) ReportSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(c.Id, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (c Chat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(c.Id, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (c Chat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(c.Id, messageIds, paidMessageStarCount, opts)
}

// RevokeInviteLink is a helper method for Client.RevokeChatInviteLink
func (c Chat) RevokeInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(c.Id, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (c Chat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, c.Id, data)
}

// SearchMembers is a helper method for Client.SearchChatMembers
func (c Chat) SearchMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(c.Id, query, limit, opts)
}

// SearchMessages is a helper method for Client.SearchChatMessages
func (c Chat) SearchMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(c.Id, query, fromMessageId, offset, limit, opts)
}

// SearchRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (c Chat) SearchRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(c.Id, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (c Chat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(c.Id, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (c Chat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, c.Id, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (c Chat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, c.Id, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (c Chat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, c.Id, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendAction is a helper method for Client.SendChatAction
func (c Chat) SendAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(c.Id, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (c Chat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(c.Id, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (c Chat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(c.Id, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (c Chat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(c.Id, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (c Chat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(c.Id, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (c Chat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(c.Id, forumTopicId, draftId, text)
}

// SetAccentColor is a helper method for Client.SetAccentColor
func (c Chat) SetAccentColor(client *Client) (*Ok, error) {
	return client.SetAccentColor(c.AccentColorId, c.BackgroundCustomEmojiId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (c Chat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, c.Id, messageId, isPinned)
}

// SetAccentColor is a helper method for Client.SetChatAccentColor
func (c Chat) SetAccentColor(client *Client) (*Ok, error) {
	return client.SetChatAccentColor(c.Id, c.AccentColorId, c.BackgroundCustomEmojiId)
}

// SetActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (c Chat) SetActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(c.Id, storyList)
}

// SetAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (c Chat) SetAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(c.Id, opts)
}

// SetAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (c Chat) SetAvailableReactions(client *Client) (*Ok, error) {
	return client.SetChatAvailableReactions(c.Id, c.AvailableReactions)
}

// SetBackground is a helper method for Client.SetChatBackground
func (c Chat) SetBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(c.Id, darkThemeDimming, onlyForSelf, opts)
}

// SetClientData is a helper method for Client.SetChatClientData
func (c Chat) SetClientData(client *Client) (*Ok, error) {
	return client.SetChatClientData(c.Id, c.ClientData)
}

// SetDescription is a helper method for Client.SetChatDescription
func (c Chat) SetDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(c.Id, description)
}

// SetDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (c Chat) SetDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(c.Id, isEnabled, paidMessageStarCount)
}

// SetDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (c Chat) SetDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(c.Id, discussionChatId)
}

// SetDraftMessage is a helper method for Client.SetChatDraftMessage
func (c Chat) SetDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(c.Id, opts)
}

// SetEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (c Chat) SetEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(c.Id, opts)
}

// SetLocation is a helper method for Client.SetChatLocation
func (c Chat) SetLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(c.Id, location)
}

// SetMemberStatus is a helper method for Client.SetChatMemberStatus
func (c Chat) SetMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(c.Id, memberId, status)
}

// SetMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (c Chat) SetMessageAutoDeleteTime(client *Client) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(c.Id, c.MessageAutoDeleteTime)
}

// SetMessageSender is a helper method for Client.SetChatMessageSender
func (c Chat) SetMessageSender(client *Client) (*Ok, error) {
	return client.SetChatMessageSender(c.Id, c.MessageSenderId)
}

// SetNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (c Chat) SetNotificationSettings(client *Client) (*Ok, error) {
	return client.SetChatNotificationSettings(c.Id, c.NotificationSettings)
}

// SetPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (c Chat) SetPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(c.Id, paidMessageStarCount)
}

// SetPermissions is a helper method for Client.SetChatPermissions
func (c Chat) SetPermissions(client *Client) (*Ok, error) {
	return client.SetChatPermissions(c.Id, c.Permissions)
}

// SetPhoto is a helper method for Client.SetChatPhoto
func (c Chat) SetPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(c.Id, opts)
}

// SetPinnedStories is a helper method for Client.SetChatPinnedStories
func (c Chat) SetPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(c.Id, storyIds)
}

// SetProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (c Chat) SetProfileAccentColor(client *Client) (*Ok, error) {
	return client.SetChatProfileAccentColor(c.Id, c.ProfileAccentColorId, c.ProfileBackgroundCustomEmojiId)
}

// SetSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (c Chat) SetSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(c.Id, slowModeDelay)
}

// SetTheme is a helper method for Client.SetChatTheme
func (c Chat) SetTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(c.Id, theme)
}

// SetTitle is a helper method for Client.SetChatTitle
func (c Chat) SetTitle(client *Client) (*Ok, error) {
	return client.SetChatTitle(c.Id, c.Title)
}

// SetDirectMessagesTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (c Chat) SetDirectMessagesTopicIsMarkedAsUnread(client *Client, topicId int64) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(c.Id, topicId, c.IsMarkedAsUnread)
}

// SetEmojiStatus is a helper method for Client.SetEmojiStatus
func (c Chat) SetEmojiStatus(client *Client) (*Ok, error) {
	return client.SetEmojiStatus(c.EmojiStatus)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (c Chat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(c.Id, forumTopicId, c.NotificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (c Chat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(c.Id, messageId, editMessage, userId, score, force)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (c Chat) SetLiveStoryMessageSender(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetLiveStoryMessageSender(groupCallId, c.MessageSenderId)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (c Chat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(c.Id, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (c Chat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(c.Id, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (c Chat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(c.Id, messageId, typeField)
}

// SetPersonal is a helper method for Client.SetPersonalChat
func (c Chat) SetPersonal(client *Client) (*Ok, error) {
	return client.SetPersonalChat(c.Id)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (c Chat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(c.Id, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (c Chat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(c.Id, messageId, optionIds)
}

// SetProfileAccentColor is a helper method for Client.SetProfileAccentColor
func (c Chat) SetProfileAccentColor(client *Client) (*Ok, error) {
	return client.SetProfileAccentColor(c.ProfileAccentColorId, c.ProfileBackgroundCustomEmojiId)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c Chat) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, c.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (c Chat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(c.Id, storyAlbumId, name)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (c Chat) SetUserEmojiStatus(client *Client, userId int64) (*Ok, error) {
	return client.SetUserEmojiStatus(userId, c.EmojiStatus)
}

// SetVideoDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (c Chat) SetVideoDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(c.Id, defaultParticipantId)
}

// SetVideoTitle is a helper method for Client.SetVideoChatTitle
func (c Chat) SetVideoTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, c.Title)
}

// ShareWithBot is a helper method for Client.ShareChatWithBot
func (c Chat) ShareWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(c.Id, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (c Chat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(c.Id, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (c Chat) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, c.Title, recordVideo, usePortraitOrientation)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (c Chat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(c.Id, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (c Chat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, c.Id, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (c Chat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(c.Id, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (c Chat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(c.Id, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (c Chat) ToggleBusinessConnectedBotIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(c.Id, isPaused)
}

// ToggleDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (c Chat) ToggleDefaultDisableNotification(client *Client) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(c.Id, c.DefaultDisableNotification)
}

// ToggleGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (c Chat) ToggleGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(c.Id, areEnabled)
}

// ToggleHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (c Chat) ToggleHasProtectedContent(client *Client) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(c.Id, c.HasProtectedContent)
}

// ToggleIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (c Chat) ToggleIsMarkedAsUnread(client *Client) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(c.Id, c.IsMarkedAsUnread)
}

// ToggleIsPinned is a helper method for Client.ToggleChatIsPinned
func (c Chat) ToggleIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, c.Id, isPinned)
}

// ToggleIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (c Chat) ToggleIsTranslatable(client *Client) (*Ok, error) {
	return client.ToggleChatIsTranslatable(c.Id, c.IsTranslatable)
}

// ToggleViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (c Chat) ToggleViewAsTopics(client *Client) (*Ok, error) {
	return client.ToggleChatViewAsTopics(c.Id, c.ViewAsTopics)
}

// ToggleDirectMessagesTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (c Chat) ToggleDirectMessagesTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(c.Id, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (c Chat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(c.Id, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (c Chat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(c.Id, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (c Chat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(c.Id, isHidden)
}

// TransferOwnership is a helper method for Client.TransferChatOwnership
func (c Chat) TransferOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(c.Id, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (c Chat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(c.Id, messageId, toLanguageCode)
}

// UnpinAllMessages is a helper method for Client.UnpinAllChatMessages
func (c Chat) UnpinAllMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(c.Id)
}

// UnpinAllDirectMessagesTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (c Chat) UnpinAllDirectMessagesTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(c.Id, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (c Chat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(c.Id, forumTopicId)
}

// UnpinMessage is a helper method for Client.UnpinChatMessage
func (c Chat) UnpinMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(c.Id, messageId)
}

// UpgradeBasicGroupToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (c Chat) UpgradeBasicGroupToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(c.Id)
}

// ViewMessages is a helper method for Client.ViewMessages
func (c Chat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(c.Id, messageIds, forceRead, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (c ChatActionBarJoinRequest) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, c.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (c ChatActionBarJoinRequest) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, c.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (c ChatActionBarJoinRequest) CreateNewSupergroupChat(client *Client, isForum bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(c.Title, isForum, c.IsChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatActionBarJoinRequest) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, c.Title, startDate, isRtmpStream)
}

// GetChatBoostFeatures is a helper method for Client.GetChatBoostFeatures
func (c ChatActionBarJoinRequest) GetChatBoostFeatures(client *Client) (*ChatBoostFeatures, error) {
	return client.GetChatBoostFeatures(c.IsChannel)
}

// GetChatBoostLevelFeatures is a helper method for Client.GetChatBoostLevelFeatures
func (c ChatActionBarJoinRequest) GetChatBoostLevelFeatures(client *Client, level int32) (*ChatBoostLevelFeatures, error) {
	return client.GetChatBoostLevelFeatures(c.IsChannel, level)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (c ChatActionBarJoinRequest) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(c.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (c ChatActionBarJoinRequest) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, c.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (c ChatActionBarJoinRequest) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, c.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (c ChatActionBarJoinRequest) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, c.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (c ChatActionBarJoinRequest) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, c.Title, recordVideo, usePortraitOrientation)
}

// GetAnimatedEmoji is a helper method for Client.GetAnimatedEmoji
func (c ChatActionWatchingAnimations) GetAnimatedEmoji(client *Client) (*AnimatedEmoji, error) {
	return client.GetAnimatedEmoji(c.Emoji)
}

// GetEmojiReaction is a helper method for Client.GetEmojiReaction
func (c ChatActionWatchingAnimations) GetEmojiReaction(client *Client) (*EmojiReaction, error) {
	return client.GetEmojiReaction(c.Emoji)
}

// AddChatMember is a helper method for Client.AddChatMember
func (c ChatActiveStories) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(c.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (c ChatActiveStories) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(c.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (c ChatActiveStories) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(c.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (c ChatActiveStories) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(c.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (c ChatActiveStories) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, c.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (c ChatActiveStories) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(c.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (c ChatActiveStories) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(c.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (c ChatActiveStories) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(c.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (c ChatActiveStories) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(c.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (c ChatActiveStories) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(c.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (c ChatActiveStories) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(c.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (c ChatActiveStories) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(c.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (c ChatActiveStories) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(c.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (c ChatActiveStories) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(c.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (c ChatActiveStories) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(c.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (c ChatActiveStories) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(c.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (c ChatActiveStories) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(c.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (c ChatActiveStories) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(c.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (c ChatActiveStories) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(c.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (c ChatActiveStories) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(c.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (c ChatActiveStories) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(c.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (c ChatActiveStories) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(c.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (c ChatActiveStories) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(c.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (c ChatActiveStories) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(c.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (c ChatActiveStories) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(c.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (c ChatActiveStories) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(c.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (c ChatActiveStories) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(c.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (c ChatActiveStories) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(c.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (c ChatActiveStories) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(c.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (c ChatActiveStories) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(c.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (c ChatActiveStories) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(c.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (c ChatActiveStories) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(c.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (c ChatActiveStories) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(c.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (c ChatActiveStories) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(c.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (c ChatActiveStories) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(c.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (c ChatActiveStories) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(c.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (c ChatActiveStories) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(c.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (c ChatActiveStories) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(c.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (c ChatActiveStories) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(c.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (c ChatActiveStories) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (c ChatActiveStories) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, c.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (c ChatActiveStories) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (c ChatActiveStories) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (c ChatActiveStories) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, c.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (c ChatActiveStories) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, c.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (c ChatActiveStories) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(c.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (c ChatActiveStories) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(c.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (c ChatActiveStories) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(c.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (c ChatActiveStories) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(c.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (c ChatActiveStories) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(c.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (c ChatActiveStories) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(c.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (c ChatActiveStories) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(c.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (c ChatActiveStories) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(c.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (c ChatActiveStories) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(c.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (c ChatActiveStories) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(c.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (c ChatActiveStories) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(c.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (c ChatActiveStories) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, c.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (c ChatActiveStories) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(c.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (c ChatActiveStories) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(c.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (c ChatActiveStories) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(c.ChatId)
}

// Get is a helper method for Client.GetChatActiveStories
func (c ChatActiveStories) Get(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(c.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (c ChatActiveStories) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(c.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (c ChatActiveStories) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(c.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (c ChatActiveStories) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(c.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (c ChatActiveStories) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(c.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (c ChatActiveStories) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(c.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (c ChatActiveStories) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(c.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (c ChatActiveStories) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(c.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (c ChatActiveStories) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(c.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (c ChatActiveStories) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(c.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (c ChatActiveStories) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(c.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (c ChatActiveStories) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(c.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (c ChatActiveStories) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(c.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (c ChatActiveStories) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(c.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (c ChatActiveStories) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(c.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (c ChatActiveStories) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(c.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (c ChatActiveStories) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(c.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (c ChatActiveStories) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(c.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (c ChatActiveStories) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(c.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (c ChatActiveStories) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(c.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (c ChatActiveStories) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(c.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (c ChatActiveStories) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(c.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (c ChatActiveStories) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(c.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (c ChatActiveStories) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(c.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (c ChatActiveStories) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(c.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (c ChatActiveStories) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(c.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (c ChatActiveStories) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(c.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (c ChatActiveStories) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(c.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (c ChatActiveStories) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(c.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (c ChatActiveStories) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(c.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (c ChatActiveStories) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(c.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (c ChatActiveStories) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(c.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (c ChatActiveStories) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(c.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (c ChatActiveStories) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(c.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (c ChatActiveStories) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(c.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (c ChatActiveStories) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(c.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (c ChatActiveStories) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(c.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (c ChatActiveStories) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(c.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (c ChatActiveStories) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(c.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (c ChatActiveStories) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(c.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (c ChatActiveStories) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(c.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (c ChatActiveStories) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(c.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (c ChatActiveStories) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(c.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (c ChatActiveStories) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, c.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (c ChatActiveStories) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(c.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (c ChatActiveStories) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(c.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (c ChatActiveStories) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(c.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (c ChatActiveStories) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(c.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (c ChatActiveStories) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, c.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (c ChatActiveStories) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(c.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (c ChatActiveStories) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(c.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (c ChatActiveStories) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(c.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (c ChatActiveStories) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(c.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (c ChatActiveStories) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(c.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (c ChatActiveStories) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(c.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (c ChatActiveStories) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(c.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (c ChatActiveStories) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(c.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (c ChatActiveStories) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(c.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (c ChatActiveStories) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(c.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (c ChatActiveStories) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(c.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (c ChatActiveStories) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(c.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (c ChatActiveStories) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(c.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (c ChatActiveStories) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(c.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (c ChatActiveStories) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(c.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (c ChatActiveStories) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(c.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (c ChatActiveStories) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(c.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (c ChatActiveStories) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(c.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (c ChatActiveStories) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(c.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (c ChatActiveStories) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(c.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (c ChatActiveStories) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, c.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (c ChatActiveStories) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(c.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (c ChatActiveStories) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(c.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (c ChatActiveStories) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(c.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (c ChatActiveStories) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(c.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (c ChatActiveStories) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(c.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (c ChatActiveStories) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(c.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (c ChatActiveStories) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(c.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}
