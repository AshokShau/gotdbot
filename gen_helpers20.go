package gotdbot

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatTitle) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatTitle) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatTitle) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatTitle) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatTitle) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatTitle) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatTitle) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatTitle) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatTitle) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatTitle) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatTitle) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatTitle) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatTitle) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatTitle) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatTitle) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatTitle) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatTitle) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatTitle) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatTitle) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatTitle) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatTitle) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatTitle) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatTitle) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatTitle) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatTitle) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatTitle) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatTitle) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatTitle) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatTitle) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatTitle) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatTitle) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatTitle) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatTitle) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatTitle) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatTitle) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatTitle) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (u UpdateChatTitle) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(u.Title)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatTitle) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatTitle) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatTitle) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatTitle) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatTitle) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatTitle) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatTitle) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatTitle) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatTitle) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatTitle) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatTitle) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatTitle) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatTitle) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatTitle) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatTitle) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatTitle) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatTitle) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatTitle) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatTitle) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatTitle) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatTitle) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatTitle) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatTitle) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatTitle) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatTitle) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatTitle) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatTitle) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatTitle) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatTitle) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatTitle) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatTitle) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatTitle) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatTitle) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatTitle) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatTitle) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatTitle) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatTitle) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatTitle) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatTitle) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatTitle) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatTitle) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatTitle) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatTitle) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatTitle) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatTitle) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatTitle) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatTitle) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatTitle) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatTitle) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatTitle) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatTitle) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatTitle) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatTitle) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatTitle) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatTitle) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatTitle) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatTitle) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatTitle) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatTitle) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatTitle) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatTitle) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatTitle) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatTitle) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatTitle) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatTitle) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatTitle) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatTitle) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatTitle) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatTitle) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatTitle) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatTitle) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatTitle) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatTitle) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatTitle) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatTitle) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatTitle) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatTitle) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatTitle) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatTitle) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatTitle) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatTitle) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatTitle) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatTitle) SetChatTitle(client *Client) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, u.Title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatTitle) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatTitle) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatTitle) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatTitle) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatTitle) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatTitle) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatTitle) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatTitle) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatTitle) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpdateChatTitle) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, u.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatTitle) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatTitle) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpdateChatTitle) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, u.Title)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatTitle) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatTitle) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpdateChatTitle) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, u.Title, recordVideo, usePortraitOrientation)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatTitle) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatTitle) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatTitle) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatTitle) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatTitle) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatTitle) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatTitle) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatTitle) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatTitle) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatTitle) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatTitle) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatTitle) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatTitle) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatTitle) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatTitle) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatTitle) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatTitle) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatTitle) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatTitle) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatTitle) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatTitle) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatTitle) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatTitle) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatTitle) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatUnreadMentionCount) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatUnreadMentionCount) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatUnreadMentionCount) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatUnreadMentionCount) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatUnreadMentionCount) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatUnreadMentionCount) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatUnreadMentionCount) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatUnreadMentionCount) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatUnreadMentionCount) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatUnreadMentionCount) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatUnreadMentionCount) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatUnreadMentionCount) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatUnreadMentionCount) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatUnreadMentionCount) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatUnreadMentionCount) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatUnreadMentionCount) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatUnreadMentionCount) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatUnreadMentionCount) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatUnreadMentionCount) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatUnreadMentionCount) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatUnreadMentionCount) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatUnreadMentionCount) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatUnreadMentionCount) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatUnreadMentionCount) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatUnreadMentionCount) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatUnreadMentionCount) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatUnreadMentionCount) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatUnreadMentionCount) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatUnreadMentionCount) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatUnreadMentionCount) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatUnreadMentionCount) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatUnreadMentionCount) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatUnreadMentionCount) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatUnreadMentionCount) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatUnreadMentionCount) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatUnreadMentionCount) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatUnreadMentionCount) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatUnreadMentionCount) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatUnreadMentionCount) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatUnreadMentionCount) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatUnreadMentionCount) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatUnreadMentionCount) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatUnreadMentionCount) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatUnreadMentionCount) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatUnreadMentionCount) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatUnreadMentionCount) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatUnreadMentionCount) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatUnreadMentionCount) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatUnreadMentionCount) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatUnreadMentionCount) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatUnreadMentionCount) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatUnreadMentionCount) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatUnreadMentionCount) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatUnreadMentionCount) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatUnreadMentionCount) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatUnreadMentionCount) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatUnreadMentionCount) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatUnreadMentionCount) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatUnreadMentionCount) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatUnreadMentionCount) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatUnreadMentionCount) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatUnreadMentionCount) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatUnreadMentionCount) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatUnreadMentionCount) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatUnreadMentionCount) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatUnreadMentionCount) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatUnreadMentionCount) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatUnreadMentionCount) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatUnreadMentionCount) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatUnreadMentionCount) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatUnreadMentionCount) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatUnreadMentionCount) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatUnreadMentionCount) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatUnreadMentionCount) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatUnreadMentionCount) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatUnreadMentionCount) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatUnreadMentionCount) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatUnreadMentionCount) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatUnreadMentionCount) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatUnreadMentionCount) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatUnreadMentionCount) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatUnreadMentionCount) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatUnreadMentionCount) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatUnreadMentionCount) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatUnreadMentionCount) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatUnreadMentionCount) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatUnreadMentionCount) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatUnreadMentionCount) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatUnreadMentionCount) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatUnreadMentionCount) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatUnreadMentionCount) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatUnreadMentionCount) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatUnreadMentionCount) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatUnreadMentionCount) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatUnreadMentionCount) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatUnreadMentionCount) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatUnreadMentionCount) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatUnreadMentionCount) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatUnreadMentionCount) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatUnreadMentionCount) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatUnreadMentionCount) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatUnreadMentionCount) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatUnreadMentionCount) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatUnreadMentionCount) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatUnreadMentionCount) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatUnreadMentionCount) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatUnreadMentionCount) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatUnreadMentionCount) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatUnreadMentionCount) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatUnreadMentionCount) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatUnreadMentionCount) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatUnreadMentionCount) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatUnreadMentionCount) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatUnreadMentionCount) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatUnreadMentionCount) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatUnreadMentionCount) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatUnreadMentionCount) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatUnreadMentionCount) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatUnreadMentionCount) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatUnreadMentionCount) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatUnreadMentionCount) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatUnreadMentionCount) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatUnreadMentionCount) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatUnreadMentionCount) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatUnreadMentionCount) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatUnreadMentionCount) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatUnreadMentionCount) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatUnreadMentionCount) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatUnreadMentionCount) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatUnreadMentionCount) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatUnreadMentionCount) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatUnreadMentionCount) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatUnreadMentionCount) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatUnreadMentionCount) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatUnreadMentionCount) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatUnreadMentionCount) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatUnreadMentionCount) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatUnreadMentionCount) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatUnreadMentionCount) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatUnreadMentionCount) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatUnreadMentionCount) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatUnreadMentionCount) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatUnreadMentionCount) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatUnreadMentionCount) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatUnreadMentionCount) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatUnreadMentionCount) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatUnreadMentionCount) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatUnreadMentionCount) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatUnreadMentionCount) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatUnreadMentionCount) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatUnreadMentionCount) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatUnreadMentionCount) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatUnreadMentionCount) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatUnreadMentionCount) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatUnreadMentionCount) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatUnreadMentionCount) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatUnreadMentionCount) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatUnreadMentionCount) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatUnreadMentionCount) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatUnreadMentionCount) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatUnreadMentionCount) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatUnreadMentionCount) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatUnreadMentionCount) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatUnreadMentionCount) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatUnreadMentionCount) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatUnreadMentionCount) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatUnreadMentionCount) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatUnreadMentionCount) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatUnreadMentionCount) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatUnreadMentionCount) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatUnreadMentionCount) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatUnreadMentionCount) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatUnreadMentionCount) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatUnreadMentionCount) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatUnreadMentionCount) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatUnreadMentionCount) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatUnreadMentionCount) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatUnreadMentionCount) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatUnreadMentionCount) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatUnreadMentionCount) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatUnreadMentionCount) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatUnreadMentionCount) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatUnreadMentionCount) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatUnreadMentionCount) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatUnreadMentionCount) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatUnreadMentionCount) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatUnreadMentionCount) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatUnreadMentionCount) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatUnreadMentionCount) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatUnreadMentionCount) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatUnreadMentionCount) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatUnreadMentionCount) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatUnreadMentionCount) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatUnreadMentionCount) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatUnreadMentionCount) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatUnreadMentionCount) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatUnreadMentionCount) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatUnreadMentionCount) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatUnreadMentionCount) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatUnreadMentionCount) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatUnreadMentionCount) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatUnreadMentionCount) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatUnreadMentionCount) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatUnreadMentionCount) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatUnreadMentionCount) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatUnreadMentionCount) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatUnreadMentionCount) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatUnreadMentionCount) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatUnreadMentionCount) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatUnreadMentionCount) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatUnreadMentionCount) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatUnreadMentionCount) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatUnreadMentionCount) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatUnreadMentionCount) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatUnreadMentionCount) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatUnreadMentionCount) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatUnreadMentionCount) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatUnreadMentionCount) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatUnreadMentionCount) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatUnreadMentionCount) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatUnreadMentionCount) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatUnreadMentionCount) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatUnreadMentionCount) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatUnreadMentionCount) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatUnreadMentionCount) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatUnreadMentionCount) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatUnreadMentionCount) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatUnreadMentionCount) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatUnreadMentionCount) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatUnreadMentionCount) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatUnreadMentionCount) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatUnreadMentionCount) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatUnreadMentionCount) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatUnreadMentionCount) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatUnreadMentionCount) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatUnreadMentionCount) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatUnreadMentionCount) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatUnreadMentionCount) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatUnreadMentionCount) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatUnreadMentionCount) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatUnreadMentionCount) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatUnreadMentionCount) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatUnreadMentionCount) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatUnreadMentionCount) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatUnreadMentionCount) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatUnreadMentionCount) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatUnreadMentionCount) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatUnreadMentionCount) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatUnreadMentionCount) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatUnreadMentionCount) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatUnreadMentionCount) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatUnreadMentionCount) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatUnreadReactionCount) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatUnreadReactionCount) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatUnreadReactionCount) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatUnreadReactionCount) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatUnreadReactionCount) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatUnreadReactionCount) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatUnreadReactionCount) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatUnreadReactionCount) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatUnreadReactionCount) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatUnreadReactionCount) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatUnreadReactionCount) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatUnreadReactionCount) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatUnreadReactionCount) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatUnreadReactionCount) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatUnreadReactionCount) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatUnreadReactionCount) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatUnreadReactionCount) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatUnreadReactionCount) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatUnreadReactionCount) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatUnreadReactionCount) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatUnreadReactionCount) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatUnreadReactionCount) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatUnreadReactionCount) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatUnreadReactionCount) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatUnreadReactionCount) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatUnreadReactionCount) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatUnreadReactionCount) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatUnreadReactionCount) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatUnreadReactionCount) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatUnreadReactionCount) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatUnreadReactionCount) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatUnreadReactionCount) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatUnreadReactionCount) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatUnreadReactionCount) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatUnreadReactionCount) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatUnreadReactionCount) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatUnreadReactionCount) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatUnreadReactionCount) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatUnreadReactionCount) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatUnreadReactionCount) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatUnreadReactionCount) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatUnreadReactionCount) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatUnreadReactionCount) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatUnreadReactionCount) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatUnreadReactionCount) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatUnreadReactionCount) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatUnreadReactionCount) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatUnreadReactionCount) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatUnreadReactionCount) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatUnreadReactionCount) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatUnreadReactionCount) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatUnreadReactionCount) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatUnreadReactionCount) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatUnreadReactionCount) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatUnreadReactionCount) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatUnreadReactionCount) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatUnreadReactionCount) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatUnreadReactionCount) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatUnreadReactionCount) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatUnreadReactionCount) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatUnreadReactionCount) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatUnreadReactionCount) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatUnreadReactionCount) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatUnreadReactionCount) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatUnreadReactionCount) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatUnreadReactionCount) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatUnreadReactionCount) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatUnreadReactionCount) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatUnreadReactionCount) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatUnreadReactionCount) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatUnreadReactionCount) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatUnreadReactionCount) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatUnreadReactionCount) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatUnreadReactionCount) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatUnreadReactionCount) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatUnreadReactionCount) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatUnreadReactionCount) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatUnreadReactionCount) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatUnreadReactionCount) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatUnreadReactionCount) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatUnreadReactionCount) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatUnreadReactionCount) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatUnreadReactionCount) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatUnreadReactionCount) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatUnreadReactionCount) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatUnreadReactionCount) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatUnreadReactionCount) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatUnreadReactionCount) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatUnreadReactionCount) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatUnreadReactionCount) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatUnreadReactionCount) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatUnreadReactionCount) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatUnreadReactionCount) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatUnreadReactionCount) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatUnreadReactionCount) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatUnreadReactionCount) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatUnreadReactionCount) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatUnreadReactionCount) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatUnreadReactionCount) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatUnreadReactionCount) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatUnreadReactionCount) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatUnreadReactionCount) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatUnreadReactionCount) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatUnreadReactionCount) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatUnreadReactionCount) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatUnreadReactionCount) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatUnreadReactionCount) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatUnreadReactionCount) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatUnreadReactionCount) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatUnreadReactionCount) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatUnreadReactionCount) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatUnreadReactionCount) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatUnreadReactionCount) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatUnreadReactionCount) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatUnreadReactionCount) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatUnreadReactionCount) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatUnreadReactionCount) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatUnreadReactionCount) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatUnreadReactionCount) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatUnreadReactionCount) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatUnreadReactionCount) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatUnreadReactionCount) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatUnreadReactionCount) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatUnreadReactionCount) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatUnreadReactionCount) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatUnreadReactionCount) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatUnreadReactionCount) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatUnreadReactionCount) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatUnreadReactionCount) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatUnreadReactionCount) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatUnreadReactionCount) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatUnreadReactionCount) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatUnreadReactionCount) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatUnreadReactionCount) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatUnreadReactionCount) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatUnreadReactionCount) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatUnreadReactionCount) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatUnreadReactionCount) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatUnreadReactionCount) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatUnreadReactionCount) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatUnreadReactionCount) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatUnreadReactionCount) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatUnreadReactionCount) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatUnreadReactionCount) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatUnreadReactionCount) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatUnreadReactionCount) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatUnreadReactionCount) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatUnreadReactionCount) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatUnreadReactionCount) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatUnreadReactionCount) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatUnreadReactionCount) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatUnreadReactionCount) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatUnreadReactionCount) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatUnreadReactionCount) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatUnreadReactionCount) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatUnreadReactionCount) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatUnreadReactionCount) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatUnreadReactionCount) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatUnreadReactionCount) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatUnreadReactionCount) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatUnreadReactionCount) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatUnreadReactionCount) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatUnreadReactionCount) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatUnreadReactionCount) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatUnreadReactionCount) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatUnreadReactionCount) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatUnreadReactionCount) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatUnreadReactionCount) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatUnreadReactionCount) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatUnreadReactionCount) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatUnreadReactionCount) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatUnreadReactionCount) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatUnreadReactionCount) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatUnreadReactionCount) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatUnreadReactionCount) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatUnreadReactionCount) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatUnreadReactionCount) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatUnreadReactionCount) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatUnreadReactionCount) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatUnreadReactionCount) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatUnreadReactionCount) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatUnreadReactionCount) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatUnreadReactionCount) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatUnreadReactionCount) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatUnreadReactionCount) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatUnreadReactionCount) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatUnreadReactionCount) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatUnreadReactionCount) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatUnreadReactionCount) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatUnreadReactionCount) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatUnreadReactionCount) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatUnreadReactionCount) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatUnreadReactionCount) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatUnreadReactionCount) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatUnreadReactionCount) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatUnreadReactionCount) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatUnreadReactionCount) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatUnreadReactionCount) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatUnreadReactionCount) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatUnreadReactionCount) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatUnreadReactionCount) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatUnreadReactionCount) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatUnreadReactionCount) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatUnreadReactionCount) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatUnreadReactionCount) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatUnreadReactionCount) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatUnreadReactionCount) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatUnreadReactionCount) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatUnreadReactionCount) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatUnreadReactionCount) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatUnreadReactionCount) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatUnreadReactionCount) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatUnreadReactionCount) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatUnreadReactionCount) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatUnreadReactionCount) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatUnreadReactionCount) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatUnreadReactionCount) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatUnreadReactionCount) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatUnreadReactionCount) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatUnreadReactionCount) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatUnreadReactionCount) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatUnreadReactionCount) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatUnreadReactionCount) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatUnreadReactionCount) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatUnreadReactionCount) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatUnreadReactionCount) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatUnreadReactionCount) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatUnreadReactionCount) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatUnreadReactionCount) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatUnreadReactionCount) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatUnreadReactionCount) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatUnreadReactionCount) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatUnreadReactionCount) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatUnreadReactionCount) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatUnreadReactionCount) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatUnreadReactionCount) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatUnreadReactionCount) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatUnreadReactionCount) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatUnreadReactionCount) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatUnreadReactionCount) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatUnreadReactionCount) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatUnreadReactionCount) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatUnreadReactionCount) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatUnreadReactionCount) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatUnreadReactionCount) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatUnreadReactionCount) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatUnreadReactionCount) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatUnreadReactionCount) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatUnreadReactionCount) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatUnreadReactionCount) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatUnreadReactionCount) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatUnreadReactionCount) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatVideoChat) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatVideoChat) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatVideoChat) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatVideoChat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatVideoChat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatVideoChat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatVideoChat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatVideoChat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatVideoChat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatVideoChat) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatVideoChat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatVideoChat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatVideoChat) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatVideoChat) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatVideoChat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatVideoChat) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatVideoChat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatVideoChat) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatVideoChat) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatVideoChat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatVideoChat) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatVideoChat) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatVideoChat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatVideoChat) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatVideoChat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatVideoChat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatVideoChat) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatVideoChat) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatVideoChat) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatVideoChat) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatVideoChat) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatVideoChat) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatVideoChat) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatVideoChat) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatVideoChat) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatVideoChat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatVideoChat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatVideoChat) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatVideoChat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatVideoChat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatVideoChat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatVideoChat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatVideoChat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatVideoChat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatVideoChat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatVideoChat) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatVideoChat) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatVideoChat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatVideoChat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatVideoChat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatVideoChat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatVideoChat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatVideoChat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatVideoChat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatVideoChat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatVideoChat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatVideoChat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatVideoChat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatVideoChat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatVideoChat) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatVideoChat) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatVideoChat) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatVideoChat) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatVideoChat) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatVideoChat) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatVideoChat) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatVideoChat) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatVideoChat) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatVideoChat) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatVideoChat) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatVideoChat) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatVideoChat) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatVideoChat) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatVideoChat) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatVideoChat) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatVideoChat) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatVideoChat) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatVideoChat) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatVideoChat) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatVideoChat) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatVideoChat) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatVideoChat) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatVideoChat) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatVideoChat) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatVideoChat) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatVideoChat) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatVideoChat) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatVideoChat) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatVideoChat) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatVideoChat) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatVideoChat) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatVideoChat) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatVideoChat) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatVideoChat) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatVideoChat) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatVideoChat) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatVideoChat) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatVideoChat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatVideoChat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatVideoChat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatVideoChat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatVideoChat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatVideoChat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatVideoChat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatVideoChat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatVideoChat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatVideoChat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatVideoChat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatVideoChat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatVideoChat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatVideoChat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatVideoChat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatVideoChat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatVideoChat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatVideoChat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatVideoChat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatVideoChat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatVideoChat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatVideoChat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatVideoChat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatVideoChat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatVideoChat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatVideoChat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatVideoChat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatVideoChat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatVideoChat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatVideoChat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatVideoChat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatVideoChat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatVideoChat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatVideoChat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatVideoChat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatVideoChat) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatVideoChat) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatVideoChat) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatVideoChat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatVideoChat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatVideoChat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatVideoChat) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatVideoChat) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatVideoChat) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatVideoChat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatVideoChat) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatVideoChat) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatVideoChat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatVideoChat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatVideoChat) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatVideoChat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatVideoChat) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatVideoChat) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatVideoChat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatVideoChat) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatVideoChat) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatVideoChat) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatVideoChat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatVideoChat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatVideoChat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatVideoChat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatVideoChat) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatVideoChat) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatVideoChat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatVideoChat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatVideoChat) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatVideoChat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatVideoChat) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatVideoChat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatVideoChat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatVideoChat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatVideoChat) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatVideoChat) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatVideoChat) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatVideoChat) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatVideoChat) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatVideoChat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatVideoChat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatVideoChat) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatVideoChat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatVideoChat) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatVideoChat) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatVideoChat) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatVideoChat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatVideoChat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatVideoChat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatVideoChat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatVideoChat) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatVideoChat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatVideoChat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatVideoChat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatVideoChat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatVideoChat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatVideoChat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatVideoChat) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatVideoChat) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatVideoChat) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatVideoChat) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatVideoChat) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatVideoChat) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatVideoChat) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatVideoChat) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatVideoChat) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatVideoChat) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatVideoChat) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatVideoChat) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatVideoChat) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatVideoChat) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatVideoChat) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatVideoChat) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatVideoChat) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatVideoChat) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatVideoChat) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatVideoChat) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatVideoChat) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatVideoChat) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatVideoChat) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatVideoChat) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatVideoChat) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatVideoChat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatVideoChat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatVideoChat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatVideoChat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatVideoChat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatVideoChat) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatVideoChat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatVideoChat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatVideoChat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatVideoChat) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatVideoChat) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatVideoChat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatVideoChat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatVideoChat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatVideoChat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatVideoChat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatVideoChat) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatVideoChat) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatVideoChat) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatVideoChat) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatVideoChat) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatVideoChat) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatVideoChat) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatVideoChat) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatVideoChat) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatVideoChat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatVideoChat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatVideoChat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatVideoChat) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatVideoChat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatVideoChat) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatVideoChat) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatVideoChat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatVideoChat) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatVideoChat) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatVideoChat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateChatViewAsTopics) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateChatViewAsTopics) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateChatViewAsTopics) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateChatViewAsTopics) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateChatViewAsTopics) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateChatViewAsTopics) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateChatViewAsTopics) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateChatViewAsTopics) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateChatViewAsTopics) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateChatViewAsTopics) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateChatViewAsTopics) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateChatViewAsTopics) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateChatViewAsTopics) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateChatViewAsTopics) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateChatViewAsTopics) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateChatViewAsTopics) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateChatViewAsTopics) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateChatViewAsTopics) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateChatViewAsTopics) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateChatViewAsTopics) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateChatViewAsTopics) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateChatViewAsTopics) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateChatViewAsTopics) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateChatViewAsTopics) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateChatViewAsTopics) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateChatViewAsTopics) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateChatViewAsTopics) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateChatViewAsTopics) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateChatViewAsTopics) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateChatViewAsTopics) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateChatViewAsTopics) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateChatViewAsTopics) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateChatViewAsTopics) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateChatViewAsTopics) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateChatViewAsTopics) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateChatViewAsTopics) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateChatViewAsTopics) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateChatViewAsTopics) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateChatViewAsTopics) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateChatViewAsTopics) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateChatViewAsTopics) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateChatViewAsTopics) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateChatViewAsTopics) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateChatViewAsTopics) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateChatViewAsTopics) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateChatViewAsTopics) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateChatViewAsTopics) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateChatViewAsTopics) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateChatViewAsTopics) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateChatViewAsTopics) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateChatViewAsTopics) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateChatViewAsTopics) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateChatViewAsTopics) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateChatViewAsTopics) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateChatViewAsTopics) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateChatViewAsTopics) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateChatViewAsTopics) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateChatViewAsTopics) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateChatViewAsTopics) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateChatViewAsTopics) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateChatViewAsTopics) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateChatViewAsTopics) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateChatViewAsTopics) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateChatViewAsTopics) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateChatViewAsTopics) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateChatViewAsTopics) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateChatViewAsTopics) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateChatViewAsTopics) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateChatViewAsTopics) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateChatViewAsTopics) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateChatViewAsTopics) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateChatViewAsTopics) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateChatViewAsTopics) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateChatViewAsTopics) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateChatViewAsTopics) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateChatViewAsTopics) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateChatViewAsTopics) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateChatViewAsTopics) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateChatViewAsTopics) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateChatViewAsTopics) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateChatViewAsTopics) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateChatViewAsTopics) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateChatViewAsTopics) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateChatViewAsTopics) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateChatViewAsTopics) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateChatViewAsTopics) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateChatViewAsTopics) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateChatViewAsTopics) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateChatViewAsTopics) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateChatViewAsTopics) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateChatViewAsTopics) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateChatViewAsTopics) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateChatViewAsTopics) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateChatViewAsTopics) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateChatViewAsTopics) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateChatViewAsTopics) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateChatViewAsTopics) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateChatViewAsTopics) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateChatViewAsTopics) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateChatViewAsTopics) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateChatViewAsTopics) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateChatViewAsTopics) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateChatViewAsTopics) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateChatViewAsTopics) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateChatViewAsTopics) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateChatViewAsTopics) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateChatViewAsTopics) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateChatViewAsTopics) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateChatViewAsTopics) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateChatViewAsTopics) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateChatViewAsTopics) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateChatViewAsTopics) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateChatViewAsTopics) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateChatViewAsTopics) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateChatViewAsTopics) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateChatViewAsTopics) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateChatViewAsTopics) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateChatViewAsTopics) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateChatViewAsTopics) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateChatViewAsTopics) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateChatViewAsTopics) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateChatViewAsTopics) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateChatViewAsTopics) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateChatViewAsTopics) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateChatViewAsTopics) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateChatViewAsTopics) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateChatViewAsTopics) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateChatViewAsTopics) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateChatViewAsTopics) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateChatViewAsTopics) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateChatViewAsTopics) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateChatViewAsTopics) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateChatViewAsTopics) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateChatViewAsTopics) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateChatViewAsTopics) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateChatViewAsTopics) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateChatViewAsTopics) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateChatViewAsTopics) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateChatViewAsTopics) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateChatViewAsTopics) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateChatViewAsTopics) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateChatViewAsTopics) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateChatViewAsTopics) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateChatViewAsTopics) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateChatViewAsTopics) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateChatViewAsTopics) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateChatViewAsTopics) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateChatViewAsTopics) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateChatViewAsTopics) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateChatViewAsTopics) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateChatViewAsTopics) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateChatViewAsTopics) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateChatViewAsTopics) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateChatViewAsTopics) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateChatViewAsTopics) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateChatViewAsTopics) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateChatViewAsTopics) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateChatViewAsTopics) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateChatViewAsTopics) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateChatViewAsTopics) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateChatViewAsTopics) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateChatViewAsTopics) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateChatViewAsTopics) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateChatViewAsTopics) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateChatViewAsTopics) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateChatViewAsTopics) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateChatViewAsTopics) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateChatViewAsTopics) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateChatViewAsTopics) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateChatViewAsTopics) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateChatViewAsTopics) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateChatViewAsTopics) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateChatViewAsTopics) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateChatViewAsTopics) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateChatViewAsTopics) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateChatViewAsTopics) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateChatViewAsTopics) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateChatViewAsTopics) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateChatViewAsTopics) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateChatViewAsTopics) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateChatViewAsTopics) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateChatViewAsTopics) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateChatViewAsTopics) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateChatViewAsTopics) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateChatViewAsTopics) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateChatViewAsTopics) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateChatViewAsTopics) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateChatViewAsTopics) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateChatViewAsTopics) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateChatViewAsTopics) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateChatViewAsTopics) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateChatViewAsTopics) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateChatViewAsTopics) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateChatViewAsTopics) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateChatViewAsTopics) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateChatViewAsTopics) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateChatViewAsTopics) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateChatViewAsTopics) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateChatViewAsTopics) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateChatViewAsTopics) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateChatViewAsTopics) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateChatViewAsTopics) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateChatViewAsTopics) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateChatViewAsTopics) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateChatViewAsTopics) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateChatViewAsTopics) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateChatViewAsTopics) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateChatViewAsTopics) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateChatViewAsTopics) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateChatViewAsTopics) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateChatViewAsTopics) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateChatViewAsTopics) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateChatViewAsTopics) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateChatViewAsTopics) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateChatViewAsTopics) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateChatViewAsTopics) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateChatViewAsTopics) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateChatViewAsTopics) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateChatViewAsTopics) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateChatViewAsTopics) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateChatViewAsTopics) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateChatViewAsTopics) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateChatViewAsTopics) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateChatViewAsTopics) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateChatViewAsTopics) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateChatViewAsTopics) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateChatViewAsTopics) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateChatViewAsTopics) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateChatViewAsTopics) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateChatViewAsTopics) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateChatViewAsTopics) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateChatViewAsTopics) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateChatViewAsTopics) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateChatViewAsTopics) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateChatViewAsTopics) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateChatViewAsTopics) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateChatViewAsTopics) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateChatViewAsTopics) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateChatViewAsTopics) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateChatViewAsTopics) ToggleChatViewAsTopics(client *Client) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, u.ViewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateChatViewAsTopics) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateChatViewAsTopics) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateChatViewAsTopics) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateChatViewAsTopics) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateChatViewAsTopics) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateChatViewAsTopics) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateChatViewAsTopics) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateChatViewAsTopics) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateChatViewAsTopics) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateChatViewAsTopics) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateChatViewAsTopics) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateChatViewAsTopics) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// DeleteDefaultBackground is a helper method for Client.DeleteDefaultBackground
func (u UpdateDefaultBackground) DeleteDefaultBackground(client *Client) (*Ok, error) {
	return client.DeleteDefaultBackground(u.ForDarkTheme)
}

