package gotdbot

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeChannelPaidReactionSend) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, s.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeChannelPaidReactionSend) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeChannelPaidReactionSend) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeChannelPaidReactionSend) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeChannelPaidReactionSend) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeChannelPaidReactionSend) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeChannelPaidReactionSend) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, s.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeChannelPaidReactionSend) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeChannelPaidReactionSend) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeChannelPaidReactionSend) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, s.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeChannelPaidReactionSend) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeChannelPaidReactionSend) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, s.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeChannelPaidReactionSend) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelPaidReactionSend) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeChannelPaidReactionSend) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (s StarTransactionTypeChannelPaidReactionSend) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(s.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeChannelPaidReactionSend) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, s.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeChannelPaidReactionSend) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeChannelPaidReactionSend) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeChannelPaidReactionSend) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeChannelPaidReactionSend) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeChannelPaidReactionSend) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeChannelPaidReactionSend) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, s.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeChannelPaidReactionSend) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, s.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeChannelPaidReactionSend) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeChannelPaidReactionSend) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeChannelPaidReactionSend) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, s.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeChannelPaidReactionSend) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, s.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeChannelPaidReactionSend) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeChannelPaidReactionSend) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeChannelPaidReactionSend) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeChannelPaidReactionSend) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeChannelPaidReactionSend) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeChannelPaidReactionSend) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeChannelPaidReactionSend) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeChannelPaidReactionSend) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeChannelPaidReactionSend) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeChannelPaidReactionSend) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, s.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeChannelPaidReactionSend) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, s.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (s StarTransactionTypeChannelPaidReactionSend) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, s.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeChannelPaidReactionSend) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeChannelPaidReactionSend) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeChannelPaidReactionSend) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeChannelPaidReactionSend) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeChannelPaidReactionSend) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeChannelPaidReactionSend) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeChannelPaidReactionSend) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeChannelPaidReactionSend) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeChannelPaidReactionSend) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeChannelPaidReactionSend) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeChannelPaidReactionSend) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeChannelPaidReactionSend) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeChannelPaidReactionSend) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeChannelPaidReactionSend) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeChannelPaidReactionSend) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeChannelPaidReactionSend) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeChannelPaidReactionSend) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, s.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeChannelPaidReactionSend) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeChannelPaidReactionSend) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeChannelPaidReactionSend) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeChannelPaidReactionSend) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeChannelPaidReactionSend) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeChannelPaidReactionSend) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeChannelPaidReactionSend) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeChannelPaidReactionSend) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeChannelPaidReactionSend) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeChannelPaidReactionSend) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeChannelPaidReactionSend) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeChannelPaidReactionSend) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeChannelPaidReactionSend) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeChannelPaidReactionSend) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeChannelPaidReactionSend) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeChannelPaidReactionSend) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeChannelPaidReactionSend) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeChannelPaidReactionSend) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeChannelPaidReactionSend) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeChannelPaidReactionSend) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeChannelPaidReactionSend) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeChannelPaidReactionSend) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeChannelPaidReactionSend) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeChannelPaidReactionSend) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeChannelPaidReactionSend) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeChannelPaidReactionSend) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelPaidReactionSend) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, s.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeChannelPaidReactionSend) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, s.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeChannelPaidReactionSend) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, s.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeChannelPaidReactionSend) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, s.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeChannelPaidReactionSend) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeChannelPaidReactionSend) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeChannelPaidReactionSend) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, s.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeChannelPaidReactionSend) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeChannelPaidReactionSend) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeChannelPaidReactionSend) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, s.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeChannelPaidReactionSend) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, s.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeChannelPaidReactionSend) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeChannelPaidReactionSend) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, s.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeChannelPaidReactionSend) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, s.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeChannelPaidReactionSend) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, s.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeChannelPaidReactionSend) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeChannelPaidReactionSend) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeChannelPaidReactionSend) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeChannelPaidReactionSend) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeChannelPaidReactionSend) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeChannelPaidReactionSend) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelPaidReactionSend) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeChannelPaidReactionSend) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, s.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeChannelPaidReactionSend) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeChannelPaidReactionSend) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeChannelPaidReactionSend) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeChannelPaidReactionSend) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, s.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeChannelPaidReactionSend) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeChannelPaidReactionSend) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelSubscriptionPurchase) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeChannelSubscriptionPurchase) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeChannelSubscriptionPurchase) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeChannelSubscriptionPurchase) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeChannelSubscriptionPurchase) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeChannelSubscriptionPurchase) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeChannelSubscriptionPurchase) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeChannelSubscriptionPurchase) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeChannelSubscriptionPurchase) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeChannelSubscriptionPurchase) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeChannelSubscriptionPurchase) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeChannelSubscriptionPurchase) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeChannelSubscriptionPurchase) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeChannelSubscriptionPurchase) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeChannelSubscriptionPurchase) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeChannelSubscriptionPurchase) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeChannelSubscriptionPurchase) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeChannelSubscriptionPurchase) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeChannelSubscriptionPurchase) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeChannelSubscriptionPurchase) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeChannelSubscriptionPurchase) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeChannelSubscriptionPurchase) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeChannelSubscriptionPurchase) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeChannelSubscriptionPurchase) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeChannelSubscriptionPurchase) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeChannelSubscriptionPurchase) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeChannelSubscriptionPurchase) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeChannelSubscriptionPurchase) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeChannelSubscriptionPurchase) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeChannelSubscriptionPurchase) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeChannelSubscriptionPurchase) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeChannelSubscriptionPurchase) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeChannelSubscriptionPurchase) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelSubscriptionPurchase) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeChannelSubscriptionPurchase) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeChannelSubscriptionPurchase) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeChannelSubscriptionPurchase) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeChannelSubscriptionPurchase) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeChannelSubscriptionPurchase) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeChannelSubscriptionPurchase) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeChannelSubscriptionPurchase) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeChannelSubscriptionPurchase) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeChannelSubscriptionPurchase) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeChannelSubscriptionPurchase) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelSubscriptionPurchase) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeChannelSubscriptionPurchase) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeChannelSubscriptionPurchase) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeChannelSubscriptionPurchase) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeChannelSubscriptionPurchase) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeChannelSubscriptionPurchase) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeChannelSubscriptionPurchase) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeChannelSubscriptionPurchase) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeChannelSubscriptionPurchase) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeChannelSubscriptionPurchase) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeChannelSubscriptionPurchase) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeChannelSubscriptionPurchase) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelSubscriptionPurchase) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeChannelSubscriptionPurchase) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeChannelSubscriptionPurchase) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeChannelSubscriptionPurchase) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeChannelSubscriptionPurchase) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeChannelSubscriptionPurchase) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeChannelSubscriptionPurchase) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeChannelSubscriptionPurchase) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeChannelSubscriptionPurchase) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeChannelSubscriptionPurchase) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeChannelSubscriptionPurchase) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeChannelSubscriptionPurchase) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeChannelSubscriptionPurchase) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeChannelSubscriptionPurchase) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeChannelSubscriptionPurchase) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeChannelSubscriptionPurchase) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeChannelSubscriptionPurchase) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeChannelSubscriptionPurchase) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeChannelSubscriptionPurchase) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeChannelSubscriptionPurchase) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeChannelSubscriptionPurchase) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelSubscriptionPurchase) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeChannelSubscriptionPurchase) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeChannelSubscriptionPurchase) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeChannelSubscriptionPurchase) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeChannelSubscriptionPurchase) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeChannelSubscriptionPurchase) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeChannelSubscriptionPurchase) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeChannelSubscriptionPurchase) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeChannelSubscriptionPurchase) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeChannelSubscriptionPurchase) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeChannelSubscriptionPurchase) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeChannelSubscriptionPurchase) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeChannelSubscriptionPurchase) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeChannelSubscriptionPurchase) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeChannelSubscriptionPurchase) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelSubscriptionPurchase) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeChannelSubscriptionPurchase) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeChannelSubscriptionPurchase) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeChannelSubscriptionPurchase) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeChannelSubscriptionPurchase) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeChannelSubscriptionSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeChannelSubscriptionSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeChannelSubscriptionSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeChannelSubscriptionSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeChannelSubscriptionSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeChannelSubscriptionSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeChannelSubscriptionSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeChannelSubscriptionSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeChannelSubscriptionSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeChannelSubscriptionSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeChannelSubscriptionSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeChannelSubscriptionSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeChannelSubscriptionSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeChannelSubscriptionSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeChannelSubscriptionSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeChannelSubscriptionSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeChannelSubscriptionSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeChannelSubscriptionSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeChannelSubscriptionSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeChannelSubscriptionSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeChannelSubscriptionSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeChannelSubscriptionSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeChannelSubscriptionSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeChannelSubscriptionSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeChannelSubscriptionSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeChannelSubscriptionSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeChannelSubscriptionSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeChannelSubscriptionSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeChannelSubscriptionSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeChannelSubscriptionSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeChannelSubscriptionSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeChannelSubscriptionSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeChannelSubscriptionSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeChannelSubscriptionSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeChannelSubscriptionSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeChannelSubscriptionSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeChannelSubscriptionSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeChannelSubscriptionSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeChannelSubscriptionSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeChannelSubscriptionSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeChannelSubscriptionSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeChannelSubscriptionSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeChannelSubscriptionSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (s StarTransactionTypeGiftAuctionBid) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarTransactionTypeGiftAuctionBid) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(s.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StarTransactionTypeGiftAuctionBid) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(s.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (s StarTransactionTypeGiftAuctionBid) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(s.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (s StarTransactionTypeGiftAuctionBid) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(s.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (s StarTransactionTypeGiftAuctionBid) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, s.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (s StarTransactionTypeGiftAuctionBid) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(s.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StarTransactionTypeGiftAuctionBid) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(s.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (s StarTransactionTypeGiftAuctionBid) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(s.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarTransactionTypeGiftAuctionBid) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(s.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (s StarTransactionTypeGiftAuctionBid) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (s StarTransactionTypeGiftAuctionBid) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (s StarTransactionTypeGiftAuctionBid) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(s.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (s StarTransactionTypeGiftAuctionBid) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, s.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StarTransactionTypeGiftAuctionBid) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(s.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StarTransactionTypeGiftAuctionBid) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, s.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StarTransactionTypeGiftAuctionBid) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(s.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (s StarTransactionTypeGiftAuctionBid) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(s.OwnerId, receivedGiftIds)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (s StarTransactionTypeGiftOriginalDetailsDrop) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarTransactionTypeGiftOriginalDetailsDrop) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(s.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StarTransactionTypeGiftOriginalDetailsDrop) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(s.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (s StarTransactionTypeGiftOriginalDetailsDrop) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(s.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(s.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, s.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(s.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(s.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(s.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarTransactionTypeGiftOriginalDetailsDrop) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(s.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (s StarTransactionTypeGiftOriginalDetailsDrop) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (s StarTransactionTypeGiftOriginalDetailsDrop) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (s StarTransactionTypeGiftOriginalDetailsDrop) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(s.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (s StarTransactionTypeGiftOriginalDetailsDrop) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, s.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StarTransactionTypeGiftOriginalDetailsDrop) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(s.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StarTransactionTypeGiftOriginalDetailsDrop) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, s.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StarTransactionTypeGiftOriginalDetailsDrop) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(s.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (s StarTransactionTypeGiftOriginalDetailsDrop) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(s.OwnerId, receivedGiftIds)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (s StarTransactionTypeGiftPurchase) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarTransactionTypeGiftPurchase) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(s.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StarTransactionTypeGiftPurchase) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(s.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (s StarTransactionTypeGiftPurchase) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(s.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (s StarTransactionTypeGiftPurchase) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(s.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (s StarTransactionTypeGiftPurchase) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, s.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (s StarTransactionTypeGiftPurchase) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(s.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StarTransactionTypeGiftPurchase) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(s.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (s StarTransactionTypeGiftPurchase) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(s.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarTransactionTypeGiftPurchase) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(s.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (s StarTransactionTypeGiftPurchase) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (s StarTransactionTypeGiftPurchase) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (s StarTransactionTypeGiftPurchase) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(s.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (s StarTransactionTypeGiftPurchase) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, s.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StarTransactionTypeGiftPurchase) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(s.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StarTransactionTypeGiftPurchase) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, s.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StarTransactionTypeGiftPurchase) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(s.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (s StarTransactionTypeGiftPurchase) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(s.OwnerId, receivedGiftIds)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeGiftSale) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeGiftSale) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeGiftSale) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeGiftSale) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeGiftSale) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeGiftSale) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeGiftSale) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeGiftSale) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeGiftSale) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeGiftSale) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeGiftSale) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeGiftSale) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeGiftSale) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeGiftSale) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeGiftSale) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeGiftSale) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeGiftSale) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeGiftSale) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeGiftSale) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeGiftSale) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeGiftSale) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeGiftSale) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeGiftSale) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeGiftSale) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeGiftSale) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeGiftSale) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeGiftSale) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeGiftSale) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeGiftSale) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeGiftSale) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeGiftSale) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeGiftSale) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeGiftSale) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeGiftSale) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeGiftSale) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeGiftSale) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeGiftSale) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeGiftSale) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeGiftSale) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeGiftSale) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeGiftSale) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeGiftSale) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeGiftSale) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (s StarTransactionTypeGiftTransfer) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarTransactionTypeGiftTransfer) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(s.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StarTransactionTypeGiftTransfer) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(s.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (s StarTransactionTypeGiftTransfer) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(s.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (s StarTransactionTypeGiftTransfer) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(s.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (s StarTransactionTypeGiftTransfer) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, s.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (s StarTransactionTypeGiftTransfer) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(s.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StarTransactionTypeGiftTransfer) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(s.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (s StarTransactionTypeGiftTransfer) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(s.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarTransactionTypeGiftTransfer) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(s.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (s StarTransactionTypeGiftTransfer) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (s StarTransactionTypeGiftTransfer) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (s StarTransactionTypeGiftTransfer) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(s.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (s StarTransactionTypeGiftTransfer) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, s.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StarTransactionTypeGiftTransfer) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(s.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StarTransactionTypeGiftTransfer) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, s.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StarTransactionTypeGiftTransfer) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(s.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (s StarTransactionTypeGiftTransfer) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(s.OwnerId, receivedGiftIds)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeGiftUpgrade) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, s.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (s StarTransactionTypeGiftUpgrade) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(s.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (s StarTransactionTypeGiftUpgrade) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(s.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (s StarTransactionTypeGiftUpgrade) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(s.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (s StarTransactionTypeGiftUpgrade) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(s.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (s StarTransactionTypeGiftUpgrade) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(s.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (s StarTransactionTypeGiftUpgrade) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(s.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (s StarTransactionTypeGiftUpgrade) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(s.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (s StarTransactionTypeGiftUpgrade) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(s.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (s StarTransactionTypeGiftUpgrade) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(s.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeGiftUpgrade) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, s.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (s StarTransactionTypeGiftUpgrade) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(s.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (s StarTransactionTypeGiftUpgrade) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, s.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (s StarTransactionTypeGiftUpgrade) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(s.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (s StarTransactionTypeGiftUpgrade) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(s.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (s StarTransactionTypeGiftUpgrade) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(s.UserId)
}

// GetUser is a helper method for Client.GetUser
func (s StarTransactionTypeGiftUpgrade) GetUser(client *Client) (*User, error) {
	return client.GetUser(s.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeGiftUpgrade) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, s.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (s StarTransactionTypeGiftUpgrade) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(s.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (s StarTransactionTypeGiftUpgrade) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(s.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (s StarTransactionTypeGiftUpgrade) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(s.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (s StarTransactionTypeGiftUpgrade) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(s.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (s StarTransactionTypeGiftUpgrade) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(s.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (s StarTransactionTypeGiftUpgrade) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, s.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (s StarTransactionTypeGiftUpgrade) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, s.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeGiftUpgrade) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, s.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (s StarTransactionTypeGiftUpgrade) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(s.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (s StarTransactionTypeGiftUpgrade) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(s.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (s StarTransactionTypeGiftUpgrade) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(s.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeGiftUpgrade) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, s.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (s StarTransactionTypeGiftUpgrade) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, s.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (s StarTransactionTypeGiftUpgrade) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(s.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (s StarTransactionTypeGiftUpgrade) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(s.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (s StarTransactionTypeGiftUpgrade) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(s.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (s StarTransactionTypeGiftUpgrade) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(s.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (s StarTransactionTypeGiftUpgrade) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(s.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (s StarTransactionTypeGiftUpgrade) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(s.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (s StarTransactionTypeGiftUpgrade) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(s.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (s StarTransactionTypeGiftUpgrade) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(s.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (s StarTransactionTypeGiftUpgrade) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(s.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (s StarTransactionTypeGiftUpgrade) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(s.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeGiftUpgrade) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, s.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (s StarTransactionTypeGiftUpgrade) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(s.UserId, stickerFormat, sticker)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (s StarTransactionTypeGiftUpgradePurchase) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (s StarTransactionTypeGiftUpgradePurchase) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(s.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (s StarTransactionTypeGiftUpgradePurchase) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(s.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (s StarTransactionTypeGiftUpgradePurchase) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(s.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (s StarTransactionTypeGiftUpgradePurchase) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(s.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (s StarTransactionTypeGiftUpgradePurchase) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, s.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (s StarTransactionTypeGiftUpgradePurchase) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(s.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (s StarTransactionTypeGiftUpgradePurchase) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(s.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (s StarTransactionTypeGiftUpgradePurchase) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(s.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (s StarTransactionTypeGiftUpgradePurchase) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(s.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (s StarTransactionTypeGiftUpgradePurchase) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (s StarTransactionTypeGiftUpgradePurchase) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(s.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (s StarTransactionTypeGiftUpgradePurchase) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(s.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (s StarTransactionTypeGiftUpgradePurchase) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, s.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (s StarTransactionTypeGiftUpgradePurchase) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(s.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (s StarTransactionTypeGiftUpgradePurchase) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, s.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (s StarTransactionTypeGiftUpgradePurchase) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(s.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (s StarTransactionTypeGiftUpgradePurchase) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(s.OwnerId, receivedGiftIds)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypeGiveawayDeposit) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypeGiveawayDeposit) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypeGiveawayDeposit) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypeGiveawayDeposit) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypeGiveawayDeposit) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypeGiveawayDeposit) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypeGiveawayDeposit) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypeGiveawayDeposit) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypeGiveawayDeposit) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypeGiveawayDeposit) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypeGiveawayDeposit) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypeGiveawayDeposit) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypeGiveawayDeposit) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypeGiveawayDeposit) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypeGiveawayDeposit) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypeGiveawayDeposit) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypeGiveawayDeposit) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypeGiveawayDeposit) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypeGiveawayDeposit) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypeGiveawayDeposit) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypeGiveawayDeposit) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypeGiveawayDeposit) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypeGiveawayDeposit) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypeGiveawayDeposit) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypeGiveawayDeposit) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypeGiveawayDeposit) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypeGiveawayDeposit) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypeGiveawayDeposit) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypeGiveawayDeposit) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypeGiveawayDeposit) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypeGiveawayDeposit) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypeGiveawayDeposit) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypeGiveawayDeposit) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypeGiveawayDeposit) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypeGiveawayDeposit) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypeGiveawayDeposit) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypeGiveawayDeposit) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypeGiveawayDeposit) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypeGiveawayDeposit) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypeGiveawayDeposit) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypeGiveawayDeposit) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypeGiveawayDeposit) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypeGiveawayDeposit) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypeGiveawayDeposit) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypeGiveawayDeposit) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypeGiveawayDeposit) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypeGiveawayDeposit) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypeGiveawayDeposit) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypeGiveawayDeposit) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypeGiveawayDeposit) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypeGiveawayDeposit) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypeGiveawayDeposit) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypeGiveawayDeposit) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypeGiveawayDeposit) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypeGiveawayDeposit) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypeGiveawayDeposit) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypeGiveawayDeposit) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypeGiveawayDeposit) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypeGiveawayDeposit) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypeGiveawayDeposit) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypeGiveawayDeposit) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypeGiveawayDeposit) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypeGiveawayDeposit) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypeGiveawayDeposit) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypeGiveawayDeposit) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypeGiveawayDeposit) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypeGiveawayDeposit) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypeGiveawayDeposit) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypeGiveawayDeposit) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypeGiveawayDeposit) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypeGiveawayDeposit) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypeGiveawayDeposit) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypeGiveawayDeposit) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypeGiveawayDeposit) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypeGiveawayDeposit) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypeGiveawayDeposit) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypeGiveawayDeposit) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypeGiveawayDeposit) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypeGiveawayDeposit) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypeGiveawayDeposit) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypeGiveawayDeposit) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypeGiveawayDeposit) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypeGiveawayDeposit) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypeGiveawayDeposit) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypeGiveawayDeposit) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypeGiveawayDeposit) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypeGiveawayDeposit) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypeGiveawayDeposit) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypeGiveawayDeposit) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypeGiveawayDeposit) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypeGiveawayDeposit) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypeGiveawayDeposit) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypeGiveawayDeposit) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypeGiveawayDeposit) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypeGiveawayDeposit) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypeGiveawayDeposit) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypeGiveawayDeposit) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypeGiveawayDeposit) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypeGiveawayDeposit) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypeGiveawayDeposit) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypeGiveawayDeposit) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypeGiveawayDeposit) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypeGiveawayDeposit) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypeGiveawayDeposit) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypeGiveawayDeposit) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypeGiveawayDeposit) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypeGiveawayDeposit) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypeGiveawayDeposit) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypeGiveawayDeposit) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypeGiveawayDeposit) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypeGiveawayDeposit) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypeGiveawayDeposit) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypeGiveawayDeposit) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypeGiveawayDeposit) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypeGiveawayDeposit) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypeGiveawayDeposit) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypeGiveawayDeposit) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypeGiveawayDeposit) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypeGiveawayDeposit) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypeGiveawayDeposit) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypeGiveawayDeposit) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypeGiveawayDeposit) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypeGiveawayDeposit) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypeGiveawayDeposit) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypeGiveawayDeposit) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypeGiveawayDeposit) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypeGiveawayDeposit) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypeGiveawayDeposit) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypeGiveawayDeposit) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypeGiveawayDeposit) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypeGiveawayDeposit) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypeGiveawayDeposit) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypeGiveawayDeposit) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypeGiveawayDeposit) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypeGiveawayDeposit) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypeGiveawayDeposit) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypeGiveawayDeposit) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypeGiveawayDeposit) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypeGiveawayDeposit) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypeGiveawayDeposit) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypeGiveawayDeposit) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypeGiveawayDeposit) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypeGiveawayDeposit) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypeGiveawayDeposit) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypeGiveawayDeposit) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypeGiveawayDeposit) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypeGiveawayDeposit) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypeGiveawayDeposit) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypeGiveawayDeposit) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypeGiveawayDeposit) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypeGiveawayDeposit) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypeGiveawayDeposit) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypeGiveawayDeposit) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypeGiveawayDeposit) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypeGiveawayDeposit) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypeGiveawayDeposit) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypeGiveawayDeposit) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypeGiveawayDeposit) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypeGiveawayDeposit) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypeGiveawayDeposit) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypeGiveawayDeposit) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypeGiveawayDeposit) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypeGiveawayDeposit) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypeGiveawayDeposit) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypeGiveawayDeposit) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypeGiveawayDeposit) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypeGiveawayDeposit) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypeGiveawayDeposit) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypeGiveawayDeposit) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypeGiveawayDeposit) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypeGiveawayDeposit) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypeGiveawayDeposit) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypeGiveawayDeposit) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypeGiveawayDeposit) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypeGiveawayDeposit) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypeGiveawayDeposit) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypeGiveawayDeposit) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypeGiveawayDeposit) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypeGiveawayDeposit) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypeGiveawayDeposit) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypeGiveawayDeposit) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypeGiveawayDeposit) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypeGiveawayDeposit) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypeGiveawayDeposit) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypeGiveawayDeposit) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypeGiveawayDeposit) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypeGiveawayDeposit) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypeGiveawayDeposit) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypeGiveawayDeposit) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypeGiveawayDeposit) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypeGiveawayDeposit) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypeGiveawayDeposit) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypeGiveawayDeposit) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypeGiveawayDeposit) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypeGiveawayDeposit) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypeGiveawayDeposit) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypeGiveawayDeposit) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypeGiveawayDeposit) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypeGiveawayDeposit) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypeGiveawayDeposit) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypeGiveawayDeposit) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypeGiveawayDeposit) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypeGiveawayDeposit) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypeGiveawayDeposit) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypeGiveawayDeposit) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypeGiveawayDeposit) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypeGiveawayDeposit) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypeGiveawayDeposit) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypeGiveawayDeposit) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypeGiveawayDeposit) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypeGiveawayDeposit) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypeGiveawayDeposit) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypeGiveawayDeposit) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypeGiveawayDeposit) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypeGiveawayDeposit) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypeGiveawayDeposit) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypeGiveawayDeposit) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypeGiveawayDeposit) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypeGiveawayDeposit) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypeGiveawayDeposit) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypeGiveawayDeposit) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypeGiveawayDeposit) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypeGiveawayDeposit) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypeGiveawayDeposit) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypeGiveawayDeposit) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypeGiveawayDeposit) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypeGiveawayDeposit) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypeGiveawayDeposit) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypeGiveawayDeposit) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypeGiveawayDeposit) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypeGiveawayDeposit) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypeGiveawayDeposit) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypeGiveawayDeposit) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypeGiveawayDeposit) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypeGiveawayDeposit) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypeGiveawayDeposit) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypeGiveawayDeposit) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypeGiveawayDeposit) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypeGiveawayDeposit) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypeGiveawayDeposit) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypeGiveawayDeposit) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypeGiveawayDeposit) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidGroupCallMessageReceive) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, s.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidGroupCallMessageReceive) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, s.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (s StarTransactionTypePaidGroupCallMessageReceive) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, s.SenderId, reportSpam)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidGroupCallMessageReceive) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, s.SenderId)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (s StarTransactionTypePaidGroupCallMessageReceive) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(s.SenderId, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypePaidGroupCallMessageSend) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypePaidGroupCallMessageSend) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypePaidGroupCallMessageSend) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypePaidGroupCallMessageSend) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypePaidGroupCallMessageSend) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidGroupCallMessageSend) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypePaidGroupCallMessageSend) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypePaidGroupCallMessageSend) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypePaidGroupCallMessageSend) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypePaidGroupCallMessageSend) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypePaidGroupCallMessageSend) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypePaidGroupCallMessageSend) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypePaidGroupCallMessageSend) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypePaidGroupCallMessageSend) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypePaidGroupCallMessageSend) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypePaidGroupCallMessageSend) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypePaidGroupCallMessageSend) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypePaidGroupCallMessageSend) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypePaidGroupCallMessageSend) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypePaidGroupCallMessageSend) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypePaidGroupCallMessageSend) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypePaidGroupCallMessageSend) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypePaidGroupCallMessageSend) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypePaidGroupCallMessageSend) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypePaidGroupCallMessageSend) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypePaidGroupCallMessageSend) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypePaidGroupCallMessageSend) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypePaidGroupCallMessageSend) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypePaidGroupCallMessageSend) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypePaidGroupCallMessageSend) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypePaidGroupCallMessageSend) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypePaidGroupCallMessageSend) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypePaidGroupCallMessageSend) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypePaidGroupCallMessageSend) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypePaidGroupCallMessageSend) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidGroupCallMessageSend) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypePaidGroupCallMessageSend) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypePaidGroupCallMessageSend) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypePaidGroupCallMessageSend) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypePaidGroupCallMessageSend) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypePaidGroupCallMessageSend) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypePaidGroupCallMessageSend) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypePaidGroupCallMessageSend) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypePaidGroupCallMessageSend) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypePaidGroupCallMessageSend) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypePaidGroupCallMessageSend) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypePaidGroupCallMessageSend) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypePaidGroupCallMessageSend) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypePaidGroupCallMessageSend) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypePaidGroupCallMessageSend) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypePaidGroupCallMessageSend) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypePaidGroupCallMessageSend) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypePaidGroupCallMessageSend) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypePaidGroupCallMessageSend) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypePaidGroupCallMessageSend) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypePaidGroupCallMessageSend) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypePaidGroupCallMessageSend) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypePaidGroupCallMessageSend) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypePaidGroupCallMessageSend) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypePaidGroupCallMessageSend) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypePaidGroupCallMessageSend) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypePaidGroupCallMessageSend) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypePaidGroupCallMessageSend) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypePaidGroupCallMessageSend) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypePaidGroupCallMessageSend) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypePaidGroupCallMessageSend) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypePaidGroupCallMessageSend) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypePaidGroupCallMessageSend) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypePaidGroupCallMessageSend) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypePaidGroupCallMessageSend) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypePaidGroupCallMessageSend) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypePaidGroupCallMessageSend) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypePaidGroupCallMessageSend) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypePaidGroupCallMessageSend) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypePaidGroupCallMessageSend) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypePaidGroupCallMessageSend) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypePaidGroupCallMessageSend) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypePaidGroupCallMessageSend) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypePaidGroupCallMessageSend) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypePaidGroupCallMessageSend) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypePaidGroupCallMessageSend) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypePaidGroupCallMessageSend) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypePaidGroupCallMessageSend) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypePaidGroupCallMessageSend) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypePaidGroupCallMessageSend) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypePaidGroupCallMessageSend) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypePaidGroupCallMessageSend) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypePaidGroupCallMessageSend) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypePaidGroupCallMessageSend) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypePaidGroupCallMessageSend) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypePaidGroupCallMessageSend) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypePaidGroupCallMessageSend) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypePaidGroupCallMessageSend) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidGroupCallMessageSend) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypePaidGroupCallMessageSend) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypePaidGroupCallMessageSend) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypePaidGroupCallMessageSend) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypePaidGroupCallMessageSend) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypePaidGroupCallMessageSend) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypePaidGroupCallMessageSend) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypePaidGroupCallMessageSend) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypePaidGroupCallMessageSend) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypePaidGroupCallMessageSend) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypePaidGroupCallMessageSend) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypePaidGroupCallMessageSend) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypePaidGroupCallMessageSend) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypePaidGroupCallMessageSend) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypePaidGroupCallMessageSend) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypePaidGroupCallMessageSend) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypePaidGroupCallMessageSend) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypePaidGroupCallMessageSend) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypePaidGroupCallMessageSend) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypePaidGroupCallMessageSend) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypePaidGroupCallMessageSend) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypePaidGroupCallMessageSend) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypePaidGroupCallMessageSend) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypePaidGroupCallMessageSend) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypePaidGroupCallMessageSend) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypePaidGroupCallMessageSend) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypePaidGroupCallMessageSend) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypePaidGroupCallMessageSend) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypePaidGroupCallMessageSend) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypePaidGroupCallMessageSend) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypePaidGroupCallMessageSend) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypePaidGroupCallMessageSend) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypePaidGroupCallMessageSend) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypePaidGroupCallMessageSend) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypePaidGroupCallMessageSend) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypePaidGroupCallMessageSend) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypePaidGroupCallMessageSend) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypePaidGroupCallMessageSend) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypePaidGroupCallMessageSend) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypePaidGroupCallMessageSend) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypePaidGroupCallMessageSend) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidGroupCallReactionReceive) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, s.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidGroupCallReactionReceive) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, s.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (s StarTransactionTypePaidGroupCallReactionReceive) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, s.SenderId, reportSpam)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidGroupCallReactionReceive) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, s.SenderId)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (s StarTransactionTypePaidGroupCallReactionReceive) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(s.SenderId, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypePaidGroupCallReactionSend) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypePaidGroupCallReactionSend) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypePaidGroupCallReactionSend) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypePaidGroupCallReactionSend) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypePaidGroupCallReactionSend) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidGroupCallReactionSend) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypePaidGroupCallReactionSend) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypePaidGroupCallReactionSend) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypePaidGroupCallReactionSend) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypePaidGroupCallReactionSend) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypePaidGroupCallReactionSend) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypePaidGroupCallReactionSend) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypePaidGroupCallReactionSend) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypePaidGroupCallReactionSend) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypePaidGroupCallReactionSend) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypePaidGroupCallReactionSend) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypePaidGroupCallReactionSend) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypePaidGroupCallReactionSend) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypePaidGroupCallReactionSend) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypePaidGroupCallReactionSend) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypePaidGroupCallReactionSend) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypePaidGroupCallReactionSend) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypePaidGroupCallReactionSend) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypePaidGroupCallReactionSend) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypePaidGroupCallReactionSend) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypePaidGroupCallReactionSend) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypePaidGroupCallReactionSend) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypePaidGroupCallReactionSend) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypePaidGroupCallReactionSend) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypePaidGroupCallReactionSend) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypePaidGroupCallReactionSend) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypePaidGroupCallReactionSend) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypePaidGroupCallReactionSend) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypePaidGroupCallReactionSend) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypePaidGroupCallReactionSend) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidGroupCallReactionSend) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypePaidGroupCallReactionSend) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypePaidGroupCallReactionSend) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypePaidGroupCallReactionSend) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypePaidGroupCallReactionSend) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypePaidGroupCallReactionSend) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypePaidGroupCallReactionSend) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypePaidGroupCallReactionSend) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypePaidGroupCallReactionSend) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypePaidGroupCallReactionSend) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypePaidGroupCallReactionSend) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypePaidGroupCallReactionSend) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypePaidGroupCallReactionSend) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypePaidGroupCallReactionSend) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypePaidGroupCallReactionSend) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypePaidGroupCallReactionSend) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypePaidGroupCallReactionSend) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypePaidGroupCallReactionSend) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypePaidGroupCallReactionSend) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypePaidGroupCallReactionSend) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypePaidGroupCallReactionSend) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypePaidGroupCallReactionSend) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypePaidGroupCallReactionSend) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypePaidGroupCallReactionSend) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypePaidGroupCallReactionSend) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypePaidGroupCallReactionSend) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypePaidGroupCallReactionSend) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypePaidGroupCallReactionSend) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypePaidGroupCallReactionSend) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypePaidGroupCallReactionSend) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (s StarTransactionTypePaidGroupCallReactionSend) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(s.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (s StarTransactionTypePaidGroupCallReactionSend) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(s.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (s StarTransactionTypePaidGroupCallReactionSend) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(s.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (s StarTransactionTypePaidGroupCallReactionSend) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(s.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (s StarTransactionTypePaidGroupCallReactionSend) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(s.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (s StarTransactionTypePaidGroupCallReactionSend) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(s.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (s StarTransactionTypePaidGroupCallReactionSend) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(s.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (s StarTransactionTypePaidGroupCallReactionSend) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(s.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (s StarTransactionTypePaidGroupCallReactionSend) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(s.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (s StarTransactionTypePaidGroupCallReactionSend) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(s.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (s StarTransactionTypePaidGroupCallReactionSend) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(s.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (s StarTransactionTypePaidGroupCallReactionSend) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(s.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (s StarTransactionTypePaidGroupCallReactionSend) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(s.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (s StarTransactionTypePaidGroupCallReactionSend) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(s.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (s StarTransactionTypePaidGroupCallReactionSend) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(s.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (s StarTransactionTypePaidGroupCallReactionSend) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(s.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (s StarTransactionTypePaidGroupCallReactionSend) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(s.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (s StarTransactionTypePaidGroupCallReactionSend) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, s.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (s StarTransactionTypePaidGroupCallReactionSend) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(s.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(s.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(s.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(s.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (s StarTransactionTypePaidGroupCallReactionSend) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(s.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(s.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (s StarTransactionTypePaidGroupCallReactionSend) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, s.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (s StarTransactionTypePaidGroupCallReactionSend) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(s.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (s StarTransactionTypePaidGroupCallReactionSend) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (s StarTransactionTypePaidGroupCallReactionSend) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(s.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(s.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (s StarTransactionTypePaidGroupCallReactionSend) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(s.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (s StarTransactionTypePaidGroupCallReactionSend) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(s.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (s StarTransactionTypePaidGroupCallReactionSend) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(s.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (s StarTransactionTypePaidGroupCallReactionSend) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(s.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidGroupCallReactionSend) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(s.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (s StarTransactionTypePaidGroupCallReactionSend) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(s.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (s StarTransactionTypePaidGroupCallReactionSend) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(s.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (s StarTransactionTypePaidGroupCallReactionSend) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, s.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (s StarTransactionTypePaidGroupCallReactionSend) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(s.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (s StarTransactionTypePaidGroupCallReactionSend) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(s.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (s StarTransactionTypePaidGroupCallReactionSend) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(s.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (s StarTransactionTypePaidGroupCallReactionSend) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(s.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, s.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (s StarTransactionTypePaidGroupCallReactionSend) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, s.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (s StarTransactionTypePaidGroupCallReactionSend) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(s.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(s.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(s.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (s StarTransactionTypePaidGroupCallReactionSend) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(s.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (s StarTransactionTypePaidGroupCallReactionSend) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(s.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (s StarTransactionTypePaidGroupCallReactionSend) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(s.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (s StarTransactionTypePaidGroupCallReactionSend) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, s.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(s.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(s.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(s.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(s.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(s.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(s.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(s.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(s.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(s.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(s.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(s.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(s.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(s.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(s.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(s.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(s.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(s.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(s.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(s.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(s.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(s.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(s.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(s.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (s StarTransactionTypePaidGroupCallReactionSend) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(s.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (s StarTransactionTypePaidGroupCallReactionSend) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(s.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (s StarTransactionTypePaidGroupCallReactionSend) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(s.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (s StarTransactionTypePaidGroupCallReactionSend) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(s.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (s StarTransactionTypePaidGroupCallReactionSend) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(s.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (s StarTransactionTypePaidGroupCallReactionSend) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(s.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (s StarTransactionTypePaidGroupCallReactionSend) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(s.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (s StarTransactionTypePaidGroupCallReactionSend) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(s.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (s StarTransactionTypePaidGroupCallReactionSend) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(s.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (s StarTransactionTypePaidGroupCallReactionSend) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(s.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (s StarTransactionTypePaidGroupCallReactionSend) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(s.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (s StarTransactionTypePaidGroupCallReactionSend) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(s.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (s StarTransactionTypePaidGroupCallReactionSend) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(s.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (s StarTransactionTypePaidGroupCallReactionSend) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(s.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (s StarTransactionTypePaidGroupCallReactionSend) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(s.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (s StarTransactionTypePaidGroupCallReactionSend) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, s.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (s StarTransactionTypePaidGroupCallReactionSend) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(s.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (s StarTransactionTypePaidGroupCallReactionSend) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(s.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(s.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(s.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(s.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(s.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(s.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, s.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(s.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(s.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(s.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(s.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(s.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (s StarTransactionTypePaidGroupCallReactionSend) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(s.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (s StarTransactionTypePaidGroupCallReactionSend) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(s.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (s StarTransactionTypePaidGroupCallReactionSend) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(s.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (s StarTransactionTypePaidGroupCallReactionSend) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(s.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (s StarTransactionTypePaidGroupCallReactionSend) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(s.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (s StarTransactionTypePaidGroupCallReactionSend) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(s.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (s StarTransactionTypePaidGroupCallReactionSend) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(s.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (s StarTransactionTypePaidGroupCallReactionSend) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(s.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (s StarTransactionTypePaidGroupCallReactionSend) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(s.ChatId, messageIds, forceRead, opts)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidMessageReceive) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, s.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidMessageReceive) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, s.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (s StarTransactionTypePaidMessageReceive) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, s.SenderId, reportSpam)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (s StarTransactionTypePaidMessageReceive) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, s.SenderId)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (s StarTransactionTypePaidMessageReceive) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(s.SenderId, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (s StarTransactionTypePaidMessageSend) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(s.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (s StarTransactionTypePaidMessageSend) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(s.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (s StarTransactionTypePaidMessageSend) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(s.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (s StarTransactionTypePaidMessageSend) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(s.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (s StarTransactionTypePaidMessageSend) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, s.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (s StarTransactionTypePaidMessageSend) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(s.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (s StarTransactionTypePaidMessageSend) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(s.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (s StarTransactionTypePaidMessageSend) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(s.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (s StarTransactionTypePaidMessageSend) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(s.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (s StarTransactionTypePaidMessageSend) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(s.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (s StarTransactionTypePaidMessageSend) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(s.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (s StarTransactionTypePaidMessageSend) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(s.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (s StarTransactionTypePaidMessageSend) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(s.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (s StarTransactionTypePaidMessageSend) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(s.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (s StarTransactionTypePaidMessageSend) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(s.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (s StarTransactionTypePaidMessageSend) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(s.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (s StarTransactionTypePaidMessageSend) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(s.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (s StarTransactionTypePaidMessageSend) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(s.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (s StarTransactionTypePaidMessageSend) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(s.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (s StarTransactionTypePaidMessageSend) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(s.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (s StarTransactionTypePaidMessageSend) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(s.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (s StarTransactionTypePaidMessageSend) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(s.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (s StarTransactionTypePaidMessageSend) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(s.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (s StarTransactionTypePaidMessageSend) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(s.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (s StarTransactionTypePaidMessageSend) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(s.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (s StarTransactionTypePaidMessageSend) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(s.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (s StarTransactionTypePaidMessageSend) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(s.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (s StarTransactionTypePaidMessageSend) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(s.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (s StarTransactionTypePaidMessageSend) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(s.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (s StarTransactionTypePaidMessageSend) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(s.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (s StarTransactionTypePaidMessageSend) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(s.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (s StarTransactionTypePaidMessageSend) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(s.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (s StarTransactionTypePaidMessageSend) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(s.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidMessageSend) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(s.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (s StarTransactionTypePaidMessageSend) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(s.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (s StarTransactionTypePaidMessageSend) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(s.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (s StarTransactionTypePaidMessageSend) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(s.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (s StarTransactionTypePaidMessageSend) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(s.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (s StarTransactionTypePaidMessageSend) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(s.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, s.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, s.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (s StarTransactionTypePaidMessageSend) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, s.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (s StarTransactionTypePaidMessageSend) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(s.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (s StarTransactionTypePaidMessageSend) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(s.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (s StarTransactionTypePaidMessageSend) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(s.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (s StarTransactionTypePaidMessageSend) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(s.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (s StarTransactionTypePaidMessageSend) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(s.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (s StarTransactionTypePaidMessageSend) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(s.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (s StarTransactionTypePaidMessageSend) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(s.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (s StarTransactionTypePaidMessageSend) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(s.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (s StarTransactionTypePaidMessageSend) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(s.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (s StarTransactionTypePaidMessageSend) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(s.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (s StarTransactionTypePaidMessageSend) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(s.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (s StarTransactionTypePaidMessageSend) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, s.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (s StarTransactionTypePaidMessageSend) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(s.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (s StarTransactionTypePaidMessageSend) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(s.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (s StarTransactionTypePaidMessageSend) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(s.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (s StarTransactionTypePaidMessageSend) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(s.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (s StarTransactionTypePaidMessageSend) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(s.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (s StarTransactionTypePaidMessageSend) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(s.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (s StarTransactionTypePaidMessageSend) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(s.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (s StarTransactionTypePaidMessageSend) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(s.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (s StarTransactionTypePaidMessageSend) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(s.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (s StarTransactionTypePaidMessageSend) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(s.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (s StarTransactionTypePaidMessageSend) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(s.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (s StarTransactionTypePaidMessageSend) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(s.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (s StarTransactionTypePaidMessageSend) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(s.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (s StarTransactionTypePaidMessageSend) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(s.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (s StarTransactionTypePaidMessageSend) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(s.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (s StarTransactionTypePaidMessageSend) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(s.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (s StarTransactionTypePaidMessageSend) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(s.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (s StarTransactionTypePaidMessageSend) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(s.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (s StarTransactionTypePaidMessageSend) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(s.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (s StarTransactionTypePaidMessageSend) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(s.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (s StarTransactionTypePaidMessageSend) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(s.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (s StarTransactionTypePaidMessageSend) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(s.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (s StarTransactionTypePaidMessageSend) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(s.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (s StarTransactionTypePaidMessageSend) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(s.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (s StarTransactionTypePaidMessageSend) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(s.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (s StarTransactionTypePaidMessageSend) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(s.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (s StarTransactionTypePaidMessageSend) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(s.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (s StarTransactionTypePaidMessageSend) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(s.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (s StarTransactionTypePaidMessageSend) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(s.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (s StarTransactionTypePaidMessageSend) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(s.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (s StarTransactionTypePaidMessageSend) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(s.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (s StarTransactionTypePaidMessageSend) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(s.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (s StarTransactionTypePaidMessageSend) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(s.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (s StarTransactionTypePaidMessageSend) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(s.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (s StarTransactionTypePaidMessageSend) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(s.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (s StarTransactionTypePaidMessageSend) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(s.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (s StarTransactionTypePaidMessageSend) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(s.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (s StarTransactionTypePaidMessageSend) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(s.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (s StarTransactionTypePaidMessageSend) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(s.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (s StarTransactionTypePaidMessageSend) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(s.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (s StarTransactionTypePaidMessageSend) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(s.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (s StarTransactionTypePaidMessageSend) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(s.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (s StarTransactionTypePaidMessageSend) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(s.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (s StarTransactionTypePaidMessageSend) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(s.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (s StarTransactionTypePaidMessageSend) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(s.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (s StarTransactionTypePaidMessageSend) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(s.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (s StarTransactionTypePaidMessageSend) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, s.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (s StarTransactionTypePaidMessageSend) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(s.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (s StarTransactionTypePaidMessageSend) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(s.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (s StarTransactionTypePaidMessageSend) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(s.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (s StarTransactionTypePaidMessageSend) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(s.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (s StarTransactionTypePaidMessageSend) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, s.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (s StarTransactionTypePaidMessageSend) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(s.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (s StarTransactionTypePaidMessageSend) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(s.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (s StarTransactionTypePaidMessageSend) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(s.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (s StarTransactionTypePaidMessageSend) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(s.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (s StarTransactionTypePaidMessageSend) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(s.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (s StarTransactionTypePaidMessageSend) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(s.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (s StarTransactionTypePaidMessageSend) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(s.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (s StarTransactionTypePaidMessageSend) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(s.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (s StarTransactionTypePaidMessageSend) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(s.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (s StarTransactionTypePaidMessageSend) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(s.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (s StarTransactionTypePaidMessageSend) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(s.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (s StarTransactionTypePaidMessageSend) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(s.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (s StarTransactionTypePaidMessageSend) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(s.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (s StarTransactionTypePaidMessageSend) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(s.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (s StarTransactionTypePaidMessageSend) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(s.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (s StarTransactionTypePaidMessageSend) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(s.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (s StarTransactionTypePaidMessageSend) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(s.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (s StarTransactionTypePaidMessageSend) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(s.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (s StarTransactionTypePaidMessageSend) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(s.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (s StarTransactionTypePaidMessageSend) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(s.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (s StarTransactionTypePaidMessageSend) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, s.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (s StarTransactionTypePaidMessageSend) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(s.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (s StarTransactionTypePaidMessageSend) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(s.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (s StarTransactionTypePaidMessageSend) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(s.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (s StarTransactionTypePaidMessageSend) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(s.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (s StarTransactionTypePaidMessageSend) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(s.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (s StarTransactionTypePaidMessageSend) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(s.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (s StarTransactionTypePaidMessageSend) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(s.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (s StarTransactionTypePaidMessageSend) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(s.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (s StarTransactionTypePaidMessageSend) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(s.ChatId)
}
