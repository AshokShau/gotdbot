package gotdbot

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateNewInlineQuery) GetChatJoinRequests(client *Client, chatId int64, inviteLink string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(chatId, inviteLink, u.Query, limit, opts)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateNewInlineQuery) GetChatRevenueTransactions(client *Client, chatId int64, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(chatId, u.Offset, limit)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (u UpdateNewInlineQuery) GetChatStoryInteractions(client *Client, storyPosterChatId int64, storyId int32, preferForwards bool, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(storyPosterChatId, storyId, preferForwards, u.Offset, limit, opts)
}

// GetConnectedAffiliatePrograms is a helper method for Client.GetConnectedAffiliatePrograms
func (u UpdateNewInlineQuery) GetConnectedAffiliatePrograms(client *Client, affiliate *AffiliateType, limit int32) (*ConnectedAffiliatePrograms, error) {
	return client.GetConnectedAffiliatePrograms(affiliate, u.Offset, limit)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateNewInlineQuery) GetForumTopics(client *Client, chatId int64, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(chatId, u.Query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGiftChatThemes is a helper method for Client.GetGiftChatThemes
func (u UpdateNewInlineQuery) GetGiftChatThemes(client *Client, limit int32) (*GiftChatThemes, error) {
	return client.GetGiftChatThemes(u.Offset, limit)
}

// GetGrossingWebAppBots is a helper method for Client.GetGrossingWebAppBots
func (u UpdateNewInlineQuery) GetGrossingWebAppBots(client *Client, limit int32) (*FoundUsers, error) {
	return client.GetGrossingWebAppBots(u.Offset, limit)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateNewInlineQuery) GetInlineQueryResults(client *Client, botUserId int64, chatId int64, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, chatId, u.Query, u.Offset, opts)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateNewInlineQuery) GetMessageAddedReactions(client *Client, chatId int64, messageId int64, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, messageId, u.Offset, limit, opts)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateNewInlineQuery) GetMessagePublicForwards(client *Client, chatId int64, messageId int64, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, messageId, u.Offset, limit)
}

// GetPublicPostSearchLimits is a helper method for Client.GetPublicPostSearchLimits
func (u UpdateNewInlineQuery) GetPublicPostSearchLimits(client *Client) (*PublicPostSearchLimits, error) {
	return client.GetPublicPostSearchLimits(u.Query)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (u UpdateNewInlineQuery) GetReceivedGifts(client *Client, businessConnectionId string, ownerId *MessageSender, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, ownerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, u.Offset, limit)
}

// GetSearchSponsoredChats is a helper method for Client.GetSearchSponsoredChats
func (u UpdateNewInlineQuery) GetSearchSponsoredChats(client *Client) (*SponsoredChats, error) {
	return client.GetSearchSponsoredChats(u.Query)
}

// GetStarSubscriptions is a helper method for Client.GetStarSubscriptions
func (u UpdateNewInlineQuery) GetStarSubscriptions(client *Client, onlyExpiring bool) (*StarSubscriptions, error) {
	return client.GetStarSubscriptions(onlyExpiring, u.Offset)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (u UpdateNewInlineQuery) GetStarTransactions(client *Client, ownerId *MessageSender, subscriptionId string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(ownerId, subscriptionId, u.Offset, limit, opts)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateNewInlineQuery) GetStickers(client *Client, stickerType *StickerType, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(stickerType, u.Query, limit, chatId)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (u UpdateNewInlineQuery) GetStoryInteractions(client *Client, storyId int32, onlyContacts bool, preferForwards bool, preferWithReaction bool, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(storyId, u.Query, onlyContacts, preferForwards, preferWithReaction, u.Offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (u UpdateNewInlineQuery) GetStoryPublicForwards(client *Client, storyPosterChatId int64, storyId int32, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(storyPosterChatId, storyId, u.Offset, limit)
}

// GetTonTransactions is a helper method for Client.GetTonTransactions
func (u UpdateNewInlineQuery) GetTonTransactions(client *Client, limit int32, opts *GetTonTransactionsOpts) (*TonTransactions, error) {
	return client.GetTonTransactions(u.Offset, limit, opts)
}

// SearchAffiliatePrograms is a helper method for Client.SearchAffiliatePrograms
func (u UpdateNewInlineQuery) SearchAffiliatePrograms(client *Client, affiliate *AffiliateType, sortOrder *AffiliateProgramSortOrder, limit int32) (*FoundAffiliatePrograms, error) {
	return client.SearchAffiliatePrograms(affiliate, sortOrder, u.Offset, limit)
}

// SearchCallMessages is a helper method for Client.SearchCallMessages
func (u UpdateNewInlineQuery) SearchCallMessages(client *Client, limit int32, onlyMissed bool) (*FoundMessages, error) {
	return client.SearchCallMessages(u.Offset, limit, onlyMissed)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateNewInlineQuery) SearchChatMembers(client *Client, chatId int64, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(chatId, u.Query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateNewInlineQuery) SearchChatMessages(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(chatId, u.Query, fromMessageId, offset, limit, opts)
}

// SearchChats is a helper method for Client.SearchChats
func (u UpdateNewInlineQuery) SearchChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchChats(u.Query, limit)
}

// SearchChatsOnServer is a helper method for Client.SearchChatsOnServer
func (u UpdateNewInlineQuery) SearchChatsOnServer(client *Client, limit int32) (*Chats, error) {
	return client.SearchChatsOnServer(u.Query, limit)
}

// SearchContacts is a helper method for Client.SearchContacts
func (u UpdateNewInlineQuery) SearchContacts(client *Client, limit int32) (*Users, error) {
	return client.SearchContacts(u.Query, limit)
}

// SearchFileDownloads is a helper method for Client.SearchFileDownloads
func (u UpdateNewInlineQuery) SearchFileDownloads(client *Client, onlyActive bool, onlyCompleted bool, limit int32) (*FoundFileDownloads, error) {
	return client.SearchFileDownloads(u.Query, onlyActive, onlyCompleted, u.Offset, limit)
}

// SearchGiftsForResale is a helper method for Client.SearchGiftsForResale
func (u UpdateNewInlineQuery) SearchGiftsForResale(client *Client, giftId string, order *GiftForResaleOrder, attributes []*UpgradedGiftAttributeId, limit int32) (*GiftsForResale, error) {
	return client.SearchGiftsForResale(giftId, order, attributes, u.Offset, limit)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (u UpdateNewInlineQuery) SearchInstalledStickerSets(client *Client, stickerType *StickerType, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(stickerType, u.Query, limit)
}

// SearchMessages is a helper method for Client.SearchMessages
func (u UpdateNewInlineQuery) SearchMessages(client *Client, limit int32, minDate int32, maxDate int32, opts *SearchMessagesOpts) (*FoundMessages, error) {
	return client.SearchMessages(u.Query, u.Offset, limit, minDate, maxDate, opts)
}

// SearchOutgoingDocumentMessages is a helper method for Client.SearchOutgoingDocumentMessages
func (u UpdateNewInlineQuery) SearchOutgoingDocumentMessages(client *Client, limit int32) (*FoundMessages, error) {
	return client.SearchOutgoingDocumentMessages(u.Query, limit)
}

// SearchPublicChats is a helper method for Client.SearchPublicChats
func (u UpdateNewInlineQuery) SearchPublicChats(client *Client) (*Chats, error) {
	return client.SearchPublicChats(u.Query)
}

// SearchPublicMessagesByTag is a helper method for Client.SearchPublicMessagesByTag
func (u UpdateNewInlineQuery) SearchPublicMessagesByTag(client *Client, tag string, limit int32) (*FoundMessages, error) {
	return client.SearchPublicMessagesByTag(tag, u.Offset, limit)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (u UpdateNewInlineQuery) SearchPublicPosts(client *Client, limit int32, starCount int64) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(u.Query, u.Offset, limit, starCount)
}

// SearchPublicStoriesByLocation is a helper method for Client.SearchPublicStoriesByLocation
func (u UpdateNewInlineQuery) SearchPublicStoriesByLocation(client *Client, address *LocationAddress, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByLocation(address, u.Offset, limit)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (u UpdateNewInlineQuery) SearchPublicStoriesByTag(client *Client, storyPosterChatId int64, tag string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(storyPosterChatId, tag, u.Offset, limit)
}

// SearchPublicStoriesByVenue is a helper method for Client.SearchPublicStoriesByVenue
func (u UpdateNewInlineQuery) SearchPublicStoriesByVenue(client *Client, venueProvider string, venueId string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByVenue(venueProvider, venueId, u.Offset, limit)
}

// SearchRecentlyFoundChats is a helper method for Client.SearchRecentlyFoundChats
func (u UpdateNewInlineQuery) SearchRecentlyFoundChats(client *Client, limit int32) (*Chats, error) {
	return client.SearchRecentlyFoundChats(u.Query, limit)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (u UpdateNewInlineQuery) SearchSavedMessages(client *Client, savedMessagesTopicId int64, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(savedMessagesTopicId, u.Query, fromMessageId, offset, limit, opts)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateNewInlineQuery) SearchSecretMessages(client *Client, chatId int64, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(chatId, u.Query, u.Offset, limit, opts)
}

// SearchStickers is a helper method for Client.SearchStickers
func (u UpdateNewInlineQuery) SearchStickers(client *Client, stickerType *StickerType, emojis string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(stickerType, emojis, u.Query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (u UpdateNewInlineQuery) SearchStickerSets(client *Client, stickerType *StickerType) (*StickerSets, error) {
	return client.SearchStickerSets(stickerType, u.Query)
}

// SearchStringsByPrefix is a helper method for Client.SearchStringsByPrefix
func (u UpdateNewInlineQuery) SearchStringsByPrefix(client *Client, strings []string, limit int32, returnNoneForEmptyQuery bool) (*FoundPositions, error) {
	return client.SearchStringsByPrefix(strings, u.Query, limit, returnNoneForEmptyQuery)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (u UpdateNewPreCheckoutQuery) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(u.Currency, amount)
}

// SendPaymentForm is a helper method for Client.SendPaymentForm
func (u UpdateNewPreCheckoutQuery) SendPaymentForm(client *Client, inputInvoice *InputInvoice, paymentFormId string, orderInfoId string, tipAmount int64, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	return client.SendPaymentForm(inputInvoice, paymentFormId, orderInfoId, u.ShippingOptionId, tipAmount, opts)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (u UpdateNewPreCheckoutQuery) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, u.Currency, amount)
}

// RemoveNotification is a helper method for Client.RemoveNotification
func (u UpdateNotification) RemoveNotification(client *Client, notificationId int32) (*Ok, error) {
	return client.RemoveNotification(u.NotificationGroupId, notificationId)
}

// RemoveNotificationGroup is a helper method for Client.RemoveNotificationGroup
func (u UpdateNotification) RemoveNotificationGroup(client *Client, maxNotificationId int32) (*Ok, error) {
	return client.RemoveNotificationGroup(u.NotificationGroupId, maxNotificationId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateNotificationGroup) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateNotificationGroup) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateNotificationGroup) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateNotificationGroup) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateNotificationGroup) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateNotificationGroup) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateNotificationGroup) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateNotificationGroup) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateNotificationGroup) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateNotificationGroup) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateNotificationGroup) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateNotificationGroup) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateNotificationGroup) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateNotificationGroup) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateNotificationGroup) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateNotificationGroup) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateNotificationGroup) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateNotificationGroup) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateNotificationGroup) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateNotificationGroup) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateNotificationGroup) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateNotificationGroup) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateNotificationGroup) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateNotificationGroup) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateNotificationGroup) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateNotificationGroup) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateNotificationGroup) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateNotificationGroup) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateNotificationGroup) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateNotificationGroup) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateNotificationGroup) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateNotificationGroup) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateNotificationGroup) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateNotificationGroup) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateNotificationGroup) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateNotificationGroup) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateNotificationGroup) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateNotificationGroup) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateNotificationGroup) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateNotificationGroup) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateNotificationGroup) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateNotificationGroup) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateNotificationGroup) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateNotificationGroup) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateNotificationGroup) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateNotificationGroup) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateNotificationGroup) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateNotificationGroup) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateNotificationGroup) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateNotificationGroup) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateNotificationGroup) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateNotificationGroup) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateNotificationGroup) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateNotificationGroup) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateNotificationGroup) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateNotificationGroup) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateNotificationGroup) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateNotificationGroup) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateNotificationGroup) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateNotificationGroup) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateNotificationGroup) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateNotificationGroup) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateNotificationGroup) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateNotificationGroup) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateNotificationGroup) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateNotificationGroup) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateNotificationGroup) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateNotificationGroup) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateNotificationGroup) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateNotificationGroup) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateNotificationGroup) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateNotificationGroup) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateNotificationGroup) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateNotificationGroup) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateNotificationGroup) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateNotificationGroup) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateNotificationGroup) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateNotificationGroup) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateNotificationGroup) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateNotificationGroup) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateNotificationGroup) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateNotificationGroup) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateNotificationGroup) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateNotificationGroup) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateNotificationGroup) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateNotificationGroup) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateNotificationGroup) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateNotificationGroup) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateNotificationGroup) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateNotificationGroup) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateNotificationGroup) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateNotificationGroup) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateNotificationGroup) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateNotificationGroup) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateNotificationGroup) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateNotificationGroup) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateNotificationGroup) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateNotificationGroup) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateNotificationGroup) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateNotificationGroup) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateNotificationGroup) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateNotificationGroup) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateNotificationGroup) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateNotificationGroup) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateNotificationGroup) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateNotificationGroup) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateNotificationGroup) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateNotificationGroup) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateNotificationGroup) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateNotificationGroup) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateNotificationGroup) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateNotificationGroup) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateNotificationGroup) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateNotificationGroup) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateNotificationGroup) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateNotificationGroup) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateNotificationGroup) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateNotificationGroup) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateNotificationGroup) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateNotificationGroup) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateNotificationGroup) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateNotificationGroup) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateNotificationGroup) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateNotificationGroup) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateNotificationGroup) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateNotificationGroup) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateNotificationGroup) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateNotificationGroup) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetSavedNotificationSound is a helper method for Client.GetSavedNotificationSound
