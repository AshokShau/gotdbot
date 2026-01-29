package gotdbot

// TransferGift is a helper method for Client.TransferGift
func (m MessageGiveawayCreated) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, m.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessageGiveawayCreated) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, m.StarCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageGiveawayPrizeStars) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, m.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageGiveawayPrizeStars) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, m.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessageGiveawayPrizeStars) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, m.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessageGiveawayPrizeStars) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, m.StarCount)
}

// GetPremiumGiveawayPaymentOptions is a helper method for Client.GetPremiumGiveawayPaymentOptions
func (m MessageGiveawayPrizeStars) GetPremiumGiveawayPaymentOptions(client *Client) (*PremiumGiveawayPaymentOptions, error) {
	return client.GetPremiumGiveawayPaymentOptions(m.BoostedChatId)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessageGiveawayPrizeStars) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, m.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageGiveawayPrizeStars) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, m.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (m MessageGiveawayPrizeStars) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiveawayPrizeStars) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, m.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageGiveawayPrizeStars) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, m.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (m MessageGiveawayPrizeStars) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, m.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (m MessageGiveawayPrizeStars) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, m.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessageGiveawayPrizeStars) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, m.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessageGiveawayPrizeStars) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, m.StarCount)
}

// GetPremiumGiveawayPaymentOptions is a helper method for Client.GetPremiumGiveawayPaymentOptions
func (m MessageGiveawayWinners) GetPremiumGiveawayPaymentOptions(client *Client) (*PremiumGiveawayPaymentOptions, error) {
	return client.GetPremiumGiveawayPaymentOptions(m.BoostedChatId)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessageGiveawayWinners) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, starCount int64) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, m.WinnerCount, starCount)
}

// CreateCall is a helper method for Client.CreateCall
func (m MessageGroupCall) CreateCall(client *Client, userId int64, protocol *CallProtocol) (*CallId, error) {
	return client.CreateCall(userId, protocol, m.IsVideo)
}

// DiscardCall is a helper method for Client.DiscardCall
func (m MessageGroupCall) DiscardCall(client *Client, callId int32, isDisconnected bool, inviteLink string, connectionId string) (*Ok, error) {
	return client.DiscardCall(callId, isDisconnected, inviteLink, m.Duration, m.IsVideo, connectionId)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageGroupCall) InviteGroupCallParticipant(client *Client, groupCallId int32, userId int64) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, userId, m.IsVideo)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageGroupCall) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, m.Duration, paidMessageStarCount)
}

// ToggleBotUsernameIsActive is a helper method for Client.ToggleBotUsernameIsActive
func (m MessageGroupCall) ToggleBotUsernameIsActive(client *Client, botUserId int64, username string) (*Ok, error) {
	return client.ToggleBotUsernameIsActive(botUserId, username, m.IsActive)
}

// ToggleSupergroupUsernameIsActive is a helper method for Client.ToggleSupergroupUsernameIsActive
func (m MessageGroupCall) ToggleSupergroupUsernameIsActive(client *Client, supergroupId int64, username string) (*Ok, error) {
	return client.ToggleSupergroupUsernameIsActive(supergroupId, username, m.IsActive)
}

// ToggleUsernameIsActive is a helper method for Client.ToggleUsernameIsActive
func (m MessageGroupCall) ToggleUsernameIsActive(client *Client, username string) (*Ok, error) {
	return client.ToggleUsernameIsActive(username, m.IsActive)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageImportInfo) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, m.Date)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageImportInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, m.Date)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (m MessageImportInfo) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, m.Date)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageInviteVideoChatParticipants) AddChatMembers(client *Client, chatId int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(chatId, m.UserIds)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessageInviteVideoChatParticipants) AddPendingLiveStoryReaction(client *Client, starCount int64) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(m.GroupCallId, starCount)
}

// BanGroupCallParticipants is a helper method for Client.BanGroupCallParticipants
func (m MessageInviteVideoChatParticipants) BanGroupCallParticipants(client *Client, userIds []string) (*Ok, error) {
	return client.BanGroupCallParticipants(m.GroupCallId, userIds)
}

// CommitPendingLiveStoryReactions is a helper method for Client.CommitPendingLiveStoryReactions
func (m MessageInviteVideoChatParticipants) CommitPendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.CommitPendingLiveStoryReactions(m.GroupCallId)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageInviteVideoChatParticipants) CreateNewBasicGroupChat(client *Client, title string, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(m.UserIds, title, messageAutoDeleteTime)
}

// DecryptGroupCallData is a helper method for Client.DecryptGroupCallData
func (m MessageInviteVideoChatParticipants) DecryptGroupCallData(client *Client, participantId *MessageSender, data string, opts *DecryptGroupCallDataOpts) (*Data, error) {
	return client.DecryptGroupCallData(m.GroupCallId, participantId, data, opts)
}

// DeleteGroupCallMessages is a helper method for Client.DeleteGroupCallMessages
func (m MessageInviteVideoChatParticipants) DeleteGroupCallMessages(client *Client, messageIds []int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessages(m.GroupCallId, messageIds, reportSpam)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageInviteVideoChatParticipants) DeleteGroupCallMessagesBySender(client *Client, senderId *MessageSender, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(m.GroupCallId, senderId, reportSpam)
}

// EncryptGroupCallData is a helper method for Client.EncryptGroupCallData
func (m MessageInviteVideoChatParticipants) EncryptGroupCallData(client *Client, dataChannel *GroupCallDataChannel, data string, unencryptedPrefixSize int32) (*Data, error) {
	return client.EncryptGroupCallData(m.GroupCallId, dataChannel, data, unencryptedPrefixSize)
}

// EndGroupCall is a helper method for Client.EndGroupCall
func (m MessageInviteVideoChatParticipants) EndGroupCall(client *Client) (*Ok, error) {
	return client.EndGroupCall(m.GroupCallId)
}

// EndGroupCallRecording is a helper method for Client.EndGroupCallRecording
func (m MessageInviteVideoChatParticipants) EndGroupCallRecording(client *Client) (*Ok, error) {
	return client.EndGroupCallRecording(m.GroupCallId)
}

// EndGroupCallScreenSharing is a helper method for Client.EndGroupCallScreenSharing
func (m MessageInviteVideoChatParticipants) EndGroupCallScreenSharing(client *Client) (*Ok, error) {
	return client.EndGroupCallScreenSharing(m.GroupCallId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageInviteVideoChatParticipants) GetChatEventLog(client *Client, chatId int64, query string, fromEventId string, limit int32, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(chatId, query, fromEventId, limit, m.UserIds, opts)
}

// GetGroupCall is a helper method for Client.GetGroupCall
func (m MessageInviteVideoChatParticipants) GetGroupCall(client *Client) (*GroupCall, error) {
	return client.GetGroupCall(m.GroupCallId)
}

// GetGroupCallStreams is a helper method for Client.GetGroupCallStreams
func (m MessageInviteVideoChatParticipants) GetGroupCallStreams(client *Client) (*GroupCallStreams, error) {
	return client.GetGroupCallStreams(m.GroupCallId)
}

// GetGroupCallStreamSegment is a helper method for Client.GetGroupCallStreamSegment
func (m MessageInviteVideoChatParticipants) GetGroupCallStreamSegment(client *Client, timeOffset int64, scale int32, channelId int32, opts *GetGroupCallStreamSegmentOpts) (*Data, error) {
	return client.GetGroupCallStreamSegment(m.GroupCallId, timeOffset, scale, channelId, opts)
}

// GetLiveStoryAvailableMessageSenders is a helper method for Client.GetLiveStoryAvailableMessageSenders
func (m MessageInviteVideoChatParticipants) GetLiveStoryAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetLiveStoryAvailableMessageSenders(m.GroupCallId)
}

// GetLiveStoryStreamer is a helper method for Client.GetLiveStoryStreamer
func (m MessageInviteVideoChatParticipants) GetLiveStoryStreamer(client *Client) (*GroupCallParticipant, error) {
	return client.GetLiveStoryStreamer(m.GroupCallId)
}

// GetLiveStoryTopDonors is a helper method for Client.GetLiveStoryTopDonors
func (m MessageInviteVideoChatParticipants) GetLiveStoryTopDonors(client *Client) (*LiveStoryDonors, error) {
	return client.GetLiveStoryTopDonors(m.GroupCallId)
}

// GetVideoChatInviteLink is a helper method for Client.GetVideoChatInviteLink
func (m MessageInviteVideoChatParticipants) GetVideoChatInviteLink(client *Client, canSelfUnmute bool) (*HttpUrl, error) {
	return client.GetVideoChatInviteLink(m.GroupCallId, canSelfUnmute)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageInviteVideoChatParticipants) InviteGroupCallParticipant(client *Client, userId int64, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(m.GroupCallId, userId, isVideo)
}

// InviteVideoChatParticipants is a helper method for Client.InviteVideoChatParticipants
func (m MessageInviteVideoChatParticipants) InviteVideoChatParticipants(client *Client) (*Ok, error) {
	return client.InviteVideoChatParticipants(m.GroupCallId, m.UserIds)
}

// JoinLiveStory is a helper method for Client.JoinLiveStory
func (m MessageInviteVideoChatParticipants) JoinLiveStory(client *Client, joinParameters *GroupCallJoinParameters) (*Text, error) {
	return client.JoinLiveStory(m.GroupCallId, joinParameters)
}

// JoinVideoChat is a helper method for Client.JoinVideoChat
func (m MessageInviteVideoChatParticipants) JoinVideoChat(client *Client, joinParameters *GroupCallJoinParameters, inviteHash string, opts *JoinVideoChatOpts) (*Text, error) {
	return client.JoinVideoChat(m.GroupCallId, joinParameters, inviteHash, opts)
}

// LeaveGroupCall is a helper method for Client.LeaveGroupCall
func (m MessageInviteVideoChatParticipants) LeaveGroupCall(client *Client) (*Ok, error) {
	return client.LeaveGroupCall(m.GroupCallId)
}

// LoadGroupCallParticipants is a helper method for Client.LoadGroupCallParticipants
func (m MessageInviteVideoChatParticipants) LoadGroupCallParticipants(client *Client, limit int32) (*Ok, error) {
	return client.LoadGroupCallParticipants(m.GroupCallId, limit)
}

// RemoveContacts is a helper method for Client.RemoveContacts
func (m MessageInviteVideoChatParticipants) RemoveContacts(client *Client) (*Ok, error) {
	return client.RemoveContacts(m.UserIds)
}

// RemovePendingLiveStoryReactions is a helper method for Client.RemovePendingLiveStoryReactions
func (m MessageInviteVideoChatParticipants) RemovePendingLiveStoryReactions(client *Client) (*Ok, error) {
	return client.RemovePendingLiveStoryReactions(m.GroupCallId)
}

// RevokeGroupCallInviteLink is a helper method for Client.RevokeGroupCallInviteLink
func (m MessageInviteVideoChatParticipants) RevokeGroupCallInviteLink(client *Client) (*Ok, error) {
	return client.RevokeGroupCallInviteLink(m.GroupCallId)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageInviteVideoChatParticipants) SendGroupCallMessage(client *Client, text *FormattedText, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(m.GroupCallId, text, paidMessageStarCount)
}

// SetCloseFriends is a helper method for Client.SetCloseFriends
func (m MessageInviteVideoChatParticipants) SetCloseFriends(client *Client) (*Ok, error) {
	return client.SetCloseFriends(m.UserIds)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessageInviteVideoChatParticipants) SetGroupCallPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(m.GroupCallId, paidMessageStarCount)
}

// SetGroupCallParticipantIsSpeaking is a helper method for Client.SetGroupCallParticipantIsSpeaking
func (m MessageInviteVideoChatParticipants) SetGroupCallParticipantIsSpeaking(client *Client, audioSource int32, isSpeaking bool) (*MessageSender, error) {
	return client.SetGroupCallParticipantIsSpeaking(m.GroupCallId, audioSource, isSpeaking)
}

// SetGroupCallParticipantVolumeLevel is a helper method for Client.SetGroupCallParticipantVolumeLevel
func (m MessageInviteVideoChatParticipants) SetGroupCallParticipantVolumeLevel(client *Client, participantId *MessageSender, volumeLevel int32) (*Ok, error) {
	return client.SetGroupCallParticipantVolumeLevel(m.GroupCallId, participantId, volumeLevel)
}

// SetLiveStoryMessageSender is a helper method for Client.SetLiveStoryMessageSender
func (m MessageInviteVideoChatParticipants) SetLiveStoryMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetLiveStoryMessageSender(m.GroupCallId, messageSenderId)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageInviteVideoChatParticipants) SetVideoChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetVideoChatTitle(m.GroupCallId, title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageInviteVideoChatParticipants) StartGroupCallRecording(client *Client, title string, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(m.GroupCallId, title, recordVideo, usePortraitOrientation)
}

// StartGroupCallScreenSharing is a helper method for Client.StartGroupCallScreenSharing
func (m MessageInviteVideoChatParticipants) StartGroupCallScreenSharing(client *Client, audioSourceId int32, payload string) (*Text, error) {
	return client.StartGroupCallScreenSharing(m.GroupCallId, audioSourceId, payload)
}

// StartScheduledVideoChat is a helper method for Client.StartScheduledVideoChat
func (m MessageInviteVideoChatParticipants) StartScheduledVideoChat(client *Client) (*Ok, error) {
	return client.StartScheduledVideoChat(m.GroupCallId)
}

// ToggleGroupCallAreMessagesAllowed is a helper method for Client.ToggleGroupCallAreMessagesAllowed
func (m MessageInviteVideoChatParticipants) ToggleGroupCallAreMessagesAllowed(client *Client, areMessagesAllowed bool) (*Ok, error) {
	return client.ToggleGroupCallAreMessagesAllowed(m.GroupCallId, areMessagesAllowed)
}

// ToggleGroupCallIsMyVideoEnabled is a helper method for Client.ToggleGroupCallIsMyVideoEnabled
func (m MessageInviteVideoChatParticipants) ToggleGroupCallIsMyVideoEnabled(client *Client, isMyVideoEnabled bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoEnabled(m.GroupCallId, isMyVideoEnabled)
}

// ToggleGroupCallIsMyVideoPaused is a helper method for Client.ToggleGroupCallIsMyVideoPaused
func (m MessageInviteVideoChatParticipants) ToggleGroupCallIsMyVideoPaused(client *Client, isMyVideoPaused bool) (*Ok, error) {
	return client.ToggleGroupCallIsMyVideoPaused(m.GroupCallId, isMyVideoPaused)
}

// ToggleGroupCallParticipantIsHandRaised is a helper method for Client.ToggleGroupCallParticipantIsHandRaised
func (m MessageInviteVideoChatParticipants) ToggleGroupCallParticipantIsHandRaised(client *Client, participantId *MessageSender, isHandRaised bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsHandRaised(m.GroupCallId, participantId, isHandRaised)
}

// ToggleGroupCallParticipantIsMuted is a helper method for Client.ToggleGroupCallParticipantIsMuted
func (m MessageInviteVideoChatParticipants) ToggleGroupCallParticipantIsMuted(client *Client, participantId *MessageSender, isMuted bool) (*Ok, error) {
	return client.ToggleGroupCallParticipantIsMuted(m.GroupCallId, participantId, isMuted)
}

// ToggleGroupCallScreenSharingIsPaused is a helper method for Client.ToggleGroupCallScreenSharingIsPaused
func (m MessageInviteVideoChatParticipants) ToggleGroupCallScreenSharingIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleGroupCallScreenSharingIsPaused(m.GroupCallId, isPaused)
}

// ToggleVideoChatEnabledStartNotification is a helper method for Client.ToggleVideoChatEnabledStartNotification
func (m MessageInviteVideoChatParticipants) ToggleVideoChatEnabledStartNotification(client *Client, enabledStartNotification bool) (*Ok, error) {
	return client.ToggleVideoChatEnabledStartNotification(m.GroupCallId, enabledStartNotification)
}

// ToggleVideoChatMuteNewParticipants is a helper method for Client.ToggleVideoChatMuteNewParticipants
func (m MessageInviteVideoChatParticipants) ToggleVideoChatMuteNewParticipants(client *Client, muteNewParticipants bool) (*Ok, error) {
	return client.ToggleVideoChatMuteNewParticipants(m.GroupCallId, muteNewParticipants)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessageInvoice) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, amount)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageInvoice) GetMainWebApp(client *Client, chatId int64, botUserId int64, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(chatId, botUserId, m.StartParameter, parameters)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageInvoice) GetWebAppLinkUrl(client *Client, chatId int64, botUserId int64, webAppShortName string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(chatId, botUserId, webAppShortName, m.StartParameter, allowWriteAccess, parameters)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessageInvoice) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, amount)
}

// ConfirmQrCodeAuthentication is a helper method for Client.ConfirmQrCodeAuthentication
func (m MessageLink) ConfirmQrCodeAuthentication(client *Client) (*Session, error) {
	return client.ConfirmQrCodeAuthentication(m.Link)
}

// DeleteBusinessChatLink is a helper method for Client.DeleteBusinessChatLink
func (m MessageLink) DeleteBusinessChatLink(client *Client) (*Ok, error) {
	return client.DeleteBusinessChatLink(m.Link)
}

// EditBusinessChatLink is a helper method for Client.EditBusinessChatLink
func (m MessageLink) EditBusinessChatLink(client *Client, linkInfo *InputBusinessChatLink) (*BusinessChatLink, error) {
	return client.EditBusinessChatLink(m.Link, linkInfo)
}

// GetDeepLinkInfo is a helper method for Client.GetDeepLinkInfo
func (m MessageLink) GetDeepLinkInfo(client *Client) (*DeepLinkInfo, error) {
	return client.GetDeepLinkInfo(m.Link)
}

// GetExternalLink is a helper method for Client.GetExternalLink
func (m MessageLink) GetExternalLink(client *Client, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetExternalLink(m.Link, allowWriteAccess)
}

// GetExternalLinkInfo is a helper method for Client.GetExternalLinkInfo
func (m MessageLink) GetExternalLinkInfo(client *Client) (*LoginUrlInfo, error) {
	return client.GetExternalLinkInfo(m.Link)
}

// GetInternalLinkType is a helper method for Client.GetInternalLinkType
func (m MessageLink) GetInternalLinkType(client *Client) (*InternalLinkType, error) {
	return client.GetInternalLinkType(m.Link)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (m MessageLink) SetBusinessAccountProfilePhoto(client *Client, businessConnectionId string, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(businessConnectionId, m.IsPublic, opts)
}

// SetProfilePhoto is a helper method for Client.SetProfilePhoto
func (m MessageLink) SetProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetProfilePhoto(photo, m.IsPublic)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageLinkInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageLinkInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m MessageLinkInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessageLinkInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessageLinkInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageLinkInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessageLinkInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessageLinkInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageLinkInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m MessageLinkInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m MessageLinkInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageLinkInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m MessageLinkInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (m MessageLinkInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (m MessageLinkInfo) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m MessageLinkInfo) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessageLinkInfo) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessageLinkInfo) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m MessageLinkInfo) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessageLinkInfo) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageLinkInfo) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageLinkInfo) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageLinkInfo) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageLinkInfo) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessageLinkInfo) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageLinkInfo) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m MessageLinkInfo) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m MessageLinkInfo) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m MessageLinkInfo) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m MessageLinkInfo) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (m MessageLinkInfo) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageLinkInfo) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessageLinkInfo) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m MessageLinkInfo) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m MessageLinkInfo) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageLinkInfo) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (m MessageLinkInfo) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m MessageLinkInfo) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m MessageLinkInfo) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageLinkInfo) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessageLinkInfo) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageLinkInfo) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessageLinkInfo) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessageLinkInfo) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessageLinkInfo) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageLinkInfo) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageLinkInfo) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageLinkInfo) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageLinkInfo) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessageLinkInfo) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageLinkInfo) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessageLinkInfo) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessageLinkInfo) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessageLinkInfo) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessageLinkInfo) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageLinkInfo) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m MessageLinkInfo) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessageLinkInfo) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessageLinkInfo) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m MessageLinkInfo) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m MessageLinkInfo) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m MessageLinkInfo) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m MessageLinkInfo) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m MessageLinkInfo) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m MessageLinkInfo) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m MessageLinkInfo) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m MessageLinkInfo) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m MessageLinkInfo) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageLinkInfo) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m MessageLinkInfo) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m MessageLinkInfo) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m MessageLinkInfo) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m MessageLinkInfo) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m MessageLinkInfo) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m MessageLinkInfo) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m MessageLinkInfo) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m MessageLinkInfo) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageLinkInfo) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (m MessageLinkInfo) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (m MessageLinkInfo) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessageLinkInfo) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (m MessageLinkInfo) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m MessageLinkInfo) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m MessageLinkInfo) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m MessageLinkInfo) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m MessageLinkInfo) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (m MessageLinkInfo) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m MessageLinkInfo) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m MessageLinkInfo) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageLinkInfo) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (m MessageLinkInfo) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m MessageLinkInfo) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m MessageLinkInfo) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m MessageLinkInfo) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m MessageLinkInfo) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageLinkInfo) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m MessageLinkInfo) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageLinkInfo) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageLinkInfo) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageLinkInfo) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m MessageLinkInfo) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageLinkInfo) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessageLinkInfo) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m MessageLinkInfo) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m MessageLinkInfo) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessageLinkInfo) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessageLinkInfo) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageLinkInfo) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageLinkInfo) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessageLinkInfo) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(m.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessageLinkInfo) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessageLinkInfo) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessageLinkInfo) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessageLinkInfo) GetMessageEmbeddingCode(client *Client, messageId int64) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, messageId, m.ForAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m MessageLinkInfo) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessageLinkInfo) GetMessageLink(client *Client, messageId int64, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, messageId, m.MediaTimestamp, m.ForAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessageLinkInfo) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessageLinkInfo) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessageLinkInfo) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessageLinkInfo) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (m MessageLinkInfo) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessageLinkInfo) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessageLinkInfo) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessageLinkInfo) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessageLinkInfo) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessageLinkInfo) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessageLinkInfo) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessageLinkInfo) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m MessageLinkInfo) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m MessageLinkInfo) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m MessageLinkInfo) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageLinkInfo) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageLinkInfo) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m MessageLinkInfo) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m MessageLinkInfo) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessageLinkInfo) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageLinkInfo) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (m MessageLinkInfo) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m MessageLinkInfo) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m MessageLinkInfo) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m MessageLinkInfo) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageLinkInfo) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m MessageLinkInfo) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m MessageLinkInfo) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessageLinkInfo) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MessageLinkInfo) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageLinkInfo) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageLinkInfo) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageLinkInfo) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m MessageLinkInfo) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessageLinkInfo) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m MessageLinkInfo) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m MessageLinkInfo) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m MessageLinkInfo) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageLinkInfo) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageLinkInfo) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessageLinkInfo) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessageLinkInfo) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m MessageLinkInfo) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m MessageLinkInfo) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessageLinkInfo) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessageLinkInfo) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m MessageLinkInfo) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m MessageLinkInfo) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m MessageLinkInfo) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m MessageLinkInfo) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m MessageLinkInfo) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m MessageLinkInfo) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m MessageLinkInfo) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m MessageLinkInfo) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageLinkInfo) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageLinkInfo) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessageLinkInfo) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageLinkInfo) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageLinkInfo) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m MessageLinkInfo) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m MessageLinkInfo) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m MessageLinkInfo) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (m MessageLinkInfo) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (m MessageLinkInfo) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (m MessageLinkInfo) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (m MessageLinkInfo) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageLinkInfo) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageLinkInfo) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m MessageLinkInfo) SendChatAction(client *Client, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, m.TopicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (m MessageLinkInfo) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (m MessageLinkInfo) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (m MessageLinkInfo) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageLinkInfo) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageLinkInfo) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessAccountProfilePhoto is a helper method for Client.SetBusinessAccountProfilePhoto
func (m MessageLinkInfo) SetBusinessAccountProfilePhoto(client *Client, businessConnectionId string, opts *SetBusinessAccountProfilePhotoOpts) (*Ok, error) {
	return client.SetBusinessAccountProfilePhoto(businessConnectionId, m.IsPublic, opts)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessageLinkInfo) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m MessageLinkInfo) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m MessageLinkInfo) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m MessageLinkInfo) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m MessageLinkInfo) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageLinkInfo) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m MessageLinkInfo) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m MessageLinkInfo) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageLinkInfo) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m MessageLinkInfo) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (m MessageLinkInfo) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m MessageLinkInfo) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m MessageLinkInfo) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m MessageLinkInfo) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageLinkInfo) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (m MessageLinkInfo) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m MessageLinkInfo) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageLinkInfo) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m MessageLinkInfo) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m MessageLinkInfo) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m MessageLinkInfo) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m MessageLinkInfo) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m MessageLinkInfo) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m MessageLinkInfo) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageLinkInfo) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m MessageLinkInfo) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageLinkInfo) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageLinkInfo) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessageLinkInfo) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessageLinkInfo) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessageLinkInfo) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m MessageLinkInfo) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m MessageLinkInfo) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessageLinkInfo) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, messageId, optionIds)
}