// GetInstalledBackgrounds is a helper method for Client.GetInstalledBackgrounds
func (u UpdateDefaultBackground) GetInstalledBackgrounds(client *Client) (*Backgrounds, error) {
	return client.GetInstalledBackgrounds(u.ForDarkTheme)
}

// SetDefaultBackground is a helper method for Client.SetDefaultBackground
func (u UpdateDefaultBackground) SetDefaultBackground(client *Client, opts *SetDefaultBackgroundOpts) (*Background, error) {
	return client.SetDefaultBackground(u.ForDarkTheme, opts)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateDefaultPaidReactionType) SetPaidMessageReactionType(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, messageId, u.Type)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateDefaultReactionType) AddMessageReaction(client *Client, chatId int64, messageId int64, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, messageId, u.ReactionType, isBig, updateRecentReactions)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateDefaultReactionType) RemoveMessageReaction(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, messageId, u.ReactionType)
}

// SetDefaultReactionType is a helper method for Client.SetDefaultReactionType
func (u UpdateDefaultReactionType) SetDefaultReactionType(client *Client) (*Ok, error) {
	return client.SetDefaultReactionType(u.ReactionType)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateDeleteMessages) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateDeleteMessages) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateDeleteMessages) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateDeleteMessages) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateDeleteMessages) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateDeleteMessages) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateDeleteMessages) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateDeleteMessages) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateDeleteMessages) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateDeleteMessages) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateDeleteMessages) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateDeleteMessages) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateDeleteMessages) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateDeleteMessages) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateDeleteMessages) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateDeleteMessages) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateDeleteMessages) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateDeleteMessages) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateDeleteMessages) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateDeleteMessages) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateDeleteMessages) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateDeleteMessages) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateDeleteMessages) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateDeleteMessages) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateDeleteMessages) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateDeleteMessages) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateDeleteMessages) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteBusinessMessages is a helper method for Client.DeleteBusinessMessages