func (u UpdateNotificationGroup) GetSavedNotificationSound(client *Client) (*NotificationSounds, error) {
	return client.GetSavedNotificationSound(u.NotificationSoundId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateNotificationGroup) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateNotificationGroup) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateNotificationGroup) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateNotificationGroup) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateNotificationGroup) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateNotificationGroup) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateNotificationGroup) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateNotificationGroup) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateNotificationGroup) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateNotificationGroup) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateNotificationGroup) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateNotificationGroup) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateNotificationGroup) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateNotificationGroup) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateNotificationGroup) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateNotificationGroup) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateNotificationGroup) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateNotificationGroup) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateNotificationGroup) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateNotificationGroup) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateNotificationGroup) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateNotificationGroup) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateNotificationGroup) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateNotificationGroup) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateNotificationGroup) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateNotificationGroup) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateNotificationGroup) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateNotificationGroup) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateNotificationGroup) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateNotificationGroup) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateNotificationGroup) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateNotificationGroup) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateNotificationGroup) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemoveNotification is a helper method for Client.RemoveNotification
func (u UpdateNotificationGroup) RemoveNotification(client *Client, notificationId int32) (*Ok, error) {
	return client.RemoveNotification(u.NotificationGroupId, notificationId)
}

// RemoveNotificationGroup is a helper method for Client.RemoveNotificationGroup
func (u UpdateNotificationGroup) RemoveNotificationGroup(client *Client, maxNotificationId int32) (*Ok, error) {
	return client.RemoveNotificationGroup(u.NotificationGroupId, maxNotificationId)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateNotificationGroup) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateNotificationGroup) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveSavedNotificationSound is a helper method for Client.RemoveSavedNotificationSound
func (u UpdateNotificationGroup) RemoveSavedNotificationSound(client *Client) (*Ok, error) {
	return client.RemoveSavedNotificationSound(u.NotificationSoundId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateNotificationGroup) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateNotificationGroup) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateNotificationGroup) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateNotificationGroup) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateNotificationGroup) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateNotificationGroup) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateNotificationGroup) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateNotificationGroup) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateNotificationGroup) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateNotificationGroup) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateNotificationGroup) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateNotificationGroup) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateNotificationGroup) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateNotificationGroup) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateNotificationGroup) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateNotificationGroup) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateNotificationGroup) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateNotificationGroup) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateNotificationGroup) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateNotificationGroup) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateNotificationGroup) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateNotificationGroup) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateNotificationGroup) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateNotificationGroup) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateNotificationGroup) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateNotificationGroup) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateNotificationGroup) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateNotificationGroup) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateNotificationGroup) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateNotificationGroup) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateNotificationGroup) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateNotificationGroup) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateNotificationGroup) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateNotificationGroup) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateNotificationGroup) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateNotificationGroup) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateNotificationGroup) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateNotificationGroup) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateNotificationGroup) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateNotificationGroup) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateNotificationGroup) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateNotificationGroup) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateNotificationGroup) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateNotificationGroup) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateNotificationGroup) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateNotificationGroup) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateNotificationGroup) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateNotificationGroup) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateNotificationGroup) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateNotificationGroup) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateNotificationGroup) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateNotificationGroup) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateNotificationGroup) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateNotificationGroup) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateNotificationGroup) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateNotificationGroup) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateNotificationGroup) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateNotificationGroup) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateNotificationGroup) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateNotificationGroup) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateNotificationGroup) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateNotificationGroup) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateNotificationGroup) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateNotificationGroup) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateNotificationGroup) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateNotificationGroup) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateNotificationGroup) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateNotificationGroup) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateNotificationGroup) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateNotificationGroup) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateNotificationGroup) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateNotificationGroup) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateNotificationGroup) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateNotificationGroup) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateNotificationGroup) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateNotificationGroup) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateNotificationGroup) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateNotificationGroup) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateNotificationGroup) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateNotificationGroup) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateNotificationGroup) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateNotificationGroup) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateNotificationGroup) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateNotificationGroup) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateNotificationGroup) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateNotificationGroup) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateNotificationGroup) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateNotificationGroup) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateNotificationGroup) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpdateOption) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, u.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (u UpdateOption) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(u.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (u UpdateOption) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(u.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UpdateOption) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, u.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateOption) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateOption) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, u.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateOption) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, u.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpdateOption) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, u.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdateOption) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, u.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpdateOption) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, u.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (u UpdateOption) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(u.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UpdateOption) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, u.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateOption) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateOption) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, u.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateOption) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, u.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (u UpdateOption) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(u.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (u UpdateOption) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(u.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (u UpdateOption) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(u.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (u UpdateOption) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(u.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpdateOption) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, u.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (u UpdateOption) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(u.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (u UpdateOption) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(u.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (u UpdateOption) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, u.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (u UpdateOption) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(u.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpdateOption) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, u.Name)
}

// SetOption is a helper method for Client.SetOption
func (u UpdateOption) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(u.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpdateOption) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, u.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpdateOption) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, u.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpdateOption) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(u.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateOption) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, u.Name)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdatePaidMediaPurchased) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, u.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (u UpdatePaidMediaPurchased) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(u.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpdatePaidMediaPurchased) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(u.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (u UpdatePaidMediaPurchased) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(u.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (u UpdatePaidMediaPurchased) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(u.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (u UpdatePaidMediaPurchased) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(u.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (u UpdatePaidMediaPurchased) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(u.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdatePaidMediaPurchased) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(u.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (u UpdatePaidMediaPurchased) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(u.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (u UpdatePaidMediaPurchased) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(u.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdatePaidMediaPurchased) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, u.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (u UpdatePaidMediaPurchased) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(u.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (u UpdatePaidMediaPurchased) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, u.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (u UpdatePaidMediaPurchased) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(u.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (u UpdatePaidMediaPurchased) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(u.UserId)
}

// GetPushReceiverId is a helper method for Client.GetPushReceiverId
func (u UpdatePaidMediaPurchased) GetPushReceiverId(client *Client) (*PushReceiverId, error) {
	return client.GetPushReceiverId(u.Payload)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (u UpdatePaidMediaPurchased) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(u.UserId)
}

// GetUser is a helper method for Client.GetUser
func (u UpdatePaidMediaPurchased) GetUser(client *Client) (*User, error) {
	return client.GetUser(u.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdatePaidMediaPurchased) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, u.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (u UpdatePaidMediaPurchased) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(u.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (u UpdatePaidMediaPurchased) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(u.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (u UpdatePaidMediaPurchased) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(u.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (u UpdatePaidMediaPurchased) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(u.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UpdatePaidMediaPurchased) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(u.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdatePaidMediaPurchased) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, u.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UpdatePaidMediaPurchased) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, u.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdatePaidMediaPurchased) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, u.UserId, approve)
}

// ProcessPushNotification is a helper method for Client.ProcessPushNotification
func (u UpdatePaidMediaPurchased) ProcessPushNotification(client *Client) (*Ok, error) {
	return client.ProcessPushNotification(u.Payload)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (u UpdatePaidMediaPurchased) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(u.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpdatePaidMediaPurchased) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(u.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (u UpdatePaidMediaPurchased) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(u.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdatePaidMediaPurchased) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, u.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (u UpdatePaidMediaPurchased) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, u.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (u UpdatePaidMediaPurchased) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(u.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (u UpdatePaidMediaPurchased) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(u.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpdatePaidMediaPurchased) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(u.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (u UpdatePaidMediaPurchased) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(u.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (u UpdatePaidMediaPurchased) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(u.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (u UpdatePaidMediaPurchased) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(u.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (u UpdatePaidMediaPurchased) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(u.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (u UpdatePaidMediaPurchased) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(u.UserId)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (u UpdatePaidMediaPurchased) StartGroupCallScreenSharing(client *Client, groupCallId int32, audioSourceId int32) (*Text, error) {
	return client.StartGroupCallScreenSharing(groupCallId, audioSourceId, u.Payload)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (u UpdatePaidMediaPurchased) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(u.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (u UpdatePaidMediaPurchased) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(u.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdatePaidMediaPurchased) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, u.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (u UpdatePaidMediaPurchased) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(u.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdatePendingTextMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdatePendingTextMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdatePendingTextMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdatePendingTextMessage) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdatePendingTextMessage) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdatePendingTextMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdatePendingTextMessage) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdatePendingTextMessage) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdatePendingTextMessage) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdatePendingTextMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdatePendingTextMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdatePendingTextMessage) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdatePendingTextMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdatePendingTextMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdatePendingTextMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdatePendingTextMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdatePendingTextMessage) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdatePendingTextMessage) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdatePendingTextMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdatePendingTextMessage) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdatePendingTextMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdatePendingTextMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdatePendingTextMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdatePendingTextMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdatePendingTextMessage) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdatePendingTextMessage) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdatePendingTextMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdatePendingTextMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdatePendingTextMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdatePendingTextMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdatePendingTextMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdatePendingTextMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdatePendingTextMessage) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdatePendingTextMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdatePendingTextMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdatePendingTextMessage) DeleteForumTopic(client *Client) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, u.ForumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdatePendingTextMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdatePendingTextMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdatePendingTextMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdatePendingTextMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdatePendingTextMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdatePendingTextMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdatePendingTextMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdatePendingTextMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdatePendingTextMessage) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdatePendingTextMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdatePendingTextMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdatePendingTextMessage) EditForumTopic(client *Client, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, u.ForumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdatePendingTextMessage) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdatePendingTextMessage) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdatePendingTextMessage) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdatePendingTextMessage) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdatePendingTextMessage) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdatePendingTextMessage) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdatePendingTextMessage) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdatePendingTextMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdatePendingTextMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdatePendingTextMessage) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdatePendingTextMessage) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdatePendingTextMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdatePendingTextMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdatePendingTextMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdatePendingTextMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdatePendingTextMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdatePendingTextMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdatePendingTextMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdatePendingTextMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdatePendingTextMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdatePendingTextMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdatePendingTextMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdatePendingTextMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdatePendingTextMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdatePendingTextMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdatePendingTextMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdatePendingTextMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdatePendingTextMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdatePendingTextMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdatePendingTextMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdatePendingTextMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdatePendingTextMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdatePendingTextMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdatePendingTextMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdatePendingTextMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdatePendingTextMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdatePendingTextMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdatePendingTextMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdatePendingTextMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdatePendingTextMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdatePendingTextMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdatePendingTextMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdatePendingTextMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdatePendingTextMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdatePendingTextMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdatePendingTextMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdatePendingTextMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdatePendingTextMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdatePendingTextMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdatePendingTextMessage) GetForumTopic(client *Client) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, u.ForumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdatePendingTextMessage) GetForumTopicHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, u.ForumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdatePendingTextMessage) GetForumTopicLink(client *Client) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, u.ForumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdatePendingTextMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdatePendingTextMessage) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdatePendingTextMessage) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdatePendingTextMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (u UpdatePendingTextMessage) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(u.Text, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdatePendingTextMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdatePendingTextMessage) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdatePendingTextMessage) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdatePendingTextMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdatePendingTextMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (u UpdatePendingTextMessage) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(u.Text)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdatePendingTextMessage) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdatePendingTextMessage) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdatePendingTextMessage) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdatePendingTextMessage) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdatePendingTextMessage) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdatePendingTextMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdatePendingTextMessage) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdatePendingTextMessage) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdatePendingTextMessage) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdatePendingTextMessage) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdatePendingTextMessage) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdatePendingTextMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdatePendingTextMessage) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdatePendingTextMessage) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdatePendingTextMessage) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdatePendingTextMessage) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdatePendingTextMessage) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdatePendingTextMessage) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdatePendingTextMessage) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdatePendingTextMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdatePendingTextMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdatePendingTextMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdatePendingTextMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdatePendingTextMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdatePendingTextMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdatePendingTextMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdatePendingTextMessage) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdatePendingTextMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UpdatePendingTextMessage) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, u.Text)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdatePendingTextMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdatePendingTextMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdatePendingTextMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdatePendingTextMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdatePendingTextMessage) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdatePendingTextMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdatePendingTextMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdatePendingTextMessage) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdatePendingTextMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (u UpdatePendingTextMessage) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(u.Text)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdatePendingTextMessage) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UpdatePendingTextMessage) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, u.Text, isPrivate)
}