// SetProfilePhoto is a helper method for Client.SetProfilePhoto
func (m MessageLinkInfo) SetProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetProfilePhoto(photo, m.IsPublic)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageLinkInfo) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m MessageLinkInfo) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageLinkInfo) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageLinkInfo) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageLinkInfo) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessageLinkInfo) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessageLinkInfo) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessageLinkInfo) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m MessageLinkInfo) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m MessageLinkInfo) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m MessageLinkInfo) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m MessageLinkInfo) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m MessageLinkInfo) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m MessageLinkInfo) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m MessageLinkInfo) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m MessageLinkInfo) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m MessageLinkInfo) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageLinkInfo) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageLinkInfo) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageLinkInfo) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageLinkInfo) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessageLinkInfo) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (m MessageLinkInfo) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m MessageLinkInfo) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageLinkInfo) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessageLinkInfo) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m MessageLinkInfo) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (m MessageLinkInfo) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageLocation) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, messageId, m.LivePeriod, m.Heading, m.ProximityAlertRadius, opts)
}

// EditInlineMessageLiveLocation is a helper method for Client.EditInlineMessageLiveLocation
func (m MessageLocation) EditInlineMessageLiveLocation(client *Client, inlineMessageId string, opts *EditInlineMessageLiveLocationOpts) (*Ok, error) {
	return client.EditInlineMessageLiveLocation(inlineMessageId, m.LivePeriod, m.Heading, m.ProximityAlertRadius, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageLocation) EditMessageLiveLocation(client *Client, chatId int64, messageId int64, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, messageId, m.LivePeriod, m.Heading, m.ProximityAlertRadius, opts)
}

// GetCurrentWeather is a helper method for Client.GetCurrentWeather
func (m MessageLocation) GetCurrentWeather(client *Client) (*CurrentWeather, error) {
	return client.GetCurrentWeather(m.Location)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageLocation) GetMapThumbnailFile(client *Client, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	return client.GetMapThumbnailFile(m.Location, zoom, width, height, scale, chatId)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageOriginChannel) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageOriginChannel) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m MessageOriginChannel) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessageOriginChannel) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, m.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessageOriginChannel) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, m.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageOriginChannel) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessageOriginChannel) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, m.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessageOriginChannel) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, m.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageOriginChannel) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, m.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m MessageOriginChannel) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m MessageOriginChannel) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageOriginChannel) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, m.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m MessageOriginChannel) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (m MessageOriginChannel) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(m.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (m MessageOriginChannel) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (m MessageOriginChannel) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m MessageOriginChannel) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessageOriginChannel) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, m.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessageOriginChannel) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, m.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m MessageOriginChannel) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessageOriginChannel) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, m.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageOriginChannel) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageOriginChannel) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageOriginChannel) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageOriginChannel) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessageOriginChannel) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, m.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageOriginChannel) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, m.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m MessageOriginChannel) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m MessageOriginChannel) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m MessageOriginChannel) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m MessageOriginChannel) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (m MessageOriginChannel) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageOriginChannel) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessageOriginChannel) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, m.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m MessageOriginChannel) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m MessageOriginChannel) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageOriginChannel) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (m MessageOriginChannel) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m MessageOriginChannel) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m MessageOriginChannel) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageOriginChannel) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessageOriginChannel) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, m.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageOriginChannel) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessageOriginChannel) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessageOriginChannel) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, m.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessageOriginChannel) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageOriginChannel) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageOriginChannel) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageOriginChannel) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageOriginChannel) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessageOriginChannel) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, m.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageOriginChannel) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessageOriginChannel) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessageOriginChannel) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, m.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessageOriginChannel) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, m.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessageOriginChannel) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (m MessageOriginChannel) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, m.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageOriginChannel) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m MessageOriginChannel) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessageOriginChannel) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, m.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessageOriginChannel) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, m.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m MessageOriginChannel) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m MessageOriginChannel) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m MessageOriginChannel) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m MessageOriginChannel) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m MessageOriginChannel) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m MessageOriginChannel) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m MessageOriginChannel) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m MessageOriginChannel) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m MessageOriginChannel) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageOriginChannel) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m MessageOriginChannel) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m MessageOriginChannel) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m MessageOriginChannel) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m MessageOriginChannel) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m MessageOriginChannel) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m MessageOriginChannel) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m MessageOriginChannel) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m MessageOriginChannel) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageOriginChannel) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (m MessageOriginChannel) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (m MessageOriginChannel) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessageOriginChannel) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, m.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (m MessageOriginChannel) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m MessageOriginChannel) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m MessageOriginChannel) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m MessageOriginChannel) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m MessageOriginChannel) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (m MessageOriginChannel) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m MessageOriginChannel) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m MessageOriginChannel) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageOriginChannel) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (m MessageOriginChannel) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m MessageOriginChannel) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m MessageOriginChannel) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m MessageOriginChannel) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m MessageOriginChannel) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageOriginChannel) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m MessageOriginChannel) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageOriginChannel) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageOriginChannel) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageOriginChannel) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m MessageOriginChannel) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageOriginChannel) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, m.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessageOriginChannel) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, m.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m MessageOriginChannel) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m MessageOriginChannel) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessageOriginChannel) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, m.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessageOriginChannel) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, m.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageOriginChannel) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageOriginChannel) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessageOriginChannel) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(m.ChatId, m.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessageOriginChannel) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, m.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessageOriginChannel) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, m.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessageOriginChannel) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, m.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessageOriginChannel) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, m.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m MessageOriginChannel) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessageOriginChannel) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, m.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessageOriginChannel) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, m.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessageOriginChannel) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, m.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessageOriginChannel) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, m.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessageOriginChannel) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, m.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (m MessageOriginChannel) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessageOriginChannel) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, m.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessageOriginChannel) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, m.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessageOriginChannel) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, m.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessageOriginChannel) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, m.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessageOriginChannel) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, m.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessageOriginChannel) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, m.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessageOriginChannel) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, m.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m MessageOriginChannel) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m MessageOriginChannel) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m MessageOriginChannel) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageOriginChannel) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageOriginChannel) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m MessageOriginChannel) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m MessageOriginChannel) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessageOriginChannel) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, m.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageOriginChannel) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (m MessageOriginChannel) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m MessageOriginChannel) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m MessageOriginChannel) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m MessageOriginChannel) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageOriginChannel) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, m.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m MessageOriginChannel) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m MessageOriginChannel) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessageOriginChannel) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, m.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MessageOriginChannel) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageOriginChannel) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, m.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageOriginChannel) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageOriginChannel) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m MessageOriginChannel) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (m MessageOriginChannel) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(m.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessageOriginChannel) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, m.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m MessageOriginChannel) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m MessageOriginChannel) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m MessageOriginChannel) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageOriginChannel) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageOriginChannel) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessageOriginChannel) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, m.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessageOriginChannel) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, m.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m MessageOriginChannel) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m MessageOriginChannel) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessageOriginChannel) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, m.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessageOriginChannel) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, m.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m MessageOriginChannel) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m MessageOriginChannel) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m MessageOriginChannel) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m MessageOriginChannel) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m MessageOriginChannel) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m MessageOriginChannel) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m MessageOriginChannel) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m MessageOriginChannel) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageOriginChannel) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageOriginChannel) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessageOriginChannel) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, m.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageOriginChannel) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, m.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m MessageOriginChannel) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, m.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageOriginChannel) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m MessageOriginChannel) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m MessageOriginChannel) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m MessageOriginChannel) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (m MessageOriginChannel) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (m MessageOriginChannel) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (m MessageOriginChannel) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (m MessageOriginChannel) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageOriginChannel) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageOriginChannel) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m MessageOriginChannel) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (m MessageOriginChannel) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (m MessageOriginChannel) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (m MessageOriginChannel) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageOriginChannel) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageOriginChannel) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessageOriginChannel) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, m.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m MessageOriginChannel) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m MessageOriginChannel) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m MessageOriginChannel) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m MessageOriginChannel) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageOriginChannel) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m MessageOriginChannel) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m MessageOriginChannel) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageOriginChannel) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m MessageOriginChannel) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (m MessageOriginChannel) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m MessageOriginChannel) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m MessageOriginChannel) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m MessageOriginChannel) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageOriginChannel) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (m MessageOriginChannel) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m MessageOriginChannel) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageOriginChannel) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m MessageOriginChannel) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m MessageOriginChannel) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m MessageOriginChannel) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m MessageOriginChannel) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m MessageOriginChannel) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m MessageOriginChannel) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageOriginChannel) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m MessageOriginChannel) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageOriginChannel) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageOriginChannel) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, m.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessageOriginChannel) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, m.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessageOriginChannel) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, m.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessageOriginChannel) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, m.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m MessageOriginChannel) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m MessageOriginChannel) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessageOriginChannel) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, m.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageOriginChannel) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m MessageOriginChannel) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageOriginChannel) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, m.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageOriginChannel) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, m.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageOriginChannel) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessageOriginChannel) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, m.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessageOriginChannel) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, m.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessageOriginChannel) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, m.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m MessageOriginChannel) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m MessageOriginChannel) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m MessageOriginChannel) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m MessageOriginChannel) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m MessageOriginChannel) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m MessageOriginChannel) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m MessageOriginChannel) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m MessageOriginChannel) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m MessageOriginChannel) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageOriginChannel) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageOriginChannel) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageOriginChannel) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageOriginChannel) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessageOriginChannel) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, m.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (m MessageOriginChannel) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m MessageOriginChannel) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageOriginChannel) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessageOriginChannel) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, m.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m MessageOriginChannel) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (m MessageOriginChannel) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessagePaidMedia) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, m.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessagePaidMedia) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, m.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessagePaidMedia) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, m.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessagePaidMedia) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, m.StarCount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessagePaidMedia) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessagePaidMedia) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (m MessagePaidMedia) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, m.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessagePaidMedia) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessagePaidMedia) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, m.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessagePaidMedia) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, m.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (m MessagePaidMedia) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessagePaidMedia) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, m.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessagePaidMedia) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, m.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (m MessagePaidMedia) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, m.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (m MessagePaidMedia) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, m.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessagePaidMedia) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, m.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessagePaidMedia) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, m.StarCount)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessagePaidMessagePriceChanged) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, m.PaidMessageStarCount, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessagePaidMessagePriceChanged) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, m.PaidMessageStarCount)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessagePaidMessagePriceChanged) SendGroupCallMessage(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, m.PaidMessageStarCount)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessagePaidMessagePriceChanged) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, m.PaidMessageStarCount)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessagePaidMessagePriceChanged) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, m.PaidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessagePaidMessagePriceChanged) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, m.PaidMessageStarCount)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessagePaidMessagePriceChanged) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, protectContent, isRtmpStream, enableMessages, m.PaidMessageStarCount)
}

// AddPendingLiveStoryReaction is a helper method for Client.AddPendingLiveStoryReaction
func (m MessagePaidMessagesRefunded) AddPendingLiveStoryReaction(client *Client, groupCallId int32) (*Ok, error) {
	return client.AddPendingLiveStoryReaction(groupCallId, m.StarCount)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessagePaidMessagesRefunded) AddPendingPaidMessageReaction(client *Client, chatId int64, messageId int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, messageId, m.StarCount, opts)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessagePaidMessagesRefunded) BuyGiftUpgrade(client *Client, ownerId *MessageSender, prepaidUpgradeHash string) (*Ok, error) {
	return client.BuyGiftUpgrade(ownerId, prepaidUpgradeHash, m.StarCount)
}

// DropGiftOriginalDetails is a helper method for Client.DropGiftOriginalDetails
func (m MessagePaidMessagesRefunded) DropGiftOriginalDetails(client *Client, receivedGiftId string) (*Ok, error) {
	return client.DropGiftOriginalDetails(receivedGiftId, m.StarCount)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessagePaidMessagesRefunded) GetStarWithdrawalUrl(client *Client, ownerId *MessageSender, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(ownerId, m.StarCount, password)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessagePaidMessagesRefunded) GiftPremiumWithStars(client *Client, userId int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, m.StarCount, monthCount, text)
}

// IncreaseGiftAuctionBid is a helper method for Client.IncreaseGiftAuctionBid
func (m MessagePaidMessagesRefunded) IncreaseGiftAuctionBid(client *Client, giftId string) (*Ok, error) {
	return client.IncreaseGiftAuctionBid(giftId, m.StarCount)
}

// LaunchPrepaidGiveaway is a helper method for Client.LaunchPrepaidGiveaway
func (m MessagePaidMessagesRefunded) LaunchPrepaidGiveaway(client *Client, giveawayId string, parameters *GiveawayParameters, winnerCount int32) (*Ok, error) {
	return client.LaunchPrepaidGiveaway(giveawayId, parameters, winnerCount, m.StarCount)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessagePaidMessagesRefunded) PlaceGiftAuctionBid(client *Client, giftId string, userId int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, m.StarCount, userId, text, isPrivate)
}

// SearchPublicPosts is a helper method for Client.SearchPublicPosts
func (m MessagePaidMessagesRefunded) SearchPublicPosts(client *Client, query string, offset string, limit int32) (*FoundPublicPosts, error) {
	return client.SearchPublicPosts(query, offset, limit, m.StarCount)
}

// TransferBusinessAccountStars is a helper method for Client.TransferBusinessAccountStars
func (m MessagePaidMessagesRefunded) TransferBusinessAccountStars(client *Client, businessConnectionId string) (*Ok, error) {
	return client.TransferBusinessAccountStars(businessConnectionId, m.StarCount)
}

// TransferGift is a helper method for Client.TransferGift
func (m MessagePaidMessagesRefunded) TransferGift(client *Client, businessConnectionId string, receivedGiftId string, newOwnerId *MessageSender) (*Ok, error) {
	return client.TransferGift(businessConnectionId, receivedGiftId, newOwnerId, m.StarCount)
}

// UpgradeGift is a helper method for Client.UpgradeGift
func (m MessagePaidMessagesRefunded) UpgradeGift(client *Client, businessConnectionId string, receivedGiftId string, keepOriginalDetails bool) (*UpgradeGiftResult, error) {
	return client.UpgradeGift(businessConnectionId, receivedGiftId, keepOriginalDetails, m.StarCount)
}

// SendPassportAuthorizationForm is a helper method for Client.SendPassportAuthorizationForm
func (m MessagePassportDataSent) SendPassportAuthorizationForm(client *Client, authorizationFormId int32) (*Ok, error) {
	return client.SendPassportAuthorizationForm(authorizationFormId, m.Types)
}

// AddGiftCollectionGifts is a helper method for Client.AddGiftCollectionGifts
func (m MessagePaymentRefunded) AddGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.AddGiftCollectionGifts(m.OwnerId, collectionId, receivedGiftIds)
}

// BuyGiftUpgrade is a helper method for Client.BuyGiftUpgrade
func (m MessagePaymentRefunded) BuyGiftUpgrade(client *Client, prepaidUpgradeHash string, starCount int64) (*Ok, error) {
	return client.BuyGiftUpgrade(m.OwnerId, prepaidUpgradeHash, starCount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessagePaymentRefunded) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, amount)
}

// CreateGiftCollection is a helper method for Client.CreateGiftCollection
func (m MessagePaymentRefunded) CreateGiftCollection(client *Client, name string, receivedGiftIds []string) (*GiftCollection, error) {
	return client.CreateGiftCollection(m.OwnerId, name, receivedGiftIds)
}

// DeleteGiftCollection is a helper method for Client.DeleteGiftCollection
func (m MessagePaymentRefunded) DeleteGiftCollection(client *Client, collectionId int32) (*Ok, error) {
	return client.DeleteGiftCollection(m.OwnerId, collectionId)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (m MessagePaymentRefunded) EditUserStarSubscription(client *Client, userId int64, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(userId, m.TelegramPaymentChargeId, isCanceled)
}

// GetGiftCollections is a helper method for Client.GetGiftCollections
func (m MessagePaymentRefunded) GetGiftCollections(client *Client) (*GiftCollections, error) {
	return client.GetGiftCollections(m.OwnerId)
}

// GetReceivedGifts is a helper method for Client.GetReceivedGifts
func (m MessagePaymentRefunded) GetReceivedGifts(client *Client, businessConnectionId string, collectionId int32, excludeUnsaved bool, excludeSaved bool, excludeUnlimited bool, excludeUpgradable bool, excludeNonUpgradable bool, excludeUpgraded bool, excludeWithoutColors bool, excludeHosted bool, sortByPrice bool, offset string, limit int32) (*ReceivedGifts, error) {
	return client.GetReceivedGifts(businessConnectionId, m.OwnerId, collectionId, excludeUnsaved, excludeSaved, excludeUnlimited, excludeUpgradable, excludeNonUpgradable, excludeUpgraded, excludeWithoutColors, excludeHosted, sortByPrice, offset, limit)
}

// GetStarAdAccountUrl is a helper method for Client.GetStarAdAccountUrl
func (m MessagePaymentRefunded) GetStarAdAccountUrl(client *Client) (*HttpUrl, error) {
	return client.GetStarAdAccountUrl(m.OwnerId)
}

// GetStarRevenueStatistics is a helper method for Client.GetStarRevenueStatistics
func (m MessagePaymentRefunded) GetStarRevenueStatistics(client *Client, isDark bool) (*StarRevenueStatistics, error) {
	return client.GetStarRevenueStatistics(m.OwnerId, isDark)
}

// GetStarTransactions is a helper method for Client.GetStarTransactions
func (m MessagePaymentRefunded) GetStarTransactions(client *Client, subscriptionId string, offset string, limit int32, opts *GetStarTransactionsOpts) (*StarTransactions, error) {
	return client.GetStarTransactions(m.OwnerId, subscriptionId, offset, limit, opts)
}

// GetStarWithdrawalUrl is a helper method for Client.GetStarWithdrawalUrl
func (m MessagePaymentRefunded) GetStarWithdrawalUrl(client *Client, starCount int64, password string) (*HttpUrl, error) {
	return client.GetStarWithdrawalUrl(m.OwnerId, starCount, password)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (m MessagePaymentRefunded) RefundStarPayment(client *Client, userId int64) (*Ok, error) {
	return client.RefundStarPayment(userId, m.TelegramPaymentChargeId)
}

// RemoveGiftCollectionGifts is a helper method for Client.RemoveGiftCollectionGifts
func (m MessagePaymentRefunded) RemoveGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.RemoveGiftCollectionGifts(m.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollectionGifts is a helper method for Client.ReorderGiftCollectionGifts
func (m MessagePaymentRefunded) ReorderGiftCollectionGifts(client *Client, collectionId int32, receivedGiftIds []string) (*GiftCollection, error) {
	return client.ReorderGiftCollectionGifts(m.OwnerId, collectionId, receivedGiftIds)
}

// ReorderGiftCollections is a helper method for Client.ReorderGiftCollections
func (m MessagePaymentRefunded) ReorderGiftCollections(client *Client, collectionIds []int32) (*Ok, error) {
	return client.ReorderGiftCollections(m.OwnerId, collectionIds)
}

// SendGift is a helper method for Client.SendGift
func (m MessagePaymentRefunded) SendGift(client *Client, giftId string, text *FormattedText, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, m.OwnerId, text, isPrivate, payForUpgrade)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessagePaymentRefunded) SendGiftPurchaseOffer(client *Client, giftName string, price *GiftResalePrice, duration int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGiftPurchaseOffer(m.OwnerId, giftName, price, duration, paidMessageStarCount)
}

// SendResoldGift is a helper method for Client.SendResoldGift
func (m MessagePaymentRefunded) SendResoldGift(client *Client, giftName string, price *GiftResalePrice) (*GiftResaleResult, error) {
	return client.SendResoldGift(giftName, m.OwnerId, price)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessagePaymentRefunded) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, amount)
}