func (u UpdateDeleteMessages) DeleteBusinessMessages(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessMessages(businessConnectionId, u.MessageIds)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateDeleteMessages) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateDeleteMessages) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateDeleteMessages) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateDeleteMessages) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateDeleteMessages) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateDeleteMessages) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateDeleteMessages) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateDeleteMessages) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateDeleteMessages) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateDeleteMessages) DeleteMessages(client *Client, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, u.MessageIds, revoke)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (u UpdateDeleteMessages) DeleteQuickReplyShortcutMessages(client *Client, shortcutId int32) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(shortcutId, u.MessageIds)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateDeleteMessages) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateDeleteMessages) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateDeleteMessages) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateDeleteMessages) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateDeleteMessages) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateDeleteMessages) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateDeleteMessages) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateDeleteMessages) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateDeleteMessages) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateDeleteMessages) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateDeleteMessages) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateDeleteMessages) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateDeleteMessages) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateDeleteMessages) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateDeleteMessages) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateDeleteMessages) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateDeleteMessages) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateDeleteMessages) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateDeleteMessages) ForwardMessages(client *Client, fromChatId int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, u.MessageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateDeleteMessages) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateDeleteMessages) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateDeleteMessages) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateDeleteMessages) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateDeleteMessages) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateDeleteMessages) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateDeleteMessages) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateDeleteMessages) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateDeleteMessages) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateDeleteMessages) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateDeleteMessages) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateDeleteMessages) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateDeleteMessages) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateDeleteMessages) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateDeleteMessages) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateDeleteMessages) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateDeleteMessages) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateDeleteMessages) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateDeleteMessages) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateDeleteMessages) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateDeleteMessages) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateDeleteMessages) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateDeleteMessages) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateDeleteMessages) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateDeleteMessages) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateDeleteMessages) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateDeleteMessages) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateDeleteMessages) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateDeleteMessages) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateDeleteMessages) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateDeleteMessages) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateDeleteMessages) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateDeleteMessages) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateDeleteMessages) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateDeleteMessages) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateDeleteMessages) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateDeleteMessages) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateDeleteMessages) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateDeleteMessages) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateDeleteMessages) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateDeleteMessages) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateDeleteMessages) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateDeleteMessages) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateDeleteMessages) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateDeleteMessages) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateDeleteMessages) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateDeleteMessages) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateDeleteMessages) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateDeleteMessages) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateDeleteMessages) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateDeleteMessages) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateDeleteMessages) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateDeleteMessages) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateDeleteMessages) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateDeleteMessages) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateDeleteMessages) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateDeleteMessages) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateDeleteMessages) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateDeleteMessages) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateDeleteMessages) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateDeleteMessages) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateDeleteMessages) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateDeleteMessages) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateDeleteMessages) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateDeleteMessages) GetMessages(client *Client) (*Messages, error) {
	return client.GetMessages(u.ChatId, u.MessageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateDeleteMessages) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateDeleteMessages) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateDeleteMessages) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateDeleteMessages) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateDeleteMessages) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateDeleteMessages) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateDeleteMessages) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateDeleteMessages) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateDeleteMessages) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateDeleteMessages) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateDeleteMessages) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateDeleteMessages) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateDeleteMessages) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateDeleteMessages) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateDeleteMessages) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateDeleteMessages) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateDeleteMessages) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateDeleteMessages) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateDeleteMessages) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateDeleteMessages) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateDeleteMessages) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateDeleteMessages) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateDeleteMessages) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateDeleteMessages) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateDeleteMessages) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateDeleteMessages) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateDeleteMessages) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateDeleteMessages) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateDeleteMessages) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateDeleteMessages) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateDeleteMessages) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateDeleteMessages) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateDeleteMessages) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateDeleteMessages) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateDeleteMessages) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateDeleteMessages) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// ReaddQuickReplyShortcutMessages is a helper method for Client.ReaddQuickReplyShortcutMessages