// PostStory is a helper method for Client.PostStory
func (u UpdatePendingTextMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdatePendingTextMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdatePendingTextMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdatePendingTextMessage) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdatePendingTextMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdatePendingTextMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdatePendingTextMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdatePendingTextMessage) ReadAllForumTopicMentions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, u.ForumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdatePendingTextMessage) ReadAllForumTopicReactions(client *Client) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, u.ForumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdatePendingTextMessage) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdatePendingTextMessage) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdatePendingTextMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdatePendingTextMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdatePendingTextMessage) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdatePendingTextMessage) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdatePendingTextMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdatePendingTextMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdatePendingTextMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdatePendingTextMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdatePendingTextMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdatePendingTextMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdatePendingTextMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdatePendingTextMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdatePendingTextMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdatePendingTextMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdatePendingTextMessage) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdatePendingTextMessage) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdatePendingTextMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdatePendingTextMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdatePendingTextMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdatePendingTextMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdatePendingTextMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdatePendingTextMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchQuote is a helper method for Client.SearchQuote
func (u UpdatePendingTextMessage) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(u.Text, quote, quotePosition)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdatePendingTextMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdatePendingTextMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdatePendingTextMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdatePendingTextMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdatePendingTextMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendGift is a helper method for Client.SendGift
func (u UpdatePendingTextMessage) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, u.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpdatePendingTextMessage) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, u.Text, paidMessageStarCount)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdatePendingTextMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdatePendingTextMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdatePendingTextMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdatePendingTextMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdatePendingTextMessage) SendTextMessageDraft(client *Client) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, u.ForumTopicId, u.DraftId, u.Text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdatePendingTextMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdatePendingTextMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdatePendingTextMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdatePendingTextMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdatePendingTextMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdatePendingTextMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdatePendingTextMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdatePendingTextMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdatePendingTextMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdatePendingTextMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdatePendingTextMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdatePendingTextMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdatePendingTextMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdatePendingTextMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdatePendingTextMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdatePendingTextMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdatePendingTextMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdatePendingTextMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdatePendingTextMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdatePendingTextMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdatePendingTextMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdatePendingTextMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdatePendingTextMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdatePendingTextMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdatePendingTextMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdatePendingTextMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdatePendingTextMessage) SetForumTopicNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, u.ForumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdatePendingTextMessage) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdatePendingTextMessage) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdatePendingTextMessage) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdatePendingTextMessage) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdatePendingTextMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdatePendingTextMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdatePendingTextMessage) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdatePendingTextMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdatePendingTextMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdatePendingTextMessage) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdatePendingTextMessage) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdatePendingTextMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdatePendingTextMessage) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdatePendingTextMessage) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdatePendingTextMessage) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdatePendingTextMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdatePendingTextMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdatePendingTextMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdatePendingTextMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdatePendingTextMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdatePendingTextMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdatePendingTextMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdatePendingTextMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdatePendingTextMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdatePendingTextMessage) ToggleForumTopicIsClosed(client *Client, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, u.ForumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdatePendingTextMessage) ToggleForumTopicIsPinned(client *Client, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, u.ForumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdatePendingTextMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdatePendingTextMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdatePendingTextMessage) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// TranslateText is a helper method for Client.TranslateText
func (u UpdatePendingTextMessage) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(u.Text, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdatePendingTextMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdatePendingTextMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdatePendingTextMessage) UnpinAllForumTopicMessages(client *Client) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, u.ForumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdatePendingTextMessage) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdatePendingTextMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdatePendingTextMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdatePollAnswer) SetPollAnswer(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.SetPollAnswer(chatId, messageId, u.OptionIds)
}

// DeleteQuickReplyShortcut is a helper method for Client.DeleteQuickReplyShortcut
func (u UpdateQuickReplyShortcutDeleted) DeleteQuickReplyShortcut(client *Client) (*Ok, error) {
	return client.DeleteQuickReplyShortcut(u.ShortcutId)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutDeleted) DeleteQuickReplyShortcutMessages(client *Client, messageIds []int64) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(u.ShortcutId, messageIds)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateQuickReplyShortcutDeleted) EditQuickReplyMessage(client *Client, messageId int64, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(u.ShortcutId, messageId, inputMessageContent)
}

// LoadQuickReplyShortcutMessages is a helper method for Client.LoadQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutDeleted) LoadQuickReplyShortcutMessages(client *Client) (*Ok, error) {
	return client.LoadQuickReplyShortcutMessages(u.ShortcutId)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutDeleted) SendQuickReplyShortcutMessages(client *Client, chatId int64, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, u.ShortcutId, sendingId)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpdateQuickReplyShortcutDeleted) SetQuickReplyShortcutName(client *Client, name string) (*Ok, error) {
	return client.SetQuickReplyShortcutName(u.ShortcutId, name)
}

// DeleteQuickReplyShortcut is a helper method for Client.DeleteQuickReplyShortcut
func (u UpdateQuickReplyShortcutMessages) DeleteQuickReplyShortcut(client *Client) (*Ok, error) {
	return client.DeleteQuickReplyShortcut(u.ShortcutId)
}

// DeleteQuickReplyShortcutMessages is a helper method for Client.DeleteQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutMessages) DeleteQuickReplyShortcutMessages(client *Client, messageIds []int64) (*Ok, error) {
	return client.DeleteQuickReplyShortcutMessages(u.ShortcutId, messageIds)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateQuickReplyShortcutMessages) EditQuickReplyMessage(client *Client, messageId int64, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(u.ShortcutId, messageId, inputMessageContent)
}

// LoadQuickReplyShortcutMessages is a helper method for Client.LoadQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutMessages) LoadQuickReplyShortcutMessages(client *Client) (*Ok, error) {
	return client.LoadQuickReplyShortcutMessages(u.ShortcutId)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateQuickReplyShortcutMessages) SendQuickReplyShortcutMessages(client *Client, chatId int64, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, u.ShortcutId, sendingId)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpdateQuickReplyShortcutMessages) SetQuickReplyShortcutName(client *Client, name string) (*Ok, error) {
	return client.SetQuickReplyShortcutName(u.ShortcutId, name)
}

// ReorderQuickReplyShortcuts is a helper method for Client.ReorderQuickReplyShortcuts
func (u UpdateQuickReplyShortcuts) ReorderQuickReplyShortcuts(client *Client) (*Ok, error) {
	return client.ReorderQuickReplyShortcuts(u.ShortcutIds)
}

// SetReactionNotificationSettings is a helper method for Client.SetReactionNotificationSettings
func (u UpdateReactionNotificationSettings) SetReactionNotificationSettings(client *Client) (*Ok, error) {
	return client.SetReactionNotificationSettings(u.NotificationSettings)
}

// AddRecentSticker is a helper method for Client.AddRecentSticker
func (u UpdateRecentStickers) AddRecentSticker(client *Client, sticker *InputFile) (*Stickers, error) {
	return client.AddRecentSticker(u.IsAttached, sticker)
}

// ClearRecentStickers is a helper method for Client.ClearRecentStickers
func (u UpdateRecentStickers) ClearRecentStickers(client *Client) (*Ok, error) {
	return client.ClearRecentStickers(u.IsAttached)
}