// SetGiftCollectionName is a helper method for Client.SetGiftCollectionName
func (m MessagePaymentRefunded) SetGiftCollectionName(client *Client, collectionId int32, name string) (*GiftCollection, error) {
	return client.SetGiftCollectionName(m.OwnerId, collectionId, name)
}

// SetPinnedGifts is a helper method for Client.SetPinnedGifts
func (m MessagePaymentRefunded) SetPinnedGifts(client *Client, receivedGiftIds []string) (*Ok, error) {
	return client.SetPinnedGifts(m.OwnerId, receivedGiftIds)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessagePaymentSuccessful) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, amount)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessagePaymentSuccessful) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, amount)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessagePaymentSuccessfulBot) CheckAuthenticationPremiumPurchase(client *Client, amount int64) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, amount)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (m MessagePaymentSuccessfulBot) EditUserStarSubscription(client *Client, userId int64, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(userId, m.TelegramPaymentChargeId, isCanceled)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (m MessagePaymentSuccessfulBot) RefundStarPayment(client *Client, userId int64) (*Ok, error) {
	return client.RefundStarPayment(userId, m.TelegramPaymentChargeId)
}

// SendPaymentForm is a helper method for Client.SendPaymentForm
func (m MessagePaymentSuccessfulBot) SendPaymentForm(client *Client, inputInvoice *InputInvoice, paymentFormId string, orderInfoId string, tipAmount int64, opts *SendPaymentFormOpts) (*PaymentResult, error) {
	return client.SendPaymentForm(inputInvoice, paymentFormId, orderInfoId, m.ShippingOptionId, tipAmount, opts)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessagePaymentSuccessfulBot) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool, amount int64) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, amount)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessagePhoto) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, messageId int64, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessagePhoto) EditBusinessStory(client *Client, storyPosterChatId int64, storyId int32, content *InputStoryContent, areas *InputStoryAreas, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(storyPosterChatId, storyId, content, areas, m.Caption, privacySettings)
}

// EditInlineMessageCaption is a helper method for Client.EditInlineMessageCaption
func (m MessagePhoto) EditInlineMessageCaption(client *Client, inlineMessageId string, opts *EditInlineMessageCaptionOpts) (*Ok, error) {
	return client.EditInlineMessageCaption(inlineMessageId, m.ShowCaptionAboveMedia, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessagePhoto) EditMessageCaption(client *Client, chatId int64, messageId int64, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, messageId, m.ShowCaptionAboveMedia, opts)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessagePinMessage) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, m.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessagePinMessage) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, m.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessagePinMessage) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, m.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessagePinMessage) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, m.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessagePinMessage) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, m.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessagePinMessage) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, m.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (m MessagePinMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(m.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessagePinMessage) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, m.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessagePinMessage) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, m.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessagePinMessage) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, m.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessagePinMessage) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, m.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessagePinMessage) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, m.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessagePinMessage) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, m.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessagePinMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessagePinMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, m.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessagePinMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessagePinMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, m.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessagePinMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, m.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessagePinMessage) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessagePinMessage) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessagePinMessage) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, m.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessagePinMessage) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessagePinMessage) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessagePinMessage) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, m.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessagePinMessage) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, m.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessagePinMessage) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, m.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (m MessagePinMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, m.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessagePinMessage) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, m.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessagePinMessage) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, m.MessageId, callbackQueryId)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessagePinMessage) GetChatMessagePosition(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, m.MessageId, opts)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessagePinMessage) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, m.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessagePinMessage) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, m.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessagePinMessage) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, m.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessagePinMessage) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, m.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessagePinMessage) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, m.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessagePinMessage) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, m.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessagePinMessage) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, m.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessagePinMessage) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, m.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessagePinMessage) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, m.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessagePinMessage) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, m.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessagePinMessage) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, m.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessagePinMessage) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, m.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessagePinMessage) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, m.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessagePinMessage) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, m.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessagePinMessage) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, m.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessagePinMessage) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, m.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessagePinMessage) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, m.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessagePinMessage) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, m.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessagePinMessage) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, m.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessagePinMessage) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, m.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessagePinMessage) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, m.MessageId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessagePinMessage) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, m.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessagePinMessage) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, m.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessagePinMessage) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, m.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessagePinMessage) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, m.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (m MessagePinMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(m.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessagePinMessage) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, m.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessagePinMessage) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, m.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessagePinMessage) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, m.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessagePinMessage) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, m.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessagePinMessage) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, m.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessagePinMessage) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, m.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessagePinMessage) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, m.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m MessagePinMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, m.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessagePinMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, m.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessagePinMessage) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, m.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessagePinMessage) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, m.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessagePinMessage) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, m.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessagePinMessage) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, m.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessagePinMessage) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, m.MessageId, optionIds)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessagePinMessage) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, m.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessagePinMessage) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, m.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessagePinMessage) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, m.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessagePinMessage) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, m.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessagePinMessage) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, m.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessagePinMessage) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, m.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessagePinMessage) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, m.MessageId)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessagePosition) AddChecklistTasks(client *Client, chatId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(chatId, m.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessagePosition) AddFileToDownloads(client *Client, fileId int32, chatId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, chatId, m.MessageId, priority)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessagePosition) AddMessageReaction(client *Client, chatId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(chatId, m.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessagePosition) AddOffer(client *Client, chatId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(chatId, m.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessagePosition) AddPendingPaidMessageReaction(client *Client, chatId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(chatId, m.MessageId, starCount, opts)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessagePosition) ApproveSuggestedPost(client *Client, chatId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, m.MessageId, sendDate)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (m MessagePosition) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(m.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessagePosition) ClickAnimatedEmojiMessage(client *Client, chatId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(chatId, m.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessagePosition) ClickChatSponsoredMessage(client *Client, chatId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(chatId, m.MessageId, isMediaClick, fromFullscreen)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessagePosition) CommitPendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(chatId, m.MessageId)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessagePosition) DeclineGroupCallInvitation(client *Client, chatId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(chatId, m.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessagePosition) DeclineSuggestedPost(client *Client, chatId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, m.MessageId, comment)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessagePosition) DeleteChatReplyMarkup(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(chatId, m.MessageId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessagePosition) EditBusinessMessageCaption(client *Client, businessConnectionId string, chatId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, chatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessagePosition) EditBusinessMessageChecklist(client *Client, businessConnectionId string, chatId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, chatId, m.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessagePosition) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, chatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessagePosition) EditBusinessMessageMedia(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, chatId, m.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessagePosition) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, chatId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, chatId, m.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessagePosition) EditBusinessMessageText(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, chatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessagePosition) EditMessageCaption(client *Client, chatId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(chatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessagePosition) EditMessageChecklist(client *Client, chatId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(chatId, m.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessagePosition) EditMessageLiveLocation(client *Client, chatId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(chatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessagePosition) EditMessageMedia(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(chatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessagePosition) EditMessageReplyMarkup(client *Client, chatId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(chatId, m.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessagePosition) EditMessageSchedulingState(client *Client, chatId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(chatId, m.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessagePosition) EditMessageText(client *Client, chatId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(chatId, m.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (m MessagePosition) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, m.MessageId, inputMessageContent)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessagePosition) GetCallbackQueryAnswer(client *Client, chatId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(chatId, m.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessagePosition) GetCallbackQueryMessage(client *Client, chatId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(chatId, m.MessageId, callbackQueryId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessagePosition) GetChatMessageByDate(client *Client, chatId int64) (*Message, error) {
	return client.GetChatMessageByDate(chatId, m.Date)
}

// GetChat is a helper method for Client.GetChatMessagePosition
func (m MessagePosition) GetChat(client *Client, chatId int64, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(chatId, filter, m.MessageId, opts)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessagePosition) GetDirectMessagesChatTopicMessageByDate(client *Client, chatId int64, topicId int64) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(chatId, topicId, m.Date)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessagePosition) GetGameHighScores(client *Client, chatId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, m.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessagePosition) GetGiveawayInfo(client *Client, chatId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(chatId, m.MessageId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessagePosition) GetLoginUrl(client *Client, chatId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(chatId, m.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessagePosition) GetLoginUrlInfo(client *Client, chatId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(chatId, m.MessageId, buttonId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessagePosition) GetMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetMessage(chatId, m.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessagePosition) GetMessageAddedReactions(client *Client, chatId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(chatId, m.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessagePosition) GetMessageAuthor(client *Client, chatId int64) (*User, error) {
	return client.GetMessageAuthor(chatId, m.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessagePosition) GetMessageAvailableReactions(client *Client, chatId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(chatId, m.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessagePosition) GetMessageEmbeddingCode(client *Client, chatId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(chatId, m.MessageId, forAlbum)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessagePosition) GetMessageLink(client *Client, chatId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(chatId, m.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessagePosition) GetMessageLocally(client *Client, chatId int64) (*Message, error) {
	return client.GetMessageLocally(chatId, m.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessagePosition) GetMessageProperties(client *Client, chatId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(chatId, m.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessagePosition) GetMessagePublicForwards(client *Client, chatId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(chatId, m.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessagePosition) GetMessageReadDate(client *Client, chatId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(chatId, m.MessageId)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessagePosition) GetMessageStatistics(client *Client, chatId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(chatId, m.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessagePosition) GetMessageThread(client *Client, chatId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(chatId, m.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessagePosition) GetMessageThreadHistory(client *Client, chatId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(chatId, m.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessagePosition) GetMessageViewers(client *Client, chatId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(chatId, m.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessagePosition) GetPaymentReceipt(client *Client, chatId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(chatId, m.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessagePosition) GetPollVoters(client *Client, chatId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(chatId, m.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessagePosition) GetRepliedMessage(client *Client, chatId int64) (*Message, error) {
	return client.GetRepliedMessage(chatId, m.MessageId)
}

// GetSavedMessagesTopicMessageByDate is a helper method for Client.GetSavedMessagesTopicMessageByDate
func (m MessagePosition) GetSavedMessagesTopicMessageByDate(client *Client, savedMessagesTopicId int64) (*Message, error) {
	return client.GetSavedMessagesTopicMessageByDate(savedMessagesTopicId, m.Date)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessagePosition) GetVideoMessageAdvertisements(client *Client, chatId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(chatId, m.MessageId)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessagePosition) MarkChecklistTasksAsDone(client *Client, chatId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(chatId, m.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessagePosition) OpenMessageContent(client *Client, chatId int64) (*Ok, error) {
	return client.OpenMessageContent(chatId, m.MessageId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessagePosition) PinChatMessage(client *Client, chatId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, m.MessageId, disableNotification, onlyForSelf)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (m MessagePosition) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(m.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessagePosition) RateSpeechRecognition(client *Client, chatId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(chatId, m.MessageId, isGood)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessagePosition) ReadBusinessMessage(client *Client, businessConnectionId string, chatId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, chatId, m.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessagePosition) RecognizeSpeech(client *Client, chatId int64) (*Ok, error) {
	return client.RecognizeSpeech(chatId, m.MessageId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessagePosition) RemoveMessageReaction(client *Client, chatId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(chatId, m.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessagePosition) RemovePendingPaidMessageReactions(client *Client, chatId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(chatId, m.MessageId)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessagePosition) ReportChatSponsoredMessage(client *Client, chatId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(chatId, m.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessagePosition) ReportMessageReactions(client *Client, chatId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(chatId, m.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m MessagePosition) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, m.MessageId)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessagePosition) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, chatId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, chatId, m.MessageId, isPinned)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessagePosition) SetGameScore(client *Client, chatId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, m.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessagePosition) SetMessageFactCheck(client *Client, chatId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(chatId, m.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessagePosition) SetMessageReactions(client *Client, chatId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(chatId, m.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessagePosition) SetPaidMessageReactionType(client *Client, chatId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(chatId, m.MessageId, typeField)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessagePosition) SetPollAnswer(client *Client, chatId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(chatId, m.MessageId, optionIds)
}

// SetStickerPositionInSet is a helper method for Client.SetStickerPositionInSet
func (m MessagePosition) SetStickerPositionInSet(client *Client, sticker *InputFile) (*Ok, error) {
	return client.SetStickerPositionInSet(sticker, m.Position)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessagePosition) ShareChatWithBot(client *Client, chatId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(chatId, m.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessagePosition) ShareUsersWithBot(client *Client, chatId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(chatId, m.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessagePosition) StopBusinessPoll(client *Client, businessConnectionId string, chatId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, chatId, m.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessagePosition) StopPoll(client *Client, chatId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(chatId, m.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessagePosition) SummarizeMessage(client *Client, chatId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(chatId, m.MessageId, translateToLanguageCode)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessagePosition) TranslateMessageText(client *Client, chatId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(chatId, m.MessageId, toLanguageCode)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessagePosition) UnpinChatMessage(client *Client, chatId int64) (*Ok, error) {
	return client.UnpinChatMessage(chatId, m.MessageId)
}

// ApplyPremiumGiftCode is a helper method for Client.ApplyPremiumGiftCode
func (m MessagePremiumGiftCode) ApplyPremiumGiftCode(client *Client) (*Ok, error) {
	return client.ApplyPremiumGiftCode(m.Code)
}

// CheckAuthenticationCode is a helper method for Client.CheckAuthenticationCode
func (m MessagePremiumGiftCode) CheckAuthenticationCode(client *Client) (*Ok, error) {
	return client.CheckAuthenticationCode(m.Code)
}

// CheckAuthenticationPremiumPurchase is a helper method for Client.CheckAuthenticationPremiumPurchase
func (m MessagePremiumGiftCode) CheckAuthenticationPremiumPurchase(client *Client) (*Ok, error) {
	return client.CheckAuthenticationPremiumPurchase(m.Currency, m.Amount)
}

// CheckEmailAddressVerificationCode is a helper method for Client.CheckEmailAddressVerificationCode
func (m MessagePremiumGiftCode) CheckEmailAddressVerificationCode(client *Client) (*Ok, error) {
	return client.CheckEmailAddressVerificationCode(m.Code)
}

// CheckPhoneNumberCode is a helper method for Client.CheckPhoneNumberCode
func (m MessagePremiumGiftCode) CheckPhoneNumberCode(client *Client) (*Ok, error) {
	return client.CheckPhoneNumberCode(m.Code)
}

// CheckPremiumGiftCode is a helper method for Client.CheckPremiumGiftCode
func (m MessagePremiumGiftCode) CheckPremiumGiftCode(client *Client) (*PremiumGiftCodeInfo, error) {
	return client.CheckPremiumGiftCode(m.Code)
}

// CheckRecoveryEmailAddressCode is a helper method for Client.CheckRecoveryEmailAddressCode
func (m MessagePremiumGiftCode) CheckRecoveryEmailAddressCode(client *Client) (*PasswordState, error) {
	return client.CheckRecoveryEmailAddressCode(m.Code)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (m MessagePremiumGiftCode) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(m.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (m MessagePremiumGiftCode) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(m.Text)
}

// GetPremiumInfoSticker is a helper method for Client.GetPremiumInfoSticker
func (m MessagePremiumGiftCode) GetPremiumInfoSticker(client *Client) (*Sticker, error) {
	return client.GetPremiumInfoSticker(m.MonthCount)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessagePremiumGiftCode) GiftPremiumWithStars(client *Client, userId int64, starCount int64) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, m.MonthCount, m.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (m MessagePremiumGiftCode) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(m.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessagePremiumGiftCode) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, m.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (m MessagePremiumGiftCode) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(m.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (m MessagePremiumGiftCode) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, m.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessagePremiumGiftCode) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, m.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessagePremiumGiftCode) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, m.Text)
}

// SetAuthenticationPremiumPurchaseTransaction is a helper method for Client.SetAuthenticationPremiumPurchaseTransaction
func (m MessagePremiumGiftCode) SetAuthenticationPremiumPurchaseTransaction(client *Client, transaction *StoreTransaction, isRestore bool) (*Ok, error) {
	return client.SetAuthenticationPremiumPurchaseTransaction(transaction, isRestore, m.Currency, m.Amount)
}

// TranslateText is a helper method for Client.TranslateText
func (m MessagePremiumGiftCode) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(m.Text, toLanguageCode)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageRefundedUpgradedGift) AddLocalMessage(client *Client, chatId int64, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, m.SenderId, disableNotification, inputMessageContent, opts)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageRefundedUpgradedGift) DeleteChatMessagesBySender(client *Client, chatId int64) (*Ok, error) {
	return client.DeleteChatMessagesBySender(chatId, m.SenderId)
}

// DeleteGroupCallMessagesBySender is a helper method for Client.DeleteGroupCallMessagesBySender
func (m MessageRefundedUpgradedGift) DeleteGroupCallMessagesBySender(client *Client, groupCallId int32, reportSpam bool) (*Ok, error) {
	return client.DeleteGroupCallMessagesBySender(groupCallId, m.SenderId, reportSpam)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageRefundedUpgradedGift) ReportMessageReactions(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ReportMessageReactions(chatId, messageId, m.SenderId)
}

// SetMessageSenderBlockList is a helper method for Client.SetMessageSenderBlockList
func (m MessageRefundedUpgradedGift) SetMessageSenderBlockList(client *Client, opts *SetMessageSenderBlockListOpts) (*Ok, error) {
	return client.SetMessageSenderBlockList(m.SenderId, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageReplyToMessage) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageReplyToMessage) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m MessageReplyToMessage) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessageReplyToMessage) AddChecklistTasks(client *Client, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, m.MessageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessageReplyToMessage) AddFileToDownloads(client *Client, fileId int32, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, m.MessageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageReplyToMessage) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessageReplyToMessage) AddMessageReaction(client *Client, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, m.MessageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessageReplyToMessage) AddOffer(client *Client, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, m.MessageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageReplyToMessage) AddPendingPaidMessageReaction(client *Client, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, m.MessageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m MessageReplyToMessage) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m MessageReplyToMessage) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageReplyToMessage) ApproveSuggestedPost(client *Client, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, m.MessageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m MessageReplyToMessage) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BlockMessageSenderFromReplies is a helper method for Client.BlockMessageSenderFromReplies
func (m MessageReplyToMessage) BlockMessageSenderFromReplies(client *Client, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	return client.BlockMessageSenderFromReplies(m.MessageId, deleteMessage, deleteAllMessages, reportSpam)
}

// BoostChat is a helper method for Client.BoostChat
func (m MessageReplyToMessage) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (m MessageReplyToMessage) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m MessageReplyToMessage) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessageReplyToMessage) ClickAnimatedEmojiMessage(client *Client) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, m.MessageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessageReplyToMessage) ClickChatSponsoredMessage(client *Client, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, m.MessageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m MessageReplyToMessage) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessageReplyToMessage) CommitPendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, m.MessageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageReplyToMessage) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageReplyToMessage) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageReplyToMessage) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageReplyToMessage) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessageReplyToMessage) DeclineGroupCallInvitation(client *Client) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, m.MessageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageReplyToMessage) DeclineSuggestedPost(client *Client, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, m.MessageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m MessageReplyToMessage) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m MessageReplyToMessage) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m MessageReplyToMessage) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m MessageReplyToMessage) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (m MessageReplyToMessage) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageReplyToMessage) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessageReplyToMessage) DeleteChatReplyMarkup(client *Client) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, m.MessageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m MessageReplyToMessage) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m MessageReplyToMessage) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageReplyToMessage) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (m MessageReplyToMessage) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m MessageReplyToMessage) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m MessageReplyToMessage) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageReplyToMessage) EditBusinessMessageCaption(client *Client, businessConnectionId string, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessageReplyToMessage) EditBusinessMessageChecklist(client *Client, businessConnectionId string, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, m.MessageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageReplyToMessage) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessageReplyToMessage) EditBusinessMessageMedia(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessageReplyToMessage) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, m.MessageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessageReplyToMessage) EditBusinessMessageText(client *Client, businessConnectionId string, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageReplyToMessage) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageReplyToMessage) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageReplyToMessage) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageReplyToMessage) EditMessageCaption(client *Client, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, m.MessageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessageReplyToMessage) EditMessageChecklist(client *Client, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, m.MessageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageReplyToMessage) EditMessageLiveLocation(client *Client, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, m.MessageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessageReplyToMessage) EditMessageMedia(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessageReplyToMessage) EditMessageReplyMarkup(client *Client, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, m.MessageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessageReplyToMessage) EditMessageSchedulingState(client *Client, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, m.MessageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessageReplyToMessage) EditMessageText(client *Client, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, m.MessageId, inputMessageContent, opts)
}