func (u UpdateDeleteMessages) ReaddQuickReplyShortcutMessages(client *Client, shortcutName string) (*QuickReplyMessages, error) {
	return client.ReaddQuickReplyShortcutMessages(shortcutName, u.MessageIds)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateDeleteMessages) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateDeleteMessages) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateDeleteMessages) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateDeleteMessages) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateDeleteMessages) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateDeleteMessages) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateDeleteMessages) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateDeleteMessages) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateDeleteMessages) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateDeleteMessages) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateDeleteMessages) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateDeleteMessages) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateDeleteMessages) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateDeleteMessages) ReportChat(client *Client, optionId string, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, u.MessageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateDeleteMessages) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateDeleteMessages) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateDeleteMessages) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (u UpdateDeleteMessages) ReportSupergroupSpam(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupSpam(supergroupId, u.MessageIds)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateDeleteMessages) ResendMessages(client *Client, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, u.MessageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateDeleteMessages) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateDeleteMessages) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateDeleteMessages) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateDeleteMessages) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateDeleteMessages) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateDeleteMessages) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateDeleteMessages) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateDeleteMessages) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateDeleteMessages) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateDeleteMessages) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateDeleteMessages) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateDeleteMessages) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateDeleteMessages) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateDeleteMessages) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateDeleteMessages) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateDeleteMessages) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateDeleteMessages) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateDeleteMessages) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateDeleteMessages) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateDeleteMessages) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateDeleteMessages) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateDeleteMessages) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateDeleteMessages) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateDeleteMessages) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateDeleteMessages) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateDeleteMessages) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateDeleteMessages) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateDeleteMessages) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateDeleteMessages) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateDeleteMessages) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateDeleteMessages) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateDeleteMessages) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateDeleteMessages) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateDeleteMessages) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateDeleteMessages) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateDeleteMessages) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateDeleteMessages) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateDeleteMessages) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateDeleteMessages) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateDeleteMessages) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateDeleteMessages) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateDeleteMessages) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateDeleteMessages) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateDeleteMessages) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateDeleteMessages) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateDeleteMessages) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateDeleteMessages) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateDeleteMessages) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateDeleteMessages) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateDeleteMessages) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateDeleteMessages) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateDeleteMessages) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateDeleteMessages) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateDeleteMessages) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateDeleteMessages) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateDeleteMessages) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateDeleteMessages) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateDeleteMessages) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateDeleteMessages) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateDeleteMessages) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateDeleteMessages) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateDeleteMessages) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateDeleteMessages) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateDeleteMessages) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateDeleteMessages) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateDeleteMessages) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateDeleteMessages) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateDeleteMessages) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateDeleteMessages) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateDeleteMessages) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateDeleteMessages) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateDeleteMessages) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateDeleteMessages) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateDeleteMessages) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateDeleteMessages) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateDeleteMessages) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateDeleteMessages) ViewMessages(client *Client, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, u.MessageIds, forceRead, opts)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateFileDownload) AddFileToDownloads(client *Client, chatId int64, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(u.FileId, chatId, messageId, priority)
}