// GetRecentStickers is a helper method for Client.GetRecentStickers
func (u UpdateRecentStickers) GetRecentStickers(client *Client) (*Stickers, error) {
	return client.GetRecentStickers(u.IsAttached)
}

// RemoveRecentSticker is a helper method for Client.RemoveRecentSticker
func (u UpdateRecentStickers) RemoveRecentSticker(client *Client, sticker *InputFile) (*Ok, error) {
	return client.RemoveRecentSticker(u.IsAttached, sticker)
}

// DeleteSavedMessagesTopicHistory is a helper method for Client.DeleteSavedMessagesTopicHistory
func (u UpdateSavedMessagesTags) DeleteSavedMessagesTopicHistory(client *Client) (*Ok, error) {
	return client.DeleteSavedMessagesTopicHistory(u.SavedMessagesTopicId)
}

// DeleteSavedMessagesTopicMessagesByDate is a helper method for Client.DeleteSavedMessagesTopicMessagesByDate
func (u UpdateSavedMessagesTags) DeleteSavedMessagesTopicMessagesByDate(client *Client, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteSavedMessagesTopicMessagesByDate(u.SavedMessagesTopicId, minDate, maxDate)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateSavedMessagesTags) GetChatSparseMessagePositions(client *Client, chatId int64, filter *SearchMessagesFilter, fromMessageId int64, limit int32) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(chatId, filter, fromMessageId, limit, u.SavedMessagesTopicId)
}

// GetSavedMessagesTags is a helper method for Client.GetSavedMessagesTags
func (u UpdateSavedMessagesTags) GetSavedMessagesTags(client *Client) (*SavedMessagesTags, error) {
	return client.GetSavedMessagesTags(u.SavedMessagesTopicId)
}

// GetSavedMessagesTopicHistory is a helper method for Client.GetSavedMessagesTopicHistory
func (u UpdateSavedMessagesTags) GetSavedMessagesTopicHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetSavedMessagesTopicHistory(u.SavedMessagesTopicId, fromMessageId, offset, limit)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (u UpdateSavedMessagesTags) GetSavedMessagesTopicMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(u.SavedMessagesTopicId, date)
}

// SearchSavedMessages is a helper method for Client.SearchSavedMessages
func (u UpdateSavedMessagesTags) SearchSavedMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchSavedMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchSavedMessages(u.SavedMessagesTopicId, query, fromMessageId, offset, limit, opts)
}

// ToggleSavedMessagesTopicIsPinned is a helper method for Client.ToggleSavedMessagesTopicIsPinned
func (u UpdateSavedMessagesTags) ToggleSavedMessagesTopicIsPinned(client *Client, isPinned bool) (*Ok, error) {
	return client.ToggleSavedMessagesTopicIsPinned(u.SavedMessagesTopicId, isPinned)
}

// GetScopeNotificationSettings is a helper method for Client.GetScopeNotificationSettings
func (u UpdateScopeNotificationSettings) GetScopeNotificationSettings(client *Client) (*ScopeNotificationSettings, error) {
	return client.GetScopeNotificationSettings(u.Scope)
}