// EditQuickReplyMessage is a helper method for Client.EditQuickReplyMessage
func (m MessageReplyToMessage) EditQuickReplyMessage(client *Client, shortcutId int32, inputMessageContent *InputMessageContent) (*Ok, error) {
	return client.EditQuickReplyMessage(shortcutId, m.MessageId, inputMessageContent)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageReplyToMessage) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m MessageReplyToMessage) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessageReplyToMessage) GetCallbackQueryAnswer(client *Client, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, m.MessageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessageReplyToMessage) GetCallbackQueryMessage(client *Client, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, m.MessageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m MessageReplyToMessage) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m MessageReplyToMessage) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m MessageReplyToMessage) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m MessageReplyToMessage) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m MessageReplyToMessage) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m MessageReplyToMessage) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m MessageReplyToMessage) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m MessageReplyToMessage) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m MessageReplyToMessage) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageReplyToMessage) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m MessageReplyToMessage) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m MessageReplyToMessage) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m MessageReplyToMessage) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m MessageReplyToMessage) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m MessageReplyToMessage) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m MessageReplyToMessage) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m MessageReplyToMessage) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m MessageReplyToMessage) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageReplyToMessage) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (m MessageReplyToMessage) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (m MessageReplyToMessage) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessageReplyToMessage) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, m.MessageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (m MessageReplyToMessage) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m MessageReplyToMessage) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m MessageReplyToMessage) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m MessageReplyToMessage) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m MessageReplyToMessage) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (m MessageReplyToMessage) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m MessageReplyToMessage) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m MessageReplyToMessage) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageReplyToMessage) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (m MessageReplyToMessage) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m MessageReplyToMessage) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m MessageReplyToMessage) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m MessageReplyToMessage) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m MessageReplyToMessage) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageReplyToMessage) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m MessageReplyToMessage) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageReplyToMessage) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageReplyToMessage) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageReplyToMessage) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m MessageReplyToMessage) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageReplyToMessage) GetGameHighScores(client *Client, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, m.MessageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessageReplyToMessage) GetGiveawayInfo(client *Client) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, m.MessageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m MessageReplyToMessage) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m MessageReplyToMessage) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessageReplyToMessage) GetLoginUrl(client *Client, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, m.MessageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessageReplyToMessage) GetLoginUrlInfo(client *Client, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, m.MessageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageReplyToMessage) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageReplyToMessage) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessageReplyToMessage) GetMessage(client *Client) (*Message, error) {
	return client.GetMessage(m.ChatId, m.MessageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessageReplyToMessage) GetMessageAddedReactions(client *Client, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, m.MessageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessageReplyToMessage) GetMessageAuthor(client *Client) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, m.MessageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessageReplyToMessage) GetMessageAvailableReactions(client *Client, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, m.MessageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessageReplyToMessage) GetMessageEmbeddingCode(client *Client, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, m.MessageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m MessageReplyToMessage) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessageReplyToMessage) GetMessageLink(client *Client, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, m.MessageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessageReplyToMessage) GetMessageLocally(client *Client) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, m.MessageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessageReplyToMessage) GetMessageProperties(client *Client) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, m.MessageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessageReplyToMessage) GetMessagePublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, m.MessageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessageReplyToMessage) GetMessageReadDate(client *Client) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, m.MessageId)
}

// GetMessages is a helper method for Client.GetMessages
func (m MessageReplyToMessage) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessageReplyToMessage) GetMessageStatistics(client *Client, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, m.MessageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessageReplyToMessage) GetMessageThread(client *Client) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, m.MessageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessageReplyToMessage) GetMessageThreadHistory(client *Client, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, m.MessageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessageReplyToMessage) GetMessageViewers(client *Client) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, m.MessageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessageReplyToMessage) GetPaymentReceipt(client *Client) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, m.MessageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessageReplyToMessage) GetPollVoters(client *Client, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, m.MessageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessageReplyToMessage) GetRepliedMessage(client *Client) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, m.MessageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m MessageReplyToMessage) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m MessageReplyToMessage) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m MessageReplyToMessage) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageReplyToMessage) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageReplyToMessage) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m MessageReplyToMessage) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m MessageReplyToMessage) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessageReplyToMessage) GetVideoMessageAdvertisements(client *Client) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, m.MessageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageReplyToMessage) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (m MessageReplyToMessage) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m MessageReplyToMessage) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m MessageReplyToMessage) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m MessageReplyToMessage) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageReplyToMessage) MarkChecklistTasksAsDone(client *Client, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, m.MessageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m MessageReplyToMessage) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m MessageReplyToMessage) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessageReplyToMessage) OpenMessageContent(client *Client) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, m.MessageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MessageReplyToMessage) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageReplyToMessage) PinChatMessage(client *Client, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, m.MessageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageReplyToMessage) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageReplyToMessage) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m MessageReplyToMessage) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// ProcessGiftPurchaseOffer is a helper method for Client.ProcessGiftPurchaseOffer
func (m MessageReplyToMessage) ProcessGiftPurchaseOffer(client *Client, accept bool) (*Ok, error) {
	return client.ProcessGiftPurchaseOffer(m.MessageId, accept)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessageReplyToMessage) RateSpeechRecognition(client *Client, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, m.MessageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m MessageReplyToMessage) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m MessageReplyToMessage) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m MessageReplyToMessage) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageReplyToMessage) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageReplyToMessage) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessageReplyToMessage) ReadBusinessMessage(client *Client, businessConnectionId string) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, m.MessageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessageReplyToMessage) RecognizeSpeech(client *Client) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, m.MessageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m MessageReplyToMessage) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m MessageReplyToMessage) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessageReplyToMessage) RemoveMessageReaction(client *Client, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, m.MessageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessageReplyToMessage) RemovePendingPaidMessageReactions(client *Client) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, m.MessageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m MessageReplyToMessage) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m MessageReplyToMessage) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m MessageReplyToMessage) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m MessageReplyToMessage) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m MessageReplyToMessage) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m MessageReplyToMessage) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m MessageReplyToMessage) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m MessageReplyToMessage) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageReplyToMessage) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageReplyToMessage) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessageReplyToMessage) ReportChatSponsoredMessage(client *Client, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, m.MessageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageReplyToMessage) ReportMessageReactions(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, m.MessageId, senderId)
}

// ReportSupergroupAntiSpamFalsePositive is a helper method for Client.ReportSupergroupAntiSpamFalsePositive
func (m MessageReplyToMessage) ReportSupergroupAntiSpamFalsePositive(client *Client, supergroupId int64) (*Ok, error) {
	return client.ReportSupergroupAntiSpamFalsePositive(supergroupId, m.MessageId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageReplyToMessage) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m MessageReplyToMessage) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m MessageReplyToMessage) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m MessageReplyToMessage) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (m MessageReplyToMessage) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (m MessageReplyToMessage) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (m MessageReplyToMessage) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (m MessageReplyToMessage) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageReplyToMessage) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageReplyToMessage) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m MessageReplyToMessage) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (m MessageReplyToMessage) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (m MessageReplyToMessage) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (m MessageReplyToMessage) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageReplyToMessage) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageReplyToMessage) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessageReplyToMessage) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, m.MessageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m MessageReplyToMessage) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m MessageReplyToMessage) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m MessageReplyToMessage) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m MessageReplyToMessage) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageReplyToMessage) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m MessageReplyToMessage) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m MessageReplyToMessage) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageReplyToMessage) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m MessageReplyToMessage) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (m MessageReplyToMessage) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m MessageReplyToMessage) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m MessageReplyToMessage) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m MessageReplyToMessage) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageReplyToMessage) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (m MessageReplyToMessage) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m MessageReplyToMessage) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageReplyToMessage) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m MessageReplyToMessage) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m MessageReplyToMessage) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m MessageReplyToMessage) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m MessageReplyToMessage) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m MessageReplyToMessage) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m MessageReplyToMessage) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageReplyToMessage) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m MessageReplyToMessage) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageReplyToMessage) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageReplyToMessage) SetGameScore(client *Client, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, m.MessageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessageReplyToMessage) SetMessageFactCheck(client *Client, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, m.MessageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessageReplyToMessage) SetMessageReactions(client *Client, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, m.MessageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessageReplyToMessage) SetPaidMessageReactionType(client *Client, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, m.MessageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m MessageReplyToMessage) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m MessageReplyToMessage) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessageReplyToMessage) SetPollAnswer(client *Client, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, m.MessageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageReplyToMessage) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m MessageReplyToMessage) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageReplyToMessage) ShareChatWithBot(client *Client, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, m.MessageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageReplyToMessage) ShareUsersWithBot(client *Client, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, m.MessageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageReplyToMessage) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessageReplyToMessage) StopBusinessPoll(client *Client, businessConnectionId string, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, m.MessageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessageReplyToMessage) StopPoll(client *Client, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, m.MessageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessageReplyToMessage) SummarizeMessage(client *Client, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, m.MessageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m MessageReplyToMessage) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m MessageReplyToMessage) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m MessageReplyToMessage) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m MessageReplyToMessage) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m MessageReplyToMessage) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m MessageReplyToMessage) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m MessageReplyToMessage) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m MessageReplyToMessage) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m MessageReplyToMessage) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageReplyToMessage) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageReplyToMessage) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageReplyToMessage) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageReplyToMessage) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessageReplyToMessage) TranslateMessageText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, m.MessageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (m MessageReplyToMessage) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m MessageReplyToMessage) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageReplyToMessage) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessageReplyToMessage) UnpinChatMessage(client *Client) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, m.MessageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m MessageReplyToMessage) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (m MessageReplyToMessage) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// CloseStory is a helper method for Client.CloseStory
