package gotdbot

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (f ForumTopicInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(f.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (f ForumTopicInfo) ReadAllForumTopicMentions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicMentions(f.ChatId, f.ForumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (f ForumTopicInfo) ReadAllForumTopicReactions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicReactions(f.ChatId, f.ForumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (f ForumTopicInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, f.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (f ForumTopicInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(f.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (f ForumTopicInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(f.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (f ForumTopicInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(f.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (f ForumTopicInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(f.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (f ForumTopicInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(f.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (f ForumTopicInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(f.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (f ForumTopicInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (f ForumTopicInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, f.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (f ForumTopicInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(f.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (f ForumTopicInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (f ForumTopicInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(f.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (f ForumTopicInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(f.ChatId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (f ForumTopicInfo) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, f.Name, oldSticker, newSticker)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (f ForumTopicInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(f.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (f ForumTopicInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(f.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (f ForumTopicInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(f.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (f ForumTopicInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(f.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (f ForumTopicInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(f.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (f ForumTopicInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(f.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (f ForumTopicInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(f.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (f ForumTopicInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, f.ChatId, data)
}

// SearchBackground is a helper method for Client.SearchBackground
func (f ForumTopicInfo) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(f.Name)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (f ForumTopicInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(f.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (f ForumTopicInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(f.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (f ForumTopicInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(f.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (f ForumTopicInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(f.ChatId, query, offset, limit, opts)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (f ForumTopicInfo) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(f.Name, ignoreCache)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (f ForumTopicInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, f.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (f ForumTopicInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, f.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (f ForumTopicInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, f.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (f ForumTopicInfo) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(f.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (f ForumTopicInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(f.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (f ForumTopicInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(f.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (f ForumTopicInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(f.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (f ForumTopicInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(f.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (f ForumTopicInfo) SendTextMessageDraft(client *Client, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(f.ChatId, f.ForumTopicId, draftId, text)
}

// SetBotName is a helper method for Client.SetBotName
func (f ForumTopicInfo) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, f.Name)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (f ForumTopicInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, f.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (f ForumTopicInfo) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(f.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (f ForumTopicInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(f.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (f ForumTopicInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(f.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (f ForumTopicInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(f.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (f ForumTopicInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(f.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (f ForumTopicInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(f.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (f ForumTopicInfo) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(f.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (f ForumTopicInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(f.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (f ForumTopicInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(f.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (f ForumTopicInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(f.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (f ForumTopicInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(f.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (f ForumTopicInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(f.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (f ForumTopicInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(f.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (f ForumTopicInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(f.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (f ForumTopicInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(f.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (f ForumTopicInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(f.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (f ForumTopicInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(f.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (f ForumTopicInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(f.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (f ForumTopicInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(f.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (f ForumTopicInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(f.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (f ForumTopicInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(f.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (f ForumTopicInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(f.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (f ForumTopicInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(f.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (f ForumTopicInfo) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(f.ChatId, title)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (f ForumTopicInfo) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(f.Name, customEmojiId)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (f ForumTopicInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(f.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (f ForumTopicInfo) SetForumTopicNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(f.ChatId, f.ForumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (f ForumTopicInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(f.ChatId, messageId, editMessage, userId, score, force)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (f ForumTopicInfo) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, f.Name)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (f ForumTopicInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(f.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (f ForumTopicInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(f.ChatId, messageId, reactionTypes, isBig)
}

// SetOption is a helper method for Client.SetOption
func (f ForumTopicInfo) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(f.Name, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (f ForumTopicInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(f.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (f ForumTopicInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(f.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (f ForumTopicInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(f.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (f ForumTopicInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(f.ChatId, messageId, optionIds)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (f ForumTopicInfo) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, f.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (f ForumTopicInfo) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, f.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (f ForumTopicInfo) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(f.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (f ForumTopicInfo) SetStoryAlbumName(client *Client, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(f.ChatId, storyAlbumId, f.Name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (f ForumTopicInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(f.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (f ForumTopicInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(f.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (f ForumTopicInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(f.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (f ForumTopicInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(f.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (f ForumTopicInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, f.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (f ForumTopicInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(f.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (f ForumTopicInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(f.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (f ForumTopicInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(f.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (f ForumTopicInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(f.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (f ForumTopicInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(f.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (f ForumTopicInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(f.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (f ForumTopicInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(f.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (f ForumTopicInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, f.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (f ForumTopicInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(f.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (f ForumTopicInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(f.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (f ForumTopicInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(f.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (f ForumTopicInfo) ToggleForumTopicIsClosed(client *Client) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(f.ChatId, f.ForumTopicId, f.IsClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (f ForumTopicInfo) ToggleForumTopicIsPinned(client *Client, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(f.ChatId, f.ForumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (f ForumTopicInfo) ToggleGeneralForumTopicIsHidden(client *Client) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(f.ChatId, f.IsHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (f ForumTopicInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(f.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (f ForumTopicInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(f.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (f ForumTopicInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(f.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (f ForumTopicInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(f.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (f ForumTopicInfo) UnpinAllForumTopicMessages(client *Client) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(f.ChatId, f.ForumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (f ForumTopicInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(f.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (f ForumTopicInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(f.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (f ForumTopicInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(f.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (f ForwardSource) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(f.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (f ForwardSource) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(f.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (f ForwardSource) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(f.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (f ForwardSource) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(f.ChatId, f.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (f ForwardSource) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, f.ChatId, f.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (f ForwardSource) AddLocalMessage(client *Client, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(f.ChatId, f.SenderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (f ForwardSource) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(f.ChatId, f.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (f ForwardSource) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(f.ChatId, f.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (f ForwardSource) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(f.ChatId, f.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (f ForwardSource) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(f.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (f ForwardSource) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (f ForwardSource) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(f.ChatId, f.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (f ForwardSource) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(f.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (f ForwardSource) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(f.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (f ForwardSource) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(f.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (f ForwardSource) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(f.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (f ForwardSource) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(f.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (f ForwardSource) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(f.ChatId, f.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (f ForwardSource) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(f.ChatId, f.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (f ForwardSource) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(f.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (f ForwardSource) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(f.ChatId, f.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (f ForwardSource) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(f.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (f ForwardSource) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(f.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (f ForwardSource) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(f.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (f ForwardSource) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(f.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (f ForwardSource) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(f.ChatId, f.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (f ForwardSource) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(f.ChatId, f.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (f ForwardSource) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(f.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (f ForwardSource) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(f.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (f ForwardSource) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(f.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (f ForwardSource) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(f.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (f ForwardSource) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(f.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (f ForwardSource) DeleteChatMessagesBySender(client *Client) (*Ok, error) {
	return client.DeleteChatMessagesBySender(f.ChatId, f.SenderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (f ForwardSource) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(f.ChatId, f.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (f ForwardSource) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(f.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (f ForwardSource) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(f.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (f ForwardSource) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(f.ChatId, forumTopicId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (f ForwardSource) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, f.SenderId, reportSpam)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (f ForwardSource) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(f.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (f ForwardSource) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(f.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (f ForwardSource) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(f.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (f ForwardSource) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, f.ChatId, f.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (f ForwardSource) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, f.ChatId, f.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (f ForwardSource) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, f.ChatId, f.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (f ForwardSource) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, f.ChatId, f.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (f ForwardSource) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, f.ChatId, f.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (f ForwardSource) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, f.ChatId, f.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (f ForwardSource) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(f.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (f ForwardSource) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(f.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (f ForwardSource) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(f.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (f ForwardSource) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(f.ChatId, f.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (f ForwardSource) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(f.ChatId, f.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (f ForwardSource) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(f.ChatId, f.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (f ForwardSource) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(f.ChatId, f.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (f ForwardSource) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(f.ChatId, f.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (f ForwardSource) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(f.ChatId, f.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (f ForwardSource) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(f.ChatId, f.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (f ForwardSource) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, f.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (f ForwardSource) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(f.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (f ForwardSource) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, f.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (f ForwardSource) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(f.ChatId, f.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (f ForwardSource) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(f.ChatId, f.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (f ForwardSource) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(f.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (f ForwardSource) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(f.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (f ForwardSource) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(f.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (f ForwardSource) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(f.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (f ForwardSource) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(f.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (f ForwardSource) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(f.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (f ForwardSource) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(f.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (f ForwardSource) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(f.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (f ForwardSource) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(f.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (f ForwardSource) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(f.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (f ForwardSource) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(f.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (f ForwardSource) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(f.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (f ForwardSource) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(f.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (f ForwardSource) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(f.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (f ForwardSource) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(f.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (f ForwardSource) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(f.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (f ForwardSource) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(f.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (f ForwardSource) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(f.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (f ForwardSource) GetChatMessageByDate(client *Client) (*Message, error) {
	return client.GetChatMessageByDate(f.ChatId, f.Date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (f ForwardSource) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(f.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (f ForwardSource) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(f.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (f ForwardSource) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(f.ChatId, filter, f.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (f ForwardSource) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(f.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (f ForwardSource) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(f.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (f ForwardSource) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(f.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (f ForwardSource) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(f.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (f ForwardSource) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(f.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (f ForwardSource) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(f.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (f ForwardSource) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(f.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (f ForwardSource) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(f.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (f ForwardSource) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(f.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (f ForwardSource) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(f.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (f ForwardSource) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(f.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (f ForwardSource) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(f.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (f ForwardSource) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(f.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (f ForwardSource) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(f.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (f ForwardSource) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(f.ChatId, topicId, f.Date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (f ForwardSource) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(f.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (f ForwardSource) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(f.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (f ForwardSource) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(f.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (f ForwardSource) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(f.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (f ForwardSource) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(f.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (f ForwardSource) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(f.ChatId, f.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (f ForwardSource) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(f.ChatId, f.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (f ForwardSource) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, f.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (f ForwardSource) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(f.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (f ForwardSource) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(f.ChatId, f.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (f ForwardSource) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(f.ChatId, f.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (f ForwardSource) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(f.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (f ForwardSource) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, f.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (f ForwardSource) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(f.ChatId, f.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (f ForwardSource) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(f.ChatId, f.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (f ForwardSource) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(f.ChatId, f.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (f ForwardSource) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(f.ChatId, f.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (f ForwardSource) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(f.ChatId, f.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (f ForwardSource) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(f.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (f ForwardSource) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(f.ChatId, f.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (f ForwardSource) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(f.ChatId, f.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (f ForwardSource) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(f.ChatId, f.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (f ForwardSource) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(f.ChatId, f.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (f ForwardSource) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(f.ChatId, f.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (f ForwardSource) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(f.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (f ForwardSource) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(f.ChatId, f.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (f ForwardSource) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(f.ChatId, f.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (f ForwardSource) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(f.ChatId, f.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (f ForwardSource) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(f.ChatId, f.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (f ForwardSource) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(f.ChatId, f.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (f ForwardSource) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(f.ChatId, f.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (f ForwardSource) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(f.ChatId, f.MessageId)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (f ForwardSource) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, f.Date)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (f ForwardSource) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(f.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (f ForwardSource) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, f.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (f ForwardSource) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(f.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (f ForwardSource) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(f.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (f ForwardSource) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(f.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (f ForwardSource) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(f.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (f ForwardSource) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(f.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (f ForwardSource) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(f.ChatId, f.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (f ForwardSource) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(f.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (f ForwardSource) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(f.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (f ForwardSource) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(f.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (f ForwardSource) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(f.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (f ForwardSource) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(f.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (f ForwardSource) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(f.ChatId, f.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (f ForwardSource) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(f.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (f ForwardSource) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(f.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (f ForwardSource) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(f.ChatId, f.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (f ForwardSource) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(f.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (f ForwardSource) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(f.ChatId, f.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (f ForwardSource) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(f.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (f ForwardSource) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(f.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (f ForwardSource) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(f.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (f ForwardSource) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(f.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (f ForwardSource) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(f.ChatId, f.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (f ForwardSource) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(f.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (f ForwardSource) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(f.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (f ForwardSource) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(f.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (f ForwardSource) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(f.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (f ForwardSource) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(f.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (f ForwardSource) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, f.ChatId, f.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (f ForwardSource) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(f.ChatId, f.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (f ForwardSource) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(f.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (f ForwardSource) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(f.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (f ForwardSource) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(f.ChatId, f.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (f ForwardSource) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(f.ChatId, f.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (f ForwardSource) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(f.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (f ForwardSource) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (f ForwardSource) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, f.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (f ForwardSource) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(f.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (f ForwardSource) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(f.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (f ForwardSource) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(f.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (f ForwardSource) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(f.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (f ForwardSource) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(f.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (f ForwardSource) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(f.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (f ForwardSource) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(f.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (f ForwardSource) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(f.ChatId, f.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (f ForwardSource) ReportMessageReactions(client *Client) (*Ok, error) {
	return client.ReportMessageReactions(f.ChatId, f.MessageId, f.SenderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (f ForwardSource) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, f.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (f ForwardSource) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(f.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (f ForwardSource) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(f.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (f ForwardSource) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, f.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (f ForwardSource) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(f.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (f ForwardSource) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(f.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (f ForwardSource) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(f.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (f ForwardSource) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(f.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (f ForwardSource) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, f.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (f ForwardSource) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, f.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (f ForwardSource) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, f.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (f ForwardSource) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(f.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (f ForwardSource) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(f.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (f ForwardSource) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(f.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (f ForwardSource) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(f.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (f ForwardSource) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(f.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (f ForwardSource) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(f.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (f ForwardSource) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, f.ChatId, f.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (f ForwardSource) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(f.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (f ForwardSource) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(f.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (f ForwardSource) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(f.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (f ForwardSource) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(f.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (f ForwardSource) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(f.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (f ForwardSource) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(f.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (f ForwardSource) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(f.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (f ForwardSource) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(f.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (f ForwardSource) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(f.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (f ForwardSource) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(f.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (f ForwardSource) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(f.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (f ForwardSource) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(f.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (f ForwardSource) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(f.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (f ForwardSource) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(f.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (f ForwardSource) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(f.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (f ForwardSource) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(f.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (f ForwardSource) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(f.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (f ForwardSource) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(f.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (f ForwardSource) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(f.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (f ForwardSource) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(f.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (f ForwardSource) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(f.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (f ForwardSource) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(f.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (f ForwardSource) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(f.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (f ForwardSource) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(f.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (f ForwardSource) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(f.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (f ForwardSource) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(f.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (f ForwardSource) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(f.ChatId, f.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (f ForwardSource) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(f.ChatId, f.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (f ForwardSource) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(f.ChatId, f.MessageId, reactionTypes, isBig)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (f ForwardSource) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(f.SenderId, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (f ForwardSource) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(f.ChatId, f.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (f ForwardSource) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(f.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (f ForwardSource) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(f.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (f ForwardSource) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(f.ChatId, f.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (f ForwardSource) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(f.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (f ForwardSource) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(f.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (f ForwardSource) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(f.ChatId, f.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (f ForwardSource) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(f.ChatId, f.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (f ForwardSource) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(f.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (f ForwardSource) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, f.ChatId, f.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (f ForwardSource) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(f.ChatId, f.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (f ForwardSource) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(f.ChatId, f.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (f ForwardSource) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(f.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (f ForwardSource) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(f.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (f ForwardSource) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(f.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (f ForwardSource) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(f.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (f ForwardSource) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(f.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (f ForwardSource) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, f.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (f ForwardSource) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(f.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (f ForwardSource) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(f.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (f ForwardSource) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(f.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (f ForwardSource) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(f.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (f ForwardSource) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(f.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (f ForwardSource) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(f.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (f ForwardSource) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(f.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (f ForwardSource) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(f.ChatId, f.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (f ForwardSource) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(f.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (f ForwardSource) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(f.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (f ForwardSource) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(f.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (f ForwardSource) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(f.ChatId, f.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (f ForwardSource) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(f.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (f ForwardSource) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(f.ChatId, messageIds, forceRead, opts)
}

// AddBotMediaPreview is a helper method for Client.AddBotMediaPreview
func (f FoundAffiliateProgram) AddBotMediaPreview(client *Client, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.AddBotMediaPreview(f.BotUserId, languageCode, content)
}

// AllowBotToSendMessages is a helper method for Client.AllowBotToSendMessages
func (f FoundAffiliateProgram) AllowBotToSendMessages(client *Client) (*Ok, error) {
	return client.AllowBotToSendMessages(f.BotUserId)
}

// CanBotSendMessages is a helper method for Client.CanBotSendMessages
func (f FoundAffiliateProgram) CanBotSendMessages(client *Client) (*Ok, error) {
	return client.CanBotSendMessages(f.BotUserId)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (f FoundAffiliateProgram) CheckWebAppFileDownload(client *Client, fileName string, url string) (*Ok, error) {
	return client.CheckWebAppFileDownload(f.BotUserId, fileName, url)
}

// ConnectAffiliateProgram is a helper method for Client.ConnectAffiliateProgram
func (f FoundAffiliateProgram) ConnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.ConnectAffiliateProgram(affiliate, f.BotUserId)
}

// DeleteBotMediaPreviews is a helper method for Client.DeleteBotMediaPreviews
func (f FoundAffiliateProgram) DeleteBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.DeleteBotMediaPreviews(f.BotUserId, languageCode, fileIds)
}

// DeleteBusinessConnectedBot is a helper method for Client.DeleteBusinessConnectedBot
func (f FoundAffiliateProgram) DeleteBusinessConnectedBot(client *Client) (*Ok, error) {
	return client.DeleteBusinessConnectedBot(f.BotUserId)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (f FoundAffiliateProgram) EditBotMediaPreview(client *Client, languageCode string, fileId int32, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(f.BotUserId, languageCode, fileId, content)
}

// GetAttachmentMenuBot is a helper method for Client.GetAttachmentMenuBot
func (f FoundAffiliateProgram) GetAttachmentMenuBot(client *Client) (*AttachmentMenuBot, error) {
	return client.GetAttachmentMenuBot(f.BotUserId)
}

// GetBotInfoDescription is a helper method for Client.GetBotInfoDescription
func (f FoundAffiliateProgram) GetBotInfoDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoDescription(f.BotUserId, languageCode)
}

// GetBotInfoShortDescription is a helper method for Client.GetBotInfoShortDescription
func (f FoundAffiliateProgram) GetBotInfoShortDescription(client *Client, languageCode string) (*Text, error) {
	return client.GetBotInfoShortDescription(f.BotUserId, languageCode)
}

// GetBotMediaPreviewInfo is a helper method for Client.GetBotMediaPreviewInfo
func (f FoundAffiliateProgram) GetBotMediaPreviewInfo(client *Client, languageCode string) (*BotMediaPreviewInfo, error) {
	return client.GetBotMediaPreviewInfo(f.BotUserId, languageCode)
}

// GetBotMediaPreviews is a helper method for Client.GetBotMediaPreviews
func (f FoundAffiliateProgram) GetBotMediaPreviews(client *Client) (*BotMediaPreviews, error) {
	return client.GetBotMediaPreviews(f.BotUserId)
}

// GetBotName is a helper method for Client.GetBotName
func (f FoundAffiliateProgram) GetBotName(client *Client, languageCode string) (*Text, error) {
	return client.GetBotName(f.BotUserId, languageCode)
}

// GetBotSimilarBotCount is a helper method for Client.GetBotSimilarBotCount
func (f FoundAffiliateProgram) GetBotSimilarBotCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetBotSimilarBotCount(f.BotUserId, returnLocal)
}

// GetBotSimilarBots is a helper method for Client.GetBotSimilarBots
func (f FoundAffiliateProgram) GetBotSimilarBots(client *Client) (*Users, error) {
	return client.GetBotSimilarBots(f.BotUserId)
}

// GetConnectedAffiliateProgram is a helper method for Client.GetConnectedAffiliateProgram
func (f FoundAffiliateProgram) GetConnectedAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.GetConnectedAffiliateProgram(affiliate, f.BotUserId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (f FoundAffiliateProgram) GetInlineQueryResults(client *Client, chatId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(f.BotUserId, chatId, query, offset, opts)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (f FoundAffiliateProgram) GetMainWebApp(client *Client, chatId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, f.BotUserId, startParameter, parameters)
}

// GetPassportAuthorizationForm is a helper method for Client.GetPassportAuthorizationForm
func (f FoundAffiliateProgram) GetPassportAuthorizationForm(client *Client, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	return client.GetPassportAuthorizationForm(f.BotUserId, scope, publicKey, nonce)
}

// GetPreparedInlineMessage is a helper method for Client.GetPreparedInlineMessage
func (f FoundAffiliateProgram) GetPreparedInlineMessage(client *Client, preparedMessageId string) (*PreparedInlineMessage, error) {
	return client.GetPreparedInlineMessage(f.BotUserId, preparedMessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (f FoundAffiliateProgram) GetWebAppLinkUrl(client *Client, chatId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, f.BotUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GetWebAppPlaceholder is a helper method for Client.GetWebAppPlaceholder
func (f FoundAffiliateProgram) GetWebAppPlaceholder(client *Client) (*Outline, error) {
	return client.GetWebAppPlaceholder(f.BotUserId)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (f FoundAffiliateProgram) GetWebAppUrl(client *Client, url string, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(f.BotUserId, url, parameters)
}

// OpenBotSimilarBot is a helper method for Client.OpenBotSimilarBot
func (f FoundAffiliateProgram) OpenBotSimilarBot(client *Client, openedBotUserId int64) (*Ok, error) {
	return client.OpenBotSimilarBot(f.BotUserId, openedBotUserId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (f FoundAffiliateProgram) OpenWebApp(client *Client, chatId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, f.BotUserId, url, parameters, opts)
}

// RemoveMessageSenderBotVerification is a helper method for Client.RemoveMessageSenderBotVerification
func (f FoundAffiliateProgram) RemoveMessageSenderBotVerification(client *Client, verifiedId *MessageSender) (*Ok, error) {
	return client.RemoveMessageSenderBotVerification(f.BotUserId, verifiedId)
}

// ReorderBotActiveUsernames is a helper method for Client.ReorderBotActiveUsernames
func (f FoundAffiliateProgram) ReorderBotActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderBotActiveUsernames(f.BotUserId, usernames)
}

// ReorderBotMediaPreviews is a helper method for Client.ReorderBotMediaPreviews
func (f FoundAffiliateProgram) ReorderBotMediaPreviews(client *Client, languageCode string, fileIds []int32) (*Ok, error) {
	return client.ReorderBotMediaPreviews(f.BotUserId, languageCode, fileIds)
}

// SearchWebApp is a helper method for Client.SearchWebApp
func (f FoundAffiliateProgram) SearchWebApp(client *Client, webAppShortName string) (*FoundWebApp, error) {
	return client.SearchWebApp(f.BotUserId, webAppShortName)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (f FoundAffiliateProgram) SendBotStartMessage(client *Client, chatId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(f.BotUserId, chatId, parameter)
}

// SendWebAppCustomRequest is a helper method for Client.SendWebAppCustomRequest
func (f FoundAffiliateProgram) SendWebAppCustomRequest(client *Client, method string, parameters string) (*CustomRequestResult, error) {
	return client.SendWebAppCustomRequest(f.BotUserId, method, parameters)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (f FoundAffiliateProgram) SendWebAppData(client *Client, buttonText string, data string) (*Ok, error) {
	return client.SendWebAppData(f.BotUserId, buttonText, data)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (f FoundAffiliateProgram) SetBotInfoDescription(client *Client, languageCode string, description string) (*Ok, error) {
	return client.SetBotInfoDescription(f.BotUserId, languageCode, description)
}

// SetBotInfoShortDescription is a helper method for Client.SetBotInfoShortDescription
func (f FoundAffiliateProgram) SetBotInfoShortDescription(client *Client, languageCode string, shortDescription string) (*Ok, error) {
	return client.SetBotInfoShortDescription(f.BotUserId, languageCode, shortDescription)
}

// SetBotName is a helper method for Client.SetBotName
func (f FoundAffiliateProgram) SetBotName(client *Client, languageCode string, name string) (*Ok, error) {
	return client.SetBotName(f.BotUserId, languageCode, name)
}

// SetBotProfilePhoto is a helper method for Client.SetBotProfilePhoto
func (f FoundAffiliateProgram) SetBotProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetBotProfilePhoto(f.BotUserId, photo)
}

// SetMessageSenderBotVerification is a helper method for Client.SetMessageSenderBotVerification
func (f FoundAffiliateProgram) SetMessageSenderBotVerification(client *Client, verifiedId *MessageSender, customDescription string) (*Ok, error) {
	return client.SetMessageSenderBotVerification(f.BotUserId, verifiedId, customDescription)
}

// ToggleBotCanManageEmojiStatus is a helper method for Client.ToggleBotCanManageEmojiStatus
func (f FoundAffiliateProgram) ToggleBotCanManageEmojiStatus(client *Client, canManageEmojiStatus bool) (*Ok, error) {
	return client.ToggleBotCanManageEmojiStatus(f.BotUserId, canManageEmojiStatus)
}

// ToggleBotIsAddedToAttachmentMenu is a helper method for Client.ToggleBotIsAddedToAttachmentMenu
func (f FoundAffiliateProgram) ToggleBotIsAddedToAttachmentMenu(client *Client, isAdded bool, allowWriteAccess bool) (*Ok, error) {
	return client.ToggleBotIsAddedToAttachmentMenu(f.BotUserId, isAdded, allowWriteAccess)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (f FoundAffiliateProgram) ToggleBotUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(f.BotUserId, username, isActive)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundAffiliatePrograms) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundChatBoosts) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundFileDownloads) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundMessages) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (f FoundPosition) SetStickerPositionInSet(client *Client, sticker *InputFile) (*Ok, error) {
	return client.SetStickerPositionInSet(sticker, f.Position)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundPublicPosts) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundStories) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (f FoundUsers) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, f.UserIds)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (f FoundUsers) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, f.NextOffset, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (f FoundUsers) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(f.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (f FoundUsers) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, f.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (f FoundUsers) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, f.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (f FoundUsers) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(f.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (f FoundUsers) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(f.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (g Game) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, g.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (g Game) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, g.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (g Game) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(g.Title, isForum, isChannel, g.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (g Game) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, g.Title, startDate, isRtmpStream)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (g Game) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(g.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (g Game) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(g.Text)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (g Game) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(g.Title)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g Game) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, g.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (g Game) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(g.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g Game) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, g.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (g Game) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(g.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (g Game) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, g.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (g Game) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, g.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (g Game) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, g.Text)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (g Game) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, g.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (g Game) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, g.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (g Game) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, g.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (g Game) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, g.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (g Game) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, g.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (g Game) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, g.Title, recordVideo, usePortraitOrientation)
}

// TranslateText is a helper method for Client.TranslateText
func (g Game) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(g.Text, toLanguageCode)
}

// AddChatMember is a helper method for Client.AddChatMember
func (g GameHighScore) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, g.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (g GameHighScore) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(g.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (g GameHighScore) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(g.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (g GameHighScore) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(g.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (g GameHighScore) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(g.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (g GameHighScore) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(g.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (g GameHighScore) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(g.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (g GameHighScore) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(g.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (g GameHighScore) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(g.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (g GameHighScore) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(g.UserId, telegramPaymentChargeId, isCanceled)
}

// Gets is a helper method for Client.GetGameHighScores
func (g GameHighScore) Gets(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, g.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (g GameHighScore) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(g.UserId, offsetChatId, limit)
}

// GetInlines is a helper method for Client.GetInlineGameHighScores
func (g GameHighScore) GetInlines(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, g.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (g GameHighScore) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(g.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (g GameHighScore) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(g.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (g GameHighScore) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(g.UserId)
}

// GetUser is a helper method for Client.GetUser
func (g GameHighScore) GetUser(client *Client) (*User, error) {
	return client.GetUser(g.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (g GameHighScore) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, g.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (g GameHighScore) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(g.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (g GameHighScore) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(g.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (g GameHighScore) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(g.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (g GameHighScore) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(g.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GameHighScore) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(g.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (g GameHighScore) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, g.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GameHighScore) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, g.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (g GameHighScore) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, g.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (g GameHighScore) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(g.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (g GameHighScore) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(g.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (g GameHighScore) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(g.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (g GameHighScore) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, g.UserId, g.Score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (g GameHighScore) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, g.UserId, g.Score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (g GameHighScore) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(g.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (g GameHighScore) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(g.UserId, errors)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (g GameHighScore) SetStickerPositionInSet(client *Client, sticker *InputFile) (*Ok, error) {
	return client.SetStickerPositionInSet(sticker, g.Position)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (g GameHighScore) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(g.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (g GameHighScore) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(g.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (g GameHighScore) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(g.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (g GameHighScore) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(g.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (g GameHighScore) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(g.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (g GameHighScore) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(g.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (g GameHighScore) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(g.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (g GameHighScore) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(g.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (g GameHighScore) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, g.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (g GameHighScore) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(g.UserId, stickerFormat, sticker)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g Gift) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g Gift) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyUpgrade is a helper method for Client.BuyGiftUpgrade
func (g Gift) BuyUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// CanSend is a helper method for Client.CanSendGift
func (g Gift) CanSend(client *Client) (*CanSendGiftResult, error) {
	return client.CanSendGift(g.Id)
}

// CloseAuction is a helper method for Client.CloseGiftAuction
func (g Gift) CloseAuction(client *Client) (*Ok, error) {
	return client.CloseGiftAuction(g.Id)
}

// DropOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g Gift) DropOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetAuctionAcquiredGifts is a helper method for Client.GetGiftAuctionAcquiredGifts
func (g Gift) GetAuctionAcquiredGifts(client *Client) (*GiftAuctionAcquiredGifts, error) {
	return client.GetGiftAuctionAcquiredGifts(g.Id)
}

// GetUpgradePreview is a helper method for Client.GetGiftUpgradePreview
func (g Gift) GetUpgradePreview(client *Client) (*GiftUpgradePreview, error) {
	return client.GetGiftUpgradePreview(g.Id)
}

// GetUpgradeVariants is a helper method for Client.GetGiftUpgradeVariants
func (g Gift) GetUpgradeVariants(client *Client) (*GiftUpgradeVariants, error) {
	return client.GetGiftUpgradeVariants(g.Id)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g Gift) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// PremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g Gift) PremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, text)
}

// IncreaseAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g Gift) IncreaseAuctionBid(client *Client) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(g.Id, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g Gift) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// OpenAuction is a helper method for Client.OpenGiftAuction
func (g Gift) OpenAuction(client *Client) (*Ok, error) {
	return client.OpenGiftAuction(g.Id)
}

// PlaceAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g Gift) PlaceAuctionBid(client *Client, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(g.Id, g.StarCount, userId, text, isPrivate)
}

// SearchsForResale is a helper method for Client.SearchGiftsForResale
func (g Gift) SearchsForResale(client *Client, order *GiftForResaleOrder, attributes []*UpgradedGiftAttributeId, offset string, limit int32) (*GiftsForResale, error) {
	return client.SearchGiftsForResale(g.Id, order, attributes, offset, limit)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g Gift) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// Send is a helper method for Client.SendGift
func (g Gift) Send(client *Client, ownerId *MessageSender, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(g.Id, ownerId, text, isPrivate, payForUpgrade)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g Gift) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// Transfer is a helper method for Client.TransferGift
func (g Gift) Transfer(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// Upgrade is a helper method for Client.UpgradeGift
func (g Gift) Upgrade(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (g GiftAuction) CreateVideoChat(client *Client, chatId int64, title string, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, title, g.StartDate, isRtmpStream)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GiftAuctionAcquiredGift) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiftAuctionAcquiredGift) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (g GiftAuctionAcquiredGift) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiftAuctionAcquiredGift) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (g GiftAuctionAcquiredGift) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, g.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (g GiftAuctionAcquiredGift) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, g.Date)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (g GiftAuctionAcquiredGift) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(g.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (g GiftAuctionAcquiredGift) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(g.Text)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (g GiftAuctionAcquiredGift) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, g.Date)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g GiftAuctionAcquiredGift) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiftAuctionAcquiredGift) GiftPremiumWithStars(client *Client, userId int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, g.Text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g GiftAuctionAcquiredGift) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiftAuctionAcquiredGift) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (g GiftAuctionAcquiredGift) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(g.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GiftAuctionAcquiredGift) PlaceGiftAuctionBid(client *Client, giftId string, userId int64) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, g.StarCount, userId, g.Text, g.IsPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g GiftAuctionAcquiredGift) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// SearchQuote is a helper method for Client.SearchQuote
func (g GiftAuctionAcquiredGift) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(g.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (g GiftAuctionAcquiredGift) SendGift(client *Client, giftId string, ownerId *MessageSender, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, g.Text, g.IsPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (g GiftAuctionAcquiredGift) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, g.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (g GiftAuctionAcquiredGift) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, g.Text)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g GiftAuctionAcquiredGift) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiftAuctionAcquiredGift) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// TranslateText is a helper method for Client.TranslateText
func (g GiftAuctionAcquiredGift) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(g.Text, toLanguageCode)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiftAuctionAcquiredGift) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (g GiftChatThemes) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, g.NextOffset, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (g GiftCollection) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, g.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (g GiftCollection) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(g.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (g GiftCollection) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(g.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (g GiftCollection) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, g.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (g GiftCollection) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, g.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (g GiftCollection) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, g.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (g GiftCollection) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, g.Name, isNameImplicit, icon)
}

// Create is a helper method for Client.CreateGiftCollection
func (g GiftCollection) Create(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, g.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (g GiftCollection) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, g.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (g GiftCollection) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, g.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (g GiftCollection) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(g.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (g GiftCollection) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, g.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (g GiftCollection) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, g.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (g GiftCollection) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, g.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (g GiftCollection) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, g.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (g GiftCollection) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(g.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (g GiftCollection) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(g.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (g GiftCollection) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(g.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (g GiftCollection) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(g.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (g GiftCollection) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, g.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (g GiftCollection) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(g.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (g GiftCollection) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(g.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (g GiftCollection) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, g.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (g GiftCollection) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(g.Name, customEmojiId)
}

// SetName is a helper method for Client.SetGiftCollectionName
func (g GiftCollection) SetName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, g.Name)
}

// SetOption is a helper method for Client.SetOption
func (g GiftCollection) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(g.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (g GiftCollection) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, g.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (g GiftCollection) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, g.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (g GiftCollection) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(g.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (g GiftCollection) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, g.Name)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiftForResale) DropGiftOriginalDetails(client *Client, starCount int64) (*Ok, error) {
	return client.DropGiftOriginalDetails(g.ReceivedGiftId, starCount)
}

// GetReceivedGift is a helper method for Client.GetReceivedGift
func (g GiftForResale) GetReceivedGift(client *Client) (*ReceivedGift, error) {
	return client.GetReceivedGift(g.ReceivedGiftId)
}

// GetUpgradedGiftWithdrawalUrl is a helper method for Client.GetUpgradedGiftWithdrawalUrl
func (g GiftForResale) GetUpgradedGiftWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetUpgradedGiftWithdrawalUrl(g.ReceivedGiftId, password)
}

// SellGift is a helper method for Client.SellGift
func (g GiftForResale) SellGift(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SellGift(businessConnectionId, g.ReceivedGiftId)
}

// SetGiftResalePrice is a helper method for Client.SetGiftResalePrice
func (g GiftForResale) SetGiftResalePrice(client *Client, opts *SetGiftResalePriceOpts) (*Ok, error) {
	return client.SetGiftResalePrice(g.ReceivedGiftId, opts)
}

// ToggleGiftIsSaved is a helper method for Client.ToggleGiftIsSaved
func (g GiftForResale) ToggleGiftIsSaved(client *Client, isSaved bool) (*Ok, error) {
	return client.ToggleGiftIsSaved(g.ReceivedGiftId, isSaved)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiftForResale) TransferGift(client *Client, businessConnectionId string, newOwnerId *MessageSender, starCount int64) (*Ok, error) {
	return client.TransferGift(businessConnectionId, g.ReceivedGiftId, newOwnerId, starCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiftForResale) UpgradeGift(client *Client, businessConnectionId string, keepOriginalDetails bool, starCount int64) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, g.ReceivedGiftId, keepOriginalDetails, starCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GiftResaleParameters) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiftResaleParameters) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (g GiftResaleParameters) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiftResaleParameters) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g GiftResaleParameters) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiftResaleParameters) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g GiftResaleParameters) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiftResaleParameters) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GiftResaleParameters) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, g.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g GiftResaleParameters) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g GiftResaleParameters) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiftResaleParameters) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiftResaleParameters) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GiftResalePriceStar) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiftResalePriceStar) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (g GiftResalePriceStar) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiftResalePriceStar) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g GiftResalePriceStar) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiftResalePriceStar) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g GiftResalePriceStar) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiftResalePriceStar) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GiftResalePriceStar) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, g.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g GiftResalePriceStar) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g GiftResalePriceStar) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiftResalePriceStar) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiftResalePriceStar) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (g GiftResaleResultPriceIncreased) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, g.Price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (g GiftResaleResultPriceIncreased) SendResoldGift(client *Client, giftName string, ownerId *MessageSender) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, ownerId, g.Price)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (g GiftsForResale) AnswerInlineQuery(client *Client, inlineQueryId string, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(inlineQueryId, isPersonal, results, cacheTime, g.NextOffset, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GiftUpgradePrice) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiftUpgradePrice) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (g GiftUpgradePrice) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiftUpgradePrice) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (g GiftUpgradePrice) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, g.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (g GiftUpgradePrice) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, g.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (g GiftUpgradePrice) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, g.Date)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g GiftUpgradePrice) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiftUpgradePrice) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g GiftUpgradePrice) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiftUpgradePrice) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GiftUpgradePrice) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, g.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g GiftUpgradePrice) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g GiftUpgradePrice) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiftUpgradePrice) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiftUpgradePrice) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiveawayInfoCompleted) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, g.WinnerCount, starCount)
}

// GetPremiumGiveawayPaymentOptions is a helper method for Client.GetPremiumGiveawayPaymentOptions
func (g GiveawayParameters) GetPremiumGiveawayPaymentOptions(client *Client) (*PremiumGiveawayPaymentOptions, error) {
	return client.GetPremiumGiveawayPaymentOptions(g.BoostedChatId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (g GiveawayParticipantStatusAdministrator) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(g.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (g GiveawayParticipantStatusAdministrator) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(g.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (g GiveawayParticipantStatusAdministrator) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(g.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (g GiveawayParticipantStatusAdministrator) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(g.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (g GiveawayParticipantStatusAdministrator) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, g.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (g GiveawayParticipantStatusAdministrator) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(g.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (g GiveawayParticipantStatusAdministrator) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(g.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (g GiveawayParticipantStatusAdministrator) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(g.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiveawayParticipantStatusAdministrator) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(g.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (g GiveawayParticipantStatusAdministrator) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(g.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (g GiveawayParticipantStatusAdministrator) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(g.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (g GiveawayParticipantStatusAdministrator) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(g.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (g GiveawayParticipantStatusAdministrator) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(g.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (g GiveawayParticipantStatusAdministrator) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(g.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (g GiveawayParticipantStatusAdministrator) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(g.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (g GiveawayParticipantStatusAdministrator) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(g.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (g GiveawayParticipantStatusAdministrator) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(g.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (g GiveawayParticipantStatusAdministrator) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(g.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (g GiveawayParticipantStatusAdministrator) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(g.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (g GiveawayParticipantStatusAdministrator) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(g.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (g GiveawayParticipantStatusAdministrator) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(g.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (g GiveawayParticipantStatusAdministrator) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(g.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (g GiveawayParticipantStatusAdministrator) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(g.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (g GiveawayParticipantStatusAdministrator) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(g.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (g GiveawayParticipantStatusAdministrator) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(g.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (g GiveawayParticipantStatusAdministrator) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(g.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (g GiveawayParticipantStatusAdministrator) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(g.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (g GiveawayParticipantStatusAdministrator) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(g.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (g GiveawayParticipantStatusAdministrator) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(g.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (g GiveawayParticipantStatusAdministrator) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(g.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (g GiveawayParticipantStatusAdministrator) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(g.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (g GiveawayParticipantStatusAdministrator) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(g.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (g GiveawayParticipantStatusAdministrator) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(g.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (g GiveawayParticipantStatusAdministrator) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(g.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (g GiveawayParticipantStatusAdministrator) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(g.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (g GiveawayParticipantStatusAdministrator) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(g.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (g GiveawayParticipantStatusAdministrator) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(g.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (g GiveawayParticipantStatusAdministrator) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(g.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (g GiveawayParticipantStatusAdministrator) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(g.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, g.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, g.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, g.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, g.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, g.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (g GiveawayParticipantStatusAdministrator) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, g.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (g GiveawayParticipantStatusAdministrator) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(g.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (g GiveawayParticipantStatusAdministrator) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(g.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (g GiveawayParticipantStatusAdministrator) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(g.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (g GiveawayParticipantStatusAdministrator) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(g.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (g GiveawayParticipantStatusAdministrator) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(g.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (g GiveawayParticipantStatusAdministrator) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(g.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (g GiveawayParticipantStatusAdministrator) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(g.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (g GiveawayParticipantStatusAdministrator) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(g.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (g GiveawayParticipantStatusAdministrator) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(g.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (g GiveawayParticipantStatusAdministrator) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(g.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (g GiveawayParticipantStatusAdministrator) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(g.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (g GiveawayParticipantStatusAdministrator) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, g.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (g GiveawayParticipantStatusAdministrator) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(g.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (g GiveawayParticipantStatusAdministrator) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(g.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (g GiveawayParticipantStatusAdministrator) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(g.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (g GiveawayParticipantStatusAdministrator) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(g.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (g GiveawayParticipantStatusAdministrator) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(g.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (g GiveawayParticipantStatusAdministrator) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(g.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (g GiveawayParticipantStatusAdministrator) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(g.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (g GiveawayParticipantStatusAdministrator) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(g.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (g GiveawayParticipantStatusAdministrator) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(g.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (g GiveawayParticipantStatusAdministrator) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(g.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (g GiveawayParticipantStatusAdministrator) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(g.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (g GiveawayParticipantStatusAdministrator) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(g.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (g GiveawayParticipantStatusAdministrator) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(g.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (g GiveawayParticipantStatusAdministrator) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(g.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (g GiveawayParticipantStatusAdministrator) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(g.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (g GiveawayParticipantStatusAdministrator) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(g.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (g GiveawayParticipantStatusAdministrator) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(g.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (g GiveawayParticipantStatusAdministrator) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(g.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (g GiveawayParticipantStatusAdministrator) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(g.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (g GiveawayParticipantStatusAdministrator) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(g.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (g GiveawayParticipantStatusAdministrator) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(g.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (g GiveawayParticipantStatusAdministrator) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(g.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (g GiveawayParticipantStatusAdministrator) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(g.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (g GiveawayParticipantStatusAdministrator) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(g.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (g GiveawayParticipantStatusAdministrator) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(g.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (g GiveawayParticipantStatusAdministrator) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(g.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (g GiveawayParticipantStatusAdministrator) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(g.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (g GiveawayParticipantStatusAdministrator) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(g.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (g GiveawayParticipantStatusAdministrator) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(g.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (g GiveawayParticipantStatusAdministrator) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(g.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (g GiveawayParticipantStatusAdministrator) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(g.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (g GiveawayParticipantStatusAdministrator) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(g.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (g GiveawayParticipantStatusAdministrator) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(g.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (g GiveawayParticipantStatusAdministrator) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(g.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (g GiveawayParticipantStatusAdministrator) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(g.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (g GiveawayParticipantStatusAdministrator) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(g.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (g GiveawayParticipantStatusAdministrator) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(g.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (g GiveawayParticipantStatusAdministrator) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(g.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (g GiveawayParticipantStatusAdministrator) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(g.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (g GiveawayParticipantStatusAdministrator) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(g.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (g GiveawayParticipantStatusAdministrator) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(g.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (g GiveawayParticipantStatusAdministrator) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(g.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (g GiveawayParticipantStatusAdministrator) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(g.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (g GiveawayParticipantStatusAdministrator) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(g.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (g GiveawayParticipantStatusAdministrator) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(g.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (g GiveawayParticipantStatusAdministrator) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(g.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (g GiveawayParticipantStatusAdministrator) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, g.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (g GiveawayParticipantStatusAdministrator) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(g.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (g GiveawayParticipantStatusAdministrator) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(g.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (g GiveawayParticipantStatusAdministrator) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(g.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (g GiveawayParticipantStatusAdministrator) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(g.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (g GiveawayParticipantStatusAdministrator) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, g.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (g GiveawayParticipantStatusAdministrator) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(g.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (g GiveawayParticipantStatusAdministrator) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(g.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (g GiveawayParticipantStatusAdministrator) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(g.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (g GiveawayParticipantStatusAdministrator) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(g.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (g GiveawayParticipantStatusAdministrator) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(g.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (g GiveawayParticipantStatusAdministrator) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(g.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (g GiveawayParticipantStatusAdministrator) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(g.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (g GiveawayParticipantStatusAdministrator) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(g.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (g GiveawayParticipantStatusAdministrator) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(g.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (g GiveawayParticipantStatusAdministrator) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(g.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (g GiveawayParticipantStatusAdministrator) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(g.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (g GiveawayParticipantStatusAdministrator) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(g.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (g GiveawayParticipantStatusAdministrator) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(g.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (g GiveawayParticipantStatusAdministrator) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(g.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (g GiveawayParticipantStatusAdministrator) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(g.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (g GiveawayParticipantStatusAdministrator) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(g.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (g GiveawayParticipantStatusAdministrator) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(g.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (g GiveawayParticipantStatusAdministrator) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(g.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (g GiveawayParticipantStatusAdministrator) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(g.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (g GiveawayParticipantStatusAdministrator) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(g.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (g GiveawayParticipantStatusAdministrator) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, g.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (g GiveawayParticipantStatusAdministrator) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(g.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (g GiveawayParticipantStatusAdministrator) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(g.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (g GiveawayParticipantStatusAdministrator) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(g.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (g GiveawayParticipantStatusAdministrator) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(g.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (g GiveawayParticipantStatusAdministrator) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(g.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (g GiveawayParticipantStatusAdministrator) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(g.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (g GiveawayParticipantStatusAdministrator) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(g.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (g GiveawayParticipantStatusAdministrator) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(g.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (g GiveawayParticipantStatusAdministrator) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(g.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (g GiveawayParticipantStatusAdministrator) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(g.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (g GiveawayParticipantStatusAdministrator) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(g.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (g GiveawayParticipantStatusAdministrator) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(g.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (g GiveawayParticipantStatusAdministrator) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(g.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (g GiveawayParticipantStatusAdministrator) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(g.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (g GiveawayParticipantStatusAdministrator) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(g.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (g GiveawayParticipantStatusAdministrator) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(g.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (g GiveawayParticipantStatusAdministrator) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(g.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (g GiveawayParticipantStatusAdministrator) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(g.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (g GiveawayParticipantStatusAdministrator) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(g.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (g GiveawayParticipantStatusAdministrator) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(g.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (g GiveawayParticipantStatusAdministrator) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(g.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (g GiveawayParticipantStatusAdministrator) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(g.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (g GiveawayParticipantStatusAdministrator) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(g.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (g GiveawayParticipantStatusAdministrator) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(g.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (g GiveawayParticipantStatusAdministrator) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(g.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (g GiveawayParticipantStatusAdministrator) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(g.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (g GiveawayParticipantStatusAdministrator) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, g.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (g GiveawayParticipantStatusAdministrator) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(g.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (g GiveawayParticipantStatusAdministrator) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(g.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (g GiveawayParticipantStatusAdministrator) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(g.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (g GiveawayParticipantStatusAdministrator) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(g.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (g GiveawayParticipantStatusAdministrator) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(g.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (g GiveawayParticipantStatusAdministrator) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(g.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (g GiveawayParticipantStatusAdministrator) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(g.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (g GiveawayParticipantStatusAdministrator) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, g.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (g GiveawayParticipantStatusAdministrator) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(g.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (g GiveawayParticipantStatusAdministrator) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(g.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (g GiveawayParticipantStatusAdministrator) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(g.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (g GiveawayParticipantStatusAdministrator) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(g.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (g GiveawayParticipantStatusAdministrator) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(g.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (g GiveawayParticipantStatusAdministrator) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(g.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (g GiveawayParticipantStatusAdministrator) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(g.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (g GiveawayParticipantStatusAdministrator) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(g.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (g GiveawayParticipantStatusAdministrator) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(g.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (g GiveawayParticipantStatusAdministrator) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(g.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (g GiveawayParticipantStatusAdministrator) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(g.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (g GiveawayParticipantStatusAdministrator) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, g.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (g GiveawayParticipantStatusAdministrator) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(g.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (g GiveawayParticipantStatusAdministrator) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(g.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (g GiveawayParticipantStatusAdministrator) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(g.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (g GiveawayParticipantStatusAdministrator) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(g.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (g GiveawayParticipantStatusAdministrator) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, g.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (g GiveawayParticipantStatusAdministrator) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, g.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (g GiveawayParticipantStatusAdministrator) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, g.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (g GiveawayParticipantStatusAdministrator) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(g.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (g GiveawayParticipantStatusAdministrator) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(g.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (g GiveawayParticipantStatusAdministrator) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(g.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (g GiveawayParticipantStatusAdministrator) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(g.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (g GiveawayParticipantStatusAdministrator) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(g.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (g GiveawayParticipantStatusAdministrator) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(g.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (g GiveawayParticipantStatusAdministrator) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, g.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (g GiveawayParticipantStatusAdministrator) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(g.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (g GiveawayParticipantStatusAdministrator) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(g.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (g GiveawayParticipantStatusAdministrator) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(g.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (g GiveawayParticipantStatusAdministrator) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(g.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (g GiveawayParticipantStatusAdministrator) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(g.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (g GiveawayParticipantStatusAdministrator) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(g.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (g GiveawayParticipantStatusAdministrator) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(g.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (g GiveawayParticipantStatusAdministrator) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(g.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (g GiveawayParticipantStatusAdministrator) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(g.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (g GiveawayParticipantStatusAdministrator) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(g.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (g GiveawayParticipantStatusAdministrator) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(g.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (g GiveawayParticipantStatusAdministrator) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(g.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (g GiveawayParticipantStatusAdministrator) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(g.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (g GiveawayParticipantStatusAdministrator) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(g.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (g GiveawayParticipantStatusAdministrator) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(g.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (g GiveawayParticipantStatusAdministrator) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(g.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (g GiveawayParticipantStatusAdministrator) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(g.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (g GiveawayParticipantStatusAdministrator) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(g.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (g GiveawayParticipantStatusAdministrator) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(g.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (g GiveawayParticipantStatusAdministrator) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(g.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (g GiveawayParticipantStatusAdministrator) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(g.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (g GiveawayParticipantStatusAdministrator) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(g.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (g GiveawayParticipantStatusAdministrator) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(g.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (g GiveawayParticipantStatusAdministrator) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(g.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (g GiveawayParticipantStatusAdministrator) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(g.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (g GiveawayParticipantStatusAdministrator) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(g.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (g GiveawayParticipantStatusAdministrator) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(g.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (g GiveawayParticipantStatusAdministrator) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(g.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (g GiveawayParticipantStatusAdministrator) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(g.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (g GiveawayParticipantStatusAdministrator) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(g.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (g GiveawayParticipantStatusAdministrator) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(g.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (g GiveawayParticipantStatusAdministrator) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(g.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (g GiveawayParticipantStatusAdministrator) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(g.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (g GiveawayParticipantStatusAdministrator) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(g.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (g GiveawayParticipantStatusAdministrator) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(g.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (g GiveawayParticipantStatusAdministrator) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(g.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (g GiveawayParticipantStatusAdministrator) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(g.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (g GiveawayParticipantStatusAdministrator) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(g.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (g GiveawayParticipantStatusAdministrator) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, g.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (g GiveawayParticipantStatusAdministrator) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(g.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (g GiveawayParticipantStatusAdministrator) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(g.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (g GiveawayParticipantStatusAdministrator) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(g.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (g GiveawayParticipantStatusAdministrator) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(g.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (g GiveawayParticipantStatusAdministrator) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(g.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (g GiveawayParticipantStatusAdministrator) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(g.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (g GiveawayParticipantStatusAdministrator) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(g.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (g GiveawayParticipantStatusAdministrator) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, g.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (g GiveawayParticipantStatusAdministrator) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(g.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (g GiveawayParticipantStatusAdministrator) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(g.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (g GiveawayParticipantStatusAdministrator) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(g.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (g GiveawayParticipantStatusAdministrator) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(g.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (g GiveawayParticipantStatusAdministrator) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(g.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (g GiveawayParticipantStatusAdministrator) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(g.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (g GiveawayParticipantStatusAdministrator) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(g.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (g GiveawayParticipantStatusAdministrator) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(g.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (g GiveawayParticipantStatusAdministrator) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(g.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (g GiveawayParticipantStatusAdministrator) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(g.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (g GiveawayParticipantStatusAdministrator) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(g.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (g GiveawayParticipantStatusAdministrator) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(g.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (g GiveawayParticipantStatusAdministrator) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(g.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (g GiveawayParticipantStatusAdministrator) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(g.ChatId, messageIds, forceRead, opts)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (g GiveawayPrizePremium) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(g.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiveawayPrizePremium) GiftPremiumWithStars(client *Client, userId int64, starCount int64, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, g.MonthCount, text)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GiveawayPrizeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, g.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (g GiveawayPrizeStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, g.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (g GiveawayPrizeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, g.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (g GiveawayPrizeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, g.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (g GiveawayPrizeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, g.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GiveawayPrizeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, g.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (g GiveawayPrizeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, g.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (g GiveawayPrizeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, g.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GiveawayPrizeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, g.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (g GiveawayPrizeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, g.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (g GiveawayPrizeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, g.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (g GiveawayPrizeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, g.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (g GiveawayPrizeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, g.StarCount)
}

// AddChatFolderByInviteLink is a helper method for Client.AddChatFolderByInviteLink
func (g GroupCall) AddChatFolderByInviteLink(client *Client, chatIds []int64) (*Ok, error) {
	return client.AddChatFolderByInviteLink(g.InviteLink, chatIds)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GroupCall) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(g.Id, starCount)
}

// BanParticipants is a helper method for Client.BanGroupCallParticipants
func (g GroupCall) BanParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(g.Id, userIds)
}

// CheckChatFolderInviteLink is a helper method for Client.CheckChatFolderInviteLink
func (g GroupCall) CheckChatFolderInviteLink(client *Client) (*ChatFolderInviteLinkInfo, error) {
	return client.CheckChatFolderInviteLink(g.InviteLink)
}

// CheckChatInviteLink is a helper method for Client.CheckChatInviteLink
func (g GroupCall) CheckChatInviteLink(client *Client) (*ChatInviteLinkInfo, error) {
	return client.CheckChatInviteLink(g.InviteLink)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (g GroupCall) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(g.Id)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (g GroupCall) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, g.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (g GroupCall) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, g.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (g GroupCall) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(g.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (g GroupCall) CreateVideoChat(client *Client, chatId int64, startDate int32) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, g.Title, startDate, g.IsRtmpStream)
}

// DecryptData is a helper method for Client.DecryptGroupCallData
func (g GroupCall) DecryptData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(g.Id, participantId, data, opts)
}

// DeleteChatFolderInviteLink is a helper method for Client.DeleteChatFolderInviteLink
func (g GroupCall) DeleteChatFolderInviteLink(client *Client, chatFolderId int32) (*Ok, error) {
	return client.DeleteChatFolderInviteLink(chatFolderId, g.InviteLink)
}

// DeleteMessages is a helper method for Client.DeleteGroupCallMessages
func (g GroupCall) DeleteMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(g.Id, messageIds, reportSpam)
}

// DeleteMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (g GroupCall) DeleteMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(g.Id, senderId, reportSpam)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (g GroupCall) DeleteRevokedChatInviteLink(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(chatId, g.InviteLink)
}

// DiscardCall is a helper method for Client.DiscardCall
func (g GroupCall) DiscardCall(client *Client, callId int32, isDisconnected bool, isVideo bool, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, g.InviteLink, g.Duration, isVideo, connectionId)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (g GroupCall) EditChatFolderInviteLink(client *Client, chatFolderId int32, name string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, g.InviteLink, name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (g GroupCall) EditChatInviteLink(client *Client, chatId int64, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, g.InviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (g GroupCall) EditChatSubscriptionInviteLink(client *Client, chatId int64, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, g.InviteLink, name)
}

// EncryptData is a helper method for Client.EncryptGroupCallData
func (g GroupCall) EncryptData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(g.Id, dataChannel, data, unencryptedPrefixSize)
}

// End is a helper method for Client.EndGroupCall
func (g GroupCall) End(client *Client) (*Ok, error) {
	return client.EndGroupCall(g.Id)
}

// EndRecording is a helper method for Client.EndGroupCallRecording
func (g GroupCall) EndRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(g.Id)
}

// EndScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (g GroupCall) EndScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(g.Id)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (g GroupCall) GetChatInviteLink(client *Client, chatId int64) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(chatId, g.InviteLink)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (g GroupCall) GetChatInviteLinkMembers(client *Client, chatId int64, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(chatId, g.InviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (g GroupCall) GetChatJoinRequests(client *Client, chatId int64, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, g.InviteLink, query, limit, opts)
}

// Get is a helper method for Client.GetGroupCall
func (g GroupCall) Get(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(g.Id)
}

// GetStreams is a helper method for Client.GetGroupCallStreams
func (g GroupCall) GetStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(g.Id)
}

// GetStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (g GroupCall) GetStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(g.Id, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (g GroupCall) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(g.Id)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (g GroupCall) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(g.Id)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (g GroupCall) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(g.Id)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (g GroupCall) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(g.Title)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (g GroupCall) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(g.Id, canSelfUnmute)
}

// InviteParticipant is a helper method for Client.InviteGroupCallParticipant
func (g GroupCall) InviteParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(g.Id, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (g GroupCall) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(g.Id, userIds)
}

// JoinChatByInviteLink is a helper method for Client.JoinChatByInviteLink
func (g GroupCall) JoinChatByInviteLink(client *Client) (*Chat, error) {
	return client.JoinChatByInviteLink(g.InviteLink)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (g GroupCall) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(g.Id, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (g GroupCall) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(g.Id, joinParameters, inviteHash, opts)
}

// Leave is a helper method for Client.LeaveGroupCall
func (g GroupCall) Leave(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(g.Id)
}

// LoadParticipants is a helper method for Client.LoadGroupCallParticipants
func (g GroupCall) LoadParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(g.Id, limit)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (g GroupCall) ProcessChatJoinRequests(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(chatId, g.InviteLink, approve)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (g GroupCall) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(g.Id)
}

// ResendMessages is a helper method for Client.ResendMessages
func (g GroupCall) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, g.PaidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (g GroupCall) RevokeChatInviteLink(client *Client, chatId int64) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(chatId, g.InviteLink)
}

// RevokeInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (g GroupCall) RevokeInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(g.Id)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (g GroupCall) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, g.Duration, g.PaidMessageStarCount)
}

// SendMessage is a helper method for Client.SendGroupCallMessage
func (g GroupCall) SendMessage(client *Client, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(g.Id, text, g.PaidMessageStarCount)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (g GroupCall) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, g.PaidMessageStarCount)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (g GroupCall) SetChatMessageSender(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatMessageSender(chatId, g.MessageSenderId)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (g GroupCall) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, g.PaidMessageStarCount)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (g GroupCall) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, g.Title)
}

// SetPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (g GroupCall) SetPaidMessageStarCount(client *Client) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(g.Id, g.PaidMessageStarCount)
}

// SetParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (g GroupCall) SetParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(g.Id, audioSource, isSpeaking)
}

// SetParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (g GroupCall) SetParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(g.Id, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (g GroupCall) SetLiveStoryMessageSender(client *Client) (*Ok, error) {
	return client.SetLiveStoryMessageSender(g.Id, g.MessageSenderId)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (g GroupCall) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, g.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (g GroupCall) SetVideoChatTitle(client *Client) (*Ok, error) {
	return client.SetVideoChatTitle(g.Id, g.Title)
}

// StartRecording is a helper method for Client.StartGroupCallRecording
func (g GroupCall) StartRecording(client *Client, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(g.Id, g.Title, recordVideo, usePortraitOrientation)
}

// StartScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (g GroupCall) StartScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(g.Id, audioSourceId, payload)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (g GroupCall) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, g.IsRtmpStream, enableMessages, g.PaidMessageStarCount)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (g GroupCall) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(g.Id)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (g GroupCall) ToggleBotUsernameIsActive(client *Client, botUserId int64, username string) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, username, g.IsActive)
}

// ToggleAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (g GroupCall) ToggleAreMessagesAllowed(client *Client) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(g.Id, g.AreMessagesAllowed)
}

// ToggleIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (g GroupCall) ToggleIsMyVideoEnabled(client *Client) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(g.Id, g.IsMyVideoEnabled)
}

// ToggleIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (g GroupCall) ToggleIsMyVideoPaused(client *Client) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(g.Id, g.IsMyVideoPaused)
}

// ToggleParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (g GroupCall) ToggleParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(g.Id, participantId, isHandRaised)
}

// ToggleParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (g GroupCall) ToggleParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(g.Id, participantId, isMuted)
}

// ToggleScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (g GroupCall) ToggleScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(g.Id, isPaused)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (g GroupCall) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, username string) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, username, g.IsActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (g GroupCall) ToggleUsernameIsActive(client *Client, username string) (*Ok, error) {
	return client.ToggleUsernameIsActive(username, g.IsActive)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (g GroupCall) ToggleVideoChatEnabledStartNotification(client *Client) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(g.Id, g.EnabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (g GroupCall) ToggleVideoChatMuteNewParticipants(client *Client) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(g.Id, g.MuteNewParticipants)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (g GroupCallInfo) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(g.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (g GroupCallInfo) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(g.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (g GroupCallInfo) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(g.GroupCallId)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (g GroupCallInfo) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(g.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (g GroupCallInfo) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(g.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (g GroupCallInfo) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(g.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (g GroupCallInfo) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(g.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (g GroupCallInfo) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(g.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (g GroupCallInfo) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(g.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (g GroupCallInfo) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(g.GroupCallId)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (g GroupCallInfo) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(g.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (g GroupCallInfo) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(g.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (g GroupCallInfo) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(g.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (g GroupCallInfo) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(g.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (g GroupCallInfo) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(g.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (g GroupCallInfo) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(g.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (g GroupCallInfo) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(g.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (g GroupCallInfo) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(g.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (g GroupCallInfo) InviteVideoChatParticipants(client *Client, userIds []int64) (*Ok, error) {
	return client.InviteVideoChatParticipants(g.GroupCallId, userIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (g GroupCallInfo) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(g.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (g GroupCallInfo) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(g.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (g GroupCallInfo) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(g.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (g GroupCallInfo) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(g.GroupCallId, limit)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (g GroupCallInfo) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(g.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (g GroupCallInfo) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(g.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (g GroupCallInfo) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(g.GroupCallId, text, paidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (g GroupCallInfo) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(g.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (g GroupCallInfo) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(g.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (g GroupCallInfo) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(g.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (g GroupCallInfo) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(g.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (g GroupCallInfo) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(g.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (g GroupCallInfo) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(g.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (g GroupCallInfo) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(g.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (g GroupCallInfo) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(g.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (g GroupCallInfo) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(g.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (g GroupCallInfo) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(g.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (g GroupCallInfo) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(g.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (g GroupCallInfo) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(g.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (g GroupCallInfo) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(g.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (g GroupCallInfo) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(g.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (g GroupCallInfo) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(g.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (g GroupCallInfo) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(g.GroupCallId, muteNewParticipants)
}

// GetPushReceiverId is a helper method for Client.GetPushReceiverId
func (g GroupCallJoinParameters) GetPushReceiverId(client *Client) (*PushReceiverId, error) {
	return client.GetPushReceiverId(g.Payload)
}

// ProcessPushNotification is a helper method for Client.ProcessPushNotification
func (g GroupCallJoinParameters) ProcessPushNotification(client *Client) (*Ok, error) {
	return client.ProcessPushNotification(g.Payload)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (g GroupCallJoinParameters) StartGroupCallScreenSharing(client *Client, groupCallId int32) (*Text, error) {
	return client.StartGroupCallScreenSharing(groupCallId, g.AudioSourceId, g.Payload)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (g GroupCallJoinParameters) ToggleGroupCallIsMyVideoEnabled(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(groupCallId, g.IsMyVideoEnabled)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (g GroupCallJoinParameters) ToggleGroupCallParticipantIsMuted(client *Client, groupCallId int32, participantId *MessageSender) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(groupCallId, participantId, g.IsMuted)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (g GroupCallMessage) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, g.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (g GroupCallMessage) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, g.SenderId)
}

// DeletesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (g GroupCallMessage) DeletesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, g.SenderId, reportSpam)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (g GroupCallMessage) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, g.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (g GroupCallMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, g.Date)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (g GroupCallMessage) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(g.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (g GroupCallMessage) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(g.Text)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (g GroupCallMessage) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, g.Date)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (g GroupCallMessage) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, g.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (g GroupCallMessage) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(g.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (g GroupCallMessage) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, g.Text, isPrivate)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (g GroupCallMessage) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, g.SenderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (g GroupCallMessage) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, g.PaidMessageStarCount, opts)
}

// SearchQuote is a helper method for Client.SearchQuote
func (g GroupCallMessage) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(g.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (g GroupCallMessage) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, g.Text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (g GroupCallMessage) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, g.PaidMessageStarCount)
}

// Send is a helper method for Client.SendGroupCallMessage
func (g GroupCallMessage) Send(client *Client, groupCallId int32) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, g.Text, g.PaidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (g GroupCallMessage) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, g.Text)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (g GroupCallMessage) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, g.PaidMessageStarCount)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (g GroupCallMessage) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, g.PaidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (g GroupCallMessage) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, g.PaidMessageStarCount)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (g GroupCallMessage) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(g.SenderId, opts)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (g GroupCallMessage) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, isRtmpStream, enableMessages, g.PaidMessageStarCount)
}

// TranslateText is a helper method for Client.TranslateText
func (g GroupCallMessage) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(g.Text, toLanguageCode)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (g GroupCallParticipant) DecryptGroupCallData(client *Client, groupCallId int32, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, g.ParticipantId, data, opts)
}

// SetBio is a helper method for Client.SetBio
func (g GroupCallParticipant) SetBio(client *Client) (*Ok, error) {
	return client.SetBio(g.Bio)
}

// SetBusinessAccountBio is a helper method for Client.SetBusinessAccountBio
func (g GroupCallParticipant) SetBusinessAccountBio(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountBio(businessConnectionId, g.Bio)
}

// SetIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (g GroupCallParticipant) SetIsSpeaking(client *Client, groupCallId int32, audioSource int32) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(groupCallId, audioSource, g.IsSpeaking)
}

// SetVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (g GroupCallParticipant) SetVolumeLevel(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(groupCallId, g.ParticipantId, g.VolumeLevel)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (g GroupCallParticipant) StartGroupCallScreenSharing(client *Client, groupCallId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(groupCallId, g.AudioSourceId, payload)
}

// ToggleIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (g GroupCallParticipant) ToggleIsHandRaised(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(groupCallId, g.ParticipantId, g.IsHandRaised)
}

// ToggleIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (g GroupCallParticipant) ToggleIsMuted(client *Client, groupCallId int32, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(groupCallId, g.ParticipantId, isMuted)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (g GroupCallParticipantVideoInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(chatId, g.IsPaused)
}

// ToggleDownloadIsPaused is a helper method for Client.ToggleDownloadIsPaused
func (g GroupCallParticipantVideoInfo) ToggleDownloadIsPaused(client *Client, fileId int32) (*Ok, error) {
	return client.ToggleDownloadIsPaused(fileId, g.IsPaused)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (g GroupCallParticipantVideoInfo) ToggleGroupCallScreenSharingIsPaused(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(groupCallId, g.IsPaused)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (g GroupCallRecentSpeaker) DecryptGroupCallData(client *Client, groupCallId int32, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, g.ParticipantId, data, opts)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (g GroupCallRecentSpeaker) SetGroupCallParticipantIsSpeaking(client *Client, groupCallId int32, audioSource int32) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(groupCallId, audioSource, g.IsSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (g GroupCallRecentSpeaker) SetGroupCallParticipantVolumeLevel(client *Client, groupCallId int32, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(groupCallId, g.ParticipantId, volumeLevel)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (g GroupCallRecentSpeaker) ToggleGroupCallParticipantIsHandRaised(client *Client, groupCallId int32, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(groupCallId, g.ParticipantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (g GroupCallRecentSpeaker) ToggleGroupCallParticipantIsMuted(client *Client, groupCallId int32, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(groupCallId, g.ParticipantId, isMuted)
}

// GetSegment is a helper method for Client.GetGroupCallStreamSegment
func (g GroupCallStream) GetSegment(client *Client, groupCallId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(groupCallId, g.TimeOffset, g.Scale, g.ChannelId, opts)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (g GroupCallStream) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, g.Scale, chatId)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (h HttpUrl) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, h.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (h HttpUrl) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, h.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (h HttpUrl) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, h.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (h HttpUrl) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(h.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (h HttpUrl) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(h.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (h HttpUrl) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, h.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (h HttpUrl) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(h.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (h HttpUrl) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, h.Url, parameters, opts)
}

// RegisterUser is a helper method for Client.RegisterUser
func (i ImportedContact) RegisterUser(client *Client, disableNotification bool) (*Ok, error) {
	return client.RegisterUser(i.FirstName, i.LastName, disableNotification)
}

// SearchUserByPhoneNumber is a helper method for Client.SearchUserByPhoneNumber
func (i ImportedContact) SearchUserByPhoneNumber(client *Client, onlyLocal bool) (*User, error) {
	return client.SearchUserByPhoneNumber(i.PhoneNumber, onlyLocal)
}

// SendPhoneNumberCode is a helper method for Client.SendPhoneNumberCode
func (i ImportedContact) SendPhoneNumberCode(client *Client, typeField *PhoneNumberCodeType, opts *SendPhoneNumberCodeOpts) (*AuthenticationCodeInfo, error) {
	return client.SendPhoneNumberCode(i.PhoneNumber, typeField, opts)
}

// SetAuthenticationPhoneNumber is a helper method for Client.SetAuthenticationPhoneNumber
func (i ImportedContact) SetAuthenticationPhoneNumber(client *Client, opts *SetAuthenticationPhoneNumberOpts) (*Ok, error) {
	return client.SetAuthenticationPhoneNumber(i.PhoneNumber, opts)
}

// SetBusinessAccountName is a helper method for Client.SetBusinessAccountName
func (i ImportedContact) SetBusinessAccountName(client *Client, businessConnectionId string) (*Ok, error) {
	return client.SetBusinessAccountName(businessConnectionId, i.FirstName, i.LastName)
}

// SetName is a helper method for Client.SetName
func (i ImportedContact) SetName(client *Client) (*Ok, error) {
	return client.SetName(i.FirstName, i.LastName)
}

// SetUserNote is a helper method for Client.SetUserNote
func (i ImportedContact) SetUserNote(client *Client, userId int64) (*Ok, error) {
	return client.SetUserNote(userId, i.Note)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i ImportedContacts) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, i.UserIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i ImportedContacts) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(i.UserIds, title, messageAutoDeleteTime)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i ImportedContacts) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, i.UserIds, opts)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (i ImportedContacts) InviteVideoChatParticipants(client *Client, groupCallId int32) (*Ok, error) {
	return client.InviteVideoChatParticipants(groupCallId, i.UserIds)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (i ImportedContacts) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(i.UserIds)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (i ImportedContacts) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(i.UserIds)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (i InlineKeyboardButton) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, i.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineKeyboardButton) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, i.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (i InlineKeyboardButton) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(i.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (i InlineKeyboardButton) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(i.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (i InlineKeyboardButton) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(i.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (i InlineKeyboardButton) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, i.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InlineKeyboardButton) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, i.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (i InlineKeyboardButton) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, i.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (i InlineKeyboardButton) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(i.Text, inputLanguageCodes)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (i InlineKeyboardButtonTypeCallback) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, i.Data, opts)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (i InlineKeyboardButtonTypeCallback) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, i.Data, unencryptedPrefixSize)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (i InlineKeyboardButtonTypeCallback) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, i.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (i InlineKeyboardButtonTypeCallback) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, i.Data)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (i InlineKeyboardButtonTypeCallbackWithPassword) DecryptGroupCallData(client *Client, groupCallId int32, participantId *MessageSender, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(groupCallId, participantId, i.Data, opts)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (i InlineKeyboardButtonTypeCallbackWithPassword) EncryptGroupCallData(client *Client, groupCallId int32, dataChannel *GroupCallDataChannel, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(groupCallId, dataChannel, i.Data, unencryptedPrefixSize)
}

// SendCallSignalingData is a helper method for Client.SendCallSignalingData
func (i InlineKeyboardButtonTypeCallbackWithPassword) SendCallSignalingData(client *Client, callId int32) (*Ok, error) {
	return client.SendCallSignalingData(callId, i.Data)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (i InlineKeyboardButtonTypeCallbackWithPassword) WriteGeneratedFilePart(client *Client, generationId string, offset int64) (*Ok, error) {
	return client.WriteGeneratedFilePart(generationId, offset, i.Data)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (i InlineKeyboardButtonTypeCopyText) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, i.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineKeyboardButtonTypeCopyText) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, i.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (i InlineKeyboardButtonTypeCopyText) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(i.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (i InlineKeyboardButtonTypeCopyText) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(i.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (i InlineKeyboardButtonTypeCopyText) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(i.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (i InlineKeyboardButtonTypeCopyText) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, i.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InlineKeyboardButtonTypeCopyText) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, i.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (i InlineKeyboardButtonTypeCopyText) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, i.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (i InlineKeyboardButtonTypeCopyText) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(i.Text, inputLanguageCodes)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineKeyboardButtonTypeLoginUrl) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InlineKeyboardButtonTypeLoginUrl) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InlineKeyboardButtonTypeLoginUrl) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InlineKeyboardButtonTypeLoginUrl) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InlineKeyboardButtonTypeLoginUrl) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InlineKeyboardButtonTypeLoginUrl) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InlineKeyboardButtonTypeLoginUrl) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InlineKeyboardButtonTypeLoginUrl) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (i InlineKeyboardButtonTypeSwitchInline) GetAllStickerEmojis(client *Client, stickerType *StickerType, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, i.Query, chatId, returnOnlyMainEmoji)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (i InlineKeyboardButtonTypeSwitchInline) GetChatEventLog(client *Client, chatId int64, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, i.Query, fromEventId, limit, userIds, opts)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (i InlineKeyboardButtonTypeSwitchInline) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, i.Query, limit, opts)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (i InlineKeyboardButtonTypeSwitchInline) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, i.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (i InlineKeyboardButtonTypeSwitchInline) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, i.Query, offset, opts)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (i InlineKeyboardButtonTypeSwitchInline) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(i.Query)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (i InlineKeyboardButtonTypeSwitchInline) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(i.Query)
}

// GetStickers is a helper method for Client.GetStickers
func (i InlineKeyboardButtonTypeSwitchInline) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, i.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (i InlineKeyboardButtonTypeSwitchInline) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, i.Query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (i InlineKeyboardButtonTypeSwitchInline) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, i.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (i InlineKeyboardButtonTypeSwitchInline) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, i.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (i InlineKeyboardButtonTypeSwitchInline) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(i.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (i InlineKeyboardButtonTypeSwitchInline) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(i.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (i InlineKeyboardButtonTypeSwitchInline) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(i.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (i InlineKeyboardButtonTypeSwitchInline) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, offset string, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(i.Query, onlyActive, onlyCompleted, offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (i InlineKeyboardButtonTypeSwitchInline) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, i.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (i InlineKeyboardButtonTypeSwitchInline) SearchMessages(client *Client, offset string, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(i.Query, offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (i InlineKeyboardButtonTypeSwitchInline) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(i.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (i InlineKeyboardButtonTypeSwitchInline) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(i.Query)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (i InlineKeyboardButtonTypeSwitchInline) SearchPublicPosts(client *Client, offset string, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(i.Query, offset, limit, starCount)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (i InlineKeyboardButtonTypeSwitchInline) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(i.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (i InlineKeyboardButtonTypeSwitchInline) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, i.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (i InlineKeyboardButtonTypeSwitchInline) SearchSecretMessages(client *Client, chatId int64, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, i.Query, offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (i InlineKeyboardButtonTypeSwitchInline) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, i.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (i InlineKeyboardButtonTypeSwitchInline) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, i.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (i InlineKeyboardButtonTypeSwitchInline) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, i.Query, limit, returnNoneForEmptyQuery)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineKeyboardButtonTypeUrl) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InlineKeyboardButtonTypeUrl) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InlineKeyboardButtonTypeUrl) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InlineKeyboardButtonTypeUrl) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InlineKeyboardButtonTypeUrl) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InlineKeyboardButtonTypeUrl) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InlineKeyboardButtonTypeUrl) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InlineKeyboardButtonTypeUrl) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InlineKeyboardButtonTypeUser) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, i.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (i InlineKeyboardButtonTypeUser) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(i.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (i InlineKeyboardButtonTypeUser) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(i.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (i InlineKeyboardButtonTypeUser) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(i.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (i InlineKeyboardButtonTypeUser) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(i.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (i InlineKeyboardButtonTypeUser) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(i.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (i InlineKeyboardButtonTypeUser) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(i.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineKeyboardButtonTypeUser) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(i.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (i InlineKeyboardButtonTypeUser) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(i.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (i InlineKeyboardButtonTypeUser) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(i.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InlineKeyboardButtonTypeUser) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, i.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (i InlineKeyboardButtonTypeUser) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(i.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (i InlineKeyboardButtonTypeUser) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, i.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (i InlineKeyboardButtonTypeUser) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(i.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (i InlineKeyboardButtonTypeUser) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(i.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (i InlineKeyboardButtonTypeUser) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(i.UserId)
}

// GetUser is a helper method for Client.GetUser
func (i InlineKeyboardButtonTypeUser) GetUser(client *Client) (*User, error) {
	return client.GetUser(i.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (i InlineKeyboardButtonTypeUser) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, i.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (i InlineKeyboardButtonTypeUser) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(i.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (i InlineKeyboardButtonTypeUser) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(i.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (i InlineKeyboardButtonTypeUser) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(i.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (i InlineKeyboardButtonTypeUser) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(i.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InlineKeyboardButtonTypeUser) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(i.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (i InlineKeyboardButtonTypeUser) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, i.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InlineKeyboardButtonTypeUser) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, i.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (i InlineKeyboardButtonTypeUser) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, i.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (i InlineKeyboardButtonTypeUser) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(i.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (i InlineKeyboardButtonTypeUser) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(i.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (i InlineKeyboardButtonTypeUser) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(i.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InlineKeyboardButtonTypeUser) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, i.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (i InlineKeyboardButtonTypeUser) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, i.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (i InlineKeyboardButtonTypeUser) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(i.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (i InlineKeyboardButtonTypeUser) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(i.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (i InlineKeyboardButtonTypeUser) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(i.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (i InlineKeyboardButtonTypeUser) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(i.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (i InlineKeyboardButtonTypeUser) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(i.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (i InlineKeyboardButtonTypeUser) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(i.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (i InlineKeyboardButtonTypeUser) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(i.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (i InlineKeyboardButtonTypeUser) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(i.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (i InlineKeyboardButtonTypeUser) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(i.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (i InlineKeyboardButtonTypeUser) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(i.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (i InlineKeyboardButtonTypeUser) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, i.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (i InlineKeyboardButtonTypeUser) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(i.UserId, stickerFormat, sticker)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineKeyboardButtonTypeWebApp) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InlineKeyboardButtonTypeWebApp) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InlineKeyboardButtonTypeWebApp) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InlineKeyboardButtonTypeWebApp) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InlineKeyboardButtonTypeWebApp) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InlineKeyboardButtonTypeWebApp) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InlineKeyboardButtonTypeWebApp) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InlineKeyboardButtonTypeWebApp) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultAnimation) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultAnimation) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultAnimation) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultAnimation) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultAnimation) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultAnimation) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultAnimation) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultAnimation) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultAnimation) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineQueryResultArticle) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InlineQueryResultArticle) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultArticle) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultArticle) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultArticle) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultArticle) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InlineQueryResultArticle) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InlineQueryResultArticle) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InlineQueryResultArticle) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultArticle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InlineQueryResultArticle) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InlineQueryResultArticle) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InlineQueryResultArticle) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InlineQueryResultArticle) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InlineQueryResultArticle) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultArticle) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultArticle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultArticle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultArticle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultDocument) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultDocument) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultDocument) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultDocument) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultDocument) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InlineQueryResultDocument) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InlineQueryResultDocument) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultDocument) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultDocument) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultDocument) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultDocument) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultLocation) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultLocation) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultLocation) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultLocation) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (i InlineQueryResultLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(i.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (i InlineQueryResultLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(i.Location, zoom, width, height, scale, chatId)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultLocation) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultLocation) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultLocation) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultLocation) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultLocation) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultPhoto) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultPhoto) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultPhoto) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultPhoto) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultPhoto) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InlineQueryResultPhoto) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InlineQueryResultPhoto) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultPhoto) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultPhoto) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultPhoto) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultPhoto) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AnswerInlineQuery is a helper method for Client.AnswerInlineQuery
func (i InlineQueryResults) AnswerInlineQuery(client *Client, isPersonal bool, results []*InputInlineQueryResult, cacheTime int32, opts *AnswerInlineQueryOpts) (*Ok, error) {
	return client.AnswerInlineQuery(i.InlineQueryId, isPersonal, results, cacheTime, i.NextOffset, opts)
}

// AddLogMessage is a helper method for Client.AddLogMessage
func (i InlineQueryResultsButton) AddLogMessage(client *Client, verbosityLevel int32) (*Ok, error) {
	return client.AddLogMessage(verbosityLevel, i.Text)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineQueryResultsButton) AnswerCallbackQuery(client *Client, callbackQueryId string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, i.Text, showAlert, url, cacheTime)
}

// GetKeywordEmojis is a helper method for Client.GetKeywordEmojis
func (i InlineQueryResultsButton) GetKeywordEmojis(client *Client, inputLanguageCodes []string) (*Emojis, error) {
	return client.GetKeywordEmojis(i.Text, inputLanguageCodes)
}

// GetTextEntities is a helper method for Client.GetTextEntities
func (i InlineQueryResultsButton) GetTextEntities(client *Client) (*TextEntities, error) {
	return client.GetTextEntities(i.Text)
}

// ParseTextEntities is a helper method for Client.ParseTextEntities
func (i InlineQueryResultsButton) ParseTextEntities(client *Client, parseMode *TextParseMode) (*FormattedText, error) {
	return client.ParseTextEntities(i.Text, parseMode)
}

// ReportChat is a helper method for Client.ReportChat
func (i InlineQueryResultsButton) ReportChat(client *Client, chatId int64, optionId string, messageIds []int64) (*ReportChatResult, error) {
	return client.ReportChat(chatId, optionId, messageIds, i.Text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (i InlineQueryResultsButton) ReportChatPhoto(client *Client, chatId int64, fileId int32, reason *ReportReason) (*Ok, error) {
	return client.ReportChatPhoto(chatId, fileId, reason, i.Text)
}

// ReportStory is a helper method for Client.ReportStory
func (i InlineQueryResultsButton) ReportStory(client *Client, storyPosterChatId int64, storyId int32, optionId string) (*ReportStoryResult, error) {
	return client.ReportStory(storyPosterChatId, storyId, optionId, i.Text)
}

// SearchEmojis is a helper method for Client.SearchEmojis
func (i InlineQueryResultsButton) SearchEmojis(client *Client, inputLanguageCodes []string) (*EmojiKeywords, error) {
	return client.SearchEmojis(i.Text, inputLanguageCodes)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (i InlineQueryResultsButtonTypeStartBot) SendBotStartMessage(client *Client, botUserId int64, chatId int64) (*Message, error) {
	return client.SendBotStartMessage(botUserId, chatId, i.Parameter)
}

// AnswerCallbackQuery is a helper method for Client.AnswerCallbackQuery
func (i InlineQueryResultsButtonTypeWebApp) AnswerCallbackQuery(client *Client, callbackQueryId string, text string, showAlert bool, cacheTime int32) (*Ok, error) {
	return client.AnswerCallbackQuery(callbackQueryId, text, showAlert, i.Url, cacheTime)
}

// CheckWebAppFileDownload is a helper method for Client.CheckWebAppFileDownload
func (i InlineQueryResultsButtonTypeWebApp) CheckWebAppFileDownload(client *Client, botUserId int64, fileName string) (*Ok, error) {
	return client.CheckWebAppFileDownload(botUserId, fileName, i.Url)
}

// DisconnectAffiliateProgram is a helper method for Client.DisconnectAffiliateProgram
func (i InlineQueryResultsButtonTypeWebApp) DisconnectAffiliateProgram(client *Client, affiliate *AffiliateType) (*ConnectedAffiliateProgram, error) {
	return client.DisconnectAffiliateProgram(affiliate, i.Url)
}

// GetChatBoostLinkInfo is a helper method for Client.GetChatBoostLinkInfo
func (i InlineQueryResultsButtonTypeWebApp) GetChatBoostLinkInfo(client *Client) (*ChatBoostLinkInfo, error) {
	return client.GetChatBoostLinkInfo(i.Url)
}

// GetMessageLinkInfo is a helper method for Client.GetMessageLinkInfo
func (i InlineQueryResultsButtonTypeWebApp) GetMessageLinkInfo(client *Client) (*MessageLinkInfo, error) {
	return client.GetMessageLinkInfo(i.Url)
}

// GetWebAppUrl is a helper method for Client.GetWebAppUrl
func (i InlineQueryResultsButtonTypeWebApp) GetWebAppUrl(client *Client, botUserId int64, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppUrl(botUserId, i.Url, parameters)
}

// GetWebPageInstantView is a helper method for Client.GetWebPageInstantView
func (i InlineQueryResultsButtonTypeWebApp) GetWebPageInstantView(client *Client, onlyLocal bool) (*WebPageInstantView, error) {
	return client.GetWebPageInstantView(i.Url, onlyLocal)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (i InlineQueryResultsButtonTypeWebApp) OpenWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(chatId, botUserId, i.Url, parameters, opts)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultVideo) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultVideo) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultVideo) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, i.Description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultVideo) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultVideo) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetBotInfoDescription is a helper method for Client.SetBotInfoDescription
func (i InlineQueryResultVideo) SetBotInfoDescription(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotInfoDescription(botUserId, languageCode, i.Description)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (i InlineQueryResultVideo) SetChatDescription(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatDescription(chatId, i.Description)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultVideo) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultVideo) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultVideo) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultVideo) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InlineQueryResultVoiceNote) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InlineQueryResultVoiceNote) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InlineQueryResultVoiceNote) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InlineQueryResultVoiceNote) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InlineQueryResultVoiceNote) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InlineQueryResultVoiceNote) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InlineQueryResultVoiceNote) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InlineQueryResultVoiceNote) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InlineQueryResultVoiceNote) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputBackgroundPrevious) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputBackgroundPrevious) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, i.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (i InputBackgroundPrevious) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, i.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (i InputBackgroundPrevious) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, i.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (i InputBackgroundPrevious) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, i.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (i InputBackgroundPrevious) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, i.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (i InputBackgroundPrevious) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(i.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (i InputBackgroundPrevious) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, i.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (i InputBackgroundPrevious) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, i.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (i InputBackgroundPrevious) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, i.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (i InputBackgroundPrevious) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, i.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (i InputBackgroundPrevious) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, i.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (i InputBackgroundPrevious) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, i.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (i InputBackgroundPrevious) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (i InputBackgroundPrevious) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, i.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (i InputBackgroundPrevious) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (i InputBackgroundPrevious) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (i InputBackgroundPrevious) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, i.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (i InputBackgroundPrevious) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (i InputBackgroundPrevious) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, i.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (i InputBackgroundPrevious) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, i.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (i InputBackgroundPrevious) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, i.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (i InputBackgroundPrevious) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, i.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (i InputBackgroundPrevious) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, i.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (i InputBackgroundPrevious) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, i.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (i InputBackgroundPrevious) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, i.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (i InputBackgroundPrevious) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, i.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (i InputBackgroundPrevious) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, i.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (i InputBackgroundPrevious) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, i.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (i InputBackgroundPrevious) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, i.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (i InputBackgroundPrevious) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, i.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (i InputBackgroundPrevious) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, i.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (i InputBackgroundPrevious) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, i.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (i InputBackgroundPrevious) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, i.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (i InputBackgroundPrevious) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, i.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (i InputBackgroundPrevious) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, i.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (i InputBackgroundPrevious) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, i.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (i InputBackgroundPrevious) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, i.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (i InputBackgroundPrevious) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, i.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (i InputBackgroundPrevious) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, i.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (i InputBackgroundPrevious) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, i.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (i InputBackgroundPrevious) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, i.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (i InputBackgroundPrevious) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, i.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (i InputBackgroundPrevious) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, i.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (i InputBackgroundPrevious) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, i.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (i InputBackgroundPrevious) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, i.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (i InputBackgroundPrevious) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, i.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (i InputBackgroundPrevious) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, i.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (i InputBackgroundPrevious) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, i.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (i InputBackgroundPrevious) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, i.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (i InputBackgroundPrevious) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, i.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (i InputBackgroundPrevious) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, i.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (i InputBackgroundPrevious) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, i.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (i InputBackgroundPrevious) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, i.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (i InputBackgroundPrevious) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, i.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (i InputBackgroundPrevious) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(i.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (i InputBackgroundPrevious) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, i.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (i InputBackgroundPrevious) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, i.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (i InputBackgroundPrevious) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, i.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (i InputBackgroundPrevious) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, i.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (i InputBackgroundPrevious) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, i.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (i InputBackgroundPrevious) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, i.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (i InputBackgroundPrevious) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, i.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (i InputBackgroundPrevious) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, i.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (i InputBackgroundPrevious) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, i.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (i InputBackgroundPrevious) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, i.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (i InputBackgroundPrevious) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, i.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (i InputBackgroundPrevious) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, i.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (i InputBackgroundPrevious) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, i.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (i InputBackgroundPrevious) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, i.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (i InputBackgroundPrevious) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, i.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (i InputBackgroundPrevious) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, i.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (i InputBackgroundPrevious) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, i.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (i InputBackgroundPrevious) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, i.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (i InputBackgroundPrevious) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, i.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (i InputBackgroundPrevious) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, i.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (i InputBackgroundPrevious) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, i.MessageId)
}

// RemoveInstalledBackground is a helper method for Client.RemoveInstalledBackground
func (i InputBackgroundRemote) RemoveInstalledBackground(client *Client) (*Ok, error) {
	return client.RemoveInstalledBackground(i.BackgroundId)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputBusinessChatLink) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputBusinessChatLink) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputBusinessChatLink) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputBusinessChatLink) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (i InputBusinessChatLink) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(i.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (i InputBusinessChatLink) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(i.Text)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputBusinessChatLink) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InputBusinessChatLink) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, i.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (i InputBusinessChatLink) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(i.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InputBusinessChatLink) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, i.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (i InputBusinessChatLink) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(i.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (i InputBusinessChatLink) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, i.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (i InputBusinessChatLink) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, i.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputBusinessChatLink) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, i.Text)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputBusinessChatLink) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputBusinessChatLink) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputBusinessChatLink) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputBusinessChatLink) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// TranslateText is a helper method for Client.TranslateText
func (i InputBusinessChatLink) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(i.Text, toLanguageCode)
}

// AddFavoriteSticker is a helper method for Client.AddFavoriteSticker
func (i InputBusinessStartPage) AddFavoriteSticker(client *Client) (*Ok, error) {
	return client.AddFavoriteSticker(i.Sticker)
}

// AddRecentSticker is a helper method for Client.AddRecentSticker
func (i InputBusinessStartPage) AddRecentSticker(client *Client, isAttached bool) (*Stickers, error) {
	return client.AddRecentSticker(isAttached, i.Sticker)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (i InputBusinessStartPage) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, i.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputBusinessStartPage) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, i.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (i InputBusinessStartPage) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(i.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (i InputBusinessStartPage) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, i.Title, startDate, isRtmpStream)
}

// GetStickerEmojis is a helper method for Client.GetStickerEmojis
func (i InputBusinessStartPage) GetStickerEmojis(client *Client) (*Emojis, error) {
	return client.GetStickerEmojis(i.Sticker)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (i InputBusinessStartPage) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(i.Title)
}

// RemoveFavoriteSticker is a helper method for Client.RemoveFavoriteSticker
func (i InputBusinessStartPage) RemoveFavoriteSticker(client *Client) (*Ok, error) {
	return client.RemoveFavoriteSticker(i.Sticker)
}

// RemoveRecentSticker is a helper method for Client.RemoveRecentSticker
func (i InputBusinessStartPage) RemoveRecentSticker(client *Client, isAttached bool) (*Ok, error) {
	return client.RemoveRecentSticker(isAttached, i.Sticker)
}

// RemoveStickerFromSet is a helper method for Client.RemoveStickerFromSet
func (i InputBusinessStartPage) RemoveStickerFromSet(client *Client) (*Ok, error) {
	return client.RemoveStickerFromSet(i.Sticker)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (i InputBusinessStartPage) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, i.Title)
}

// SetStickerEmojis is a helper method for Client.SetStickerEmojis
func (i InputBusinessStartPage) SetStickerEmojis(client *Client, emojis string) (*Ok, error) {
	return client.SetStickerEmojis(i.Sticker, emojis)
}

// SetStickerKeywords is a helper method for Client.SetStickerKeywords
func (i InputBusinessStartPage) SetStickerKeywords(client *Client, keywords []string) (*Ok, error) {
	return client.SetStickerKeywords(i.Sticker, keywords)
}

// SetStickerMaskPosition is a helper method for Client.SetStickerMaskPosition
func (i InputBusinessStartPage) SetStickerMaskPosition(client *Client, opts *SetStickerMaskPositionOpts) (*Ok, error) {
	return client.SetStickerMaskPosition(i.Sticker, opts)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (i InputBusinessStartPage) SetStickerPositionInSet(client *Client, position int32) (*Ok, error) {
	return client.SetStickerPositionInSet(i.Sticker, position)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputBusinessStartPage) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, i.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (i InputBusinessStartPage) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, i.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (i InputBusinessStartPage) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, i.Title, recordVideo, usePortraitOrientation)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (i InputBusinessStartPage) UploadStickerFile(client *Client, userId int64, stickerFormat *StickerFormat) (*File, error) {
	return client.UploadStickerFile(userId, stickerFormat, i.Sticker)
}

// AddSavedAnimation is a helper method for Client.AddSavedAnimation
func (i InputChatPhotoAnimation) AddSavedAnimation(client *Client) (*Ok, error) {
	return client.AddSavedAnimation(i.Animation)
}

// RemoveSavedAnimation is a helper method for Client.RemoveSavedAnimation
func (i InputChatPhotoAnimation) RemoveSavedAnimation(client *Client) (*Ok, error) {
	return client.RemoveSavedAnimation(i.Animation)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (i InputChatThemeEmoji) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, i.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (i InputChatThemeEmoji) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(i.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (i InputChatThemeEmoji) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(i.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (i InputChatThemeEmoji) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, i.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputChatThemeEmoji) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputChatThemeEmoji) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, i.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputChatThemeEmoji) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, i.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (i InputChatThemeEmoji) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, i.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputChatThemeEmoji) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, i.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InputChatThemeEmoji) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, i.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (i InputChatThemeEmoji) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(i.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InputChatThemeEmoji) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, i.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputChatThemeEmoji) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputChatThemeEmoji) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, i.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputChatThemeEmoji) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, i.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (i InputChatThemeEmoji) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(i.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (i InputChatThemeEmoji) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(i.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (i InputChatThemeEmoji) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(i.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (i InputChatThemeEmoji) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(i.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (i InputChatThemeEmoji) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, i.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (i InputChatThemeEmoji) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(i.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (i InputChatThemeEmoji) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(i.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (i InputChatThemeEmoji) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, i.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (i InputChatThemeEmoji) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(i.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (i InputChatThemeEmoji) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, i.Name)
}

// SetOption is a helper method for Client.SetOption
func (i InputChatThemeEmoji) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(i.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (i InputChatThemeEmoji) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, i.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (i InputChatThemeEmoji) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, i.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputChatThemeEmoji) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(i.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputChatThemeEmoji) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, i.Name)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (i InputChatThemeGift) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, i.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (i InputChatThemeGift) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(i.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (i InputChatThemeGift) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(i.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (i InputChatThemeGift) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, i.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (i InputChatThemeGift) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (i InputChatThemeGift) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, i.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (i InputChatThemeGift) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, i.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (i InputChatThemeGift) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, i.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (i InputChatThemeGift) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, i.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (i InputChatThemeGift) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, i.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (i InputChatThemeGift) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(i.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (i InputChatThemeGift) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, i.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (i InputChatThemeGift) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, i.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (i InputChatThemeGift) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, i.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (i InputChatThemeGift) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, i.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (i InputChatThemeGift) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(i.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (i InputChatThemeGift) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(i.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (i InputChatThemeGift) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(i.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (i InputChatThemeGift) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(i.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (i InputChatThemeGift) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, i.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (i InputChatThemeGift) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(i.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (i InputChatThemeGift) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(i.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (i InputChatThemeGift) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, i.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (i InputChatThemeGift) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(i.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (i InputChatThemeGift) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, i.Name)
}

// SetOption is a helper method for Client.SetOption
func (i InputChatThemeGift) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(i.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (i InputChatThemeGift) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, i.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (i InputChatThemeGift) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, i.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (i InputChatThemeGift) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(i.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (i InputChatThemeGift) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, i.Name)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputChecklist) AddChecklistTasks(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.AddChecklistTasks(chatId, messageId, i.Tasks)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (i InputChecklistTask) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(i.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (i InputChecklistTask) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(i.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (i InputChecklistTask) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, i.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (i InputChecklistTask) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(i.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (i InputChecklistTask) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, i.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (i InputChecklistTask) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(i.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (i InputChecklistTask) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, i.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (i InputChecklistTask) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, i.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (i InputChecklistTask) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, i.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (i InputChecklistTask) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(i.Text, toLanguageCode)
}

// AnswerCustomQuery is a helper method for Client.AnswerCustomQuery
func (i InputCredentialsApplePay) AnswerCustomQuery(client *Client, customQueryId string) (*Ok, error) {
	return client.AnswerCustomQuery(customQueryId, i.Data)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (i InputCredentialsApplePay) SendWebAppData(client *Client, botUserId int64, buttonText string) (*Ok, error) {
	return client.SendWebAppData(botUserId, buttonText, i.Data)
}

// AnswerCustomQuery is a helper method for Client.AnswerCustomQuery
func (i InputCredentialsGooglePay) AnswerCustomQuery(client *Client, customQueryId string) (*Ok, error) {
	return client.AnswerCustomQuery(customQueryId, i.Data)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (i InputCredentialsGooglePay) SendWebAppData(client *Client, botUserId int64, buttonText string) (*Ok, error) {
	return client.SendWebAppData(botUserId, buttonText, i.Data)
}

// AnswerCustomQuery is a helper method for Client.AnswerCustomQuery
func (i InputCredentialsNew) AnswerCustomQuery(client *Client, customQueryId string) (*Ok, error) {
	return client.AnswerCustomQuery(customQueryId, i.Data)
}

// SendWebAppData is a helper method for Client.SendWebAppData
func (i InputCredentialsNew) SendWebAppData(client *Client, botUserId int64, buttonText string) (*Ok, error) {
	return client.SendWebAppData(botUserId, buttonText, i.Data)
}

// ValidateOrderInfo is a helper method for Client.ValidateOrderInfo
func (i InputCredentialsNew) ValidateOrderInfo(client *Client, inputInvoice *InputInvoice, opts *ValidateOrderInfoOpts) (*ValidatedOrderInfo, error) {
	return client.ValidateOrderInfo(inputInvoice, i.AllowSave, opts)
}

// SetFileGenerationProgress is a helper method for Client.SetFileGenerationProgress
func (i InputFileGenerated) SetFileGenerationProgress(client *Client, generationId string, localPrefixSize int64) (*Ok, error) {
	return client.SetFileGenerationProgress(generationId, i.ExpectedSize, localPrefixSize)
}

// ConfirmQrCodeAuthentication is a helper method for Client.ConfirmQrCodeAuthentication
func (i InputGroupCallLink) ConfirmQrCodeAuthentication(client *Client) (*Session, error) {
	return client.ConfirmQrCodeAuthentication(i.Link)
}

// DeleteBusinessChatLink is a helper method for Client.DeleteBusinessChatLink
func (i InputGroupCallLink) DeleteBusinessChatLink(client *Client) (*Ok, error) {
	return client.DeleteBusinessChatLink(i.Link)
}

// EditBusinessChatLink is a helper method for Client.EditBusinessChatLink
func (i InputGroupCallLink) EditBusinessChatLink(client *Client, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	return client.EditBusinessChatLink(i.Link, linkInfo)
}

// GetDeepLinkInfo is a helper method for Client.GetDeepLinkInfo
func (i InputGroupCallLink) GetDeepLinkInfo(client *Client) (*DeepLinkInfo, error) {
	return client.GetDeepLinkInfo(i.Link)
}

// GetExternalLink is a helper method for Client.GetExternalLink
func (i InputGroupCallLink) GetExternalLink(client *Client, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetExternalLink(i.Link, allowWriteAccess)
}

// GetExternalLinkInfo is a helper method for Client.GetExternalLinkInfo
func (i InputGroupCallLink) GetExternalLinkInfo(client *Client) (*LoginUrlInfo, error) {
	return client.GetExternalLinkInfo(i.Link)
}

// GetInternalLinkType is a helper method for Client.GetInternalLinkType
func (i InputGroupCallLink) GetInternalLinkType(client *Client) (*InternalLinkType, error) {
	return client.GetInternalLinkType(i.Link)
}

// AddChatMember is a helper method for Client.AddChatMember
func (i InputGroupCallMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(i.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (i InputGroupCallMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(i.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (i InputGroupCallMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(i.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (i InputGroupCallMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(i.ChatId, i.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (i InputGroupCallMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, i.ChatId, i.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (i InputGroupCallMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(i.ChatId, senderId, disableNotification, inputMessageContent, opts)
}