// SetScopeNotificationSettings is a helper method for Client.SetScopeNotificationSettings
func (u UpdateScopeNotificationSettings) SetScopeNotificationSettings(client *Client) (*Ok, error) {
	return client.SetScopeNotificationSettings(u.Scope, u.NotificationSettings)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateServiceNotification) SaveApplicationLogEvent(client *Client, chatId int64, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(u.Type, chatId, data)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (u UpdateStarRevenueStatus) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (u UpdateStarRevenueStatus) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(u.OwnerId, prepaidUpgradeHash, starCount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpdateStarRevenueStatus) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(u.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (u UpdateStarRevenueStatus) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(u.OwnerId, collectionId)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (u UpdateStarRevenueStatus) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(u.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (u UpdateStarRevenueStatus) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, u.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (u UpdateStarRevenueStatus) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(u.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (u UpdateStarRevenueStatus) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(u.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (u UpdateStarRevenueStatus) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(u.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (u UpdateStarRevenueStatus) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(u.OwnerId, starCount, password)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (u UpdateStarRevenueStatus) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (u UpdateStarRevenueStatus) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (u UpdateStarRevenueStatus) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(u.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (u UpdateStarRevenueStatus) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, u.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u UpdateStarRevenueStatus) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(u.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (u UpdateStarRevenueStatus) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, u.OwnerId, price)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpdateStarRevenueStatus) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(u.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (u UpdateStarRevenueStatus) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(u.OwnerId, receivedGiftIds)
}

// CloseStory is a helper method for Client.CloseStory
func (u UpdateStoryDeleted) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(u.StoryPosterChatId, u.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpdateStoryDeleted) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(u.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (u UpdateStoryDeleted) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, u.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (u UpdateStoryDeleted) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(u.StoryPosterChatId, u.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (u UpdateStoryDeleted) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(u.StoryPosterChatId, u.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (u UpdateStoryDeleted) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(u.StoryPosterChatId, u.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (u UpdateStoryDeleted) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(u.StoryPosterChatId, u.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (u UpdateStoryDeleted) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(u.StoryPosterChatId, u.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (u UpdateStoryDeleted) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(u.StoryPosterChatId, u.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (u UpdateStoryDeleted) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(u.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (u UpdateStoryDeleted) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(u.StoryPosterChatId, u.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateStoryDeleted) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, u.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (u UpdateStoryDeleted) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(u.StoryPosterChatId, u.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (u UpdateStoryDeleted) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(u.StoryPosterChatId, u.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (u UpdateStoryDeleted) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(u.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (u UpdateStoryDeleted) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(u.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (u UpdateStoryDeleted) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(u.StoryPosterChatId, u.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (u UpdateStoryDeleted) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(u.StoryPosterChatId, u.StoryId, isPostedToChatPage)
}

// LoadActiveStories is a helper method for Client.LoadActiveStories
func (u UpdateStoryListChatCount) LoadActiveStories(client *Client) (*Ok, error) {
	return client.LoadActiveStories(u.StoryList)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateStoryListChatCount) SetChatActiveStoriesList(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatActiveStoriesList(chatId, u.StoryList)
}

// TestReturnError is a helper method for Client.TestReturnError
func (u UpdateStoryPostFailed) TestReturnError(client *Client) (*Error, error) {
	return client.TestReturnError(u.Error)
}

// CreateSupergroupChat is a helper method for Client.CreateSupergroupChat
func (u UpdateSupergroupFullInfo) CreateSupergroupChat(client *Client, force bool) (*Chat, error) {
	return client.CreateSupergroupChat(u.SupergroupId, force)
}

// DisableAllSupergroupUsernames is a helper method for Client.DisableAllSupergroupUsernames
func (u UpdateSupergroupFullInfo) DisableAllSupergroupUsernames(client *Client) (*Ok, error) {
	return client.DisableAllSupergroupUsernames(u.SupergroupId)
}

// GetSupergroup is a helper method for Client.GetSupergroup
func (u UpdateSupergroupFullInfo) GetSupergroup(client *Client) (*Supergroup, error) {
	return client.GetSupergroup(u.SupergroupId)
}

// GetSupergroupFullInfo is a helper method for Client.GetSupergroupFullInfo
func (u UpdateSupergroupFullInfo) GetSupergroupFullInfo(client *Client) (*SupergroupFullInfo, error) {
	return client.GetSupergroupFullInfo(u.SupergroupId)
}

// GetSupergroupMembers is a helper method for Client.GetSupergroupMembers
func (u UpdateSupergroupFullInfo) GetSupergroupMembers(client *Client, offset int32, limit int32, opts *GetSupergroupMembersOpts) (*ChatMembers, error) {
	return client.GetSupergroupMembers(u.SupergroupId, offset, limit, opts)
}

// ReorderSupergroupActiveUsernames is a helper method for Client.ReorderSupergroupActiveUsernames
func (u UpdateSupergroupFullInfo) ReorderSupergroupActiveUsernames(client *Client, usernames []string) (*Ok, error) {
	return client.ReorderSupergroupActiveUsernames(u.SupergroupId, usernames)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateSupergroupFullInfo) ReportSupergroupAntiSpamFalsePositive(client *Client, messageId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(u.SupergroupId, messageId)
}

// ReportSupergroupSpam is a helper method for Client.ReportSupergroupSpam
func (u UpdateSupergroupFullInfo) ReportSupergroupSpam(client *Client, messageIds []int64) (*Ok, error) {
	return client.ReportSupergroupSpam(u.SupergroupId, messageIds)
}

// SetSupergroupCustomEmojiStickerSet is a helper method for Client.SetSupergroupCustomEmojiStickerSet
func (u UpdateSupergroupFullInfo) SetSupergroupCustomEmojiStickerSet(client *Client, customEmojiStickerSetId string) (*Ok, error) {
	return client.SetSupergroupCustomEmojiStickerSet(u.SupergroupId, customEmojiStickerSetId)
}

// SetSupergroupMainProfileTab is a helper method for Client.SetSupergroupMainProfileTab
func (u UpdateSupergroupFullInfo) SetSupergroupMainProfileTab(client *Client, mainProfileTab *ProfileTab) (*Ok, error) {
	return client.SetSupergroupMainProfileTab(u.SupergroupId, mainProfileTab)
}

// SetSupergroupStickerSet is a helper method for Client.SetSupergroupStickerSet
func (u UpdateSupergroupFullInfo) SetSupergroupStickerSet(client *Client, stickerSetId string) (*Ok, error) {
	return client.SetSupergroupStickerSet(u.SupergroupId, stickerSetId)
}

// SetSupergroupUnrestrictBoostCount is a helper method for Client.SetSupergroupUnrestrictBoostCount
func (u UpdateSupergroupFullInfo) SetSupergroupUnrestrictBoostCount(client *Client, unrestrictBoostCount int32) (*Ok, error) {
	return client.SetSupergroupUnrestrictBoostCount(u.SupergroupId, unrestrictBoostCount)
}

// SetSupergroupUsername is a helper method for Client.SetSupergroupUsername
func (u UpdateSupergroupFullInfo) SetSupergroupUsername(client *Client, username string) (*Ok, error) {
	return client.SetSupergroupUsername(u.SupergroupId, username)
}

// ToggleSupergroupCanHaveSponsoredMessages is a helper method for Client.ToggleSupergroupCanHaveSponsoredMessages
func (u UpdateSupergroupFullInfo) ToggleSupergroupCanHaveSponsoredMessages(client *Client, canHaveSponsoredMessages bool) (*Ok, error) {
	return client.ToggleSupergroupCanHaveSponsoredMessages(u.SupergroupId, canHaveSponsoredMessages)
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled is a helper method for Client.ToggleSupergroupHasAggressiveAntiSpamEnabled
func (u UpdateSupergroupFullInfo) ToggleSupergroupHasAggressiveAntiSpamEnabled(client *Client, hasAggressiveAntiSpamEnabled bool) (*Ok, error) {
	return client.ToggleSupergroupHasAggressiveAntiSpamEnabled(u.SupergroupId, hasAggressiveAntiSpamEnabled)
}

// ToggleSupergroupHasAutomaticTranslation is a helper method for Client.ToggleSupergroupHasAutomaticTranslation
func (u UpdateSupergroupFullInfo) ToggleSupergroupHasAutomaticTranslation(client *Client, hasAutomaticTranslation bool) (*Ok, error) {
	return client.ToggleSupergroupHasAutomaticTranslation(u.SupergroupId, hasAutomaticTranslation)
}

// ToggleSupergroupHasHiddenMembers is a helper method for Client.ToggleSupergroupHasHiddenMembers
func (u UpdateSupergroupFullInfo) ToggleSupergroupHasHiddenMembers(client *Client, hasHiddenMembers bool) (*Ok, error) {
	return client.ToggleSupergroupHasHiddenMembers(u.SupergroupId, hasHiddenMembers)
}

// ToggleSupergroupIsAllHistoryAvailable is a helper method for Client.ToggleSupergroupIsAllHistoryAvailable
func (u UpdateSupergroupFullInfo) ToggleSupergroupIsAllHistoryAvailable(client *Client, isAllHistoryAvailable bool) (*Ok, error) {
	return client.ToggleSupergroupIsAllHistoryAvailable(u.SupergroupId, isAllHistoryAvailable)
}

// ToggleSupergroupIsBroadcastGroup is a helper method for Client.ToggleSupergroupIsBroadcastGroup
func (u UpdateSupergroupFullInfo) ToggleSupergroupIsBroadcastGroup(client *Client) (*Ok, error) {
	return client.ToggleSupergroupIsBroadcastGroup(u.SupergroupId)
}

// ToggleSupergroupIsForum is a helper method for Client.ToggleSupergroupIsForum
func (u UpdateSupergroupFullInfo) ToggleSupergroupIsForum(client *Client, isForum bool, hasForumTabs bool) (*Ok, error) {
	return client.ToggleSupergroupIsForum(u.SupergroupId, isForum, hasForumTabs)
}

// ToggleSupergroupJoinByRequest is a helper method for Client.ToggleSupergroupJoinByRequest
func (u UpdateSupergroupFullInfo) ToggleSupergroupJoinByRequest(client *Client, joinByRequest bool) (*Ok, error) {
	return client.ToggleSupergroupJoinByRequest(u.SupergroupId, joinByRequest)
}

// ToggleSupergroupJoinToSendMessages is a helper method for Client.ToggleSupergroupJoinToSendMessages
func (u UpdateSupergroupFullInfo) ToggleSupergroupJoinToSendMessages(client *Client, joinToSendMessages bool) (*Ok, error) {
	return client.ToggleSupergroupJoinToSendMessages(u.SupergroupId, joinToSendMessages)
}

// ToggleSupergroupSignMessages is a helper method for Client.ToggleSupergroupSignMessages
func (u UpdateSupergroupFullInfo) ToggleSupergroupSignMessages(client *Client, signMessages bool, showMessageSender bool) (*Ok, error) {
	return client.ToggleSupergroupSignMessages(u.SupergroupId, signMessages, showMessageSender)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (u UpdateSupergroupFullInfo) ToggleSupergroupUsernameIsActive(client *Client, username string, isActive bool) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(u.SupergroupId, username, isActive)
}

// AcceptTermsOfService is a helper method for Client.AcceptTermsOfService
func (u UpdateTermsOfService) AcceptTermsOfService(client *Client) (*Ok, error) {
	return client.AcceptTermsOfService(u.TermsOfServiceId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateTopicMessageCount) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateTopicMessageCount) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateTopicMessageCount) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateTopicMessageCount) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateTopicMessageCount) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateTopicMessageCount) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateTopicMessageCount) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateTopicMessageCount) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateTopicMessageCount) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateTopicMessageCount) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateTopicMessageCount) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateTopicMessageCount) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateTopicMessageCount) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateTopicMessageCount) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateTopicMessageCount) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateTopicMessageCount) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateTopicMessageCount) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateTopicMessageCount) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateTopicMessageCount) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateTopicMessageCount) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateTopicMessageCount) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateTopicMessageCount) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateTopicMessageCount) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateTopicMessageCount) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateTopicMessageCount) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateTopicMessageCount) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateTopicMessageCount) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateTopicMessageCount) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateTopicMessageCount) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateTopicMessageCount) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateTopicMessageCount) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateTopicMessageCount) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateTopicMessageCount) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateTopicMessageCount) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateTopicMessageCount) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateTopicMessageCount) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateTopicMessageCount) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateTopicMessageCount) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateTopicMessageCount) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateTopicMessageCount) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateTopicMessageCount) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateTopicMessageCount) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateTopicMessageCount) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateTopicMessageCount) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateTopicMessageCount) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateTopicMessageCount) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateTopicMessageCount) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateTopicMessageCount) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateTopicMessageCount) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateTopicMessageCount) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateTopicMessageCount) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateTopicMessageCount) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateTopicMessageCount) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateTopicMessageCount) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateTopicMessageCount) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateTopicMessageCount) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateTopicMessageCount) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateTopicMessageCount) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateTopicMessageCount) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateTopicMessageCount) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateTopicMessageCount) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateTopicMessageCount) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateTopicMessageCount) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateTopicMessageCount) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateTopicMessageCount) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateTopicMessageCount) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateTopicMessageCount) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateTopicMessageCount) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateTopicMessageCount) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateTopicMessageCount) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateTopicMessageCount) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateTopicMessageCount) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateTopicMessageCount) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateTopicMessageCount) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateTopicMessageCount) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateTopicMessageCount) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateTopicMessageCount) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateTopicMessageCount) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateTopicMessageCount) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateTopicMessageCount) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateTopicMessageCount) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateTopicMessageCount) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateTopicMessageCount) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateTopicMessageCount) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateTopicMessageCount) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateTopicMessageCount) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateTopicMessageCount) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateTopicMessageCount) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateTopicMessageCount) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateTopicMessageCount) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateTopicMessageCount) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateTopicMessageCount) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateTopicMessageCount) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateTopicMessageCount) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateTopicMessageCount) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateTopicMessageCount) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateTopicMessageCount) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateTopicMessageCount) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateTopicMessageCount) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateTopicMessageCount) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateTopicMessageCount) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateTopicMessageCount) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateTopicMessageCount) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateTopicMessageCount) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateTopicMessageCount) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateTopicMessageCount) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateTopicMessageCount) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateTopicMessageCount) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateTopicMessageCount) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateTopicMessageCount) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(u.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateTopicMessageCount) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateTopicMessageCount) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateTopicMessageCount) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateTopicMessageCount) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateTopicMessageCount) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateTopicMessageCount) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateTopicMessageCount) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateTopicMessageCount) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateTopicMessageCount) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateTopicMessageCount) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateTopicMessageCount) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateTopicMessageCount) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateTopicMessageCount) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateTopicMessageCount) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateTopicMessageCount) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateTopicMessageCount) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateTopicMessageCount) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateTopicMessageCount) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateTopicMessageCount) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateTopicMessageCount) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateTopicMessageCount) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateTopicMessageCount) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateTopicMessageCount) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateTopicMessageCount) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateTopicMessageCount) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateTopicMessageCount) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateTopicMessageCount) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateTopicMessageCount) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateTopicMessageCount) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateTopicMessageCount) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateTopicMessageCount) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateTopicMessageCount) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateTopicMessageCount) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateTopicMessageCount) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateTopicMessageCount) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateTopicMessageCount) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateTopicMessageCount) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateTopicMessageCount) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateTopicMessageCount) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateTopicMessageCount) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateTopicMessageCount) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateTopicMessageCount) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateTopicMessageCount) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateTopicMessageCount) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateTopicMessageCount) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateTopicMessageCount) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateTopicMessageCount) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateTopicMessageCount) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateTopicMessageCount) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateTopicMessageCount) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateTopicMessageCount) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateTopicMessageCount) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateTopicMessageCount) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateTopicMessageCount) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateTopicMessageCount) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateTopicMessageCount) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateTopicMessageCount) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateTopicMessageCount) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateTopicMessageCount) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateTopicMessageCount) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateTopicMessageCount) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateTopicMessageCount) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateTopicMessageCount) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateTopicMessageCount) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateTopicMessageCount) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateTopicMessageCount) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateTopicMessageCount) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateTopicMessageCount) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateTopicMessageCount) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateTopicMessageCount) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateTopicMessageCount) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateTopicMessageCount) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateTopicMessageCount) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateTopicMessageCount) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateTopicMessageCount) SendChatAction(client *Client, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, u.TopicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateTopicMessageCount) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateTopicMessageCount) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateTopicMessageCount) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateTopicMessageCount) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateTopicMessageCount) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateTopicMessageCount) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateTopicMessageCount) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateTopicMessageCount) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateTopicMessageCount) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateTopicMessageCount) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateTopicMessageCount) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateTopicMessageCount) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateTopicMessageCount) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateTopicMessageCount) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateTopicMessageCount) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateTopicMessageCount) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateTopicMessageCount) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateTopicMessageCount) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateTopicMessageCount) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateTopicMessageCount) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateTopicMessageCount) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateTopicMessageCount) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateTopicMessageCount) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateTopicMessageCount) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateTopicMessageCount) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateTopicMessageCount) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateTopicMessageCount) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateTopicMessageCount) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateTopicMessageCount) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateTopicMessageCount) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateTopicMessageCount) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateTopicMessageCount) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateTopicMessageCount) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateTopicMessageCount) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateTopicMessageCount) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateTopicMessageCount) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateTopicMessageCount) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateTopicMessageCount) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateTopicMessageCount) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateTopicMessageCount) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateTopicMessageCount) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateTopicMessageCount) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateTopicMessageCount) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateTopicMessageCount) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateTopicMessageCount) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateTopicMessageCount) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateTopicMessageCount) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateTopicMessageCount) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateTopicMessageCount) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateTopicMessageCount) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateTopicMessageCount) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateTopicMessageCount) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateTopicMessageCount) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateTopicMessageCount) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateTopicMessageCount) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateTopicMessageCount) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateTopicMessageCount) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateTopicMessageCount) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateTopicMessageCount) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateTopicMessageCount) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateTopicMessageCount) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateTopicMessageCount) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateTopicMessageCount) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateTopicMessageCount) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateTopicMessageCount) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateTopicMessageCount) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateTopicMessageCount) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdateTrendingStickerSets) CreateNewStickerSet(client *Client, userId int64, title string, name string, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, name, u.StickerType, needsRepainting, stickers, source)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateTrendingStickerSets) GetAllStickerEmojis(client *Client, query string, chatId int64, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(u.StickerType, query, chatId, returnOnlyMainEmoji)
}

// GetArchivedStickerSets is a helper method for Client.GetArchivedStickerSets
func (u UpdateTrendingStickerSets) GetArchivedStickerSets(client *Client, offsetStickerSetId string, limit int32) (*StickerSets, error) {
	return client.GetArchivedStickerSets(u.StickerType, offsetStickerSetId, limit)
}