func (m MessageReplyToStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(m.StoryPosterChatId, m.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (m MessageReplyToStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(m.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (m MessageReplyToStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, m.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (m MessageReplyToStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(m.StoryPosterChatId, m.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageReplyToStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(m.StoryPosterChatId, m.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (m MessageReplyToStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(m.StoryPosterChatId, m.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (m MessageReplyToStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(m.StoryPosterChatId, m.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (m MessageReplyToStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(m.StoryPosterChatId, m.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (m MessageReplyToStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(m.StoryPosterChatId, m.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (m MessageReplyToStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(m.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (m MessageReplyToStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(m.StoryPosterChatId, m.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageReplyToStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, m.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (m MessageReplyToStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(m.StoryPosterChatId, m.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (m MessageReplyToStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(m.StoryPosterChatId, m.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (m MessageReplyToStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(m.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (m MessageReplyToStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(m.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (m MessageReplyToStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(m.StoryPosterChatId, m.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (m MessageReplyToStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(m.StoryPosterChatId, m.StoryId, isPostedToChatPage)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageSchedulingStateSendAtDate) ApproveSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, messageId, m.SendDate)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageSchedulingStateSendWhenVideoProcessed) ApproveSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, messageId, m.SendDate)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageSenderChat) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageSenderChat) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m MessageSenderChat) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessageSenderChat) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessageSenderChat) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageSenderChat) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessageSenderChat) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessageSenderChat) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageSenderChat) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m MessageSenderChat) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m MessageSenderChat) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageSenderChat) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m MessageSenderChat) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (m MessageSenderChat) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}

// CanPostStory is a helper method for Client.CanPostStory
func (m MessageSenderChat) CanPostStory(client *Client) (*CanPostStoryResult, error) {
	return client.CanPostStory(m.ChatId)
}

// CheckChatUsername is a helper method for Client.CheckChatUsername
func (m MessageSenderChat) CheckChatUsername(client *Client, username string) (*CheckChatUsernameResult, error) {
	return client.CheckChatUsername(m.ChatId, username)
}

// ClickAnimatedEmojiMessage is a helper method for Client.ClickAnimatedEmojiMessage
func (m MessageSenderChat) ClickAnimatedEmojiMessage(client *Client, messageId int64) (*Sticker, error) {
	return client.ClickAnimatedEmojiMessage(m.ChatId, messageId)
}

// ClickChatSponsoredMessage is a helper method for Client.ClickChatSponsoredMessage
func (m MessageSenderChat) ClickChatSponsoredMessage(client *Client, messageId int64, isMediaClick bool, fromFullscreen bool) (*Ok, error) {
	return client.ClickChatSponsoredMessage(m.ChatId, messageId, isMediaClick, fromFullscreen)
}

// CloseChat is a helper method for Client.CloseChat
func (m MessageSenderChat) CloseChat(client *Client) (*Ok, error) {
	return client.CloseChat(m.ChatId)
}

// CommitPendingPaidMessageReactions is a helper method for Client.CommitPendingPaidMessageReactions
func (m MessageSenderChat) CommitPendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.CommitPendingPaidMessageReactions(m.ChatId, messageId)
}

// CreateChatInviteLink is a helper method for Client.CreateChatInviteLink
func (m MessageSenderChat) CreateChatInviteLink(client *Client, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.CreateChatInviteLink(m.ChatId, name, expirationDate, memberLimit, createsJoinRequest)
}

// CreateChatSubscriptionInviteLink is a helper method for Client.CreateChatSubscriptionInviteLink
func (m MessageSenderChat) CreateChatSubscriptionInviteLink(client *Client, name string, subscriptionPricing *StarSubscriptionPricing) (*ChatInviteLink, error) {
	return client.CreateChatSubscriptionInviteLink(m.ChatId, name, subscriptionPricing)
}

// CreateForumTopic is a helper method for Client.CreateForumTopic
func (m MessageSenderChat) CreateForumTopic(client *Client, name string, isNameImplicit bool, icon *ForumTopicIcon) (*ForumTopicInfo, error) {
	return client.CreateForumTopic(m.ChatId, name, isNameImplicit, icon)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageSenderChat) CreateVideoChat(client *Client, title string, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(m.ChatId, title, startDate, isRtmpStream)
}

// DeclineGroupCallInvitation is a helper method for Client.DeclineGroupCallInvitation
func (m MessageSenderChat) DeclineGroupCallInvitation(client *Client, messageId int64) (*Ok, error) {
	return client.DeclineGroupCallInvitation(m.ChatId, messageId)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageSenderChat) DeclineSuggestedPost(client *Client, messageId int64, comment string) (*Ok, error) {
	return client.DeclineSuggestedPost(m.ChatId, messageId, comment)
}

// DeleteAllRevokedChatInviteLinks is a helper method for Client.DeleteAllRevokedChatInviteLinks
func (m MessageSenderChat) DeleteAllRevokedChatInviteLinks(client *Client, creatorUserId int64) (*Ok, error) {
	return client.DeleteAllRevokedChatInviteLinks(m.ChatId, creatorUserId)
}

// DeleteChat is a helper method for Client.DeleteChat
func (m MessageSenderChat) DeleteChat(client *Client) (*Ok, error) {
	return client.DeleteChat(m.ChatId)
}

// DeleteChatBackground is a helper method for Client.DeleteChatBackground
func (m MessageSenderChat) DeleteChatBackground(client *Client, restorePrevious bool) (*Ok, error) {
	return client.DeleteChatBackground(m.ChatId, restorePrevious)
}

// DeleteChatHistory is a helper method for Client.DeleteChatHistory
func (m MessageSenderChat) DeleteChatHistory(client *Client, removeFromChatList bool, revoke bool) (*Ok, error) {
	return client.DeleteChatHistory(m.ChatId, removeFromChatList, revoke)
}

// DeleteChatMessagesByDate is a helper method for Client.DeleteChatMessagesByDate
func (m MessageSenderChat) DeleteChatMessagesByDate(client *Client, minDate int32, maxDate int32, revoke bool) (*Ok, error) {
	return client.DeleteChatMessagesByDate(m.ChatId, minDate, maxDate, revoke)
}

// DeleteChatMessagesBySender is a helper method for Client.DeleteChatMessagesBySender
func (m MessageSenderChat) DeleteChatMessagesBySender(client *Client, senderId *MessageSender) (*Ok, error) {
	return client.DeleteChatMessagesBySender(m.ChatId, senderId)
}

// DeleteChatReplyMarkup is a helper method for Client.DeleteChatReplyMarkup
func (m MessageSenderChat) DeleteChatReplyMarkup(client *Client, messageId int64) (*Ok, error) {
	return client.DeleteChatReplyMarkup(m.ChatId, messageId)
}

// DeleteDirectMessagesChatTopicHistory is a helper method for Client.DeleteDirectMessagesChatTopicHistory
func (m MessageSenderChat) DeleteDirectMessagesChatTopicHistory(client *Client, topicId int64) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicHistory(m.ChatId, topicId)
}

// DeleteDirectMessagesChatTopicMessagesByDate is a helper method for Client.DeleteDirectMessagesChatTopicMessagesByDate
func (m MessageSenderChat) DeleteDirectMessagesChatTopicMessagesByDate(client *Client, topicId int64, minDate int32, maxDate int32) (*Ok, error) {
	return client.DeleteDirectMessagesChatTopicMessagesByDate(m.ChatId, topicId, minDate, maxDate)
}

// DeleteForumTopic is a helper method for Client.DeleteForumTopic
func (m MessageSenderChat) DeleteForumTopic(client *Client, forumTopicId int32) (*Ok, error) {
	return client.DeleteForumTopic(m.ChatId, forumTopicId)
}

// DeleteMessages is a helper method for Client.DeleteMessages
func (m MessageSenderChat) DeleteMessages(client *Client, messageIds []int64, revoke bool) (*Ok, error) {
	return client.DeleteMessages(m.ChatId, messageIds, revoke)
}

// DeleteRevokedChatInviteLink is a helper method for Client.DeleteRevokedChatInviteLink
func (m MessageSenderChat) DeleteRevokedChatInviteLink(client *Client, inviteLink string) (*Ok, error) {
	return client.DeleteRevokedChatInviteLink(m.ChatId, inviteLink)
}

// DeleteStoryAlbum is a helper method for Client.DeleteStoryAlbum
func (m MessageSenderChat) DeleteStoryAlbum(client *Client, storyAlbumId int32) (*Ok, error) {
	return client.DeleteStoryAlbum(m.ChatId, storyAlbumId)
}

// EditBusinessMessageCaption is a helper method for Client.EditBusinessMessageCaption
func (m MessageSenderChat) EditBusinessMessageCaption(client *Client, businessConnectionId string, messageId int64, showCaptionAboveMedia bool, opts *EditBusinessMessageCaptionOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageCaption(businessConnectionId, m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditBusinessMessageChecklist is a helper method for Client.EditBusinessMessageChecklist
func (m MessageSenderChat) EditBusinessMessageChecklist(client *Client, businessConnectionId string, messageId int64, checklist *InputChecklist, opts *EditBusinessMessageChecklistOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageChecklist(businessConnectionId, m.ChatId, messageId, checklist, opts)
}

// EditBusinessMessageLiveLocation is a helper method for Client.EditBusinessMessageLiveLocation
func (m MessageSenderChat) EditBusinessMessageLiveLocation(client *Client, businessConnectionId string, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditBusinessMessageLiveLocationOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageLiveLocation(businessConnectionId, m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditBusinessMessageMedia is a helper method for Client.EditBusinessMessageMedia
func (m MessageSenderChat) EditBusinessMessageMedia(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageMediaOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageMedia(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditBusinessMessageReplyMarkup is a helper method for Client.EditBusinessMessageReplyMarkup
func (m MessageSenderChat) EditBusinessMessageReplyMarkup(client *Client, businessConnectionId string, messageId int64, opts *EditBusinessMessageReplyMarkupOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageReplyMarkup(businessConnectionId, m.ChatId, messageId, opts)
}

// EditBusinessMessageText is a helper method for Client.EditBusinessMessageText
func (m MessageSenderChat) EditBusinessMessageText(client *Client, businessConnectionId string, messageId int64, inputMessageContent *InputMessageContent, opts *EditBusinessMessageTextOpts) (*BusinessMessage, error) {
	return client.EditBusinessMessageText(businessConnectionId, m.ChatId, messageId, inputMessageContent, opts)
}

// EditChatInviteLink is a helper method for Client.EditChatInviteLink
func (m MessageSenderChat) EditChatInviteLink(client *Client, inviteLink string, name string, expirationDate int32, memberLimit int32, createsJoinRequest bool) (*ChatInviteLink, error) {
	return client.EditChatInviteLink(m.ChatId, inviteLink, name, expirationDate, memberLimit, createsJoinRequest)
}

// EditChatSubscriptionInviteLink is a helper method for Client.EditChatSubscriptionInviteLink
func (m MessageSenderChat) EditChatSubscriptionInviteLink(client *Client, inviteLink string, name string) (*ChatInviteLink, error) {
	return client.EditChatSubscriptionInviteLink(m.ChatId, inviteLink, name)
}

// EditForumTopic is a helper method for Client.EditForumTopic
func (m MessageSenderChat) EditForumTopic(client *Client, forumTopicId int32, name string, editIconCustomEmoji bool, iconCustomEmojiId string) (*Ok, error) {
	return client.EditForumTopic(m.ChatId, forumTopicId, name, editIconCustomEmoji, iconCustomEmojiId)
}

// EditMessageCaption is a helper method for Client.EditMessageCaption
func (m MessageSenderChat) EditMessageCaption(client *Client, messageId int64, showCaptionAboveMedia bool, opts *EditMessageCaptionOpts) (*Message, error) {
	return client.EditMessageCaption(m.ChatId, messageId, showCaptionAboveMedia, opts)
}

// EditMessageChecklist is a helper method for Client.EditMessageChecklist
func (m MessageSenderChat) EditMessageChecklist(client *Client, messageId int64, checklist *InputChecklist, opts *EditMessageChecklistOpts) (*Message, error) {
	return client.EditMessageChecklist(m.ChatId, messageId, checklist, opts)
}

// EditMessageLiveLocation is a helper method for Client.EditMessageLiveLocation
func (m MessageSenderChat) EditMessageLiveLocation(client *Client, messageId int64, livePeriod int32, heading int32, proximityAlertRadius int32, opts *EditMessageLiveLocationOpts) (*Message, error) {
	return client.EditMessageLiveLocation(m.ChatId, messageId, livePeriod, heading, proximityAlertRadius, opts)
}

// EditMessageMedia is a helper method for Client.EditMessageMedia
func (m MessageSenderChat) EditMessageMedia(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageMediaOpts) (*Message, error) {
	return client.EditMessageMedia(m.ChatId, messageId, inputMessageContent, opts)
}

// EditMessageReplyMarkup is a helper method for Client.EditMessageReplyMarkup
func (m MessageSenderChat) EditMessageReplyMarkup(client *Client, messageId int64, opts *EditMessageReplyMarkupOpts) (*Message, error) {
	return client.EditMessageReplyMarkup(m.ChatId, messageId, opts)
}

// EditMessageSchedulingState is a helper method for Client.EditMessageSchedulingState
func (m MessageSenderChat) EditMessageSchedulingState(client *Client, messageId int64, opts *EditMessageSchedulingStateOpts) (*Ok, error) {
	return client.EditMessageSchedulingState(m.ChatId, messageId, opts)
}

// EditMessageText is a helper method for Client.EditMessageText
func (m MessageSenderChat) EditMessageText(client *Client, messageId int64, inputMessageContent *InputMessageContent, opts *EditMessageTextOpts) (*Message, error) {
	return client.EditMessageText(m.ChatId, messageId, inputMessageContent, opts)
}

// ForwardMessages is a helper method for Client.ForwardMessages
func (m MessageSenderChat) ForwardMessages(client *Client, fromChatId int64, messageIds []int64, sendCopy bool, removeCaption bool, opts *ForwardMessagesOpts) (*Messages, error) {
	return client.ForwardMessages(m.ChatId, fromChatId, messageIds, sendCopy, removeCaption, opts)
}

// GetAllStickerEmojis is a helper method for Client.GetAllStickerEmojis
func (m MessageSenderChat) GetAllStickerEmojis(client *Client, stickerType *StickerType, query string, returnOnlyMainEmoji bool) (*Emojis, error) {
	return client.GetAllStickerEmojis(stickerType, query, m.ChatId, returnOnlyMainEmoji)
}

// GetCallbackQueryAnswer is a helper method for Client.GetCallbackQueryAnswer
func (m MessageSenderChat) GetCallbackQueryAnswer(client *Client, messageId int64, payload *CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	return client.GetCallbackQueryAnswer(m.ChatId, messageId, payload)
}

// GetCallbackQueryMessage is a helper method for Client.GetCallbackQueryMessage
func (m MessageSenderChat) GetCallbackQueryMessage(client *Client, messageId int64, callbackQueryId string) (*Message, error) {
	return client.GetCallbackQueryMessage(m.ChatId, messageId, callbackQueryId)
}

// GetChat is a helper method for Client.GetChat
func (m MessageSenderChat) GetChat(client *Client) (*Chat, error) {
	return client.GetChat(m.ChatId)
}

// GetChatActiveStories is a helper method for Client.GetChatActiveStories
func (m MessageSenderChat) GetChatActiveStories(client *Client) (*ChatActiveStories, error) {
	return client.GetChatActiveStories(m.ChatId)
}

// GetChatAdministrators is a helper method for Client.GetChatAdministrators
func (m MessageSenderChat) GetChatAdministrators(client *Client) (*ChatAdministrators, error) {
	return client.GetChatAdministrators(m.ChatId)
}

// GetChatArchivedStories is a helper method for Client.GetChatArchivedStories
func (m MessageSenderChat) GetChatArchivedStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatArchivedStories(m.ChatId, fromStoryId, limit)
}

// GetChatAvailableMessageSenders is a helper method for Client.GetChatAvailableMessageSenders
func (m MessageSenderChat) GetChatAvailableMessageSenders(client *Client) (*ChatMessageSenders, error) {
	return client.GetChatAvailableMessageSenders(m.ChatId)
}

// GetChatAvailablePaidMessageReactionSenders is a helper method for Client.GetChatAvailablePaidMessageReactionSenders
func (m MessageSenderChat) GetChatAvailablePaidMessageReactionSenders(client *Client) (*MessageSenders, error) {
	return client.GetChatAvailablePaidMessageReactionSenders(m.ChatId)
}

// GetChatBoostLink is a helper method for Client.GetChatBoostLink
func (m MessageSenderChat) GetChatBoostLink(client *Client) (*ChatBoostLink, error) {
	return client.GetChatBoostLink(m.ChatId)
}

// GetChatBoosts is a helper method for Client.GetChatBoosts
func (m MessageSenderChat) GetChatBoosts(client *Client, onlyGiftCodes bool, offset string, limit int32) (*FoundChatBoosts, error) {
	return client.GetChatBoosts(m.ChatId, onlyGiftCodes, offset, limit)
}

// GetChatBoostStatus is a helper method for Client.GetChatBoostStatus
func (m MessageSenderChat) GetChatBoostStatus(client *Client) (*ChatBoostStatus, error) {
	return client.GetChatBoostStatus(m.ChatId)
}

// GetChatEventLog is a helper method for Client.GetChatEventLog
func (m MessageSenderChat) GetChatEventLog(client *Client, query string, fromEventId string, limit int32, userIds []int64, opts *GetChatEventLogOpts) (*ChatEvents, error) {
	return client.GetChatEventLog(m.ChatId, query, fromEventId, limit, userIds, opts)
}

// GetChatHistory is a helper method for Client.GetChatHistory
func (m MessageSenderChat) GetChatHistory(client *Client, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	return client.GetChatHistory(m.ChatId, fromMessageId, offset, limit, onlyLocal)
}

// GetChatInviteLink is a helper method for Client.GetChatInviteLink
func (m MessageSenderChat) GetChatInviteLink(client *Client, inviteLink string) (*ChatInviteLink, error) {
	return client.GetChatInviteLink(m.ChatId, inviteLink)
}

// GetChatInviteLinkCounts is a helper method for Client.GetChatInviteLinkCounts
func (m MessageSenderChat) GetChatInviteLinkCounts(client *Client) (*ChatInviteLinkCounts, error) {
	return client.GetChatInviteLinkCounts(m.ChatId)
}

// GetChatInviteLinkMembers is a helper method for Client.GetChatInviteLinkMembers
func (m MessageSenderChat) GetChatInviteLinkMembers(client *Client, inviteLink string, onlyWithExpiredSubscription bool, limit int32, opts *GetChatInviteLinkMembersOpts) (*ChatInviteLinkMembers, error) {
	return client.GetChatInviteLinkMembers(m.ChatId, inviteLink, onlyWithExpiredSubscription, limit, opts)
}

// GetChatInviteLinks is a helper method for Client.GetChatInviteLinks
func (m MessageSenderChat) GetChatInviteLinks(client *Client, creatorUserId int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*ChatInviteLinks, error) {
	return client.GetChatInviteLinks(m.ChatId, creatorUserId, isRevoked, offsetDate, offsetInviteLink, limit)
}

// GetChatJoinRequests is a helper method for Client.GetChatJoinRequests
func (m MessageSenderChat) GetChatJoinRequests(client *Client, inviteLink string, query string, limit int32, opts *GetChatJoinRequestsOpts) (*ChatJoinRequests, error) {
	return client.GetChatJoinRequests(m.ChatId, inviteLink, query, limit, opts)
}

// GetChatListsToAddChat is a helper method for Client.GetChatListsToAddChat
func (m MessageSenderChat) GetChatListsToAddChat(client *Client) (*ChatLists, error) {
	return client.GetChatListsToAddChat(m.ChatId)
}

// GetChatMember is a helper method for Client.GetChatMember
func (m MessageSenderChat) GetChatMember(client *Client, memberId *MessageSender) (*ChatMember, error) {
	return client.GetChatMember(m.ChatId, memberId)
}

// GetChatMessageByDate is a helper method for Client.GetChatMessageByDate
func (m MessageSenderChat) GetChatMessageByDate(client *Client, date int32) (*Message, error) {
	return client.GetChatMessageByDate(m.ChatId, date)
}

// GetChatMessageCalendar is a helper method for Client.GetChatMessageCalendar
func (m MessageSenderChat) GetChatMessageCalendar(client *Client, filter *SearchMessagesFilter, fromMessageId int64, opts *GetChatMessageCalendarOpts) (*MessageCalendar, error) {
	return client.GetChatMessageCalendar(m.ChatId, filter, fromMessageId, opts)
}

// GetChatMessageCount is a helper method for Client.GetChatMessageCount
func (m MessageSenderChat) GetChatMessageCount(client *Client, filter *SearchMessagesFilter, returnLocal bool, opts *GetChatMessageCountOpts) (*Count, error) {
	return client.GetChatMessageCount(m.ChatId, filter, returnLocal, opts)
}

// GetChatMessagePosition is a helper method for Client.GetChatMessagePosition
func (m MessageSenderChat) GetChatMessagePosition(client *Client, filter *SearchMessagesFilter, messageId int64, opts *GetChatMessagePositionOpts) (*Count, error) {
	return client.GetChatMessagePosition(m.ChatId, filter, messageId, opts)
}

// GetChatPinnedMessage is a helper method for Client.GetChatPinnedMessage
func (m MessageSenderChat) GetChatPinnedMessage(client *Client) (*Message, error) {
	return client.GetChatPinnedMessage(m.ChatId)
}

// GetChatPostedToChatPageStories is a helper method for Client.GetChatPostedToChatPageStories
func (m MessageSenderChat) GetChatPostedToChatPageStories(client *Client, fromStoryId int32, limit int32) (*Stories, error) {
	return client.GetChatPostedToChatPageStories(m.ChatId, fromStoryId, limit)
}

// GetChatRevenueStatistics is a helper method for Client.GetChatRevenueStatistics
func (m MessageSenderChat) GetChatRevenueStatistics(client *Client, isDark bool) (*ChatRevenueStatistics, error) {
	return client.GetChatRevenueStatistics(m.ChatId, isDark)
}

// GetChatRevenueTransactions is a helper method for Client.GetChatRevenueTransactions
func (m MessageSenderChat) GetChatRevenueTransactions(client *Client, offset string, limit int32) (*ChatRevenueTransactions, error) {
	return client.GetChatRevenueTransactions(m.ChatId, offset, limit)
}

// GetChatRevenueWithdrawalUrl is a helper method for Client.GetChatRevenueWithdrawalUrl
func (m MessageSenderChat) GetChatRevenueWithdrawalUrl(client *Client, password string) (*HttpUrl, error) {
	return client.GetChatRevenueWithdrawalUrl(m.ChatId, password)
}

// GetChatScheduledMessages is a helper method for Client.GetChatScheduledMessages
func (m MessageSenderChat) GetChatScheduledMessages(client *Client) (*Messages, error) {
	return client.GetChatScheduledMessages(m.ChatId)
}

// GetChatSimilarChatCount is a helper method for Client.GetChatSimilarChatCount
func (m MessageSenderChat) GetChatSimilarChatCount(client *Client, returnLocal bool) (*Count, error) {
	return client.GetChatSimilarChatCount(m.ChatId, returnLocal)
}

// GetChatSimilarChats is a helper method for Client.GetChatSimilarChats
func (m MessageSenderChat) GetChatSimilarChats(client *Client) (*Chats, error) {
	return client.GetChatSimilarChats(m.ChatId)
}

// GetChatSparseMessagePositions is a helper method for Client.GetChatSparseMessagePositions
func (m MessageSenderChat) GetChatSparseMessagePositions(client *Client, filter *SearchMessagesFilter, fromMessageId int64, limit int32, savedMessagesTopicId int64) (*MessagePositions, error) {
	return client.GetChatSparseMessagePositions(m.ChatId, filter, fromMessageId, limit, savedMessagesTopicId)
}

// GetChatSponsoredMessages is a helper method for Client.GetChatSponsoredMessages
func (m MessageSenderChat) GetChatSponsoredMessages(client *Client) (*SponsoredMessages, error) {
	return client.GetChatSponsoredMessages(m.ChatId)
}

// GetChatStatistics is a helper method for Client.GetChatStatistics
func (m MessageSenderChat) GetChatStatistics(client *Client, isDark bool) (*ChatStatistics, error) {
	return client.GetChatStatistics(m.ChatId, isDark)
}

// GetChatStoryAlbums is a helper method for Client.GetChatStoryAlbums
func (m MessageSenderChat) GetChatStoryAlbums(client *Client) (*StoryAlbums, error) {
	return client.GetChatStoryAlbums(m.ChatId)
}

// GetDirectMessagesChatTopic is a helper method for Client.GetDirectMessagesChatTopic
func (m MessageSenderChat) GetDirectMessagesChatTopic(client *Client, topicId int64) (*DirectMessagesChatTopic, error) {
	return client.GetDirectMessagesChatTopic(m.ChatId, topicId)
}

// GetDirectMessagesChatTopicHistory is a helper method for Client.GetDirectMessagesChatTopicHistory
func (m MessageSenderChat) GetDirectMessagesChatTopicHistory(client *Client, topicId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetDirectMessagesChatTopicHistory(m.ChatId, topicId, fromMessageId, offset, limit)
}

// GetDirectMessagesChatTopicMessageByDate is a helper method for Client.GetDirectMessagesChatTopicMessageByDate
func (m MessageSenderChat) GetDirectMessagesChatTopicMessageByDate(client *Client, topicId int64, date int32) (*Message, error) {
	return client.GetDirectMessagesChatTopicMessageByDate(m.ChatId, topicId, date)
}

// GetDirectMessagesChatTopicRevenue is a helper method for Client.GetDirectMessagesChatTopicRevenue
func (m MessageSenderChat) GetDirectMessagesChatTopicRevenue(client *Client, topicId int64) (*StarCount, error) {
	return client.GetDirectMessagesChatTopicRevenue(m.ChatId, topicId)
}

// GetForumTopic is a helper method for Client.GetForumTopic
func (m MessageSenderChat) GetForumTopic(client *Client, forumTopicId int32) (*ForumTopic, error) {
	return client.GetForumTopic(m.ChatId, forumTopicId)
}

// GetForumTopicHistory is a helper method for Client.GetForumTopicHistory
func (m MessageSenderChat) GetForumTopicHistory(client *Client, forumTopicId int32, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetForumTopicHistory(m.ChatId, forumTopicId, fromMessageId, offset, limit)
}

// GetForumTopicLink is a helper method for Client.GetForumTopicLink
func (m MessageSenderChat) GetForumTopicLink(client *Client, forumTopicId int32) (*MessageLink, error) {
	return client.GetForumTopicLink(m.ChatId, forumTopicId)
}

// GetForumTopics is a helper method for Client.GetForumTopics
func (m MessageSenderChat) GetForumTopics(client *Client, query string, offsetDate int32, offsetMessageId int64, offsetForumTopicId int32, limit int32) (*ForumTopics, error) {
	return client.GetForumTopics(m.ChatId, query, offsetDate, offsetMessageId, offsetForumTopicId, limit)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageSenderChat) GetGameHighScores(client *Client, messageId int64, userId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(m.ChatId, messageId, userId)
}

// GetGiveawayInfo is a helper method for Client.GetGiveawayInfo
func (m MessageSenderChat) GetGiveawayInfo(client *Client, messageId int64) (*GiveawayInfo, error) {
	return client.GetGiveawayInfo(m.ChatId, messageId)
}

// GetInlineQueryResults is a helper method for Client.GetInlineQueryResults
func (m MessageSenderChat) GetInlineQueryResults(client *Client, botUserId int64, query string, offset string, opts *GetInlineQueryResultsOpts) (*InlineQueryResults, error) {
	return client.GetInlineQueryResults(botUserId, m.ChatId, query, offset, opts)
}

// GetLiveStoryRtmpUrl is a helper method for Client.GetLiveStoryRtmpUrl
func (m MessageSenderChat) GetLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetLiveStoryRtmpUrl(m.ChatId)
}

// GetLoginUrl is a helper method for Client.GetLoginUrl
func (m MessageSenderChat) GetLoginUrl(client *Client, messageId int64, buttonId int64, allowWriteAccess bool) (*HttpUrl, error) {
	return client.GetLoginUrl(m.ChatId, messageId, buttonId, allowWriteAccess)
}

// GetLoginUrlInfo is a helper method for Client.GetLoginUrlInfo
func (m MessageSenderChat) GetLoginUrlInfo(client *Client, messageId int64, buttonId int64) (*LoginUrlInfo, error) {
	return client.GetLoginUrlInfo(m.ChatId, messageId, buttonId)
}

// GetMainWebApp is a helper method for Client.GetMainWebApp
func (m MessageSenderChat) GetMainWebApp(client *Client, botUserId int64, startParameter string, parameters *WebAppOpenParameters) (*MainWebApp, error) {
	return client.GetMainWebApp(m.ChatId, botUserId, startParameter, parameters)
}

// GetMapThumbnailFile is a helper method for Client.GetMapThumbnailFile
func (m MessageSenderChat) GetMapThumbnailFile(client *Client, location *Location, zoom int32, width int32, height int32, scale int32) (*File, error) {
	return client.GetMapThumbnailFile(location, zoom, width, height, scale, m.ChatId)
}

// GetMessage is a helper method for Client.GetMessage
func (m MessageSenderChat) GetMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetMessage(m.ChatId, messageId)
}

// GetMessageAddedReactions is a helper method for Client.GetMessageAddedReactions
func (m MessageSenderChat) GetMessageAddedReactions(client *Client, messageId int64, offset string, limit int32, opts *GetMessageAddedReactionsOpts) (*AddedReactions, error) {
	return client.GetMessageAddedReactions(m.ChatId, messageId, offset, limit, opts)
}

// GetMessageAuthor is a helper method for Client.GetMessageAuthor
func (m MessageSenderChat) GetMessageAuthor(client *Client, messageId int64) (*User, error) {
	return client.GetMessageAuthor(m.ChatId, messageId)
}

// GetMessageAvailableReactions is a helper method for Client.GetMessageAvailableReactions
func (m MessageSenderChat) GetMessageAvailableReactions(client *Client, messageId int64, rowSize int32) (*AvailableReactions, error) {
	return client.GetMessageAvailableReactions(m.ChatId, messageId, rowSize)
}

// GetMessageEmbeddingCode is a helper method for Client.GetMessageEmbeddingCode
func (m MessageSenderChat) GetMessageEmbeddingCode(client *Client, messageId int64, forAlbum bool) (*Text, error) {
	return client.GetMessageEmbeddingCode(m.ChatId, messageId, forAlbum)
}

// GetMessageImportConfirmationText is a helper method for Client.GetMessageImportConfirmationText
func (m MessageSenderChat) GetMessageImportConfirmationText(client *Client) (*Text, error) {
	return client.GetMessageImportConfirmationText(m.ChatId)
}

// GetMessageLink is a helper method for Client.GetMessageLink
func (m MessageSenderChat) GetMessageLink(client *Client, messageId int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*MessageLink, error) {
	return client.GetMessageLink(m.ChatId, messageId, mediaTimestamp, forAlbum, inMessageThread)
}

// GetMessageLocally is a helper method for Client.GetMessageLocally
func (m MessageSenderChat) GetMessageLocally(client *Client, messageId int64) (*Message, error) {
	return client.GetMessageLocally(m.ChatId, messageId)
}

// GetMessageProperties is a helper method for Client.GetMessageProperties
func (m MessageSenderChat) GetMessageProperties(client *Client, messageId int64) (*MessageProperties, error) {
	return client.GetMessageProperties(m.ChatId, messageId)
}

// GetMessagePublicForwards is a helper method for Client.GetMessagePublicForwards
func (m MessageSenderChat) GetMessagePublicForwards(client *Client, messageId int64, offset string, limit int32) (*PublicForwards, error) {
	return client.GetMessagePublicForwards(m.ChatId, messageId, offset, limit)
}

// GetMessageReadDate is a helper method for Client.GetMessageReadDate
func (m MessageSenderChat) GetMessageReadDate(client *Client, messageId int64) (*MessageReadDate, error) {
	return client.GetMessageReadDate(m.ChatId, messageId)
}

// GetMessages is a helper method for Client.GetMessages
func (m MessageSenderChat) GetMessages(client *Client, messageIds []int64) (*Messages, error) {
	return client.GetMessages(m.ChatId, messageIds)
}

// GetMessageStatistics is a helper method for Client.GetMessageStatistics
func (m MessageSenderChat) GetMessageStatistics(client *Client, messageId int64, isDark bool) (*MessageStatistics, error) {
	return client.GetMessageStatistics(m.ChatId, messageId, isDark)
}

// GetMessageThread is a helper method for Client.GetMessageThread
func (m MessageSenderChat) GetMessageThread(client *Client, messageId int64) (*MessageThreadInfo, error) {
	return client.GetMessageThread(m.ChatId, messageId)
}

// GetMessageThreadHistory is a helper method for Client.GetMessageThreadHistory
func (m MessageSenderChat) GetMessageThreadHistory(client *Client, messageId int64, fromMessageId int64, offset int32, limit int32) (*Messages, error) {
	return client.GetMessageThreadHistory(m.ChatId, messageId, fromMessageId, offset, limit)
}

// GetMessageViewers is a helper method for Client.GetMessageViewers
func (m MessageSenderChat) GetMessageViewers(client *Client, messageId int64) (*MessageViewers, error) {
	return client.GetMessageViewers(m.ChatId, messageId)
}

// GetPaymentReceipt is a helper method for Client.GetPaymentReceipt
func (m MessageSenderChat) GetPaymentReceipt(client *Client, messageId int64) (*PaymentReceipt, error) {
	return client.GetPaymentReceipt(m.ChatId, messageId)
}

// GetPollVoters is a helper method for Client.GetPollVoters
func (m MessageSenderChat) GetPollVoters(client *Client, messageId int64, optionId int32, offset int32, limit int32) (*MessageSenders, error) {
	return client.GetPollVoters(m.ChatId, messageId, optionId, offset, limit)
}

// GetRepliedMessage is a helper method for Client.GetRepliedMessage
func (m MessageSenderChat) GetRepliedMessage(client *Client, messageId int64) (*Message, error) {
	return client.GetRepliedMessage(m.ChatId, messageId)
}

// GetStatisticalGraph is a helper method for Client.GetStatisticalGraph
func (m MessageSenderChat) GetStatisticalGraph(client *Client, token string, x int64) (*StatisticalGraph, error) {
	return client.GetStatisticalGraph(m.ChatId, token, x)
}

// GetStickers is a helper method for Client.GetStickers
func (m MessageSenderChat) GetStickers(client *Client, stickerType *StickerType, query string, limit int32) (*Stickers, error) {
	return client.GetStickers(stickerType, query, limit, m.ChatId)
}

// GetStoryAlbumStories is a helper method for Client.GetStoryAlbumStories
func (m MessageSenderChat) GetStoryAlbumStories(client *Client, storyAlbumId int32, offset int32, limit int32) (*Stories, error) {
	return client.GetStoryAlbumStories(m.ChatId, storyAlbumId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageSenderChat) GetStoryStatistics(client *Client, storyId int32, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(m.ChatId, storyId, isDark)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageSenderChat) GetUserChatBoosts(client *Client, userId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(m.ChatId, userId)
}

// GetVideoChatAvailableParticipants is a helper method for Client.GetVideoChatAvailableParticipants
func (m MessageSenderChat) GetVideoChatAvailableParticipants(client *Client) (*MessageSenders, error) {
	return client.GetVideoChatAvailableParticipants(m.ChatId)
}

// GetVideoChatRtmpUrl is a helper method for Client.GetVideoChatRtmpUrl
func (m MessageSenderChat) GetVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.GetVideoChatRtmpUrl(m.ChatId)
}

// GetVideoMessageAdvertisements is a helper method for Client.GetVideoMessageAdvertisements
func (m MessageSenderChat) GetVideoMessageAdvertisements(client *Client, messageId int64) (*VideoMessageAdvertisements, error) {
	return client.GetVideoMessageAdvertisements(m.ChatId, messageId)
}

// GetWebAppLinkUrl is a helper method for Client.GetWebAppLinkUrl
func (m MessageSenderChat) GetWebAppLinkUrl(client *Client, botUserId int64, webAppShortName string, startParameter string, allowWriteAccess bool, parameters *WebAppOpenParameters) (*HttpUrl, error) {
	return client.GetWebAppLinkUrl(m.ChatId, botUserId, webAppShortName, startParameter, allowWriteAccess, parameters)
}

// ImportMessages is a helper method for Client.ImportMessages
func (m MessageSenderChat) ImportMessages(client *Client, messageFile *InputFile, attachedFiles []*InputFile) (*Ok, error) {
	return client.ImportMessages(m.ChatId, messageFile, attachedFiles)
}

// JoinChat is a helper method for Client.JoinChat
func (m MessageSenderChat) JoinChat(client *Client) (*Ok, error) {
	return client.JoinChat(m.ChatId)
}

// LeaveChat is a helper method for Client.LeaveChat
func (m MessageSenderChat) LeaveChat(client *Client) (*Ok, error) {
	return client.LeaveChat(m.ChatId)
}

// LoadDirectMessagesChatTopics is a helper method for Client.LoadDirectMessagesChatTopics
func (m MessageSenderChat) LoadDirectMessagesChatTopics(client *Client, limit int32) (*Ok, error) {
	return client.LoadDirectMessagesChatTopics(m.ChatId, limit)
}

// MarkChecklistTasksAsDone is a helper method for Client.MarkChecklistTasksAsDone
func (m MessageSenderChat) MarkChecklistTasksAsDone(client *Client, messageId int64, markedAsDoneTaskIds []int32, markedAsNotDoneTaskIds []int32) (*Ok, error) {
	return client.MarkChecklistTasksAsDone(m.ChatId, messageId, markedAsDoneTaskIds, markedAsNotDoneTaskIds)
}

// OpenChat is a helper method for Client.OpenChat
func (m MessageSenderChat) OpenChat(client *Client) (*Ok, error) {
	return client.OpenChat(m.ChatId)
}

// OpenChatSimilarChat is a helper method for Client.OpenChatSimilarChat
func (m MessageSenderChat) OpenChatSimilarChat(client *Client, openedChatId int64) (*Ok, error) {
	return client.OpenChatSimilarChat(m.ChatId, openedChatId)
}

// OpenMessageContent is a helper method for Client.OpenMessageContent
func (m MessageSenderChat) OpenMessageContent(client *Client, messageId int64) (*Ok, error) {
	return client.OpenMessageContent(m.ChatId, messageId)
}

// OpenWebApp is a helper method for Client.OpenWebApp
func (m MessageSenderChat) OpenWebApp(client *Client, botUserId int64, url string, parameters *WebAppOpenParameters, opts *OpenWebAppOpts) (*WebAppInfo, error) {
	return client.OpenWebApp(m.ChatId, botUserId, url, parameters, opts)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageSenderChat) PinChatMessage(client *Client, messageId int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(m.ChatId, messageId, disableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageSenderChat) PostStory(client *Client, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, protectContent bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(m.ChatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, protectContent, opts)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageSenderChat) ProcessChatJoinRequest(client *Client, userId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(m.ChatId, userId, approve)
}

// ProcessChatJoinRequests is a helper method for Client.ProcessChatJoinRequests
func (m MessageSenderChat) ProcessChatJoinRequests(client *Client, inviteLink string, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequests(m.ChatId, inviteLink, approve)
}

// RateSpeechRecognition is a helper method for Client.RateSpeechRecognition
func (m MessageSenderChat) RateSpeechRecognition(client *Client, messageId int64, isGood bool) (*Ok, error) {
	return client.RateSpeechRecognition(m.ChatId, messageId, isGood)
}

// ReadAllChatMentions is a helper method for Client.ReadAllChatMentions
func (m MessageSenderChat) ReadAllChatMentions(client *Client) (*Ok, error) {
	return client.ReadAllChatMentions(m.ChatId)
}

// ReadAllChatReactions is a helper method for Client.ReadAllChatReactions
func (m MessageSenderChat) ReadAllChatReactions(client *Client) (*Ok, error) {
	return client.ReadAllChatReactions(m.ChatId)
}

// ReadAllDirectMessagesChatTopicReactions is a helper method for Client.ReadAllDirectMessagesChatTopicReactions
func (m MessageSenderChat) ReadAllDirectMessagesChatTopicReactions(client *Client, topicId int64) (*Ok, error) {
	return client.ReadAllDirectMessagesChatTopicReactions(m.ChatId, topicId)
}

// ReadAllForumTopicMentions is a helper method for Client.ReadAllForumTopicMentions
func (m MessageSenderChat) ReadAllForumTopicMentions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicMentions(m.ChatId, forumTopicId)
}

// ReadAllForumTopicReactions is a helper method for Client.ReadAllForumTopicReactions
func (m MessageSenderChat) ReadAllForumTopicReactions(client *Client, forumTopicId int32) (*Ok, error) {
	return client.ReadAllForumTopicReactions(m.ChatId, forumTopicId)
}

// ReadBusinessMessage is a helper method for Client.ReadBusinessMessage
func (m MessageSenderChat) ReadBusinessMessage(client *Client, businessConnectionId string, messageId int64) (*Ok, error) {
	return client.ReadBusinessMessage(businessConnectionId, m.ChatId, messageId)
}

// RecognizeSpeech is a helper method for Client.RecognizeSpeech
func (m MessageSenderChat) RecognizeSpeech(client *Client, messageId int64) (*Ok, error) {
	return client.RecognizeSpeech(m.ChatId, messageId)
}

// RemoveBusinessConnectedBotFromChat is a helper method for Client.RemoveBusinessConnectedBotFromChat
func (m MessageSenderChat) RemoveBusinessConnectedBotFromChat(client *Client) (*Ok, error) {
	return client.RemoveBusinessConnectedBotFromChat(m.ChatId)
}

// RemoveChatActionBar is a helper method for Client.RemoveChatActionBar
func (m MessageSenderChat) RemoveChatActionBar(client *Client) (*Ok, error) {
	return client.RemoveChatActionBar(m.ChatId)
}

// RemoveMessageReaction is a helper method for Client.RemoveMessageReaction
func (m MessageSenderChat) RemoveMessageReaction(client *Client, messageId int64, reactionType *ReactionType) (*Ok, error) {
	return client.RemoveMessageReaction(m.ChatId, messageId, reactionType)
}

// RemovePendingPaidMessageReactions is a helper method for Client.RemovePendingPaidMessageReactions
func (m MessageSenderChat) RemovePendingPaidMessageReactions(client *Client, messageId int64) (*Ok, error) {
	return client.RemovePendingPaidMessageReactions(m.ChatId, messageId)
}

// RemoveRecentlyFoundChat is a helper method for Client.RemoveRecentlyFoundChat
func (m MessageSenderChat) RemoveRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.RemoveRecentlyFoundChat(m.ChatId)
}

// RemoveStoryAlbumStories is a helper method for Client.RemoveStoryAlbumStories
func (m MessageSenderChat) RemoveStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.RemoveStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// RemoveTopChat is a helper method for Client.RemoveTopChat
func (m MessageSenderChat) RemoveTopChat(client *Client, category *TopChatCategory) (*Ok, error) {
	return client.RemoveTopChat(category, m.ChatId)
}

// ReorderStoryAlbums is a helper method for Client.ReorderStoryAlbums
func (m MessageSenderChat) ReorderStoryAlbums(client *Client, storyAlbumIds []int32) (*Ok, error) {
	return client.ReorderStoryAlbums(m.ChatId, storyAlbumIds)
}

// ReorderStoryAlbumStories is a helper method for Client.ReorderStoryAlbumStories
func (m MessageSenderChat) ReorderStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.ReorderStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ReplaceLiveStoryRtmpUrl is a helper method for Client.ReplaceLiveStoryRtmpUrl
func (m MessageSenderChat) ReplaceLiveStoryRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceLiveStoryRtmpUrl(m.ChatId)
}

// ReplacePrimaryChatInviteLink is a helper method for Client.ReplacePrimaryChatInviteLink
func (m MessageSenderChat) ReplacePrimaryChatInviteLink(client *Client) (*ChatInviteLink, error) {
	return client.ReplacePrimaryChatInviteLink(m.ChatId)
}

// ReplaceVideoChatRtmpUrl is a helper method for Client.ReplaceVideoChatRtmpUrl
func (m MessageSenderChat) ReplaceVideoChatRtmpUrl(client *Client) (*RtmpUrl, error) {
	return client.ReplaceVideoChatRtmpUrl(m.ChatId)
}

// ReportChat is a helper method for Client.ReportChat
func (m MessageSenderChat) ReportChat(client *Client, optionId string, messageIds []int64, text string) (*ReportChatResult, error) {
	return client.ReportChat(m.ChatId, optionId, messageIds, text)
}

// ReportChatPhoto is a helper method for Client.ReportChatPhoto
func (m MessageSenderChat) ReportChatPhoto(client *Client, fileId int32, reason *ReportReason, text string) (*Ok, error) {
	return client.ReportChatPhoto(m.ChatId, fileId, reason, text)
}

// ReportChatSponsoredMessage is a helper method for Client.ReportChatSponsoredMessage
func (m MessageSenderChat) ReportChatSponsoredMessage(client *Client, messageId int64, optionId string) (*ReportSponsoredResult, error) {
	return client.ReportChatSponsoredMessage(m.ChatId, messageId, optionId)
}

// ReportMessageReactions is a helper method for Client.ReportMessageReactions
func (m MessageSenderChat) ReportMessageReactions(client *Client, messageId int64, senderId *MessageSender) (*Ok, error) {
	return client.ReportMessageReactions(m.ChatId, messageId, senderId)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageSenderChat) ResendMessages(client *Client, messageIds []int64, paidMessageStarCount int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(m.ChatId, messageIds, paidMessageStarCount, opts)
}

// RevokeChatInviteLink is a helper method for Client.RevokeChatInviteLink
func (m MessageSenderChat) RevokeChatInviteLink(client *Client, inviteLink string) (*ChatInviteLinks, error) {
	return client.RevokeChatInviteLink(m.ChatId, inviteLink)
}

// SaveApplicationLogEvent is a helper method for Client.SaveApplicationLogEvent
func (m MessageSenderChat) SaveApplicationLogEvent(client *Client, typeField string, data *JsonValue) (*Ok, error) {
	return client.SaveApplicationLogEvent(typeField, m.ChatId, data)
}

// SearchChatMembers is a helper method for Client.SearchChatMembers
func (m MessageSenderChat) SearchChatMembers(client *Client, query string, limit int32, opts *SearchChatMembersOpts) (*ChatMembers, error) {
	return client.SearchChatMembers(m.ChatId, query, limit, opts)
}

// SearchChatMessages is a helper method for Client.SearchChatMessages
func (m MessageSenderChat) SearchChatMessages(client *Client, query string, fromMessageId int64, offset int32, limit int32, opts *SearchChatMessagesOpts) (*FoundChatMessages, error) {
	return client.SearchChatMessages(m.ChatId, query, fromMessageId, offset, limit, opts)
}

// SearchChatRecentLocationMessages is a helper method for Client.SearchChatRecentLocationMessages
func (m MessageSenderChat) SearchChatRecentLocationMessages(client *Client, limit int32) (*Messages, error) {
	return client.SearchChatRecentLocationMessages(m.ChatId, limit)
}

// SearchSecretMessages is a helper method for Client.SearchSecretMessages
func (m MessageSenderChat) SearchSecretMessages(client *Client, query string, offset string, limit int32, opts *SearchSecretMessagesOpts) (*FoundMessages, error) {
	return client.SearchSecretMessages(m.ChatId, query, offset, limit, opts)
}

// SendBotStartMessage is a helper method for Client.SendBotStartMessage
func (m MessageSenderChat) SendBotStartMessage(client *Client, botUserId int64, parameter string) (*Message, error) {
	return client.SendBotStartMessage(botUserId, m.ChatId, parameter)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageSenderChat) SendBusinessMessage(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageSenderChat) SendBusinessMessageAlbum(client *Client, businessConnectionId string, disableNotification bool, protectContent bool, effectId string, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, m.ChatId, disableNotification, protectContent, effectId, inputMessageContents, opts)
}

// SendChatAction is a helper method for Client.SendChatAction
func (m MessageSenderChat) SendChatAction(client *Client, topicId *MessageTopic, businessConnectionId string, opts *SendChatActionOpts) (*Ok, error) {
	return client.SendChatAction(m.ChatId, topicId, businessConnectionId, opts)
}

// SendInlineQueryResultMessage is a helper method for Client.SendInlineQueryResultMessage
func (m MessageSenderChat) SendInlineQueryResultMessage(client *Client, queryId string, resultId string, hideViaBot bool, opts *SendInlineQueryResultMessageOpts) (*Message, error) {
	return client.SendInlineQueryResultMessage(m.ChatId, queryId, resultId, hideViaBot, opts)
}

// SendMessage is a helper method for Client.SendMessage
func (m MessageSenderChat) SendMessage(client *Client, inputMessageContent *InputMessageContent, opts *SendMessageOpts) (*Message, error) {
	return client.SendMessage(m.ChatId, inputMessageContent, opts)
}

// SendMessageAlbum is a helper method for Client.SendMessageAlbum
func (m MessageSenderChat) SendMessageAlbum(client *Client, inputMessageContents []*InputMessageContent, opts *SendMessageAlbumOpts) (*Messages, error) {
	return client.SendMessageAlbum(m.ChatId, inputMessageContents, opts)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageSenderChat) SendQuickReplyShortcutMessages(client *Client, shortcutId int32, sendingId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(m.ChatId, shortcutId, sendingId)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageSenderChat) SendTextMessageDraft(client *Client, forumTopicId int32, draftId string, text *FormattedText) (*Ok, error) {
	return client.SendTextMessageDraft(m.ChatId, forumTopicId, draftId, text)
}

// SetBusinessMessageIsPinned is a helper method for Client.SetBusinessMessageIsPinned
func (m MessageSenderChat) SetBusinessMessageIsPinned(client *Client, businessConnectionId string, messageId int64, isPinned bool) (*Ok, error) {
	return client.SetBusinessMessageIsPinned(businessConnectionId, m.ChatId, messageId, isPinned)
}

// SetChatAccentColor is a helper method for Client.SetChatAccentColor
func (m MessageSenderChat) SetChatAccentColor(client *Client, accentColorId int32, backgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatAccentColor(m.ChatId, accentColorId, backgroundCustomEmojiId)
}

// SetChatActiveStoriesList is a helper method for Client.SetChatActiveStoriesList
func (m MessageSenderChat) SetChatActiveStoriesList(client *Client, storyList *StoryList) (*Ok, error) {
	return client.SetChatActiveStoriesList(m.ChatId, storyList)
}

// SetChatAffiliateProgram is a helper method for Client.SetChatAffiliateProgram
func (m MessageSenderChat) SetChatAffiliateProgram(client *Client, opts *SetChatAffiliateProgramOpts) (*Ok, error) {
	return client.SetChatAffiliateProgram(m.ChatId, opts)
}

// SetChatAvailableReactions is a helper method for Client.SetChatAvailableReactions
func (m MessageSenderChat) SetChatAvailableReactions(client *Client, availableReactions *ChatAvailableReactions) (*Ok, error) {
	return client.SetChatAvailableReactions(m.ChatId, availableReactions)
}

// SetChatBackground is a helper method for Client.SetChatBackground
func (m MessageSenderChat) SetChatBackground(client *Client, darkThemeDimming int32, onlyForSelf bool, opts *SetChatBackgroundOpts) (*Ok, error) {
	return client.SetChatBackground(m.ChatId, darkThemeDimming, onlyForSelf, opts)
}

// SetChatClientData is a helper method for Client.SetChatClientData
func (m MessageSenderChat) SetChatClientData(client *Client, clientData string) (*Ok, error) {
	return client.SetChatClientData(m.ChatId, clientData)
}

// SetChatDescription is a helper method for Client.SetChatDescription
func (m MessageSenderChat) SetChatDescription(client *Client, description string) (*Ok, error) {
	return client.SetChatDescription(m.ChatId, description)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageSenderChat) SetChatDirectMessagesGroup(client *Client, isEnabled bool, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(m.ChatId, isEnabled, paidMessageStarCount)
}

// SetChatDiscussionGroup is a helper method for Client.SetChatDiscussionGroup
func (m MessageSenderChat) SetChatDiscussionGroup(client *Client, discussionChatId int64) (*Ok, error) {
	return client.SetChatDiscussionGroup(m.ChatId, discussionChatId)
}

// SetChatDraftMessage is a helper method for Client.SetChatDraftMessage
func (m MessageSenderChat) SetChatDraftMessage(client *Client, opts *SetChatDraftMessageOpts) (*Ok, error) {
	return client.SetChatDraftMessage(m.ChatId, opts)
}

// SetChatEmojiStatus is a helper method for Client.SetChatEmojiStatus
func (m MessageSenderChat) SetChatEmojiStatus(client *Client, opts *SetChatEmojiStatusOpts) (*Ok, error) {
	return client.SetChatEmojiStatus(m.ChatId, opts)
}

// SetChatLocation is a helper method for Client.SetChatLocation
func (m MessageSenderChat) SetChatLocation(client *Client, location *ChatLocation) (*Ok, error) {
	return client.SetChatLocation(m.ChatId, location)
}

// SetChatMemberStatus is a helper method for Client.SetChatMemberStatus
func (m MessageSenderChat) SetChatMemberStatus(client *Client, memberId *MessageSender, status *ChatMemberStatus) (*Ok, error) {
	return client.SetChatMemberStatus(m.ChatId, memberId, status)
}

// SetChatMessageAutoDeleteTime is a helper method for Client.SetChatMessageAutoDeleteTime
func (m MessageSenderChat) SetChatMessageAutoDeleteTime(client *Client, messageAutoDeleteTime int32) (*Ok, error) {
	return client.SetChatMessageAutoDeleteTime(m.ChatId, messageAutoDeleteTime)
}

// SetChatMessageSender is a helper method for Client.SetChatMessageSender
func (m MessageSenderChat) SetChatMessageSender(client *Client, messageSenderId *MessageSender) (*Ok, error) {
	return client.SetChatMessageSender(m.ChatId, messageSenderId)
}

// SetChatNotificationSettings is a helper method for Client.SetChatNotificationSettings
func (m MessageSenderChat) SetChatNotificationSettings(client *Client, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetChatNotificationSettings(m.ChatId, notificationSettings)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageSenderChat) SetChatPaidMessageStarCount(client *Client, paidMessageStarCount int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(m.ChatId, paidMessageStarCount)
}

// SetChatPermissions is a helper method for Client.SetChatPermissions
func (m MessageSenderChat) SetChatPermissions(client *Client, permissions *ChatPermissions) (*Ok, error) {
	return client.SetChatPermissions(m.ChatId, permissions)
}

// SetChatPhoto is a helper method for Client.SetChatPhoto
func (m MessageSenderChat) SetChatPhoto(client *Client, opts *SetChatPhotoOpts) (*Ok, error) {
	return client.SetChatPhoto(m.ChatId, opts)
}

// SetChatPinnedStories is a helper method for Client.SetChatPinnedStories
func (m MessageSenderChat) SetChatPinnedStories(client *Client, storyIds []int32) (*Ok, error) {
	return client.SetChatPinnedStories(m.ChatId, storyIds)
}

// SetChatProfileAccentColor is a helper method for Client.SetChatProfileAccentColor
func (m MessageSenderChat) SetChatProfileAccentColor(client *Client, profileAccentColorId int32, profileBackgroundCustomEmojiId string) (*Ok, error) {
	return client.SetChatProfileAccentColor(m.ChatId, profileAccentColorId, profileBackgroundCustomEmojiId)
}

// SetChatSlowModeDelay is a helper method for Client.SetChatSlowModeDelay
func (m MessageSenderChat) SetChatSlowModeDelay(client *Client, slowModeDelay int32) (*Ok, error) {
	return client.SetChatSlowModeDelay(m.ChatId, slowModeDelay)
}

// SetChatTheme is a helper method for Client.SetChatTheme
func (m MessageSenderChat) SetChatTheme(client *Client, theme *InputChatTheme) (*Ok, error) {
	return client.SetChatTheme(m.ChatId, theme)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageSenderChat) SetChatTitle(client *Client, title string) (*Ok, error) {
	return client.SetChatTitle(m.ChatId, title)
}

// SetDirectMessagesChatTopicIsMarkedAsUnread is a helper method for Client.SetDirectMessagesChatTopicIsMarkedAsUnread
func (m MessageSenderChat) SetDirectMessagesChatTopicIsMarkedAsUnread(client *Client, topicId int64, isMarkedAsUnread bool) (*Ok, error) {
	return client.SetDirectMessagesChatTopicIsMarkedAsUnread(m.ChatId, topicId, isMarkedAsUnread)
}

// SetForumTopicNotificationSettings is a helper method for Client.SetForumTopicNotificationSettings
func (m MessageSenderChat) SetForumTopicNotificationSettings(client *Client, forumTopicId int32, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	return client.SetForumTopicNotificationSettings(m.ChatId, forumTopicId, notificationSettings)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageSenderChat) SetGameScore(client *Client, messageId int64, editMessage bool, userId int64, score int32, force bool) (*Message, error) {
	return client.SetGameScore(m.ChatId, messageId, editMessage, userId, score, force)
}

// SetMessageFactCheck is a helper method for Client.SetMessageFactCheck
func (m MessageSenderChat) SetMessageFactCheck(client *Client, messageId int64, opts *SetMessageFactCheckOpts) (*Ok, error) {
	return client.SetMessageFactCheck(m.ChatId, messageId, opts)
}

// SetMessageReactions is a helper method for Client.SetMessageReactions
func (m MessageSenderChat) SetMessageReactions(client *Client, messageId int64, reactionTypes []*ReactionType, isBig bool) (*Ok, error) {
	return client.SetMessageReactions(m.ChatId, messageId, reactionTypes, isBig)
}

// SetPaidMessageReactionType is a helper method for Client.SetPaidMessageReactionType
func (m MessageSenderChat) SetPaidMessageReactionType(client *Client, messageId int64, typeField *PaidReactionType) (*Ok, error) {
	return client.SetPaidMessageReactionType(m.ChatId, messageId, typeField)
}

// SetPersonalChat is a helper method for Client.SetPersonalChat
func (m MessageSenderChat) SetPersonalChat(client *Client) (*Ok, error) {
	return client.SetPersonalChat(m.ChatId)
}

// SetPinnedForumTopics is a helper method for Client.SetPinnedForumTopics
func (m MessageSenderChat) SetPinnedForumTopics(client *Client, forumTopicIds []int32) (*Ok, error) {
	return client.SetPinnedForumTopics(m.ChatId, forumTopicIds)
}

// SetPollAnswer is a helper method for Client.SetPollAnswer
func (m MessageSenderChat) SetPollAnswer(client *Client, messageId int64, optionIds []int32) (*Ok, error) {
	return client.SetPollAnswer(m.ChatId, messageId, optionIds)
}

// SetStoryAlbumName is a helper method for Client.SetStoryAlbumName
func (m MessageSenderChat) SetStoryAlbumName(client *Client, storyAlbumId int32, name string) (*StoryAlbum, error) {
	return client.SetStoryAlbumName(m.ChatId, storyAlbumId, name)
}

// SetVideoChatDefaultParticipant is a helper method for Client.SetVideoChatDefaultParticipant
func (m MessageSenderChat) SetVideoChatDefaultParticipant(client *Client, defaultParticipantId *MessageSender) (*Ok, error) {
	return client.SetVideoChatDefaultParticipant(m.ChatId, defaultParticipantId)
}

// ShareChatWithBot is a helper method for Client.ShareChatWithBot
func (m MessageSenderChat) ShareChatWithBot(client *Client, messageId int64, buttonId int32, sharedChatId int64, onlyCheck bool) (*Ok, error) {
	return client.ShareChatWithBot(m.ChatId, messageId, buttonId, sharedChatId, onlyCheck)
}

// ShareUsersWithBot is a helper method for Client.ShareUsersWithBot
func (m MessageSenderChat) ShareUsersWithBot(client *Client, messageId int64, buttonId int32, sharedUserIds []int64, onlyCheck bool) (*Ok, error) {
	return client.ShareUsersWithBot(m.ChatId, messageId, buttonId, sharedUserIds, onlyCheck)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageSenderChat) StartLiveStory(client *Client, privacySettings *StoryPrivacySettings, protectContent bool, isRtmpStream bool, enableMessages bool, paidMessageStarCount int64) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(m.ChatId, privacySettings, protectContent, isRtmpStream, enableMessages, paidMessageStarCount)
}

// StopBusinessPoll is a helper method for Client.StopBusinessPoll
func (m MessageSenderChat) StopBusinessPoll(client *Client, businessConnectionId string, messageId int64, opts *StopBusinessPollOpts) (*BusinessMessage, error) {
	return client.StopBusinessPoll(businessConnectionId, m.ChatId, messageId, opts)
}

// StopPoll is a helper method for Client.StopPoll
func (m MessageSenderChat) StopPoll(client *Client, messageId int64, opts *StopPollOpts) (*Ok, error) {
	return client.StopPoll(m.ChatId, messageId, opts)
}

// SummarizeMessage is a helper method for Client.SummarizeMessage
func (m MessageSenderChat) SummarizeMessage(client *Client, messageId int64, translateToLanguageCode string) (*FormattedText, error) {
	return client.SummarizeMessage(m.ChatId, messageId, translateToLanguageCode)
}

// ToggleBusinessConnectedBotChatIsPaused is a helper method for Client.ToggleBusinessConnectedBotChatIsPaused
func (m MessageSenderChat) ToggleBusinessConnectedBotChatIsPaused(client *Client, isPaused bool) (*Ok, error) {
	return client.ToggleBusinessConnectedBotChatIsPaused(m.ChatId, isPaused)
}

// ToggleChatDefaultDisableNotification is a helper method for Client.ToggleChatDefaultDisableNotification
func (m MessageSenderChat) ToggleChatDefaultDisableNotification(client *Client, defaultDisableNotification bool) (*Ok, error) {
	return client.ToggleChatDefaultDisableNotification(m.ChatId, defaultDisableNotification)
}

// ToggleChatGiftNotifications is a helper method for Client.ToggleChatGiftNotifications
func (m MessageSenderChat) ToggleChatGiftNotifications(client *Client, areEnabled bool) (*Ok, error) {
	return client.ToggleChatGiftNotifications(m.ChatId, areEnabled)
}

// ToggleChatHasProtectedContent is a helper method for Client.ToggleChatHasProtectedContent
func (m MessageSenderChat) ToggleChatHasProtectedContent(client *Client, hasProtectedContent bool) (*Ok, error) {
	return client.ToggleChatHasProtectedContent(m.ChatId, hasProtectedContent)
}

// ToggleChatIsMarkedAsUnread is a helper method for Client.ToggleChatIsMarkedAsUnread
func (m MessageSenderChat) ToggleChatIsMarkedAsUnread(client *Client, isMarkedAsUnread bool) (*Ok, error) {
	return client.ToggleChatIsMarkedAsUnread(m.ChatId, isMarkedAsUnread)
}

// ToggleChatIsPinned is a helper method for Client.ToggleChatIsPinned
func (m MessageSenderChat) ToggleChatIsPinned(client *Client, chatList *ChatList, isPinned bool) (*Ok, error) {
	return client.ToggleChatIsPinned(chatList, m.ChatId, isPinned)
}

// ToggleChatIsTranslatable is a helper method for Client.ToggleChatIsTranslatable
func (m MessageSenderChat) ToggleChatIsTranslatable(client *Client, isTranslatable bool) (*Ok, error) {
	return client.ToggleChatIsTranslatable(m.ChatId, isTranslatable)
}

// ToggleChatViewAsTopics is a helper method for Client.ToggleChatViewAsTopics
func (m MessageSenderChat) ToggleChatViewAsTopics(client *Client, viewAsTopics bool) (*Ok, error) {
	return client.ToggleChatViewAsTopics(m.ChatId, viewAsTopics)
}

// ToggleDirectMessagesChatTopicCanSendUnpaidMessages is a helper method for Client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages
func (m MessageSenderChat) ToggleDirectMessagesChatTopicCanSendUnpaidMessages(client *Client, topicId int64, canSendUnpaidMessages bool, refundPayments bool) (*Ok, error) {
	return client.ToggleDirectMessagesChatTopicCanSendUnpaidMessages(m.ChatId, topicId, canSendUnpaidMessages, refundPayments)
}

// ToggleForumTopicIsClosed is a helper method for Client.ToggleForumTopicIsClosed
func (m MessageSenderChat) ToggleForumTopicIsClosed(client *Client, forumTopicId int32, isClosed bool) (*Ok, error) {
	return client.ToggleForumTopicIsClosed(m.ChatId, forumTopicId, isClosed)
}

// ToggleForumTopicIsPinned is a helper method for Client.ToggleForumTopicIsPinned
func (m MessageSenderChat) ToggleForumTopicIsPinned(client *Client, forumTopicId int32, isPinned bool) (*Ok, error) {
	return client.ToggleForumTopicIsPinned(m.ChatId, forumTopicId, isPinned)
}

// ToggleGeneralForumTopicIsHidden is a helper method for Client.ToggleGeneralForumTopicIsHidden
func (m MessageSenderChat) ToggleGeneralForumTopicIsHidden(client *Client, isHidden bool) (*Ok, error) {
	return client.ToggleGeneralForumTopicIsHidden(m.ChatId, isHidden)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageSenderChat) TransferChatOwnership(client *Client, userId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(m.ChatId, userId, password)
}

// TranslateMessageText is a helper method for Client.TranslateMessageText
func (m MessageSenderChat) TranslateMessageText(client *Client, messageId int64, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateMessageText(m.ChatId, messageId, toLanguageCode)
}

// UnpinAllChatMessages is a helper method for Client.UnpinAllChatMessages
func (m MessageSenderChat) UnpinAllChatMessages(client *Client) (*Ok, error) {
	return client.UnpinAllChatMessages(m.ChatId)
}

// UnpinAllDirectMessagesChatTopicMessages is a helper method for Client.UnpinAllDirectMessagesChatTopicMessages
func (m MessageSenderChat) UnpinAllDirectMessagesChatTopicMessages(client *Client, topicId int64) (*Ok, error) {
	return client.UnpinAllDirectMessagesChatTopicMessages(m.ChatId, topicId)
}

// UnpinAllForumTopicMessages is a helper method for Client.UnpinAllForumTopicMessages
func (m MessageSenderChat) UnpinAllForumTopicMessages(client *Client, forumTopicId int32) (*Ok, error) {
	return client.UnpinAllForumTopicMessages(m.ChatId, forumTopicId)
}

// UnpinChatMessage is a helper method for Client.UnpinChatMessage
func (m MessageSenderChat) UnpinChatMessage(client *Client, messageId int64) (*Ok, error) {
	return client.UnpinChatMessage(m.ChatId, messageId)
}

// UpgradeBasicGroupChatToSupergroupChat is a helper method for Client.UpgradeBasicGroupChatToSupergroupChat
func (m MessageSenderChat) UpgradeBasicGroupChatToSupergroupChat(client *Client) (*Chat, error) {
	return client.UpgradeBasicGroupChatToSupergroupChat(m.ChatId)
}

// ViewMessages is a helper method for Client.ViewMessages
func (m MessageSenderChat) ViewMessages(client *Client, messageIds []int64, forceRead bool, opts *ViewMessagesOpts) (*Ok, error) {
	return client.ViewMessages(m.ChatId, messageIds, forceRead, opts)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageSenderUser) AddChatMember(client *Client, chatId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(chatId, m.UserId, forwardLimit)
}

// AddContact is a helper method for Client.AddContact
func (m MessageSenderUser) AddContact(client *Client, contact *ImportedContact, sharePhoneNumber bool) (*Ok, error) {
	return client.AddContact(m.UserId, contact, sharePhoneNumber)
}

// AddStickerToSet is a helper method for Client.AddStickerToSet
func (m MessageSenderUser) AddStickerToSet(client *Client, name string, sticker *InputSticker) (*Ok, error) {
	return client.AddStickerToSet(m.UserId, name, sticker)
}

// AllowUnpaidMessagesFromUser is a helper method for Client.AllowUnpaidMessagesFromUser
func (m MessageSenderUser) AllowUnpaidMessagesFromUser(client *Client, refundPayments bool) (*Ok, error) {
	return client.AllowUnpaidMessagesFromUser(m.UserId, refundPayments)
}

// CanSendMessageToUser is a helper method for Client.CanSendMessageToUser
func (m MessageSenderUser) CanSendMessageToUser(client *Client, onlyLocal bool) (*CanSendMessageToUserResult, error) {
	return client.CanSendMessageToUser(m.UserId, onlyLocal)
}

// CreateCall is a helper method for Client.CreateCall
func (m MessageSenderUser) CreateCall(client *Client, protocol *CallProtocol, isVideo bool) (*CallId, error) {
	return client.CreateCall(m.UserId, protocol, isVideo)
}

// CreateNewSecretChat is a helper method for Client.CreateNewSecretChat
func (m MessageSenderUser) CreateNewSecretChat(client *Client) (*Chat, error) {
	return client.CreateNewSecretChat(m.UserId)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageSenderUser) CreateNewStickerSet(client *Client, title string, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(m.UserId, title, name, stickerType, needsRepainting, stickers, source)
}

// CreatePrivateChat is a helper method for Client.CreatePrivateChat
func (m MessageSenderUser) CreatePrivateChat(client *Client, force bool) (*Chat, error) {
	return client.CreatePrivateChat(m.UserId, force)
}

// EditUserStarSubscription is a helper method for Client.EditUserStarSubscription
func (m MessageSenderUser) EditUserStarSubscription(client *Client, telegramPaymentChargeId string, isCanceled bool) (*Ok, error) {
	return client.EditUserStarSubscription(m.UserId, telegramPaymentChargeId, isCanceled)
}

// GetGameHighScores is a helper method for Client.GetGameHighScores
func (m MessageSenderUser) GetGameHighScores(client *Client, chatId int64, messageId int64) (*GameHighScores, error) {
	return client.GetGameHighScores(chatId, messageId, m.UserId)
}

// GetGroupsInCommon is a helper method for Client.GetGroupsInCommon
func (m MessageSenderUser) GetGroupsInCommon(client *Client, offsetChatId int64, limit int32) (*Chats, error) {
	return client.GetGroupsInCommon(m.UserId, offsetChatId, limit)
}

// GetInlineGameHighScores is a helper method for Client.GetInlineGameHighScores
func (m MessageSenderUser) GetInlineGameHighScores(client *Client, inlineMessageId string) (*GameHighScores, error) {
	return client.GetInlineGameHighScores(inlineMessageId, m.UserId)
}

// GetMenuButton is a helper method for Client.GetMenuButton
func (m MessageSenderUser) GetMenuButton(client *Client) (*BotMenuButton, error) {
	return client.GetMenuButton(m.UserId)
}

// GetPaidMessageRevenue is a helper method for Client.GetPaidMessageRevenue
func (m MessageSenderUser) GetPaidMessageRevenue(client *Client) (*StarCount, error) {
	return client.GetPaidMessageRevenue(m.UserId)
}

// GetStarGiftPaymentOptions is a helper method for Client.GetStarGiftPaymentOptions
func (m MessageSenderUser) GetStarGiftPaymentOptions(client *Client) (*StarPaymentOptions, error) {
	return client.GetStarGiftPaymentOptions(m.UserId)
}

// GetUser is a helper method for Client.GetUser
func (m MessageSenderUser) GetUser(client *Client) (*User, error) {
	return client.GetUser(m.UserId)
}

// GetUserChatBoosts is a helper method for Client.GetUserChatBoosts
func (m MessageSenderUser) GetUserChatBoosts(client *Client, chatId int64) (*FoundChatBoosts, error) {
	return client.GetUserChatBoosts(chatId, m.UserId)
}

// GetUserFullInfo is a helper method for Client.GetUserFullInfo
func (m MessageSenderUser) GetUserFullInfo(client *Client) (*UserFullInfo, error) {
	return client.GetUserFullInfo(m.UserId)
}

// GetUserProfileAudios is a helper method for Client.GetUserProfileAudios
func (m MessageSenderUser) GetUserProfileAudios(client *Client, offset int32, limit int32) (*Audios, error) {
	return client.GetUserProfileAudios(m.UserId, offset, limit)
}

// GetUserProfilePhotos is a helper method for Client.GetUserProfilePhotos
func (m MessageSenderUser) GetUserProfilePhotos(client *Client, offset int32, limit int32) (*ChatPhotos, error) {
	return client.GetUserProfilePhotos(m.UserId, offset, limit)
}

// GetUserSupportInfo is a helper method for Client.GetUserSupportInfo
func (m MessageSenderUser) GetUserSupportInfo(client *Client) (*UserSupportInfo, error) {
	return client.GetUserSupportInfo(m.UserId)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageSenderUser) GiftPremiumWithStars(client *Client, starCount int64, monthCount int32, text *FormattedText) (*Ok, error) {
	return client.GiftPremiumWithStars(m.UserId, starCount, monthCount, text)
}

// InviteGroupCallParticipant is a helper method for Client.InviteGroupCallParticipant
func (m MessageSenderUser) InviteGroupCallParticipant(client *Client, groupCallId int32, isVideo bool) (*InviteGroupCallParticipantResult, error) {
	return client.InviteGroupCallParticipant(groupCallId, m.UserId, isVideo)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageSenderUser) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, text *FormattedText, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, m.UserId, text, isPrivate)
}

// ProcessChatJoinRequest is a helper method for Client.ProcessChatJoinRequest
func (m MessageSenderUser) ProcessChatJoinRequest(client *Client, chatId int64, approve bool) (*Ok, error) {
	return client.ProcessChatJoinRequest(chatId, m.UserId, approve)
}

// RefundStarPayment is a helper method for Client.RefundStarPayment
func (m MessageSenderUser) RefundStarPayment(client *Client, telegramPaymentChargeId string) (*Ok, error) {
	return client.RefundStarPayment(m.UserId, telegramPaymentChargeId)
}

// ReplaceStickerInSet is a helper method for Client.ReplaceStickerInSet
func (m MessageSenderUser) ReplaceStickerInSet(client *Client, name string, oldSticker *InputFile, newSticker *InputSticker) (*Ok, error) {
	return client.ReplaceStickerInSet(m.UserId, name, oldSticker, newSticker)
}

// SavePreparedInlineMessage is a helper method for Client.SavePreparedInlineMessage
func (m MessageSenderUser) SavePreparedInlineMessage(client *Client, result *InputInlineQueryResult, chatTypes *TargetChatTypes) (*PreparedInlineMessageId, error) {
	return client.SavePreparedInlineMessage(m.UserId, result, chatTypes)
}

// SetGameScore is a helper method for Client.SetGameScore
func (m MessageSenderUser) SetGameScore(client *Client, chatId int64, messageId int64, editMessage bool, score int32, force bool) (*Message, error) {
	return client.SetGameScore(chatId, messageId, editMessage, m.UserId, score, force)
}

// SetInlineGameScore is a helper method for Client.SetInlineGameScore
func (m MessageSenderUser) SetInlineGameScore(client *Client, inlineMessageId string, editMessage bool, score int32, force bool) (*Ok, error) {
	return client.SetInlineGameScore(inlineMessageId, editMessage, m.UserId, score, force)
}

// SetMenuButton is a helper method for Client.SetMenuButton
func (m MessageSenderUser) SetMenuButton(client *Client, menuButton *BotMenuButton) (*Ok, error) {
	return client.SetMenuButton(m.UserId, menuButton)
}

// SetPassportElementErrors is a helper method for Client.SetPassportElementErrors
func (m MessageSenderUser) SetPassportElementErrors(client *Client, errors []*InputPassportElementError) (*Ok, error) {
	return client.SetPassportElementErrors(m.UserId, errors)
}

// SetStickerSetThumbnail is a helper method for Client.SetStickerSetThumbnail
func (m MessageSenderUser) SetStickerSetThumbnail(client *Client, name string, opts *SetStickerSetThumbnailOpts) (*Ok, error) {
	return client.SetStickerSetThumbnail(m.UserId, name, opts)
}

// SetUserEmojiStatus is a helper method for Client.SetUserEmojiStatus
func (m MessageSenderUser) SetUserEmojiStatus(client *Client, emojiStatus *EmojiStatus) (*Ok, error) {
	return client.SetUserEmojiStatus(m.UserId, emojiStatus)
}

// SetUserNote is a helper method for Client.SetUserNote
func (m MessageSenderUser) SetUserNote(client *Client, note *FormattedText) (*Ok, error) {
	return client.SetUserNote(m.UserId, note)
}

// SetUserPersonalProfilePhoto is a helper method for Client.SetUserPersonalProfilePhoto
func (m MessageSenderUser) SetUserPersonalProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SetUserPersonalProfilePhoto(m.UserId, photo)
}

// SetUserSupportInfo is a helper method for Client.SetUserSupportInfo
func (m MessageSenderUser) SetUserSupportInfo(client *Client, message *FormattedText) (*UserSupportInfo, error) {
	return client.SetUserSupportInfo(m.UserId, message)
}

// SharePhoneNumber is a helper method for Client.SharePhoneNumber
func (m MessageSenderUser) SharePhoneNumber(client *Client) (*Ok, error) {
	return client.SharePhoneNumber(m.UserId)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (m MessageSenderUser) SuggestUserBirthdate(client *Client, birthdate *Birthdate) (*Ok, error) {
	return client.SuggestUserBirthdate(m.UserId, birthdate)
}

// SuggestUserProfilePhoto is a helper method for Client.SuggestUserProfilePhoto
func (m MessageSenderUser) SuggestUserProfilePhoto(client *Client, photo *InputChatPhoto) (*Ok, error) {
	return client.SuggestUserProfilePhoto(m.UserId, photo)
}

// TransferChatOwnership is a helper method for Client.TransferChatOwnership
func (m MessageSenderUser) TransferChatOwnership(client *Client, chatId int64, password string) (*Ok, error) {
	return client.TransferChatOwnership(chatId, m.UserId, password)
}

// UploadStickerFile is a helper method for Client.UploadStickerFile
func (m MessageSenderUser) UploadStickerFile(client *Client, stickerFormat *StickerFormat, sticker *InputFile) (*File, error) {
	return client.UploadStickerFile(m.UserId, stickerFormat, sticker)
}

// TestReturnError is a helper method for Client.TestReturnError
func (m MessageSendingStateFailed) TestReturnError(client *Client) (*Error, error) {
	return client.TestReturnError(m.Error)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageSendingStatePending) SendQuickReplyShortcutMessages(client *Client, chatId int64, shortcutId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, shortcutId, m.SendingId)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageSendOptions) AddLocalMessage(client *Client, chatId int64, senderId *MessageSender, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(chatId, senderId, m.DisableNotification, inputMessageContent, opts)
}

// GetMessageEffect is a helper method for Client.GetMessageEffect
func (m MessageSendOptions) GetMessageEffect(client *Client) (*MessageEffect, error) {
	return client.GetMessageEffect(m.EffectId)
}

// PinChatMessage is a helper method for Client.PinChatMessage
func (m MessageSendOptions) PinChatMessage(client *Client, chatId int64, messageId int64, onlyForSelf bool) (*Ok, error) {
	return client.PinChatMessage(chatId, messageId, m.DisableNotification, onlyForSelf)
}

// PostStory is a helper method for Client.PostStory
func (m MessageSendOptions) PostStory(client *Client, chatId int64, content *InputStoryContent, privacySettings *StoryPrivacySettings, albumIds []int32, activePeriod int32, isPostedToChatPage bool, opts *PostStoryOpts) (*Story, error) {
	return client.PostStory(chatId, content, privacySettings, albumIds, activePeriod, isPostedToChatPage, m.ProtectContent, opts)
}

// RegisterUser is a helper method for Client.RegisterUser
func (m MessageSendOptions) RegisterUser(client *Client, firstName string, lastName string) (*Ok, error) {
	return client.RegisterUser(firstName, lastName, m.DisableNotification)
}

// ResendMessages is a helper method for Client.ResendMessages
func (m MessageSendOptions) ResendMessages(client *Client, chatId int64, messageIds []int64, opts *ResendMessagesOpts) (*Messages, error) {
	return client.ResendMessages(chatId, messageIds, m.PaidMessageStarCount, opts)
}

// SendBusinessMessage is a helper method for Client.SendBusinessMessage
func (m MessageSendOptions) SendBusinessMessage(client *Client, businessConnectionId string, chatId int64, inputMessageContent *InputMessageContent, opts *SendBusinessMessageOpts) (*BusinessMessage, error) {
	return client.SendBusinessMessage(businessConnectionId, chatId, m.DisableNotification, m.ProtectContent, m.EffectId, inputMessageContent, opts)
}

// SendBusinessMessageAlbum is a helper method for Client.SendBusinessMessageAlbum
func (m MessageSendOptions) SendBusinessMessageAlbum(client *Client, businessConnectionId string, chatId int64, inputMessageContents []*InputMessageContent, opts *SendBusinessMessageAlbumOpts) (*BusinessMessages, error) {
	return client.SendBusinessMessageAlbum(businessConnectionId, chatId, m.DisableNotification, m.ProtectContent, m.EffectId, inputMessageContents, opts)
}

// SendGiftPurchaseOffer is a helper method for Client.SendGiftPurchaseOffer
func (m MessageSendOptions) SendGiftPurchaseOffer(client *Client, ownerId *MessageSender, giftName string, price *GiftResalePrice, duration int32) (*Ok, error) {
	return client.SendGiftPurchaseOffer(ownerId, giftName, price, duration, m.PaidMessageStarCount)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageSendOptions) SendGroupCallMessage(client *Client, groupCallId int32, text *FormattedText) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, text, m.PaidMessageStarCount)
}

// SendQuickReplyShortcutMessages is a helper method for Client.SendQuickReplyShortcutMessages
func (m MessageSendOptions) SendQuickReplyShortcutMessages(client *Client, chatId int64, shortcutId int32) (*Messages, error) {
	return client.SendQuickReplyShortcutMessages(chatId, shortcutId, m.SendingId)
}

// SetChatDirectMessagesGroup is a helper method for Client.SetChatDirectMessagesGroup
func (m MessageSendOptions) SetChatDirectMessagesGroup(client *Client, chatId int64, isEnabled bool) (*Ok, error) {
	return client.SetChatDirectMessagesGroup(chatId, isEnabled, m.PaidMessageStarCount)
}

// SetChatPaidMessageStarCount is a helper method for Client.SetChatPaidMessageStarCount
func (m MessageSendOptions) SetChatPaidMessageStarCount(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatPaidMessageStarCount(chatId, m.PaidMessageStarCount)
}

// SetGroupCallPaidMessageStarCount is a helper method for Client.SetGroupCallPaidMessageStarCount
func (m MessageSendOptions) SetGroupCallPaidMessageStarCount(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetGroupCallPaidMessageStarCount(groupCallId, m.PaidMessageStarCount)
}

// StartLiveStory is a helper method for Client.StartLiveStory
func (m MessageSendOptions) StartLiveStory(client *Client, chatId int64, privacySettings *StoryPrivacySettings, isRtmpStream bool, enableMessages bool) (*StartLiveStoryResult, error) {
	return client.StartLiveStory(chatId, privacySettings, m.ProtectContent, isRtmpStream, enableMessages, m.PaidMessageStarCount)
}

// CloseStory is a helper method for Client.CloseStory
func (m MessageStory) CloseStory(client *Client) (*Ok, error) {
	return client.CloseStory(m.StoryPosterChatId, m.StoryId)
}

// CreateStoryAlbum is a helper method for Client.CreateStoryAlbum
func (m MessageStory) CreateStoryAlbum(client *Client, name string, storyIds []int32) (*StoryAlbum, error) {
	return client.CreateStoryAlbum(m.StoryPosterChatId, name, storyIds)
}

// DeleteBusinessStory is a helper method for Client.DeleteBusinessStory
func (m MessageStory) DeleteBusinessStory(client *Client, businessConnectionId string) (*Ok, error) {
	return client.DeleteBusinessStory(businessConnectionId, m.StoryId)
}

// DeleteStory is a helper method for Client.DeleteStory
func (m MessageStory) DeleteStory(client *Client) (*Ok, error) {
	return client.DeleteStory(m.StoryPosterChatId, m.StoryId)
}

// EditBusinessStory is a helper method for Client.EditBusinessStory
func (m MessageStory) EditBusinessStory(client *Client, content *InputStoryContent, areas *InputStoryAreas, caption *FormattedText, privacySettings *StoryPrivacySettings) (*Story, error) {
	return client.EditBusinessStory(m.StoryPosterChatId, m.StoryId, content, areas, caption, privacySettings)
}

// EditStory is a helper method for Client.EditStory
func (m MessageStory) EditStory(client *Client, opts *EditStoryOpts) (*Ok, error) {
	return client.EditStory(m.StoryPosterChatId, m.StoryId, opts)
}

// EditStoryCover is a helper method for Client.EditStoryCover
func (m MessageStory) EditStoryCover(client *Client, coverFrameTimestamp float64) (*Ok, error) {
	return client.EditStoryCover(m.StoryPosterChatId, m.StoryId, coverFrameTimestamp)
}

// GetChatStoryInteractions is a helper method for Client.GetChatStoryInteractions
func (m MessageStory) GetChatStoryInteractions(client *Client, preferForwards bool, offset string, limit int32, opts *GetChatStoryInteractionsOpts) (*StoryInteractions, error) {
	return client.GetChatStoryInteractions(m.StoryPosterChatId, m.StoryId, preferForwards, offset, limit, opts)
}

// GetStory is a helper method for Client.GetStory
func (m MessageStory) GetStory(client *Client, onlyLocal bool) (*Story, error) {
	return client.GetStory(m.StoryPosterChatId, m.StoryId, onlyLocal)
}

// GetStoryInteractions is a helper method for Client.GetStoryInteractions
func (m MessageStory) GetStoryInteractions(client *Client, query string, onlyContacts bool, preferForwards bool, preferWithReaction bool, offset string, limit int32) (*StoryInteractions, error) {
	return client.GetStoryInteractions(m.StoryId, query, onlyContacts, preferForwards, preferWithReaction, offset, limit)
}

// GetStoryPublicForwards is a helper method for Client.GetStoryPublicForwards
func (m MessageStory) GetStoryPublicForwards(client *Client, offset string, limit int32) (*PublicForwards, error) {
	return client.GetStoryPublicForwards(m.StoryPosterChatId, m.StoryId, offset, limit)
}

// GetStoryStatistics is a helper method for Client.GetStoryStatistics
func (m MessageStory) GetStoryStatistics(client *Client, chatId int64, isDark bool) (*StoryStatistics, error) {
	return client.GetStoryStatistics(chatId, m.StoryId, isDark)
}

// OpenStory is a helper method for Client.OpenStory
func (m MessageStory) OpenStory(client *Client) (*Ok, error) {
	return client.OpenStory(m.StoryPosterChatId, m.StoryId)
}

// ReportStory is a helper method for Client.ReportStory
func (m MessageStory) ReportStory(client *Client, optionId string, text string) (*ReportStoryResult, error) {
	return client.ReportStory(m.StoryPosterChatId, m.StoryId, optionId, text)
}

// SearchPublicStoriesByTag is a helper method for Client.SearchPublicStoriesByTag
func (m MessageStory) SearchPublicStoriesByTag(client *Client, tag string, offset string, limit int32) (*FoundStories, error) {
	return client.SearchPublicStoriesByTag(m.StoryPosterChatId, tag, offset, limit)
}

// SetStoryPrivacySettings is a helper method for Client.SetStoryPrivacySettings
func (m MessageStory) SetStoryPrivacySettings(client *Client, privacySettings *StoryPrivacySettings) (*Ok, error) {
	return client.SetStoryPrivacySettings(m.StoryId, privacySettings)
}

// SetStoryReaction is a helper method for Client.SetStoryReaction
func (m MessageStory) SetStoryReaction(client *Client, updateRecentReactions bool, opts *SetStoryReactionOpts) (*Ok, error) {
	return client.SetStoryReaction(m.StoryPosterChatId, m.StoryId, updateRecentReactions, opts)
}

// ToggleStoryIsPostedToChatPage is a helper method for Client.ToggleStoryIsPostedToChatPage
func (m MessageStory) ToggleStoryIsPostedToChatPage(client *Client, isPostedToChatPage bool) (*Ok, error) {
	return client.ToggleStoryIsPostedToChatPage(m.StoryPosterChatId, m.StoryId, isPostedToChatPage)
}

// SetBirthdate is a helper method for Client.SetBirthdate
func (m MessageSuggestBirthdate) SetBirthdate(client *Client) (*Ok, error) {
	return client.SetBirthdate(m.Birthdate)
}

// SuggestUserBirthdate is a helper method for Client.SuggestUserBirthdate
func (m MessageSuggestBirthdate) SuggestUserBirthdate(client *Client, userId int64) (*Ok, error) {
	return client.SuggestUserBirthdate(userId, m.Birthdate)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageSuggestedPostApproved) ApproveSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.ApproveSuggestedPost(chatId, messageId, m.SendDate)
}

// DeclineSuggestedPost is a helper method for Client.DeclineSuggestedPost
func (m MessageSuggestedPostDeclined) DeclineSuggestedPost(client *Client, chatId int64, messageId int64) (*Ok, error) {
	return client.DeclineSuggestedPost(chatId, messageId, m.Comment)
}

// SendCallRating is a helper method for Client.SendCallRating
func (m MessageSuggestedPostDeclined) SendCallRating(client *Client, callId int32, rating int32, problems []*CallProblem) (*Ok, error) {
	return client.SendCallRating(callId, rating, m.Comment, problems)
}

// CreateNewBasicGroupChat is a helper method for Client.CreateNewBasicGroupChat
func (m MessageSupergroupChatCreate) CreateNewBasicGroupChat(client *Client, userIds []int64, messageAutoDeleteTime int32) (*CreatedBasicGroupChat, error) {
	return client.CreateNewBasicGroupChat(userIds, m.Title, messageAutoDeleteTime)
}

// CreateNewStickerSet is a helper method for Client.CreateNewStickerSet
func (m MessageSupergroupChatCreate) CreateNewStickerSet(client *Client, userId int64, name string, stickerType *StickerType, needsRepainting bool, stickers []*InputSticker, source string) (*StickerSet, error) {
	return client.CreateNewStickerSet(userId, m.Title, name, stickerType, needsRepainting, stickers, source)
}

// CreateNewSupergroupChat is a helper method for Client.CreateNewSupergroupChat
func (m MessageSupergroupChatCreate) CreateNewSupergroupChat(client *Client, isForum bool, isChannel bool, description string, messageAutoDeleteTime int32, forImport bool, opts *CreateNewSupergroupChatOpts) (*Chat, error) {
	return client.CreateNewSupergroupChat(m.Title, isForum, isChannel, description, messageAutoDeleteTime, forImport, opts)
}

// CreateVideoChat is a helper method for Client.CreateVideoChat
func (m MessageSupergroupChatCreate) CreateVideoChat(client *Client, chatId int64, startDate int32, isRtmpStream bool) (*GroupCallId, error) {
	return client.CreateVideoChat(chatId, m.Title, startDate, isRtmpStream)
}

// GetSuggestedStickerSetName is a helper method for Client.GetSuggestedStickerSetName
func (m MessageSupergroupChatCreate) GetSuggestedStickerSetName(client *Client) (*Text, error) {
	return client.GetSuggestedStickerSetName(m.Title)
}

// SetChatTitle is a helper method for Client.SetChatTitle
func (m MessageSupergroupChatCreate) SetChatTitle(client *Client, chatId int64) (*Ok, error) {
	return client.SetChatTitle(chatId, m.Title)
}

// SetStickerSetTitle is a helper method for Client.SetStickerSetTitle
func (m MessageSupergroupChatCreate) SetStickerSetTitle(client *Client, name string) (*Ok, error) {
	return client.SetStickerSetTitle(name, m.Title)
}

// SetVideoChatTitle is a helper method for Client.SetVideoChatTitle
func (m MessageSupergroupChatCreate) SetVideoChatTitle(client *Client, groupCallId int32) (*Ok, error) {
	return client.SetVideoChatTitle(groupCallId, m.Title)
}

// StartGroupCallRecording is a helper method for Client.StartGroupCallRecording
func (m MessageSupergroupChatCreate) StartGroupCallRecording(client *Client, groupCallId int32, recordVideo bool, usePortraitOrientation bool) (*Ok, error) {
	return client.StartGroupCallRecording(groupCallId, m.Title, recordVideo, usePortraitOrientation)
}

// GetLinkPreview is a helper method for Client.GetLinkPreview
func (m MessageText) GetLinkPreview(client *Client, opts *GetLinkPreviewOpts) (*LinkPreview, error) {
	return client.GetLinkPreview(m.Text, opts)
}

// GetMarkdownText is a helper method for Client.GetMarkdownText
func (m MessageText) GetMarkdownText(client *Client) (*FormattedText, error) {
	return client.GetMarkdownText(m.Text)
}

// GiftPremiumWithStars is a helper method for Client.GiftPremiumWithStars
func (m MessageText) GiftPremiumWithStars(client *Client, userId int64, starCount int64, monthCount int32) (*Ok, error) {
	return client.GiftPremiumWithStars(userId, starCount, monthCount, m.Text)
}

// ParseMarkdown is a helper method for Client.ParseMarkdown
func (m MessageText) ParseMarkdown(client *Client) (*FormattedText, error) {
	return client.ParseMarkdown(m.Text)
}

// PlaceGiftAuctionBid is a helper method for Client.PlaceGiftAuctionBid
func (m MessageText) PlaceGiftAuctionBid(client *Client, giftId string, starCount int64, userId int64, isPrivate bool) (*Ok, error) {
	return client.PlaceGiftAuctionBid(giftId, starCount, userId, m.Text, isPrivate)
}

// SearchQuote is a helper method for Client.SearchQuote
func (m MessageText) SearchQuote(client *Client, quote *FormattedText, quotePosition int32) (*FoundPosition, error) {
	return client.SearchQuote(m.Text, quote, quotePosition)
}

// SendGift is a helper method for Client.SendGift
func (m MessageText) SendGift(client *Client, giftId string, ownerId *MessageSender, isPrivate bool, payForUpgrade bool) (*Ok, error) {
	return client.SendGift(giftId, ownerId, m.Text, isPrivate, payForUpgrade)
}

// SendGroupCallMessage is a helper method for Client.SendGroupCallMessage
func (m MessageText) SendGroupCallMessage(client *Client, groupCallId int32, paidMessageStarCount int64) (*Ok, error) {
	return client.SendGroupCallMessage(groupCallId, m.Text, paidMessageStarCount)
}

// SendTextMessageDraft is a helper method for Client.SendTextMessageDraft
func (m MessageText) SendTextMessageDraft(client *Client, chatId int64, forumTopicId int32, draftId string) (*Ok, error) {
	return client.SendTextMessageDraft(chatId, forumTopicId, draftId, m.Text)
}

// TranslateText is a helper method for Client.TranslateText
func (m MessageText) TranslateText(client *Client, toLanguageCode string) (*FormattedText, error) {
	return client.TranslateText(m.Text, toLanguageCode)
}

// AddChatMember is a helper method for Client.AddChatMember
func (m MessageThreadInfo) AddChatMember(client *Client, userId int64, forwardLimit int32) (*FailedToAddMembers, error) {
	return client.AddChatMember(m.ChatId, userId, forwardLimit)
}

// AddChatMembers is a helper method for Client.AddChatMembers
func (m MessageThreadInfo) AddChatMembers(client *Client, userIds []int64) (*FailedToAddMembers, error) {
	return client.AddChatMembers(m.ChatId, userIds)
}

// AddChatToList is a helper method for Client.AddChatToList
func (m MessageThreadInfo) AddChatToList(client *Client, chatList *ChatList) (*Ok, error) {
	return client.AddChatToList(m.ChatId, chatList)
}

// AddChecklistTasks is a helper method for Client.AddChecklistTasks
func (m MessageThreadInfo) AddChecklistTasks(client *Client, messageId int64, tasks []*InputChecklistTask) (*Ok, error) {
	return client.AddChecklistTasks(m.ChatId, messageId, tasks)
}

// AddFileToDownloads is a helper method for Client.AddFileToDownloads
func (m MessageThreadInfo) AddFileToDownloads(client *Client, fileId int32, messageId int64, priority int32) (*File, error) {
	return client.AddFileToDownloads(fileId, m.ChatId, messageId, priority)
}

// AddLocalMessage is a helper method for Client.AddLocalMessage
func (m MessageThreadInfo) AddLocalMessage(client *Client, senderId *MessageSender, disableNotification bool, inputMessageContent *InputMessageContent, opts *AddLocalMessageOpts) (*Message, error) {
	return client.AddLocalMessage(m.ChatId, senderId, disableNotification, inputMessageContent, opts)
}

// AddMessageReaction is a helper method for Client.AddMessageReaction
func (m MessageThreadInfo) AddMessageReaction(client *Client, messageId int64, reactionType *ReactionType, isBig bool, updateRecentReactions bool) (*Ok, error) {
	return client.AddMessageReaction(m.ChatId, messageId, reactionType, isBig, updateRecentReactions)
}

// AddOffer is a helper method for Client.AddOffer
func (m MessageThreadInfo) AddOffer(client *Client, messageId int64, options *MessageSendOptions) (*Message, error) {
	return client.AddOffer(m.ChatId, messageId, options)
}

// AddPendingPaidMessageReaction is a helper method for Client.AddPendingPaidMessageReaction
func (m MessageThreadInfo) AddPendingPaidMessageReaction(client *Client, messageId int64, starCount int64, opts *AddPendingPaidMessageReactionOpts) (*Ok, error) {
	return client.AddPendingPaidMessageReaction(m.ChatId, messageId, starCount, opts)
}

// AddRecentlyFoundChat is a helper method for Client.AddRecentlyFoundChat
func (m MessageThreadInfo) AddRecentlyFoundChat(client *Client) (*Ok, error) {
	return client.AddRecentlyFoundChat(m.ChatId)
}

// AddStoryAlbumStories is a helper method for Client.AddStoryAlbumStories
func (m MessageThreadInfo) AddStoryAlbumStories(client *Client, storyAlbumId int32, storyIds []int32) (*StoryAlbum, error) {
	return client.AddStoryAlbumStories(m.ChatId, storyAlbumId, storyIds)
}

// ApproveSuggestedPost is a helper method for Client.ApproveSuggestedPost
func (m MessageThreadInfo) ApproveSuggestedPost(client *Client, messageId int64, sendDate int32) (*Ok, error) {
	return client.ApproveSuggestedPost(m.ChatId, messageId, sendDate)
}

// BanChatMember is a helper method for Client.BanChatMember
func (m MessageThreadInfo) BanChatMember(client *Client, memberId *MessageSender, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	return client.BanChatMember(m.ChatId, memberId, bannedUntilDate, revokeMessages)
}

// BoostChat is a helper method for Client.BoostChat
func (m MessageThreadInfo) BoostChat(client *Client, slotIds []int32) (*ChatBoostSlots, error) {
	return client.BoostChat(m.ChatId, slotIds)
}