// AddProfileAudio is a helper method for Client.AddProfileAudio
func (u UpdateFileDownload) AddProfileAudio(client *Client) (*Ok, error) {
	return client.AddProfileAudio(u.FileId)
}

// CancelDownloadFile is a helper method for Client.CancelDownloadFile
func (u UpdateFileDownload) CancelDownloadFile(client *Client, onlyIfPending bool) (*Ok, error) {
	return client.CancelDownloadFile(u.FileId, onlyIfPending)
}

// CancelPreliminaryUploadFile is a helper method for Client.CancelPreliminaryUploadFile
func (u UpdateFileDownload) CancelPreliminaryUploadFile(client *Client) (*Ok, error) {
	return client.CancelPreliminaryUploadFile(u.FileId)
}

// DeleteFile is a helper method for Client.DeleteFile
func (u UpdateFileDownload) DeleteFile(client *Client) (*Ok, error) {
	return client.DeleteFile(u.FileId)
}

// DownloadFile is a helper method for Client.DownloadFile
func (u UpdateFileDownload) DownloadFile(client *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	return client.DownloadFile(u.FileId, priority, offset, limit, synchronous)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (u UpdateFileDownload) EditBotMediaPreview(client *Client, botUserId int64, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(botUserId, languageCode, u.FileId, content)
}

// GetAttachedStickerSets is a helper method for Client.GetAttachedStickerSets
func (u UpdateFileDownload) GetAttachedStickerSets(client *Client) (*StickerSets, error) {
	return client.GetAttachedStickerSets(u.FileId)
}

// GetFile is a helper method for Client.GetFile
func (u UpdateFileDownload) GetFile(client *Client) (*File, error) {
	return client.GetFile(u.FileId)
}

// GetFileDownloadedPrefixSize is a helper method for Client.GetFileDownloadedPrefixSize
func (u UpdateFileDownload) GetFileDownloadedPrefixSize(client *Client, offset int64) (*FileDownloadedPrefixSize, error) {
	return client.GetFileDownloadedPrefixSize(u.FileId, offset)
}

// GetSuggestedFileName is a helper method for Client.GetSuggestedFileName
func (u UpdateFileDownload) GetSuggestedFileName(client *Client, directory string) (*Text, error) {
	return client.GetSuggestedFileName(u.FileId, directory)
}

// IsProfileAudio is a helper method for Client.IsProfileAudio
func (u UpdateFileDownload) IsProfileAudio(client *Client) (*Ok, error) {
	return client.IsProfileAudio(u.FileId)
}

// ReadFilePart is a helper method for Client.ReadFilePart
func (u UpdateFileDownload) ReadFilePart(client *Client, offset int64, count int64) (*Data, error) {
	return client.ReadFilePart(u.FileId, offset, count)
}

// RemoveFileFromDownloads is a helper method for Client.RemoveFileFromDownloads
func (u UpdateFileDownload) RemoveFileFromDownloads(client *Client, deleteFromCache bool) (*Ok, error) {
	return client.RemoveFileFromDownloads(u.FileId, deleteFromCache)
}

// RemoveProfileAudio is a helper method for Client.RemoveProfileAudio
func (u UpdateFileDownload) RemoveProfileAudio(client *Client) (*Ok, error) {
	return client.RemoveProfileAudio(u.FileId)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateFileDownload) ReportChatPhoto(client *Client, chatId int64, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(chatId, u.FileId, reason, text)
}

// SetProfileAudioPosition is a helper method for Client.SetProfileAudioPosition
func (u UpdateFileDownload) SetProfileAudioPosition(client *Client, afterFileId int32) (*Ok, error) {
	return client.SetProfileAudioPosition(u.FileId, afterFileId)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateFileDownload) ToggleBusinessConnectedBotChatIsPaused(client *Client, chatId int64) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(chatId, u.IsPaused)
}