// GetInstalledStickerSets is a helper method for Client.GetInstalledStickerSets
func (u UpdateTrendingStickerSets) GetInstalledStickerSets(client *Client) (*StickerSets, error) {
	return client.GetInstalledStickerSets(u.StickerType)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateTrendingStickerSets) GetStickers(client *Client, query string, limit int32, chatId int64) (*Stickers, error) {
	return client.GetStickers(u.StickerType, query, limit, chatId)
}

// GetTrendingStickerSets is a helper method for Client.GetTrendingStickerSets
func (u UpdateTrendingStickerSets) GetTrendingStickerSets(client *Client, offset int32, limit int32) (*TrendingStickerSets, error) {
	return client.GetTrendingStickerSets(u.StickerType, offset, limit)
}

// ReorderInstalledStickerSets is a helper method for Client.ReorderInstalledStickerSets
func (u UpdateTrendingStickerSets) ReorderInstalledStickerSets(client *Client, stickerSetIds []string) (*Ok, error) {
	return client.ReorderInstalledStickerSets(u.StickerType, stickerSetIds)
}

// SearchInstalledStickerSets is a helper method for Client.SearchInstalledStickerSets
func (u UpdateTrendingStickerSets) SearchInstalledStickerSets(client *Client, query string, limit int32) (*StickerSets, error) {
	return client.SearchInstalledStickerSets(u.StickerType, query, limit)
}

// SearchStickers is a helper method for Client.SearchStickers
func (u UpdateTrendingStickerSets) SearchStickers(client *Client, emojis string, query string, inputLanguageCodes []string, offset int32, limit int32) (*Stickers, error) {
	return client.SearchStickers(u.StickerType, emojis, query, inputLanguageCodes, offset, limit)
}

// SearchStickerSets is a helper method for Client.SearchStickerSets
func (u UpdateTrendingStickerSets) SearchStickerSets(client *Client, query string) (*StickerSets, error) {
	return client.SearchStickerSets(u.StickerType, query)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateUnreadChatCount) AddChatToList(client *Client, chatId int64) (*Ok, error) {
	return client.AddChatToList(chatId, u.ChatList)
}

// ReadChatList is a helper method for Client.ReadChatList
func (u UpdateUnreadChatCount) ReadChatList(client *Client) (*Ok, error) {
	return client.ReadChatList(u.ChatList)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (u UpdateUnreadChatCount) SetPinnedChats(client *Client, chatIds []int64) (*Ok, error) {
	return client.SetPinnedChats(u.ChatList, chatIds)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateUnreadChatCount) ToggleChatIsPinned(client *Client, chatId int64, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(u.ChatList, chatId, isPinned)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateUnreadMessageCount) AddChatToList(client *Client, chatId int64) (*Ok, error) {
	return client.AddChatToList(chatId, u.ChatList)
}

// ReadChatList is a helper method for Client.ReadChatList
func (u UpdateUnreadMessageCount) ReadChatList(client *Client) (*Ok, error) {
	return client.ReadChatList(u.ChatList)
}

// SetPinnedChats is a helper method for Client.SetPinnedChats
func (u UpdateUnreadMessageCount) SetPinnedChats(client *Client, chatIds []int64) (*Ok, error) {
	return client.SetPinnedChats(u.ChatList, chatIds)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateUnreadMessageCount) ToggleChatIsPinned(client *Client, chatId int64, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(u.ChatList, chatId, isPinned)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateUserFullInfo) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, u.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (u UpdateUserFullInfo) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(u.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpdateUserFullInfo) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(u.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (u UpdateUserFullInfo) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(u.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (u UpdateUserFullInfo) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(u.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (u UpdateUserFullInfo) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(u.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (u UpdateUserFullInfo) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(u.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdateUserFullInfo) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(u.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (u UpdateUserFullInfo) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(u.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (u UpdateUserFullInfo) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(u.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateUserFullInfo) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, u.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (u UpdateUserFullInfo) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(u.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (u UpdateUserFullInfo) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, u.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (u UpdateUserFullInfo) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(u.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (u UpdateUserFullInfo) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(u.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (u UpdateUserFullInfo) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(u.UserId)
}

// GetUser is a helper method for Client.GetUser
func (u UpdateUserFullInfo) GetUser(client *Client) (*User, error) {
	return client.GetUser(u.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateUserFullInfo) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, u.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (u UpdateUserFullInfo) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(u.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (u UpdateUserFullInfo) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(u.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (u UpdateUserFullInfo) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(u.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (u UpdateUserFullInfo) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(u.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UpdateUserFullInfo) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(u.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateUserFullInfo) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, u.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UpdateUserFullInfo) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, u.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateUserFullInfo) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, u.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (u UpdateUserFullInfo) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(u.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpdateUserFullInfo) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(u.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (u UpdateUserFullInfo) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(u.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateUserFullInfo) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, u.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (u UpdateUserFullInfo) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, u.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (u UpdateUserFullInfo) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(u.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (u UpdateUserFullInfo) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(u.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpdateUserFullInfo) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(u.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (u UpdateUserFullInfo) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(u.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (u UpdateUserFullInfo) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(u.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (u UpdateUserFullInfo) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(u.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (u UpdateUserFullInfo) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(u.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (u UpdateUserFullInfo) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(u.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (u UpdateUserFullInfo) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(u.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (u UpdateUserFullInfo) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(u.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateUserFullInfo) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, u.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (u UpdateUserFullInfo) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(u.UserId, stickerFormat, sticker)
}

// GetUserPrivacySettingRules is a helper method for Client.GetUserPrivacySettingRules
func (u UpdateUserPrivacySettingRules) GetUserPrivacySettingRules(client *Client) (*UserPrivacySettingRules, error) {
	return client.GetUserPrivacySettingRules(u.Setting)
}

// SetUserPrivacySettingRules is a helper method for Client.SetUserPrivacySettingRules
func (u UpdateUserPrivacySettingRules) SetUserPrivacySettingRules(client *Client) (*Ok, error) {
	return client.SetUserPrivacySettingRules(u.Setting, u.Rules)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateUserStatus) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, u.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (u UpdateUserStatus) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(u.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpdateUserStatus) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(u.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (u UpdateUserStatus) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(u.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (u UpdateUserStatus) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(u.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (u UpdateUserStatus) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(u.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (u UpdateUserStatus) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(u.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpdateUserStatus) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(u.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (u UpdateUserStatus) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(u.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (u UpdateUserStatus) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(u.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateUserStatus) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, u.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (u UpdateUserStatus) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(u.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (u UpdateUserStatus) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, u.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (u UpdateUserStatus) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(u.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (u UpdateUserStatus) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(u.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (u UpdateUserStatus) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(u.UserId)
}

// GetUser is a helper method for Client.GetUser
func (u UpdateUserStatus) GetUser(client *Client) (*User, error) {
	return client.GetUser(u.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateUserStatus) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, u.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (u UpdateUserStatus) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(u.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (u UpdateUserStatus) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(u.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (u UpdateUserStatus) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(u.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (u UpdateUserStatus) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(u.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UpdateUserStatus) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(u.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (u UpdateUserStatus) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, u.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UpdateUserStatus) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, u.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateUserStatus) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, u.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (u UpdateUserStatus) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(u.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpdateUserStatus) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(u.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (u UpdateUserStatus) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(u.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateUserStatus) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, u.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (u UpdateUserStatus) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, u.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (u UpdateUserStatus) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(u.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (u UpdateUserStatus) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(u.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpdateUserStatus) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(u.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (u UpdateUserStatus) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(u.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (u UpdateUserStatus) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(u.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (u UpdateUserStatus) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(u.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (u UpdateUserStatus) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(u.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (u UpdateUserStatus) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(u.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (u UpdateUserStatus) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(u.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (u UpdateUserStatus) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(u.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateUserStatus) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, u.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (u UpdateUserStatus) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(u.UserId, stickerFormat, sticker)
}

// AddChatMember is a helper method for Client.AddChatMember
func (u UpdateVideoPublished) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(u.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (u UpdateVideoPublished) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(u.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (u UpdateVideoPublished) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(u.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (u UpdateVideoPublished) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(u.ChatId, u.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (u UpdateVideoPublished) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, u.ChatId, u.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpdateVideoPublished) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(u.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (u UpdateVideoPublished) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(u.ChatId, u.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (u UpdateVideoPublished) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(u.ChatId, u.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (u UpdateVideoPublished) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(u.ChatId, u.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (u UpdateVideoPublished) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(u.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (u UpdateVideoPublished) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (u UpdateVideoPublished) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(u.ChatId, u.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (u UpdateVideoPublished) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(u.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (u UpdateVideoPublished) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(u.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (u UpdateVideoPublished) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(u.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (u UpdateVideoPublished) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(u.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (u UpdateVideoPublished) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(u.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (u UpdateVideoPublished) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(u.ChatId, u.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (u UpdateVideoPublished) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(u.ChatId, u.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (u UpdateVideoPublished) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(u.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (u UpdateVideoPublished) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpdateVideoPublished) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(u.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpdateVideoPublished) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(u.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpdateVideoPublished) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(u.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpdateVideoPublished) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(u.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (u UpdateVideoPublished) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(u.ChatId, u.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (u UpdateVideoPublished) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(u.ChatId, u.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (u UpdateVideoPublished) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(u.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (u UpdateVideoPublished) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(u.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (u UpdateVideoPublished) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(u.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (u UpdateVideoPublished) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(u.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (u UpdateVideoPublished) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(u.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpdateVideoPublished) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(u.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (u UpdateVideoPublished) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(u.ChatId, u.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (u UpdateVideoPublished) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(u.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (u UpdateVideoPublished) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(u.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (u UpdateVideoPublished) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(u.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (u UpdateVideoPublished) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(u.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (u UpdateVideoPublished) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(u.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (u UpdateVideoPublished) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(u.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (u UpdateVideoPublished) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (u UpdateVideoPublished) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, u.ChatId, u.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (u UpdateVideoPublished) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (u UpdateVideoPublished) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (u UpdateVideoPublished) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (u UpdateVideoPublished) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpdateVideoPublished) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(u.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpdateVideoPublished) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(u.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpdateVideoPublished) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(u.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (u UpdateVideoPublished) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(u.ChatId, u.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (u UpdateVideoPublished) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(u.ChatId, u.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (u UpdateVideoPublished) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(u.ChatId, u.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (u UpdateVideoPublished) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (u UpdateVideoPublished) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(u.ChatId, u.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (u UpdateVideoPublished) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(u.ChatId, u.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (u UpdateVideoPublished) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(u.ChatId, u.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (u UpdateVideoPublished) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, u.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (u UpdateVideoPublished) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(u.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (u UpdateVideoPublished) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, u.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (u UpdateVideoPublished) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(u.ChatId, u.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (u UpdateVideoPublished) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(u.ChatId, u.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (u UpdateVideoPublished) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(u.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (u UpdateVideoPublished) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(u.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (u UpdateVideoPublished) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(u.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (u UpdateVideoPublished) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(u.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (u UpdateVideoPublished) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(u.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (u UpdateVideoPublished) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(u.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (u UpdateVideoPublished) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(u.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (u UpdateVideoPublished) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(u.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (u UpdateVideoPublished) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(u.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (u UpdateVideoPublished) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(u.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (u UpdateVideoPublished) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(u.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (u UpdateVideoPublished) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(u.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (u UpdateVideoPublished) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(u.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (u UpdateVideoPublished) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(u.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (u UpdateVideoPublished) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(u.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (u UpdateVideoPublished) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(u.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (u UpdateVideoPublished) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(u.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (u UpdateVideoPublished) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(u.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpdateVideoPublished) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(u.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (u UpdateVideoPublished) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(u.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (u UpdateVideoPublished) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(u.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (u UpdateVideoPublished) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(u.ChatId, filter, u.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (u UpdateVideoPublished) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(u.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (u UpdateVideoPublished) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(u.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (u UpdateVideoPublished) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(u.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (u UpdateVideoPublished) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(u.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (u UpdateVideoPublished) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(u.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (u UpdateVideoPublished) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(u.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (u UpdateVideoPublished) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(u.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (u UpdateVideoPublished) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(u.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (u UpdateVideoPublished) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(u.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (u UpdateVideoPublished) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(u.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (u UpdateVideoPublished) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(u.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (u UpdateVideoPublished) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(u.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (u UpdateVideoPublished) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(u.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (u UpdateVideoPublished) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(u.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpdateVideoPublished) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(u.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (u UpdateVideoPublished) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(u.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (u UpdateVideoPublished) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(u.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (u UpdateVideoPublished) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(u.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (u UpdateVideoPublished) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(u.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (u UpdateVideoPublished) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(u.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (u UpdateVideoPublished) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(u.ChatId, u.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (u UpdateVideoPublished) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(u.ChatId, u.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (u UpdateVideoPublished) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, u.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (u UpdateVideoPublished) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(u.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (u UpdateVideoPublished) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(u.ChatId, u.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (u UpdateVideoPublished) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(u.ChatId, u.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (u UpdateVideoPublished) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(u.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (u UpdateVideoPublished) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, u.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (u UpdateVideoPublished) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(u.ChatId, u.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (u UpdateVideoPublished) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(u.ChatId, u.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (u UpdateVideoPublished) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(u.ChatId, u.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (u UpdateVideoPublished) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(u.ChatId, u.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (u UpdateVideoPublished) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(u.ChatId, u.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (u UpdateVideoPublished) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(u.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (u UpdateVideoPublished) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(u.ChatId, u.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (u UpdateVideoPublished) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(u.ChatId, u.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (u UpdateVideoPublished) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(u.ChatId, u.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (u UpdateVideoPublished) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(u.ChatId, u.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (u UpdateVideoPublished) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(u.ChatId, u.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (u UpdateVideoPublished) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(u.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (u UpdateVideoPublished) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(u.ChatId, u.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (u UpdateVideoPublished) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(u.ChatId, u.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (u UpdateVideoPublished) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(u.ChatId, u.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (u UpdateVideoPublished) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(u.ChatId, u.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (u UpdateVideoPublished) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(u.ChatId, u.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (u UpdateVideoPublished) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(u.ChatId, u.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (u UpdateVideoPublished) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(u.ChatId, u.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (u UpdateVideoPublished) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(u.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (u UpdateVideoPublished) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, u.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (u UpdateVideoPublished) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(u.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (u UpdateVideoPublished) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(u.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (u UpdateVideoPublished) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(u.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (u UpdateVideoPublished) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(u.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (u UpdateVideoPublished) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(u.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (u UpdateVideoPublished) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(u.ChatId, u.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (u UpdateVideoPublished) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(u.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (u UpdateVideoPublished) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(u.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (u UpdateVideoPublished) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(u.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (u UpdateVideoPublished) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(u.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (u UpdateVideoPublished) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(u.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (u UpdateVideoPublished) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(u.ChatId, u.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (u UpdateVideoPublished) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(u.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (u UpdateVideoPublished) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(u.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (u UpdateVideoPublished) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(u.ChatId, u.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (u UpdateVideoPublished) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(u.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (u UpdateVideoPublished) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(u.ChatId, u.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (u UpdateVideoPublished) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(u.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (u UpdateVideoPublished) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(u.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (u UpdateVideoPublished) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(u.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (u UpdateVideoPublished) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(u.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (u UpdateVideoPublished) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(u.ChatId, u.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (u UpdateVideoPublished) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(u.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (u UpdateVideoPublished) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(u.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (u UpdateVideoPublished) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(u.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (u UpdateVideoPublished) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(u.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (u UpdateVideoPublished) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(u.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (u UpdateVideoPublished) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, u.ChatId, u.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (u UpdateVideoPublished) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(u.ChatId, u.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (u UpdateVideoPublished) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(u.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (u UpdateVideoPublished) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(u.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (u UpdateVideoPublished) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(u.ChatId, u.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (u UpdateVideoPublished) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(u.ChatId, u.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (u UpdateVideoPublished) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(u.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (u UpdateVideoPublished) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (u UpdateVideoPublished) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, u.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (u UpdateVideoPublished) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(u.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (u UpdateVideoPublished) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(u.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (u UpdateVideoPublished) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(u.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (u UpdateVideoPublished) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(u.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (u UpdateVideoPublished) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(u.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (u UpdateVideoPublished) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(u.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (u UpdateVideoPublished) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(u.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (u UpdateVideoPublished) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(u.ChatId, u.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpdateVideoPublished) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(u.ChatId, u.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (u UpdateVideoPublished) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, u.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (u UpdateVideoPublished) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(u.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (u UpdateVideoPublished) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(u.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (u UpdateVideoPublished) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, u.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (u UpdateVideoPublished) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(u.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (u UpdateVideoPublished) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(u.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (u UpdateVideoPublished) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(u.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (u UpdateVideoPublished) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(u.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (u UpdateVideoPublished) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, u.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (u UpdateVideoPublished) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (u UpdateVideoPublished) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, u.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (u UpdateVideoPublished) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(u.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (u UpdateVideoPublished) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(u.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (u UpdateVideoPublished) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(u.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (u UpdateVideoPublished) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(u.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (u UpdateVideoPublished) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(u.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpdateVideoPublished) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(u.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (u UpdateVideoPublished) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, u.ChatId, u.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (u UpdateVideoPublished) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(u.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (u UpdateVideoPublished) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(u.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (u UpdateVideoPublished) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(u.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (u UpdateVideoPublished) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(u.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (u UpdateVideoPublished) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(u.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (u UpdateVideoPublished) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(u.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (u UpdateVideoPublished) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(u.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (u UpdateVideoPublished) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(u.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (u UpdateVideoPublished) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(u.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (u UpdateVideoPublished) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(u.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (u UpdateVideoPublished) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(u.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (u UpdateVideoPublished) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(u.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (u UpdateVideoPublished) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(u.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (u UpdateVideoPublished) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(u.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (u UpdateVideoPublished) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(u.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (u UpdateVideoPublished) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(u.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (u UpdateVideoPublished) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(u.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (u UpdateVideoPublished) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(u.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (u UpdateVideoPublished) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(u.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (u UpdateVideoPublished) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(u.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (u UpdateVideoPublished) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(u.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (u UpdateVideoPublished) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(u.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (u UpdateVideoPublished) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(u.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpdateVideoPublished) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(u.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (u UpdateVideoPublished) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(u.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (u UpdateVideoPublished) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(u.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (u UpdateVideoPublished) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(u.ChatId, u.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (u UpdateVideoPublished) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(u.ChatId, u.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (u UpdateVideoPublished) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(u.ChatId, u.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (u UpdateVideoPublished) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(u.ChatId, u.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (u UpdateVideoPublished) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(u.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (u UpdateVideoPublished) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(u.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (u UpdateVideoPublished) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(u.ChatId, u.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpdateVideoPublished) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(u.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (u UpdateVideoPublished) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(u.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (u UpdateVideoPublished) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(u.ChatId, u.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (u UpdateVideoPublished) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(u.ChatId, u.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (u UpdateVideoPublished) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(u.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (u UpdateVideoPublished) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, u.ChatId, u.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (u UpdateVideoPublished) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(u.ChatId, u.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (u UpdateVideoPublished) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(u.ChatId, u.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (u UpdateVideoPublished) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(u.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (u UpdateVideoPublished) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(u.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (u UpdateVideoPublished) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(u.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (u UpdateVideoPublished) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(u.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (u UpdateVideoPublished) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(u.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (u UpdateVideoPublished) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, u.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (u UpdateVideoPublished) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(u.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (u UpdateVideoPublished) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(u.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (u UpdateVideoPublished) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(u.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (u UpdateVideoPublished) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(u.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (u UpdateVideoPublished) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(u.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (u UpdateVideoPublished) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(u.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (u UpdateVideoPublished) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(u.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (u UpdateVideoPublished) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(u.ChatId, u.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (u UpdateVideoPublished) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(u.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (u UpdateVideoPublished) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(u.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (u UpdateVideoPublished) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(u.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (u UpdateVideoPublished) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(u.ChatId, u.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (u UpdateVideoPublished) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(u.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (u UpdateVideoPublished) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(u.ChatId, messageIds, forceRead, opts)
}

// CloseWebApp is a helper method for Client.CloseWebApp
func (u UpdateWebAppMessageSent) CloseWebApp(client *Client) (*Ok, error) {
	return client.CloseWebApp(u.WebAppLaunchId)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (u UpgradedGift) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpgradedGift) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, u.Name, sticker)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (u UpgradedGift) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(u.OwnerId, prepaidUpgradeHash, starCount)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (u UpgradedGift) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(u.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (u UpgradedGift) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(u.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UpgradedGift) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, u.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpgradedGift) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpgradedGift) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, u.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpgradedGift) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, u.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpgradedGift) CreateGiftCollection(client *Client, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(u.OwnerId, u.Name, receivedGiftIds)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (u UpgradedGift) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, u.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpgradedGift) CreateNewStickerSet(client *Client, userId int64, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, u.Title, u.Name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (u UpgradedGift) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(u.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpgradedGift) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, u.Name, storyIds)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (u UpgradedGift) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, u.Title, startDate, isRtmpStream)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (u UpgradedGift) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(u.OwnerId, collectionId)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (u UpgradedGift) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(u.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UpgradedGift) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, u.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpgradedGift) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpgradedGift) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, u.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpgradedGift) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, u.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (u UpgradedGift) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(u.Name, typeField)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (u UpgradedGift) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(u.OwnerId)
}

// GetOption is a helper method for Client.GetOption
func (u UpgradedGift) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(u.Name)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (u UpgradedGift) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, u.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (u UpgradedGift) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(u.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (u UpgradedGift) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(u.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (u UpgradedGift) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(u.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (u UpgradedGift) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(u.OwnerId, starCount, password)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (u UpgradedGift) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(u.Title)
}

// Get is a helper method for Client.GetUpgradedGift
func (u UpgradedGift) Get(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(u.Name)
}

// GetValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (u UpgradedGift) GetValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(u.Name)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (u UpgradedGift) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (u UpgradedGift) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(u.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (u UpgradedGift) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(u.OwnerId, collectionIds)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpgradedGift) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, u.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (u UpgradedGift) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(u.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (u UpgradedGift) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(u.Name, ignoreCache)
}

// SendGift is a helper method for Client.SendGift
func (u UpgradedGift) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, u.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u UpgradedGift) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(u.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (u UpgradedGift) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, u.OwnerId, price)
}

// SetBotName is a helper method for Client.SetBotName
func (u UpgradedGift) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, u.Name)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (u UpgradedGift) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, u.Title)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (u UpgradedGift) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(u.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpgradedGift) SetGiftCollectionName(client *Client, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(u.OwnerId, collectionId, u.Name)
}

// SetOption is a helper method for Client.SetOption
func (u UpgradedGift) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(u.Name, opts)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (u UpgradedGift) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(u.OwnerId, receivedGiftIds)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpgradedGift) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, u.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpgradedGift) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, u.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpgradedGift) SetStickerSetTitle(client *Client) (*Ok, error) {
	return client.SetStickerSetTitle(u.Name, u.Title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpgradedGift) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, u.Name)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (u UpgradedGift) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, u.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (u UpgradedGift) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, u.Title, recordVideo, usePortraitOrientation)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpgradedGiftBackdrop) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, u.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (u UpgradedGiftBackdrop) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(u.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (u UpgradedGiftBackdrop) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(u.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UpgradedGiftBackdrop) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, u.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpgradedGiftBackdrop) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpgradedGiftBackdrop) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, u.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpgradedGiftBackdrop) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, u.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpgradedGiftBackdrop) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, u.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpgradedGiftBackdrop) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, u.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpgradedGiftBackdrop) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, u.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (u UpgradedGiftBackdrop) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(u.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UpgradedGiftBackdrop) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, u.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpgradedGiftBackdrop) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpgradedGiftBackdrop) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, u.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpgradedGiftBackdrop) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, u.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (u UpgradedGiftBackdrop) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(u.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (u UpgradedGiftBackdrop) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(u.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (u UpgradedGiftBackdrop) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(u.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (u UpgradedGiftBackdrop) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(u.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpgradedGiftBackdrop) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, u.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (u UpgradedGiftBackdrop) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(u.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (u UpgradedGiftBackdrop) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(u.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (u UpgradedGiftBackdrop) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, u.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (u UpgradedGiftBackdrop) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(u.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpgradedGiftBackdrop) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, u.Name)
}

// SetOption is a helper method for Client.SetOption
func (u UpgradedGiftBackdrop) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(u.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpgradedGiftBackdrop) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, u.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpgradedGiftBackdrop) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, u.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpgradedGiftBackdrop) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(u.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpgradedGiftBackdrop) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, u.Name)
}

// Set is a helper method for Client.SetUpgradedGiftColors
func (u UpgradedGiftColors) Set(client *Client) (*Ok, error) {
	return client.SetUpgradedGiftColors(u.Id)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpgradedGiftModel) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, u.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (u UpgradedGiftModel) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(u.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (u UpgradedGiftModel) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(u.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UpgradedGiftModel) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, u.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpgradedGiftModel) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpgradedGiftModel) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, u.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpgradedGiftModel) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, u.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpgradedGiftModel) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, u.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpgradedGiftModel) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, u.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpgradedGiftModel) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, u.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (u UpgradedGiftModel) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(u.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UpgradedGiftModel) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, u.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpgradedGiftModel) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpgradedGiftModel) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, u.Name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (u UpgradedGiftModel) EditForumTopic(client *Client, chatId int64, forumTopicId int32, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(chatId, forumTopicId, u.Name, editIconCustomEmoji, iconCustomEmojiId)
}

// GetBackgroundUrl is a helper method for Client.GetBackgroundUrl
func (u UpgradedGiftModel) GetBackgroundUrl(client *Client, typeField *BackgroundType) (*HttpUrl, error) {
	return client.GetBackgroundUrl(u.Name, typeField)
}

// GetOption is a helper method for Client.GetOption
func (u UpgradedGiftModel) GetOption(client *Client) (*OptionValue, error) {
	return client.GetOption(u.Name)
}

// GetUpgradedGift is a helper method for Client.GetUpgradedGift
func (u UpgradedGiftModel) GetUpgradedGift(client *Client) (*UpgradedGift, error) {
	return client.GetUpgradedGift(u.Name)
}

// GetUpgradedGiftValueInfo is a helper method for Client.GetUpgradedGiftValueInfo
func (u UpgradedGiftModel) GetUpgradedGiftValueInfo(client *Client) (*UpgradedGiftValueInfo, error) {
	return client.GetUpgradedGiftValueInfo(u.Name)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (u UpgradedGiftModel) ReplaceStickerInSet(client *Client, userId int64, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(userId, u.Name, oldSticker, newSticker)
}

// SearchBackground is a helper method for Client.SearchBackground
func (u UpgradedGiftModel) SearchBackground(client *Client) (*Background, error) {
	return client.SearchBackground(u.Name)
}

// SearchStickerSet is a helper method for Client.SearchStickerSet
func (u UpgradedGiftModel) SearchStickerSet(client *Client, ignoreCache bool) (*StickerSet, error) {
	return client.SearchStickerSet(u.Name, ignoreCache)
}

// SetBotName is a helper method for Client.SetBotName
func (u UpgradedGiftModel) SetBotName(client *Client, botUserId int64, languageCode string) (*Ok, error) {
	return client.SetBotName(botUserId, languageCode, u.Name)
}

// SetCustomEmojiStickerSetThumbnail is a helper method for Client.SetCustomEmojiStickerSetThumbnail
func (u UpgradedGiftModel) SetCustomEmojiStickerSetThumbnail(client *Client, customEmojiId string) (*Ok, error) {
	return client.SetCustomEmojiStickerSetThumbnail(u.Name, customEmojiId)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (u UpgradedGiftModel) SetGiftCollectionName(client *Client, ownerId *MessageSender, collectionId int32) (*GiftCollection, error) {
	return client.SetGiftCollectionName(ownerId, collectionId, u.Name)
}

// SetOption is a helper method for Client.SetOption
func (u UpgradedGiftModel) SetOption(client *Client, opts *SetOptionOpts) (*Ok, error) {
	return client.SetOption(u.Name, opts)
}

// SetQuickReplyShortcutName is a helper method for Client.SetQuickReplyShortcutName
func (u UpgradedGiftModel) SetQuickReplyShortcutName(client *Client, shortcutId int32) (*Ok, error) {
	return client.SetQuickReplyShortcutName(shortcutId, u.Name)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (u UpgradedGiftModel) SetStickerSetThumbnail(client *Client, userId int64, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(userId, u.Name, opts)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (u UpgradedGiftModel) SetStickerSetTitle(client *Client, title string) (*Ok, error) {
	return client.SetStickerSetTitle(u.Name, title)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (u UpgradedGiftModel) SetStoryAlbumName(client *Client, chatId int64, storyAlbumId int32) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(chatId, storyAlbumId, u.Name)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (u UpgradedGiftOriginalDetails) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, u.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (u UpgradedGiftOriginalDetails) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, u.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (u UpgradedGiftOriginalDetails) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, u.SenderId, reportSpam)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (u UpgradedGiftOriginalDetails) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, u.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (u UpgradedGiftOriginalDetails) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, u.Date)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (u UpgradedGiftOriginalDetails) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(u.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (u UpgradedGiftOriginalDetails) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(u.Text)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (u UpgradedGiftOriginalDetails) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, u.Date)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (u UpgradedGiftOriginalDetails) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, u.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (u UpgradedGiftOriginalDetails) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(u.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (u UpgradedGiftOriginalDetails) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, u.Text, isPrivate)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (u UpgradedGiftOriginalDetails) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, u.SenderId)
}

// SearchQuote is a helper method for Client.SearchQuote
func (u UpgradedGiftOriginalDetails) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(u.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (u UpgradedGiftOriginalDetails) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, u.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (u UpgradedGiftOriginalDetails) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, u.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (u UpgradedGiftOriginalDetails) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, u.Text)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (u UpgradedGiftOriginalDetails) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(u.SenderId, opts)
}

// TranslateText is a helper method for Client.TranslateText
func (u UpgradedGiftOriginalDetails) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(u.Text, toLanguageCode)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u UpgradedGiftOriginOffer) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, u.Price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (u UpgradedGiftOriginOffer) SendResoldGift(client *Client, giftName string, ownerId *MessageSender) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, ownerId, u.Price)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (u UpgradedGiftOriginResale) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, u.Price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (u UpgradedGiftOriginResale) SendResoldGift(client *Client, giftName string, ownerId *MessageSender) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, ownerId, u.Price)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (u UpgradedGiftSymbol) AddStickerToSet(client *Client, userId int64, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(userId, u.Name, sticker)
}

// CheckQuickReplyShortcutName is a helper method for Client.CheckQuickReplyShortcutName
func (u UpgradedGiftSymbol) CheckQuickReplyShortcutName(client *Client) (*Ok, error) {
	return client.CheckQuickReplyShortcutName(u.Name)
}

// CheckStickerSetName is a helper method for Client.CheckStickerSetName
func (u UpgradedGiftSymbol) CheckStickerSetName(client *Client) (*CheckStickerSetNameResult, error) {
	return client.CheckStickerSetName(u.Name)
}

// CreateChatFolderInviteLink is a helper method for Client.CreateChatFolderInviteLink
func (u UpgradedGiftSymbol) CreateChatFolderInviteLink(client *Client, chatFolderId int32, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.CreateChatFolderInviteLink(chatFolderId, u.Name, chatIds)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (u UpgradedGiftSymbol) CreateChatInviteLink(client *Client, chatId int64, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(chatId, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (u UpgradedGiftSymbol) CreateChatSubscriptionInviteLink(client *Client, chatId int64, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(chatId, u.Name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (u UpgradedGiftSymbol) CreateForumTopic(client *Client, chatId int64, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(chatId, u.Name, isNameImplicit, icon)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (u UpgradedGiftSymbol) CreateGiftCollection(client *Client, ownerId *MessageSender, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(ownerId, u.Name, receivedGiftIds)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (u UpgradedGiftSymbol) CreateNewStickerSet(client *Client, userId int64, title string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, title, u.Name, stickerType, needsRepainting, stickers, source)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (u UpgradedGiftSymbol) CreateStoryAlbum(client *Client, storyPosterChatId int64, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(storyPosterChatId, u.Name, storyIds)
}

// DeleteStickerSet is a helper method for Client.DeleteStickerSet
func (u UpgradedGiftSymbol) DeleteStickerSet(client *Client) (*Ok, error) {
	return client.DeleteStickerSet(u.Name)
}

// EditChatFolderInviteLink is a helper method for Client.EditChatFolderInviteLink
func (u UpgradedGiftSymbol) EditChatFolderInviteLink(client *Client, chatFolderId int32, inviteLink string, chatIds []int64) (*ChatFolderInviteLink, error) {
	return client.EditChatFolderInviteLink(chatFolderId, inviteLink, u.Name, chatIds)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (u UpgradedGiftSymbol) EditChatInviteLink(client *Client, chatId int64, inviteLink string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(chatId, inviteLink, u.Name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (u UpgradedGiftSymbol) EditChatSubscriptionInviteLink(client *Client, chatId int64, inviteLink string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(chatId, inviteLink, u.Name)
}