// ToggleDownloadIsPaused is a helper method for Client.ToggleDownloadIsPaused
func (u UpdateFileDownload) ToggleDownloadIsPaused(client *Client) (*Ok, error) {
	return client.ToggleDownloadIsPaused(u.FileId, u.IsPaused)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (u UpdateFileDownload) ToggleGroupCallScreenSharingIsPaused(client *Client, groupCallId int32) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(groupCallId, u.IsPaused)
}

// FinishFileGeneration is a helper method for Client.FinishFileGeneration
func (u UpdateFileGenerationStart) FinishFileGeneration(client *Client, opts *FinishFileGenerationOpts) (*Ok, error) {
	return client.FinishFileGeneration(u.GenerationId, opts)
}

// SetFileGenerationProgress is a helper method for Client.SetFileGenerationProgress
func (u UpdateFileGenerationStart) SetFileGenerationProgress(client *Client, expectedSize int64, localPrefixSize int64) (*Ok, error) {
	return client.SetFileGenerationProgress(u.GenerationId, expectedSize, localPrefixSize)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (u UpdateFileGenerationStart) WriteGeneratedFilePart(client *Client, offset int64, data string) (*Ok, error) {
	return client.WriteGeneratedFilePart(u.GenerationId, offset, data)
}

// FinishFileGeneration is a helper method for Client.FinishFileGeneration
func (u UpdateFileGenerationStop) FinishFileGeneration(client *Client, opts *FinishFileGenerationOpts) (*Ok, error) {
	return client.FinishFileGeneration(u.GenerationId, opts)
}

// SetFileGenerationProgress is a helper method for Client.SetFileGenerationProgress
func (u UpdateFileGenerationStop) SetFileGenerationProgress(client *Client, expectedSize int64, localPrefixSize int64) (*Ok, error) {
	return client.SetFileGenerationProgress(u.GenerationId, expectedSize, localPrefixSize)
}

// WriteGeneratedFilePart is a helper method for Client.WriteGeneratedFilePart
func (u UpdateFileGenerationStop) WriteGeneratedFilePart(client *Client, offset int64, data string) (*Ok, error) {
	return client.WriteGeneratedFilePart(u.GenerationId, offset, data)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateFileRemovedFromDownloads) AddFileToDownloads(client *Client, chatId int64, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(u.FileId, chatId, messageId, priority)
}

// AddProfileAudio is a helper method for Client.AddProfileAudio
func (u UpdateFileRemovedFromDownloads) AddProfileAudio(client *Client) (*Ok, error) {
	return client.AddProfileAudio(u.FileId)
}

// CancelDownloadFile is a helper method for Client.CancelDownloadFile
func (u UpdateFileRemovedFromDownloads) CancelDownloadFile(client *Client, onlyIfPending bool) (*Ok, error) {
	return client.CancelDownloadFile(u.FileId, onlyIfPending)
}

// CancelPreliminaryUploadFile is a helper method for Client.CancelPreliminaryUploadFile
func (u UpdateFileRemovedFromDownloads) CancelPreliminaryUploadFile(client *Client) (*Ok, error) {
	return client.CancelPreliminaryUploadFile(u.FileId)
}

// DeleteFile is a helper method for Client.DeleteFile
func (u UpdateFileRemovedFromDownloads) DeleteFile(client *Client) (*Ok, error) {
	return client.DeleteFile(u.FileId)
}

// DownloadFile is a helper method for Client.DownloadFile
func (u UpdateFileRemovedFromDownloads) DownloadFile(client *Client, priority int32, offset int64, limit int64, synchronous bool) (*File, error) {
	return client.DownloadFile(u.FileId, priority, offset, limit, synchronous)
}

// EditBotMediaPreview is a helper method for Client.EditBotMediaPreview
func (u UpdateFileRemovedFromDownloads) EditBotMediaPreview(client *Client, botUserId int64, languageCode string, content *InputStoryContent) (*BotMediaPreview, error) {
	return client.EditBotMediaPreview(botUserId, languageCode, u.FileId, content)
}

// GetAttachedStickerSets is a helper method for Client.GetAttachedStickerSets
func (u UpdateFileRemovedFromDownloads) GetAttachedStickerSets(client *Client) (*StickerSets, error) {
	return client.GetAttachedStickerSets(u.FileId)
}

// GetFile is a helper method for Client.GetFile
func (u UpdateFileRemovedFromDownloads) GetFile(client *Client) (*File, error) {
	return client.GetFile(u.FileId)
}

// GetFileDownloadedPrefixSize is a helper method for Client.GetFileDownloadedPrefixSize
func (u UpdateFileRemovedFromDownloads) GetFileDownloadedPrefixSize(client *Client, offset int64) (*FileDownloadedPrefixSize, error) {
	return client.GetFileDownloadedPrefixSize(u.FileId, offset)
}

// GetSuggestedFileName is a helper method for Client.GetSuggestedFileName
func (u UpdateFileRemovedFromDownloads) GetSuggestedFileName(client *Client, directory string) (*Text, error) {
	return client.GetSuggestedFileName(u.FileId, directory)
}

// IsProfileAudio is a helper method for Client.IsProfileAudio
func (u UpdateFileRemovedFromDownloads) IsProfileAudio(client *Client) (*Ok, error) {
	return client.IsProfileAudio(u.FileId)
}

// ReadFilePart is a helper method for Client.ReadFilePart
func (u UpdateFileRemovedFromDownloads) ReadFilePart(client *Client, offset int64, count int64) (*Data, error) {
	return client.ReadFilePart(u.FileId, offset, count)
}

// RemoveFileFromDownloads is a helper method for Client.RemoveFileFromDownloads
func (u UpdateFileRemovedFromDownloads) RemoveFileFromDownloads(client *Client, deleteFromCache bool) (*Ok, error) {
	return client.RemoveFileFromDownloads(u.FileId, deleteFromCache)
}

// RemoveProfileAudio is a helper method for Client.RemoveProfileAudio
func (u UpdateFileRemovedFromDownloads) RemoveProfileAudio(client *Client) (*Ok, error) {
	return client.RemoveProfileAudio(u.FileId)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateFileRemovedFromDownloads) ReportChatPhoto(client *Client, chatId int64, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(chatId, u.FileId, reason, text)
}

// SetProfileAudioPosition is a helper method for Client.SetProfileAudioPosition
func (u UpdateFileRemovedFromDownloads) SetProfileAudioPosition(client *Client, afterFileId int32) (*Ok, error) {
	return client.SetProfileAudioPosition(u.FileId, afterFileId)
}

// ToggleDownloadIsPaused is a helper method for Client.ToggleDownloadIsPaused
func (u UpdateFileRemovedFromDownloads) ToggleDownloadIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleDownloadIsPaused(u.FileId, isPaused)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateForumTopic) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateForumTopic) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateForumTopic) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateForumTopic) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateForumTopic) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateForumTopic) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateForumTopic) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateForumTopic) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateForumTopic) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateForumTopic) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateForumTopic) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateForumTopic) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateForumTopic) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateForumTopic) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateForumTopic) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateForumTopic) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateForumTopic) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateForumTopic) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateForumTopic) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateForumTopic) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateForumTopic) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateForumTopic) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateForumTopic) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateForumTopic) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateForumTopic) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}
